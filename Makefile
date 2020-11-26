TMPLINK = tools/tmplink/tmplink
TMPLINK_SOURCES = $(wildcard $(basename $(TMPLINK))/*.go)

VIEWS = build/gen/views/views.go
VIEW_SOURCES = $(wildcard internal/server/views/*.gohtml)

MAILARC = cmd/mailarc/mailarc
MAILARC_SOURCES = \
	$(shell find . -name *.go) \
	$(VIEWS)


all: $(MAILARC)

install: install-mailarc

clean: clean-tmplink clean-views clean-mailarc

#
# Template code generator
#

$(TMPLINK): $(TMPLINK_SOURCES)
	go build -o $@ mailarc/$(dir $@)

clean-tmplink:
	rm -f $(TMPLINK)

#
# Generate code for all views
#

$(VIEWS): $(TMPLINK) $(VIEW_SOURCES)
	mkdir -p build/gen/views || :
	$(TMPLINK) -package views -path internal/server/views/ -output $@

clean-views:
	rm -f $(VIEWS)


#
# mailarc
#

$(MAILARC): $(MAILARC_SOURCES)
	go build -o $@ mailarc/$(dir $@)

install-mailarc: $(MAILARC)
	go install mailarc/cmd/mailarc

clean-mailarc:
	rm -f $(MAILARC)


