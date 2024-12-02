# Use the Ubuntu base image
FROM ubuntu:latest

# Set environment variables for Go
ENV GO_VERSION=1.23.3
ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Install basic tools
RUN apt-get update && \
    apt-get install -y \
    curl \
    wget \
    git \
    vim \
    nano \
    build-essential

SHELL ["/bin/bash", "-c"]

# Set working directory
WORKDIR /workspace