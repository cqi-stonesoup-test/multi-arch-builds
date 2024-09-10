FROM registry.access.redhat.com/ubi9/go-toolset@sha256:ae17d73e70a966f39ef4dfca74241e3ca4374cd1198b02c30ea0748b8dcc83a6 AS builder
# FROM registry.fedoraproject.org/fedora:40
WORKDIR /src
COPY bing.go .
RUN go build -o bing bing.go

FROM registry.access.redhat.com/ubi9/ubi:9.4-1214.1725849297
COPY --from=builder /src/bing /bin/bing
CMD ["/bin/bing"]
