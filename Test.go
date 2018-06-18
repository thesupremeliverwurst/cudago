package main
/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
typedef void (*cFunc) ();
typedef void (*c2Func) (cFunc);
void begink(c2Func a, cFunc b) {
	(*a)(*b);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
	//"reflect"
)
type tobe func()
func needsGPU(){
	fmt.Println("It worked!")
}

func main() {
	var tr tobe=needsGPU
	f:=&tr
	handle := C.dlopen(C.CString("../cudago/libpar.so"), C.RTLD_LAZY)
	csbf:=C.CString("Start")
	fmt.Println(csbf)
	dbf:=C.dlsym(handle,csbf)
	fmt.Println(dbf)
	csbft:=C.CString("gFunc")
	fmt.Println(csbft)
	dbft:=C.dlsym(handle,csbf)
	fmt.Println(dbft)
	//gFake:=(unsafe.Pointer(&dbft))
	//gFakeType:=*(*func())(unsafe.Pointer(&dbf))
	//fmt.Println(gFakeType)
	fake:=*(*func(*tobe))(unsafe.Pointer(&dbf))//express gFakeType as type
	//fake(*(&f))
	C.begink(unsafe.Pointer(&fake),unsafe.Pointer(&f))
}
