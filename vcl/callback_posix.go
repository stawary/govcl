// +build !windows
// +build cgo

//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

// extern void* doEventCallbackProc(void* f, void* args, long argcount);
// static void* doGetEventCallbackAddr() {
//    return &doEventCallbackProc;
// }
//
// extern void* doMessageCallbackProc(void* f, void* msg);
// static void* doGetMessageCallbackAddr() {
//    return &doMessageCallbackProc;
// }
//
// extern void* doThreadSyncCallbackProc();
// static void* doGetThreadSyncCallbackAddr() {
//    return &doThreadSyncCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doEventCallbackProc
func doEventCallbackProc(f unsafe.Pointer, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	eventCallbackProc(uintptr(f), uintptr(args), int(argcount))
	return unsafe.Pointer(uintptr(0))
}

//export doMessageCallbackProc
func doMessageCallbackProc(f unsafe.Pointer, msg unsafe.Pointer) unsafe.Pointer {
	messageCallbackProc(uintptr(f), uintptr(msg))
	return unsafe.Pointer(uintptr(0))
}

//export doThreadSyncCallbackProc
func doThreadSyncCallbackProc() unsafe.Pointer {
	threadSyncCallbackProc()
	return unsafe.Pointer(uintptr(0))
}

var (
	eventCallback      = uintptr(C.doGetEventCallbackAddr())
	messageCallback    = uintptr(C.doGetMessageCallbackAddr())
	threadSyncCallback = uintptr(C.doGetThreadSyncCallbackAddr())
)
