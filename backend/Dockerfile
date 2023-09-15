# Build
FROM golang:alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-w -s" -o /app/bin/myapp cmd/is-lab/main.go

# Run
FROM scratch

COPY --from=builder /app/bin/myapp /app/bin/myapp

EXPOSE 8888

ENTRYPOINT ["/app/bin/myapp"]