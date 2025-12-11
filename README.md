# Student CI/CD Challenge

A robust Go web application designed to demonstrate a complete CI/CD pipeline using GitHub Actions, Docker, and Kubernetes.

## 🚀 Project Overview

This project showcases a full DevOps workflow, taking a simple Go application from development to deployment.

### Key Components

#### 1. Go Web Server (`main.go`)
- **Functionality**: A lightweight HTTP server listening on port `8080`.
- **Features**:
  - Uses `net/http` for handling web requests.
  - Uses `github.com/fatih/color` to print colored logs to the console.
  - Returns a welcome message "Hello, DevelopersFoundry fellows!" to visitors.
- **Testing**: Includes `main_test.go` to verify the HTTP handler returns the correct status code (200 OK) and message body.

#### 2. Docker Containerization (`Dockerfile`)
- **Multi-Stage Build**:
  - **Builder Stage**: Uses `golang:1.25` to compile the application. This keeps the build tools separate from the final image.
  - **Final Stage**: Uses a minimal `debian:stable-slim` image to run the binary, reducing the image size and attack surface.
- **Port Exposure**: Exposes port `8080` to allow traffic to reach the application.

#### 3. Kubernetes Orchestration (`k8s/deployment.yaml`)
- **Deployment**: Manages the application pods, ensuring 1 replica is always running.
  - Uses the image `krizeal/my-repo:go-app_v2.0`.
  - Defines resource limits (CPU/Memory) to prevent resource hogging.
- **Service**: Exposes the application internally within the cluster on port 80.
- **Private Registry Support**: Configured with `imagePullSecrets` to securely pull images from a private Docker Hub repository.

#### 4. CI/CD Pipeline (`.github/workflows/ci.yml`)
Automated workflow triggered on pushes to `main`.
- **Build & Test**: Compiles the code and runs unit tests.
- **Docker Build**: Builds the image and pushes it to Docker Hub with a custom tag (`go-app_v2.0`).
- **Deployment Verification**:
  - Spins up an **ephemeral Minikube cluster** inside the CI runner.
  - Creates a Kubernetes secret (`regcred`) for private repo access.
  - Deploys the application.
  - **Verifies Success**: Waits for rollout, port-forwards, and `curl`s the app to ensure it's actually serving traffic.
  - **Debugs Failures**: Automatically prints pod logs and events if the deployment fails.

---

## 🛠️ Prerequisites

- [Go 1.25+](https://go.dev/dl/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) (for local Kubernetes testing)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

---

## 🏃‍♂️ Running Locally

### 1. Run with Go
You can run the server directly on your machine.
```bash
# Install dependencies
go mod tidy

# Run the app
go run main.go
```
*Visit `http://localhost:8080` in your browser.*

### 2. Run with Docker
Build and run the containerized version.
```bash
# Build the image
docker build -t student-ci-challenge .

# Run the container (mapping port 8080)
docker run -p 8080:8080 --rm student-ci-challenge
```

---

## 🌿 Dev Branch Workflow (How to Contribute)

Follow these steps to make changes and submit a Pull Request (PR) using a feature branch.

### Step 1: Create a Dev Branch
Never work directly on `main`. Create a new branch for your changes.
```bash
# Create and switch to a new branch named 'dev' (or feature/your-feature)
git checkout -b dev
```

### Step 2: Make Changes
Edit files, add features, or fix bugs.
```bash
# Check status of changed files
git status

# Add files to staging
git add .

# Commit your changes
git commit -m "feat: added new amazing feature"
```

### Step 3: Push to GitHub
Push your new branch to the remote repository.
```bash
# Push the 'dev' branch to origin
git push -u origin dev
```

### Step 4: Create a Pull Request (PR)
1. Go to your repository on **GitHub**.
2. You will see a banner saying **"dev had recent pushes"**. Click **"Compare & pull request"**.
3. **Base**: Set to `main`. **Compare**: Set to `dev`.
4. Add a title and description explaining your changes.
5. Click **"Create pull request"**.

Once created, your CI pipeline will run to verify your code. If it passes, you can merge the PR into `main`.

### END ....
