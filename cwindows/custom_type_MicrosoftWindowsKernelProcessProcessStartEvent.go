// Code generated DO NOT EDIT

package cwindows

import (
	"encoding/binary"
	"unsafe"
)

type MicrosoftWindowsKernelProcessProcessStartEvent struct {
	ProcessID                   ULONG
	ProcessSequenceNumber       ULONGLONG
	CreateTime                  FILETIME
	ParentProcessID             ULONG
	ParentProcessSequenceNumber ULONGLONG
	SessionID                   ULONG
	Flags                       ULONG
	ProcessTokenElevationType   ULONG
	ProcessTokenIsElevated      ULONG
	MandatoryLabel              SID
	ImageName                   []WCHAR
	ImageChecksum               ULONG
	TimeDateStamp               ULONG
	PackageFullName             []WCHAR
	PackageRelativeAppId        []WCHAR
}

func NewMicrosoftWindowsKernelProcessProcessStartEvent(UserData PVOID, UserDataLength USHORT) MicrosoftWindowsKernelProcessProcessStartEvent {
	parseULONG := func(UserData []byte, offset *int) ULONG {
		result := ULONG(binary.LittleEndian.Uint32(UserData[*offset : *offset+4]))
		*offset += 4
		return result
	}
	parseULONGLONG := func(UserData []byte, offset *int) ULONGLONG {
		result := ULONGLONG(binary.LittleEndian.Uint64(UserData[*offset : *offset+8]))
		*offset += 8
		return result
	}
	parseDWORD := func(UserData []byte, offset *int) DWORD {
		result := DWORD(binary.LittleEndian.Uint32(UserData[*offset : *offset+4]))
		*offset += 4
		return result
	}
	parseFILETIME := func(UserData []byte, offset *int) FILETIME {
		lowDateTime := parseDWORD(UserData, offset)
		highDateTime := parseDWORD(UserData, offset)
		return FILETIME{
			LowDateTime:  lowDateTime,
			HighDateTime: highDateTime,
		}
	}
	parseSID := func(UserData []byte, offset *int) SID {
		revision := UserData[*offset]
		subAuthorityCount := UserData[*offset+1]
		identifierAuthority := *((*SIDIdentifierAuthority)(unsafe.Pointer(&UserData[*offset+2])))
		*offset += 8
		subAuthorities := make([]DWORD, subAuthorityCount)
		for i := 0; i < int(subAuthorityCount); i++ {
			subAuthorities[i] = DWORD(binary.LittleEndian.Uint32(UserData[*offset : *offset+4]))
			*offset += 4
		}
		return SID{
			Revision:          BYTE(revision),
			SubAuthorityCount: BYTE(subAuthorityCount),
			IdentifierAuthority: SID_IDENTIFIER_AUTHORITY{
				Value: identifierAuthority,
			},
			SubAuthority: subAuthorities,
		}
	}
	parseWCHARSlice := func(UserData []byte, offset *int) []WCHAR {
		start := *offset
		for i := start; i < len(UserData)-1; i += 2 {
			char := binary.LittleEndian.Uint16(UserData[i : i+2])
			if char == 0 {
				chars := unsafe.Slice(
					(*WCHAR)(unsafe.Pointer(&UserData[start])),
					(i-start)/2,
				)
				*offset = i + 2
				return chars
			}
		}
		*offset = len(UserData)
		return nil
	}

	data := unsafe.Slice((*byte)(UserData), UserDataLength)
	offset := 0
	processID := parseULONG(data, &offset)
	processSequenceNumber := parseULONGLONG(data, &offset)
	createTime := parseFILETIME(data, &offset)
	parentProcessID := parseULONG(data, &offset)
	parentProcessSequenceNumber := parseULONGLONG(data, &offset)
	sessionID := parseULONG(data, &offset)
	flags := parseULONG(data, &offset)
	processTokenElevationType := parseULONG(data, &offset)
	processTokenIsElevated := parseULONG(data, &offset)
	mandatoryLabel := parseSID(data, &offset)
	imageName := parseWCHARSlice(data, &offset)
	imageChecksum := parseULONG(data, &offset)
	timeDateStamp := parseULONG(data, &offset)
	packageFullName := parseWCHARSlice(data, &offset)
	packageRelativeAppId := parseWCHARSlice(data, &offset)
	return MicrosoftWindowsKernelProcessProcessStartEvent{
		ProcessID:                   processID,
		ProcessSequenceNumber:       processSequenceNumber,
		CreateTime:                  createTime,
		ParentProcessID:             parentProcessID,
		ParentProcessSequenceNumber: parentProcessSequenceNumber,
		SessionID:                   sessionID,
		Flags:                       flags,
		ProcessTokenElevationType:   processTokenElevationType,
		ProcessTokenIsElevated:      processTokenIsElevated,
		MandatoryLabel:              mandatoryLabel,
		ImageName:                   imageName,
		ImageChecksum:               imageChecksum,
		TimeDateStamp:               timeDateStamp,
		PackageFullName:             packageFullName,
		PackageRelativeAppId:        packageRelativeAppId,
	}
}
