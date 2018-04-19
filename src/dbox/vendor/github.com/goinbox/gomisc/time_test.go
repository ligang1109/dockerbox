package gomisc

import (
	"fmt"
	"testing"
	"time"
)

func TestRandByTime(t *testing.T) {
	tm := time.Now()
	fmt.Println(RandByTime(&tm), RandByTime(&tm), RandByTime(nil))
}
