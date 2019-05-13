// Code generated by 'go generate'; DO NOT EDIT.

package service

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modwtsapi32 = windows.NewLazySystemDLL("wtsapi32.dll")
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procWTSQueryUserToken          = modwtsapi32.NewProc("WTSQueryUserToken")
	procWTSEnumerateSessionsW      = modwtsapi32.NewProc("WTSEnumerateSessionsW")
	procWTSFreeMemory              = modwtsapi32.NewProc("WTSFreeMemory")
	procNotifyServiceStatusChangeW = modadvapi32.NewProc("NotifyServiceStatusChangeW")
	procSleepEx                    = modkernel32.NewProc("SleepEx")
)

func wtsQueryUserToken(session uint32, token *windows.Token) (err error) {
	r1, _, e1 := syscall.Syscall(procWTSQueryUserToken.Addr(), 2, uintptr(session), uintptr(unsafe.Pointer(token)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func wtsEnumerateSessions(handle windows.Handle, reserved uint32, version uint32, sessions **WTS_SESSION_INFO, count *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procWTSEnumerateSessionsW.Addr(), 5, uintptr(handle), uintptr(reserved), uintptr(version), uintptr(unsafe.Pointer(sessions)), uintptr(unsafe.Pointer(count)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func wtsFreeMemory(ptr uintptr) {
	syscall.Syscall(procWTSFreeMemory.Addr(), 1, uintptr(ptr), 0, 0)
	return
}

func notifyServiceStatusChange(service windows.Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) {
	r0, _, _ := syscall.Syscall(procNotifyServiceStatusChangeW.Addr(), 3, uintptr(service), uintptr(notifyMask), uintptr(unsafe.Pointer(notifier)))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func sleepEx(milliseconds uint32, alertable bool) (ret uint32) {
	var _p0 uint32
	if alertable {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, _ := syscall.Syscall(procSleepEx.Addr(), 2, uintptr(milliseconds), uintptr(_p0), 0)
	ret = uint32(r0)
	return
}
