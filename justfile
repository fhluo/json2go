set shell := ["nu", "-c"]
goos := `go env GOOS`
goarch := `go env GOARCH`
extension := if goos == "windows" {
  ".exe"
} else {
  ""
}
app := "json2go"
version := "0.5.2"
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
setup-7z:
  winget install -e --id 7zip.7zip --accept-source-agreements --accept-package-agreements

[group: 'setup']
[windows]
setup-7z-env:
  #!pwsh -File
  $7zPath = "C:\Program Files\7-Zip"
  if (-not (Test-Path $7zPath)) {
    return
  }

  $Path = [Environment]::GetEnvironmentVariable("Path", "User")
  [Environment]::SetEnvironmentVariable("Path", $7zPath + ";" + $Path, "User")

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
  # Monaco files have hashed names that change across versions; clean first to avoid stale assets accumulating
  rm -rf $dst
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
  $env.JSON2GO_DEV = true; $env.JSON2GO_DEBUG = true; wails3 dev

[working-directory: 'bin']
run:
  {{app}}{{extension}}

test:
  go test -v ./...

[group: 'build']
check-version version:
  #!nu
  if "{{version}}" not-like '^v?\d+\.\d+\.\d+$' {
    print $"(ansi red)invalid version: {{version}}, expected a version like 0.5.0 or v0.5.0(ansi reset)"
    exit 1
  }

[group: 'build']
set-version version=version: (check-version version)
  "{{version}}" | str replace -r '^v' '' | just sync-version $in

[group: 'build']
bump-version version: (check-version version)
  #!nu
  let status = (git status --porcelain --untracked-files=no)
  if ($status | is-not-empty) {
    print $"(ansi red)working tree is not clean, please commit or stash changes first(ansi reset)"
    exit 1
  }

  let ver = ("{{version}}" | str replace -r '^v' '')
  print $"(ansi light_gray)Syncing version to ($ver)...(ansi reset)"
  just sync-version $ver

  print $"(ansi light_gray)Committing...(ansi reset)"
  git add justfile internal/version/version.go app/build/windows/info.json ui/package.json app/build/config.yaml app/build/windows/json2go.exe.manifest app/build/windows/json2go.iss
  git commit -m $"chore: bump version to ($ver)"

  print $"(ansi light_gray)Tagging v($ver)...(ansi reset)"
  git tag $"v($ver)" -m $"v($ver)"

  print $"(ansi green)✓ Bumped to ($ver)(ansi reset)"

[group: 'build']
sync-version version=version:
  #!nu
  open justfile
  | str replace -r 'version := "\d+\.\d+\.\d+"' 'version := "{{version}}"'
  | save -f justfile

  open internal/version/version.go
  | str replace -r 'const version = "\d+\.\d+\.\d+"' 'const version = "{{version}}"'
  | save -f internal/version/version.go

  open --raw app/build/windows/info.json
  | str replace -a -r '"\d+\.\d+\.\d+"' '"{{version}}"'
  | save -f app/build/windows/info.json

  open --raw ui/package.json
  | str replace -r '"version": "\d+\.\d+\.\d+"' '"version": "{{version}}"'
  | save -f ui/package.json

  open --raw app/build/config.yaml
  | str replace -r 'version: "\d+\.\d+\.\d+"' 'version: "{{version}}"'
  | save -f app/build/config.yaml

  open app/build/windows/json2go.exe.manifest
  | str replace -r 'name="json2go" version="\d+\.\d+\.\d+"' 'name="json2go" version="{{version}}"'
  | save -f app/build/windows/json2go.exe.manifest

  open app/build/windows/json2go.iss
  | str replace -r '#define AppVersion "\d+\.\d+\.\d+"' '#define AppVersion "{{version}}"'
  | save -f app/build/windows/json2go.iss

[group: 'build']
build: build-cli build-wails

[group: 'build']
build-release *args: build-cli build-wails-prod
  def main [ --upx ] { if $upx { just upx } }; main {{args}}

[group: 'build']
[working-directory: 'bin']
upx: upx-cli upx-app

[group: 'build']
[working-directory: 'bin']
upx-cli:
  upx {{cli_file}}

