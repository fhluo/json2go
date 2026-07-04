#define AppName "json2go"
#define AppExeName AppName + ".exe"
#define AppVersion "0.5.0"
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
Source: {#BinDir}\json2go.exe; DestDir: {app}

[Tasks]
Name: desktopicon; Description: {cm:CreateDesktopIcon}; GroupDescription: {cm:AdditionalIcons}

[Icons]
Name: {autodesktop}\{#AppName}; Filename: {app}\{#AppExeName}; Tasks: desktopicon

[Run]
Filename: {app}\{#AppExeName}; Description: {cm:LaunchProgram,{#AppName}}; Flags: nowait postinstall skipifsilent
