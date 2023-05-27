// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spec/spec.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Spec_ProvidersTypes int32

const (
	Spec_dynamic Spec_ProvidersTypes = 0
	Spec_static  Spec_ProvidersTypes = 1
)

var Spec_ProvidersTypes_name = map[int32]string{
	0: "dynamic",
	1: "static",
}

var Spec_ProvidersTypes_value = map[string]int32{
	"dynamic": 0,
	"static":  1,
}

func (x Spec_ProvidersTypes) String() string {
	return proto.EnumName(Spec_ProvidersTypes_name, int32(x))
}

func (Spec_ProvidersTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4cc771ffab81d0a, []int{0, 0}
}

type Spec struct {
	Index                         string              `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name                          string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Imports                       []string            `protobuf:"bytes,15,rep,name=imports,proto3" json:"imports,omitempty"`
	Apis                          []ServiceApi        `protobuf:"bytes,3,rep,name=apis,proto3" json:"apis"`
	Enabled                       bool                `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ReliabilityThreshold          uint32              `protobuf:"varint,5,opt,name=reliability_threshold,json=reliabilityThreshold,proto3" json:"reliability_threshold,omitempty"`
	DataReliabilityEnabled        bool                `protobuf:"varint,6,opt,name=data_reliability_enabled,json=dataReliabilityEnabled,proto3" json:"data_reliability_enabled,omitempty"`
	BlockDistanceForFinalizedData uint32              `protobuf:"varint,7,opt,name=block_distance_for_finalized_data,json=blockDistanceForFinalizedData,proto3" json:"block_distance_for_finalized_data,omitempty"`
	BlocksInFinalizationProof     uint32              `protobuf:"varint,8,opt,name=blocks_in_finalization_proof,json=blocksInFinalizationProof,proto3" json:"blocks_in_finalization_proof,omitempty"`
	AverageBlockTime              int64               `protobuf:"varint,9,opt,name=average_block_time,json=averageBlockTime,proto3" json:"average_block_time,omitempty"`
	AllowedBlockLagForQosSync     int64               `protobuf:"varint,10,opt,name=allowed_block_lag_for_qos_sync,json=allowedBlockLagForQosSync,proto3" json:"allowed_block_lag_for_qos_sync,omitempty"`
	BlockLastUpdated              uint64              `protobuf:"varint,11,opt,name=block_last_updated,json=blockLastUpdated,proto3" json:"block_last_updated,omitempty"`
	MinStakeProvider              types.Coin          `protobuf:"bytes,12,opt,name=min_stake_provider,json=minStakeProvider,proto3" json:"min_stake_provider"`
	MinStakeClient                types.Coin          `protobuf:"bytes,13,opt,name=min_stake_client,json=minStakeClient,proto3" json:"min_stake_client"`
	ProvidersTypes                Spec_ProvidersTypes `protobuf:"varint,14,opt,name=providers_types,json=providersTypes,proto3,enum=lavanet.lava.spec.Spec_ProvidersTypes" json:"providers_types,omitempty"`
	ChainId                       string              `protobuf:"bytes,16,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4cc771ffab81d0a, []int{0}
}
func (m *Spec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(m, src)
}
func (m *Spec) XXX_Size() int {
	return m.Size()
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

func (m *Spec) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Spec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Spec) GetImports() []string {
	if m != nil {
		return m.Imports
	}
	return nil
}

func (m *Spec) GetApis() []ServiceApi {
	if m != nil {
		return m.Apis
	}
	return nil
}

func (m *Spec) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Spec) GetReliabilityThreshold() uint32 {
	if m != nil {
		return m.ReliabilityThreshold
	}
	return 0
}

func (m *Spec) GetDataReliabilityEnabled() bool {
	if m != nil {
		return m.DataReliabilityEnabled
	}
	return false
}

func (m *Spec) GetBlockDistanceForFinalizedData() uint32 {
	if m != nil {
		return m.BlockDistanceForFinalizedData
	}
	return 0
}

func (m *Spec) GetBlocksInFinalizationProof() uint32 {
	if m != nil {
		return m.BlocksInFinalizationProof
	}
	return 0
}

func (m *Spec) GetAverageBlockTime() int64 {
	if m != nil {
		return m.AverageBlockTime
	}
	return 0
}

func (m *Spec) GetAllowedBlockLagForQosSync() int64 {
	if m != nil {
		return m.AllowedBlockLagForQosSync
	}
	return 0
}

func (m *Spec) GetBlockLastUpdated() uint64 {
	if m != nil {
		return m.BlockLastUpdated
	}
	return 0
}

func (m *Spec) GetMinStakeProvider() types.Coin {
	if m != nil {
		return m.MinStakeProvider
	}
	return types.Coin{}
}

func (m *Spec) GetMinStakeClient() types.Coin {
	if m != nil {
		return m.MinStakeClient
	}
	return types.Coin{}
}

func (m *Spec) GetProvidersTypes() Spec_ProvidersTypes {
	if m != nil {
		return m.ProvidersTypes
	}
	return Spec_dynamic
}

func (m *Spec) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func init() {
	proto.RegisterEnum("lavanet.lava.spec.Spec_ProvidersTypes", Spec_ProvidersTypes_name, Spec_ProvidersTypes_value)
	proto.RegisterType((*Spec)(nil), "lavanet.lava.spec.Spec")
}

func init() { proto.RegisterFile("spec/spec.proto", fileDescriptor_c4cc771ffab81d0a) }

var fileDescriptor_c4cc771ffab81d0a = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6e, 0xd3, 0x3e,
	0x1c, 0x6f, 0x7e, 0xed, 0xda, 0xcd, 0xfd, 0xad, 0x2b, 0xd6, 0x98, 0xdc, 0x89, 0x85, 0x30, 0x21,
	0x14, 0x24, 0x94, 0x68, 0xdb, 0x01, 0x6e, 0x68, 0xdd, 0xa8, 0x98, 0x04, 0x62, 0xa4, 0xe3, 0xc2,
	0xc5, 0x72, 0x1c, 0xaf, 0xb5, 0x96, 0xd8, 0x21, 0xf6, 0xca, 0xca, 0x53, 0xf0, 0x18, 0x3c, 0xca,
	0x8e, 0x3b, 0x72, 0x42, 0xa8, 0x3b, 0xf0, 0x1a, 0xc8, 0x4e, 0xa2, 0x6d, 0x82, 0x03, 0x17, 0xdb,
	0x5f, 0x7f, 0xfe, 0xf8, 0x63, 0xe5, 0xeb, 0x80, 0x35, 0x95, 0x33, 0x1a, 0x9a, 0x21, 0xc8, 0x0b,
	0xa9, 0x25, 0xbc, 0x97, 0x92, 0x19, 0x11, 0x4c, 0x07, 0x66, 0x0e, 0x0c, 0xb0, 0xb9, 0x3e, 0x91,
	0x13, 0x69, 0xd1, 0xd0, 0xac, 0x4a, 0xe2, 0xe6, 0x46, 0xa9, 0x64, 0xc5, 0x8c, 0x53, 0x86, 0x49,
	0xce, 0xab, 0x7d, 0x97, 0x4a, 0x95, 0x49, 0x15, 0xc6, 0x44, 0xb1, 0x70, 0xb6, 0x13, 0x33, 0x4d,
	0x76, 0x42, 0x2a, 0xb9, 0x28, 0xf1, 0xed, 0x5f, 0x6d, 0xd0, 0x1a, 0xe7, 0x8c, 0xc2, 0x75, 0xb0,
	0xc4, 0x45, 0xc2, 0x2e, 0x90, 0xe3, 0x39, 0xfe, 0x4a, 0x54, 0x16, 0x10, 0x82, 0x96, 0x20, 0x19,
	0x43, 0xff, 0xd9, 0x4d, 0xbb, 0x86, 0x08, 0x74, 0x78, 0x96, 0xcb, 0x42, 0x2b, 0xb4, 0xe6, 0x35,
	0xfd, 0x95, 0xa8, 0x2e, 0xe1, 0x73, 0xd0, 0x22, 0x39, 0x57, 0xa8, 0xe9, 0x35, 0xfd, 0xee, 0xee,
	0x56, 0xf0, 0x47, 0xf8, 0x60, 0x5c, 0x06, 0xdc, 0xcf, 0xf9, 0xb0, 0x75, 0xf9, 0xe3, 0x61, 0x23,
	0xb2, 0x02, 0x63, 0xc9, 0x04, 0x89, 0x53, 0x96, 0xa0, 0x96, 0xe7, 0xf8, 0xcb, 0x51, 0x5d, 0xc2,
	0x3d, 0x70, 0xbf, 0x60, 0x29, 0x27, 0x31, 0x4f, 0xb9, 0x9e, 0x63, 0x3d, 0x2d, 0x98, 0x9a, 0xca,
	0x34, 0x41, 0x4b, 0x9e, 0xe3, 0xaf, 0x46, 0xeb, 0xb7, 0xc0, 0x93, 0x1a, 0x83, 0x2f, 0x00, 0x4a,
	0x88, 0x26, 0xf8, 0xb6, 0xb2, 0xf6, 0x6f, 0x5b, 0xff, 0x0d, 0x83, 0x47, 0x37, 0xf0, 0xab, 0xea,
	0xb8, 0xd7, 0xe0, 0x51, 0x9c, 0x4a, 0x7a, 0x86, 0x13, 0xae, 0x34, 0x11, 0x94, 0xe1, 0x53, 0x59,
	0xe0, 0x53, 0x2e, 0x48, 0xca, 0xbf, 0xb0, 0x04, 0x1b, 0x19, 0xea, 0xd8, 0xa3, 0xb7, 0x2c, 0xf1,
	0xb0, 0xe2, 0x8d, 0x64, 0x31, 0xaa, 0x59, 0x87, 0x44, 0x13, 0xf8, 0x12, 0x3c, 0xb0, 0x04, 0x85,
	0xb9, 0xa8, 0x0d, 0x88, 0xe6, 0x52, 0xe0, 0xbc, 0x90, 0xf2, 0x14, 0x2d, 0x5b, 0x93, 0x41, 0xc9,
	0x39, 0x12, 0xa3, 0x5b, 0x8c, 0x63, 0x43, 0x80, 0xcf, 0x00, 0x24, 0x33, 0x56, 0x90, 0x09, 0xc3,
	0x65, 0x24, 0xcd, 0x33, 0x86, 0x56, 0x3c, 0xc7, 0x6f, 0x46, 0xfd, 0x0a, 0x19, 0x1a, 0xe0, 0x84,
	0x67, 0x0c, 0xee, 0x03, 0x97, 0xa4, 0xa9, 0xfc, 0xcc, 0x92, 0x8a, 0x9d, 0x92, 0x89, 0xcd, 0xfe,
	0x49, 0x2a, 0xac, 0xe6, 0x82, 0x22, 0x60, 0x95, 0x83, 0x8a, 0x65, 0x95, 0x6f, 0xc8, 0x64, 0x24,
	0x8b, 0xf7, 0x52, 0x8d, 0xe7, 0x82, 0x9a, 0x03, 0x6b, 0xa9, 0xd2, 0xf8, 0x3c, 0x4f, 0x88, 0x66,
	0x09, 0xea, 0x7a, 0x8e, 0xdf, 0x8a, 0xfa, 0x71, 0xc9, 0x57, 0xfa, 0x43, 0xb9, 0x0f, 0xdf, 0x02,
	0x98, 0x71, 0x81, 0x95, 0x26, 0x67, 0xcc, 0x5c, 0x69, 0xc6, 0x13, 0x56, 0xa0, 0xff, 0x3d, 0xc7,
	0xef, 0xee, 0x0e, 0x82, 0xb2, 0xeb, 0x02, 0xd3, 0x75, 0x41, 0xd5, 0x75, 0xc1, 0x81, 0xe4, 0xa2,
	0xfa, 0xea, 0xfd, 0x8c, 0x8b, 0xb1, 0x51, 0x1e, 0x57, 0x42, 0x78, 0x04, 0xfa, 0x37, 0x76, 0x34,
	0xe5, 0x4c, 0x68, 0xb4, 0xfa, 0x6f, 0x66, 0xbd, 0xda, 0xec, 0xc0, 0xca, 0xe0, 0x3b, 0xb0, 0x56,
	0xe7, 0x51, 0x58, 0xcf, 0x73, 0xa6, 0x50, 0xcf, 0x73, 0xfc, 0xde, 0xee, 0x93, 0xbf, 0x35, 0xa4,
	0x19, 0xea, 0x14, 0xea, 0xc4, 0xb0, 0xa3, 0x5e, 0x7e, 0xa7, 0x86, 0x03, 0xb0, 0x4c, 0xa7, 0x84,
	0x0b, 0xcc, 0x13, 0xd4, 0xb7, 0x0f, 0xa1, 0x63, 0xeb, 0xa3, 0x64, 0xfb, 0x29, 0xe8, 0xdd, 0x15,
	0xc3, 0x2e, 0xe8, 0x24, 0x73, 0x41, 0x32, 0x4e, 0xfb, 0x0d, 0x08, 0x40, 0x5b, 0x69, 0xa2, 0x39,
	0xed, 0x3b, 0xc3, 0xe1, 0xb7, 0x85, 0xeb, 0x5c, 0x2e, 0x5c, 0xe7, 0x6a, 0xe1, 0x3a, 0x3f, 0x17,
	0xae, 0xf3, 0xf5, 0xda, 0x6d, 0x5c, 0x5d, 0xbb, 0x8d, 0xef, 0xd7, 0x6e, 0xe3, 0xe3, 0xe3, 0x09,
	0xd7, 0xd3, 0xf3, 0x38, 0xa0, 0x32, 0x0b, 0xab, 0x94, 0x76, 0x0e, 0x2f, 0xec, 0xef, 0x20, 0xb4,
	0xf7, 0x88, 0xdb, 0xf6, 0xd1, 0xee, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xca, 0x75, 0xf7,
	0x28, 0x04, 0x00, 0x00,
}

func (this *Spec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Spec)
	if !ok {
		that2, ok := that.(Spec)
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
	if this.Name != that1.Name {
		return false
	}
	if len(this.Imports) != len(that1.Imports) {
		return false
	}
	for i := range this.Imports {
		if this.Imports[i] != that1.Imports[i] {
			return false
		}
	}
	if len(this.Apis) != len(that1.Apis) {
		return false
	}
	for i := range this.Apis {
		if !this.Apis[i].Equal(&that1.Apis[i]) {
			return false
		}
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.ReliabilityThreshold != that1.ReliabilityThreshold {
		return false
	}
	if this.DataReliabilityEnabled != that1.DataReliabilityEnabled {
		return false
	}
	if this.BlockDistanceForFinalizedData != that1.BlockDistanceForFinalizedData {
		return false
	}
	if this.BlocksInFinalizationProof != that1.BlocksInFinalizationProof {
		return false
	}
	if this.AverageBlockTime != that1.AverageBlockTime {
		return false
	}
	if this.AllowedBlockLagForQosSync != that1.AllowedBlockLagForQosSync {
		return false
	}
	if this.BlockLastUpdated != that1.BlockLastUpdated {
		return false
	}
	if !this.MinStakeProvider.Equal(&that1.MinStakeProvider) {
		return false
	}
	if !this.MinStakeClient.Equal(&that1.MinStakeClient) {
		return false
	}
	if this.ProvidersTypes != that1.ProvidersTypes {
		return false
	}
	if this.ChainId != that1.ChainId {
		return false
	}
	return true
}
func (m *Spec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Spec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Spec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintSpec(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.Imports) > 0 {
		for iNdEx := len(m.Imports) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Imports[iNdEx])
			copy(dAtA[i:], m.Imports[iNdEx])
			i = encodeVarintSpec(dAtA, i, uint64(len(m.Imports[iNdEx])))
			i--
			dAtA[i] = 0x7a
		}
	}
	if m.ProvidersTypes != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.ProvidersTypes))
		i--
		dAtA[i] = 0x70
	}
	{
		size, err := m.MinStakeClient.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSpec(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size, err := m.MinStakeProvider.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSpec(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.BlockLastUpdated != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.BlockLastUpdated))
		i--
		dAtA[i] = 0x58
	}
	if m.AllowedBlockLagForQosSync != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.AllowedBlockLagForQosSync))
		i--
		dAtA[i] = 0x50
	}
	if m.AverageBlockTime != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.AverageBlockTime))
		i--
		dAtA[i] = 0x48
	}
	if m.BlocksInFinalizationProof != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.BlocksInFinalizationProof))
		i--
		dAtA[i] = 0x40
	}
	if m.BlockDistanceForFinalizedData != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.BlockDistanceForFinalizedData))
		i--
		dAtA[i] = 0x38
	}
	if m.DataReliabilityEnabled {
		i--
		if m.DataReliabilityEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.ReliabilityThreshold != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.ReliabilityThreshold))
		i--
		dAtA[i] = 0x28
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Apis) > 0 {
		for iNdEx := len(m.Apis) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Apis[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSpec(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSpec(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSpec(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpec(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Spec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSpec(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSpec(uint64(l))
	}
	if len(m.Apis) > 0 {
		for _, e := range m.Apis {
			l = e.Size()
			n += 1 + l + sovSpec(uint64(l))
		}
	}
	if m.Enabled {
		n += 2
	}
	if m.ReliabilityThreshold != 0 {
		n += 1 + sovSpec(uint64(m.ReliabilityThreshold))
	}
	if m.DataReliabilityEnabled {
		n += 2
	}
	if m.BlockDistanceForFinalizedData != 0 {
		n += 1 + sovSpec(uint64(m.BlockDistanceForFinalizedData))
	}
	if m.BlocksInFinalizationProof != 0 {
		n += 1 + sovSpec(uint64(m.BlocksInFinalizationProof))
	}
	if m.AverageBlockTime != 0 {
		n += 1 + sovSpec(uint64(m.AverageBlockTime))
	}
	if m.AllowedBlockLagForQosSync != 0 {
		n += 1 + sovSpec(uint64(m.AllowedBlockLagForQosSync))
	}
	if m.BlockLastUpdated != 0 {
		n += 1 + sovSpec(uint64(m.BlockLastUpdated))
	}
	l = m.MinStakeProvider.Size()
	n += 1 + l + sovSpec(uint64(l))
	l = m.MinStakeClient.Size()
	n += 1 + l + sovSpec(uint64(l))
	if m.ProvidersTypes != 0 {
		n += 1 + sovSpec(uint64(m.ProvidersTypes))
	}
	if len(m.Imports) > 0 {
		for _, s := range m.Imports {
			l = len(s)
			n += 1 + l + sovSpec(uint64(l))
		}
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 2 + l + sovSpec(uint64(l))
	}
	return n
}

func sovSpec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpec(x uint64) (n int) {
	return sovSpec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Spec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpec
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
			return fmt.Errorf("proto: Spec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Spec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Apis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Apis = append(m.Apis, ServiceApi{})
			if err := m.Apis[len(m.Apis)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
			m.Enabled = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReliabilityThreshold", wireType)
			}
			m.ReliabilityThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReliabilityThreshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataReliabilityEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
			m.DataReliabilityEnabled = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockDistanceForFinalizedData", wireType)
			}
			m.BlockDistanceForFinalizedData = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockDistanceForFinalizedData |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksInFinalizationProof", wireType)
			}
			m.BlocksInFinalizationProof = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksInFinalizationProof |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageBlockTime", wireType)
			}
			m.AverageBlockTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AverageBlockTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedBlockLagForQosSync", wireType)
			}
			m.AllowedBlockLagForQosSync = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllowedBlockLagForQosSync |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockLastUpdated", wireType)
			}
			m.BlockLastUpdated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockLastUpdated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStakeProvider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinStakeProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStakeClient", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinStakeClient.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvidersTypes", wireType)
			}
			m.ProvidersTypes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProvidersTypes |= Spec_ProvidersTypes(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Imports", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Imports = append(m.Imports, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpec
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
func skipSpec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpec
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
					return 0, ErrIntOverflowSpec
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
					return 0, ErrIntOverflowSpec
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
				return 0, ErrInvalidLengthSpec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpec = fmt.Errorf("proto: unexpected end of group")
)
