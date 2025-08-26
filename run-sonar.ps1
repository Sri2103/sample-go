param(
    [switch]$Stop,
    [switch]$Clean,
    [string]$TokenName = "jenkins-token",
    [string]$SonarUser = "admin",
    [string]$SonarPassword = "admin"
)

$composeFile = "docker-compose.sonarqube.yml"
# $sonarUrl = "http://localhost:9000"

# function Wait-ForSonar {
#     Write-Host "Waiting for SonarQube to start..."
#     $maxAttempts = 30
#     for ($i = 0; $i -lt $maxAttempts; $i++) {
#         try {
#             $response = Invoke-WebRequest -Uri "$sonarUrl/api/system/status" -UseBasicParsing -TimeoutSec 5
#             $json = $response.Content | ConvertFrom-Json
#             if ($json.status -eq "UP") {
#                 Write-Host "SonarQube is UP ✅"
#                 return $true
#             }
#         }
#         catch {
#             Start-Sleep -Seconds 5
#         }
#     }
#     throw "SonarQube did not start in time ❌"
# }

if ($Stop) {
    Write-Host "Stopping SonarQube stack..."
    docker compose -f $composeFile down
    exit
}

if ($Clean) {
    Write-Host "Removing SonarQube stack with volumes..."
    docker compose -f $composeFile down -v
    exit
}

Write-Host "Starting SonarQube stack..."
docker compose -f $composeFile up -d

# Wait until Sonar is ready
# Wait-ForSonar

# Create Jenkins token
# Write-Host "Creating token '$TokenName' for Jenkins..."
# $authHeader = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$SonarUser`:$SonarPassword"))
# $response = Invoke-RestMethod -Uri "$sonarUrl/api/user_tokens/generate?name=$TokenName" `
#     -Method Post -Headers @{ Authorization = "Basic $authHeader" }

# $token = $response.token
# Write-Host "`nGenerated Jenkins Token: $token"
# Write-Host "⚠️ Save this token in Jenkins > Manage Jenkins > Credentials > Secret Text"
