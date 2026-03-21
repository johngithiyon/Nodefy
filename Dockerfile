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

# Install packages and ngrok
RUN apt-get update && \
    apt-get install -y git curl unzip ${PACKAGES} && \
    # Download and install ngrok
    curl -sSL https://bin.equinox.io/c/bNyj1mQVY4c/ngrok-v4-stable-linux-amd64.zip -o /tmp/ngrok.zip && \
    unzip /tmp/ngrok.zip -d /usr/local/bin/ && \
    chmod +x /usr/local/bin/ngrok && \
    rm /tmp/ngrok.zip && \
    mkdir -p /home/${USERNAME} && \
    rm -rf /var/lib/apt/lists/*

# Expose code-server port
EXPOSE 8080

# Empty entrypoint to allow CMD
ENTRYPOINT []

# Runtime command: clone repo, start ngrok (if authtoken), then code-server
CMD ["/bin/sh", "-c", "\
if [ -n \"$REPO_URL\" ]; then \
  echo 'Cloning repo...'; \
  git clone $REPO_URL /home/${USERNAME}/ || true; \
fi; \
if [ -n \"$NGROK_AUTHTOKEN\" ]; then \
  echo 'Adding ngrok authtoken...'; \
  ngrok config add-authtoken $NGROK_AUTHTOKEN; \
  echo 'Starting ngrok tunnel...'; \
  ngrok http 8080 & \
fi; \
exec code-server --bind-addr 0.0.0.0:8080 --auth none"]