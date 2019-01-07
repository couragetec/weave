// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/nft/username/codec.proto

/*
	Package username is a generated protocol buffer package.

	It is generated from these files:
		x/nft/username/codec.proto

	It has these top-level messages:
		UsernameToken
		TokenDetails
		ChainAddress
		IssueTokenMsg
		AddChainAddressMsg
		RemoveChainAddressMsg
*/
package username

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import nft "github.com/iov-one/weave/x/nft"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UsernameToken struct {
	Base    *nft.NonFungibleToken `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	Details *TokenDetails         `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
}

func (m *UsernameToken) Reset()                    { *m = UsernameToken{} }
func (m *UsernameToken) String() string            { return proto.CompactTextString(m) }
func (*UsernameToken) ProtoMessage()               {}
func (*UsernameToken) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{0} }

func (m *UsernameToken) GetBase() *nft.NonFungibleToken {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *UsernameToken) GetDetails() *TokenDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

type TokenDetails struct {
	Addresses []ChainAddress `protobuf:"bytes,1,rep,name=addresses" json:"addresses"`
}

func (m *TokenDetails) Reset()                    { *m = TokenDetails{} }
func (m *TokenDetails) String() string            { return proto.CompactTextString(m) }
func (*TokenDetails) ProtoMessage()               {}
func (*TokenDetails) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{1} }

func (m *TokenDetails) GetAddresses() []ChainAddress {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// ChainAddress is an address on a blockchain
type ChainAddress struct {
	// blockchainID is the reference to a blockchain nft token.
	BlockchainID []byte `protobuf:"bytes,1,opt,name=blockchain_id,json=blockchainId,proto3" json:"blockchain_id,omitempty"`
	// address is the unique identifier of an account on the referenced blockchain
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ChainAddress) Reset()                    { *m = ChainAddress{} }
func (m *ChainAddress) String() string            { return proto.CompactTextString(m) }
func (*ChainAddress) ProtoMessage()               {}
func (*ChainAddress) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{2} }

func (m *ChainAddress) GetBlockchainID() []byte {
	if m != nil {
		return m.BlockchainID
	}
	return nil
}

