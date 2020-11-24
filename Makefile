TMPLINK = tools/tmplink/tmplink
TMPLINK_SOURCES = $(wildcard $(basename $(TMPLINK))/*.go)

TEMPLATES = $(patsubst %, %.go, $(shell find . -type d -name templates))

MAILARC = cmd/mailarc/mailarc
MAILARC_SOURCES = \
	$(shell find . -name *.go) \
	$(TEMPLATES)


all: $(MAILARC)

install: install-mailarc

clean: clean-tmplink clean-templates clean-mailarc

#
# Template code generator
#

$(TMPLINK): $(TMPLINK_SOURCES)
	go build -o $@ mailarc/$(dir $@)

clean-tmplink:
	rm -f $(TMPLINK)

#
# Generate code for all templates
#


%/templates.go: $(TMPLINK) $(wildcard %/templates/*.gohtml)
	# $*/templates/*.gohtml
	# $(wildcard %/templates/*.gohtml)
	# $^
	$(TMPLINK) -package $(notdir $*) -path $*/templates/ -output $@

mailarc: $(MAILARC_SOURCES)
	go build -o $@ mailarc/cmd/mailarc

clean-templates:
	rm -f $(TEMPLATES)


#
# mailarc
#

$(MAILARC): $(MAILARC_SOURCES)
	go build -o $@ mailarc/$(dir $@)

install-mailarc: $(MAILARC)
	go install mailarc/cmd/mailarc

clean-mailarc:
	rm -f $(MAILARC)


