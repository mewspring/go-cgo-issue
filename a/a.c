#include <stdio.h>

#include "go-a.h"

int a1() {
	printf("a1\n");
	printf("   argv: %s\n", Args());
	printf("   envp: %s\n", Environ());
	printf("   exe:  %s\n", Executable());
	return 0;
}
