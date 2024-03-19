# Use Ubuntu 20.04 as base image
FROM ubuntu:22.04

COPY . /app

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive

# Update package lists and install necessary dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    build-essential \
    curl 
    # && rm -rf /var/lib/apt/lists/*

# Install Rust using Rustup
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y

# Add cargo and rust to PATH
ENV PATH="/root/.cargo/bin:${PATH}"


# Set the default working directory
WORKDIR /app

# Display installed versions
RUN rustup --version && \
    rustc --version && \
    cargo --version

# CMD [ "bash" ] # If you want to enter into the container shell when it starts
