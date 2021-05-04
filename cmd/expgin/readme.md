## Small REST API server

#### PostgreSQL
- docker pull postgres
- docker run --name gopostgres -e POSTGRES_PASSWORD=rowssap -p 5432:5432 -d postgres
- docker exec -it gopostgres psql -U postgres

#### psql
- psql -> pgAdmin -> create DB "offersapp"
- psql offersapp < database/schema.sql

#### GUI 
- HeidiSQL - free / win10
- CREATE DATABASE "godbapi"

#### Request Postman
- Post to localhost:3000/users/register  -> { "user_id": "someid" }

#### Run
- docker start gopostgres 
- run test in second_test.go
