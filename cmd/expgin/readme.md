## Small REST API server

### PostgreSQL
- Скачать образ базы данных:  
  ``` docker pull postgres```
- Запуск с созданием БД gopostgres:   
  ``` docker run --name gopostgres -e POSTGRES_PASSWORD=rowssap -p 5432:5432 -d postgres```
- Запустить утилиту psql:   
  ``` docker exec -it gopostgres psql -U postgres ```
- Если необходимо применить schema.sql:   
  ``` cat ./database/schema.sql | docker exec -i gopostgres psql -U postgres -a ```

#### psql
- psql -> pgAdmin -> create DB "offersapp"
- psql offersapp < database/schema.sql

#### GUI 
- HeidiSQL - free / win10
- CREATE DATABASE "godbapi"

#### Request Postman
- Post to localhost:3000/users/register  -> { "user_id": "someid" }

### Run
- docker start gopostgres 
- run test in server_test.go
