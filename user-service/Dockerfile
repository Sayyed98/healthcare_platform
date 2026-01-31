FROM golang:1.25.5-alpine

WORKDIR /app

# Copy ONLY what is needed from build context
COPY proto ./proto
COPY user-service ./user-service

WORKDIR /app/user-service

RUN go mod download
RUN go build -o app

EXPOSE 8080 9090
CMD ["./app"]
