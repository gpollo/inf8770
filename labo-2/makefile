export JPEG2000_PYTHON=$(PWD)/jpeg2000/python

JPEG2000=bin/jpeg2000

all: $(JPEG2000)

$(JPEG2000): $(shell find . -type f -name "*.go")
	mkdir -p bin/
	cd jpeg2000/main && go build
	mv jpeg2000/main/main $@

##########
# tests  #
##########

ORIGINAL=tests/1-original
ENCODED=tests/2-encoded
DECODED=tests/3-decoded

tests: subsampling-yuv-tests \
       subsampling-rgb-tests \
       haar-wavelet-tests \
       deadzone-width-tests \
       deadzone-delta-tests \
       deadzone-various-tests \
       midthread-delta-tests \
       mixed-tests \
       various-tests

###########################
# subsampling (YUV) tests #
###########################

$(ENCODED)/yuv410.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 410 --wavelet haar:4

$(ENCODED)/yuv420.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 420 --wavelet haar:4

$(ENCODED)/yuv422.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 422 --wavelet haar:4

$(ENCODED)/yuv444.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 444 --wavelet haar:4

$(DECODED)/yuv410.png: $(ENCODED)/yuv410.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/yuv420.png: $(ENCODED)/yuv420.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/yuv422.png: $(ENCODED)/yuv422.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/yuv444.png: $(ENCODED)/yuv444.data
	$(JPEG2000) decode -i $< -o $@

subsampling-yuv-tests: $(DECODED)/yuv410.png \
                       $(DECODED)/yuv420.png \
                       $(DECODED)/yuv422.png \
                       $(DECODED)/yuv444.png
.PHONY: subsampling-yuv-tests

###########################
# subsampling (RGB) tests #
###########################

$(ENCODED)/rgb410.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 410 --wavelet haar:4 --no-conversion

$(ENCODED)/rgb420.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 420 --wavelet haar:4 --no-conversion

$(ENCODED)/rgb422.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 422 --wavelet haar:4 --no-conversion

$(ENCODED)/rgb444.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 444 --wavelet haar:4 --no-conversion

$(DECODED)/rgb410.png: $(ENCODED)/rgb410.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/rgb420.png: $(ENCODED)/rgb420.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/rgb422.png: $(ENCODED)/rgb422.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/rgb444.png: $(ENCODED)/rgb444.data
	$(JPEG2000) decode -i $< -o $@

subsampling-rgb-tests: $(DECODED)/rgb410.png \
                       $(DECODED)/rgb420.png \
                       $(DECODED)/rgb422.png \
                       $(DECODED)/rgb444.png
.PHONY: subsampling-rgb-tests

################################
# haar wavelet recursion tests #
################################

$(ENCODED)/haar-1.data: $(ORIGINAL)/claude3.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet haar:1

$(ENCODED)/haar-2.data: $(ORIGINAL)/claude3.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet haar:2

$(ENCODED)/haar-4.data: $(ORIGINAL)/claude3.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet haar:4

$(ENCODED)/haar-8.data: $(ORIGINAL)/claude3.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet haar:8

$(ENCODED)/haar-10.data: $(ORIGINAL)/claude3.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet haar:10

$(DECODED)/haar-1.png: $(ENCODED)/haar-1.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/haar-2.png: $(ENCODED)/haar-2.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/haar-4.png: $(ENCODED)/haar-4.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/haar-8.png: $(ENCODED)/haar-8.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/haar-10.png: $(ENCODED)/haar-10.data
	$(JPEG2000) decode -i $< -o $@

haar-wavelet-tests: $(DECODED)/haar-1.png \
                    $(DECODED)/haar-2.png \
                    $(DECODED)/haar-4.png \
                    $(DECODED)/haar-8.png \
                    $(DECODED)/haar-10.png
.PHONY: haar-wavelet-tests

########################
# deadzone width tests #
########################

$(ENCODED)/deadzone-width-1.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:1:1:0.0

$(ENCODED)/deadzone-width-64.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:64:1:0.0

$(ENCODED)/deadzone-width-128.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:128:1:0.0

$(ENCODED)/deadzone-width-256.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:256:1:0.0

