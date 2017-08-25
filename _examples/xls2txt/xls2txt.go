package main

import (
	"fmt"

	"github.com/yanghuxiao/goxls2txt"
)

func main() {
	fmt.Println("convert start")
	gx2t := goxls2txt.NewGoXls2Txt()
	_, err := gx2t.Xls2txt("../../Workbook1.xls", "../../Workbook1.txt")
	if err == nil {
		fmt.Println("convert xls to txt successfull")
	}
}