func (m *ChainAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type IssueTokenMsg struct {
	ID        []byte                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner     []byte                `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Approvals []nft.ActionApprovals `protobuf:"bytes,3,rep,name=approvals" json:"approvals"`
	Details   TokenDetails          `protobuf:"bytes,4,opt,name=details" json:"details"`
}

func (m *IssueTokenMsg) Reset()                    { *m = IssueTokenMsg{} }
func (m *IssueTokenMsg) String() string            { return proto.CompactTextString(m) }
func (*IssueTokenMsg) ProtoMessage()               {}
func (*IssueTokenMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{3} }

func (m *IssueTokenMsg) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *IssueTokenMsg) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *IssueTokenMsg) GetApprovals() []nft.ActionApprovals {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func (m *IssueTokenMsg) GetDetails() TokenDetails {
	if m != nil {
		return m.Details
	}
	return TokenDetails{}
}

type AddChainAddressMsg struct {
	UsernameID   []byte `protobuf:"bytes,1,opt,name=username_id,json=usernameId,proto3" json:"username_id,omitempty"`
	BlockchainID []byte `protobuf:"bytes,2,opt,name=blockchain_id,json=blockchainId,proto3" json:"blockchain_id,omitempty"`
	Address      string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *AddChainAddressMsg) Reset()                    { *m = AddChainAddressMsg{} }
func (m *AddChainAddressMsg) String() string            { return proto.CompactTextString(m) }
func (*AddChainAddressMsg) ProtoMessage()               {}
func (*AddChainAddressMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{4} }

func (m *AddChainAddressMsg) GetUsernameID() []byte {
	if m != nil {
		return m.UsernameID
	}
	return nil
}

func (m *AddChainAddressMsg) GetBlockchainID() []byte {
	if m != nil {
		return m.BlockchainID
	}
	return nil
}

func (m *AddChainAddressMsg) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type RemoveChainAddressMsg struct {
	UsernameID   []byte `protobuf:"bytes,1,opt,name=username_id,json=usernameId,proto3" json:"username_id,omitempty"`
	BlockchainID []byte `protobuf:"bytes,2,opt,name=blockchain_id,json=blockchainId,proto3" json:"blockchain_id,omitempty"`
	Address      string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *RemoveChainAddressMsg) Reset()                    { *m = RemoveChainAddressMsg{} }
func (m *RemoveChainAddressMsg) String() string            { return proto.CompactTextString(m) }
func (*RemoveChainAddressMsg) ProtoMessage()               {}
func (*RemoveChainAddressMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{5} }

func (m *RemoveChainAddressMsg) GetUsernameID() []byte {
	if m != nil {
		return m.UsernameID
	}
	return nil
}

func (m *RemoveChainAddressMsg) GetBlockchainID() []byte {
	if m != nil {
		return m.BlockchainID
	}
	return nil
}

func (m *RemoveChainAddressMsg) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*UsernameToken)(nil), "username.UsernameToken")
	proto.RegisterType((*TokenDetails)(nil), "username.TokenDetails")
	proto.RegisterType((*ChainAddress)(nil), "username.ChainAddress")
	proto.RegisterType((*IssueTokenMsg)(nil), "username.IssueTokenMsg")
	proto.RegisterType((*AddChainAddressMsg)(nil), "username.AddChainAddressMsg")
	proto.RegisterType((*RemoveChainAddressMsg)(nil), "username.RemoveChainAddressMsg")
}
func (m *UsernameToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UsernameToken) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Base != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Base.Size()))
		n1, err := m.Base.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Details != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Details.Size()))
		n2, err := m.Details.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *TokenDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, msg := range m.Addresses {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ChainAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainAddress) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.BlockchainID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.BlockchainID)))
		i += copy(dAtA[i:], m.BlockchainID)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	return i, nil
}

func (m *IssueTokenMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IssueTokenMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Owner) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if len(m.Approvals) > 0 {
		for _, msg := range m.Approvals {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintCodec(dAtA, i, uint64(m.Details.Size()))
	n3, err := m.Details.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *AddChainAddressMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddChainAddressMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UsernameID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.UsernameID)))
		i += copy(dAtA[i:], m.UsernameID)
	}
	if len(m.BlockchainID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.BlockchainID)))
		i += copy(dAtA[i:], m.BlockchainID)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	return i, nil
}

func (m *RemoveChainAddressMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveChainAddressMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UsernameID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.UsernameID)))
		i += copy(dAtA[i:], m.UsernameID)
	}
	if len(m.BlockchainID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.BlockchainID)))
		i += copy(dAtA[i:], m.BlockchainID)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UsernameToken) Size() (n int) {
	var l int
	_ = l
	if m.Base != nil {
		l = m.Base.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Details != nil {
		l = m.Details.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *TokenDetails) Size() (n int) {
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, e := range m.Addresses {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *ChainAddress) Size() (n int) {
	var l int
	_ = l
	l = len(m.BlockchainID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *IssueTokenMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	l = m.Details.Size()
	n += 1 + l + sovCodec(uint64(l))
	return n
}

func (m *AddChainAddressMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.UsernameID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.BlockchainID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *RemoveChainAddressMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.UsernameID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.BlockchainID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UsernameToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UsernameToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UsernameToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Base == nil {
				m.Base = &nft.NonFungibleToken{}
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Details == nil {
				m.Details = &TokenDetails{}
			}
			if err := m.Details.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *TokenDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, ChainAddress{})
			if err := m.Addresses[len(m.Addresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *ChainAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockchainID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockchainID = append(m.BlockchainID[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockchainID == nil {
				m.BlockchainID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *IssueTokenMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IssueTokenMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueTokenMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, nft.ActionApprovals{})
			if err := m.Approvals[len(m.Approvals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Details.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *AddChainAddressMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddChainAddressMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddChainAddressMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsernameID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UsernameID = append(m.UsernameID[:0], dAtA[iNdEx:postIndex]...)
			if m.UsernameID == nil {
				m.UsernameID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockchainID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockchainID = append(m.BlockchainID[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockchainID == nil {
				m.BlockchainID = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *RemoveChainAddressMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoveChainAddressMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveChainAddressMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsernameID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UsernameID = append(m.UsernameID[:0], dAtA[iNdEx:postIndex]...)
			if m.UsernameID == nil {
				m.UsernameID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockchainID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockchainID = append(m.BlockchainID[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockchainID == nil {
				m.BlockchainID = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("x/nft/username/codec.proto", fileDescriptorCodec) }

var fileDescriptorCodec = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x3b, 0xd9, 0xb5, 0x75, 0x5f, 0xb3, 0x52, 0x86, 0xb6, 0x84, 0x3d, 0xec, 0x96, 0x3d,
	0xa9, 0xd0, 0x8c, 0x28, 0x8a, 0x78, 0xdb, 0x75, 0x11, 0x22, 0xe8, 0x61, 0xd0, 0x73, 0x99, 0x64,
	0x66, 0xd3, 0xa1, 0xd9, 0x99, 0x25, 0x93, 0x6c, 0xfd, 0x18, 0x9e, 0x04, 0x3f, 0x85, 0x5f, 0xa3,
	0x47, 0x3f, 0xc1, 0x22, 0xf1, 0x8b, 0x48, 0x27, 0x3b, 0x49, 0x8a, 0x20, 0x78, 0xeb, 0x2d, 0xef,
	0xfd, 0x7f, 0xef, 0xe5, 0xfd, 0xff, 0x21, 0x30, 0xfa, 0x42, 0xd4, 0xb2, 0x20, 0xa5, 0x11, 0xb9,
	0x62, 0x2b, 0x41, 0x12, 0xcd, 0x45, 0x12, 0xae, 0x73, 0x5d, 0x68, 0xfc, 0xd0, 0x75, 0x47, 0x4f,
	0x53, 0x59, 0x5c, 0x96, 0x71, 0x98, 0xe8, 0x15, 0x91, 0x7a, 0x73, 0xae, 0x95, 0x20, 0xd7, 0x82,
	0x6d, 0x04, 0xa9, 0xc7, 0x3b, 0x53, 0xa3, 0xf3, 0x0e, 0x9b, 0xea, 0x54, 0x13, 0xdb, 0x8e, 0xcb,
	0xa5, 0xad, 0x6c, 0x61, 0x9f, 0x6a, 0x7c, 0x9a, 0xc1, 0xf0, 0xf3, 0xee, 0x35, 0x9f, 0xf4, 0x95,
	0x50, 0xf8, 0x09, 0xf4, 0x63, 0x66, 0x44, 0x80, 0xce, 0xd0, 0xe3, 0xc3, 0xe7, 0x27, 0xa1, 0x5a,
	0x16, 0xe1, 0x47, 0xad, 0xde, 0x95, 0x2a, 0x95, 0x71, 0x56, 0x43, 0xd4, 0x22, 0xf8, 0x19, 0x1c,
	0x70, 0x51, 0x30, 0x99, 0x99, 0xc0, 0xb3, 0xf4, 0x69, 0xe8, 0x4e, 0x0e, 0x2d, 0xb7, 0xa8, 0x55,
	0xea, 0xb0, 0xe9, 0x7b, 0xf0, 0xbb, 0x02, 0x7e, 0x03, 0x03, 0xc6, 0x79, 0x2e, 0x8c, 0x11, 0x26,
	0x40, 0x67, 0xbd, 0xbb, 0x3b, 0xde, 0x5e, 0x32, 0xa9, 0x66, 0xb5, 0x3e, 0xef, 0xdf, 0x6c, 0x27,
	0x7b, 0xb4, 0xc5, 0xa7, 0x17, 0xe0, 0x77, 0x01, 0xfc, 0x12, 0x86, 0x71, 0xa6, 0x93, 0xab, 0xe4,
	0xb6, 0x79, 0x21, 0xb9, 0x75, 0xe0, 0xcf, 0x8f, 0xaa, 0xed, 0xc4, 0x9f, 0x37, 0x42, 0xb4, 0xa0,
	0x7e, 0x8b, 0x45, 0x1c, 0x07, 0x70, 0xb0, 0xdb, 0x69, 0x4d, 0x0c, 0xa8, 0x2b, 0xa7, 0x3f, 0x10,
	0x0c, 0x23, 0x63, 0xca, 0xda, 0xf3, 0x07, 0x93, 0xe2, 0x53, 0xf0, 0x9a, 0xbd, 0xfb, 0xd5, 0x76,
	0xe2, 0x45, 0x0b, 0xea, 0x49, 0x8e, 0x8f, 0xe1, 0x81, 0xbe, 0x56, 0x22, 0xb7, 0x1b, 0x7c, 0x5a,
	0x17, 0xf8, 0x35, 0x0c, 0xd8, 0x7a, 0x9d, 0xeb, 0x0d, 0xcb, 0x4c, 0xd0, 0xb3, 0xe6, 0x8e, 0x6d,
	0x9c, 0xb3, 0xa4, 0x90, 0x5a, 0xcd, 0x9c, 0xd6, 0x58, 0x73, 0x0d, 0xfc, 0xaa, 0x0d, 0xb6, 0xff,
	0xaf, 0x60, 0x77, 0x93, 0x4d, 0xbc, 0xdf, 0x10, 0xe0, 0x19, 0xe7, 0xdd, 0x58, 0x6e, 0xcf, 0x26,
	0x70, 0xe8, 0xc6, 0xdb, 0x5c, 0x1e, 0x55, 0xdb, 0x09, 0xb8, 0x4f, 0x1f, 0x2d, 0x28, 0x38, 0x24,
	0xe2, 0x7f, 0x47, 0xe9, 0xfd, 0x6f, 0x94, 0xbd, 0xbb, 0x51, 0x7e, 0x47, 0x70, 0x42, 0xc5, 0x4a,
	0x6f, 0xc4, 0xbd, 0xbb, 0x6d, 0x7e, 0x74, 0x53, 0x8d, 0xd1, 0xcf, 0x6a, 0x8c, 0x7e, 0x55, 0x63,
	0xf4, 0xf5, 0xf7, 0x78, 0x2f, 0xde, 0xb7, 0xbf, 0xc6, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x87, 0xec, 0xc1, 0xdc, 0x9d, 0x03, 0x00, 0x00,
}
