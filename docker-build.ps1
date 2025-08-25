# build-jenkins.ps1
# Builds a custom Jenkins image that can access the Docker socket

$ImageName = "jenkins:local"
$DockerFile = "DockerJenkins"
$DockerGroupId = 998   # <-- change this if your host's docker group GID is different

Write-Host "Building custom Jenkins image..."
docker build --build-arg DOCKER_GID=$DockerGroupId -t $ImageName $DockerFile .

Write-Host "Custom Jenkins image '$ImageName' built successfully."
