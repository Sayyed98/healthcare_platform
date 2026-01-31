FROM golang:1.25.5-alpine

WORKDIR /app

COPY proto ./proto
COPY hospital-service ./hospital-service

WORKDIR /app/hospital-service

RUN go mod download
RUN go build -o app

EXPOSE 8081
CMD ["./app"]