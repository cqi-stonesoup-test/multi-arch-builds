FROM registry.access.redhat.com/ubi9/go-toolset@sha256:bee74f192b88c76f4b1a41f9892cff62b346ce6224fdbf167e225918f8e9d216 AS builder
# FROM registry.fedoraproject.org/fedora:40
WORKDIR /src
COPY bing.go .
RUN go build -o bing bing.go

FROM registry.access.redhat.com/ubi9/ubi:9.4-1123
COPY --from=builder /src/bing /bin/bing
CMD ["/bin/bing"]
