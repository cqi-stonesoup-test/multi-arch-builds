FROM registry.access.redhat.com/ubi9/go-toolset@sha256:f001ad1001a22fe5f6fc7d876fc172b01c1b7dcd6c498f83a07b425e24275a79 AS builder
# FROM registry.fedoraproject.org/fedora:40
WORKDIR /src
COPY bing.go .
RUN go build -o bing bing.go

FROM registry.access.redhat.com/ubi9/ubi:9.3-1361.1699548029
COPY --from=builder /src/bing /bin/bing
CMD ["/bin/bing"]
