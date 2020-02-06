FROM golang:latest

RUN mkdir -p /go/src/github.com/backend

# Set WorkDir
WORKDIR /go/src/github.com/backend

# Copy codes on server into docker
COPY . /go/src/github.com/backend

# Set Environment Variables
ENV ENV=product
    GO111MODULE=on
    GOPROXY=https://goproxy.io

# Expose Port
EXPOSE 9000

# Check files
RUN pwd && cd /go/src/github.com/backend && ls
