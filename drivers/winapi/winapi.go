package winapi

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	PROCESS_QUERY_INFORMATION = 0x0400
	PROCESS_VM_READ           = 0x0010
	MAX_PATH                  = 260
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
	closeHandle      = kernel32.NewProc("CloseHandle")
	enumProcess      = psapi.NewProc("EnumProcesses")
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
func getProcessIds() ([]uint32, error) {
	var processIds [4096]uint32
	var bytesReturned uint32
	ret, _, _ := enumProcess.Call(
		uintptr(unsafe.Pointer(&processIds[0])),
		uintptr(len(processIds)*4),
		uintptr(unsafe.Pointer(&bytesReturned)),
	)
	if ret == 0 {
		return nil, fmt.Errorf("EnumProcess error")
	}
	return processIds[:bytesReturned/4], nil
}
func isProcessRunning(pid uint32, program string) bool {
	handle, _, _ := openProcess.Call(
		PROCESS_QUERY_INFORMATION|PROCESS_VM_READ,
		0,
		uintptr(pid),
	)
	if handle == 0 {
		return false
	}
	defer closeHandle.Call(handle)
	var buffer [MAX_PATH]uint16
	var bufferSize = uint32(len(buffer))
	ret, _, _ := getModuleBase.Call(
		handle,
		0,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(bufferSize),
	)
	if ret == 0 {
		return false
	}
	processName := syscall.UTF16ToString(buffer[:ret])
	return processName == program
}
func GetRunningPids(program string) []uint32 {
	pids, err := getProcessIds()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var retPids []uint32
	for _, pid := range pids {
		if isProcessRunning(pid, program) {
			retPids = append(retPids, pid)
		}
	}
	return retPids
}

func StopPid(pid uint32) {
	handle, err := syscall.OpenProcess(syscall.PROCESS_TERMINATE, false, pid)
	if err != nil {
		fmt.Println("Failed to open", err)
		return
	}
	defer syscall.CloseHandle(handle)
	err = syscall.TerminateProcess(handle, 0)
	if err != nil {
		fmt.Println("Failed to terminate", err)
		return
	}
	return
}
