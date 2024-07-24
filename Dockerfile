FROM registry.access.redhat.com/ubi9/go-toolset@sha256:d128c3d36878251f039606f144ef2608746c3203708b722295e6c571df1d8613 AS builder
# FROM registry.fedoraproject.org/fedora:40
WORKDIR /src
COPY bing.go .
RUN go build -o bing bing.go

FROM registry.access.redhat.com/ubi9/ubi:9.4-1123
COPY --from=builder /src/bing /bin/bing
CMD ["/bin/bing"]
