package fileops

import (
	"os/user"
	"runtime"
)

func IsAdmin() bool {
	currentUser, err := user.Current()
	if err != nil {
		return false // if the user can not be verrified as admin, it defaults to false (access denied)
	}
	if runtime.GOOS == "windows" {
		return currentUser.Gid == "S-1-5-32-544" // the local Administrators group in Windows
	}
	return currentUser.Uid == "0" // admin check for Unix-like systems
}
