package string

import (
	"fmt"
	"strings"
)

// 别名
var p = fmt.Println

func main() {
	p("contains:", strings.Contains("test", "es"))
	p("count:", strings.Count("test", "t"))
	p("hasPrefix:", strings.HasPrefix("test", "te"))
	p("hasSuffix:", strings.HasSuffix("test", "st"))
	p("index:", strings.Index("test", "s"))
	p("join:", strings.Join([]string{"a", "b"}, "-"))
	//重复
	p("repeat:", strings.Repeat("test", 5))
	// n : 要替换的个数,-1代表替换所有
	p("replace:", strings.Replace("test", "t", "T", 1))
	p("replace:", strings.Replace("test", "t", "T", -1))
	p("split:", strings.Split("t,e,s,t", ","))
	p("toLower:", strings.ToLower("TEST"))
	p("toUpper:", strings.ToUpper("test"))
	p()
	p("len:", len("test"))
	p("char", "test"[1])
}
