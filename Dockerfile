FROM golang:1.21-alpine AS builder


ENV CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=${GOARCH}

WORKDIR /app

COPY ./go.* ./
RUN go mod download

# Copy local code to the container image.
COPY .. ./

# Build the binary.
RUN go build -v -o monolith ./cmd/monolith

FROM alpine:3.18 AS runtime

WORKDIR /app

COPY ./wait-for.sh .
RUN chmod +x wait-for.sh

COPY --from=builder /app/monolith /app/monolith

CMD ["/app/monolith"]
