package color

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Println(string(Black([]byte("Black"))))
	fmt.Println(string(Red([]byte("Red"))))
	fmt.Println(string(Green([]byte("Green"))))
	fmt.Println(string(Yellow([]byte("Yellow"))))
	fmt.Println(string(Blue([]byte("Blue"))))
	fmt.Println(string(Maganta([]byte("Maganta"))))
	fmt.Println(string(Cyan([]byte("Cyan"))))
	fmt.Println(string(White([]byte("White"))))
}
