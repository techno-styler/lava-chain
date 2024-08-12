// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/rewards/base_pay.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// aggregated rewards for the provider through out the month
type BasePay struct {
	Total         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=total,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total"`
	TotalAdjusted github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=totalAdjusted,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"totalAdjusted"`
	IprpcCu       uint64                                 `protobuf:"varint,3,opt,name=iprpc_cu,json=iprpcCu,proto3" json:"iprpc_cu,omitempty"`
}

func (m *BasePay) Reset()         { *m = BasePay{} }
func (m *BasePay) String() string { return proto.CompactTextString(m) }
func (*BasePay) ProtoMessage()    {}
func (*BasePay) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2fb0eb917a4ee4e, []int{0}
}
func (m *BasePay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasePay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasePay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasePay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasePay.Merge(m, src)
}
func (m *BasePay) XXX_Size() int {
	return m.Size()
}
func (m *BasePay) XXX_DiscardUnknown() {
	xxx_messageInfo_BasePay.DiscardUnknown(m)
}

var xxx_messageInfo_BasePay proto.InternalMessageInfo

func (m *BasePay) GetIprpcCu() uint64 {
	if m != nil {
		return m.IprpcCu
	}
	return 0
}

// aggregated rewards for the provider through out the month
type BasePayWithIndex struct {
	Provider string  `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	ChainId  string  `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	BasePay  BasePay `protobuf:"bytes,3,opt,name=base_pay,json=basePay,proto3" json:"base_pay"`
}

func (m *BasePayWithIndex) Reset()         { *m = BasePayWithIndex{} }
func (m *BasePayWithIndex) String() string { return proto.CompactTextString(m) }
func (*BasePayWithIndex) ProtoMessage()    {}
func (*BasePayWithIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2fb0eb917a4ee4e, []int{1}
}
func (m *BasePayWithIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasePayWithIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasePayWithIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasePayWithIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasePayWithIndex.Merge(m, src)
}
func (m *BasePayWithIndex) XXX_Size() int {
	return m.Size()
}
func (m *BasePayWithIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_BasePayWithIndex.DiscardUnknown(m)
}

var xxx_messageInfo_BasePayWithIndex proto.InternalMessageInfo

func (m *BasePayWithIndex) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *BasePayWithIndex) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *BasePayWithIndex) GetBasePay() BasePay {
	if m != nil {
		return m.BasePay
	}
	return BasePay{}
}

func init() {
	proto.RegisterType((*BasePay)(nil), "lavanet.lava.rewards.BasePay")
	proto.RegisterType((*BasePayWithIndex)(nil), "lavanet.lava.rewards.BasePayWithIndex")
}

func init() {
	proto.RegisterFile("lavanet/lava/rewards/base_pay.proto", fileDescriptor_a2fb0eb917a4ee4e)
}

var fileDescriptor_a2fb0eb917a4ee4e = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0xcd, 0xbc, 0xd7, 0xf7, 0x52, 0x47, 0x04, 0x0d, 0x5d, 0xb4, 0x05, 0xd3, 0x52, 0x41, 0x8a,
	0xd0, 0x19, 0x5a, 0xb7, 0x22, 0x58, 0xbb, 0xc9, 0x4e, 0x82, 0x20, 0xb8, 0x09, 0x93, 0x64, 0x68,
	0x47, 0xdb, 0x4c, 0x48, 0x26, 0xb1, 0xfd, 0x04, 0x77, 0x7e, 0x8c, 0x1f, 0xd1, 0x95, 0x14, 0x57,
	0xe2, 0xa2, 0x48, 0xfb, 0x23, 0x92, 0x99, 0xb1, 0x58, 0x70, 0xe3, 0x26, 0xf7, 0x9e, 0x7b, 0x6f,
	0xce, 0x9d, 0xc3, 0xb9, 0xf0, 0x68, 0x4c, 0x72, 0x12, 0x51, 0x81, 0x8b, 0x88, 0x13, 0xfa, 0x40,
	0x92, 0x30, 0xc5, 0x3e, 0x49, 0xa9, 0x17, 0x93, 0x19, 0x8a, 0x13, 0x2e, 0xb8, 0x55, 0xd1, 0x43,
	0xa8, 0x88, 0x48, 0x0f, 0xd5, 0x0f, 0xc8, 0x84, 0x45, 0x1c, 0xcb, 0xaf, 0x1a, 0xac, 0xdb, 0x01,
	0x4f, 0x27, 0x5c, 0xfd, 0x8f, 0xf3, 0xae, 0x4f, 0x05, 0xe9, 0xe2, 0x80, 0xb3, 0x48, 0xf7, 0x6b,
	0xaa, 0xef, 0x49, 0x84, 0x15, 0xd0, 0xad, 0xca, 0x90, 0x0f, 0xb9, 0xaa, 0x17, 0x99, 0xaa, 0xb6,
	0x5e, 0x00, 0x34, 0xfb, 0x24, 0xa5, 0x57, 0x64, 0x66, 0xb9, 0xf0, 0x9f, 0xe0, 0x82, 0x8c, 0xab,
	0xa0, 0x09, 0xda, 0x3b, 0xfd, 0xb3, 0xf9, 0xb2, 0x61, 0xbc, 0x2f, 0x1b, 0xc7, 0x43, 0x26, 0x46,
	0x99, 0x8f, 0x02, 0x3e, 0xd1, 0x8c, 0x3a, 0x74, 0xd2, 0xf0, 0x1e, 0x8b, 0x59, 0x4c, 0x53, 0xe4,
	0x44, 0xe2, 0xf5, 0xb9, 0x03, 0xf5, 0x42, 0x27, 0x12, 0xae, 0xa2, 0xb2, 0xae, 0xe1, 0x9e, 0x4c,
	0x2e, 0xc2, 0xbb, 0x2c, 0x15, 0x34, 0xac, 0xfe, 0x91, 0xdc, 0xe8, 0x17, 0xdc, 0x03, 0x1a, 0xb8,
	0xdb, 0x24, 0x56, 0x0d, 0x96, 0x59, 0x9c, 0xc4, 0x81, 0x17, 0x64, 0xd5, 0xbf, 0x4d, 0xd0, 0x2e,
	0xb9, 0xa6, 0xc4, 0x97, 0x59, 0xeb, 0x11, 0xc0, 0x7d, 0x2d, 0xe8, 0x86, 0x89, 0x91, 0x13, 0x85,
	0x74, 0x6a, 0xd5, 0x61, 0x39, 0x4e, 0x78, 0xce, 0x42, 0x9a, 0x28, 0x71, 0xee, 0x06, 0x17, 0x5c,
	0xc1, 0x88, 0xb0, 0xc8, 0x63, 0xfa, 0x71, 0xae, 0x29, 0xb1, 0x13, 0x5a, 0xe7, 0xb0, 0xfc, 0x65,
	0x94, 0x5c, 0xb3, 0xdb, 0x3b, 0x44, 0x3f, 0x39, 0x85, 0xf4, 0xc2, 0x7e, 0xa9, 0x90, 0xe5, 0x9a,
	0xbe, 0x86, 0x83, 0xf9, 0xca, 0x06, 0x8b, 0x95, 0x0d, 0x3e, 0x56, 0x36, 0x78, 0x5a, 0xdb, 0xc6,
	0x62, 0x6d, 0x1b, 0x6f, 0x6b, 0xdb, 0xb8, 0x3d, 0xf9, 0xa6, 0x7b, 0xeb, 0x40, 0xf2, 0x1e, 0x9e,
	0x6e, 0xae, 0x44, 0xea, 0xf7, 0xff, 0x4b, 0xa7, 0x4e, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x57,
	0x67, 0x61, 0x46, 0x4a, 0x02, 0x00, 0x00,
}

func (m *BasePay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasePay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasePay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IprpcCu != 0 {
		i = encodeVarintBasePay(dAtA, i, uint64(m.IprpcCu))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.TotalAdjusted.Size()
		i -= size
		if _, err := m.TotalAdjusted.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBasePay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Total.Size()
		i -= size
		if _, err := m.Total.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBasePay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BasePayWithIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasePayWithIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasePayWithIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BasePay.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBasePay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintBasePay(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintBasePay(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBasePay(dAtA []byte, offset int, v uint64) int {
	offset -= sovBasePay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BasePay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Total.Size()
	n += 1 + l + sovBasePay(uint64(l))
	l = m.TotalAdjusted.Size()
	n += 1 + l + sovBasePay(uint64(l))
	if m.IprpcCu != 0 {
		n += 1 + sovBasePay(uint64(m.IprpcCu))
	}
	return n
}

func (m *BasePayWithIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovBasePay(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovBasePay(uint64(l))
	}
	l = m.BasePay.Size()
	n += 1 + l + sovBasePay(uint64(l))
	return n
}

func sovBasePay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBasePay(x uint64) (n int) {
	return sovBasePay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BasePay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBasePay
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
			return fmt.Errorf("proto: BasePay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasePay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Total.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAdjusted", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalAdjusted.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IprpcCu", wireType)
			}
			m.IprpcCu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IprpcCu |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBasePay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBasePay
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
func (m *BasePayWithIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBasePay
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
			return fmt.Errorf("proto: BasePayWithIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasePayWithIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasePay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BasePay.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBasePay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBasePay
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
func skipBasePay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBasePay
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
					return 0, ErrIntOverflowBasePay
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
					return 0, ErrIntOverflowBasePay
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
				return 0, ErrInvalidLengthBasePay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBasePay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBasePay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBasePay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBasePay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBasePay = fmt.Errorf("proto: unexpected end of group")
)
