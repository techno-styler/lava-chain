// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/rewards/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	MinBondedTarget                     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=min_bonded_target,json=minBondedTarget,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_bonded_target" yaml:"min_bonded_target"`
	MaxBondedTarget                     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=max_bonded_target,json=maxBondedTarget,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_bonded_target" yaml:"max_bonded_target"`
	LowFactor                           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=low_factor,json=lowFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"low_factor" yaml:"low_factor"`
	LeftoverBurnRate                    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=leftover_burn_rate,json=leftoverBurnRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"leftover_burn_rate" yaml:"leftover_burn_rate"`
	MaxRewardBoost                      uint64                                 `protobuf:"varint,5,opt,name=max_reward_boost,json=maxRewardBoost,proto3" json:"max_reward_boost,omitempty"`
	ValidatorsSubscriptionParticipation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=validators_subscription_participation,json=validatorsSubscriptionParticipation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"validators_subscription_participation" yaml:"validators_subscription_participation"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_12687c5fbcde5c39, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxRewardBoost() uint64 {
	if m != nil {
		return m.MaxRewardBoost
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "lavanet.lava.rewards.Params")
}

func init() { proto.RegisterFile("lavanet/lava/rewards/params.proto", fileDescriptor_12687c5fbcde5c39) }

var fileDescriptor_12687c5fbcde5c39 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xbf, 0x8e, 0x94, 0x50,
	0x14, 0xc6, 0x41, 0x67, 0x27, 0xd9, 0x5b, 0xe8, 0x2e, 0xd9, 0x02, 0x2d, 0x60, 0xc5, 0x68, 0x26,
	0x46, 0x21, 0xd1, 0x6e, 0x4b, 0xdc, 0x58, 0x68, 0xb3, 0x41, 0x2b, 0x0b, 0x6f, 0x0e, 0x70, 0x17,
	0x6f, 0x04, 0x0e, 0xb9, 0xf7, 0xc2, 0x30, 0x6f, 0x61, 0x69, 0xe9, 0x33, 0x58, 0xfb, 0x00, 0x5b,
	0x6e, 0x69, 0x2c, 0x26, 0x66, 0xe6, 0x0d, 0x7c, 0x02, 0xc3, 0x05, 0x9d, 0x7f, 0x8d, 0x93, 0xad,
	0xce, 0xe1, 0xcb, 0xc7, 0xf7, 0xfb, 0x02, 0x39, 0xe4, 0x41, 0x0e, 0x0d, 0x94, 0x4c, 0x05, 0xdd,
	0x0c, 0x04, 0x9b, 0x82, 0x48, 0x65, 0x50, 0x81, 0x80, 0x42, 0xfa, 0x95, 0x40, 0x85, 0xd6, 0xc9,
	0x60, 0xf1, 0xbb, 0xe9, 0x0f, 0x96, 0xfb, 0x27, 0x19, 0x66, 0xa8, 0x0d, 0x41, 0xb7, 0xf5, 0x5e,
	0xef, 0xfb, 0x01, 0x19, 0x5f, 0xe8, 0x97, 0xad, 0x86, 0x1c, 0x17, 0xbc, 0xa4, 0x31, 0x96, 0x29,
	0x4b, 0xa9, 0x02, 0x91, 0x31, 0x65, 0x9b, 0xa7, 0xe6, 0xe4, 0x30, 0x7c, 0x7d, 0x35, 0x77, 0x8d,
	0x9f, 0x73, 0xf7, 0x71, 0xc6, 0xd5, 0xc7, 0x3a, 0xf6, 0x13, 0x2c, 0x82, 0x04, 0x65, 0x81, 0x72,
	0x18, 0xcf, 0x64, 0xfa, 0x29, 0x50, 0xb3, 0x8a, 0x49, 0xff, 0x9c, 0x25, 0xbf, 0xe7, 0xae, 0x3d,
	0x83, 0x22, 0x3f, 0xf3, 0x76, 0x02, 0xbd, 0xe8, 0x6e, 0xc1, 0xcb, 0x50, 0x4b, 0xef, 0xb4, 0xa2,
	0xb9, 0xd0, 0x6e, 0x71, 0x6f, 0xdd, 0x90, 0xbb, 0x1d, 0xd8, 0x71, 0xa1, 0xdd, 0xe0, 0xc6, 0x84,
	0xe4, 0x38, 0xa5, 0x97, 0x90, 0x28, 0x14, 0xf6, 0x6d, 0x0d, 0x7c, 0xb9, 0x37, 0xf0, 0xb8, 0x07,
	0xae, 0x92, 0xbc, 0xe8, 0x30, 0xc7, 0xe9, 0x2b, 0xbd, 0x5b, 0x33, 0x62, 0xe5, 0xec, 0x52, 0x61,
	0xc3, 0x04, 0x8d, 0x6b, 0x51, 0x52, 0x01, 0x8a, 0xd9, 0x23, 0xcd, 0x7a, 0xb3, 0x37, 0xeb, 0xde,
	0xc0, 0xda, 0x49, 0xf4, 0xa2, 0xa3, 0xbf, 0x62, 0x58, 0x8b, 0x32, 0x02, 0xc5, 0xac, 0x09, 0x39,
	0xea, 0xbe, 0x42, 0xff, 0xfb, 0x69, 0x8c, 0x28, 0x95, 0x7d, 0x70, 0x6a, 0x4e, 0x46, 0xd1, 0x9d,
	0x02, 0xda, 0x48, 0xcb, 0x61, 0xa7, 0x5a, 0xdf, 0x4c, 0xf2, 0xa8, 0x81, 0x9c, 0xa7, 0xa0, 0x50,
	0x48, 0x2a, 0xeb, 0x58, 0x26, 0x82, 0x57, 0x8a, 0x63, 0x49, 0x2b, 0x10, 0x8a, 0x27, 0xbc, 0x82,
	0xee, 0xc9, 0x1e, 0xeb, 0xe2, 0x1f, 0xf6, 0x2e, 0xfe, 0xb4, 0x2f, 0xfe, 0x5f, 0x10, 0x2f, 0x7a,
	0xb8, 0xf2, 0xbd, 0x5d, 0xb3, 0x5d, 0xac, 0xbb, 0xce, 0x46, 0x5f, 0xbe, 0xba, 0x46, 0x78, 0x7e,
	0xb5, 0x70, 0xcc, 0xeb, 0x85, 0x63, 0xfe, 0x5a, 0x38, 0xe6, 0xe7, 0xa5, 0x63, 0x5c, 0x2f, 0x1d,
	0xe3, 0xc7, 0xd2, 0x31, 0xde, 0x3f, 0x59, 0x2b, 0xb7, 0x71, 0x32, 0xcd, 0xf3, 0xa0, 0xfd, 0x77,
	0x37, 0xba, 0x64, 0x3c, 0xd6, 0xb7, 0xf0, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x9f,
	0x83, 0x73, 0x5c, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ValidatorsSubscriptionParticipation.Size()
		i -= size
		if _, err := m.ValidatorsSubscriptionParticipation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.MaxRewardBoost != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxRewardBoost))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.LeftoverBurnRate.Size()
		i -= size
		if _, err := m.LeftoverBurnRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.LowFactor.Size()
		i -= size
		if _, err := m.LowFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MaxBondedTarget.Size()
		i -= size
		if _, err := m.MaxBondedTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinBondedTarget.Size()
		i -= size
		if _, err := m.MinBondedTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinBondedTarget.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxBondedTarget.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.LowFactor.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.LeftoverBurnRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MaxRewardBoost != 0 {
		n += 1 + sovParams(uint64(m.MaxRewardBoost))
	}
	l = m.ValidatorsSubscriptionParticipation.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBondedTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBondedTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBondedTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxBondedTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LowFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeftoverBurnRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LeftoverBurnRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRewardBoost", wireType)
			}
			m.MaxRewardBoost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRewardBoost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorsSubscriptionParticipation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValidatorsSubscriptionParticipation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
