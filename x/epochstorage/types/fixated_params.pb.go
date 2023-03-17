// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: epochstorage/fixated_params.proto

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

type FixatedParams struct {
	Index         string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Parameter     []byte `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty"`
	FixationBlock uint64 `protobuf:"varint,3,opt,name=fixationBlock,proto3" json:"fixationBlock,omitempty"`
}

func (m *FixatedParams) Reset()         { *m = FixatedParams{} }
func (m *FixatedParams) String() string { return proto.CompactTextString(m) }
func (*FixatedParams) ProtoMessage()    {}
func (*FixatedParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_899e892bbf90433a, []int{0}
}
func (m *FixatedParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FixatedParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FixatedParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FixatedParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixatedParams.Merge(m, src)
}
func (m *FixatedParams) XXX_Size() int {
	return m.Size()
}
func (m *FixatedParams) XXX_DiscardUnknown() {
	xxx_messageInfo_FixatedParams.DiscardUnknown(m)
}

var xxx_messageInfo_FixatedParams proto.InternalMessageInfo

func (m *FixatedParams) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *FixatedParams) GetParameter() []byte {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func (m *FixatedParams) GetFixationBlock() uint64 {
	if m != nil {
		return m.FixationBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*FixatedParams)(nil), "lavanet.lava.epochstorage.FixatedParams")
}

func init() { proto.RegisterFile("epochstorage/fixated_params.proto", fileDescriptor_899e892bbf90433a) }

var fileDescriptor_899e892bbf90433a = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2d, 0xc8, 0x4f,
	0xce, 0x28, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x4f, 0xcb, 0xac, 0x48, 0x2c, 0x49, 0x4d,
	0x89, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xcc,
	0x49, 0x2c, 0x4b, 0xcc, 0x4b, 0x2d, 0xd1, 0x03, 0xd1, 0x7a, 0xc8, 0xea, 0xa5, 0x44, 0xd2, 0xf3,
	0xd3, 0xf3, 0xc1, 0xaa, 0xf4, 0x41, 0x2c, 0x88, 0x06, 0x29, 0x49, 0x14, 0x33, 0x91, 0xcd, 0x52,
	0xca, 0xe4, 0xe2, 0x75, 0x83, 0xd8, 0x11, 0x00, 0x16, 0x16, 0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x4b,
	0x49, 0xad, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x64, 0xb8, 0x38, 0xc1,
	0xda, 0x52, 0x4b, 0x52, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0x10, 0x02, 0x42, 0x2a,
	0x5c, 0xbc, 0x60, 0x87, 0x66, 0xe6, 0xe7, 0x39, 0xe5, 0xe4, 0x27, 0x67, 0x4b, 0x30, 0x2b, 0x30,
	0x6a, 0xb0, 0x04, 0xa1, 0x0a, 0x3a, 0xb9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x4e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd4, 0x6f,
	0x60, 0x5a, 0xbf, 0x42, 0x1f, 0xc5, 0xe5, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x97,
	0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x02, 0x58, 0x8d, 0xb0, 0x2a, 0x01, 0x00, 0x00,
}

func (m *FixatedParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FixatedParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FixatedParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FixationBlock != 0 {
		i = encodeVarintFixatedParams(dAtA, i, uint64(m.FixationBlock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Parameter) > 0 {
		i -= len(m.Parameter)
		copy(dAtA[i:], m.Parameter)
		i = encodeVarintFixatedParams(dAtA, i, uint64(len(m.Parameter)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintFixatedParams(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFixatedParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovFixatedParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FixatedParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovFixatedParams(uint64(l))
	}
	l = len(m.Parameter)
	if l > 0 {
		n += 1 + l + sovFixatedParams(uint64(l))
	}
	if m.FixationBlock != 0 {
		n += 1 + sovFixatedParams(uint64(m.FixationBlock))
	}
	return n
}

func sovFixatedParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFixatedParams(x uint64) (n int) {
	return sovFixatedParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FixatedParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFixatedParams
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
			return fmt.Errorf("proto: FixatedParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FixatedParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixatedParams
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
				return ErrInvalidLengthFixatedParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFixatedParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixatedParams
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
				return ErrInvalidLengthFixatedParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFixatedParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameter = append(m.Parameter[:0], dAtA[iNdEx:postIndex]...)
			if m.Parameter == nil {
				m.Parameter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixationBlock", wireType)
			}
			m.FixationBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixatedParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FixationBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFixatedParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFixatedParams
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
func skipFixatedParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFixatedParams
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
					return 0, ErrIntOverflowFixatedParams
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
					return 0, ErrIntOverflowFixatedParams
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
				return 0, ErrInvalidLengthFixatedParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFixatedParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFixatedParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFixatedParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFixatedParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFixatedParams = fmt.Errorf("proto: unexpected end of group")
)
