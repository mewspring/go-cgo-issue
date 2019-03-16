all: b.so

b.so: b.c go-b.a
	clang -shared -o $@ $^ -lpthread

go-b.a: go-b.go
	CGO_ENABLED=1 go build -buildmode=c-archive -o $@ $<

.PHONY: all clean

clean:
	$(RM) -v b.so go-b.a go-b.h
