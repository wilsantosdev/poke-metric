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