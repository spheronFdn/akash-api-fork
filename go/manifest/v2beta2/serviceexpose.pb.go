// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/manifest/v2beta2/serviceexpose.proto

package v2beta2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// ServiceExpose stores exposed ports and hosts details
type ServiceExpose struct {
	// port on the container
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port" yaml:"port"`
	// port on the service definition
	ExternalPort uint32                   `protobuf:"varint,2,opt,name=external_port,json=externalPort,proto3" json:"externalPort" yaml:"externalPort"`
	Proto        ServiceProtocol          `protobuf:"bytes,3,opt,name=proto,proto3,casttype=ServiceProtocol" json:"proto" yaml:"proto"`
	Service      string                   `protobuf:"bytes,4,opt,name=service,proto3" json:"service" yaml:"service"`
	Global       bool                     `protobuf:"varint,5,opt,name=global,proto3" json:"global" yaml:"global"`
	Hosts        []string                 `protobuf:"bytes,6,rep,name=hosts,proto3" json:"hosts" yaml:"hosts"`
	HTTPOptions  ServiceExposeHTTPOptions `protobuf:"bytes,7,opt,name=http_options,json=httpOptions,proto3" json:"httpOptions" yaml:"httpOptions"`
	// The name of the IP address associated with this, if any
	IP string `protobuf:"bytes,8,opt,name=ip,proto3" json:"ip" yaml:"ip"`
	// The sequence number of the associated endpoint in the on-chain data
	EndpointSequenceNumber uint32 `protobuf:"varint,9,opt,name=endpoint_sequence_number,json=endpointSequenceNumber,proto3" json:"endpointSequenceNumber" yaml:"endpointSequenceNumber"`
}

func (m *ServiceExpose) Reset()      { *m = ServiceExpose{} }
func (*ServiceExpose) ProtoMessage() {}
func (*ServiceExpose) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8dad39b3d78f39d, []int{0}
}
func (m *ServiceExpose) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceExpose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceExpose.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceExpose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceExpose.Merge(m, src)
}
func (m *ServiceExpose) XXX_Size() int {
	return m.Size()
}
func (m *ServiceExpose) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceExpose.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceExpose proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ServiceExpose)(nil), "akash.manifest.v2beta2.ServiceExpose")
}

func init() {
	proto.RegisterFile("akash/manifest/v2beta2/serviceexpose.proto", fileDescriptor_e8dad39b3d78f39d)
}

