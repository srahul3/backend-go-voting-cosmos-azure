FROM golang:1.17

WORKDIR /usr/src/app

COPY ./app ./

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 8080

CMD ["app"]
