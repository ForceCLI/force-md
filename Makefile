VERSION=$(shell git describe --abbrev=0 --always)
LDFLAGS = -ldflags "-X github.com/octoberswimmer/force-md/cmd.version=${VERSION}"
EXECUTABLE=force-md
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
OSX=$(EXECUTABLE)_osx_amd64
ALL=$(WINDOWS) $(LINUX) $(OSX)

default:
	go build ${LDFLAGS}

install:
	go install ${LDFLAGS}

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) ${LDFLAGS}

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) ${LDFLAGS}

$(OSX):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(OSX) ${LDFLAGS}

$(basename $(WINDOWS)).zip: $(WINDOWS)
	zip $@ $<
	7za rn $@ $< $(EXECUTABLE)$(suffix $<)

%.zip: %
	zip $@ $<
	7za rn $@ $< $(EXECUTABLE)

dist: $(addsuffix .zip,$(basename $(ALL)))

docs:
	go run docs/mkdocs.go

clean:
	-rm -f $(EXECUTABLE) $(EXECUTABLE)_*

.PHONY: default dist clean docs
