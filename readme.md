# Blog 2020

## Требования
- собственный блог
- несоколько направлений
- возможность сбора и отображения статистики
- админака для написания публикаций
- желательно маркдаун стиль
- возможность комментировать - через соцсети авторизацию
- оформление кода - цвет
- многостраничный переход
- поиск
- тегирование
- возможность экспериментальных проектов - сбор новостей, агрегация и тд
- возможность рассылки в соцсети (телега, twitter)
- система активного мониторинга и информирования в слак
- возможно мобильное приложение 

## Этапы
1. Верстка + отображение одной тематики
2. Статистика посещений
3. Добавить другие тематики
4. Админка - показывать посещения
5. Админка - добавление статей
6. Добавить возможнсть комментировать
7. Удаление комментариев, бан, и тд
8. Добавить работу с чужим API
	https://openweathermap.org/current
	https://sv443.net/jokeapi/v2/
9. можно добавить агригатор интересных новостей
10. Попробовать пушить в телегу и вк
11. Добавить бота
12. Аналитики и статистика посещений


## Технологии
- React - build static
- TS
- Rest API - Go
- PostgreSQL
- Nginx
- Swagger
- Figma под создание интерфейса

## Реализация
1. Начать с простого API на GO
	https://github.com/gin-gonic/gin
2. Работа с базой - получение данных
3. Залить в базу тестовые документы
4. React базовое приложение
5. Мобильная верстка
6. Подготовить виртуальную машину
7. Выкатить
8. Вопрос с аналитикой
9. Логирование
10. Документирование процесса.
	https://github.com/swaggo/gin-swagger
11. Последовательный этап превращение в чистую архитектурв
	Начать пока с простого и двигаться к этому
	https://www.youtube.com/watch?v=dyvYXidvc8g	
12. Прикрутить линтер - как часть процсса
	- Go
	- JS
	- TS

IPTables -> 443 + SSH 
Client -> HTTPS -> NGINX -> Static page React
                         -> /api -> :5000 golang  -> PostgreSQL

## Install
Golang:
	$ go version
	go version go1.15.3 windows/amd64
	$ go get -u github.com/gin-gonic/gin
Linter:
	https://golangci-lint.run/usage/install/
	$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.32.2
	$ golangci-lint --version
	golangci-lint has version 1.32.2 built from 52d26a3 on 2020-11-03T01:15:38Z
	$ golangci-lint run
Test:
	$ go get github.com/onsi/ginkgo/ginkgo
	$ go get github.com/onsi/gomega/...	
	$ go mod vendor
React:
	$ npx create-react-app client
Golang:
	$ cd goblog
	$ go run main.go
	Start on port :5000
Browser:
	http://localhost:3000
GnuWin:
	make for Windows


## Step by step
1. Поставить Golang
2. VSCode
3. Создаем проект и стурктуру папок под Front и Back End
4. В Go весь код в /cmd/goblog
5. Тесты в /tests
6. Берем пример из документации gin вместе с тестом
7. Запускаем 
	$ go run ./cmd/goblog/main.go
8. Работа с модулями 
	$ go mod init goblog
9. Добавить все зависимости в файл go.mod
	$ go mod vendor
10. Делаем Makefile с разделми 
	make run, test, lint
11. Ставимь make for Windows и прописываем в path
12. Запускаем тесты
	$ make test
13. Переделвыаем тесты на https://onsi.github.io/ginkgo/
	$ go get github.com/onsi/ginkgo/ginkgo
	$ go get github.com/onsi/gomega/...
	добавил в тесты сделующие строки

	и запустил 
	$ go mod vendor
	запуск тестов через 
	$ ginkgo ./cmd/goblog/



14. Прикручиваем линтер https://golangci-lint.run/usage/install/
	$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.32.2
	$ golangci-lint --version
		golangci-lint has version 1.32.2 built from 52d26a3 on 2020-11-03T01:15:38Z
15. Запуск линтера и добавление в make
	$ golangci-lint run
	$ make lint	
15. Прикручиваем логирование

16. Прикручиваем Swagger
	Берем библиотеку swag https://github.com/swaggo/swag
	$ go get -u github.com/swaggo/swag/cmd/swag
17. Для взаимодействия с gin необходимо еще 
	$ go get -u github.com/swaggo/gin-swagger	
	$ go get -u github.com/swaggo/files

