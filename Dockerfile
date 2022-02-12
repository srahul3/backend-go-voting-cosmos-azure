FROM golang:1.16-alpine

WORKDIR /app

COPY ./app ./

RUN ls -ltrR
RUN go mod download

EXPOSE 8080

CMD ["go","run","main.go" ]
