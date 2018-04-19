/**
* @file color.go
* @brief set text color
* @author ligang
* @version
* @date 2016-02-06
 */

package color

import (
	"bytes"
)

func Black(msg []byte) []byte {
	return Color([]byte("\033[01;0m"), msg)
}

func Red(msg []byte) []byte {
	return Color([]byte("\033[01;31m"), msg)
}

func Green(msg []byte) []byte {
	return Color([]byte("\033[01;32m"), msg)
}

func Yellow(msg []byte) []byte {
	return Color([]byte("\033[01;33m"), msg)
}

func Blue(msg []byte) []byte {
	return Color([]byte("\033[01;34m"), msg)
}

func Maganta(msg []byte) []byte {
	return Color([]byte("\033[01;35m"), msg)
}

func Cyan(msg []byte) []byte {
	return Color([]byte("\033[01;36m"), msg)
}

func White(msg []byte) []byte {
	return Color([]byte("\033[01;37m"), msg)
}

func Color(color []byte, msg []byte) []byte {
	buf := bytes.NewBuffer(color)

	buf.Write(msg)
	buf.Write([]byte("\033[0m"))

	return buf.Bytes()
}
