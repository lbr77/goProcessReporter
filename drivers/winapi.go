package drivers

import (
	"syscall"
	"unsafe"
)

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	winmm               = syscall.NewLazyDLL("winmm.dll")
	getForegroundWindow = user32.NewProc("GetForegroundWindow")
	getWindowText       = user32.NewProc("GetWindowTextW")
	getProcessId        = user32.NewProc("GetWindowThreadProcessId")
	getModuleFileName   = kernel32.NewProc("K32GetModuleFileNameExW")
	mciSendString       = winmm.NewProc("mciSendStringW")
)

func getWindow2Text(hwnd syscall.Handle) string {
	var buffer [512]uint16
	getWindowText.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buffer)), uintptr(len(buffer)))
	return syscall.UTF16ToString(buffer[:])
}

func getProcessFileName(processId uint32) string {
	var buffer [syscall.MAX_PATH]uint16
	getModuleFileName.Call(uintptr(0), uintptr(processId), uintptr(unsafe.Pointer(&buffer)), uintptr(len(buffer)))
	return syscall.UTF16ToString(buffer[:])
}

func getForegroundWindowHandle() syscall.Handle {
	hwnd, _, _ := getForegroundWindow.Call()
	return syscall.Handle(hwnd)
}

func getWindowProcessId(hwnd syscall.Handle) uint32 {
	var processId uint32
	getProcessId.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&processId)))
	return processId
}

func getPlayingMusicByWindowsAPI() string { //maybe get nothing. || always get noting
	var buffer [512]uint16
	mciCommand := "status MediaFile mode"
	mciSendString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(mciCommand))), uintptr(unsafe.Pointer(&buffer)), uintptr(len(buffer)), 0)
	return syscall.UTF16ToString(buffer[:])
}

func getApplicationForeground() string {
	hwnd := getForegroundWindowHandle()
	if hwnd == 0 {
		return ""
	}
	return getWindow2Text(hwnd)
}
