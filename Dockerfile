FROM golang:1.23-alpine AS build

WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /catalog ./cmd/catalog

FROM alpine:3.20

RUN addgroup -S app && adduser -S app -G app
COPY --from=build /catalog /catalog
USER app
EXPOSE 8080
HEALTHCHECK --interval=5s --timeout=2s --retries=5 CMD wget -qO- http://127.0.0.1:8080/health || exit 1
ENTRYPOINT ["/catalog"]
