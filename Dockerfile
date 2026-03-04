# Allow dynamic base image
ARG BASE_IMAGE
FROM ${BASE_IMAGE}

# Packages to install (space separated)
ARG PACKAGES=""

WORKDIR /app

# Install packages depending on detected OS
RUN set -e; \
    if command -v apt-get >/dev/null 2>&1; then \
        echo "Detected Debian/Ubuntu"; \
        DEBIAN_FRONTEND=noninteractive \
        apt-get update; \
        apt-get install -y --no-install-recommends ${PACKAGES}; \
        rm -rf /var/lib/apt/lists/*; \
    elif command -v apk >/dev/null 2>&1; then \
        echo "Detected Alpine"; \
        apk update; \
        apk add --no-cache ${PACKAGES}; \
    elif command -v dnf >/dev/null 2>&1; then \
        echo "Detected RHEL/Fedora"; \
        dnf install -y ${PACKAGES}; \
        dnf clean all; \
    elif command -v yum >/dev/null 2>&1; then \
        echo "Detected CentOS"; \
        yum install -y ${PACKAGES}; \
        yum clean all; \
    else \
        echo "Unsupported Linux distribution"; \
        exit 1; \
    fi    

# Default command (can be overridden at runtime)
CMD ["sh"]