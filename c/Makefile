all: c.so

c.so: c.c
	clang -m32 -shared -o c.so c.c

.PHONY: all clean

clean:
	$(RM) -v c.so
