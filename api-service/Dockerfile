FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o api-service

EXPOSE 8081 

CMD [ "./api-service" ]
