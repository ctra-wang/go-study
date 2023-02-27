package services

import "github.com/zalando/go-keyring"

const (
	serviceName = "go-gpt3-chat-cli"
)

// go-keyring是一个与操作系统无关的库，用于设置，获取和删除系统密钥环中的机密。 它支持OS X ， Linux（dbus）和Windows 。
// OS X的实现依赖于/usr/bin/security二进制文件与OS X的钥匙串连接。 默认情况下应可用。

// demo 详见 下面项目：kering

var keyMagNow *KeyMag

// KeyMag 是从service+name 换取key，主要参数就这三个
type KeyMag struct {
	Service string
}

func (k *KeyMag) ClearKey() {

}

func (k *KeyMag) GetKey(user string) string {
	key, err := keyring.Get(k.Service, user)
	if err != nil {
		return ""
	}
	return key
}

func (k *KeyMag) SetKey(user string, pwd string) {
	err := keyring.Set(k.Service, user, pwd)
	if err != nil {
		panic(err)
	}
}

func (k *KeyMag) DelKey(user string) {
	err := keyring.Delete(k.Service, user)
	if err != nil {
		panic(err)
	}
}

func NewKeyMag() {
	keyMagNow = &KeyMag{Service: serviceName}
}

func InitKeyMag() *KeyMag {
	if keyMagNow == nil {
		NewKeyMag()
	}
	return keyMagNow
}

type (
	KeyMagInterface interface {
		GetKey(user string) string
		SetKey(user string, pwd string)
		DelKey(user string)
		ClearKey()
	}
)

var _ KeyMagInterface = &KeyMag{}
