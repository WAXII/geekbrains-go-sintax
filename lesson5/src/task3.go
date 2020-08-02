package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Test struct {
	FieldA string
	FieldB string
	FieldC int
}

func main() {
	source, err := ReadAppFile("/tmp/somefile.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	result := make([]Test, 0)
	var reader = csv.NewReader(strings.NewReader(source))
	reader.Comma = ';'
	for {
		var test Test
		err := Unmarshal(reader, &test)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		result = append(result, test)
	}
	rB, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	fmt.Println(string(rB))
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

func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}

type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}
