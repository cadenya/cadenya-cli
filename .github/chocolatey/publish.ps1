# Packs and pushes the Chocolatey package for an already-published GitHub
# release. The install script downloads the windows_amd64 zip goreleaser
# attached to the release and verifies it against goreleaser's checksums file,
# so nothing is rebuilt here — this only renders the templates under package/,
# runs `choco pack`, and pushes to the community feed.
#
# Requires: choco (windows runners), gh (authenticated via GH_TOKEN),
# CHOCOLATEY_API_KEY. Run from the repository root.
param(
    [Parameter(Mandatory = $true)][string]$Tag
)
$ErrorActionPreference = 'Stop'

$version = $Tag.TrimStart('v')
$repo = if ($env:GITHUB_REPOSITORY) { $env:GITHUB_REPOSITORY } else { 'cadenya/cadenya-cli' }
$zip = "cadenya_${version}_windows_amd64.zip"
$url = "https://github.com/$repo/releases/download/$Tag/$zip"
$checksums = "cadenya_${version}_checksums.txt"
$work = Join-Path ([System.IO.Path]::GetTempPath()) "cadenya-choco-$version"

if (Test-Path $work) { Remove-Item -Recurse -Force $work }
New-Item -ItemType Directory -Path $work | Out-Null

gh release download $Tag --repo $repo --pattern $checksums --dir $work --clobber
if ($LASTEXITCODE -ne 0) { throw "gh release download failed for $Tag" }

$line = Get-Content (Join-Path $work $checksums) | Where-Object { $_ -match "\s$([regex]::Escape($zip))$" }
if (-not $line) { throw "$zip not listed in $checksums" }
$checksum = ($line -split '\s+')[0]

$pkg = Join-Path $work 'package'
Copy-Item -Recurse (Join-Path $PSScriptRoot 'package') $pkg
foreach ($rel in @('cadenya.nuspec', 'tools/chocolateyinstall.ps1')) {
    $path = Join-Path $pkg $rel
    $content = Get-Content -Raw $path
    $content = $content.Replace('__VERSION__', $version)
    $content = $content.Replace('__URL__', $url)
    $content = $content.Replace('__CHECKSUM__', $checksum)
    $content = $content.Replace('__TAG__', $Tag)
    Set-Content -Path $path -Value $content -NoNewline -Encoding utf8
}

choco pack (Join-Path $pkg 'cadenya.nuspec') --outputdirectory $work
if ($LASTEXITCODE -ne 0) { throw 'choco pack failed' }

$nupkg = Join-Path $work "cadenya.$version.nupkg"
choco push $nupkg --source https://push.chocolatey.org/ --api-key $env:CHOCOLATEY_API_KEY
if ($LASTEXITCODE -ne 0) { throw 'choco push failed' }
Write-Host "Pushed $nupkg (community moderation may take a while for new versions)."
