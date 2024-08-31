set dotenv-load

run:
  sqlc generate
  go run .

watch:
  watchexec --exts '.go' just run

migrate direction:
  cd "sql/schema" && goose postgres $CONN {{direction}}

