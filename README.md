

```bash
# Build the container
docker build -t hello-go .

# Run the container
docker run -p 8080:8080 hello-go

# docker run -it --rm hello-go-alpine bash


docker build -t hello-go:v1.0.0 .

docker run -p 8080:8080 hello-go:v1.0.0

docker login

docker tag hello-go:v1.0.0 sirajudheenam/hello-go:v1.0.0

docker push sirajudheenam/hello-go:v1.0.0

# create a file .github/workflows/docker-build.yml
name: Build and Push Docker Image

on:
  push:
    branches:
      - main

jobs:
  docker:
    runs-on: ubuntu-latest

    steps:
      # Checkout the repo
      - name: Checkout code
        uses: actions/checkout@v4

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
          images: sirajudheenam/hello-go

      # Build and push the Docker image
      - name: Build and push Docker image
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: |
            ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}


# Generate SSH Keys on your macOS or compatible Systems
ssh-keygen -t ed25519 -C "your_email@example.com"

# on legacy systems use,
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"

# Some changes to the code

```