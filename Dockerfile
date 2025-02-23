FROM golang:1.17

WORKDIR /usr/src/app

COPY ./app ./

RUN go mod download && go mod verify

ENV CGO_ENABLED 0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -v -o /dist

RUN apt update && apt install net-tools && apt install curl

EXPOSE 8080

CMD ["/dist"]