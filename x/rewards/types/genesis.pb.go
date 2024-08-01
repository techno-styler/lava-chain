// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/rewards/genesis.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/v2/x/timerstore/types"
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

// GenesisState defines the rewards module's genesis state.
type GenesisState struct {
	Params              Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	RefillRewardsTS     types.GenesisState `protobuf:"bytes,2,opt,name=refillRewardsTS,proto3" json:"refillRewardsTS"`
	BasePays            []BasePayGenesis   `protobuf:"bytes,3,rep,name=base_pays,json=basePays,proto3" json:"base_pays"`
	IprpcSubscriptions  []string           `protobuf:"bytes,4,rep,name=iprpc_subscriptions,json=iprpcSubscriptions,proto3" json:"iprpc_subscriptions,omitempty"`
	MinIprpcCost        types1.Coin        `protobuf:"bytes,5,opt,name=min_iprpc_cost,json=minIprpcCost,proto3" json:"min_iprpc_cost"`
	IprpcRewards        []IprpcReward      `protobuf:"bytes,6,rep,name=iprpc_rewards,json=iprpcRewards,proto3" json:"iprpc_rewards"`
	IprpcRewardsCurrent uint64             `protobuf:"varint,7,opt,name=iprpc_rewards_current,json=iprpcRewardsCurrent,proto3" json:"iprpc_rewards_current,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_02c24f4df31ca14e, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRefillRewardsTS() types.GenesisState {
	if m != nil {
		return m.RefillRewardsTS
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetBasePays() []BasePayGenesis {
	if m != nil {
		return m.BasePays
	}
	return nil
}

func (m *GenesisState) GetIprpcSubscriptions() []string {
	if m != nil {
		return m.IprpcSubscriptions
	}
	return nil
}

func (m *GenesisState) GetMinIprpcCost() types1.Coin {
	if m != nil {
		return m.MinIprpcCost
	}
	return types1.Coin{}
}

func (m *GenesisState) GetIprpcRewards() []IprpcReward {
	if m != nil {
		return m.IprpcRewards
	}
	return nil
}

func (m *GenesisState) GetIprpcRewardsCurrent() uint64 {
	if m != nil {
		return m.IprpcRewardsCurrent
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.rewards.GenesisState")
}

func init() {
	proto.RegisterFile("lavanet/lava/rewards/genesis.proto", fileDescriptor_02c24f4df31ca14e)
}

var fileDescriptor_02c24f4df31ca14e = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x80, 0x1b, 0x52, 0x0a, 0xf3, 0x06, 0x48, 0xde, 0x90, 0x42, 0x85, 0x42, 0x36, 0x40, 0x8a,
	0x38, 0xd8, 0x5a, 0xb8, 0x71, 0x6c, 0x41, 0x13, 0x12, 0x87, 0x29, 0x85, 0x0b, 0x97, 0xc8, 0x09,
	0x26, 0x58, 0x6a, 0xec, 0xc8, 0xcf, 0x2d, 0xf4, 0x5f, 0xf0, 0xb3, 0x76, 0xdc, 0x91, 0x13, 0x42,
	0x2d, 0x3f, 0x04, 0xc5, 0xf6, 0x68, 0x3b, 0xe5, 0x64, 0xcb, 0xef, 0x7b, 0x9f, 0xdf, 0x7b, 0x7a,
	0xe8, 0x6c, 0xce, 0x96, 0x4c, 0x72, 0x43, 0xbb, 0x93, 0x6a, 0xfe, 0x9d, 0xe9, 0x2f, 0x40, 0x6b,
	0x2e, 0x39, 0x08, 0x20, 0xad, 0x56, 0x46, 0xe1, 0x13, 0xcf, 0x90, 0xee, 0x24, 0x9e, 0x19, 0x9f,
	0xd4, 0xaa, 0x56, 0x16, 0xa0, 0xdd, 0xcd, 0xb1, 0xe3, 0xd3, 0x5e, 0x5f, 0xcb, 0x34, 0x6b, 0xbc,
	0x6e, 0xfc, 0xbc, 0x17, 0x29, 0x19, 0xf0, 0xa2, 0x65, 0x2b, 0x0f, 0x25, 0xbd, 0x90, 0x68, 0x75,
	0x5b, 0xf5, 0x6a, 0x8c, 0x68, 0xb8, 0x06, 0xa3, 0x34, 0x77, 0x57, 0x0f, 0xc5, 0x95, 0x82, 0x46,
	0x39, 0x3b, 0x5d, 0x9e, 0x97, 0xdc, 0xb0, 0x73, 0x5a, 0x29, 0x21, 0x5d, 0xfc, 0xec, 0x6f, 0x88,
	0x8e, 0x2e, 0x5c, 0xb3, 0x33, 0xc3, 0x0c, 0xc7, 0x6f, 0xd0, 0xc8, 0x15, 0x1b, 0x05, 0x49, 0x90,
	0x1e, 0x66, 0x4f, 0x49, 0x5f, 0xf3, 0xe4, 0xd2, 0x32, 0x93, 0xe1, 0xd5, 0xef, 0x67, 0x83, 0xdc,
	0x67, 0xe0, 0x4f, 0xe8, 0x91, 0xe6, 0x5f, 0xc5, 0x7c, 0x9e, 0x3b, 0xea, 0xe3, 0x2c, 0xba, 0x63,
	0x25, 0x2f, 0xf7, 0x25, 0xdb, 0x5a, 0xc9, 0xee, 0xdf, 0xde, 0x76, 0xdb, 0x81, 0x2f, 0xd0, 0xc1,
	0xcd, 0x70, 0x20, 0x0a, 0x93, 0x30, 0x3d, 0xcc, 0x5e, 0xf4, 0x57, 0x35, 0x61, 0xc0, 0x2f, 0xd9,
	0xca, 0x4b, 0xbd, 0xef, 0x7e, 0xe9, 0x5e, 0x01, 0x53, 0x74, 0x6c, 0x07, 0x58, 0xc0, 0xa2, 0x84,
	0x4a, 0x8b, 0xd6, 0x08, 0x25, 0x21, 0x1a, 0x26, 0x61, 0x7a, 0x90, 0x63, 0x1b, 0x9a, 0xed, 0x46,
	0xf0, 0x3b, 0xf4, 0xb0, 0x11, 0xb2, 0x70, 0x49, 0x95, 0x02, 0x13, 0xdd, 0xb5, 0xfd, 0x3c, 0x21,
	0x6e, 0xac, 0xa4, 0x53, 0x13, 0x3f, 0x56, 0x32, 0x55, 0x42, 0xfa, 0x3f, 0x8f, 0x1a, 0x21, 0xdf,
	0x77, 0x59, 0x53, 0x05, 0x06, 0x7f, 0x40, 0x0f, 0x9c, 0xc2, 0xd7, 0x19, 0x8d, 0x6c, 0x13, 0xa7,
	0xfd, 0x4d, 0xd8, 0x3c, 0xd7, 0xfd, 0x8d, 0x4d, 0x6c, 0x9f, 0x00, 0x67, 0xe8, 0xf1, 0x9e, 0xad,
	0xa8, 0x16, 0x5a, 0x73, 0x69, 0xa2, 0x7b, 0x49, 0x90, 0x0e, 0xf3, 0xe3, 0x5d, 0x78, 0xea, 0x42,
	0x93, 0xb7, 0x57, 0xeb, 0x38, 0xb8, 0x5e, 0xc7, 0xc1, 0x9f, 0x75, 0x1c, 0xfc, 0xdc, 0xc4, 0x83,
	0xeb, 0x4d, 0x3c, 0xf8, 0xb5, 0x89, 0x07, 0x9f, 0x5f, 0xd5, 0xc2, 0x7c, 0x5b, 0x94, 0xa4, 0x52,
	0x0d, 0xdd, 0x5b, 0xa8, 0x65, 0x46, 0x7f, 0xfc, 0xdf, 0x3b, 0xb3, 0x6a, 0x39, 0x94, 0x23, 0xbb,
	0x33, 0xaf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x44, 0x82, 0x77, 0xee, 0x34, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IprpcRewardsCurrent != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.IprpcRewardsCurrent))
		i--
		dAtA[i] = 0x38
	}
	if len(m.IprpcRewards) > 0 {
		for iNdEx := len(m.IprpcRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IprpcRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size, err := m.MinIprpcCost.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.IprpcSubscriptions) > 0 {
		for iNdEx := len(m.IprpcSubscriptions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.IprpcSubscriptions[iNdEx])
			copy(dAtA[i:], m.IprpcSubscriptions[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.IprpcSubscriptions[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.BasePays) > 0 {
		for iNdEx := len(m.BasePays) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BasePays[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.RefillRewardsTS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.RefillRewardsTS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.BasePays) > 0 {
		for _, e := range m.BasePays {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.IprpcSubscriptions) > 0 {
		for _, s := range m.IprpcSubscriptions {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.MinIprpcCost.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.IprpcRewards) > 0 {
		for _, e := range m.IprpcRewards {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.IprpcRewardsCurrent != 0 {
		n += 1 + sovGenesis(uint64(m.IprpcRewardsCurrent))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefillRewardsTS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RefillRewardsTS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasePays", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BasePays = append(m.BasePays, BasePayGenesis{})
			if err := m.BasePays[len(m.BasePays)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IprpcSubscriptions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IprpcSubscriptions = append(m.IprpcSubscriptions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinIprpcCost", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinIprpcCost.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IprpcRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IprpcRewards = append(m.IprpcRewards, IprpcReward{})
			if err := m.IprpcRewards[len(m.IprpcRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IprpcRewardsCurrent", wireType)
			}
			m.IprpcRewardsCurrent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IprpcRewardsCurrent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
