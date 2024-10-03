FROM registry.access.redhat.com/ubi9/go-toolset@sha256:fd41c001abc243076cc28b63c409ae6d9cbcad401c8124fb67d20fe57a2aa63a AS builder
# FROM registry.fedoraproject.org/fedora:40
WORKDIR /src
COPY bing.go .
RUN go build -o bing bing.go

FROM registry.access.redhat.com/ubi9/ubi:9.4-1181
COPY --from=builder /src/bing /bin/bing
CMD ["/bin/bing"]
