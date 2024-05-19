FROM golang:1.22 AS build
WORKDIR /app
COPY ./ ./
RUN go mod tidy
RUN go build -o /app/app main.go

FROM alpine:latest
COPY ./schema ./
COPY --from=build /app/app ./
RUN apk add --no-cache ca-certificates migrate
CMD migrate -path ./schema -database 'postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable' up && ./app
