#define AppName "json2go"
#define AppExeName AppName + ".exe"
#define AppVersion "0.5.6"
#define GOOS "windows"
#define GOARCH "amd64"
#define RootDir "..\..\.."
#define BinDir RootDir + "\bin"

[Setup]
AppId={{FFE55700-E4DA-48D0-A746-C77AA87403D8}
AppName={#AppName}
AppVersion={#AppVersion}

DefaultDirName={autopf}\{#AppName}
PrivilegesRequired=lowest

OutputBaseFilename=json2go-{#AppVersion}-{#GOOS}-{#GOARCH}-setup
OutputDir={#BinDir}

SetupIconFile=icon.ico
WizardStyle=modern dynamic
SolidCompression=yes

ArchitecturesAllowed=x64compatible
ArchitecturesInstallIn64BitMode=x64compatible
SetupArchitecture=x64

[Files]
Source: {#BinDir}\json2go.exe; DestDir: {app}; Flags: ignoreversion
Source: MicrosoftEdgeWebview2Setup.exe; Flags: dontcopy

[Tasks]
Name: desktopicon; Description: {cm:CreateDesktopIcon}; GroupDescription: {cm:AdditionalIcons}

[Icons]
Name: {autodesktop}\{#AppName}; Filename: {app}\{#AppExeName}; Tasks: desktopicon

[Run]
Filename: {app}\{#AppExeName}; Description: {cm:LaunchProgram,{#AppName}}; Flags: nowait postinstall skipifsilent

[Languages]
Name: english; MessagesFile: compiler:Default.isl
Name: chinesesimplified; MessagesFile: compiler:Languages\ChineseSimplified.isl

[Code]
function IsWebView2Installed: Boolean;
var
  Version: string;
begin
  Result := False;
  if RegQueryStringValue(HKLM64, 'SOFTWARE\WOW6432Node\Microsoft\EdgeUpdate\Clients\{F3017226-FE2A-4295-8BDF-00C3A9A7E4C5}', 'pv', Version) then
    if Version <> '' then
    begin
      Result := True;
      Exit;
    end;
  if RegQueryStringValue(HKCU, 'Software\Microsoft\EdgeUpdate\Clients\{F3017226-FE2A-4295-8BDF-00C3A9A7E4C5}', 'pv', Version) then
    if Version <> '' then
      Result := True;
end;

procedure CurStepChanged(CurStep: TSetupStep);
var
  ResultCode: Integer;
begin
  if CurStep = ssPostInstall then
    if not IsWebView2Installed then
    begin
      ExtractTemporaryFile('MicrosoftEdgeWebview2Setup.exe');
      Exec(ExpandConstant('{tmp}\MicrosoftEdgeWebview2Setup.exe'), '/silent /install', '', SW_HIDE, ewWaitUntilTerminated, ResultCode);
    end;
end;
