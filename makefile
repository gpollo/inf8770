BINARIES=bin/arithmetic bin/dictionnary

TESTS= \
	test/binary/10 \
	test/binary/100 \
	test/binary/1000 \
	test/binary/4000 \
	test/letters/2 \
	test/letters/3 \
	test/letters/4 \
	test/letters/5 \
	test/letters/10 \
	test/letters/20 \
	test/other/pattern \
	test/text/short \
	test/text/average \
	test/text/long

all: $(BINARIES)

bin/arithmetic: $(wildcard arithmetic/*.go)
	go build -o $@ $?

bin/dictionnary: $(wildcard dictionnary/*.go)
	go build -o $@ $?

encode: $(BINARIES) \
	$(TESTS:=.encoded.arithmetic) \
	$(TESTS:=.encoded.dictionnary)

decode: encode \
	$(TESTS:=.decoded.arithmetic) \
	$(TESTS:=.decoded.dictionnary)

check: decode \
	$(TESTS:=.decoded.arithmetic.check) \
	$(TESTS:=.decoded.dictionnary.check)

clean:
	rm -vf $(BINARIES)

mrproper: clean
	find -type f -iname "*.arithmetic"  -print0 | xargs -0 rm -vf
	find -type f -iname "*.dictionnary" -print0 | xargs -0 rm -vf

##############
# Arithmetic #
##############

$(TESTS:=.encoded.arithmetic):
	$(eval src := $(@:.encoded.arithmetic=.decoded))
	$(eval dst := $@)
	cat $(src) | bin/arithmetic --encode --parallel --workers 16 > $(dst)

$(TESTS:=.decoded.arithmetic): 
	$(eval src := $(@:.decoded.arithmetic=.encoded.arithmetic))
	$(eval dst := $@)
	cat $(src) | bin/arithmetic --decode --parallel --workers 16 > $(dst)

$(TESTS:=.decoded.arithmetic.check): 
	$(eval file1 := $(@:.check=))
	$(eval file2 := $(@:.arithmetic.check=))
	@if ! cmp --silent $(file1) $(file2); then \
		echo "error: $(file1) and $(file2) are different"; \
	else \
		echo "ok: $(file1) and $(file2) are identical"; \
	fi

###############
# Dictionnary #
###############

$(TESTS:=.encoded.dictionnary):
	$(eval SRC := $(@:.encoded.dictionnary=.decoded))
	$(eval DST := $@)
	cat $(SRC) | bin/dictionnary --encode > $(DST)

$(TESTS:=.decoded.dictionnary): 
	$(eval SRC := $(@:.decoded.dictionnary=.encoded.dictionnary))
	$(eval DST := $@)
	cat $(SRC) | bin/dictionnary --decode > $(DST)

$(TESTS:=.decoded.dictionnary.check):
	$(eval file1 := $(@:.check=))
	$(eval file2 := $(@:.dictionnary.check=))
	@if ! cmp --silent $(file1) $(file2); then \
		echo "error: $(file1) and $(file2) are differential"; \
	else \
		echo "ok: $(file1) and $(file2) are identical"; \
	fi
