package astiav

//#cgo pkg-config: libavcodec libavutil
//#include <libavcodec/avcodec.h>
//#include <libavutil/dict.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

// https://github.com/FFmpeg/FFmpeg/blob/n5.0/libavutil/dict.h#L84
type Dictionary struct {
	c *C.struct_AVDictionary
}

func NewDictionary() *Dictionary {
	return &Dictionary{}
}

func newDictionaryFromC(c *C.struct_AVDictionary) *Dictionary {
	if c == nil {
		return nil
	}
	return &Dictionary{c: c}
}

func (d *Dictionary) Set(key, value string, flags DictionaryFlags) error {
	ck := C.CString(key)
	defer C.free(unsafe.Pointer(ck))
	cv := C.CString(value)
	defer C.free(unsafe.Pointer(cv))
	return newError(C.av_dict_set(&d.c, ck, cv, C.int(flags)))
}

func (d *Dictionary) ParseString(i, keyValSep, pairsSep string, flags DictionaryFlags) error {
	ci := C.CString(i)
	defer C.free(unsafe.Pointer(ci))
	ck := C.CString(keyValSep)
	defer C.free(unsafe.Pointer(ck))
	cp := C.CString(pairsSep)
	defer C.free(unsafe.Pointer(cp))
	return newError(C.av_dict_parse_string(&d.c, ci, ck, cp, C.int(flags)))
}

func (d *Dictionary) Get(key string, prev *DictionaryEntry, flags DictionaryFlags) *DictionaryEntry {
	ck := C.CString(key)
	defer C.free(unsafe.Pointer(ck))
	var cp *C.struct_AVDictionaryEntry
	if prev != nil {
		cp = prev.c
	}
	if e := C.av_dict_get(d.c, ck, cp, C.int(flags)); e != nil {
		return newDictionaryEntryFromC(e)
	}
	return nil
}

func (d *Dictionary) Free() {
	C.av_dict_free(&d.c)
}

func (d *Dictionary) Pack() []byte {
	return bytesFromC(func(size *C.ulong) *C.uint8_t {
		return C.av_packet_pack_dictionary(d.c, size)
	})
}

func (d *Dictionary) Unpack(b []byte) error {
	return bytesToC(b, func(b *C.uint8_t, size C.ulong) error {
		return newError(C.av_packet_unpack_dictionary(b, size, &d.c))
	})
}