[group: 'build']
[working-directory: 'bin']
upx-app:
  upx {{app}}{{extension}}

[group: 'build']
[working-directory: 'cmd/json2go']
build-cli:
  go build -o ../../bin/{{cli_file}}

[group: 'build']
tidy:
  go mod tidy

[group: 'build']
[working-directory: 'app']
generate-bindings:
  wails3 generate bindings -d ../ui/bindings -ts -i -clean=true
  just web-fmt

[group: 'build']
[working-directory: 'app']
[windows]
generate-icons:
  wails3 generate icons -input build/appicon.png -windowsfilename build/windows/icons.ico

[group: 'build']
[working-directory: 'app']
[macos]
generate-icons:
  wails3 generate icons -input build/appicon.png -macfilename build/darwin/icons.icns

[group: 'build']
[working-directory: 'app']
[windows]
build-wails: tidy generate-bindings web-build generate-icons
  wails3 generate syso -arch {{goarch}} -icon build/windows/icon.ico -manifest build/windows/json2go.exe.manifest -info build/windows/info.json -out wails_windows_{{goarch}}.syso
  go build -buildvcs=false -gcflags=all=-l -o ../bin/{{app}}{{extension}}
  rm *.syso

[group: 'build']
[working-directory: 'app']
[windows]
build-wails-prod: tidy generate-bindings web-build generate-icons
  wails3 generate syso -arch {{goarch}} -icon build/windows/icon.ico -manifest build/windows/json2go.exe.manifest -info build/windows/info.json -out wails_windows_{{goarch}}.syso
  go build -tags production -trimpath -buildvcs=false -ldflags="-w -s -H windowsgui" -o ../bin/{{app}}{{extension}}
  rm *.syso

[group: 'build']
[working-directory: 'app']
[linux]
build-wails: tidy generate-bindings web-build
  go build -buildvcs=false -gcflags=all=-l -o ../bin/{{app}}{{extension}}

[group: 'build']
[working-directory: 'app']
[linux]
build-wails-prod: tidy generate-bindings web-build
  go build -tags production -trimpath -buildvcs=false -ldflags="-w -s" -o ../bin/{{app}}{{extension}}

[group: 'build']
[working-directory: 'app']
[macos]
build-wails: tidy generate-bindings web-build generate-icons
  go build -buildvcs=false -gcflags=all=-l -o ../bin/{{app}}{{extension}}

[group: 'build']
[working-directory: 'app']
[macos]
build-wails-prod: tidy generate-bindings web-build generate-icons
  go build -tags production -trimpath -buildvcs=false -ldflags="-w -s" -o ../bin/{{app}}{{extension}}

[group: 'packge']
package: package-cli package-app

[group: 'packge']
[working-directory: 'bin']
[windows]
package-cli: build-cli upx-cli
   7z a -mx9 {{cli}}-{{version}}-{{goos}}-{{goarch}}.zip {{cli_file}}

[group: 'packge']
[working-directory: 'bin']
[unix]
package-cli: build-cli
  tar -czf {{cli}}-{{version}}-{{goos}}-{{goarch}}.tar.gz {{cli_file}}

[group: 'packge']
[working-directory: 'app/build/windows']
[windows]
iscc:
  #!nu
  let inno = (
    ["C:\\Program Files\\Inno Setup 7", $"($env.LOCALAPPDATA)\\Programs\\Inno Setup 7"]
    | where {|p| $p | path exists }
    | first
  )

  if ($inno != null) {
    $env.Path = ($env.Path | prepend $inno)

    let lang = ($inno | path join "Languages" "ChineseSimplified.isl")
    if not ($lang | path exists) {
      print $"(ansi light_gray)Downloading Simplified Chinese translation for Inno Setup...(ansi reset)"
      http get "https://raw.githubusercontent.com/kira-96/Inno-Setup-Chinese-Simplified-Translation/main/ChineseSimplified.isl"
      | save $lang
      print $"(ansi green)✓ Saved to ($lang)(ansi reset)"
    }
  }

  wails3 generate webview2bootstrapper -dir . e>| ignore
  ISCC json2go.iss

[group: 'packge']
[working-directory: 'app']
[windows]
package-app: build-wails-prod upx-app iscc
