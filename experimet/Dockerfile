FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go build -o /go/src/app logger.go
#RUN go get -d -v ./...
#RUN go install -v ./...

CMD ["app"]

# docker build -t logger:v1 .
# docker run logger:v1