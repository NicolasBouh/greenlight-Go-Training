migrate create -seq -ext=.sql -dir=./migrations create_movies_table

migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up

migrate -path=./migrations -database=$EXAMPLE_DSN force 1

 