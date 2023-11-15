package system

import (
	"os"
	"runtime"
)

func GetRuntimeSystemHostName() string {
	if runtime.GOOS == "windows" {
		name := os.Getenv("COMPUTERNAME")
		if name == "" {
			return "windows"
		}
		return name
	}

	hostname, _ := os.Hostname()

	if hostname == "" {
		return runtime.GOOS
	}
	return hostname
}
