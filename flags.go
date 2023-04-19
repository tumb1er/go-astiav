// Code generated by astiav. DO NOT EDIT.
package astiav

type BuffersinkFlags flags

func NewBuffersinkFlags(fs ...BuffersinkFlag) BuffersinkFlags {
	o := BuffersinkFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs BuffersinkFlags) Add(f BuffersinkFlag) BuffersinkFlags {
	return BuffersinkFlags(flags(fs).add(int(f)))
}

func (fs BuffersinkFlags) Del(f BuffersinkFlag) BuffersinkFlags {
	return BuffersinkFlags(flags(fs).del(int(f)))
}

func (fs BuffersinkFlags) Has(f BuffersinkFlag) bool { return flags(fs).has(int(f)) }

type BuffersrcFlags flags

func NewBuffersrcFlags(fs ...BuffersrcFlag) BuffersrcFlags {
	o := BuffersrcFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs BuffersrcFlags) Add(f BuffersrcFlag) BuffersrcFlags {
	return BuffersrcFlags(flags(fs).add(int(f)))
}

func (fs BuffersrcFlags) Del(f BuffersrcFlag) BuffersrcFlags {
	return BuffersrcFlags(flags(fs).del(int(f)))
}

func (fs BuffersrcFlags) Has(f BuffersrcFlag) bool { return flags(fs).has(int(f)) }

type CodecContextFlags flags

func NewCodecContextFlags(fs ...CodecContextFlag) CodecContextFlags {
	o := CodecContextFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs CodecContextFlags) Add(f CodecContextFlag) CodecContextFlags {
	return CodecContextFlags(flags(fs).add(int(f)))
}

func (fs CodecContextFlags) Del(f CodecContextFlag) CodecContextFlags {
	return CodecContextFlags(flags(fs).del(int(f)))
}

func (fs CodecContextFlags) Has(f CodecContextFlag) bool { return flags(fs).has(int(f)) }

type CodecContextFlags2 flags

