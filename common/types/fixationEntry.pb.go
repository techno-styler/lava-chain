// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/fixationEntry.proto

package types

import (
	fmt "fmt"
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

type Entry struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Block    uint64 `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	StaleAt  uint64 `protobuf:"varint,3,opt,name=stale_at,json=staleAt,proto3" json:"stale_at,omitempty"`
	Refcount uint64 `protobuf:"varint,4,opt,name=refcount,proto3" json:"refcount,omitempty"`
	Data     []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_702478ef0512c95d, []int{0}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return m.Size()
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Entry) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Entry) GetStaleAt() uint64 {
	if m != nil {
		return m.StaleAt
	}
	return 0
}

func (m *Entry) GetRefcount() uint64 {
	if m != nil {
		return m.Refcount
	}
	return 0
}

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Entry)(nil), "lavanet.lava.common.Entry")
}

func init() { proto.RegisterFile("common/fixationEntry.proto", fileDescriptor_702478ef0512c95d) }

var fileDescriptor_702478ef0512c95d = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x4f, 0xcb, 0xac, 0x48, 0x2c, 0xc9, 0xcc, 0xcf, 0x73, 0xcd, 0x2b, 0x29, 0xaa,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xce, 0x49, 0x2c, 0x4b, 0xcc, 0x4b, 0x2d, 0xd1,
	0x03, 0xd1, 0x7a, 0x10, 0x85, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x79, 0x7d, 0x10, 0x0b,
	0xa2, 0x54, 0xa9, 0x8e, 0x8b, 0x15, 0xac, 0x53, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5,
	0x42, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x01, 0x89, 0x26, 0xe5, 0xe4, 0x27, 0x67,
	0x4b, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x04, 0x41, 0x38, 0x42, 0x92, 0x5c, 0x1c, 0xc5, 0x25, 0x89,
	0x39, 0xa9, 0xf1, 0x89, 0x25, 0x12, 0xcc, 0x60, 0x09, 0x76, 0x30, 0xdf, 0xb1, 0x44, 0x48, 0x8a,
	0x8b, 0xa3, 0x28, 0x35, 0x2d, 0x39, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x05, 0x2c, 0x05, 0xe7, 0x0b,
	0x09, 0x71, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x4a, 0xb0, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9,
	0x4e, 0x76, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x92, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0x04, 0x72, 0xbe, 0x3e, 0xd4, 0x3f, 0x60, 0x5a, 0x1f, 0xea, 0xf1, 0x92, 0xca,
	0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x37, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0xfc,
	0x8b, 0xbd, 0x0f, 0x01, 0x00, 0x00,
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Entry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintFixationEntry(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Refcount != 0 {
		i = encodeVarintFixationEntry(dAtA, i, uint64(m.Refcount))
		i--
		dAtA[i] = 0x20
	}
	if m.StaleAt != 0 {
		i = encodeVarintFixationEntry(dAtA, i, uint64(m.StaleAt))
		i--
		dAtA[i] = 0x18
	}
	if m.Block != 0 {
		i = encodeVarintFixationEntry(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintFixationEntry(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFixationEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovFixationEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovFixationEntry(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovFixationEntry(uint64(m.Block))
	}
	if m.StaleAt != 0 {
		n += 1 + sovFixationEntry(uint64(m.StaleAt))
	}
	if m.Refcount != 0 {
		n += 1 + sovFixationEntry(uint64(m.Refcount))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovFixationEntry(uint64(l))
	}
	return n
}

func sovFixationEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFixationEntry(x uint64) (n int) {
	return sovFixationEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFixationEntry
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
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixationEntry
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
				return ErrInvalidLengthFixationEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFixationEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixationEntry
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaleAt", wireType)
			}
			m.StaleAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixationEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StaleAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refcount", wireType)
			}
			m.Refcount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixationEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Refcount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixationEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFixationEntry
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFixationEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFixationEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFixationEntry
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
func skipFixationEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFixationEntry
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
					return 0, ErrIntOverflowFixationEntry
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
					return 0, ErrIntOverflowFixationEntry
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
				return 0, ErrInvalidLengthFixationEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFixationEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFixationEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFixationEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFixationEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFixationEntry = fmt.Errorf("proto: unexpected end of group")
)
