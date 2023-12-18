// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/inventory/v1/gpu.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GPUInfo struct {
	Vendor     string `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor" yaml:"vendor"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" yaml:"name"`
	ModelID    string `protobuf:"bytes,3,opt,name=modelid,proto3" json:"model_id" yaml:"model_id"`
	Interface  string `protobuf:"bytes,4,opt,name=interface,proto3" json:"interface" yaml:"interface"`
	MemorySize string `protobuf:"bytes,5,opt,name=memory_size,json=memorySize,proto3" json:"memory_size" yaml:"memory_size"`
}

func (m *GPUInfo) Reset()         { *m = GPUInfo{} }
func (m *GPUInfo) String() string { return proto.CompactTextString(m) }
func (*GPUInfo) ProtoMessage()    {}
func (*GPUInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc01b12bd00ffcc, []int{0}
}
func (m *GPUInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GPUInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GPUInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GPUInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPUInfo.Merge(m, src)
}
func (m *GPUInfo) XXX_Size() int {
	return m.Size()
}
func (m *GPUInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GPUInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GPUInfo proto.InternalMessageInfo

func (m *GPUInfo) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *GPUInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GPUInfo) GetModelID() string {
	if m != nil {
		return m.ModelID
	}
	return ""
}

func (m *GPUInfo) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *GPUInfo) GetMemorySize() string {
	if m != nil {
		return m.MemorySize
	}
	return ""
}

type GPU struct {
	Quantity ResourcePair `protobuf:"bytes,1,opt,name=quantity,proto3" json:"quantity" yaml:"quantity"`
	Info     GPUInfo      `protobuf:"bytes,2,opt,name=info,proto3" json:"info" yaml:"info"`
}

func (m *GPU) Reset()         { *m = GPU{} }
func (m *GPU) String() string { return proto.CompactTextString(m) }
func (*GPU) ProtoMessage()    {}
func (*GPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc01b12bd00ffcc, []int{1}
}
func (m *GPU) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GPU.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPU.Merge(m, src)
}
func (m *GPU) XXX_Size() int {
	return m.Size()
}
func (m *GPU) XXX_DiscardUnknown() {
	xxx_messageInfo_GPU.DiscardUnknown(m)
}

var xxx_messageInfo_GPU proto.InternalMessageInfo

func (m *GPU) GetQuantity() ResourcePair {
	if m != nil {
		return m.Quantity
	}
	return ResourcePair{}
}

func (m *GPU) GetInfo() GPUInfo {
	if m != nil {
		return m.Info
	}
	return GPUInfo{}
}

func init() {
	proto.RegisterType((*GPUInfo)(nil), "akash.inventory.v1.GPUInfo")
	proto.RegisterType((*GPU)(nil), "akash.inventory.v1.GPU")
}

func init() { proto.RegisterFile("akash/inventory/v1/gpu.proto", fileDescriptor_2cc01b12bd00ffcc) }

