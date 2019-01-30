all:
	go build -o build/specgen

win:
	GOOS=windows GOARCH=amd64 go build -o build/specgen.exe

install:
	go build -i -o build/specgen && cp build/specgen ${HOME}/go/bin/

test: all
	build/specgen --file=drafts/foo.gspec --lang=c > test/foo.h && gcc test/main.c -o build/stub.out

md_test: all
	build/specgen --file=drafts/foo.gspec --lang=md > test/foo.md
