// Copyright (C) 2017 Yanghuxiao <yanghuxiao@vip.qq.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package goxls2txt

/*
#cgo CFLAGS: -O2 -g -Wall
#cgo LDFLAGS: -lm
#include "xls2txt.h"
#include <stdio.h>
#ifdef linux
# include <getopt.h>
#endif
#include <time.h>
#include <math.h>

*/
import "C"

import (
	"errors"
)

type GoXls2Txt struct {
}

func NewGoXls2Txt() *GoXls2Txt {
	return &GoXls2Txt{}
}

func (gx2t *GoXls2Txt) Xls2txt(src, dst string) (bool, error) {
	if src == "" || dst == "" {
		return false, nil
	}

	srcname := C.CString(src)
	dstname := C.CString(dst)
	rs := C.xls2cvs(srcname, dstname)
	if rs == 1 {
		return true, nil
	} else {
		return false, errors.New("check excel xls exits")
	}
}