var fileDescriptor_e8dad39b3d78f39d = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xed, 0xb4, 0x49, 0x9b, 0x4b, 0x02, 0xd2, 0x81, 0x8a, 0x5b, 0x54, 0x5f, 0xe4, 0x05,
	0xf3, 0xa7, 0x36, 0x4a, 0x07, 0xaa, 0xb2, 0x59, 0x02, 0x81, 0x84, 0x20, 0x72, 0x3b, 0x75, 0x89,
	0x9c, 0xf4, 0x9a, 0x9c, 0x9a, 0xf8, 0x0e, 0xfb, 0x52, 0x95, 0x8d, 0x91, 0x05, 0x89, 0x8f, 0xc0,
	0xc2, 0x77, 0xe9, 0xd8, 0xb1, 0xd3, 0x09, 0x9c, 0xcd, 0xa3, 0x47, 0x26, 0xe4, 0xbb, 0xb3, 0x9a,
	0x8a, 0x76, 0xf3, 0xfb, 0xbc, 0xbf, 0xe7, 0xbd, 0x47, 0xf7, 0x9e, 0xc1, 0xb3, 0xe8, 0x34, 0x4a,
	0x27, 0xfe, 0x2c, 0x8a, 0xc9, 0x09, 0x4e, 0xb9, 0x7f, 0xd6, 0x1b, 0x62, 0x1e, 0xf5, 0xfc, 0x14,
	0x27, 0x67, 0x64, 0x84, 0xf1, 0x39, 0xa3, 0x29, 0xf6, 0x58, 0x42, 0x39, 0x85, 0x1b, 0x92, 0xf5,
	0x2a, 0xd6, 0xd3, 0xec, 0xd6, 0xc3, 0x31, 0x1d, 0x53, 0x89, 0xf8, 0xe5, 0x97, 0xa2, 0xb7, 0xdc,
	0x3b, 0x26, 0x4f, 0x38, 0x67, 0x94, 0x71, 0x42, 0xe3, 0x54, 0x91, 0xce, 0xaf, 0x3a, 0xe8, 0x1c,
	0xa8, 0xf3, 0xde, 0xc8, 0xf3, 0xe0, 0x73, 0xb0, 0xca, 0x68, 0xc2, 0x2d, 0xb3, 0x6b, 0xba, 0x9d,
	0xe0, 0x51, 0x2e, 0x90, 0xac, 0x0b, 0x81, 0x5a, 0x5f, 0xa2, 0xd9, 0x74, 0xdf, 0x29, 0x2b, 0x27,
	0x94, 0x22, 0xfc, 0x00, 0x3a, 0xf8, 0x9c, 0xe3, 0x24, 0x8e, 0xa6, 0x03, 0xe9, 0xaa, 0x49, 0xd7,
	0x93, 0x5c, 0xa0, 0x76, 0xd5, 0xe8, 0x2b, 0xf7, 0x03, 0xe5, 0x5e, 0x56, 0x9d, 0xf0, 0x06, 0x04,
	0x03, 0x50, 0x97, 0xa9, 0xac, 0x95, 0xae, 0xe9, 0x36, 0x83, 0x17, 0xb9, 0x40, 0x4a, 0x28, 0x04,
	0x6a, 0xeb, 0xc3, 0x65, 0xea, 0xbf, 0x02, 0xdd, 0xd7, 0xa9, 0xfb, 0xa5, 0x30, 0xa2, 0xd3, 0x50,
	0x91, 0xf0, 0x15, 0x58, 0xd3, 0xf7, 0x67, 0xad, 0xca, 0x29, 0xdb, 0xb9, 0x40, 0x95, 0x54, 0x08,
	0x74, 0x4f, 0xcd, 0xd1, 0x82, 0x13, 0x56, 0x2d, 0xb8, 0x0b, 0x1a, 0xe3, 0x29, 0x1d, 0x46, 0x53,
	0xab, 0xde, 0x35, 0xdd, 0xf5, 0xe0, 0x71, 0x2e, 0x90, 0x56, 0x0a, 0x81, 0x3a, 0xca, 0xa6, 0x6a,
	0x27, 0xd4, 0x0d, 0xe8, 0x83, 0xfa, 0x84, 0xa6, 0x3c, 0xb5, 0x1a, 0xdd, 0x15, 0xb7, 0x19, 0x6c,
	0x96, 0x89, 0xa5, 0x70, 0x9d, 0x58, 0x96, 0x4e, 0xa8, 0x64, 0xf8, 0xdd, 0x04, 0xed, 0x72, 0x0b,
	0x03, 0xbd, 0x06, 0x6b, 0xad, 0x6b, 0xba, 0xad, 0xde, 0x4b, 0xef, 0xf6, 0xfd, 0x7a, 0x37, 0x76,
	0xf3, 0xee, 0xf0, 0xb0, 0xff, 0x49, 0xf9, 0x82, 0xbd, 0x0b, 0x81, 0x8c, 0x4c, 0xa0, 0xd6, 0x92,
	0x98, 0x0b, 0xd4, 0x2a, 0x87, 0xeb, 0xb2, 0x10, 0x08, 0xea, 0x0c, 0xd7, 0xa2, 0x13, 0x2e, 0x23,
	0xf0, 0x29, 0xa8, 0x11, 0x66, 0xad, 0xcb, 0x9b, 0xda, 0xcc, 0x04, 0xaa, 0xbd, 0xef, 0xe7, 0x02,
	0xd5, 0x08, 0x2b, 0x04, 0x6a, 0x2a, 0x33, 0x61, 0x4e, 0x58, 0x23, 0x0c, 0xce, 0x81, 0x85, 0xe3,
	0x63, 0x46, 0x49, 0xcc, 0x07, 0x29, 0xfe, 0x3c, 0xc7, 0xf1, 0x08, 0x0f, 0xe2, 0xf9, 0x6c, 0x88,
	0x13, 0xab, 0x29, 0xd7, 0xfe, 0x3a, 0x17, 0x68, 0xa3, 0x62, 0x0e, 0x34, 0xf2, 0x51, 0x12, 0x85,
	0x40, 0xdb, 0xfa, 0x01, 0xdc, 0xda, 0x77, 0xc2, 0x3b, 0x8c, 0xfb, 0xab, 0xdf, 0x7e, 0x22, 0x23,
	0x38, 0xba, 0xfa, 0x63, 0x1b, 0x5f, 0x33, 0xdb, 0xbc, 0xc8, 0x6c, 0xf3, 0x32, 0xb3, 0xcd, 0xdf,
	0x99, 0x6d, 0xfe, 0x58, 0xd8, 0xc6, 0xe5, 0xc2, 0x36, 0xae, 0x16, 0xb6, 0x71, 0xb4, 0x37, 0x26,
	0x7c, 0x32, 0x1f, 0x7a, 0x23, 0x3a, 0xf3, 0x53, 0x36, 0xc1, 0x09, 0x8d, 0xdf, 0x1e, 0xc7, 0xbe,
	0xbc, 0xd7, 0x9d, 0x88, 0x91, 0x9d, 0x13, 0x9a, 0x9c, 0xfa, 0x63, 0xfa, 0xdf, 0x5f, 0x31, 0x6c,
	0xc8, 0x97, 0xb3, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xc6, 0x90, 0x8e, 0x90, 0x03, 0x00,
	0x00,
}

