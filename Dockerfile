FROM golang:1.19.1

RUN go version
ENV GOPATH=/

COPY . /todo-app
WORKDIR /todo-app

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o /app /cmd/main.go

CMD ["./app"]
