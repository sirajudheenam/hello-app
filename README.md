
# hello-go-app 
Simple Go web app that responds with "Hello World" and current time.

## Updates
v1.0.13 - Updated logging middleware to display latency, removed PR with update README.md


### Build Status

![Docker Build](https://github.com/sirajudheenam/hello-app/actions/workflows/docker-build.yml/badge.svg?branch=main)
![Docker Image Version](https://img.shields.io/docker/v/sirajudheenam/hello-go-app?sort=semver)



```bash
# Build the container
docker build -t hello-go-app .

# Run the container
docker run -p 8080:8080 hello-go-app

# docker run -it --rm hello-go-app-alpine bash


docker build -t hello-go-app:v1.0.0 .

docker run -p 8080:8080 hello-go-app:v1.0.0

docker login

docker tag hello-go-app:v1.0.0 sirajudheenam/hello-go-app:v1.0.0

docker push sirajudheenam/hello-go-app:v1.0.0

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

```


<!-- **Latest Docker Image:** `sirajudheenam/hello-go-app:{{TAG}}` -->

