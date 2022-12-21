FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o auth-service

EXPOSE 8082

CMD [ "./auth-service" ]
