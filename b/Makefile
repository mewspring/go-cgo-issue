all: b.so

b.so: b.c
	clang -m32 -shared -o b.so b.c

.PHONY: all clean

clean:
	$(RM) -v b.so
