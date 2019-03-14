all: c.so

c.so: c.c go-c.a
	clang -m32 -shared -o $@ $^ -lpthread

go-c.a: go-c.go
	CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-archive -o $@ $<

.PHONY: all clean

clean:
	$(RM) -v c.so go-c.a go-c.h
