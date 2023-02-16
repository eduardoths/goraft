FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
  --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
  go mod download
COPY . .


ENV GOGOARCH=amd64 GOOS=linux
RUN --mount=type=cache,target=/root/.cache/go-build \
  go build -tags musl -ldflags '-w -extldflags "-static"' -a -installsuffix cgo -o main cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
