; Script generated by the Inno Setup Script Wizard.
; SEE THE DOCUMENTATION FOR DETAILS ON CREATING INNO SETUP SCRIPT FILES!

[Setup]
; NOTE: The value of AppId uniquely identifies this application.
; Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{DC598D8E-8A9B-4CAD-AFD8-0324FDF4E0F1}
AppName=Phrase CLI
AppVersion=1.17.1
AppPublisher=Phrase GmbH
AppPublisherURL=https://phrase.com/cli
AppSupportURL=https://phrase.com/cli
AppUpdatesURL=https://phrase.com/cli
DefaultDirName={pf}\Phrase
DefaultGroupName=Phrase-CLI
DisableProgramGroupPage=yes
LicenseFile=LICENSE
InfoAfterFile=postinstall.rtf
OutputBaseFilename=phrase_setup_386
SetupIconFile=parrot.ico
Compression=lzma
SolidCompression=yes
WizardStyle=modern

[Files]
Source: "../../dist/phrase_windows_386.exe"; DestDir: "{app}"; DestName: "phrase.exe"; Flags: ignoreversion
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Registry]
Root: HKLM; Subkey: "SYSTEM\CurrentControlSet\Control\Session Manager\Environment"; \
    ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; \

[Setup]
AlwaysRestart = yes

[Icons]
Name: "{group}\Phrase CLI"; Filename: "{app}"
Name: "{group}\{cm:UninstallProgram,Phrase CLI}"; Filename: "{uninstallexe}"
