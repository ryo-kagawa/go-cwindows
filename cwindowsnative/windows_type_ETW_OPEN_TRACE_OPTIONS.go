// Code generated DO NOT EDIT

package cwindowsnative

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type ETW_OPEN_TRACE_OPTIONS struct {
	ProcessTraceModes     ETW_PROCESS_TRACE_MODES
	EventCallback         uintptr
	EventCallbackContext  unsafe.Pointer
	BufferCallback        uintptr
	BufferCallbackContext unsafe.Pointer
}

func (e *ETW_OPEN_TRACE_OPTIONS) SetEventCallback(value PEVENT_RECORD_CALLBACK) {
	e.EventCallback = windows.NewCallback(func(EventRecord PEVENT_RECORD) uintptr {
		value(EventRecord)
		return 0
	})
}

func (e *ETW_OPEN_TRACE_OPTIONS) SetBufferCallback(value PETW_BUFFER_CALLBACK) {
	e.BufferCallback = windows.NewCallback(func(Buffer *ETW_BUFFER_HEADER, BufferSize ULONG, ConsumerInfo *ETW_BUFFER_CALLBACK_INFORMATION, CallbackContext unsafe.Pointer) BOOL {
		return value(Buffer, BufferSize, ConsumerInfo, CallbackContext)
	})
}
