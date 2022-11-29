package gmd5

import (
	"crypto/md5"
	"fmt"
	"rustdesk-api-server/utils/gconv"
)

// MD5 加密
func Encrypt(data interface{}) (encrypt string, err error) {
	return EncryptBytes(gconv.Bytes(data))
}
func EncryptNE(data interface{}) (encrypt string) {
	encrypt, _ = EncryptBytes(gconv.Bytes(data))
	return
}

// MD5 字节集计算
func EncryptBytes(bytes []byte) (encrypt string, err error) {
	m := md5.New()
	if _, err = m.Write(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", m.Sum(nil)), nil
}
