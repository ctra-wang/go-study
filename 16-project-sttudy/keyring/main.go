package main

import (
	"github.com/zalando/go-keyring"
	"log"
)

// go-keyring是一个与操作系统无关的库，用于设置，获取和删除系统密钥环中的机密。 它支持OS X ， Linux（dbus）和Windows 。
// OS X的实现依赖于/usr/bin/security二进制文件与OS X的钥匙串连接。 默认情况下应可用。

func main() {
	service := "my-app"
	user := "anon"
	password := "secret123"

	// set password
	err := keyring.Set(service, user, password)
	if err != nil {
		log.Fatal(err)
	}

	// get password
	secret, err := keyring.Get(service, user)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(secret)
}
