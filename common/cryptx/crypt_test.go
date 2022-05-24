package cryptx

import (
	"fmt"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	salt := "HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe"
	password := "123456"
	encrypt := PasswordEncrypt(salt, password)
	fmt.Println(encrypt)
}