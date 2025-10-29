FROM golang:1.24-bullseye AS build

# Install build dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    librdkafka-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /bin/bumpy .

# Runtime stage
FROM debian:bullseye-slim AS release
RUN apt-get update && apt-get install -y librdkafka1 && rm -rf /var/lib/apt/lists/*

COPY --from=build /bin/bumpy /usr/bin/bumpy
ENTRYPOINT ["bumpy","server"]
