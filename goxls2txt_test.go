// Copyright (C) 2017 Yanghuxiao <yanghuxiao@vip.qq.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package goxls2txt

import (
	"io/ioutil"
	"strings"
	"testing"
)

var gx2t = NewGoXls2Txt()

func Test_Xls2txt(t *testing.T) {
	_, err := gx2t.Xls2txt("Workbook1.xls", "Workbook1.txt")
	if err != nil {
		t.Error("Test GoXls2Txt Xls2txt not pass")
	} else {
		t.Log("GoXls2Txt Xls2txt pass successfull")
	}
	dat, err := ioutil.ReadFile("Workbook1.txt")
	if "Wetter	Sonne	Regen	Schnee" == strings.TrimSpace(string(dat)) {
		t.Log("GoXls2Txt Xls2txt Parse pass successfull")
	} else {
		t.Error("Test GoXls2Txt Xls2txt Parse not pass")
	}

}
