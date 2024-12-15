# Use Alpine 3.16 as Basic Image for Root Certs
FROM alpine:3.16 AS root-certs
RUN echo "http://dl-cdn.alpinelinux.org/alpine/v3.16/main" > /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/v3.16/community" >> /etc/apk/repositories && \
    apk update && \
    apk add --no-cache ca-certificates && \
    addgroup -g 1001 app && \
    adduser app -u 1001 -D -G app -h /home/app app

# Use Go 1.22 as Build Image
FROM golang:1.22 AS builder
WORKDIR /social-api-files
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/social-api ./cmd/api

# Use Scratch as Runtime Image
FROM scratch
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --from=root-certs /etc/group /etc/group
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/social-api /bin/social-api

# Set User
USER app

# Set Entrypoint
ENTRYPOINT ["/bin/social-api"]

# Expose Port
EXPOSE 8080