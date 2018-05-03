// +build windows

/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package homedir

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// HomeDir returns the home directory for the current user
func HomeDir() string {
	if runtime.GOOS == "windows" {

		// SHGetKnownFolderPath should return a valid homedir but we have fallbacks if it does not
		if shellHome := readWindowsHomeDirFromShellAPI(); len(shellHome) > 0 {
			if _, err := os.Stat(shellHome); err == nil {
				return shellHome
			}
		}

		// First prefer the HOME environmental variable
		if home := os.Getenv("HOME"); len(home) > 0 {
			if _, err := os.Stat(home); err == nil {
				return home
			}
		}
		// Next try the USERPROFILE environmental variable
		if userProfile := os.Getenv("USERPROFILE"); len(userProfile) > 0 {
			if _, err := os.Stat(userProfile); err == nil {
				return userProfile
			}
		}
		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			homeDir := homeDrive + homePath
			if _, err := os.Stat(homeDir); err == nil {
				return homeDir
			}
		}
	}
	return os.Getenv("HOME")
}

func readWindowsHomeDirFromShellAPI() string {
	type guid struct {
		Data1 uint32
		Data2 uint16
		Data3 uint16
		Data4 [8]byte
	}
	var folderIDProfile = guid{0x5E6C858F, 0x0E22, 0x4760, [8]byte{0x9A, 0xFE, 0xEA, 0x33, 0x17, 0xB6, 0x71, 0x73}} // 5E6C858F-0E22-4760-9AFE-EA3317B67173

	procSHGetKnownFolderPath := windows.NewLazySystemDLL("Shell32.dll").NewProc("SHGetKnownFolderPath")
	procCoTaskMemFree := windows.NewLazySystemDLL("Ole32.dll").NewProc("CoTaskMemFree")

	var pszPath *uint16

	r0, _, _ := procSHGetKnownFolderPath.Call(uintptr(unsafe.Pointer(&folderIDProfile)), uintptr(0), uintptr(0), uintptr(unsafe.Pointer(&pszPath)))
	if r0 != 0 { // S_OK == 0
		return ""
	}

	defer syscall.Syscall(procCoTaskMemFree.Addr(), 1, uintptr(unsafe.Pointer(pszPath)), 0, 0)

	return syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(pszPath))[:])
}
