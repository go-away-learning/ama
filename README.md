<!-- gen-readme start - generated by https://github.com/jetify-com/devbox/ -->
## Getting Started
This project uses [devbox](https://github.com/jetify-com/devbox) to manage its development environment.

Install devbox:
```sh
curl -fsSL https://get.jetpack.io/devbox | bash
```

Start the devbox shell:
```sh 
devbox shell
```

Run a script in the devbox environment:
```sh
devbox run <script>
```
## Scripts
Scripts are custom commands that can be run using this project's environment. This project has the following scripts:

* [generate_code_from_sql](#devbox-run-generate_code_from_sql)
* [migrations](#devbox-run-migrations)
* [start_db](#devbox-run-start_db)
* [stop_db](#devbox-run-stop_db)

## Environment

```sh
GOENV="off"
```

## Shell Init Hook
The Shell Init Hook is a script that runs whenever the devbox environment is instantiated. It runs 
on `devbox shell` and on `devbox run`.
```sh
cp -n .env.example .env 2>/dev/null
```

## Packages

* [sqlc@1.26.0](https://www.nixhub.io/packages/sqlc)
* [atlas@0.25.0](https://www.nixhub.io/packages/atlas)
* [dotenv-cli@7.4.3](https://www.nixhub.io/packages/dotenv-cli)
* [go@1.22.5](https://www.nixhub.io/packages/go)

## Script Details

### devbox run generate_code_from_sql
```sh
sqlc generate --file database/sqlc.yaml
```
&ensp;

### devbox run migrations
```sh
dotenv -- atlas schema apply --config "file://database/atlas.hcl" --env local
```
&ensp;

### devbox run start_db
```sh
docker-compose -p ama -f database/docker-compose.yaml --env-file .env up -d
```
&ensp;

### devbox run stop_db
```sh
docker-compose -p ama -f database/docker-compose.yaml --env-file .env down
```
&ensp;



<!-- gen-readme end -->
