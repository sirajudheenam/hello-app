
# hello-go-app 
Simple Go web app that responds with "Hello World" and current time.

### Build Status

![Docker Build](https://github.com/sirajudheenam/hello-app/actions/workflows/docker-build.yml/badge.svg?branch=main)
![Docker Image Version](https://img.shields.io/docker/v/sirajudheenam/hello-go-app?sort=semver)



```bash

go mod init hello-go-app

go get gopkg.in/natefinch/lumberjack.v2

go run main.go

# optionally 
go build -o hello-server main.go
./hello-server

# create local dir called logs
mkdir logs


# Build the container
docker build -t hello-go-app:latest .

# Run the container
docker run -p 8080:8080 hello-go-app:latest

# run with local logs directory
# docker run -p 8080:8080 -v $(pwd)/logs:/app/logs hello-go-app:latest

# 
# docker run -it --rm hello-go-app-alpine bash


docker build -t hello-go-app:v1.0.0 .

docker run -p 8080:8080 hello-go-app:v1.0.0

docker login

docker tag hello-go-app:v1.0.0 sirajudheenam/hello-go-app:v1.0.0

docker push sirajudheenam/hello-go-app:v1.0.0

# extract container ID from docker ps command
CONTAINER_ID=$(docker ps -q -f "ancestor=hello-go-app:latest" | head -n 1)
echo $CONTAINER_ID

# Make a Reques to our server on command line or Web Browser
curl localhost:8080

# To view logs the request
docker exec -it $CONTAINER_ID cat /app/logs/server.json.log

Starting server on :8080
{"timestamp":"2025-09-09T01:04:18Z","remote_addr":"173.194.69.82:22615","method":"GET","path":"/","status":200,"latency_ms":0,"response_size":40}
{"timestamp":"2025-09-09T01:04:18Z","remote_addr":"173.194.69.82:22615","method":"GET","path":"/favicon.ico","status":200,"latency_ms":0,"response_size":40}
{"timestamp":"2025-09-09T01:10:22Z","remote_addr":"173.194.69.82:33847","method":"GET","path":"/","status":200,"latency_ms":0,"response_size":40}


# create a file .github/workflows/docker-build.yml
name: Build and Push Docker Image

on:
  push:
    branches:
      # - main
      - 'v*'   # only run when pushing version tags like v1.0.0

jobs:
  docker:
    runs-on: ubuntu-latest
    # Secrets are located under an environment gh secret list -e DOCKER
    environment: DOCKER # 🔑 use DOCKER environment

    steps:
      # Checkout the repo
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set image tags
        id: vars
        run: |
          echo "date=$(date +'%Y%m%d')" >> $GITHUB_OUTPUT
          echo "sha=$(echo $GITHUB_SHA | cut -c1-7)" >> $GITHUB_OUTPUT
          echo "tag=${echo $GITHUB_REF_NAME}" >> $GITHUB_OUTPUT

      # Log in to Docker Hub (set secrets in repo settings)
      - name: Log in to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      # Extract metadata (tags, labels) for Docker
      - name: Extract Docker metadata
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: sirajudheenam/hello-go-app

      # Build and push the Docker image
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: |
            sirajudheenam/hello-go-app:latest
            sirajudheenam/hello-go-app:${{ steps.vars.outputs.date }}-${{ steps.vars.outputs.sha }}
            sirajudheenam/hello-go-app:${{ steps.vars.outputs.tag }}

# Generate SSH Keys on your macOS or compatible Systems
ssh-keygen -t ed25519 -C "your_email@example.com"

# on legacy systems use,
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"

# Copy the public key ed25519.pub or id_rsa.pub content to your github under 
# https://github.com/settings/profile
# https://github.com/settings/ssh/new 
# Title: <any title you like>
# Key type: Authentication Key
# Key: <PASTE your ed25519.pub or id_rsa.pub content>
# Click on Add SSH Key button to save it.

```


<!-- **Latest Docker Image:** `sirajudheenam/hello-go-app:{{TAG}}` -->

## Updates
v1.0.23 - Updated logging middleware to display latency, removed PR with update README.md