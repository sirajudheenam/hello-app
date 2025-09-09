#!/bin/bash

# making a curl request to the application
echo "Making a curl request to the application at http://localhost:8080"
echo "--------------------------------------------------------------------------------"
curl localhost:8080

# Get container ID
CONTAINER_ID=$(docker ps -q -f "ancestor=hello-go-app:latest" | head -n 1)

if [ -z "$CONTAINER_ID" ]; then
    echo "No container found for image hello-go-app:latest"
    exit 1
fi

echo "Container ID: $CONTAINER_ID"

# Example: view logs
docker logs $CONTAINER_ID

echo "--------------------------------------------------------------------------------"
echo "Viewing logs from /app/logs/server.json.log for $CONTAINER_ID "
echo "--------------------------------------------------------------------------------"
# View logs with timestamps on log file
# docker exec -it $CONTAINER_ID tail -f /app/logs/server.json.log
docker exec -it $CONTAINER_ID cat /app/logs/server.json.log

echo "--------------------------------------------------------------------------------"
echo "Checking size of /app/logs/server.json.log for $CONTAINER_ID "
docker exec $CONTAINER_ID du -h /app/logs/server.json.log
echo "--------------------------------------------------------------------------------"

# Get log size in bytes
LOG_SIZE=$(docker exec $CONTAINER_ID stat -c%s /app/logs/server.json.log)
echo "Log file size: $LOG_SIZE bytes"

echo "--------------------------------------------------------------------------------"
# Get the latest tag
LATEST_TAG=$(git tag --sort=-v:refname | head -n 1)

if [ -z "$LATEST_TAG" ]; then
  echo "No tags found"
  exit 1
fi

echo "Latest tag: $LATEST_TAG"
