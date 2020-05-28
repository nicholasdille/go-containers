# go-containers

go library for talking to container registries

## Building

Locally:

```bash
make
```

Containerized:

```bash
docker build --file Dockerfile.build --build-arg GO_PACKAGE=github.com/nicholasdille/go-containers --build-arg GO_OUTPUT=go-containers --tag go-containers .
```