func NewCodecContextFlags2(fs ...CodecContextFlag2) CodecContextFlags2 {
	o := CodecContextFlags2(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs CodecContextFlags2) Add(f CodecContextFlag2) CodecContextFlags2 {
	return CodecContextFlags2(flags(fs).add(int(f)))
}

func (fs CodecContextFlags2) Del(f CodecContextFlag2) CodecContextFlags2 {
	return CodecContextFlags2(flags(fs).del(int(f)))
}

func (fs CodecContextFlags2) Has(f CodecContextFlag2) bool { return flags(fs).has(int(f)) }

type DeviceTypeFlags flags

func NewDeviceTypeFlags(fs ...DeviceTypeFlag) DeviceTypeFlags {
	o := DeviceTypeFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs DeviceTypeFlags) Add(f DeviceTypeFlag) DeviceTypeFlags {
	return DeviceTypeFlags(flags(fs).add(int(f)))
}

func (fs DeviceTypeFlags) Del(f DeviceTypeFlag) DeviceTypeFlags {
	return DeviceTypeFlags(flags(fs).del(int(f)))
}

func (fs DeviceTypeFlags) Has(f DeviceTypeFlag) bool { return flags(fs).has(int(f)) }

type DictionaryFlags flags

func NewDictionaryFlags(fs ...DictionaryFlag) DictionaryFlags {
	o := DictionaryFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs DictionaryFlags) Add(f DictionaryFlag) DictionaryFlags {
	return DictionaryFlags(flags(fs).add(int(f)))
}

func (fs DictionaryFlags) Del(f DictionaryFlag) DictionaryFlags {
	return DictionaryFlags(flags(fs).del(int(f)))
}

func (fs DictionaryFlags) Has(f DictionaryFlag) bool { return flags(fs).has(int(f)) }

type FilterCommandFlags flags

func NewFilterCommandFlags(fs ...FilterCommandFlag) FilterCommandFlags {
	o := FilterCommandFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs FilterCommandFlags) Add(f FilterCommandFlag) FilterCommandFlags {
	return FilterCommandFlags(flags(fs).add(int(f)))
}

func (fs FilterCommandFlags) Del(f FilterCommandFlag) FilterCommandFlags {
	return FilterCommandFlags(flags(fs).del(int(f)))
}

func (fs FilterCommandFlags) Has(f FilterCommandFlag) bool { return flags(fs).has(int(f)) }

type FormatContextCtxFlags flags

func NewFormatContextCtxFlags(fs ...FormatContextCtxFlag) FormatContextCtxFlags {
	o := FormatContextCtxFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs FormatContextCtxFlags) Add(f FormatContextCtxFlag) FormatContextCtxFlags {
	return FormatContextCtxFlags(flags(fs).add(int(f)))
}

func (fs FormatContextCtxFlags) Del(f FormatContextCtxFlag) FormatContextCtxFlags {
	return FormatContextCtxFlags(flags(fs).del(int(f)))
}

func (fs FormatContextCtxFlags) Has(f FormatContextCtxFlag) bool { return flags(fs).has(int(f)) }

type FormatContextFlags flags

func NewFormatContextFlags(fs ...FormatContextFlag) FormatContextFlags {
	o := FormatContextFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs FormatContextFlags) Add(f FormatContextFlag) FormatContextFlags {
	return FormatContextFlags(flags(fs).add(int(f)))
}

func (fs FormatContextFlags) Del(f FormatContextFlag) FormatContextFlags {
	return FormatContextFlags(flags(fs).del(int(f)))
}

func (fs FormatContextFlags) Has(f FormatContextFlag) bool { return flags(fs).has(int(f)) }

type FormatEventFlags flags

func NewFormatEventFlags(fs ...FormatEventFlag) FormatEventFlags {
	o := FormatEventFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs FormatEventFlags) Add(f FormatEventFlag) FormatEventFlags {
	return FormatEventFlags(flags(fs).add(int(f)))
}

func (fs FormatEventFlags) Del(f FormatEventFlag) FormatEventFlags {
	return FormatEventFlags(flags(fs).del(int(f)))
}

func (fs FormatEventFlags) Has(f FormatEventFlag) bool { return flags(fs).has(int(f)) }

type IOContextFlags flags

func NewIOContextFlags(fs ...IOContextFlag) IOContextFlags {
	o := IOContextFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs IOContextFlags) Add(f IOContextFlag) IOContextFlags {
	return IOContextFlags(flags(fs).add(int(f)))
}

func (fs IOContextFlags) Del(f IOContextFlag) IOContextFlags {
	return IOContextFlags(flags(fs).del(int(f)))
}

func (fs IOContextFlags) Has(f IOContextFlag) bool { return flags(fs).has(int(f)) }

type IOFormatFlags flags

func NewIOFormatFlags(fs ...IOFormatFlag) IOFormatFlags {
	o := IOFormatFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs IOFormatFlags) Add(f IOFormatFlag) IOFormatFlags {
	return IOFormatFlags(flags(fs).add(int(f)))
}

func (fs IOFormatFlags) Del(f IOFormatFlag) IOFormatFlags {
	return IOFormatFlags(flags(fs).del(int(f)))
}

func (fs IOFormatFlags) Has(f IOFormatFlag) bool { return flags(fs).has(int(f)) }

type PacketFlags flags

func NewPacketFlags(fs ...PacketFlag) PacketFlags {
	o := PacketFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs PacketFlags) Add(f PacketFlag) PacketFlags {
	return PacketFlags(flags(fs).add(int(f)))
}

func (fs PacketFlags) Del(f PacketFlag) PacketFlags {
	return PacketFlags(flags(fs).del(int(f)))
}

func (fs PacketFlags) Has(f PacketFlag) bool { return flags(fs).has(int(f)) }

type SeekFlags flags

func NewSeekFlags(fs ...SeekFlag) SeekFlags {
	o := SeekFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs SeekFlags) Add(f SeekFlag) SeekFlags {
	return SeekFlags(flags(fs).add(int(f)))
}

func (fs SeekFlags) Del(f SeekFlag) SeekFlags {
	return SeekFlags(flags(fs).del(int(f)))
}

func (fs SeekFlags) Has(f SeekFlag) bool { return flags(fs).has(int(f)) }

type StreamEventFlags flags

func NewStreamEventFlags(fs ...StreamEventFlag) StreamEventFlags {
	o := StreamEventFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs StreamEventFlags) Add(f StreamEventFlag) StreamEventFlags {
	return StreamEventFlags(flags(fs).add(int(f)))
}

func (fs StreamEventFlags) Del(f StreamEventFlag) StreamEventFlags {
	return StreamEventFlags(flags(fs).del(int(f)))
}

func (fs StreamEventFlags) Has(f StreamEventFlag) bool { return flags(fs).has(int(f)) }
