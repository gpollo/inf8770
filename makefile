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

encode: $(TESTS:=.encoded.arithmetic)

decode: encode $(TESTS:=.decoded.arithmetic)

bin/arithmetic: arithmetic/*.go
	go build -o $@ $?

$(TESTS:=.encoded.arithmetic):
	$(eval SRC := $(@:.encoded.arithmetic=.decoded))
	$(eval DST := $@)
	cat $(SRC) | bin/arithmetic --encode --parallel --workers 16 > $(DST)

$(TESTS:=.decoded.arithmetic): 
	$(eval SRC := $(@:.decoded.arithmetic=.encoded.arithmetic))
	$(eval DST := $@)
	cat $(SRC) | bin/arithmetic --decode --parallel --workers 16 > $(DST)

