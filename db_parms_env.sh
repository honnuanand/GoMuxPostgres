export APP_DB_USERNAME=postgres
export APP_DB_PASSWORD=pgpass
export APP_DB_NAME=postgres


docker run -d \
	--name postgres \
	-e POSTGRES_PASSWORD=pgpass \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
    -e POSTGRES_DB=postgres \
    -e "DB_HOST=127.0.0.1" \
	postgres