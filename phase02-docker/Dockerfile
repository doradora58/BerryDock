FROM golang:1.26-alpine AS build
WORKDIR /app
COPY phase02-docker/go.mod ./
COPY phase01-info-api/main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/info-api

FROM scratch
COPY --from=build /app/info-api /info-api
EXPOSE 8080
ENTRYPOINT ["/info-api"]
