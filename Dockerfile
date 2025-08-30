FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/wievielnoch cmd/wievielnoch/main.go

FROM scratch
COPY --from=builder /build/bin/wievielnoch /wievielnoch
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/wievielnoch"]