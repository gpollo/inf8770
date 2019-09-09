BINARIES=bin/arithmetic

TESTS=                \
	test/binary/10    \
	test/binary/100   \
	test/binary/1000  \
	test/binary/4000  \
	test/text/short   \
	test/text/average \
	test/text/long

all: $(BINARIES)

encode: $(TESTS:=.arithmetic)

decode: $(TESTS:=.arithmetic.decoded)

bin/arithmetic: arithmetic/*.go
	go build -o $@ $?

$(TESTS:=.arithmetic): 
	cat $(@:.arithmetic=) | bin/arithmetic --encode --parallel --workers 16 > $@

$(TESTS:=.arithmetic.decoded): 
	cat $(@:.decoded=) | bin/arithmetic --decode --parallel --workers 16 > $@

