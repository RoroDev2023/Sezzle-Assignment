FROM golang:1.22-alpine AS build

WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN go test ./... && go build -o /calculator-api ./cmd/server

FROM alpine:3.20

RUN adduser -D appuser
USER appuser
COPY --from=build /calculator-api /calculator-api
EXPOSE 8080
CMD ["/calculator-api"]