$(ENCODED)/deadzone-width-512.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:512:1:0.0

$(ENCODED)/deadzone-width-1024.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:1024:1:0.0

$(DECODED)/deadzone-width-1.png: $(ENCODED)/deadzone-width-1.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-width-64.png: $(ENCODED)/deadzone-width-64.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-width-128.png: $(ENCODED)/deadzone-width-128.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-width-256.png: $(ENCODED)/deadzone-width-256.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-width-512.png: $(ENCODED)/deadzone-width-512.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-width-1024.png: $(ENCODED)/deadzone-width-1024.data
	$(JPEG2000) decode -i $< -o $@

deadzone-width-tests: $(DECODED)/deadzone-width-1.png \
                      $(DECODED)/deadzone-width-64.png \
                      $(DECODED)/deadzone-width-128.png \
                      $(DECODED)/deadzone-width-256.png \
                      $(DECODED)/deadzone-width-512.png \
                      $(DECODED)/deadzone-width-1024.png
.PHONY: deadzone-width-tests

#########################
# dead zone delta tests #
#########################

$(ENCODED)/deadzone-delta-1.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:1:1:0.0

$(ENCODED)/deadzone-delta-64.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:1:64:0.0

$(ENCODED)/deadzone-delta-128.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:1:128:0.0

$(ENCODED)/deadzone-delta-256.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:1:256:0.0

$(ENCODED)/deadzone-delta-512.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:1:512:0.0

$(ENCODED)/deadzone-delta-1024.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:1:1024:0.0

$(DECODED)/deadzone-delta-1.png: $(ENCODED)/deadzone-delta-1.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-delta-64.png: $(ENCODED)/deadzone-delta-64.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-delta-128.png: $(ENCODED)/deadzone-delta-128.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-delta-256.png: $(ENCODED)/deadzone-delta-256.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-delta-512.png: $(ENCODED)/deadzone-delta-512.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-delta-1024.png: $(ENCODED)/deadzone-delta-1024.data
	$(JPEG2000) decode -i $< -o $@

deadzone-delta-tests: $(DECODED)/deadzone-delta-1.png \
                      $(DECODED)/deadzone-delta-64.png \
                      $(DECODED)/deadzone-delta-128.png \
                      $(DECODED)/deadzone-delta-256.png \
                      $(DECODED)/deadzone-delta-512.png \
                      $(DECODED)/deadzone-delta-1024.png
.PHONY: deadzone-delta-tests

###########################
# dead zone various tests #
###########################

$(ENCODED)/deadzone-various-128-64.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:128:64:0.0

$(ENCODED)/deadzone-various-128-128.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier deadzone:128:128:0.0

$(DECODED)/deadzone-various-128-64.png: $(ENCODED)/deadzone-various-128-64.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/deadzone-various-128-128.png: $(ENCODED)/deadzone-various-128-128.data
	$(JPEG2000) decode -i $< -o $@

deadzone-various-tests: $(DECODED)/deadzone-various-128-64.png \
                        $(DECODED)/deadzone-various-128-128.png
.PHONY: deadzone-various-tests

##########################
# mid-thread delta tests #
##########################

$(ENCODED)/midthread-delta-1.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier midthread:1

$(ENCODED)/midthread-delta-64.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier midthread:64

$(ENCODED)/midthread-delta-128.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier midthread:128

$(ENCODED)/midthread-delta-256.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier midthread:256

$(ENCODED)/midthread-delta-512.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier midthread:512

$(ENCODED)/midthread-delta-1024.data: $(ORIGINAL)/claude2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --quantifier midthread:1024

$(DECODED)/midthread-delta-1.png: $(ENCODED)/midthread-delta-1.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/midthread-delta-64.png: $(ENCODED)/midthread-delta-64.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/midthread-delta-128.png: $(ENCODED)/midthread-delta-128.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/midthread-delta-256.png: $(ENCODED)/midthread-delta-256.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/midthread-delta-512.png: $(ENCODED)/midthread-delta-512.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/midthread-delta-1024.png: $(ENCODED)/midthread-delta-1024.data
	$(JPEG2000) decode -i $< -o $@

