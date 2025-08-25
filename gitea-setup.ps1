# setup-gitea.ps1
# PowerShell script to run Gitea locally with Docker

Write-Host "Pulling latest Gitea image..."
docker pull gitea/gitea:latest

Write-Host "Stopping any old Gitea container..."
docker stop gitea -ErrorAction SilentlyContinue | Out-Null
docker rm gitea -ErrorAction SilentlyContinue | Out-Null

Write-Host "Starting new Gitea container..."
docker run -d `
  --name gitea `
  -p 3000:3000 `
  -p 222:22 `
  -v gitea-data:/data `
  gitea/gitea:latest

Write-Host "`n✅ Gitea is running!"
Write-Host "Open: http://localhost:3000"
Write-Host "SSH access: ssh://git@localhost:222/<your-repo>.git"

# Optional: instructions for adding Gitea as a remote
$repoName = "sample-go"
$remoteUrl = "http://localhost:3000/git/$repoName.git"

Write-Host "`nTo add your repo in Gitea, first create it in the web UI."
Write-Host "Then run this inside your project folder:"
Write-Host "    git remote add gitea $remoteUrl"
