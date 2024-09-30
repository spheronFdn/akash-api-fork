// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta2/resource.proto

package v1beta2

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1beta2 "github.com/spheronFdn/akash-api-fork/go/node/types/v1beta2"
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

// Resource stores unit, total count and price of resource
type Resource struct {
	Resources v1beta2.ResourceUnits `protobuf:"bytes,1,opt,name=resources,proto3" json:"unit" yaml:"unit"`
	Count     uint32                `protobuf:"varint,2,opt,name=count,proto3" json:"count" yaml:"count"`
	Price     types.DecCoin         `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_93085c4a2e404198, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetResources() v1beta2.ResourceUnits {
	if m != nil {
		return m.Resources
	}
	return v1beta2.ResourceUnits{}
}

func (m *Resource) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Resource) GetPrice() types.DecCoin {
	if m != nil {
		return m.Price
	}
	return types.DecCoin{}
}

func init() {
	proto.RegisterType((*Resource)(nil), "akash.deployment.v1beta2.Resource")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta2/resource.proto", fileDescriptor_93085c4a2e404198)
}

var fileDescriptor_93085c4a2e404198 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x4f, 0x83, 0x40,
	0x14, 0xc7, 0x41, 0x5b, 0xa3, 0x54, 0x17, 0xe2, 0x80, 0x55, 0x8f, 0xca, 0xa0, 0x5d, 0x7a, 0x97,
	0xd6, 0xad, 0x8b, 0x09, 0x1a, 0x77, 0x49, 0x1c, 0x34, 0x2e, 0x70, 0x3d, 0x29, 0x69, 0xb9, 0x47,
	0x38, 0x30, 0xe9, 0xb7, 0xf0, 0x23, 0xf8, 0x71, 0x3a, 0x76, 0x74, 0x22, 0xa6, 0x5d, 0x4c, 0xc7,
	0x7e, 0x02, 0xc3, 0x1d, 0x0d, 0x1d, 0xdc, 0x78, 0xff, 0xf7, 0x7b, 0xff, 0xf7, 0xe7, 0x9d, 0x71,
	0xe3, 0x4f, 0x7c, 0x31, 0x26, 0x23, 0x96, 0x4c, 0x61, 0x16, 0x33, 0x9e, 0x91, 0x8f, 0x7e, 0xc0,
	0x32, 0x7f, 0x40, 0x52, 0x26, 0x20, 0x4f, 0x29, 0xc3, 0x49, 0x0a, 0x19, 0x98, 0x96, 0x04, 0x71,
	0x0d, 0xe2, 0x0a, 0x6c, 0x9f, 0x86, 0x10, 0x82, 0x84, 0x48, 0xf9, 0xa5, 0xf8, 0x36, 0xa2, 0x20,
	0x62, 0x10, 0x24, 0xf0, 0x05, 0xab, 0x3c, 0xfb, 0x84, 0x42, 0xc4, 0xab, 0xfe, 0xb5, 0x5a, 0xbc,
	0xd3, 0xae, 0x57, 0xe6, 0x3c, 0xca, 0x84, 0xe2, 0x9c, 0x8d, 0x6e, 0x1c, 0x7a, 0x95, 0x6e, 0xbe,
	0x19, 0x47, 0x5b, 0x46, 0x58, 0x7a, 0x47, 0xef, 0xb6, 0x06, 0x57, 0x58, 0x05, 0x2b, 0x8d, 0xb6,
	0x91, 0xf0, 0x76, 0xe0, 0xb9, 0x34, 0x72, 0xcf, 0xe7, 0x85, 0xad, 0xad, 0x0b, 0xbb, 0x51, 0xfa,
	0x6e, 0x0a, 0xbb, 0x35, 0xf3, 0xe3, 0xe9, 0xd0, 0x29, 0x2b, 0xc7, 0xab, 0x0d, 0x4d, 0x62, 0x34,
	0x29, 0xe4, 0x3c, 0xb3, 0xf6, 0x3a, 0x7a, 0xf7, 0xc4, 0x3d, 0x5b, 0x17, 0xb6, 0x12, 0x36, 0x85,
	0x7d, 0xac, 0x66, 0x64, 0xe9, 0x78, 0x4a, 0x36, 0x9f, 0x8c, 0x66, 0x92, 0x46, 0x94, 0x59, 0xfb,
	0x32, 0xca, 0x05, 0x56, 0xff, 0xbc, 0x9b, 0xa5, 0x8f, 0x1f, 0x18, 0xbd, 0x87, 0x88, 0xbb, 0x97,
	0x55, 0x0a, 0x35, 0x52, 0x5b, 0xca, 0xd2, 0xf1, 0x94, 0x3c, 0x6c, 0xfc, 0x7e, 0xd9, 0x9a, 0xfb,
	0x32, 0x5f, 0x22, 0x7d, 0xb1, 0x44, 0xfa, 0xcf, 0x12, 0xe9, 0x9f, 0x2b, 0xa4, 0x2d, 0x56, 0x48,
	0xfb, 0x5e, 0x21, 0xed, 0xf5, 0x2e, 0x8c, 0xb2, 0x71, 0x1e, 0x60, 0x0a, 0x31, 0x11, 0xc9, 0x98,
	0xa5, 0xc0, 0x1f, 0x47, 0x9c, 0xc8, 0x1b, 0xf4, 0xfc, 0x24, 0xea, 0xbd, 0x43, 0x3a, 0x21, 0x21,
	0x10, 0x0e, 0x23, 0xf6, 0xcf, 0xb3, 0x06, 0x07, 0xf2, 0xac, 0xb7, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x40, 0xbb, 0xc3, 0x68, 0xf9, 0x01, 0x00, 0x00,
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintResource(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Count != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Resources.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResource(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	offset -= sovResource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Resources.Size()
	n += 1 + l + sovResource(uint64(l))
	if m.Count != 0 {
		n += 1 + sovResource(uint64(m.Count))
	}
	l = m.Price.Size()
	n += 1 + l + sovResource(uint64(l))
	return n
}

func sovResource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
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
					return ErrIntOverflowResource
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
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
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
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
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
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
				return 0, ErrInvalidLengthResource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResource = fmt.Errorf("proto: unexpected end of group")
)
