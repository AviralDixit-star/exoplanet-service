# Dockerfile
FROM golang:1.16-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o /exoplanet-service ./cmd/main.go

EXPOSE 8080

# Command to run the executable
CMD [ "/exoplanet-service" ]
