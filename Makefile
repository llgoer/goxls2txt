# gmake

NAME = xls2txt
VERSION = 0.15
BINDEST = /usr/local/lib
PKG=$(NAME)-$(VERSION)
FILES = Makefile xls2txt.[ch] ole.c cp.c ummap.[ch] ieee754.c list.h myerr.h

CFLAGS ?= -O2 -g -Wall
LDFLAGS = -lm

xls2txt: xls2txt.c ole.c cp.c ummap.c ieee754.c
	$(CC) $(CFLAGS) $^ -o $@.so -fPIC -shared $(LDFLAGS)
	go build

install: xls2txt
	install -s $< $(BINDEST)

clean:
	rm -f xls2txt $(addsuffix .o,$(basename $(filter %.c %.[ch],$(FILES))))
	rm -f xls2txt.so
	rm -f xls2csv

dist:
	ln -s . $(PKG)
	tar czf $(PKG).tar.gz --group=root --owner=root $(addprefix $(PKG)/, $(FILES)); \
	rm $(PKG)

check: xls2txt
	./$< -l Workbook1.xls
	./$< Workbook1.xls

.PHONY: install clean dist check
