# goxls2txt

Package for golang to convert excel xls files to text.It's a go binding for [hroptatyr/xls2txt](https://github.com/hroptatyr/xls2txt)


## Installation

This package can be installed with the go get command:

```go get github.com/yanghuxiao/goxls2txt```

goxls2txt is cgo package. If you want to build your app using goxls2txt, you need gcc. 

## Documentation

This package is easy to use，more example see [Examples](https://github.com/yanghuxiao/goxls2txt/tree/master/_examples)

### convert xls to text

```go
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

```

## Author

YangHuxiao

## Thanks for

Jan Bobrowski

Sebastian Freundt