18.	В main.go
	добавить в проект два импорта:
		ginSwagger "github.com/swaggo/gin-swagger"
		swaggerFiles "github.com/swaggo/files"

	Так же перед фунйией - заголовок:
		// @title Goblog Swagger API
		// @version 1.0
		// @description Swagger API for Golang Project Goblog.
		// @termsOfService http://swagger.io/terms/

		// @contact.name API Support
		// @contact.email vvuri1978@gmail.com

		// @license.name MIT

		// @BasePath /api/v1

		// @securityDefinitions.apikey ApiKeyAuth
		// @in header
		// @name Authorization

	$ go mod vendor
	
18. Формируем структуру
	$ cd cmd/goblog/
	$ swag init
	появился каталог /docs

	$ swag init .../cmd/goblog

19. Доступ к документации swagger
	 http://localhost:5000/swagger/index.html	
	 swag init .../cmd/goblog







-. React проект запускаем
	$ npx create-react-app client
-. Прописываем в package.json 
	"proxy": "http://localhost:5000"	
-. Запускаем	
	$ cd client
	$ npm start
-. Открываем http://localhost:3000
-.
-.


## Make
- install   Install missing dependencies. Runs `go get` internally.
- start     Start in development mode. Auto-starts when code changes.
- stop      Stop development mode.
- compile   Compile the binary.
- watch     Run given command when code changes. e.g; make watch run="go test ./..."
- exec      Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
- clean     Clean build files. Runs `go clean` internally.


