# Use Ubuntu 20.04 as base image
FROM ubuntu:jammy

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive

# Update package lists and install necessary dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    build-essential \
    llvm \
    # clang \
    curl 
    # && rm -rf /var/lib/apt/lists/*

# Install Rust using Rustup
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y

# Add cargo and rust to PATH
ENV PATH="/root/.cargo/bin:${PATH}"

# Display installed versions
RUN rustup --version && \
    rustc --version && \
    cargo --version

# aya setting 
RUN rustup install stable
RUN rustup toolchain install nightly --component rust-src
RUN cargo install bpf-linker
RUN cargo install --no-default-features bpf-linker
RUN cargo install cargo-generate
RUN cargo generate https://github.com/aya-rs/aya-template

# Set the default working directory
WORKDIR /app

# CMD [ "bash" ] # If you want to enter into the container shell when it starts
