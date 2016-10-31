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
	QuakeTechnical
	RealQuantity
	StationMagnitude
	Magnitude
	Pick
	Waveform
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

// QuakeTechnical for encoding technical information about a quake.
// More information than haz.Quake.
type QuakeTechnical struct {
	PublicID         string        `protobuf:"bytes,1,opt,name=public_iD,json=publicID" json:"public_iD,omitempty"`
	Type             string        `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Agency           string        `protobuf:"bytes,3,opt,name=agency" json:"agency,omitempty"`
	Time             *Timestamp    `protobuf:"bytes,4,opt,name=time" json:"time,omitempty"`
	ModificationTime *Timestamp    `protobuf:"bytes,5,opt,name=modification_time,json=modificationTime" json:"modification_time,omitempty"`
	Latitude         *RealQuantity `protobuf:"bytes,6,opt,name=latitude" json:"latitude,omitempty"`
	Longitude        *RealQuantity `protobuf:"bytes,7,opt,name=longitude" json:"longitude,omitempty"`
	Depth            *RealQuantity `protobuf:"bytes,8,opt,name=depth" json:"depth,omitempty"`
	DepthType        string        `protobuf:"bytes,9,opt,name=depth_type,json=depthType" json:"depth_type,omitempty"`
	Method           string        `protobuf:"bytes,10,opt,name=method" json:"method,omitempty"`
	EarthModel       string        `protobuf:"bytes,11,opt,name=earth_model,json=earthModel" json:"earth_model,omitempty"`
	EvaluationMode   string        `protobuf:"bytes,12,opt,name=evaluation_mode,json=evaluationMode" json:"evaluation_mode,omitempty"`
	EvaluationStatus string        `protobuf:"bytes,13,opt,name=evaluation_status,json=evaluationStatus" json:"evaluation_status,omitempty"`
	UsedPhaseCount   int64         `protobuf:"varint,14,opt,name=used_phase_count,json=usedPhaseCount" json:"used_phase_count,omitempty"`
	UsedStationCount int64         `protobuf:"varint,15,opt,name=used_station_count,json=usedStationCount" json:"used_station_count,omitempty"`
	StandardError    float64       `protobuf:"fixed64,16,opt,name=standard_error,json=standardError" json:"standard_error,omitempty"`
	AzimuthalGap     float64       `protobuf:"fixed64,17,opt,name=azimuthal_gap,json=azimuthalGap" json:"azimuthal_gap,omitempty"`
	MinimumDistance  float64       `protobuf:"fixed64,18,opt,name=minimum_distance,json=minimumDistance" json:"minimum_distance,omitempty"`
	MaximumDistance  float64       `protobuf:"fixed64,19,opt,name=maximum_distance,json=maximumDistance" json:"maximum_distance,omitempty"`
	MedianDistance   float64       `protobuf:"fixed64,20,opt,name=median_distance,json=medianDistance" json:"median_distance,omitempty"`
	Pick             []*Pick       `protobuf:"bytes,21,rep,name=pick" json:"pick,omitempty"`
	Magnitude        *RealQuantity `protobuf:"bytes,22,opt,name=magnitude" json:"magnitude,omitempty"`
	MagnitudeType    string        `protobuf:"bytes,23,opt,name=magnitude_type,json=magnitudeType" json:"magnitude_type,omitempty"`
	Magnitudes       []*Magnitude  `protobuf:"bytes,24,rep,name=magnitudes" json:"magnitudes,omitempty"`
}

func (m *QuakeTechnical) Reset()                    { *m = QuakeTechnical{} }
func (m *QuakeTechnical) String() string            { return proto.CompactTextString(m) }
func (*QuakeTechnical) ProtoMessage()               {}
func (*QuakeTechnical) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *QuakeTechnical) GetTime() *Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *QuakeTechnical) GetModificationTime() *Timestamp {
	if m != nil {
		return m.ModificationTime
	}
	return nil
}

func (m *QuakeTechnical) GetLatitude() *RealQuantity {
	if m != nil {
		return m.Latitude
	}
	return nil
}

func (m *QuakeTechnical) GetLongitude() *RealQuantity {
	if m != nil {
		return m.Longitude
	}
	return nil
}

func (m *QuakeTechnical) GetDepth() *RealQuantity {
	if m != nil {
		return m.Depth
	}
	return nil
}

func (m *QuakeTechnical) GetPick() []*Pick {
	if m != nil {
		return m.Pick
	}
	return nil
}

func (m *QuakeTechnical) GetMagnitude() *RealQuantity {
	if m != nil {
		return m.Magnitude
	}
	return nil
}

func (m *QuakeTechnical) GetMagnitudes() []*Magnitude {
	if m != nil {
		return m.Magnitudes
	}
	return nil
}

type RealQuantity struct {
	Value       float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	Uncertainty float64 `protobuf:"fixed64,2,opt,name=uncertainty" json:"uncertainty,omitempty"`
}

func (m *RealQuantity) Reset()                    { *m = RealQuantity{} }
func (m *RealQuantity) String() string            { return proto.CompactTextString(m) }
func (*RealQuantity) ProtoMessage()               {}
func (*RealQuantity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type StationMagnitude struct {
	Waveform  *Waveform     `protobuf:"bytes,1,opt,name=waveform" json:"waveform,omitempty"`
	Magnitude *RealQuantity `protobuf:"bytes,2,opt,name=magnitude" json:"magnitude,omitempty"`
	Type      string        `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Azimuth   float64       `protobuf:"fixed64,4,opt,name=azimuth" json:"azimuth,omitempty"`
	Distance  float64       `protobuf:"fixed64,5,opt,name=distance" json:"distance,omitempty"`
	Residual  float64       `protobuf:"fixed64,6,opt,name=residual" json:"residual,omitempty"`
	Weight    float64       `protobuf:"fixed64,7,opt,name=weight" json:"weight,omitempty"`
	Amplitude *RealQuantity `protobuf:"bytes,8,opt,name=amplitude" json:"amplitude,omitempty"`
}

