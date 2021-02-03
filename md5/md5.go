package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// CalcFileMD5 to get md5-32
func CalcFileMD5(filename string) string {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return "ErrorOpenFile"
	}
	defer f.Close()

	md5Handle := md5.New()         //创建 md5 句柄
	_, err = io.Copy(md5Handle, f) //将文件内容拷贝到 md5 句柄中
	if nil != err {
		fmt.Println(err)
		return "ErrorCpenFile"
	}
	md := md5Handle.Sum(nil)        //计算 MD5 值，返回 []byte
	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str
}

func main() {
	path := ""
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		fmt.Println("传入文件数量有误！目前只能计算一个md5")
		return
	}
	result := CalcFileMD5(path)
	fmt.Println(result)
}