midthread-delta-tests: $(DECODED)/midthread-delta-1.png \
                       $(DECODED)/midthread-delta-64.png \
                       $(DECODED)/midthread-delta-128.png \
                       $(DECODED)/midthread-delta-256.png \
                       $(DECODED)/midthread-delta-512.png \
                       $(DECODED)/midthread-delta-1024.png
.PHONY: midthread-delta-tests

###############
# mixed tests #
###############

$(ENCODED)/mixed-h1-dz512.data: $(ORIGINAL)/claude4.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet haar:1 --quantifier deadzone:512:512:0.0

$(ENCODED)/mixed-h1-mt512.data: $(ORIGINAL)/claude4.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet haar:1 --quantifier midthread:512

$(ENCODED)/mixed-h4-dz512.data: $(ORIGINAL)/claude4.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet haar:4 --quantifier deadzone:512:512:0.0

$(ENCODED)/mixed-h4-mt512.data: $(ORIGINAL)/claude4.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet haar:4 --quantifier midthread:512

$(ENCODED)/mixed-daub-dz512.data: $(ORIGINAL)/claude4.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet daubechies:1:2 --quantifier deadzone:512:512:0.0

$(ENCODED)/mixed-daub-mt512.data: $(ORIGINAL)/claude4.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --wavelet daubechies:1:2 --quantifier midthread:512

$(DECODED)/mixed-h1-dz512.png: $(ENCODED)/mixed-h1-dz512.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/mixed-h1-mt512.png: $(ENCODED)/mixed-h1-mt512.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/mixed-h4-dz512.png: $(ENCODED)/mixed-h4-dz512.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/mixed-h4-mt512.png: $(ENCODED)/mixed-h4-mt512.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/mixed-daub-dz512.png: $(ENCODED)/mixed-daub-dz512.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/mixed-daub-mt512.png: $(ENCODED)/mixed-daub-mt512.data
	$(JPEG2000) decode -i $< -o $@

mixed-tests: $(DECODED)/mixed-h1-dz512.png \
             $(DECODED)/mixed-h1-mt512.png \
             $(DECODED)/mixed-h4-dz512.png \
             $(DECODED)/mixed-h4-mt512.png \
             $(DECODED)/mixed-daub-dz512.png \
             $(DECODED)/mixed-daub-mt512.png
.PHONY: mixed-tests

#######################
# various image tests #
#######################

OPTIMAL_SETTINGS=--subsampling 410 \
                 --wavelet daubechies:1:2 \
                 --quantifier deadzone:256:256:0.0

$(ENCODED)/various-black.data: $(ORIGINAL)/black.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ $(OPTIMAL_SETTINGS)

$(ENCODED)/various-colorful.data: $(ORIGINAL)/colorful.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ $(OPTIMAL_SETTINGS)

$(ENCODED)/various-dark.data: $(ORIGINAL)/dark.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ $(OPTIMAL_SETTINGS)

$(ENCODED)/various-normal.data: $(ORIGINAL)/normal.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ $(OPTIMAL_SETTINGS)

$(ENCODED)/various-gradient1.data: $(ORIGINAL)/gradient1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ $(OPTIMAL_SETTINGS)

$(ENCODED)/various-gradient2.data: $(ORIGINAL)/gradient2.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ $(OPTIMAL_SETTINGS)

$(ENCODED)/various-gradient3.data: $(ORIGINAL)/gradient3.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ $(OPTIMAL_SETTINGS)

$(ENCODED)/various-white.data: $(ORIGINAL)/white.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ $(OPTIMAL_SETTINGS)

$(DECODED)/various-black.png: $(ENCODED)/various-black.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/various-colorful.png: $(ENCODED)/various-colorful.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/various-dark.png: $(ENCODED)/various-dark.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/various-normal.png: $(ENCODED)/various-normal.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/various-gradient1.png: $(ENCODED)/various-gradient1.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/various-gradient2.png: $(ENCODED)/various-gradient2.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/various-gradient3.png: $(ENCODED)/various-gradient3.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/various-white.png: $(ENCODED)/various-white.data
	$(JPEG2000) decode -i $< -o $@

