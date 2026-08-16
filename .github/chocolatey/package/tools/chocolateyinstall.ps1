# Rendered by .github/chocolatey/publish.ps1; __PLACEHOLDERS__ are filled per release.
$ErrorActionPreference = 'Stop'
$toolsDir = Split-Path -Parent $MyInvocation.MyCommand.Definition

$packageArgs = @{
  packageName    = 'cadenya'
  unzipLocation  = $toolsDir
  url64bit       = '__URL__'
  checksum64     = '__CHECKSUM__'
  checksumType64 = 'sha256'
}
# cadenya.exe lands in tools\ and Chocolatey shims it onto PATH automatically.
Install-ChocolateyZipPackage @packageArgs
