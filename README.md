# HereWeGo

Messing around with web dev, docker, devcontainers, sveltekit, postgres, and more.

## Go

### [Go-Blueprint](https://github.com/Melkeydev/go-blueprint)

Setting up the go project using Go-Blueprint. Goblueprint is a cli tool that 
allows a user to spin up a go project.

## Svelte

Set up using 

```bash
bun create svelte@latest .
```

## Database

Connecting to the dev container database 

``` bash
docker exec -it herewego_devcontainer-db-1 bash
```

Running `pqsl`

```bash
psql -U $POSTGRES_USER -d $POSTGRES_DB
```

The `POSTGRES_USER` and `POSTGRES_DB` will be created when the container is 
created in the environment variables section.


## TODO

On container set up

* Install air for watch `go install github.com/cosmtrek/air@latest`
* Install go migrate `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
* Install echo-swagger `go install github.com/swaggo/swag/cmd/swag@latest`

