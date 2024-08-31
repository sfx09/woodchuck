set dotenv-load

run:
  sqlc generate
  go run .

watch:
  watchexec --exts '.go' just run

migrate direction:
  cd "sql/schema" && goose postgres $CONN {{direction}}

start-db:
  docker run --name postgres-docker -p 5432:5432 -e POSTGRES_DB=$DB -e POSTGRES_PASSWORD=$PASS -d postgres

reset-db:
  docker kill postgres-docker 
  docker rm postgres-docker 

query-db:
  docker exec -it postgres-docker su -c psql postgres
