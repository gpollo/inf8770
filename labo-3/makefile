DECOMPOSE=bin/decompose

all: $(DECOMPOSE)

$(DECOMPOSE): $(shell find . -type f -name "*.go")
	mkdir -p bin/
	cd decompose/main && go build
	mv decompose/main/main $@

############
# document #
############

rapport.pdf: README.asc
	asciidoctor-pdf $< -a pdf-themesdir=$(PWD)/themes -a pdf-theme=basic
	mv README.pdf rapport.pdf

#################
# other targets #
#################

clean:
	rm -vf $(DECOMPOSE)
.PHONY: clean

remise: rapport.pdf
	tar cvf - rapport.pdf makefile decompose | gzip -9 - > INF8770_1837776_LABO3.tar.gz
