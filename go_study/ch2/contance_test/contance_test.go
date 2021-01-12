package contance_test

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)
const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestContest(t *testing.T)  {
	t.Log(Monday,Tuesday,Wednesday)
}

func TestContestTry(t *testing.T){
	a := 7  // 二进制 0111
	//做与运算
	t.Log(a&Readable==Readable,a&Writeable==Writeable,a&Executable==Executable)

}

