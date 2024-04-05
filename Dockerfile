FROM golang:bookworm

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
ENV GIN_MODE=release
COPY . .
RUN go build -v -o /usr/local/bin/terrapak ./cmd/main.go

ENTRYPOINT ["/usr/local/bin/terrapak"]