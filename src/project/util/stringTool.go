package util

import (
	"crypto/md5"
	"fmt"
)

//将字符串加密成 md5
func Str2md5(str string) string  {
	data :=[]byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