func (m *ServiceExpose) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceExpose) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceExpose) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndpointSequenceNumber != 0 {
		i = encodeVarintServiceexpose(dAtA, i, uint64(m.EndpointSequenceNumber))
		i--
		dAtA[i] = 0x48
	}
	if len(m.IP) > 0 {
		i -= len(m.IP)
		copy(dAtA[i:], m.IP)
		i = encodeVarintServiceexpose(dAtA, i, uint64(len(m.IP)))
		i--
		dAtA[i] = 0x42
	}
	{
		size, err := m.HTTPOptions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintServiceexpose(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Hosts) > 0 {
		for iNdEx := len(m.Hosts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Hosts[iNdEx])
			copy(dAtA[i:], m.Hosts[iNdEx])
			i = encodeVarintServiceexpose(dAtA, i, uint64(len(m.Hosts[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Global {
		i--
		if m.Global {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Service) > 0 {
		i -= len(m.Service)
		copy(dAtA[i:], m.Service)
		i = encodeVarintServiceexpose(dAtA, i, uint64(len(m.Service)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Proto) > 0 {
		i -= len(m.Proto)
		copy(dAtA[i:], m.Proto)
		i = encodeVarintServiceexpose(dAtA, i, uint64(len(m.Proto)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ExternalPort != 0 {
		i = encodeVarintServiceexpose(dAtA, i, uint64(m.ExternalPort))
		i--
		dAtA[i] = 0x10
	}
	if m.Port != 0 {
		i = encodeVarintServiceexpose(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintServiceexpose(dAtA []byte, offset int, v uint64) int {
	offset -= sovServiceexpose(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ServiceExpose) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Port != 0 {
		n += 1 + sovServiceexpose(uint64(m.Port))
	}
	if m.ExternalPort != 0 {
		n += 1 + sovServiceexpose(uint64(m.ExternalPort))
	}
	l = len(m.Proto)
	if l > 0 {
		n += 1 + l + sovServiceexpose(uint64(l))
	}
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovServiceexpose(uint64(l))
	}
	if m.Global {
		n += 2
	}
	if len(m.Hosts) > 0 {
		for _, s := range m.Hosts {
			l = len(s)
			n += 1 + l + sovServiceexpose(uint64(l))
		}
	}
	l = m.HTTPOptions.Size()
	n += 1 + l + sovServiceexpose(uint64(l))
	l = len(m.IP)
	if l > 0 {
		n += 1 + l + sovServiceexpose(uint64(l))
	}
	if m.EndpointSequenceNumber != 0 {
		n += 1 + sovServiceexpose(uint64(m.EndpointSequenceNumber))
	}
	return n
}

func sovServiceexpose(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozServiceexpose(x uint64) (n int) {
	return sovServiceexpose(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ServiceExpose) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServiceExpose{`,
		`Port:` + fmt.Sprintf("%v", this.Port) + `,`,
		`ExternalPort:` + fmt.Sprintf("%v", this.ExternalPort) + `,`,
		`Proto:` + fmt.Sprintf("%v", this.Proto) + `,`,
		`Service:` + fmt.Sprintf("%v", this.Service) + `,`,
		`Global:` + fmt.Sprintf("%v", this.Global) + `,`,
		`Hosts:` + fmt.Sprintf("%v", this.Hosts) + `,`,
		`HTTPOptions:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.HTTPOptions), "ServiceExposeHTTPOptions", "ServiceExposeHTTPOptions", 1), `&`, ``, 1) + `,`,
		`IP:` + fmt.Sprintf("%v", this.IP) + `,`,
		`EndpointSequenceNumber:` + fmt.Sprintf("%v", this.EndpointSequenceNumber) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringServiceexpose(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ServiceExpose) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceexpose
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
			return fmt.Errorf("proto: ServiceExpose: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceExpose: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceexpose
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalPort", wireType)
			}
			m.ExternalPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceexpose
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExternalPort |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proto", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceexpose
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
				return ErrInvalidLengthServiceexpose
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceexpose
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proto = ServiceProtocol(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceexpose
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
				return ErrInvalidLengthServiceexpose
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceexpose
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Global", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceexpose
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
			m.Global = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hosts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceexpose
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
				return ErrInvalidLengthServiceexpose
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceexpose
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hosts = append(m.Hosts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HTTPOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceexpose
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
				return ErrInvalidLengthServiceexpose
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServiceexpose
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HTTPOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceexpose
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
				return ErrInvalidLengthServiceexpose
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceexpose
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndpointSequenceNumber", wireType)
			}
			m.EndpointSequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceexpose
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndpointSequenceNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipServiceexpose(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServiceexpose
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
func skipServiceexpose(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServiceexpose
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
					return 0, ErrIntOverflowServiceexpose
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
					return 0, ErrIntOverflowServiceexpose
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
				return 0, ErrInvalidLengthServiceexpose
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupServiceexpose
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthServiceexpose
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthServiceexpose        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServiceexpose          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupServiceexpose = fmt.Errorf("proto: unexpected end of group")
)
