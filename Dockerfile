FROM golang:1.23 AS build-stage
WORKDIR /app
COPY *.go ./
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /server /
EXPOSE 9090
ENTRYPOINT ["/server"]
