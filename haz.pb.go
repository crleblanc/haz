// Code generated by protoc-gen-go.
// source: haz.proto
// DO NOT EDIT!

/*
Package haz is a generated protocol buffer package.

It is generated from these files:
	haz.proto

It has these top-level messages:
	Quake
	Timestamp
	Quakes
	Volcano
	VAL
	Volcanoes
	MMI
	Shaking
	Story
	News
	Rate
	QuakeStats
*/
package haz

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Quake is for earthquake information.
type Quake struct {
	// the unique public identifier for this quake.
	PublicID string `protobuf:"bytes,1,opt,name=public_iD,json=publicID" json:"public_iD,omitempty"`
	// the origin time of the quake.
	Time *Timestamp `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	// the modification time of the quake information.
	ModificationTime *Timestamp `protobuf:"bytes,3,opt,name=modification_time,json=modificationTime" json:"modification_time,omitempty"`
	// latitude of the quake location.
	Latitude float64 `protobuf:"fixed64,4,opt,name=latitude" json:"latitude,omitempty"`
	// longitude of the quake location.
	Longitude float64 `protobuf:"fixed64,5,opt,name=longitude" json:"longitude,omitempty"`
	// the depth of the quake in km.
	Depth float64 `protobuf:"fixed64,6,opt,name=depth" json:"depth,omitempty"`
	// magnitude of the quake.
	Magnitude float64 `protobuf:"fixed64,7,opt,name=magnitude" json:"magnitude,omitempty"`
	// distance and direction to the nearest locality.
	Locality string `protobuf:"bytes,8,opt,name=locality" json:"locality,omitempty"`
	// the quality of this information; `best`, `good`, `caution`, `deleted`.
	Quality string `protobuf:"bytes,9,opt,name=quality" json:"quality,omitempty"`
	// the calculated MMI shaking at the closest locality in the New Zealand region.
	Mmi int32 `protobuf:"varint,10,opt,name=mmi" json:"mmi,omitempty"`
}

func (m *Quake) Reset()                    { *m = Quake{} }
func (m *Quake) String() string            { return proto.CompactTextString(m) }
func (*Quake) ProtoMessage()               {}
func (*Quake) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Quake) GetTime() *Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Quake) GetModificationTime() *Timestamp {
	if m != nil {
		return m.ModificationTime
	}
	return nil
}

// Timestamp for encoding time stamps.
type Timestamp struct {
	// Unix time in seconds
	Sec int64 `protobuf:"varint,1,opt,name=sec" json:"sec,omitempty"`
	// Fractional part of time in nanoseconds.
	Nsec int64 `protobuf:"varint,2,opt,name=nsec" json:"nsec,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Quakes struct {
	Quakes []*Quake `protobuf:"bytes,1,rep,name=quakes" json:"quakes,omitempty"`
}

func (m *Quakes) Reset()                    { *m = Quakes{} }
func (m *Quakes) String() string            { return proto.CompactTextString(m) }
func (*Quakes) ProtoMessage()               {}
func (*Quakes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Quakes) GetQuakes() []*Quake {
	if m != nil {
		return m.Quakes
	}
	return nil
}

type Volcano struct {
	// a unique identifier for the volcano.
	VolcanoID string `protobuf:"bytes,1,opt,name=volcano_iD,json=volcanoID" json:"volcano_iD,omitempty"`
	// the volcano title.
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// latitude of the volcano.
	Latitude float64 `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	// longitude of the volcano.
	Longitude float64 `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
	Val       *VAL    `protobuf:"bytes,5,opt,name=val" json:"val,omitempty"`
}

func (m *Volcano) Reset()                    { *m = Volcano{} }
func (m *Volcano) String() string            { return proto.CompactTextString(m) }
func (*Volcano) ProtoMessage()               {}
func (*Volcano) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Volcano) GetVal() *VAL {
	if m != nil {
		return m.Val
	}
	return nil
}

