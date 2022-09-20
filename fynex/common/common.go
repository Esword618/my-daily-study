package common

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// http://www.361way.com/golang-win-process/6105.html
// OpenWithDefaultApp 使用操作系统默认程序打开指定的文件
func OpenWithDefaultApp(path string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", path).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Start()
	case "darwin":
		err = exec.Command("open", path).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		// log.Fatal(err)
	}
}
func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