## Погнали
	База тут https://habr.com/ru/post/455302/


	gorilla / mux — мощный URL-маршрутизатор и диспетчер. Мы будем использовать этот пакет для сопоставления путей URL с их обработчиками.
	jinzhu / gorm — отличная ORM-библиотека для Golang. Мы используем этот пакет, чтобы взаимодействовать с базой данных.
	dgrijalva / jwt-go — используется для подписи и проверки токенов JWT.
	joho / godotenv — используется для загрузки файлов .env в проект.

	JWT JSON Web Tokens — это открытый стандарт RFC 7519 для создания токенов доступа. Используется в передаче данных для аутентификации в клиент-серверных приложениях. В обычных веб-приложениях легко идентифицировать пользователей с помощью сессий, однако, когда API вашего веб-приложения взаимодействует, скажем, с клиентом Android или IOS, сессии становятся малопригодными для использования. С помощью JWT мы можем создать уникальный токен для каждого аутентифицированного пользователя. Этот токен будет включён в заголовок последующего запроса к API. 

	Настройка HTTPS NGINX
		https://nginx.org/ru/docs/http/configuring_https_servers.html
		+ SSL сертификат - но там чисто на сервер зайти и получить

	go modules позволяет в том числе и вендорить зависимости.


	Routing

		направление / blog / {id}

		- devops
		- java
		- qa

	Разделы
		- Blog	
		- About
		- Experiment

	API
		/api
		/api/post         - все заголовки постов + возможность перехода
			{	"count": 36, 
    			"next": "http://swapi.dev/api/starships/?page=2", 
    			"previous": null, 
    			"results": [
        			{...

		/api/post/{:id}   - конкретную статью
			- тут прилетает { с полным постом }
		/api/post/?search=devops
			- 
		/api/post/?tag=devops
		/api/tags/		  - списбок тегов
			

	DB
		post 
		anotation
		tag

=====================================================================================================
	^admin/
	^$
	^documentation$
	^about$
	^stats$
	^stripe/donation
	^api/people/schema$
	^api/planets/schema$
	^api/films/schema$
	^api/species/schema$
	^api/vehicles/schema$
	^api/starships/schema$
	^api/ ^$ [name='api-root']
	^api/ ^\.(?P<format>[a-z0-9]+)$ [name='api-root']
	^api/ ^people/$ [name='people-list']
	^api/ ^people/\.(?P<format>[a-z0-9]+)$ [name='people-list']
	^api/ ^people/(?P<pk>[^/.]+)/$ [name='people-detail']
	^api/ ^people/(?P<pk>[^/.]+)/\.(?P<format>[a-z0-9]+)$ [name='people-detail']
	^api/ ^planets/$ [name='planet-list']
	^api/ ^planets/\.(?P<format>[a-z0-9]+)$ [name='planet-list']
	^api/ ^planets/(?P<pk>[^/.]+)/$ [name='planet-detail']
	^api/ ^planets/(?P<pk>[^/.]+)/\.(?P<format>[a-z0-9]+)$ [name='planet-detail']
	^api/ ^films/$ [name='film-list']
	^api/ ^films/\.(?P<format>[a-z0-9]+)$ [name='film-list']
	^api/ ^films/(?P<pk>[^/.]+)/$ [name='film-detail']
	^api/ ^films/(?P<pk>[^/.]+)/\.(?P<format>[a-z0-9]+)$ [name='film-detail']
	^api/ ^species/$ [name='species-list']
	^api/ ^species/\.(?P<format>[a-z0-9]+)$ [name='species-list']
	^api/ ^species/(?P<pk>[^/.]+)/$ [name='species-detail']
	^api/ ^species/(?P<pk>[^/.]+)/\.(?P<format>[a-z0-9]+)$ [name='species-detail']
	^api/ ^vehicles/$ [name='vehicle-list']
	^api/ ^vehicles/\.(?P<format>[a-z0-9]+)$ [name='vehicle-list']
	^api/ ^vehicles/(?P<pk>[^/.]+)/$ [name='vehicle-detail']
	^api/ ^vehicles/(?P<pk>[^/.]+)/\.(?P<format>[a-z0-9]+)$ [name='vehicle-detail']
	^api/ ^starships/$ [name='starship-list']
	^api/ ^starships/\.(?P<format>[a-z0-9]+)$ [name='starship-list']
	^api/ ^starships/(?P<pk>[^/.]+)/$ [name='starship-detail']
	^api/ ^starships/(?P<pk>[^/.]+)/\.(?P<format>[a-z0-9]+)$ [name='starship-detail']

### ToDo List

- базовые запросы
- с ID
- с поиском по тегу

- gin.BasicAuth
- by tocken временный на ...

- CORS

- загрузка файлов 
    https://zhashkevych.medium.com/api-%D0%BD%D0%B0-go-gin-gonic-%D0%B4%D0%BB%D1%8F-%D0%B7%D0%B0%D0%B3%D1%80%D1%83%D0%B7%D0%BA%D0%B8-%D1%84%D0%B0%D0%B9%D0%BB%D0%BE%D0%B2-99e3622a3d6a

- рпботы с БД

- Add React Static
- Собрать все в контейнере для тестов

- DNS in hosts - или есть другие варианты

### Test framework
	$ go get -u github.com/stretchr/testify/assert - чисто assert
	or
	$ go get github.com/smartystreets/goconvey - отличный UI + как сервер отслеживает локально + BDD
	$ goconvey  + HTTP://127.0.0.1:8080
	or
	https://onsi.github.io/ginkgo/ - этот в стиле Jest




### Doc
- Request -> Route Parser -> [Optional Middleware] -> Route Handler -> [Optional Middleware] -> Response
- Middleware часть кода, запускается в любой части, например авторизация


### Lib
- Goose - для миграций 
    $ go get bitbucket.org/liamstask/goose/cmd/goose
- package for email validation.
    $ go get github.com/badoux/checkmail
- GORM - ORM library
    $ go get github.com/jinzhu/gorm
- Provos and Mazières's bcrypt adaptive hashing algorithm
    $ go get golang.org/x/crypto/bcrypt
- implementation of JSON Web Tokens    
    $ go get github.com/dgrijalva/jwt-go
- ???    
    $ go get github.com/jinzhu/gorm/dialects/postgres
- Ggolang port of the Ruby dotenv project    
    $ go get github.com/joho/godotenv
- basic assertion functions for testing
    $ go get gopkg.in/go-playground/assert.v1
- Gin middleware/handler to enable CORS support.    
    $ go get github.com/gin-contrib/cors 
- ??
    $ go get github.com/gin-gonic/contrib
- official AWS SDK
go get github.com/aws/aws-sdk-go 
go get github.com/sendgrid/sendgrid-go
go get github.com/stretchr/testify
go get github.com/twinj/uuid
github.com/matcornic/hermes/v2



### Решения

- CORS
    Существует также официальное промежуточное программное обеспечение gin для обработки CORS запросов github.com/gin-contrib/cors .

    Вы можете установить его с помощью $ go get github.com/gin-contrib/cors , а затем добавить это промежуточное программное обеспечение в свое приложение следующим образом: пакет основной

    import (
        "time"

        "github.com/gin-contrib/cors"
        "github.com/gin-gonic/gin"
    )

    func main() {
        router := gin.Default()
        // CORS for https://foo.com and https://github.com origins, allowing:
        // - PUT and PATCH methods
        // - Origin header
        // - Credentials share
        // - Preflight requests cached for 12 hours
        router.Use(cors.New(cors.Config{
            AllowOrigins:     []string{"https://foo.com"},
            AllowMethods:     []string{"PUT", "PATCH"},
            AllowHeaders:     []string{"Origin"},
            ExposeHeaders:    []string{"Content-Length"},
            AllowCredentials: true,
            AllowOriginFunc: func(origin string) bool {
                return origin == "https://github.com"
            },
            MaxAge: 12 * time.Hour,
        }))
        router.Run()
    }

## Links 
	project-layout - https://github.com/golang-standards/project-layout

