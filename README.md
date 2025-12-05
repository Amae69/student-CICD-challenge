# Student CI/CD Challenge

A simple Go application designed to demonstrate a complete CI/CD pipeline using GitHub Actions, Docker, and Kubernetes.

## 🚀 Project Overview

This project consists of:
- A basic **Go CLI application** that prints a welcome message.
- A **Dockerfile** for containerizing the application.
- A **Kubernetes Deployment** configuration for orchestration.
- A **GitHub Actions Workflow** that automates:
  - Building and testing the Go code.
  - Building and pushing the Docker image to Docker Hub.
  - Deploying the application to an ephemeral Minikube cluster for verification.

## 🛠️ Prerequisites

- [Go 1.25+](https://go.dev/dl/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) (for local Kubernetes testing)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

## 🏃‍♂️ Running Locally

### Standard Go Run
```bash
go run main.go
```

### Build Binary
```bash
go build -o app.exe
./app.exe
```

### Run with Docker
Build the image:
```bash
docker build -t student-ci-challenge .
```

Run the container:
```bash
docker run --rm student-ci-challenge
```

## ☸️ Kubernetes Deployment

To deploy locally to Minikube:

1. Start Minikube:
   ```bash
   minikube start
   ```

2. Apply the deployment manifest:
   ```bash
   kubectl apply -f k8s/deployment.yaml
   ```

3. Check the status:
   ```bash
   kubectl get pods
   ```

## 🔄 CI/CD Pipeline

The pipeline is defined in `.github/workflows/ci.yml` and triggers on pushes to the `main` branch.

### Stages:
1. **Build**: Compiles the Go application.
2. **Test**: Runs unit tests (if any).
3. **Docker**: Builds the Docker image and pushes it to Docker Hub (requires `DOCKER_USERNAME` and `DOCKER_PASSWORD` secrets).
4. **Deploy**: Spins up an ephemeral Minikube cluster inside the CI runner and deploys the application to verify the manifests.
