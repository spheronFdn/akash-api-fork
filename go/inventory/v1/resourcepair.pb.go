// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/inventory/v1/resourcepair.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_spheronFdn_akash_api_fork_go_node_types_v1beta3 "github.com/spheronFdn/akash-api-fork/go/node/types/v1beta3"
	v1beta3 "github.com/spheronFdn/akash-api-fork/go/node/types/v1beta3"
	io "io"
	resource "k8s.io/apimachinery/pkg/api/resource"
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

// ResourcePair to extents resource.Quantity to provide total and available units of the resource
type ResourcePair struct {
	Allocatable *resource.Quantity                                                    `protobuf:"bytes,1,opt,name=allocatable,proto3" json:"allocatable" yaml:"allocatable"`
	Allocated   *resource.Quantity                                                    `protobuf:"bytes,2,opt,name=allocated,proto3" json:"allocated" yaml:"allocated"`
	Attributes  github_com_spheronFdn_akash_api_fork_go_node_types_v1beta3.Attributes `protobuf:"bytes,3,rep,name=attributes,proto3,castrepeated=github.com/spheronFdn/akash-api-fork/go/node/types/v1beta3.Attributes" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

func (m *ResourcePair) Reset()         { *m = ResourcePair{} }
func (m *ResourcePair) String() string { return proto.CompactTextString(m) }
func (*ResourcePair) ProtoMessage()    {}
func (*ResourcePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_995cee7bf7b692e7, []int{0}
}
func (m *ResourcePair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourcePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourcePair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourcePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcePair.Merge(m, src)
}
func (m *ResourcePair) XXX_Size() int {
	return m.Size()
}
func (m *ResourcePair) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcePair.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcePair proto.InternalMessageInfo

func (m *ResourcePair) GetAllocatable() *resource.Quantity {
	if m != nil {
		return m.Allocatable
	}
	return nil
}

func (m *ResourcePair) GetAllocated() *resource.Quantity {
	if m != nil {
		return m.Allocated
	}
	return nil
}

func (m *ResourcePair) GetAttributes() github_com_spheronFdn_akash_api_fork_go_node_types_v1beta3.Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourcePair)(nil), "akash.inventory.v1.ResourcePair")
}

func init() {
	proto.RegisterFile("akash/inventory/v1/resourcepair.proto", fileDescriptor_995cee7bf7b692e7)
}

var fileDescriptor_995cee7bf7b692e7 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x8f, 0xd3, 0x30,
	0x14, 0xc6, 0x63, 0x8a, 0x90, 0x2e, 0x65, 0x40, 0xd1, 0x0d, 0xd5, 0x21, 0xe2, 0x53, 0x24, 0xa4,
	0x1b, 0x38, 0x5b, 0xed, 0xdd, 0x70, 0xba, 0xad, 0x95, 0x60, 0x44, 0xd0, 0x91, 0xed, 0xa5, 0x79,
	0xa4, 0x56, 0x9a, 0xd8, 0xb2, 0xdd, 0x4a, 0xd9, 0x18, 0x98, 0x98, 0xf8, 0x13, 0x98, 0xf9, 0x3b,
	0x18, 0x3a, 0xde, 0xc8, 0x64, 0x50, 0xbb, 0xa0, 0xb2, 0xf5, 0x2f, 0x40, 0x6d, 0xd2, 0x26, 0x08,
	0x06, 0xc4, 0x96, 0xbc, 0xef, 0xf3, 0xef, 0xfb, 0xe4, 0x67, 0xff, 0x29, 0x64, 0x60, 0xa6, 0x5c,
	0x14, 0x0b, 0x2c, 0xac, 0xd4, 0x25, 0x5f, 0xf4, 0xb9, 0x46, 0x23, 0xe7, 0x7a, 0x82, 0x0a, 0x84,
	0x66, 0x4a, 0x4b, 0x2b, 0x83, 0x60, 0x6f, 0x63, 0x47, 0x1b, 0x5b, 0xf4, 0xcf, 0x4e, 0x53, 0x99,
	0xca, 0xbd, 0xcc, 0x77, 0x5f, 0x95, 0xf3, 0x2c, 0xaa, 0x80, 0x31, 0x18, 0xe4, 0x8b, 0x7e, 0x8c,
	0x16, 0xae, 0x38, 0x58, 0xab, 0x45, 0x3c, 0xb7, 0x58, 0x7b, 0xae, 0xb3, 0x1b, 0xc3, 0x84, 0xe4,
	0xa0, 0x44, 0x0e, 0x93, 0xa9, 0x28, 0x50, 0x97, 0x5c, 0x65, 0xe9, 0x6e, 0x70, 0x4c, 0xe7, 0x29,
	0x16, 0xa8, 0xc1, 0x62, 0x52, 0x9d, 0x8a, 0x7e, 0x76, 0xfc, 0x87, 0xe3, 0x5a, 0x7c, 0x05, 0x42,
	0x07, 0x1f, 0x88, 0xdf, 0x85, 0xd9, 0x4c, 0x4e, 0xc0, 0x42, 0x3c, 0xc3, 0x1e, 0x39, 0x27, 0x17,
	0xdd, 0x01, 0x63, 0x15, 0x9d, 0xb5, 0xe9, 0x4c, 0x65, 0xe9, 0x6e, 0xc0, 0x0e, 0x74, 0xf6, 0x7a,
	0x0e, 0x85, 0x15, 0xb6, 0x1c, 0xdd, 0x2c, 0x1d, 0x25, 0x2b, 0x47, 0xbb, 0xc3, 0x06, 0xb5, 0x71,
	0xb4, 0x4d, 0xde, 0x3a, 0x1a, 0x94, 0x90, 0xcf, 0x6e, 0xa3, 0xd6, 0x30, 0x1a, 0xb7, 0x2d, 0xc1,
	0x3b, 0xe2, 0x9f, 0xd4, 0xff, 0x98, 0xf4, 0xee, 0xfd, 0x57, 0x95, 0x41, 0x5d, 0xe5, 0x64, 0x78,
	0x00, 0x6d, 0x1c, 0x6d, 0xa8, 0x5b, 0x47, 0x1f, 0xfd, 0x56, 0x03, 0x93, 0x68, 0xdc, 0xc8, 0xc1,
	0x17, 0xe2, 0xfb, 0xc7, 0xab, 0x36, 0xbd, 0xce, 0x79, 0xe7, 0xa2, 0x3b, 0x78, 0xc2, 0xaa, 0xd5,
	0xed, 0x16, 0xc2, 0xea, 0x85, 0xb0, 0xe1, 0xc1, 0x35, 0x7a, 0x4f, 0x96, 0x8e, 0x7a, 0x1b, 0x47,
	0x4f, 0x9b, 0x93, 0xcf, 0x64, 0x2e, 0x2c, 0xe6, 0xca, 0x96, 0x5b, 0x47, 0x1f, 0xd7, 0x89, 0x7f,
	0x51, 0xa3, 0xcf, 0xdf, 0xe8, 0xf3, 0x54, 0xd8, 0xe9, 0x3c, 0x66, 0x13, 0x99, 0x73, 0xa3, 0xa6,
	0xa8, 0x65, 0xf1, 0x22, 0x29, 0xf8, 0x3e, 0xf3, 0x12, 0x94, 0xb8, 0x7c, 0x2b, 0x75, 0xc6, 0x53,
	0xc9, 0x0b, 0x99, 0x20, 0xb7, 0xa5, 0x42, 0xc3, 0xff, 0x68, 0x61, 0xc6, 0xad, 0xde, 0xb7, 0xf7,
	0x7f, 0x7c, 0xa2, 0xde, 0xe8, 0xe5, 0x72, 0x15, 0x92, 0xbb, 0x55, 0x48, 0xbe, 0xaf, 0x42, 0xf2,
	0x71, 0x1d, 0x7a, 0x77, 0xeb, 0xd0, 0xfb, 0xba, 0x0e, 0xbd, 0x37, 0xd7, 0xff, 0x1a, 0xd9, 0x7e,
	0xd4, 0xf1, 0x83, 0xfd, 0x23, 0xba, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xd3, 0x8e, 0x12,
	0xf1, 0x02, 0x00, 0x00,
}

func (m *ResourcePair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourcePair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourcePair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResourcepair(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Allocated != nil {
		{
			size, err := m.Allocated.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResourcepair(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Allocatable != nil {
		{
			size, err := m.Allocatable.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResourcepair(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResourcepair(dAtA []byte, offset int, v uint64) int {
	offset -= sovResourcepair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResourcePair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Allocatable != nil {
		l = m.Allocatable.Size()
		n += 1 + l + sovResourcepair(uint64(l))
	}
	if m.Allocated != nil {
		l = m.Allocated.Size()
		n += 1 + l + sovResourcepair(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovResourcepair(uint64(l))
		}
	}
	return n
}

func sovResourcepair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResourcepair(x uint64) (n int) {
	return sovResourcepair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourcePair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourcepair
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
			return fmt.Errorf("proto: ResourcePair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourcePair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocatable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcepair
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
				return ErrInvalidLengthResourcepair
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourcepair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allocatable == nil {
				m.Allocatable = &resource.Quantity{}
			}
			if err := m.Allocatable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcepair
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
				return ErrInvalidLengthResourcepair
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourcepair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allocated == nil {
				m.Allocated = &resource.Quantity{}
			}
			if err := m.Allocated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcepair
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
				return ErrInvalidLengthResourcepair
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourcepair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, v1beta3.Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourcepair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResourcepair
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
func skipResourcepair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourcepair
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
					return 0, ErrIntOverflowResourcepair
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
					return 0, ErrIntOverflowResourcepair
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
				return 0, ErrInvalidLengthResourcepair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResourcepair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResourcepair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResourcepair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourcepair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResourcepair = fmt.Errorf("proto: unexpected end of group")
)
