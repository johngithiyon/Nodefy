# Use code-server base image
FROM codercom/code-server:latest

# Build arguments
ARG PACKAGES=""
ARG USERNAME=coder

# Convert ARG → ENV
ENV USERNAME=${USERNAME}

# Set working directory
WORKDIR /home/${USERNAME}

# Switch to root
USER root

# Install packages and cloudflared (ngrok removed)
RUN apt-get update && \
    apt-get install -y git curl unzip ${PACKAGES} && \
    # Download and install cloudflared
    curl -L https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64 \
    -o /usr/local/bin/cloudflared && \
    chmod +x /usr/local/bin/cloudflared && \
    mkdir -p /home/${USERNAME} && \
    rm -rf /var/lib/apt/lists/*

# Expose code-server port
EXPOSE 8080

# Empty entrypoint to allow CMD
ENTRYPOINT []

# Runtime command (UNCHANGED except ngrok block removed)
CMD ["/bin/sh", "-c", "\
if [ -n \"$REPO_URL\" ]; then \
  echo 'Cloning repo...'; \
  git clone $REPO_URL /home/${USERNAME}/ || true; \
fi; \
exec code-server --bind-addr 0.0.0.0:8080 --auth none"]