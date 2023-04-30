# bassoon

Bassoon app

## Development: build & run

To run the app locally use this command:

```bash
  $ make run
```

To run app using Docker use this command:

```bash
  $ docker-compose up
```

## Tests

To run tests for the app use this command:

```bash
  $ make test
```
 or to be able to see coverage report:

```bash
  $ make test-cover
```

## Linter

To run linter use this command:

```bash
  $ make lint
```

## Config

App configurable with ENVs and `.env` file:   

```.dotenv
    BASSOON_PORTS_FILEPATH: /bassoon/ports.json
    BASSOON_HTTP_PORT: :8000
```

## Endpoints

Create new entry:

POST /v1/ports

Payload example:

```json
{
    "id": "KWK",
    "name": "Harare",
    "city": "Harare",
    "country": "Zimbabwe",
    "alias": [],
    "regions": [],
    "coordinates": [
      31.03351,
      -17.8251657
    ],
    "province": "Harare",
    "timezone": "Africa/Harare",
    "unlocs": [
      "ZWHRE"
    ]
}
```

Retrieve entry by `id`:

GET /v1/ports/{id}