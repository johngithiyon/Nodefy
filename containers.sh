#!/bin/bash
sudo usermod -aG docker "${USER}
echo "Your containers:"

docker ps \
--filter "label=owner=$USER" \
--format "{{.Names}}"

echo ""
echo "Enter container name:"
read CONTAINER

docker exec -it "$CONTAINER" /bin/bash