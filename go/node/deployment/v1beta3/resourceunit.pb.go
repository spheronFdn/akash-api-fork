// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta3/resourceunit.proto

package v1beta3

import (
	fmt "fmt"
	v1beta3 "github.com/akash-network/akash-api/go/node/types/v1beta3"
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

// ResourceUnit extends Resources and adds Count along with the Price
type ResourceUnit struct {
	v1beta3.Resources `protobuf:"bytes,1,opt,name=resource,proto3,embedded=resource" json:"resource" yaml:"resource"`
	Count             uint32        `protobuf:"varint,2,opt,name=count,proto3" json:"count" yaml:"count"`
	Price             types.DecCoin `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *ResourceUnit) Reset()         { *m = ResourceUnit{} }
func (m *ResourceUnit) String() string { return proto.CompactTextString(m) }
func (*ResourceUnit) ProtoMessage()    {}
func (*ResourceUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb431767d5aa2e0f, []int{0}
}
func (m *ResourceUnit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceUnit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceUnit.Merge(m, src)
}
func (m *ResourceUnit) XXX_Size() int {
	return m.Size()
}
func (m *ResourceUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceUnit.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceUnit proto.InternalMessageInfo

func (m *ResourceUnit) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ResourceUnit) GetPrice() types.DecCoin {
	if m != nil {
		return m.Price
	}
	return types.DecCoin{}
}

func init() {
	proto.RegisterType((*ResourceUnit)(nil), "akash.deployment.v1beta3.ResourceUnit")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta3/resourceunit.proto", fileDescriptor_fb431767d5aa2e0f)
}

var fileDescriptor_fb431767d5aa2e0f = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x33, 0xf7, 0xde, 0x5e, 0x2e, 0xb9, 0x15, 0x21, 0xb8, 0x88, 0xc5, 0x26, 0x25, 0x1b,
	0x0b, 0xe2, 0x0c, 0xb5, 0xbb, 0x82, 0x9b, 0xe8, 0x0b, 0x18, 0x10, 0xc1, 0x5d, 0x92, 0x0e, 0xe9,
	0xd0, 0x66, 0x4e, 0xc8, 0x4c, 0x94, 0xbe, 0x85, 0x8f, 0xe0, 0xe3, 0x74, 0xd9, 0xa5, 0xab, 0x20,
	0xed, 0x46, 0xba, 0xec, 0xda, 0x85, 0x64, 0x26, 0x69, 0x0b, 0xba, 0xcb, 0xf9, 0xcf, 0xf7, 0xff,
	0xfc, 0x27, 0x63, 0x5e, 0x84, 0xd3, 0x50, 0x4c, 0xc8, 0x98, 0x66, 0x33, 0x98, 0xa7, 0x94, 0x4b,
	0xf2, 0x34, 0x88, 0xa8, 0x0c, 0x87, 0x24, 0xa7, 0x02, 0x8a, 0x3c, 0xa6, 0x05, 0x67, 0x12, 0x67,
	0x39, 0x48, 0xb0, 0x6c, 0x05, 0xe3, 0x3d, 0x8c, 0x6b, 0xb8, 0x73, 0x92, 0x40, 0x02, 0x0a, 0x22,
	0xd5, 0x97, 0xe6, 0x3b, 0x9e, 0x0e, 0x8f, 0x42, 0x41, 0xbf, 0xc5, 0x8a, 0x9a, 0x71, 0x62, 0x10,
	0x29, 0x88, 0x43, 0x68, 0x40, 0x62, 0x60, 0x5c, 0xef, 0xbd, 0x4f, 0x64, 0xb6, 0x83, 0xda, 0x73,
	0xcf, 0x99, 0xb4, 0x22, 0xf3, 0x5f, 0x93, 0x61, 0xa3, 0x1e, 0xea, 0xff, 0xbf, 0xea, 0x62, 0xdd,
	0xab, 0x8a, 0x68, 0x1a, 0xe1, 0xc6, 0x23, 0xfc, 0xf3, 0x45, 0xe9, 0x1a, 0xcb, 0xd2, 0x45, 0x9b,
	0xd2, 0xdd, 0x59, 0xb7, 0xa5, 0x7b, 0x3c, 0x0f, 0xd3, 0xd9, 0xc8, 0x6b, 0x14, 0x2f, 0xd8, 0x2d,
	0x2d, 0x62, 0xb6, 0x62, 0x28, 0xb8, 0xb4, 0x7f, 0xf5, 0x50, 0xff, 0xc8, 0x3f, 0xdd, 0x94, 0xae,
	0x16, 0xb6, 0xa5, 0xdb, 0xd6, 0x36, 0x35, 0x7a, 0x81, 0x96, 0xad, 0x3b, 0xb3, 0x95, 0xe5, 0x2c,
	0xa6, 0xf6, 0x6f, 0xd5, 0xe8, 0x0c, 0xeb, 0xab, 0x0e, 0x2b, 0x0d, 0xf0, 0x2d, 0x8d, 0x6f, 0x80,
	0x71, 0xbf, 0x5b, 0x15, 0xaa, 0x22, 0x95, 0x65, 0x1f, 0xa9, 0x46, 0x2f, 0xd0, 0xf2, 0xe8, 0xcf,
	0xc7, 0xab, 0x6b, 0xf8, 0x0f, 0x8b, 0x95, 0x83, 0x96, 0x2b, 0x07, 0xbd, 0xaf, 0x1c, 0xf4, 0xb2,
	0x76, 0x8c, 0xe5, 0xda, 0x31, 0xde, 0xd6, 0x8e, 0xf1, 0x78, 0x9d, 0x30, 0x39, 0x29, 0x22, 0x1c,
	0x43, 0x4a, 0xd4, 0xfd, 0x97, 0x9c, 0xca, 0x67, 0xc8, 0xa7, 0xf5, 0x14, 0x66, 0x8c, 0x24, 0x40,
	0x38, 0x8c, 0xe9, 0x0f, 0xcf, 0x1b, 0xfd, 0x55, 0xbf, 0x77, 0xf8, 0x15, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0x9f, 0xec, 0xd3, 0x01, 0x02, 0x00, 0x00,
}

func (m *ResourceUnit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceUnit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceUnit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResourceunit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Count != 0 {
		i = encodeVarintResourceunit(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Resources.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResourceunit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintResourceunit(dAtA []byte, offset int, v uint64) int {
	offset -= sovResourceunit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResourceUnit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Resources.Size()
	n += 1 + l + sovResourceunit(uint64(l))
	if m.Count != 0 {
		n += 1 + sovResourceunit(uint64(m.Count))
	}
	l = m.Price.Size()
	n += 1 + l + sovResourceunit(uint64(l))
	return n
}

func sovResourceunit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResourceunit(x uint64) (n int) {
	return sovResourceunit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceUnit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceunit
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
			return fmt.Errorf("proto: ResourceUnit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceUnit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunit
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
				return ErrInvalidLengthResourceunit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resources.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunit
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
				return ErrInvalidLengthResourceunit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourceunit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResourceunit
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
func skipResourceunit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourceunit
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
					return 0, ErrIntOverflowResourceunit
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
					return 0, ErrIntOverflowResourceunit
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
				return 0, ErrInvalidLengthResourceunit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResourceunit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResourceunit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResourceunit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourceunit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResourceunit = fmt.Errorf("proto: unexpected end of group")
)