package main

import (
	"fmt"

	"github.com/beautwc/pkg"
	thirdparty "github.com/beautwc/third_party"
)

func main() {
	linesCount := pkg.LinesCount{}
	err := linesCount.GetLinesCount("../../examples/text1.txt", "../../examples/text2.txt", "../../examples/text3.txt")
	if err != nil {
		fmt.Println(err)
	}
	prefix, value := thirdparty.GetValueColors()
	fmt.Println("total lines: ", linesCount.TotalLines)
	for _, lines := range linesCount.PerFileLines {
		p, v := prefix.Sprint(lines.Filename), value.Sprint(lines.Count)
		fmt.Printf("%s : %s\n", p, v)
	}
}