// volcanic alert level.
type VAL struct {
	Level    int32  `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
	Activity string `protobuf:"bytes,2,opt,name=activity" json:"activity,omitempty"`
	Hazards  string `protobuf:"bytes,3,opt,name=hazards" json:"hazards,omitempty"`
}

func (m *VAL) Reset()                    { *m = VAL{} }
func (m *VAL) String() string            { return proto.CompactTextString(m) }
func (*VAL) ProtoMessage()               {}
func (*VAL) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Volcanoes struct {
	Volcanoes []*Volcano `protobuf:"bytes,1,rep,name=volcanoes" json:"volcanoes,omitempty"`
}

func (m *Volcanoes) Reset()                    { *m = Volcanoes{} }
func (m *Volcanoes) String() string            { return proto.CompactTextString(m) }
func (*Volcanoes) ProtoMessage()               {}
func (*Volcanoes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Volcanoes) GetVolcanoes() []*Volcano {
	if m != nil {
		return m.Volcanoes
	}
	return nil
}

type MMI struct {
	// latitude of the mmi.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	// longitude of the mmi.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
	// the mmi at the location.  Currently the max mmi.
	Mmi int32 `protobuf:"varint,3,opt,name=mmi" json:"mmi,omitempty"`
	// count of mmi values at the location.
	Count int32 `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *MMI) Reset()                    { *m = MMI{} }
