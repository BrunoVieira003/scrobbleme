!include "MUI2.nsh"

Name "Scrobbleme"
OutFile "ScrobblemeInstaller.exe"
InstallDir "$PROGRAMFILES\Scrobbleme"
RequestExecutionLevel admin

; =========================
; Installer Pages
; =========================
!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_PAGE_FINISH
!define MUI_HEADERIMAGE_BITMAP "icon.bmp"

!insertmacro MUI_LANGUAGE "English"

Section "Install"

	SetOutPath "$INSTDIR"

	; Compiled Go binary
	File "..\scrobbleme.exe"

	CreateDirectory "$APPDATA\Scrobbleme"
	FileOpen $0 "$APPDATA\Scrobbleme\config.json" w
	FileWrite $0 '{ "Session": { "name": "", "key": "" } }'
	FileClose $0

	; Add uninstall entry
	WriteUninstaller "$INSTDIR\uninstall.exe"

	; Store install path
	WriteRegStr HKLM "Software\Scrobbleme" "InstallDir" "$INSTDIR"

	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Scrobbleme" "DisplayName" "Scrobbleme"
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Scrobbleme" "Publisher" "BrunoVieira003"
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Scrobbleme" "DisplayVersion" "0.1.1"

	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Scrobbleme" "UninstallString" \
		"$INSTDIR\uninstall.exe"

	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Scrobbleme" "InstallLocation" "$INSTDIR"

	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Scrobbleme" "DisplayIcon" \
		"$INSTDIR\scrobbleme.exe"

	WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Scrobbleme" "NoModify" 1

	WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Scrobbleme" "NoRepair" 1

	; =========================
	; Context menu for .mp3
	; =========================

	; Add right-click entry
	WriteRegStr HKCR "SystemFileAssociations\.mp3\shell\Scrobbleme" "" "Scrobble me"
	WriteRegStr HKCR "SystemFileAssociations\.mp3\shell\Scrobbleme\command" "" '"$INSTDIR\scrobbleme.exe" "%1"'

SectionEnd

Section "Uninstall"

	Delete "$INSTDIR\scrobbleme.exe"
	Delete "$INSTDIR\uninstall.exe"
	RMDir "$INSTDIR"

	; remove app registry key
	DeleteRegKey HKLM "Software\Scrobbleme"

	; remove Windows Apps entry
	DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Scrobbleme"

	; remove context menu
	DeleteRegKey HKCR "SystemFileAssociations\.mp3\shell\Scrobbleme"

SectionEnd
