FROM golang:latest as build
RUN apt-get update && apt-get install -y nocache git ca-certificates && update-ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY="https://goproxy.io,direct"
RUN go mod download
COPY . .
RUN go build -o /app/bin/light-house-command .



FROM debian:buster-slim
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=build /app/bin /app
EXPOSE 8080
# Run the executable
CMD ["./light-house-command"]