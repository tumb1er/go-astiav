package astiav

//#cgo pkg-config: libavutil
//#include <libavutil/channel_layout.h>
//#include <libavutil/frame.h>
//#include <libavutil/imgutils.h>
//#include <libavutil/samplefmt.h>
import "C"

const NumDataPointers = uint(C.AV_NUM_DATA_POINTERS)

// https://github.com/FFmpeg/FFmpeg/blob/n5.0/libavutil/frame.h#L317
type Frame struct {
	c *C.struct_AVFrame
}

func newFrameFromC(c *C.struct_AVFrame) *Frame {
	if c == nil {
		return nil
	}
	return &Frame{c: c}
}

func AllocFrame() *Frame {
	return newFrameFromC(C.av_frame_alloc())
}

func (f *Frame) AllocBuffer(align int) error {
	return newError(C.av_frame_get_buffer(f.c, C.int(align)))
}

func (f *Frame) AllocImage(align int) error {
	return newError(C.av_image_alloc(&f.c.data[0], &f.c.linesize[0], f.c.width, f.c.height, (C.enum_AVPixelFormat)(f.c.format), C.int(align)))
}

func (f *Frame) AllocSamples(align int) error {
	return newError(C.av_samples_alloc(&f.c.data[0], &f.c.linesize[0], f.c.ch_layout.nb_channels, f.c.nb_samples, (C.enum_AVSampleFormat)(f.c.format), C.int(align)))
}

func (f *Frame) ChannelLayout() *ChannelLayout {
	return newChannelLayoutFromC(&f.c.ch_layout)
}

func (f *Frame) SetChannelLayout(l *ChannelLayout) {
	l.copy(&f.c.ch_layout) //nolint: errcheck
}

func (f *Frame) ColorRange() ColorRange {
	return ColorRange(f.c.color_range)
}

func (f *Frame) SetColorRange(r ColorRange) {
	f.c.color_range = C.enum_AVColorRange(r)
}

func (f *Frame) Data() [NumDataPointers][]byte {
	b := [NumDataPointers][]byte{}
	for i := 0; i < int(NumDataPointers); i++ {
		b[i] = bytesFromC(func(size *C.ulong) *C.uint8_t {
			*size = C.ulong(f.c.linesize[i])
			if f.c.height > 0 {
				*size = *size * C.ulong(f.c.height)
			} else if f.c.channels > 0 {
				*size = *size * C.ulong(f.c.channels)
			}
			return f.c.data[i]
		})
	}
	return b
}

func (f *Frame) Height() int {
	return int(f.c.height)
}

func (f *Frame) SetHeight(h int) {
	f.c.height = C.int(h)
}

func (f *Frame) KeyFrame() bool {
	return int(f.c.key_frame) > 0
}

func (f *Frame) SetKeyFrame(k bool) {
	i := 0
	if k {
		i = 1
	}
	f.c.key_frame = C.int(i)
}

func (f *Frame) ImageFillBlack() error {
	linesize := [NumDataPointers]C.long{}
	for i := 0; i < int(NumDataPointers); i++ {
		linesize[i] = C.long(f.c.linesize[i])
	}
	return newError(C.av_image_fill_black(&f.c.data[0], &linesize[0], (C.enum_AVPixelFormat)(f.c.format), (C.enum_AVColorRange)(f.c.color_range), f.c.width, f.c.height))
}

func (f *Frame) Linesize() [NumDataPointers]int {
	o := [NumDataPointers]int{}
	for i := 0; i < int(NumDataPointers); i++ {
		o[i] = int(f.c.linesize[i])
	}
	return o
}

func (f *Frame) NbSamples() int {
	return int(f.c.nb_samples)
}

func (f *Frame) SetNbSamples(n int) {
	f.c.nb_samples = C.int(n)
}

func (f *Frame) PictureType() PictureType {
	return PictureType(f.c.pict_type)
}

func (f *Frame) SetPictureType(t PictureType) {
	f.c.pict_type = C.enum_AVPictureType(t)
}

func (f *Frame) PixelFormat() PixelFormat {
	return PixelFormat(f.c.format)
}

func (f *Frame) SetPixelFormat(pf PixelFormat) {
	f.c.format = C.int(pf)
}

func (f *Frame) PktDts() int64 {
	return int64(f.c.pkt_dts)
}

func (f *Frame) Pts() int64 {
	return int64(f.c.pts)
}

func (f *Frame) SetPts(i int64) {
	f.c.pts = C.int64_t(i)
}

func (f *Frame) SampleFormat() SampleFormat {
	return SampleFormat(f.c.format)
}

func (f *Frame) SetSampleFormat(sf SampleFormat) {
	f.c.format = C.int(sf)
}

func (f *Frame) SampleRate() int {
	return int(f.c.sample_rate)
}

func (f *Frame) SetSampleRate(r int) {
	f.c.sample_rate = C.int(r)
}

func (f *Frame) NewSideData(t FrameSideDataType, size uint64) *FrameSideData {
	return newFrameSideDataFromC(C.av_frame_new_side_data(f.c, (C.enum_AVFrameSideDataType)(t), C.ulong(size)))
}

func (f *Frame) SideData(t FrameSideDataType) *FrameSideData {
	return newFrameSideDataFromC(C.av_frame_get_side_data(f.c, (C.enum_AVFrameSideDataType)(t)))
}

func (f *Frame) Width() int {
	return int(f.c.width)
}

func (f *Frame) SetWidth(w int) {
	f.c.width = C.int(w)
}

func (f *Frame) Free() {
	C.av_frame_free(&f.c)
}

func (f *Frame) Ref(src *Frame) error {
	return newError(C.av_frame_ref(f.c, src.c))
}

func (f *Frame) Clone() *Frame {
	return newFrameFromC(C.av_frame_clone(f.c))
}

func (f *Frame) Unref() {
	C.av_frame_unref(f.c)
}

func (f *Frame) MoveRef(src *Frame) {
	C.av_frame_move_ref(f.c, src.c)
}
