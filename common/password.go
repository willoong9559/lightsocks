package common

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
)

type Lspasswd [256]byte

// 采用base64编码把密码转换为字符串
func Pta(lsPW *Lspasswd) string {
	return base64.StdEncoding.EncodeToString(lsPW[:])
}

// 解析采用base64编码的字符串获取密码
func Atp(password string) (*Lspasswd, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(password))
	if err != nil {
		return new(Lspasswd), errors.New("不合法的密码")
	}
	Lspasswd := new(Lspasswd)
	copy(Lspasswd[:], bs)
	return Lspasswd, nil
}

func NewRandPasswdStr() string {
	return Pta(NewRandPassword())
}

func GetDecodePasswdStr(lspasswd *Lspasswd) *Lspasswd {
	encodePasswd := new(Lspasswd)
	for i, v := range lspasswd {
		encodePasswd[v] = byte(i)
	}
	return encodePasswd
}

// 产生 256个byte随机组合的密码
func NewRandPassword() *Lspasswd {
	// 随机生成一个由  0~255 组成的 byte 数组
	intArr := rand.Perm(256)
	lspasswd := new(Lspasswd)
	for i, v := range intArr {
		lspasswd[i] = byte(v)
		if i == v {
			return NewRandPassword()
		}
	}
	return lspasswd
}
