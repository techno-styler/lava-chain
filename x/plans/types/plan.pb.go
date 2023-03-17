// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plans/plan.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Plan struct {
	Index                    string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Block                    uint64     `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
	Price                    types.Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
	ComputeUnits             uint64     `protobuf:"varint,5,opt,name=compute_units,json=computeUnits,proto3" json:"compute_units,omitempty"`
	ComputeUnitsPerEpoch     uint64     `protobuf:"varint,6,opt,name=compute_units_per_epoch,json=computeUnitsPerEpoch,proto3" json:"compute_units_per_epoch,omitempty"`
	MaxProvidersToPair       uint64     `protobuf:"varint,7,opt,name=max_providers_to_pair,json=maxProvidersToPair,proto3" json:"max_providers_to_pair,omitempty"`
	AllowOveruse             bool       `protobuf:"varint,8,opt,name=allow_overuse,json=allowOveruse,proto3" json:"allow_overuse,omitempty"`
	OveruseRate              uint64     `protobuf:"varint,9,opt,name=overuse_rate,json=overuseRate,proto3" json:"overuse_rate,omitempty"`
	Description              string     `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Type                     string     `protobuf:"bytes,12,opt,name=type,proto3" json:"type,omitempty"`
	AnnualDiscountPercentage uint64     `protobuf:"varint,13,opt,name=annual_discount_percentage,json=annualDiscountPercentage,proto3" json:"annual_discount_percentage,omitempty"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5909a10cd0e3497, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Plan) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Plan) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

func (m *Plan) GetComputeUnits() uint64 {
	if m != nil {
		return m.ComputeUnits
	}
	return 0
}

func (m *Plan) GetComputeUnitsPerEpoch() uint64 {
	if m != nil {
		return m.ComputeUnitsPerEpoch
	}
	return 0
}

func (m *Plan) GetMaxProvidersToPair() uint64 {
	if m != nil {
		return m.MaxProvidersToPair
	}
	return 0
}

func (m *Plan) GetAllowOveruse() bool {
	if m != nil {
		return m.AllowOveruse
	}
	return false
}

func (m *Plan) GetOveruseRate() uint64 {
	if m != nil {
		return m.OveruseRate
	}
	return 0
}

func (m *Plan) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Plan) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Plan) GetAnnualDiscountPercentage() uint64 {
	if m != nil {
		return m.AnnualDiscountPercentage
	}
	return 0
}

func init() {
	proto.RegisterType((*Plan)(nil), "lavanet.lava.plans.Plan")
}

func init() { proto.RegisterFile("plans/plan.proto", fileDescriptor_e5909a10cd0e3497) }

var fileDescriptor_e5909a10cd0e3497 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x6d, 0x68, 0xba, 0x74, 0x9d, 0xae, 0x54, 0x59, 0x45, 0x98, 0x1e, 0x42, 0x00, 0x21, 0xf5,
	0x94, 0xa8, 0xa0, 0xbd, 0x71, 0xda, 0x85, 0xcb, 0x5e, 0x88, 0x22, 0xb8, 0x70, 0xb1, 0x1c, 0xd7,
	0xea, 0x5a, 0x24, 0x1e, 0xcb, 0x76, 0x4a, 0xf9, 0x0b, 0x3e, 0x83, 0xcf, 0xe0, 0xb8, 0xc7, 0x3d,
	0x72, 0x42, 0xa8, 0xfd, 0x11, 0x64, 0x27, 0xa0, 0x72, 0xf1, 0xcc, 0xbc, 0xf7, 0x3c, 0x6f, 0x6c,
	0x0d, 0x9a, 0xeb, 0x86, 0x29, 0x5b, 0xf8, 0x33, 0xd7, 0x06, 0x1c, 0x60, 0xdc, 0xb0, 0x1d, 0x53,
	0xc2, 0xe5, 0x3e, 0xe6, 0x81, 0x5e, 0x2e, 0xb6, 0xb0, 0x85, 0x40, 0x17, 0x3e, 0xeb, 0x95, 0xcb,
	0x94, 0x83, 0x6d, 0xc1, 0x16, 0x35, 0xb3, 0xa2, 0xd8, 0xad, 0x6b, 0xe1, 0xd8, 0xba, 0xe0, 0x20,
	0x87, 0x4e, 0xcf, 0x7f, 0x8c, 0x51, 0x5c, 0x36, 0x4c, 0xe1, 0x05, 0x9a, 0x48, 0xb5, 0x11, 0x7b,
	0x12, 0x65, 0xd1, 0xea, 0xbc, 0xea, 0x0b, 0x8f, 0xd6, 0x0d, 0xf0, 0xcf, 0x64, 0x9c, 0x45, 0xab,
	0xb8, 0xea, 0x0b, 0x7c, 0x89, 0x26, 0xda, 0x48, 0x2e, 0x48, 0x9c, 0x45, 0xab, 0xe4, 0xd5, 0x93,
	0xbc, 0x37, 0xc9, 0xbd, 0x49, 0x3e, 0x98, 0xe4, 0xd7, 0x20, 0xd5, 0x55, 0x7c, 0xf7, 0xeb, 0xe9,
	0xa8, 0xea, 0xd5, 0xf8, 0x05, 0xba, 0xe0, 0xd0, 0xea, 0xce, 0x09, 0xda, 0x29, 0xe9, 0x2c, 0x99,
	0x84, 0xa6, 0xb3, 0x01, 0xfc, 0xe8, 0x31, 0x7c, 0x89, 0x1e, 0xff, 0x27, 0xa2, 0x5a, 0x18, 0x2a,
	0x34, 0xf0, 0x5b, 0x72, 0x16, 0xe4, 0x8b, 0x53, 0x79, 0x29, 0xcc, 0x3b, 0xcf, 0xe1, 0x35, 0x7a,
	0xd4, 0xb2, 0x3d, 0xd5, 0x06, 0x76, 0x72, 0x23, 0x8c, 0xa5, 0x0e, 0xa8, 0x66, 0xd2, 0x90, 0x87,
	0xe1, 0x12, 0x6e, 0xd9, 0xbe, 0xfc, 0xcb, 0x7d, 0x80, 0x92, 0x49, 0xe3, 0xc7, 0x61, 0x4d, 0x03,
	0x5f, 0x28, 0xec, 0x84, 0xe9, 0xac, 0x20, 0xd3, 0x2c, 0x5a, 0x4d, 0xab, 0x59, 0x00, 0xdf, 0xf7,
	0x18, 0x7e, 0x86, 0x66, 0x03, 0x4d, 0x0d, 0x73, 0x82, 0x9c, 0x87, 0x76, 0xc9, 0x80, 0x55, 0xcc,
	0x09, 0x9c, 0xa1, 0x64, 0x23, 0x2c, 0x37, 0x52, 0x3b, 0x09, 0x8a, 0x24, 0xe1, 0xff, 0x4e, 0x21,
	0x8c, 0x51, 0xec, 0xbe, 0x6a, 0x41, 0x66, 0x81, 0x0a, 0x39, 0x7e, 0x83, 0x96, 0x4c, 0xa9, 0x8e,
	0x35, 0x74, 0x23, 0x2d, 0x87, 0x4e, 0x39, 0xff, 0x52, 0x2e, 0x94, 0x63, 0x5b, 0x41, 0x2e, 0x82,
	0x0d, 0xe9, 0x15, 0x6f, 0x07, 0x41, 0xf9, 0x8f, 0xbf, 0x89, 0xa7, 0x68, 0x9e, 0xdc, 0xc4, 0xd3,
	0x07, 0xf3, 0xf1, 0xd5, 0xf5, 0xf7, 0x43, 0x1a, 0xdd, 0x1d, 0xd2, 0xe8, 0xfe, 0x90, 0x46, 0xbf,
	0x0f, 0x69, 0xf4, 0xed, 0x98, 0x8e, 0xee, 0x8f, 0xe9, 0xe8, 0xe7, 0x31, 0x1d, 0x7d, 0x7a, 0xb9,
	0x95, 0xee, 0xb6, 0xab, 0x73, 0x0e, 0x6d, 0x31, 0x6c, 0x4d, 0x88, 0xc5, 0xbe, 0xe8, 0xd7, 0xca,
	0x4f, 0x63, 0xeb, 0xb3, 0xb0, 0x0e, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x19, 0x0a, 0x79,
	0xc2, 0x6c, 0x02, 0x00, 0x00,
}

func (this *Plan) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Plan)
	if !ok {
		that2, ok := that.(Plan)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if this.Block != that1.Block {
		return false
	}
	if !this.Price.Equal(&that1.Price) {
		return false
	}
	if this.ComputeUnits != that1.ComputeUnits {
		return false
	}
	if this.ComputeUnitsPerEpoch != that1.ComputeUnitsPerEpoch {
		return false
	}
	if this.MaxProvidersToPair != that1.MaxProvidersToPair {
		return false
	}
	if this.AllowOveruse != that1.AllowOveruse {
		return false
	}
	if this.OveruseRate != that1.OveruseRate {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.AnnualDiscountPercentage != that1.AnnualDiscountPercentage {
		return false
	}
	return true
}
func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AnnualDiscountPercentage != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.AnnualDiscountPercentage))
		i--
		dAtA[i] = 0x68
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x5a
	}
	if m.OveruseRate != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.OveruseRate))
		i--
		dAtA[i] = 0x48
	}
	if m.AllowOveruse {
		i--
		if m.AllowOveruse {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.MaxProvidersToPair != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.MaxProvidersToPair))
		i--
		dAtA[i] = 0x38
	}
	if m.ComputeUnitsPerEpoch != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.ComputeUnitsPerEpoch))
		i--
		dAtA[i] = 0x30
	}
	if m.ComputeUnits != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.ComputeUnits))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Block != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovPlan(uint64(m.Block))
	}
	l = m.Price.Size()
	n += 1 + l + sovPlan(uint64(l))
	if m.ComputeUnits != 0 {
		n += 1 + sovPlan(uint64(m.ComputeUnits))
	}
	if m.ComputeUnitsPerEpoch != 0 {
		n += 1 + sovPlan(uint64(m.ComputeUnitsPerEpoch))
	}
	if m.MaxProvidersToPair != 0 {
		n += 1 + sovPlan(uint64(m.MaxProvidersToPair))
	}
	if m.AllowOveruse {
		n += 2
	}
	if m.OveruseRate != 0 {
		n += 1 + sovPlan(uint64(m.OveruseRate))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.AnnualDiscountPercentage != 0 {
		n += 1 + sovPlan(uint64(m.AnnualDiscountPercentage))
	}
	return n
}

func sovPlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlan(x uint64) (n int) {
	return sovPlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnits", wireType)
			}
			m.ComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnitsPerEpoch", wireType)
			}
			m.ComputeUnitsPerEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnitsPerEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxProvidersToPair", wireType)
			}
			m.MaxProvidersToPair = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxProvidersToPair |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowOveruse", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowOveruse = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OveruseRate", wireType)
			}
			m.OveruseRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OveruseRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualDiscountPercentage", wireType)
			}
			m.AnnualDiscountPercentage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnnualDiscountPercentage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlan
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPlan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlan = fmt.Errorf("proto: unexpected end of group")
)
