FROM golang:latest

ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update

# build go app
RUN go build -o convapi ./cmd/main.go

CMD ["./convapi"]