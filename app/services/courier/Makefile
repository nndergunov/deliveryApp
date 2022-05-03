migrate -path database/migration -database "postgresql://postgres:postgres@0.0.0.0:5433/courier_db?sslmode=disable" -verbose up

Get-content courier.sql |  psql --username postgres courier_db

createdb --username=postgres --owner=postgres courier_db