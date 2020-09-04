# Step 1
# Build the vans binary
FROM golang:1.13 as builder
WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY core/ core/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o vans main.go

# Step 2
FROM alpine:3.8
COPY --from=builder /workspace/vans /bin/
COPY 3rdparty/csvq /bin/
COPY csv-render-wrapper.sh /bin/
