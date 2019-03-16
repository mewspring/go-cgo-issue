//#include <stdio.h>

#include "go-c.h"

char * c1() {
	//printf("c1\n");
	// NOTE: removing this call to FooBar is enough to remove the Go panic.
	//
	//    function symbol table not sorted by program counter: 0xf7eec3f0 _cgo_topofstack > 0xf7dc5f20 runtime.goexit
	//    ...
	//    fatal error: cgo callback before cgo call
	char *s = FooBar();
	//printf("   foobar: %s\n", s);
	return s;
}
