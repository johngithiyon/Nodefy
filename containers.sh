#!/bin/bash
# Automatically attach the SSH user to the container with the same name
sudo usermod -aG docker "${USER}"

echo "Your containers:"

docker ps \
--filter "label=owner=$USER" \
--format "{{.Label \"instance\"}}"

echo ""
echo "Enter instance name:"
read INSTANCE

# Find container name using labels
CONTAINER=$(docker ps \
--filter "label=owner=$USER" \
--filter "label=instance=$INSTANCE" \
--format "{{.Names}}")

if [ -z "$CONTAINER" ]; then
    echo "Container not found or not owned by you"
    exit 1
fi

docker exec -it "$CONTAINER" /bin/bash
