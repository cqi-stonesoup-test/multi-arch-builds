FROM registry.access.redhat.com/ubi9/ubi:9.3-1361.1699548029 AS builder
RUN dnf install -y golang
WORKDIR /src
COPY bing.go .
RUN go build -o bing bing.go

FROM registry.access.redhat.com/ubi9/ubi:9.3-1361.1699548029
COPY --from=builder /src/bing /bin/bing
CMD ["/bin/bing"]
