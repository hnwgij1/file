package file

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()        //关闭文件
	file.WriteString(content) //写入文件
}

func ReadFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileSize,_ := file.Seek(0, io.SeekEnd)
	content := make([]byte, fileSize)
	file.Seek(0, io.SeekStart)
	sumReadNum := 0
	for {
		readNum, err := file.Read(content)
		if err != nil && err != io.EOF {
			panic(err) //有错误抛出异常
		}
		if 0 == readNum {
			break //当读取完毕时候退出循环
		}
		sumReadNum += readNum
	}
	return string(content[:sumReadNum])
}