func (m *MMI) String() string            { return proto.CompactTextString(m) }
func (*MMI) ProtoMessage()               {}
func (*MMI) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Shaking struct {
	Mmi        []*MMI          `protobuf:"bytes,1,rep,name=mmi" json:"mmi,omitempty"`
	MmiSummary map[int32]int32 `protobuf:"bytes,2,rep,name=mmi_summary,json=mmiSummary" json:"mmi_summary,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	MmiTotal   int32           `protobuf:"varint,3,opt,name=mmi_total,json=mmiTotal" json:"mmi_total,omitempty"`
}

func (m *Shaking) Reset()                    { *m = Shaking{} }
func (m *Shaking) String() string            { return proto.CompactTextString(m) }
func (*Shaking) ProtoMessage()               {}
func (*Shaking) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Shaking) GetMmi() []*MMI {
	if m != nil {
		return m.Mmi
	}
	return nil
}

func (m *Shaking) GetMmiSummary() map[int32]int32 {
	if m != nil {
		return m.MmiSummary
	}
	return nil
}

type Story struct {
	Title     string     `protobuf:"bytes,1,opt,name=Title,json=title" json:"Title,omitempty"`
	Link      string     `protobuf:"bytes,2,opt,name=link" json:"link,omitempty"`
	Published *Timestamp `protobuf:"bytes,3,opt,name=published" json:"published,omitempty"`
}

func (m *Story) Reset()                    { *m = Story{} }
func (m *Story) String() string            { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()               {}
func (*Story) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Story) GetPublished() *Timestamp {
	if m != nil {
		return m.Published
	}
	return nil
}

type News struct {
	Stories []*Story `protobuf:"bytes,1,rep,name=stories" json:"stories,omitempty"`
}

func (m *News) Reset()                    { *m = News{} }
func (m *News) String() string            { return proto.CompactTextString(m) }
func (*News) ProtoMessage()               {}
func (*News) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *News) GetStories() []*Story {
	if m != nil {
		return m.Stories
	}
	return nil
}

type Rate struct {
	Time  *Timestamp `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
	Count int32      `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *Rate) Reset()                    { *m = Rate{} }
func (m *Rate) String() string            { return proto.CompactTextString(m) }
func (*Rate) ProtoMessage()               {}
func (*Rate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Rate) GetTime() *Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type QuakeStats struct {
	// quakes per day
	PerDay []*Rate `protobuf:"bytes,1,rep,name=per_day,json=perDay" json:"per_day,omitempty"`
	// magnitude count over the last 7 days
	Week map[int32]int32 `protobuf:"bytes,2,rep,name=week" json:"week,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// magnitude count over the last 28 days
	Month map[int32]int32 `protobuf:"bytes,3,rep,name=month" json:"month,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// magnitude count over the last 365 days
	Year map[int32]int32 `protobuf:"bytes,4,rep,name=year" json:"year,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *QuakeStats) Reset()                    { *m = QuakeStats{} }
func (m *QuakeStats) String() string            { return proto.CompactTextString(m) }
func (*QuakeStats) ProtoMessage()               {}
func (*QuakeStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *QuakeStats) GetPerDay() []*Rate {
	if m != nil {
		return m.PerDay
	}
	return nil
}

func (m *QuakeStats) GetWeek() map[int32]int32 {
	if m != nil {
		return m.Week
	}
	return nil
}

func (m *QuakeStats) GetMonth() map[int32]int32 {
	if m != nil {
		return m.Month
	}
	return nil
}

func (m *QuakeStats) GetYear() map[int32]int32 {
	if m != nil {
		return m.Year
	}
	return nil
}

func init() {
	proto.RegisterType((*Quake)(nil), "haz.Quake")
	proto.RegisterType((*Timestamp)(nil), "haz.Timestamp")
	proto.RegisterType((*Quakes)(nil), "haz.Quakes")
	proto.RegisterType((*Volcano)(nil), "haz.Volcano")
	proto.RegisterType((*VAL)(nil), "haz.VAL")
	proto.RegisterType((*Volcanoes)(nil), "haz.Volcanoes")
	proto.RegisterType((*MMI)(nil), "haz.MMI")
	proto.RegisterType((*Shaking)(nil), "haz.Shaking")
	proto.RegisterType((*Story)(nil), "haz.Story")
	proto.RegisterType((*News)(nil), "haz.News")
	proto.RegisterType((*Rate)(nil), "haz.Rate")
	proto.RegisterType((*QuakeStats)(nil), "haz.QuakeStats")
}

func init() { proto.RegisterFile("haz.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xe3, 0x38, 0x89, 0xa7, 0x08, 0xca, 0x8a, 0x43, 0x08, 0x45, 0x42, 0x16, 0x07, 0x84,
	0x4a, 0x04, 0xe5, 0x40, 0x05, 0xaa, 0x04, 0xa8, 0x1c, 0x2a, 0x11, 0xa4, 0x6e, 0xaa, 0x22, 0xb8,
	0x44, 0x5b, 0x67, 0x49, 0xac, 0xf8, 0x27, 0xb5, 0xd7, 0xa9, 0xc2, 0x73, 0xf0, 0x36, 0x5c, 0x78,
	0x20, 0x1e, 0x82, 0x99, 0x59, 0xc7, 0x69, 0xa3, 0x52, 0xa9, 0xa7, 0xcc, 0xcc, 0x37, 0x9f, 0x67,
	0xe6, 0x9b, 0xdd, 0x0d, 0xf8, 0x53, 0xf5, 0xb3, 0x3f, 0xcf, 0x33, 0x93, 0x09, 0x17, 0xcd, 0xe0,
	0x77, 0x03, 0xbc, 0xe3, 0x52, 0xcd, 0xb4, 0x78, 0x04, 0xfe, 0xbc, 0x3c, 0x8b, 0xa3, 0x70, 0x14,
	0x1d, 0x76, 0x9d, 0x27, 0xce, 0x33, 0x5f, 0x76, 0x6c, 0xe0, 0xe8, 0x50, 0x04, 0xd0, 0x34, 0x51,
	0xa2, 0xbb, 0x0d, 0x8c, 0x6f, 0xed, 0xdd, 0xed, 0xd3, 0x57, 0x4e, 0x30, 0x50, 0x18, 0x95, 0xcc,
	0x25, 0x63, 0xe2, 0x1d, 0xdc, 0x4f, 0xb2, 0x71, 0xf4, 0x23, 0x0a, 0x95, 0x89, 0xb2, 0x74, 0xc4,
	0x04, 0xf7, 0x5a, 0xc2, 0xf6, 0xe5, 0x44, 0x0a, 0x8b, 0x1e, 0x74, 0x62, 0xf4, 0x4c, 0x39, 0xd6,
	0xdd, 0x26, 0x72, 0x1c, 0x59, 0xfb, 0x62, 0x07, 0xfc, 0x38, 0x4b, 0x27, 0x16, 0xf4, 0x18, 0x5c,
	0x07, 0xc4, 0x03, 0xf0, 0xc6, 0x7a, 0x6e, 0xa6, 0xdd, 0x16, 0x23, 0xd6, 0x21, 0x4e, 0xa2, 0x26,
	0xa9, 0xe5, 0xb4, 0x2d, 0xa7, 0x0e, 0x70, 0xb5, 0x2c, 0x54, 0x71, 0x64, 0x96, 0xdd, 0x8e, 0x1d,
	0x75, 0xe5, 0x8b, 0x2e, 0xb4, 0xcf, 0x4b, 0x0b, 0xf9, 0x0c, 0xad, 0x5c, 0xb1, 0x0d, 0x6e, 0x92,
	0x44, 0x5d, 0xc0, 0xa8, 0x27, 0xc9, 0x0c, 0x5e, 0x81, 0x5f, 0x0f, 0x45, 0x70, 0xa1, 0x43, 0x96,
	0xce, 0x95, 0x64, 0x0a, 0x01, 0xcd, 0x94, 0x42, 0x0d, 0x0e, 0xb1, 0x1d, 0xec, 0x42, 0x8b, 0xf5,
	0x2e, 0x50, 0xd3, 0xd6, 0x39, 0x5b, 0x48, 0x71, 0x51, 0x24, 0x60, 0x91, 0x18, 0x94, 0x15, 0x12,
	0xfc, 0x72, 0xa0, 0x7d, 0x9a, 0xc5, 0xa1, 0x4a, 0x33, 0xf1, 0x18, 0x60, 0x61, 0xcd, 0xf5, 0x86,
	0xfc, 0x2a, 0x82, 0x2b, 0x42, 0x1d, 0x50, 0xaf, 0xd8, 0xee, 0xc8, 0x97, 0xd6, 0xb9, 0xa2, 0xab,
	0x7b, 0x93, 0xae, 0xcd, 0x4d, 0x5d, 0x7b, 0xe0, 0x2e, 0x54, 0xcc, 0x7a, 0x6f, 0xed, 0x75, 0xb8,
	0xb7, 0xd3, 0x0f, 0x9f, 0x25, 0x05, 0x83, 0x63, 0x70, 0xd1, 0xa6, 0x92, 0xb1, 0x5e, 0xe8, 0x98,
	0x9b, 0xf1, 0xa4, 0x75, 0xa8, 0xa4, 0x0a, 0x4d, 0xb4, 0x20, 0x05, 0x6d, 0x2f, 0xb5, 0x4f, 0xe2,
	0xe2, 0x87, 0x54, 0x3e, 0x2e, 0xb8, 0x1b, 0x14, 0xb7, 0x72, 0x83, 0x37, 0xe0, 0x57, 0x83, 0xa2,
	0x34, 0xcf, 0x61, 0x35, 0x58, 0xad, 0xce, 0x1d, 0xdb, 0x81, 0x8d, 0xca, 0x35, 0x1c, 0x4c, 0xc0,
	0x1d, 0x0c, 0x8e, 0xae, 0x0c, 0xea, 0xdc, 0x34, 0x68, 0x63, 0x73, 0xd0, 0x6a, 0xad, 0x6e, 0xbd,
	0x56, 0x9a, 0x2b, 0xcc, 0xca, 0xd4, 0xb0, 0x28, 0x38, 0x17, 0x3b, 0xc1, 0x1f, 0xdc, 0xc5, 0x70,
	0xaa, 0x66, 0x51, 0x3a, 0x21, 0x71, 0x88, 0x63, 0x5b, 0xb3, 0xe2, 0x60, 0x13, 0x96, 0x7d, 0x00,
	0x5b, 0xf8, 0x33, 0x2a, 0xca, 0x24, 0x51, 0x39, 0x49, 0x40, 0x39, 0x3b, 0x9c, 0x53, 0xd1, 0xfb,
	0x83, 0x24, 0x1a, 0x5a, 0xf8, 0x53, 0x6a, 0xf2, 0xa5, 0x84, 0xa4, 0x0e, 0xd0, 0x3d, 0x24, 0xba,
	0xc9, 0x0c, 0xaa, 0x6f, 0x9b, 0xea, 0x60, 0xe0, 0x84, 0xfc, 0xde, 0x01, 0xdc, 0xdb, 0xe0, 0x52,
	0xfb, 0x33, 0xbd, 0xac, 0x56, 0x40, 0x26, 0xb5, 0x8f, 0x4b, 0x2a, 0xed, 0xa8, 0xd8, 0x3e, 0x3b,
	0x6f, 0x1b, 0xfb, 0x4e, 0x30, 0x02, 0x6f, 0x68, 0xb2, 0x9c, 0x53, 0x4e, 0xf8, 0xb0, 0x38, 0x97,
	0x0f, 0x0b, 0x9e, 0xd7, 0x38, 0x4a, 0x67, 0xd5, 0xd6, 0xd8, 0x16, 0xbb, 0xd5, 0xb3, 0x50, 0x4c,
	0xf5, 0xf8, 0x3f, 0xb7, 0x79, 0x9d, 0x80, 0xa7, 0xbb, 0xf9, 0x45, 0x5f, 0x14, 0xe2, 0x29, 0xb4,
	0x0b, 0x2c, 0x14, 0x6d, 0x1c, 0x6e, 0x2e, 0x2e, 0x57, 0x50, 0xf0, 0x1e, 0x9a, 0x52, 0x19, 0x5d,
	0xbf, 0x2e, 0xce, 0x0d, 0xaf, 0x4b, 0xbd, 0x93, 0xc6, 0xe5, 0x9d, 0xfc, 0x6d, 0x00, 0xf0, 0x8d,
	0x19, 0x1a, 0x65, 0xe8, 0x4a, 0xb5, 0xe7, 0x3a, 0x1f, 0x8d, 0xd5, 0xb2, 0x2a, 0xeb, 0xf3, 0xb7,
	0xa8, 0x88, 0x6c, 0x21, 0x72, 0xa8, 0x96, 0xe2, 0x05, 0x34, 0x2f, 0xb4, 0x9e, 0x55, 0x7b, 0x79,
	0xb8, 0xbe, 0x74, 0xfc, 0x89, 0xfe, 0x57, 0xc4, 0xec, 0x52, 0x38, 0x4d, 0xbc, 0x04, 0x2f, 0xc9,
	0x52, 0x7c, 0x5e, 0x5c, 0xce, 0xef, 0x6d, 0xe6, 0x0f, 0x08, 0xb4, 0x04, 0x9b, 0x48, 0x05, 0x96,
	0x5a, 0xe5, 0x78, 0x78, 0xae, 0x2d, 0xf0, 0x0d, 0xb1, 0xaa, 0x00, 0xa5, 0xf5, 0xf0, 0xe0, 0xd7,
	0x35, 0x6f, 0xb3, 0xcc, 0xde, 0x3e, 0xc0, 0xba, 0xf8, 0xad, 0x98, 0x58, 0xb2, 0xee, 0xe2, 0x36,
	0xc4, 0x8f, 0xde, 0x77, 0xfa, 0xd3, 0x38, 0x6b, 0xf1, 0x1f, 0xc8, 0xeb, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0x9a, 0xf3, 0x01, 0x4d, 0x06, 0x00, 0x00,
}