var fileDescriptor_2cc01b12bd00ffcc = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xe3, 0xd6, 0xcb, 0x1f, 0x85, 0xb1, 0x21, 0x76, 0x30, 0xdd, 0xb0, 0x8a, 0x46, 0x61,
	0x0c, 0x66, 0xd3, 0xe4, 0x36, 0xd8, 0x25, 0x14, 0x4a, 0x60, 0x1d, 0x99, 0xd7, 0xee, 0xb0, 0x4b,
	0x51, 0x13, 0xc5, 0x15, 0xad, 0xa4, 0x4c, 0x71, 0x32, 0xd2, 0x4f, 0xb1, 0xe3, 0x3e, 0x52, 0xd9,
	0xa9, 0xc7, 0x9d, 0xc4, 0x50, 0x0e, 0x83, 0x1c, 0xfd, 0x09, 0x86, 0x64, 0x3b, 0xde, 0x68, 0x6f,
	0x7e, 0x7f, 0xcf, 0xfb, 0x3c, 0x32, 0x7a, 0x04, 0x5e, 0x90, 0x2b, 0x32, 0xbf, 0x8c, 0x99, 0x58,
	0x52, 0x91, 0x49, 0xb5, 0x8a, 0x97, 0x87, 0x71, 0x3a, 0x5b, 0x44, 0x33, 0x25, 0x33, 0x09, 0xa1,
	0x53, 0xa3, 0xad, 0x1a, 0x2d, 0x0f, 0xf7, 0x9e, 0xa5, 0x32, 0x95, 0x4e, 0x8e, 0xed, 0x57, 0xb1,
	0xb9, 0x77, 0xf0, 0x40, 0x8e, 0xa2, 0x73, 0xb9, 0x50, 0x63, 0x3a, 0x23, 0x4c, 0x15, 0x6b, 0xf8,
	0xcf, 0x0e, 0x68, 0x1d, 0x8f, 0xce, 0x86, 0x62, 0x2a, 0xe1, 0x3b, 0xd0, 0x5c, 0x52, 0x31, 0x91,
	0x2a, 0xf0, 0xf6, 0xbd, 0x57, 0x9d, 0xc1, 0x81, 0xd1, 0xa8, 0xf9, 0xd9, 0x91, 0x8d, 0x46, 0xa5,
	0x96, 0x6b, 0xf4, 0x78, 0x45, 0xf8, 0xf5, 0x5b, 0x5c, 0xcc, 0x38, 0x29, 0x05, 0xd8, 0x07, 0xbe,
	0x20, 0x9c, 0x06, 0x3b, 0xce, 0x8c, 0x8c, 0x46, 0xfe, 0x07, 0xc2, 0xe9, 0x46, 0x23, 0xc7, 0x73,
	0x8d, 0xba, 0x85, 0xd1, 0x4e, 0x38, 0x71, 0x10, 0x1e, 0x81, 0x16, 0x97, 0x13, 0x7a, 0xcd, 0x26,
	0xc1, 0xae, 0xf3, 0xbd, 0x36, 0x1a, 0xb5, 0x4e, 0x2c, 0x1a, 0x1e, 0x6d, 0x34, 0x6a, 0x3b, 0xf5,
	0x9c, 0x4d, 0x72, 0x8d, 0x9e, 0x14, 0xf6, 0x8a, 0xe0, 0xa4, 0xb2, 0xc2, 0xf7, 0xa0, 0xc3, 0x44,
	0x46, 0xd5, 0x94, 0x8c, 0x69, 0xe0, 0xbb, 0x9c, 0xc8, 0x68, 0xd4, 0x19, 0x56, 0x70, 0xa3, 0x51,
	0xbd, 0x91, 0x6b, 0xf4, 0xb4, 0x88, 0xda, 0x22, 0x9c, 0xd4, 0x32, 0x3c, 0x05, 0x5d, 0x4e, 0xb9,
	0x54, 0xab, 0xf3, 0x39, 0xbb, 0xa1, 0xc1, 0x23, 0x97, 0xd7, 0x37, 0x1a, 0x81, 0x13, 0x87, 0x3f,
	0xb1, 0x1b, 0x1b, 0xf8, 0xef, 0x52, 0xae, 0x11, 0x2c, 0xff, 0xae, 0x86, 0x38, 0x01, 0x7c, 0x6b,
	0xc0, 0x3f, 0x3d, 0xb0, 0x7b, 0x3c, 0x3a, 0x83, 0x1c, 0xb4, 0xbf, 0x2e, 0x88, 0xc8, 0x58, 0xb6,
	0x72, 0xf7, 0xdc, 0xed, 0xed, 0x47, 0xf7, 0x5b, 0x8d, 0x92, 0xb2, 0xab, 0x11, 0x61, 0x6a, 0x10,
	0xdf, 0x6a, 0xd4, 0x30, 0x1a, 0xb5, 0x3f, 0x96, 0x4e, 0x7b, 0x33, 0x55, 0x4a, 0x7d, 0x33, 0x15,
	0xc1, 0xc9, 0x56, 0x84, 0xa7, 0xc0, 0x67, 0x62, 0x2a, 0x5d, 0x2b, 0xdd, 0xde, 0xf3, 0x87, 0x8e,
	0x2a, 0xfb, 0x1f, 0xbc, 0x2c, 0x4f, 0xf1, 0xed, 0x64, 0x6b, 0xb3, 0xc6, 0xba, 0x36, 0x3b, 0xe1,
	0xc4, 0xc1, 0xc1, 0xe8, 0xd6, 0x84, 0xde, 0x9d, 0x09, 0xbd, 0xdf, 0x26, 0xf4, 0xbe, 0xaf, 0xc3,
	0xc6, 0x8f, 0x75, 0xd8, 0xb8, 0x5b, 0x87, 0x8d, 0x5f, 0xeb, 0xb0, 0xf1, 0xa5, 0x97, 0xb2, 0xec,
	0x72, 0x71, 0x11, 0x8d, 0x25, 0x8f, 0xdd, 0x79, 0x6f, 0x04, 0xcd, 0xbe, 0x49, 0x75, 0x55, 0x4e,
	0x64, 0xc6, 0xe2, 0x54, 0xfe, 0xf7, 0x36, 0x2f, 0x9a, 0xee, 0x3d, 0xf6, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x2c, 0x99, 0x43, 0xfe, 0x00, 0x03, 0x00, 0x00,
}