various-tests: $(DECODED)/various-black.png \
               $(DECODED)/various-colorful.png \
               $(DECODED)/various-dark.png \
               $(DECODED)/various-normal.png \
               $(DECODED)/various-gradient1.png \
               $(DECODED)/various-gradient2.png \
               $(DECODED)/various-gradient3.png \
               $(DECODED)/various-white.png
.PHONY: various-tests

##########
# tables #
##########

table/stats-subsampling.asc: export TABLE_TITLE=Comparaison de différentes méthodes de sous-échantillonnage
table/stats-subsampling.asc: export ADDITIONAL_COLUMNS=colorspace
table/stats-subsampling.asc: export ORIGINAL_IMAGE=$(ORIGINAL)/claude1.bmp
table/stats-subsampling.asc: scripts/compress-stats.sh \
                             $(ENCODED)/yuv410.data \
                             $(ENCODED)/yuv420.data \
                             $(ENCODED)/yuv422.data \
                             $(ENCODED)/yuv444.data \
                             $(ENCODED)/rgb410.data \
                             $(ENCODED)/rgb420.data \
                             $(ENCODED)/rgb444.data \
                             $(ENCODED)/rgb422.data
	$< $^ > $@

table/stats-haar-recursion.asc: export TABLE_TITLE=Comparaison de différentes récursions pour les ondelettes de Haar
table/stats-haar-recursion.asc: export ADDITIONAL_COLUMNS=haar_recursion
table/stats-haar-recursion.asc: export ORIGINAL_IMAGE=$(ORIGINAL)/claude3.bmp
table/stats-haar-recursion.asc: scripts/compress-stats.sh \
                                $(ENCODED)/haar-1.data \
                                $(ENCODED)/haar-2.data \
                                $(ENCODED)/haar-4.data \
                                $(ENCODED)/haar-8.data \
                                $(ENCODED)/haar-10.data
	$< $^ > $@

table/stats-deadzone-width.asc: export TABLE_TITLE=Comparaison de différentes largeurs de zone morte
table/stats-deadzone-width.asc: export ADDITIONAL_COLUMNS=deadzone_width
table/stats-deadzone-width.asc: export ORIGINAL_IMAGE=$(ORIGINAL)/claude2.bmp
table/stats-deadzone-width.asc: scripts/compress-stats.sh \
                                $(ENCODED)/deadzone-width-1.data \
                                $(ENCODED)/deadzone-width-64.data \
                                $(ENCODED)/deadzone-width-128.data \
                                $(ENCODED)/deadzone-width-256.data \
                                $(ENCODED)/deadzone-width-512.data \
                                $(ENCODED)/deadzone-width-1024.data
	$< $^ > $@

table/stats-deadzone-delta.asc: export TABLE_TITLE=Comparaison de différents pas pour la quantification à zone morte
table/stats-deadzone-delta.asc: export ADDITIONAL_COLUMNS=deadzone_delta
table/stats-deadzone-delta.asc: export ORIGINAL_IMAGE=$(ORIGINAL)/claude2.bmp
table/stats-deadzone-delta.asc: scripts/compress-stats.sh \
                                $(ENCODED)/deadzone-delta-1.data \
                                $(ENCODED)/deadzone-delta-64.data \
                                $(ENCODED)/deadzone-delta-128.data \
                                $(ENCODED)/deadzone-delta-256.data \
                                $(ENCODED)/deadzone-delta-512.data \
                                $(ENCODED)/deadzone-delta-1024.data
	$< $^ > $@

table/stats-midthread-delta.asc: export TABLE_TITLE=Comparaison de différents pas pour la quantification à moyenne
table/stats-midthread-delta.asc: export ADDITIONAL_COLUMNS=midthread_delta
table/stats-midthread-delta.asc: export ORIGINAL_IMAGE=$(ORIGINAL)/claude2.bmp
table/stats-midthread-delta.asc: scripts/compress-stats.sh \
                                 $(ENCODED)/midthread-delta-1.data \
                                 $(ENCODED)/midthread-delta-64.data \
                                 $(ENCODED)/midthread-delta-128.data \
                                 $(ENCODED)/midthread-delta-256.data \
                                 $(ENCODED)/midthread-delta-512.data \
                                 $(ENCODED)/midthread-delta-1024.data
	$< $^ > $@

