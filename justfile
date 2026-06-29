set shell := ["nu", "-c"]
os := env("GOOS", os())
arch := env("GOARCH", arch())
extension := if os == "windows" {
  ".exe"
} else {
  ""
}
app := "json2go"
cli := app + "-cli"
cli_file := cli + extension
ci := if env("GITHUB_ACTIONS", "false") == "true" {
  "true"
} else {
  "false"
}

default:
  @just --list

[group: 'setup']
setup: setup-wails setup-web

[group: 'setup']
setup-wails:
  go install github.com/wailsapp/wails/v3/cmd/wails3@latest

[group: 'setup']
[windows]
setup-inno:
  winget install -e --id JRSoftware.InnoSetup.7 --accept-source-agreements --accept-package-agreements
  if {{ci}} { "C:\\Program Files\\Inno Setup 7" | $in + "\n" | save --append $env.GITHUB_PATH }

[group: 'setup']
[windows]
setup-inno-env:
  #!pwsh -File
  $InnoPath = "C:\Program Files\Inno Setup 7"
  if (-not (Test-Path $InnoPath)) {
    $InnoPath = Join-Path -Path $env:LOCALAPPDATA -ChildPath "Programs\Inno Setup 7"
    if (-not (Test-Path $InnoPath)) {
      return
    }
  }

  $Path = [Environment]::GetEnvironmentVariable("Path", "User")
  [Environment]::SetEnvironmentVariable("Path", $InnoPath + ";" + $Path, "User")


[group: 'setup']
[windows]
setup-upx:
  winget install -e --id UPX.UPX --accept-source-agreements --accept-package-agreements
  if {{ci}} { $env.LOCALAPPDATA | path join Microsoft\WinGet\Links | $in + "\n" | save --append $env.GITHUB_PATH }

[group: 'setup']
[group: 'web']
[working-directory: 'ui']
setup-web:
  bun install

[group: 'web']
[working-directory: 'ui']
copy-monaco:
  #!nu
  let src = ("node_modules/monaco-editor/min/vs" | path expand)
  let dst = ("public/monaco-editor/min/vs" | path expand)
  mkdir $dst
  cd $src
  cp -r assets editor basic-languages language/json $dst
  cp loader.js nls.messages-loader.js nls.messages.zh-cn.js $dst
  cp `_commonjsHelpers-*.js` `go-*.js` `editor.api-*.js` `jsonMode-*.js` `lspLanguageFeatures-*.js` `workers-*.js` $dst

[group: 'web']
[working-directory: 'ui']
web-build: setup-web copy-monaco
  bun run build

[group: 'web']
[working-directory: 'ui']
web-dev: setup-web copy-monaco
  bun run dev

[group: 'web']
[working-directory: 'ui']
web-lint: setup-web
  bun lint

[group: 'web']
[working-directory: 'ui']
web-fmt: setup-web
  bun fmt

fmt: web-fmt

[working-directory: 'app']
dev:
  $env.JSON2GO_DEV = true; $env.JSON2GO_DEBUG = true; wails3 task dev

test:
  go test -v ./...

[group: 'build']
build: build-cli build-wails

[group: 'build']
build-release: build-cli build-wails-prod

[group: 'build']
[working-directory: 'cmd/json2go']
build-cli:
  go build -o ../../bin/{{cli_file}}

[group: 'build']
[working-directory: 'app']
build-wails:
  wails3 task build

[group: 'build']
[working-directory: 'app']
build-wails-prod:
  $env.PRODUCTION = "true"; wails3 task build

[group: 'packge']
[working-directory: 'bin']
[windows]
package-cli: build-cli
   powershell Compress-Archive -Force {{cli_file}} {{cli}}-{{os}}-{{arch}}.zip

[group: 'packge']
[working-directory: 'bin']
[unix]
package-cli: build-cli
  tar -czf {{cli}}-{{os}}-{{arch}}.tar.gz {{cli_file}}

[group: 'packge']
[working-directory: 'app']
package-app:
  wails3 task package
