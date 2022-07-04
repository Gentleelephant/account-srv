package biz

import (
	"crypto/md5"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"testing"
)

func TestPasswordEncode(t *testing.T) {

	options := password.Options{
		SaltLen:      256,
		Iterations:   128,
		KeyLen:       128,
		HashFunction: md5.New,
	}

	s, encode := password.Encode("hahaha", &options)
	fmt.Println("slat:", s)
	fmt.Println("encode:", encode)
	fmt.Println("slat length:", len(s))
	fmt.Println("encode length:", len(encode))
	//
	verify := password.Verify("hahaha", s, encode, &options)
	fmt.Println("verify:", verify)

}

func TestCustomPassword(t *testing.T) {
	encode, s := PasswordEncode("hahaha")
	fmt.Println("encode:", encode)
	fmt.Println("slat:", s)
}
