all: foo

foo: foo.c ../a/a.so ../b/b.so ../c/c.so
	clang -I../a -I../b -I../c -o foo foo.c ../a/a.so ../b/b.so ../c/c.so

../a/a.so:
	make -C ../a

../b/b.so:
	make -C ../b

../c/c.so:
	make -C ../c

.PHONY: all clean

clean:
	$(RM) -v foo
