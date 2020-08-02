package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
1. Добавить текущаю директорию. Если приложение не расчитано на работу с внешними файлами, нечего давать возмоджность
2. Возвращение кода отличного от 0 при ошибках
*/

func main() {
	body, err := ReadAppFile("/tmp/somefile.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	fmt.Println(body)
}

func ReadAppFile(filename string) (body string, err error) {
	CurrentPath, e := os.Executable()
	if e != nil {
		return "", e
	}
	splitted := strings.Split(CurrentPath, "/")
	splitted = splitted[:len(splitted)-1]
	CurrentPath = strings.Join(splitted, "/")
	bodyByte, err := ioutil.ReadFile(CurrentPath + filename)
	return string(bodyByte), err
}
