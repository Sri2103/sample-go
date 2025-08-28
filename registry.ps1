$containerName = "registry"
$exportPort = 5000


$exists = docker ps -a --format "{{.Names}}" | Select-String $containerName

if ($exists) {
    Write-Host "Container $containerName already exists removing..."
    docker stop $containerName

    Start-Sleep -Seconds 5
    docker rm -f $containerName
}


Write-Host "Starting registry container"

docker run -d `
    --name $containerName `
    -p ${exportPort}:5000 `
    registry:2

Write-Host "registry is starting ...."

Write-Host "User http://localhost:5000 as based of the image repo"
