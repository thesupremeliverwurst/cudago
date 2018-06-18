extern "C"{
	#include "Kernalize.h"
}
// typedef void (*gFunc) ();
// __global__
// void Kernalize::kernalize(gFunc *ptr) {
//     (*ptr)();
// }
// void Kernalize::start(gFunc *ptr) {
// 	kernalize<<<1,1>>>(ptr);
// }
extern "C" {
	typedef void (*gFunc) () __asm__ ("cudago.tobe");
	__global__ void kernalize(gFunc* ptr) {
	    (*ptr)();
	}
	void Start(gFunc* ptr) {
		kernalize<<<1,1>>>(ptr);
	}
}