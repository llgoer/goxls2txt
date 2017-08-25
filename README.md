# goxls2txt

## Description

Package for golang to convert excel xls files to text.It's a go binding for [hroptatyr/xls2txt](https://github.com/hroptatyr/xls2txt)


## Installation

This package can be installed with the go get command:

```go get github.com/yanghuxiao/goxls2txt```

goxls2txt is cgo package. If you want to build your app using goxls2txt, you need gcc. 

## Documentation

This package is easy to use

### convert xls to text

```go
gx2t := NewGoXls2Txt()
_, err := gx2t.Xls2txt("Workbook1.xls", "Workbook1.txt")
if err != nil {
	fmt.Println("ok!")	
}
```

## Author

YangHuxiao

## Thanks for

Jan Bobrowski

Sebastian Freundt