package monster

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func writeTofile(data interface{}, path string) error {
	bytes, _ := json.Marshal(data)
	//打开文件
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	//文件的写入
	n, err := f.Write(bytes)
	if err == nil && n < len(bytes) {
		err = io.ErrShortWrite
	}
	//关闭文件
	if err1 := f.Close(); err == nil {
		err = err1
	}
	log.Println("CHENGGON")
	return err
}
