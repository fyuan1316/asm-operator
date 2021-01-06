# Build the manager binary
FROM golang:1.13 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
ENV GOPROXY https://athens.alauda.cn

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/
COPY local/ local/

RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
#FROM gcr.io/distroless/static:nonroot
FROM alpine
WORKDIR /
COPY --from=builder /workspace/manager .
COPY --from=builder /workspace/pkg ./pkg
COPY files/ files/
#USER nonroot:nonroot
COPY --from=alpine/k8s:1.14.9 /usr/bin/kubectl /usr/local/bin/kubectl
RUN apk add --no-cache bash && rm -rf /var/cache/apk/*
ENTRYPOINT ["/manager"]
