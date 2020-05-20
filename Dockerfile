FROM golang
RUN mkdir -p /go/src/github.com/nicholasdille/registry
WORKDIR /go/src/github.com/nicholasdille/registry
COPY . .