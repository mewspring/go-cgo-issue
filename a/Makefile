all: a.so

a.so: a.c go-a.a
	clang -m32 -shared -o $@ $^ -lpthread

go-a.a: go-a.go
	CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-archive -o $@ $<

.PHONY: all clean

clean:
	$(RM) -v a.so go-a.a go-a.h
