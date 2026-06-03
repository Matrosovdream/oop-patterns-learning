# Go 1.26 environment for the Go track.
FROM golang:1.26

WORKDIR /app

# Pre-download modules if a go.mod is copied in; the source itself is mounted
# at runtime via docker-compose, so `go run ./cmd/<lesson>` just works.
ENV CGO_ENABLED=0

CMD ["go", "version"]
