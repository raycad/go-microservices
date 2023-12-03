FROM golang:1.20-alpine AS builder

# Start build process
RUN mkdir -p /usr/local/go/src/movie-microservice
WORKDIR /usr/local/go/src/movie-microservice

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o movie-microservice

# --------------------------- Second stage: FINAL IMAGE BUILD -----------------------
FROM alpine:latest

ARG version
ENV VERSION="$version"
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /usr/local/go/src/movie-microservice .

CMD ["./movie-microservice"]
