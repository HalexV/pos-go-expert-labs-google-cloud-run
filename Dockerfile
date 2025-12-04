FROM golang:1.24.3 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./main.go ./main.go
COPY configs ./configs
COPY dto ./dto

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /webserver .

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /webserver /webserver

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/webserver"]