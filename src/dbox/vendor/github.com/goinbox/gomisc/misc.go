/**
* @file misc.go
* @brief misc supermarket
* @author ligang
* @date 2015-12-11
 */

package gomisc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func IntSliceUnique(s []int) []int {
	m := make(map[int]bool)
	r := make([]int, 0, cap(s))

	for _, k := range s {
		_, ok := m[k]
		if !ok {
			r = append(r, k)
			m[k] = true
		}
	}

	return r
}

func StringSliceUnique(s []string) []string {
	m := make(map[string]bool)
	r := make([]string, 0, cap(s))

	for _, k := range s {
		_, ok := m[k]
		if !ok {
			r = append(r, k)
			m[k] = true
		}
	}

	return r
}

func FileExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func DirExist(path string) bool {
	fi, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	if fi.IsDir() {
		return true
	}
	return false
}

func AppendBytes(b []byte, elems ...[]byte) []byte {
	buf := bytes.NewBuffer(b)
	for _, e := range elems {
		buf.Write(e)
	}

	return buf.Bytes()
}

func ListFilesInDir(rootDir string) ([]string, error) {
	rootDir = strings.TrimRight(rootDir, "/")
	if !DirExist(rootDir) {
		return nil, errors.New("Dir not exists")
	}

	var fileList []string
	dirList := []string{rootDir}

	for i := 0; i < len(dirList); i++ {
		curDir := dirList[i]
		file, err := os.Open(dirList[i])
		if err != nil {
			return nil, err
		}

		fis, err := file.Readdir(-1)
		if err != nil {
			return nil, err
		}

		for _, fi := range fis {
			path := curDir + "/" + fi.Name()
			if fi.IsDir() {
				dirList = append(dirList, path)
			} else {
				fileList = append(fileList, path)
			}
		}
	}

	return fileList, nil
}

func SaveJsonFile(filePath string, v interface{}) error {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, jsonBytes, 0644)
}

func ParseJsonFile(filePath string, v interface{}) error {
	if !FileExist(filePath) {
		return errors.New("confFile " + filePath + " not exists")
	}

	jsonBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonBytes, v)
}

func SubString(str string, start, cnt int) (string, error) {
	end := start + cnt
	if len(str) < end {
		return "", errors.New("end too long")
	}

	rs := []rune(str)

	return string(rs[start:end]), nil
}
