# Poke Metrics
![Cover](./cover.png "Cover" )

This project was created with the purpose of demonstrating metric generation and tracing in different scenarios: API, gRPC, and Messaging.

To simplify the business domain, the use cases are reduced to:

- Create a Pokémon Trainer

- Retrieve Trainer Information

- Go on a Hunt

## Trainer Creation:
Each trainer has a Name and a Favorite Pokémon Type.

## Retrieve Trainer Information:
Once created, the trainer receives an identifier. With this identifier, we can retrieve their information: Name, Favorite Pokémon Type, and Captured Pokémon.

## Hunt:
When going on a hunt, the trainer will have X capture attempts.
If the encountered Pokémon matches the trainer's selected Favorite Pokémon Type, the capture chance is **100%**; otherwise, it is **50%**.


# Running
With docker/docker-compose installed run the following command 

```
docker-compose up 
```

# API

#### Create Trainer

<details>
 <summary><code>POST</code> <code><b>/trainer</b></code> <code>(Create a new trainer)</code></summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | name      |  required | string                  | Trainer's name                                                        |
> | favorite_pokemon_type | optional                | Trainer's Favorite Pokemon type | When hunting favorite types, there is a 100% chance of successful capture |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`                | `{"id": uuid, "name": string, "favorite_pokemon_type": string, pokemons []Pokemon {"id": int32, "name": string, "pokemon-types": []string} }`                                      |
> | `400`         | `application/json`                | `{"message":"Bad Request"}`                            |
> | `500`         | `application/json`                | ``
</details>

### Get Trainer information
<details>
 <summary><code>GET</code> <code><b>/trainer/:trainer-id</b></code> <code>(Get trainer information)</code></summary>