func (m *GPUInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GPUInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GPUInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MemorySize) > 0 {
		i -= len(m.MemorySize)
		copy(dAtA[i:], m.MemorySize)
		i = encodeVarintGpu(dAtA, i, uint64(len(m.MemorySize)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Interface) > 0 {
		i -= len(m.Interface)
		copy(dAtA[i:], m.Interface)
		i = encodeVarintGpu(dAtA, i, uint64(len(m.Interface)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ModelID) > 0 {
		i -= len(m.ModelID)
		copy(dAtA[i:], m.ModelID)
		i = encodeVarintGpu(dAtA, i, uint64(len(m.ModelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGpu(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Vendor) > 0 {
		i -= len(m.Vendor)
		copy(dAtA[i:], m.Vendor)
		i = encodeVarintGpu(dAtA, i, uint64(len(m.Vendor)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GPU) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GPU) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GPU) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGpu(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Quantity.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGpu(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGpu(dAtA []byte, offset int, v uint64) int {
	offset -= sovGpu(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GPUInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Vendor)
	if l > 0 {
		n += 1 + l + sovGpu(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGpu(uint64(l))
	}
	l = len(m.ModelID)
	if l > 0 {
		n += 1 + l + sovGpu(uint64(l))
	}
	l = len(m.Interface)
	if l > 0 {
		n += 1 + l + sovGpu(uint64(l))
	}
	l = len(m.MemorySize)
	if l > 0 {
		n += 1 + l + sovGpu(uint64(l))
	}
	return n
}

func (m *GPU) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Quantity.Size()
	n += 1 + l + sovGpu(uint64(l))
	l = m.Info.Size()
	n += 1 + l + sovGpu(uint64(l))
	return n
}

func sovGpu(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGpu(x uint64) (n int) {
	return sovGpu(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GPUInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGpu
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
			return fmt.Errorf("proto: GPUInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GPUInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vendor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGpu
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
				return ErrInvalidLengthGpu
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGpu
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vendor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGpu
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
				return ErrInvalidLengthGpu
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGpu
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGpu
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
				return ErrInvalidLengthGpu
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGpu
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interface", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGpu
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
				return ErrInvalidLengthGpu
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGpu
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Interface = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemorySize", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGpu
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
				return ErrInvalidLengthGpu
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGpu
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MemorySize = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGpu(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGpu
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
func (m *GPU) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGpu
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
			return fmt.Errorf("proto: GPU: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GPU: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGpu
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
				return ErrInvalidLengthGpu
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGpu
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGpu
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
				return ErrInvalidLengthGpu
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGpu
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGpu(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGpu
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
func skipGpu(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGpu
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
					return 0, ErrIntOverflowGpu
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
					return 0, ErrIntOverflowGpu
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
				return 0, ErrInvalidLengthGpu
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGpu
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGpu
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGpu        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGpu          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGpu = fmt.Errorf("proto: unexpected end of group")
)
