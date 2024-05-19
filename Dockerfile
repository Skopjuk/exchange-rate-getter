FROM golang:1.22 AS build
WORKDIR /app
COPY ./ ./
RUN go mod tidy
RUN go build -o /app/app main.go

FROM alpine:latest
COPY --from=build /app/app /app
RUN apk add --no-cache ca-certificates
CMD /app
