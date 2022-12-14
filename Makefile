VERSION=$(shell git describe --abbrev=0 --always)
LDFLAGS = -ldflags "-X github.com/octoberswimmer/force-md/cmd.version=${VERSION}"
EXECUTABLE=force-md
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
OSX_AMD64=$(EXECUTABLE)_osx_amd64
OSX_ARM64=$(EXECUTABLE)_osx_arm64
ALL=$(WINDOWS) $(LINUX) $(OSX_ARM64) $(OSX_AMD64)

default:
	go build ${LDFLAGS}

install:
	go install ${LDFLAGS}

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) ${LDFLAGS}

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) ${LDFLAGS}

$(OSX_ARM64):
	env GOOS=darwin GOARCH=arm64 go build -v -o $(OSX_ARM64) ${LDFLAGS}

$(OSX_AMD64):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(OSX_AMD64) ${LDFLAGS}

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
