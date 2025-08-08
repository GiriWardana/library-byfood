FROM golang:1.23-alpine AS backend

WORKDIR /app

# Add timezone database
RUN apk add --no-cache tzdata

# Copy go mod and download deps
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the app
RUN go build -o /app/app

EXPOSE 8080
CMD ["/app/app"]
