FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apk add --no-cache upx ca-certificates && update-ca-certificates && \
    go version && \
    go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags '-w -s' -o /go/bin/api cmd/web/main.go && upx /go/bin/api

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/api /go/bin/api
ENTRYPOINT ["/go/bin/api"]
