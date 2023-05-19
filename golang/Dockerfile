## Dev stage: Install all tools needed for a production developer environment
FROM golang:1.20-bullseye AS dev

RUN go install -v golang.org/x/tools/gopls@latest

## Build stage: Create build artifacts
FROM dev AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN \
    --mount=type=cache,id=dl_cache,target=/go/pkg/mod \
    go mod download -x

COPY . .
RUN \
    --mount=type=cache,id=build_cache,target=/root/.cache/go-build \
    --mount=type=cache,id=dl_cache,target=/go/pkg/mod \
    go build -v -o build/hello ./cmd/hello/

## Production release stage
FROM golang:1.20-bullseye AS release

RUN \
    --mount=type=cache,id=apt_cache,target=/var/lib/apt \
    apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates \
    tzdata

COPY --from=build /app/build/hello /hello

ENTRYPOINT [ "/hello" ]
