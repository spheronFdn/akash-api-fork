// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta3/group.proto

package v1beta3

import (
	fmt "fmt"
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

// State is an enum which refers to state of group
type Group_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	GroupStateInvalid Group_State = 0
	// GroupOpen denotes state for group open
	GroupOpen Group_State = 1
	// GroupOrdered denotes state for group ordered
	GroupPaused Group_State = 2
	// GroupInsufficientFunds denotes state for group insufficient_funds
	GroupInsufficientFunds Group_State = 3
	// GroupClosed denotes state for group closed
	GroupClosed Group_State = 4
)

var Group_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "paused",
	3: "insufficient_funds",
	4: "closed",
}

var Group_State_value = map[string]int32{
	"invalid":            0,
	"open":               1,
	"paused":             2,
	"insufficient_funds": 3,
	"closed":             4,
}

func (x Group_State) String() string {
	return proto.EnumName(Group_State_name, int32(x))
}

func (Group_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6dcebaf6050ef7f1, []int{0, 0}
}

// Group stores group id, state and specifications of group
type Group struct {
	GroupID   GroupID     `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"id" yaml:"id"`
	State     Group_State `protobuf:"varint,2,opt,name=state,proto3,enum=akash.deployment.v1beta3.Group_State" json:"state" yaml:"state"`
	GroupSpec GroupSpec   `protobuf:"bytes,3,opt,name=group_spec,json=groupSpec,proto3" json:"spec" yaml:"spec"`
	CreatedAt int64       `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dcebaf6050ef7f1, []int{0}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Group.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return m.Size()
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetGroupID() GroupID {
	if m != nil {
		return m.GroupID
	}
	return GroupID{}
}

func (m *Group) GetState() Group_State {
	if m != nil {
		return m.State
	}
	return GroupStateInvalid
}

func (m *Group) GetGroupSpec() GroupSpec {
	if m != nil {
		return m.GroupSpec
	}
	return GroupSpec{}
}

func (m *Group) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterEnum("akash.deployment.v1beta3.Group_State", Group_State_name, Group_State_value)
	proto.RegisterType((*Group)(nil), "akash.deployment.v1beta3.Group")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta3/group.proto", fileDescriptor_6dcebaf6050ef7f1)
}

var fileDescriptor_6dcebaf6050ef7f1 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3d, 0x6f, 0xd3, 0x4e,
	0x1c, 0xc7, 0xed, 0xc4, 0x69, 0xfe, 0xb9, 0xfc, 0x81, 0x70, 0xe2, 0xc1, 0xb8, 0xc2, 0x36, 0xe6,
	0x41, 0x59, 0x6a, 0x8b, 0x74, 0xeb, 0x82, 0x08, 0xa8, 0x28, 0x13, 0x28, 0x95, 0x90, 0x60, 0x09,
	0x17, 0xdf, 0xc5, 0x39, 0x35, 0xb9, 0x3b, 0xd9, 0xe7, 0x4a, 0x5d, 0x99, 0x50, 0x26, 0xde, 0x40,
	0x24, 0x24, 0xde, 0x04, 0x2f, 0xa1, 0x63, 0x47, 0x26, 0x0b, 0x25, 0x0b, 0xca, 0x98, 0x57, 0x80,
	0xee, 0x6c, 0x44, 0x07, 0x68, 0x37, 0xfb, 0xfb, 0xfd, 0xfc, 0x3e, 0xfe, 0xf9, 0x74, 0xe0, 0x11,
	0x3a, 0x46, 0xd9, 0x34, 0xc2, 0x44, 0xcc, 0xf8, 0xe9, 0x9c, 0x30, 0x19, 0x9d, 0x3c, 0x1d, 0x13,
	0x89, 0xf6, 0xa3, 0x24, 0xe5, 0xb9, 0x08, 0x45, 0xca, 0x25, 0x87, 0xb6, 0xa6, 0xc2, 0x3f, 0x54,
	0x58, 0x51, 0xce, 0xad, 0x84, 0x27, 0x5c, 0x43, 0x91, 0x7a, 0x2a, 0x79, 0xe7, 0xc9, 0xe5, 0x56,
	0x8a, 0x2b, 0xae, 0x7b, 0x39, 0x97, 0x09, 0x12, 0x97, 0x64, 0xf0, 0xd1, 0x02, 0x8d, 0x57, 0x2a,
	0x83, 0x1f, 0xc0, 0x7f, 0xba, 0x1c, 0x51, 0x6c, 0x9b, 0xbe, 0xd9, 0x6d, 0xf7, 0x1e, 0x84, 0xff,
	0x5a, 0x2f, 0xd4, 0x23, 0x83, 0x97, 0xfd, 0xe0, 0xac, 0xf0, 0x8c, 0x55, 0xe1, 0x35, 0xab, 0x60,
	0x53, 0x78, 0x35, 0x8a, 0xb7, 0x85, 0xd7, 0x3a, 0x45, 0xf3, 0xd9, 0x41, 0x40, 0x71, 0x30, 0x6c,
	0x6a, 0xed, 0x00, 0xc3, 0xb7, 0xa0, 0x91, 0x49, 0x24, 0x89, 0x5d, 0xf3, 0xcd, 0xee, 0xf5, 0xde,
	0xe3, 0x2b, 0xf4, 0xe1, 0x91, 0x82, 0xfb, 0xf7, 0x36, 0x85, 0x57, 0xce, 0x6d, 0x0b, 0xef, 0xff,
	0x52, 0xab, 0x5f, 0x83, 0x61, 0x19, 0xc3, 0x11, 0x00, 0xe5, 0xe6, 0xea, 0xbf, 0xec, 0xba, 0xde,
	0xfd, 0xe1, 0x15, 0xf2, 0x23, 0x41, 0xe2, 0xfe, 0xae, 0xda, 0x7e, 0x53, 0x78, 0x96, 0x1a, 0xdc,
	0x16, 0x5e, 0xbb, 0xb2, 0x0b, 0x12, 0x07, 0xc3, 0x56, 0xf2, 0x9b, 0x83, 0xf7, 0x01, 0x88, 0x53,
	0x82, 0x24, 0xc1, 0x23, 0x24, 0x6d, 0xcb, 0x37, 0xbb, 0xf5, 0x61, 0xab, 0x4a, 0x9e, 0xcb, 0xe0,
	0x9b, 0x09, 0x1a, 0x7a, 0x57, 0x18, 0x80, 0x26, 0x65, 0x27, 0x68, 0x46, 0x71, 0xc7, 0x70, 0x6e,
	0x2f, 0x96, 0xfe, 0xcd, 0xf2, 0x63, 0xaa, 0x1c, 0x94, 0x05, 0xbc, 0x0b, 0x2c, 0x2e, 0x08, 0xeb,
	0x98, 0xce, 0xb5, 0xc5, 0xd2, 0x6f, 0x69, 0xe0, 0xb5, 0x20, 0x0c, 0xee, 0x82, 0x1d, 0x81, 0xf2,
	0x8c, 0xe0, 0x4e, 0xcd, 0xb9, 0xb1, 0x58, 0xfa, 0x6d, 0x5d, 0xbd, 0xd1, 0x11, 0xec, 0x01, 0x48,
	0x59, 0x96, 0x4f, 0x26, 0x34, 0xa6, 0x84, 0xc9, 0xd1, 0x24, 0x67, 0x38, 0xeb, 0xd4, 0x1d, 0x67,
	0xb1, 0xf4, 0xef, 0x94, 0x87, 0x7f, 0xa1, 0x3e, 0x54, 0xad, 0x12, 0xc6, 0x33, 0xae, 0x84, 0xd6,
	0x05, 0xe1, 0x0b, 0x1d, 0x39, 0xd6, 0xa7, 0xaf, 0xae, 0x71, 0x60, 0xfd, 0xfc, 0xe2, 0x19, 0xfd,
	0x77, 0x67, 0x2b, 0xd7, 0x3c, 0x5f, 0xb9, 0xe6, 0x8f, 0x95, 0x6b, 0x7e, 0x5e, 0xbb, 0xc6, 0xf9,
	0xda, 0x35, 0xbe, 0xaf, 0x5d, 0xe3, 0xfd, 0xb3, 0x84, 0xca, 0x69, 0x3e, 0x0e, 0x63, 0x3e, 0x8f,
	0x32, 0x31, 0x25, 0x29, 0x67, 0x87, 0x98, 0x45, 0xfa, 0x6c, 0xf7, 0x90, 0xa0, 0x7b, 0x13, 0x9e,
	0x1e, 0x47, 0x09, 0x8f, 0x18, 0xc7, 0xe4, 0x2f, 0xf7, 0x6d, 0xbc, 0xa3, 0xaf, 0xd9, 0xfe, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0xdf, 0x87, 0xd6, 0x10, 0x03, 0x00, 0x00,
}

func (m *Group) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Group) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.GroupSpec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.GroupID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Group) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GroupID.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.State != 0 {
		n += 1 + sovGroup(uint64(m.State))
	}
	l = m.GroupSpec.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.CreatedAt != 0 {
		n += 1 + sovGroup(uint64(m.CreatedAt))
	}
	return n
}

func sovGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroup(x uint64) (n int) {
	return sovGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: Group: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Group: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Group_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroup
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
func skipGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
				return 0, ErrInvalidLengthGroup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroup = fmt.Errorf("proto: unexpected end of group")
)
