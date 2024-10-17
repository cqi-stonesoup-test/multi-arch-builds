FROM registry.access.redhat.com/ubi9/go-toolset@sha256:b08eed026ac732d125e61ebfe3a0ef1dea7293b6a1c6d8b71c007a3088c4ed3b AS builder
# FROM registry.fedoraproject.org/fedora:40
WORKDIR /src
COPY bing.go .
RUN go build -o bing bing.go

FROM registry.access.redhat.com/ubi9/ubi:9.4-1181
COPY --from=builder /src/bing /bin/bing
CMD ["/bin/bing"]
