package drivers

import (
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	psapi            = syscall.NewLazyDLL("psapi.dll")
	getForegroundWin = user32.NewProc("GetForegroundWindow")
	getWindowText    = user32.NewProc("GetWindowTextW")
	getWindowThread  = user32.NewProc("GetWindowThreadProcessId")
	openProcess      = kernel32.NewProc("OpenProcess")
	getModuleBase    = psapi.NewProc("GetModuleBaseNameW")
)

func GetActiveWindowProcessAndTitle() (string, string) {
	hwnd, _, _ := getForegroundWin.Call()
	windowTitle := make([]uint16, 255)
	getWindowText.Call(hwnd, uintptr(unsafe.Pointer(&windowTitle[0])), 255)
	var processID uint32
	getWindowThread.Call(hwnd, uintptr(unsafe.Pointer(&processID)), 0)
	processHandle, _, _ := openProcess.Call(0x0400|0x0010, 0, uintptr(processID))
	processName := make([]uint16, 255)
	getModuleBase.Call(processHandle, 0, uintptr(unsafe.Pointer(&processName[0])), 255)
	return syscall.UTF16ToString(processName), syscall.UTF16ToString(windowTitle)
}
