// Code generated DO NOT EDIT

package cwindows

import "unsafe"

type PETW_BUFFER_CALLBACK func(Buffer *ETW_BUFFER_HEADER, BufferSize ULONG, ConsumerInfo *ETW_BUFFER_CALLBACK_INFORMATION, CallbackContext unsafe.Pointer) BOOL
