# PowerShell script for Docker build, tag and push

# Check if enough parameters are provided
param (
    [Parameter(Mandatory=$true)][string]$IMAGE_NAME,
    [Parameter(Mandatory=$true)][string]$TAG,
    [Parameter(Mandatory=$true)][string]$DOCKERFILE_PATH
)

# Fixed Docker repository address
$REPO_URL = "registry.cn-hangzhou.aliyuncs.com/pf94514203"

# Build Docker image
Write-Host "Building Docker image..."
docker build -t "$IMAGE_NAME`:$TAG" -f $DOCKERFILE_PATH .

# Check if build was successful
if ($LASTEXITCODE -ne 0) {
    Write-Host "Docker image build failed"
    exit 1
}

Write-Host "Build successful: $IMAGE_NAME`:$TAG"

# Tag the image
Write-Host "Tagging image..."
docker tag "$IMAGE_NAME`:$TAG" "$REPO_URL/$IMAGE_NAME`:$TAG"

# Check if tagging was successful
if ($LASTEXITCODE -ne 0) {
    Write-Host "Image tagging failed"
    exit 1
}

Write-Host "Tagging successful: $REPO_URL/$IMAGE_NAME`:$TAG"

# Push image to repository
Write-Host "Pushing image to repository..."
docker push "$REPO_URL/$IMAGE_NAME`:$TAG"

# Check if push was successful
if ($LASTEXITCODE -ne 0) {
    Write-Host "Image push failed"
    exit 1
}

Write-Host "Image push successful: $REPO_URL/$IMAGE_NAME`:$TAG"

# Complete
Write-Host "Task completed"




## ./deploy/dockerfile/run.ps1 blog bff.001 ./deploy/dockerfile/bff.Dockerfile
## ./deploy/dockerfile/run.ps1 blog rpc.account.001 ./deploy/dockerfile/rpc.account.Dockerfile
## ./deploy/dockerfile/run.ps1 blog rpc.blogs.001 ./deploy/dockerfile/rpc.blogs.Dockerfile
## ./deploy/dockerfile/run.ps1 blog rpc.tools.001 ./deploy/dockerfile/rpc.tools.Dockerfile
## ./deploy/dockerfile/run.ps1 blog rpc.user.001 ./deploy/dockerfile/rpc.user.Dockerfile
## ./deploy/dockerfile/run.ps1 blog rpc.travel.001 ./deploy/dockerfile/rpc.travel.Dockerfile
## ./deploy/dockerfile/run.ps1 blog rpc.palaces.001 ./deploy/dockerfile/rpc.palaces.Dockerfile