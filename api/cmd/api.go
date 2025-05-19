package cmd

import (
	"api/internal/domain"
	"api/internal/http/request"
	"api/internal/http/response"
	"api/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	customMiddleware "api/internal/http/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type api struct {
	Port           string
	Router         *chi.Mux
	TrainerService domain.TrainerService
	HuntService    domain.HuntService
	context        context.Context
}

func NewAPI(ctx context.Context, port string) *api {
	return &api{
		Port:           port,
		context:        ctx,
		Router:         chi.NewRouter(),
		TrainerService: service.NewTrainerService(),
		HuntService:    service.NewHuntService(),
	}
}

func (a *api) setupRoutes() {
	a.Router.Post("/trainer", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		createTrainerRequest, err := request.NewCreateTrainerRequest(r.Body)
		if err != nil {
			response, err := response.NewErrorResponse(
				err,
				http.StatusBadRequest,
			).ToJSON()
			if err != nil {
				log.Fatal(err)
				http.Error(w, "Failed to create error response", http.StatusInternalServerError)
				return
			}
			http.Error(w, string(response), http.StatusBadRequest)
			return
		}

		trainer, err := a.TrainerService.CreateTrainer(
			r.Context(),
			createTrainerRequest.Name,
			createTrainerRequest.FavoritePokemonType,
		)

		if err != nil {
			log.Fatal(err)
			http.Error(w, "Failed to create trainer", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.NewCreateTrainerResponse(*trainer))
	})

	a.Router.Get("/trainer/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		trainer, err := a.TrainerService.GetTrainer(r.Context(), id)

		if err != nil {
			json.NewEncoder(w).Encode(response.NewErrorResponse(fmt.Errorf("trainer not found"), http.StatusNotFound))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.NewGetTrainerResponse(*trainer))
	})

	a.Router.Post("/trainer/{id}/hunt", func(w http.ResponseWriter, r *http.Request) {
		trainerId := chi.URLParam(r, "id")

		trainer, err := a.TrainerService.GetTrainer(r.Context(), trainerId)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Failed to get trainer", http.StatusInternalServerError)
			return
		}

		err = a.HuntService.HuntPokemon(r.Context(), *trainer)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Failed to hunt pokemon", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}

func (a *api) setupMiddleware() {
	a.Router.Use(
		middleware.Logger,
		middleware.RequestID,
		customMiddleware.ApplicationJsonMiddleware,
	)
}

func (a *api) Run() error {
	shutdown := service.NewTracerService(a.context)
	defer shutdown()
	a.setupMiddleware()
	a.setupRoutes()
	return http.ListenAndServe(fmt.Sprintf(":%v", a.Port), a.Router)
}
