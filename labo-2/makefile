export JPEG2000_PYTHON="$(PWD)/jpeg2000/python"

JPEG2000=bin/jpeg2000

all: $(JPEG2000)

$(JPEG2000): $(shell find . -type f -name "*.go")
	mkdir -p bin/
	cd jpeg2000/main && go build
	mv jpeg2000/main/main $@

#########
# test  #
#########

ORIGINAL=tests/1-original
ENCODED=tests/2-encoded
DECODED=tests/3-decoded

tests: subsampling-tests \
       deadzone-width-tests \
       deadzone-delta-tests \
       deadzone-various-tests

#####################
# subsampling tests #
#####################

$(ENCODED)/yuv410.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 410

$(ENCODED)/yuv420.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 420

$(ENCODED)/yuv422.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 422

$(ENCODED)/yuv444.data: $(ORIGINAL)/claude1.bmp $(JPEG2000)
	$(JPEG2000) encode -i $< -o $@ --subsampling 444

$(DECODED)/yuv410.png: $(ENCODED)/yuv410.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/yuv420.png: $(ENCODED)/yuv420.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/yuv422.png: $(ENCODED)/yuv422.data
	$(JPEG2000) decode -i $< -o $@

$(DECODED)/yuv444.png: $(ENCODED)/yuv444.data
	$(JPEG2000) decode -i $< -o $@

subsampling-tests: $(DECODED)/yuv410.png \
                   $(DECODED)/yuv420.png \
                   $(DECODED)/yuv422.png \
		           $(DECODED)/yuv444.png

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

########################
# deadzone delta tests #
########################

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

##########################
# deadzone various tests #
##########################

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

#################
# other targets #
#################

clean:
	rm -vf $(JPEG2000)
	rm -vf $(ENCODED)/*
	rm -vf $(DECODED)/*
