# Use code-server base image
FROM codercom/code-server:latest

# Build arguments
ARG PACKAGES=""
ARG USERNAME=coder

# Convert ARG → ENV (VERY IMPORTANT)
ENV USERNAME=${USERNAME}

# Set working directory
WORKDIR /home/${USERNAME}

# Switch to root
USER root

# Install packages
RUN apt-get update && \
    apt-get install -y git curl ${PACKAGES} && \
    mkdir -p /home/${USERNAME} && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT []

# Expose port
EXPOSE 8080

# Runtime command
CMD ["/bin/sh", "-c", "\
if [ -n \"$REPO_URL\" ]; then \
  echo 'Cloning repo...'; \
  git clone $REPO_URL /home/${USERNAME}/ || true; \
fi; \
exec code-server --bind-addr 0.0.0.0:8080 --auth none"]