func (m *StationMagnitude) Reset()                    { *m = StationMagnitude{} }
func (m *StationMagnitude) String() string            { return proto.CompactTextString(m) }
func (*StationMagnitude) ProtoMessage()               {}
func (*StationMagnitude) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *StationMagnitude) GetWaveform() *Waveform {
	if m != nil {
		return m.Waveform
	}
	return nil
}

func (m *StationMagnitude) GetMagnitude() *RealQuantity {
	if m != nil {
		return m.Magnitude
	}
	return nil
}

func (m *StationMagnitude) GetAmplitude() *RealQuantity {
	if m != nil {
		return m.Amplitude
	}
	return nil
}

type Magnitude struct {
	Magnitude        *RealQuantity       `protobuf:"bytes,1,opt,name=magnitude" json:"magnitude,omitempty"`
	Type             string              `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	StationCount     int64               `protobuf:"varint,3,opt,name=station_count,json=stationCount" json:"station_count,omitempty"`
	StationMagnitude []*StationMagnitude `protobuf:"bytes,4,rep,name=station_magnitude,json=stationMagnitude" json:"station_magnitude,omitempty"`
}

func (m *Magnitude) Reset()                    { *m = Magnitude{} }
func (m *Magnitude) String() string            { return proto.CompactTextString(m) }
func (*Magnitude) ProtoMessage()               {}
func (*Magnitude) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Magnitude) GetMagnitude() *RealQuantity {
	if m != nil {
		return m.Magnitude
	}
	return nil
}

func (m *Magnitude) GetStationMagnitude() []*StationMagnitude {
	if m != nil {
		return m.StationMagnitude
	}
	return nil
}

type Pick struct {
	Waveform         *Waveform  `protobuf:"bytes,1,opt,name=waveform" json:"waveform,omitempty"`
	Time             *Timestamp `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	Phase            string     `protobuf:"bytes,3,opt,name=phase" json:"phase,omitempty"`
	Azimuth          float64    `protobuf:"fixed64,4,opt,name=azimuth" json:"azimuth,omitempty"`
	Distance         float64    `protobuf:"fixed64,5,opt,name=distance" json:"distance,omitempty"`
	Residual         float64    `protobuf:"fixed64,6,opt,name=residual" json:"residual,omitempty"`
	Weight           float64    `protobuf:"fixed64,7,opt,name=weight" json:"weight,omitempty"`
	EvaluationMode   string     `protobuf:"bytes,8,opt,name=evaluation_mode,json=evaluationMode" json:"evaluation_mode,omitempty"`
	EvaluationStatus string     `protobuf:"bytes,9,opt,name=evaluation_status,json=evaluationStatus" json:"evaluation_status,omitempty"`
}

