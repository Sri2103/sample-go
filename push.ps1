param(
    [string]$branch = "main"
)

Write-Host "Pushing to GitHub..."
# git push github $branch

Write-Host "Pushing to Gitea..."
git push gitea $branch

Write-Host "`n✅ Push complete to both remotes."
