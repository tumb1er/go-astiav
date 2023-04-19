package astiav

// #cgo pkg-config: libavutil
// #include <libavutil/buffer.h>
// #include <libavutil/hwcontext.h>
import "C"

type DeviceType uint32

const (
	DeviceTypeVAAPI = DeviceType(C.AV_HWDEVICE_TYPE_VAAPI)
	DeviceTypeQSV   = DeviceType(C.AV_HWDEVICE_TYPE_QSV)
	DeviceTypeCUDA  = DeviceType(C.AV_HWDEVICE_TYPE_CUDA)
	DeviceTypeNone  = DeviceType(C.AV_HWDEVICE_TYPE_NONE)
)
