package advanced

/*
#cgo CFLAGS: -I${SRCDIR}/lib
#cgo LDFLAGS: -L${SRCDIR}/lib -lsum
#include "sum.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func CGoMain() {
	// win 下需要使用 -shared -o libsum.dll 生成动态链接库
	// gcc -c sum.c -o sum.o
	// gcc -shared -o libsum.dll sum.o
	// Linux 下需要使用 -shared -o libsum.so 生成动态链接库
	// gcc -c sum.c -o sum.o
	// gcc -shared -o libsum.so sum.o

	cs := C.CString("hello CGo")
	defer C.free(unsafe.Pointer(cs))
	C.puts(cs)
	C.fflush(C.stdout) // stdout 缓冲 → puts 的内容延迟出现，造成“先打印 20”的假象
	//
	//retSum := C.getSum(5, 15)
	//fmt.Println(retSum)
}
