FROM golang:1.23 AS build

WORKDIR /app

COPY . .

RUN go build -v ./cmd/main/main.go

FROM golang:1.23

WORKDIR /app

COPY --from=build /app/main main

EXPOSE 8080

ENTRYPOINT ["./main"]