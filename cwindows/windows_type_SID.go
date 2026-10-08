package cwindows

type SID struct {
	Revision            BYTE
	SubAuthorityCount   BYTE
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
	SubAuthority        []DWORD
}
