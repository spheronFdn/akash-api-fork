// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta3/params.proto

package v1beta3

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the x/deployment package
type Params struct {
	MinDeposits github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=min_deposits,json=minDeposits,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_deposits" yaml:"min_deposits"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a677e8b392d6c91, []int{0}
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

func (m *Params) GetMinDeposits() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.MinDeposits
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "akash.deployment.v1beta3.Params")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta3/params.proto", fileDescriptor_9a677e8b392d6c91)
}

var fileDescriptor_9a677e8b392d6c91 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x63, 0x5d, 0xa9, 0x43, 0x7b, 0xa7, 0xc2, 0x50, 0x3a, 0x38, 0x28, 0x12, 0x52, 0x97,
	0xd8, 0x0a, 0xdd, 0x58, 0x90, 0x0a, 0x62, 0x43, 0x42, 0x6c, 0xb0, 0x20, 0x27, 0x71, 0x13, 0x2b,
	0xb5, 0x8f, 0x15, 0x1b, 0xa4, 0xbc, 0x05, 0x0f, 0xc1, 0xc4, 0xc2, 0x6b, 0x74, 0xec, 0xc8, 0x14,
	0x50, 0xb2, 0x31, 0xf2, 0x04, 0xa8, 0x4e, 0xa4, 0x16, 0x89, 0xc9, 0x96, 0xfd, 0x9d, 0xef, 0x3f,
	0xfa, 0x87, 0x27, 0xac, 0x60, 0x26, 0xa7, 0x29, 0xd7, 0x2b, 0xa8, 0x24, 0x57, 0x96, 0x3e, 0x45,
	0x31, 0xb7, 0x6c, 0x4e, 0x35, 0x2b, 0x99, 0x34, 0x44, 0x97, 0x60, 0x61, 0x3c, 0x71, 0x18, 0xd9,
	0x61, 0xa4, 0xc7, 0xa6, 0x87, 0x19, 0x64, 0xe0, 0x20, 0xba, 0xbd, 0x75, 0xfc, 0x14, 0x27, 0x60,
	0x24, 0x18, 0x1a, 0x33, 0xc3, 0x7b, 0x63, 0x44, 0x13, 0x10, 0xaa, 0xfb, 0x0f, 0xde, 0xd0, 0x70,
	0x70, 0xe3, 0x02, 0xc6, 0x2f, 0x68, 0xf8, 0x5f, 0x0a, 0xf5, 0x90, 0x72, 0x0d, 0x46, 0x58, 0x33,
	0x41, 0xc7, 0xff, 0x66, 0xa3, 0xd3, 0x23, 0xd2, 0x29, 0xc8, 0x56, 0xd1, 0xa7, 0x45, 0xe4, 0x02,
	0x84, 0x5a, 0x2c, 0xd7, 0xb5, 0xef, 0x35, 0xb5, 0x3f, 0xba, 0x16, 0xea, 0xb2, 0x9f, 0xfa, 0xaa,
	0xfd, 0x5f, 0x96, 0xef, 0xda, 0x3f, 0xa8, 0x98, 0x5c, 0x9d, 0x05, 0xfb, 0xaf, 0xc1, 0xeb, 0x87,
	0x3f, 0xcb, 0x84, 0xcd, 0x1f, 0x63, 0x92, 0x80, 0xa4, 0xfd, 0x96, 0xdd, 0x11, 0x9a, 0xb4, 0xa0,
	0xb6, 0xd2, 0xdc, 0xb8, 0x18, 0x73, 0x3b, 0x92, 0x3b, 0xff, 0xe2, 0x6e, 0xdd, 0x60, 0xb4, 0x69,
	0x30, 0xfa, 0x6c, 0x30, 0x7a, 0x6e, 0xb1, 0xb7, 0x69, 0xb1, 0xf7, 0xde, 0x62, 0xef, 0xfe, 0x7c,
	0x4f, 0x68, 0x74, 0xce, 0x4b, 0x50, 0x57, 0xa9, 0xa2, 0xae, 0xb1, 0x90, 0x69, 0x11, 0x2e, 0xa1,
	0x2c, 0x68, 0x06, 0x54, 0x41, 0xca, 0xff, 0x68, 0x3a, 0x1e, 0xb8, 0x4e, 0xe6, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xa2, 0x3b, 0x48, 0x53, 0x8c, 0x01, 0x00, 0x00,
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
	if len(m.MinDeposits) > 0 {
		for iNdEx := len(m.MinDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
	if len(m.MinDeposits) > 0 {
		for _, e := range m.MinDeposits {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinDeposits = append(m.MinDeposits, types.Coin{})
			if err := m.MinDeposits[len(m.MinDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
