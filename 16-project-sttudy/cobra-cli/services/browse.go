package services

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenLinkInBrowser(url string) error {
	var err error
	//获取操作系统,根据不同操作系统打开网址方式
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
