package gomisc

import (
	"fmt"
	"testing"
)

func TestIntSliceUnique(t *testing.T) {
	s := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}

	fmt.Println("origin slice is: ", s)

	s = IntSliceUnique(s)

	fmt.Println("after call slice is: ", s)
}

func TestStringSliceUnique(t *testing.T) {
	s := []string{"a", "ab", "ab", "abc", "abc", "abc", "abcd", "abcd", "abcd", "abcd", "abcd"}

	fmt.Println("origin slice is: ", s)

	s = StringSliceUnique(s)

	fmt.Println("after call slice is: ", s)
}

func TestFileExist(t *testing.T) {
	f := "/etc/passwd"

	r := FileExist(f)
	if r {
		fmt.Println(f, "is exist")
	} else {
		fmt.Println(f, "is not exist")
	}
}

func TestDirExist(t *testing.T) {
	d := "/home/ligang/devspace"

	r := DirExist(d)
	if r {
		fmt.Println(d, "is exist")
	} else {
		fmt.Println(d, "is not exist")
	}
}

func TestAppendBytes(t *testing.T) {
	b := []byte("abc")
	b = AppendBytes(b, []byte("def"), []byte("ghi"))

	fmt.Println(string(b))
}

func TestListFilesInDir(t *testing.T) {
	fileList, err := ListFilesInDir("/home/ligang/tmp")
	if err != nil {
		t.Log(err)
		return
	}

	for _, path := range fileList {
		t.Log(path)
	}
}

func TestSaveParseJsonFile(t *testing.T) {
	filePath := "/tmp/test_save_parse_json_file.json"

	v1 := make(map[string]string)
	v1["k1"] = "a"
	v1["k2"] = "b"
	v1["k3"] = "c"

	err := SaveJsonFile(filePath, v1)
	if err != nil {
		t.Error("save json file failed: " + err.Error())
	}

	v2 := make(map[string]string)
	err = ParseJsonFile(filePath, &v2)
	if err != nil {
		t.Error("parse json file failed: " + err.Error())
	}

	for k, v := range v2 {
		if v != v1[k] {
			t.Error("save parse json file error, k: " + k + " not equal")
		}
	}
}

func TestSubString(t *testing.T) {
	s := "abcdefg"

	_, err := SubString(s, 3, 20)
	t.Log(err)

	_, err = SubString(s, 10, 3)
	t.Log(err)

	ss, _ := SubString(s, 3, 4)
	t.Log(ss)
}
