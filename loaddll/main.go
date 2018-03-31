package main

import "C"
import (
	"log"
	"syscall"
	"unsafe"

	"github.com/vitaminwater/cgo.wchar"
)

func main() {
	h, err := syscall.LoadDLL("speechSynthesisWin.dll")
	if err != nil {
		log.Println("Can not found the tts dll file:", err)
		return
	}

	tts, err := h.FindProc("TTS")
	if err != nil {
		log.Println("Can not found the function:", err)
		return
	}

	k1, k2, err := tts.Call(uintptr(unsafe.Pointer((*uint16)("123"))), uintptr((*uint16)("D:/12.wav")))
	if err != nil {
		log.Panicln("Can not run the function:", err)
		return
	}

	log.Println(k1, "   ", k2)
}

func CwcharToString(p uintptr, maxchars int) string {

	if p == 0 {
		return ""
	}
	uints := make([]uint16, 0, maxchars)
	for i, p := 0, uintptr(unsafe.Pointer(p)); i < maxchars; p += 2 {

		u := *(*uint16)(unsafe.Pointer(p))
		if u == 0 {
			break
		}
		uints = append(uints, u)
		i++
	}
	return string(utf16.Decode(uints))
}

func StringToC(p uintptr, maxchars int) string {

	if p == 0 {
		return ""
	}
	uints := make([]uint16, 0, maxchars)
	for i, p := 0, uintptr(unsafe.Pointer(p)); i < maxchars; p += 2 {

		u := *(*uint16)(unsafe.Pointer(p))
		if u == 0 {
			break
		}
		uints = append(uints, u)
		i++
	}
	return string(utf16.Decode(uints))
}
