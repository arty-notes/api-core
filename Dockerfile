FROM golang:1.22.2 as build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY ./ .
RUN go mod download

# Build api
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api cmd/api/main.go

# Start from scratch and use a separate build
# using alpine os as the foundation for my container
FROM alpine:latest

WORKDIR /

# Bring the executable from the build
COPY --from=build /bin/api /api

# Expose the port
EXPOSE 4000

# Run the API as the entry point
ENTRYPOINT ["/api"]