table/stats-mixed.asc: export TABLE_TITLE=Comparaison de combinaisons d'ondelette et quantificateur différents
#table/stats-mixed.asc: export ADDITIONAL_COLUMNS=midthread_delta TODO
table/stats-mixed.asc: export ORIGINAL_IMAGE=$(ORIGINAL)/claude4.bmp
table/stats-mixed.asc: scripts/compress-stats.sh \
                       $(ENCODED)/mixed-h1-dz512.data \
                       $(ENCODED)/mixed-h1-mt512.data \
                       $(ENCODED)/mixed-h4-dz512.data \
                       $(ENCODED)/mixed-h4-mt512.data \
                       $(ENCODED)/mixed-daub-dz512.data \
                       $(ENCODED)/mixed-daub-mt512.data
	$< $^ > $@

table/stats-various-images.asc: export TABLE_TITLE=Comparaison des différents types d'images utilisées
table/stats-various-images.asc: export ORIGINAL_IMAGE=$(ORIGINAL)/black.bmp
table/stats-various-images.asc: scripts/compress-stats.sh \
                                $(ENCODED)/various-black.data \
                                $(ENCODED)/various-colorful.data \
                                $(ENCODED)/various-dark.data \
                                $(ENCODED)/various-normal.data \
                                $(ENCODED)/various-gradient1.data \
                                $(ENCODED)/various-gradient2.data \
                                $(ENCODED)/various-gradient3.data \
                                $(ENCODED)/various-white.data
	$< $^ > $@

table/various-images.asc: export TABLE_TITLE=Liste des différents types d'images utilisées
table/various-images.asc: scripts/table-images.sh \
                          $(DECODED)/various-black.png \
                          $(DECODED)/various-colorful.png \
                          $(DECODED)/various-dark.png \
                          $(DECODED)/various-normal.png \
                          $(DECODED)/various-gradient1.png \
                          $(DECODED)/various-gradient2.png \
                          $(DECODED)/various-gradient3.png \
                          $(DECODED)/various-white.png
	$< $^ > $@

############
# document #
############

rapport.pdf: README.asc \
             themes/basic-theme.yml \
             table/stats-subsampling.asc \
             table/stats-haar-recursion.asc \
             table/stats-deadzone-width.asc \
             table/stats-deadzone-delta.asc \
             table/stats-midthread-delta.asc \
             table/stats-mixed.asc \
             table/stats-various-images.asc \
             table/various-images.asc \
             $(DECODED)/yuv410.png \
             $(DECODED)/yuv444.png \
             $(DECODED)/rgb410.png \
             $(DECODED)/haar-1.png \
             $(DECODED)/haar-10.png \
             $(DECODED)/deadzone-width-1.png \
             $(DECODED)/deadzone-width-256.png \
             $(DECODED)/deadzone-width-1024.png \
             $(DECODED)/deadzone-delta-1.png \
             $(DECODED)/deadzone-delta-256.png \
             $(DECODED)/deadzone-delta-1024.png \
             $(DECODED)/midthread-delta-1.png \
             $(DECODED)/midthread-delta-256.png \
             $(DECODED)/midthread-delta-1024.png \
             $(DECODED)/mixed-h1-dz512.png \
             $(DECODED)/mixed-h1-mt512.png \
             $(DECODED)/mixed-h4-dz512.png \
             $(DECODED)/mixed-h4-mt512.png \
             $(DECODED)/mixed-daub-dz512.png \
             $(DECODED)/mixed-daub-mt512.png
	asciidoctor-pdf -r asciidoctor-mathematical $< -a pdf-themesdir=$(PWD)/themes -a pdf-theme=basic
	mv README.pdf rapport.pdf

#################
# other targets #
#################

clean:
	rm -vf $(JPEG2000)
	rm -vf $(ENCODED)/*
	rm -vf $(DECODED)/*

distclean: clean
	rm -vf table/*
