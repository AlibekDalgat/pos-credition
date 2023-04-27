FROM golang:1.19.1

RUN go version

COPY . /pos_credition/
WORKDIR /pos_credition/

RUN apt-get update && apt-get -y install postgresql-client

RUN go mod download
RUN GOOS=linux go build -o app ./cmd/main.go

RUN sed -i -e 's/\r$//' *.sh
RUN chmod +x wait-for-postgres.sh

CMD ["./app"]