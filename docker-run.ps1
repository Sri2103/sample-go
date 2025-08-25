# run-jenkins.ps1
# This script runs Jenkins inside Docker with persistent storage and Docker access

$containerName = "jenkins"
$jenkinsPort = 8080
$agentPort = 50000
$jenkinsHome = "jenkins_home"

# Check if container already exists
$exists = docker ps -a --format "{{.Names}}" | Select-String $containerName

if ($exists) {
    Write-Host "Container $containerName already exists. Removing..."
    docker stop $containerName
    docker rm -f $containerName
}

Write-Host "Starting Jenkins container..."
docker run -d `
    --name $containerName `
    -p ${jenkinsPort}:8080 -p ${agentPort}:50000 `
    -v ${jenkinsHome}:/var/jenkins_home `
    -e DOCKER_HOST=tcp://host.docker.internal:2375 `
    jenkins:local

Write-Host "Jenkins is starting..."
Write-Host "Open http://localhost:$jenkinsPort to access Jenkins"
