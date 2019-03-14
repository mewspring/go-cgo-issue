# PoC for Cgo issue

To reproduce, run

```bash
cd foo
make -B
./foo
```

This should produce the follow crash.

```
u@x1 ~/D/f/foo> ./foo
function symbol table not sorted by program counter: 0xf7ef1f00 _cgo_topofstack > 0xf7dcbf20 runtime.goexit
    0xf7d80780 __x86.get_pc_thunk.ax
    0xf7d80790 __x86.get_pc_thunk.cx
..
    0xf7d8c840 runtime.evacuate_fast32
    0xf7d8cc30 runtime.mapassign_faststr
a1
       args: ./foo
b1
c1
fatal error: cgo callback before cgo call
```

The crash is present on both Go 1.12, and master.

```
$ go version
go version devel +04f1b65cc6 Wed Mar 13 15:11:37 2019 +0000 linux/amd64
```