func (m *Pick) Reset()                    { *m = Pick{} }
func (m *Pick) String() string            { return proto.CompactTextString(m) }
func (*Pick) ProtoMessage()               {}
func (*Pick) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Pick) GetWaveform() *Waveform {
	if m != nil {
		return m.Waveform
	}
	return nil
}

func (m *Pick) GetTime() *Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type Waveform struct {
	Network  string `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	Station  string `protobuf:"bytes,2,opt,name=station" json:"station,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	Channel  string `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
}

func (m *Waveform) Reset()                    { *m = Waveform{} }
func (m *Waveform) String() string            { return proto.CompactTextString(m) }
func (*Waveform) ProtoMessage()               {}
func (*Waveform) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

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
	proto.RegisterType((*QuakeTechnical)(nil), "haz.QuakeTechnical")
	proto.RegisterType((*RealQuantity)(nil), "haz.RealQuantity")
	proto.RegisterType((*StationMagnitude)(nil), "haz.StationMagnitude")
	proto.RegisterType((*Magnitude)(nil), "haz.Magnitude")
	proto.RegisterType((*Pick)(nil), "haz.Pick")
	proto.RegisterType((*Waveform)(nil), "haz.Waveform")
}

func init() { proto.RegisterFile("haz.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x57, 0xdb, 0x6e, 0x1c, 0x45,
	0x13, 0xd6, 0xee, 0xec, 0x69, 0xca, 0xf6, 0x7a, 0xdd, 0x7f, 0x92, 0x7f, 0x30, 0x89, 0x88, 0x06,
	0x10, 0x01, 0x12, 0x03, 0xe1, 0x82, 0x08, 0x14, 0x09, 0x82, 0x03, 0x8a, 0x84, 0x51, 0x32, 0xb6,
	0x12, 0xc1, 0xcd, 0xaa, 0xb3, 0xd3, 0xd9, 0x1d, 0xed, 0x1c, 0x36, 0x33, 0xbd, 0x36, 0x9b, 0xe7,
	0xe0, 0x96, 0x27, 0xe0, 0x09, 0x50, 0x6e, 0x78, 0x20, 0x1e, 0x82, 0xaa, 0xea, 0x9e, 0xc3, 0x5a,
	0xb6, 0x65, 0xdf, 0x70, 0xb5, 0x5d, 0x5f, 0x7d, 0x35, 0xdd, 0xfd, 0x55, 0x75, 0x75, 0x2f, 0xb8,
	0x33, 0xf9, 0x66, 0x6f, 0x91, 0x67, 0x3a, 0x13, 0x0e, 0x0e, 0xfd, 0xb7, 0x6d, 0xe8, 0x3e, 0x5b,
	0xca, 0xb9, 0x12, 0xef, 0x82, 0xbb, 0x58, 0xbe, 0x8c, 0xa3, 0xc9, 0x38, 0xda, 0xf7, 0x5a, 0xb7,
	0x5b, 0x77, 0xdc, 0x60, 0x60, 0x80, 0x27, 0xfb, 0xc2, 0x87, 0x8e, 0x8e, 0x12, 0xe5, 0xb5, 0x11,
	0xdf, 0xb8, 0x3f, 0xdc, 0xa3, 0xaf, 0x1c, 0x21, 0x50, 0x68, 0x99, 0x2c, 0x02, 0xf6, 0x89, 0x6f,
	0x60, 0x27, 0xc9, 0xc2, 0xe8, 0x55, 0x34, 0x91, 0x3a, 0xca, 0xd2, 0x31, 0x07, 0x38, 0x67, 0x06,
	0x8c, 0x9a, 0x44, 0x82, 0xc5, 0x2e, 0x0c, 0x62, 0xb4, 0xf4, 0x32, 0x54, 0x5e, 0x07, 0x63, 0x5a,
	0x41, 0x65, 0x8b, 0x9b, 0xe0, 0xc6, 0x59, 0x3a, 0x35, 0xce, 0x2e, 0x3b, 0x6b, 0x40, 0x5c, 0x83,
	0x6e, 0xa8, 0x16, 0x7a, 0xe6, 0xf5, 0xd8, 0x63, 0x0c, 0x8a, 0x49, 0xe4, 0x34, 0x35, 0x31, 0x7d,
	0x13, 0x53, 0x01, 0x3c, 0x5b, 0x36, 0x91, 0x71, 0xa4, 0x57, 0xde, 0xc0, 0x6c, 0xb5, 0xb4, 0x85,
	0x07, 0xfd, 0xd7, 0x4b, 0xe3, 0x72, 0xd9, 0x55, 0x9a, 0x62, 0x04, 0x4e, 0x92, 0x44, 0x1e, 0x20,
	0xda, 0x0d, 0x68, 0xe8, 0x7f, 0x01, 0x6e, 0xb5, 0x29, 0x72, 0x17, 0x6a, 0xc2, 0xd2, 0x39, 0x01,
	0x0d, 0x85, 0x80, 0x4e, 0x4a, 0x50, 0x9b, 0x21, 0x1e, 0xfb, 0x77, 0xa1, 0xc7, 0x7a, 0x17, 0xa8,
	0x69, 0xef, 0x35, 0x8f, 0x30, 0xc4, 0x41, 0x91, 0x80, 0x45, 0x62, 0x67, 0x60, 0x3d, 0xfe, 0xef,
	0x2d, 0xe8, 0x3f, 0xcf, 0xe2, 0x89, 0x4c, 0x33, 0x71, 0x0b, 0xe0, 0xd8, 0x0c, 0xeb, 0x0c, 0xb9,
	0x16, 0xc1, 0x14, 0xa1, 0x0e, 0xa8, 0x57, 0x6c, 0x72, 0xe4, 0x06, 0xc6, 0x58, 0xd3, 0xd5, 0xb9,
	0x48, 0xd7, 0xce, 0x69, 0x5d, 0x77, 0xc1, 0x39, 0x96, 0x31, 0xeb, 0xbd, 0x71, 0x7f, 0xc0, 0x6b,
	0x7b, 0xfe, 0xdd, 0x4f, 0x01, 0x81, 0xfe, 0x33, 0x70, 0x70, 0x4c, 0x53, 0xc6, 0xea, 0x58, 0xc5,
	0xbc, 0x98, 0x6e, 0x60, 0x0c, 0x9a, 0x52, 0x4e, 0x74, 0x74, 0x4c, 0x0a, 0x9a, 0xb5, 0x54, 0x36,
	0x89, 0x8b, 0x1f, 0x92, 0x79, 0x58, 0xf0, 0x6a, 0x50, 0x5c, 0x6b, 0xfa, 0x5f, 0x81, 0x6b, 0x37,
	0x8a, 0xd2, 0x7c, 0x02, 0xe5, 0xc6, 0x2a, 0x75, 0x36, 0xcd, 0x0a, 0x0c, 0x1a, 0xd4, 0x6e, 0x7f,
	0x0a, 0xce, 0xc1, 0xc1, 0x93, 0xb5, 0x8d, 0xb6, 0x2e, 0xda, 0x68, 0xfb, 0xf4, 0x46, 0x6d, 0x5a,
	0x9d, 0x2a, 0xad, 0xb4, 0xaf, 0x49, 0xb6, 0x4c, 0x35, 0x8b, 0x82, 0xfb, 0x62, 0xc3, 0xff, 0x1b,
	0x73, 0x71, 0x38, 0x93, 0xf3, 0x28, 0x9d, 0x92, 0x38, 0x14, 0x63, 0x96, 0x66, 0xc4, 0xc1, 0x45,
	0x98, 0xe8, 0x87, 0xb0, 0x81, 0x3f, 0xe3, 0x62, 0x99, 0x24, 0x32, 0x27, 0x09, 0x88, 0x73, 0x93,
	0x39, 0x36, 0x7c, 0xef, 0x20, 0x89, 0x0e, 0x8d, 0xfb, 0x71, 0xaa, 0xf3, 0x55, 0x00, 0x49, 0x05,
	0xd0, 0x39, 0xa4, 0x70, 0x9d, 0x69, 0x54, 0xdf, 0x2c, 0x6a, 0x80, 0xc0, 0x11, 0xd9, 0xbb, 0x0f,
	0x61, 0xfb, 0x54, 0x2c, 0x2d, 0x7f, 0xae, 0x56, 0x36, 0x05, 0x34, 0xa4, 0xe5, 0x63, 0x92, 0x96,
	0x66, 0xab, 0xb8, 0x7c, 0x36, 0xbe, 0x6e, 0x3f, 0x68, 0xf9, 0x63, 0xe8, 0x1e, 0xea, 0x2c, 0x67,
	0xca, 0x11, 0x17, 0x4b, 0xab, 0x59, 0x2c, 0x58, 0xaf, 0x71, 0x94, 0xce, 0x6d, 0xd6, 0x78, 0x2c,
	0xee, 0xda, 0xb6, 0x50, 0xcc, 0x54, 0x78, 0xce, 0x69, 0xae, 0x09, 0x58, 0xdd, 0x9d, 0x9f, 0xd5,
	0x49, 0x21, 0x3e, 0x80, 0x7e, 0x81, 0x13, 0x45, 0xa7, 0x8a, 0x9b, 0x27, 0x0f, 0x4a, 0x97, 0xff,
	0x2d, 0x74, 0x02, 0xa9, 0x55, 0xd5, 0x5d, 0x5a, 0x17, 0x74, 0x97, 0x2a, 0x27, 0xed, 0x66, 0x4e,
	0xfe, 0x69, 0x03, 0xf0, 0x89, 0x39, 0xd4, 0x52, 0xd3, 0x91, 0xea, 0x2f, 0x54, 0x3e, 0x0e, 0xe5,
	0xca, 0x4e, 0xeb, 0xf2, 0xb7, 0x68, 0x92, 0xa0, 0x87, 0x9e, 0x7d, 0xb9, 0x12, 0xf7, 0xa0, 0x73,
	0xa2, 0xd4, 0xdc, 0xe6, 0xe5, 0x9d, 0xfa, 0xd0, 0xf1, 0x27, 0xf6, 0x5e, 0xa0, 0xcf, 0x24, 0x85,
	0x69, 0xe2, 0x73, 0xe8, 0x26, 0x59, 0x8a, 0xed, 0xc5, 0x61, 0xfe, 0xee, 0x69, 0xfe, 0x01, 0x39,
	0x4d, 0x80, 0x21, 0xd2, 0x04, 0x2b, 0x25, 0x73, 0x2c, 0x9e, 0x33, 0x27, 0xf8, 0x05, 0x7d, 0x76,
	0x02, 0xa2, 0xed, 0x62, 0xe1, 0x57, 0x73, 0x5e, 0x25, 0x99, 0xbb, 0x0f, 0x00, 0xea, 0xc9, 0xaf,
	0x14, 0x89, 0x53, 0x56, 0xab, 0xb8, 0x52, 0xfd, 0xbc, 0xed, 0xc3, 0x90, 0xb7, 0x72, 0xa4, 0x26,
	0xb3, 0x14, 0xfb, 0x77, 0x7c, 0xf1, 0xb5, 0x81, 0x05, 0xa5, 0x57, 0x8b, 0xb2, 0x25, 0xf1, 0x58,
	0xdc, 0x80, 0x9e, 0x9c, 0xaa, 0x74, 0xb2, 0xb2, 0x1d, 0xc0, 0x5a, 0x55, 0x11, 0x74, 0xae, 0x7a,
	0xc5, 0x74, 0x2f, 0x79, 0xc5, 0xdc, 0x6b, 0x74, 0x88, 0x1e, 0xc7, 0xec, 0x98, 0xea, 0x50, 0x32,
	0xc6, 0x4d, 0xa5, 0xe8, 0x5a, 0x35, 0x9a, 0xc6, 0x67, 0xcd, 0xa6, 0xd1, 0x3f, 0x8f, 0xdf, 0xe8,
	0x23, 0x1f, 0x95, 0x17, 0xd1, 0xe0, 0x3c, 0xb2, 0xbd, 0x9b, 0xb0, 0x91, 0xf3, 0x60, 0xcc, 0xda,
	0x98, 0x4b, 0xc6, 0x65, 0xe4, 0xc8, 0x0a, 0x94, 0x28, 0x3d, 0xcb, 0x42, 0xbe, 0x69, 0x50, 0x20,
	0x63, 0x89, 0xf7, 0x60, 0x03, 0x93, 0x86, 0x61, 0xb8, 0x33, 0xec, 0xb9, 0x1b, 0xec, 0x04, 0x86,
	0x0e, 0x08, 0xc1, 0x05, 0x6c, 0x2b, 0xca, 0x95, 0xd1, 0x86, 0x58, 0xde, 0x26, 0x93, 0x86, 0x35,
	0x4c, 0x4c, 0xf1, 0x29, 0xec, 0x34, 0x88, 0xa8, 0x97, 0x5e, 0x16, 0xde, 0x16, 0x53, 0x47, 0xb5,
	0xe3, 0x90, 0x71, 0x71, 0x07, 0x46, 0xcb, 0x42, 0x85, 0xe3, 0xc5, 0x4c, 0x16, 0x6a, 0x6c, 0xce,
	0xe0, 0x90, 0x2f, 0xb4, 0x21, 0xe1, 0x4f, 0x09, 0xfe, 0x9e, 0x50, 0x6c, 0x15, 0x82, 0x99, 0xf4,
	0x41, 0xfa, 0xb0, 0xe1, 0x6e, 0x33, 0x97, 0xbf, 0x71, 0x68, 0x1c, 0x86, 0xfd, 0x21, 0x0c, 0x91,
	0x98, 0x86, 0xd8, 0xfd, 0xc7, 0x2a, 0xcf, 0xb3, 0xdc, 0x1b, 0x71, 0x67, 0xde, 0x2a, 0xd1, 0xc7,
	0x04, 0x8a, 0xf7, 0x61, 0x4b, 0xbe, 0x89, 0x92, 0xa5, 0x9e, 0xc9, 0x78, 0x3c, 0x95, 0x0b, 0x6f,
	0x87, 0x59, 0x9b, 0x15, 0xf8, 0xa3, 0x5c, 0x88, 0x8f, 0x61, 0x94, 0x44, 0x29, 0x02, 0xc9, 0x38,
	0x8c, 0x28, 0x7e, 0xa2, 0x3c, 0xc1, 0xbc, 0x6d, 0x8b, 0xef, 0x5b, 0x98, 0xa9, 0xf2, 0xb7, 0x75,
	0xea, 0xff, 0x2c, 0xd5, 0xe0, 0x15, 0x15, 0xf5, 0x4c, 0x54, 0x18, 0xc9, 0xb4, 0x66, 0x5e, 0x63,
	0xe6, 0xd0, 0xc0, 0x15, 0xf1, 0x16, 0x74, 0x16, 0xd1, 0x64, 0xee, 0x5d, 0x6f, 0xf4, 0x9c, 0xa7,
	0x08, 0x04, 0x0c, 0x53, 0x25, 0xd5, 0x6f, 0x91, 0x1b, 0xe7, 0x56, 0x52, 0xfd, 0x3c, 0x41, 0x69,
	0x2a, 0xc3, 0x14, 0xc9, 0xff, 0x39, 0x39, 0x5b, 0x15, 0xca, 0x85, 0xb2, 0x07, 0x50, 0x01, 0x85,
	0xe7, 0xf1, 0xe4, 0xe6, 0x18, 0x1c, 0x94, 0x70, 0xd0, 0x60, 0xf8, 0x3f, 0xc0, 0x66, 0x73, 0xc6,
	0xfa, 0x9c, 0x9b, 0xfb, 0xd2, 0x18, 0xe2, 0x36, 0x6c, 0x2c, 0x71, 0x53, 0xb9, 0x96, 0x51, 0x6a,
	0x6f, 0xf0, 0x56, 0xd0, 0x84, 0xfc, 0x3f, 0xda, 0x30, 0xb2, 0xa9, 0xac, 0x26, 0x42, 0x5d, 0x07,
	0x27, 0xf2, 0x58, 0xbd, 0xca, 0xf2, 0xc4, 0xf6, 0xf1, 0x2d, 0x5e, 0xca, 0x0b, 0x0b, 0x06, 0x95,
	0x7b, 0x5d, 0x8f, 0xf6, 0x25, 0xf4, 0x28, 0xdb, 0x88, 0xd3, 0x68, 0x23, 0xf8, 0x92, 0xb0, 0x25,
	0x60, 0x9f, 0x2e, 0xa5, 0x49, 0x2f, 0x81, 0x2a, 0x5f, 0xe6, 0xb5, 0x58, 0xd9, 0xe4, 0xcb, 0x55,
	0x11, 0x85, 0xf8, 0xa2, 0xb3, 0xef, 0xc5, 0xca, 0xa6, 0x73, 0x77, 0xa2, 0xa2, 0xe9, 0x4c, 0xdb,
	0xf7, 0xa2, 0xb5, 0x68, 0xb9, 0xd8, 0x50, 0x62, 0xb3, 0xdc, 0x73, 0xcf, 0x76, 0xcd, 0xf1, 0xff,
	0x6a, 0x81, 0x5b, 0x0b, 0xb3, 0xb6, 0xdb, 0xd6, 0x15, 0x76, 0xdb, 0x6c, 0x9a, 0x78, 0x0a, 0xd6,
	0x4f, 0x95, 0xc3, 0xa7, 0x6a, 0xb3, 0x68, 0x9e, 0xa8, 0x47, 0xb0, 0x53, 0x92, 0xea, 0x19, 0xcd,
	0x2d, 0x74, 0xdd, 0x5e, 0xbf, 0xeb, 0x49, 0x0b, 0x46, 0xc5, 0x29, 0xc4, 0xff, 0xb3, 0x0d, 0x1d,
	0x2a, 0xdd, 0xab, 0xe4, 0xf3, 0x32, 0x7f, 0x0e, 0xb0, 0xd6, 0xb8, 0x81, 0xd8, 0x1c, 0x1a, 0xe3,
	0x3f, 0x4c, 0xe2, 0x19, 0xbd, 0x71, 0x70, 0xf9, 0xde, 0xe8, 0x9e, 0xdd, 0x1b, 0x7d, 0x0d, 0x83,
	0x52, 0x0f, 0xda, 0x4b, 0xaa, 0xf4, 0x49, 0x96, 0xcf, 0xed, 0x35, 0x58, 0x9a, 0xe4, 0xb1, 0x3a,
	0xdb, 0x9c, 0x96, 0x66, 0xf9, 0x3f, 0x84, 0x5d, 0x4e, 0xfd, 0x3f, 0x84, 0x7d, 0x18, 0x35, 0x99,
	0xc9, 0x34, 0xc5, 0x56, 0xdf, 0x31, 0x51, 0xd6, 0x7c, 0xd4, 0xfd, 0x95, 0xfe, 0xba, 0xbd, 0xec,
	0xf1, 0xdf, 0xb8, 0x2f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x99, 0xb2, 0x1d, 0x22, 0xd3, 0x0d,
	0x00, 0x00,
}
