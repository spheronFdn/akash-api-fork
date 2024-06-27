// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1beta3/order.proto

package v1beta3

import (
	fmt "fmt"
	v1beta3 "github.com/spheronFdn/akash-api-fork/go/node/deployment/v1beta3"
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

// State is an enum which refers to state of order
type Order_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	OrderStateInvalid Order_State = 0
	// OrderOpen denotes state for order open
	OrderOpen Order_State = 1
	// OrderMatched denotes state for order matched
	OrderActive Order_State = 2
	// OrderClosed denotes state for order lost
	OrderClosed Order_State = 3
)

var Order_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "active",
	3: "closed",
}

var Order_State_value = map[string]int32{
	"invalid": 0,
	"open":    1,
	"active":  2,
	"closed":  3,
}

func (x Order_State) String() string {
	return proto.EnumName(Order_State_name, int32(x))
}

func (Order_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cffb7c6e7c4cdcbf, []int{1, 0}
}

// OrderID stores owner and all other seq numbers
type OrderID struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq  uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq  uint32 `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	OSeq  uint32 `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
}

func (m *OrderID) Reset()      { *m = OrderID{} }
func (*OrderID) ProtoMessage() {}
func (*OrderID) Descriptor() ([]byte, []int) {
	return fileDescriptor_cffb7c6e7c4cdcbf, []int{0}
}
func (m *OrderID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderID.Merge(m, src)
}
func (m *OrderID) XXX_Size() int {
	return m.Size()
}
func (m *OrderID) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderID.DiscardUnknown(m)
}

var xxx_messageInfo_OrderID proto.InternalMessageInfo

func (m *OrderID) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *OrderID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *OrderID) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *OrderID) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

// Order stores orderID, state of order and other details
type Order struct {
	OrderID   OrderID           `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"id" yaml:"id"`
	State     Order_State       `protobuf:"varint,2,opt,name=state,proto3,enum=akash.market.v1beta3.Order_State" json:"state" yaml:"state"`
	Spec      v1beta3.GroupSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec" yaml:"spec"`
	CreatedAt int64             `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *Order) Reset()      { *m = Order{} }
func (*Order) ProtoMessage() {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cffb7c6e7c4cdcbf, []int{1}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderID() OrderID {
	if m != nil {
		return m.OrderID
	}
	return OrderID{}
}

func (m *Order) GetState() Order_State {
	if m != nil {
		return m.State
	}
	return OrderStateInvalid
}

func (m *Order) GetSpec() v1beta3.GroupSpec {
	if m != nil {
		return m.Spec
	}
	return v1beta3.GroupSpec{}
}

func (m *Order) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

// OrderFilters defines flags for order list filter
type OrderFilters struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq  uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq  uint32 `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	OSeq  uint32 `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
	State string `protobuf:"bytes,5,opt,name=state,proto3" json:"state" yaml:"state"`
}

func (m *OrderFilters) Reset()         { *m = OrderFilters{} }
func (m *OrderFilters) String() string { return proto.CompactTextString(m) }
func (*OrderFilters) ProtoMessage()    {}
func (*OrderFilters) Descriptor() ([]byte, []int) {
	return fileDescriptor_cffb7c6e7c4cdcbf, []int{2}
}
func (m *OrderFilters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderFilters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderFilters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderFilters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFilters.Merge(m, src)
}
func (m *OrderFilters) XXX_Size() int {
	return m.Size()
}
func (m *OrderFilters) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFilters.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFilters proto.InternalMessageInfo

func (m *OrderFilters) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *OrderFilters) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *OrderFilters) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *OrderFilters) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

func (m *OrderFilters) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterEnum("akash.market.v1beta3.Order_State", Order_State_name, Order_State_value)
	proto.RegisterType((*OrderID)(nil), "akash.market.v1beta3.OrderID")
	proto.RegisterType((*Order)(nil), "akash.market.v1beta3.Order")
	proto.RegisterType((*OrderFilters)(nil), "akash.market.v1beta3.OrderFilters")
}

func init() { proto.RegisterFile("akash/market/v1beta3/order.proto", fileDescriptor_cffb7c6e7c4cdcbf) }

var fileDescriptor_cffb7c6e7c4cdcbf = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0x13, 0xa7, 0x6d, 0x2e, 0x2d, 0x04, 0xab, 0x88, 0xe2, 0xaa, 0x3e, 0x63, 0x96, 0x2c,
	0xd8, 0xa2, 0x9d, 0xc8, 0xd6, 0x50, 0x51, 0x65, 0x8a, 0xe4, 0x30, 0x21, 0xa4, 0xca, 0xf1, 0x9d,
	0x5c, 0x2b, 0x89, 0xcf, 0xb5, 0xaf, 0xa9, 0xba, 0x33, 0xa0, 0x4c, 0x2c, 0x48, 0x2c, 0x91, 0x2a,
	0xf1, 0x43, 0x58, 0x3b, 0x76, 0x64, 0xb2, 0x50, 0xb2, 0xa0, 0x8c, 0xf9, 0x05, 0xe8, 0xde, 0x19,
	0xdc, 0x22, 0xd4, 0x1f, 0xc0, 0x64, 0xbf, 0xef, 0x7d, 0xdf, 0xbb, 0x7b, 0xdf, 0x7b, 0x3a, 0x64,
	0xf9, 0x43, 0x3f, 0x3b, 0x75, 0xc7, 0x7e, 0x3a, 0xa4, 0xdc, 0x9d, 0xbc, 0x1c, 0x50, 0xee, 0x1f,
	0xb8, 0x2c, 0x25, 0x34, 0x75, 0x92, 0x94, 0x71, 0xa6, 0x6f, 0x03, 0xc3, 0x91, 0x0c, 0xa7, 0x60,
	0x18, 0xdb, 0x21, 0x0b, 0x19, 0x10, 0x5c, 0xf1, 0x27, 0xb9, 0x46, 0x4b, 0x56, 0x23, 0x34, 0x19,
	0xb1, 0xcb, 0x31, 0x8d, 0xcb, 0x8a, 0x61, 0xca, 0xce, 0x93, 0x2c, 0xa1, 0x81, 0x64, 0xda, 0x0b,
	0x15, 0xad, 0xf7, 0xc4, 0x29, 0xdd, 0x23, 0xdd, 0x45, 0x35, 0x76, 0x11, 0xd3, 0x74, 0x47, 0xb5,
	0xd4, 0x56, 0xbd, 0xf3, 0x74, 0x99, 0x63, 0x09, 0xac, 0x72, 0xbc, 0x79, 0xe9, 0x8f, 0x47, 0x6d,
	0x1b, 0x42, 0xdb, 0x93, 0xb0, 0x7e, 0x80, 0x34, 0x92, 0xd1, 0xb3, 0x9d, 0x8a, 0xa5, 0xb6, 0xb4,
	0x0e, 0x9e, 0xe7, 0x58, 0x3b, 0xea, 0xd3, 0xb3, 0x65, 0x8e, 0x01, 0x5f, 0xe5, 0xb8, 0x21, 0x65,
	0x22, 0xb2, 0x3d, 0x00, 0x85, 0x28, 0x14, 0xa2, 0xaa, 0xa5, 0xb6, 0xb6, 0xa4, 0xe8, 0xb8, 0x10,
	0x85, 0x77, 0x44, 0xa1, 0x14, 0x85, 0x85, 0x88, 0x09, 0x91, 0x56, 0x8a, 0x7a, 0x85, 0x88, 0xdd,
	0x11, 0x31, 0x29, 0x12, 0x9f, 0xf6, 0xc6, 0x97, 0x2b, 0xac, 0xfc, 0xbc, 0xc2, 0x8a, 0xfd, 0xad,
	0x8a, 0x6a, 0xd0, 0xa5, 0xfe, 0x1e, 0x6d, 0x80, 0xa9, 0x27, 0x11, 0x81, 0x36, 0x1b, 0xfb, 0x7b,
	0xce, 0xbf, 0x8c, 0x75, 0x0a, 0x53, 0x3a, 0xf6, 0x75, 0x8e, 0x95, 0x79, 0x8e, 0x7f, 0xbb, 0xb4,
	0xcc, 0x71, 0x25, 0x22, 0xab, 0x1c, 0xd7, 0xe5, 0x81, 0x11, 0xb1, 0xbd, 0x75, 0x28, 0xd9, 0x25,
	0xba, 0x87, 0x6a, 0x19, 0xf7, 0x39, 0x05, 0x47, 0x1e, 0xec, 0x3f, 0xbb, 0xa7, 0xb4, 0xd3, 0x17,
	0x44, 0x69, 0x32, 0x68, 0x4a, 0x93, 0x21, 0xb4, 0x3d, 0x09, 0xeb, 0x6f, 0x91, 0x26, 0xe6, 0x05,
	0x7e, 0x35, 0xf6, 0x9f, 0x17, 0x25, 0xcb, 0xd1, 0xfe, 0x29, 0x7b, 0x2c, 0x46, 0xdb, 0x4f, 0x68,
	0xd0, 0xd9, 0x15, 0x77, 0x16, 0xde, 0x08, 0x61, 0xe9, 0x8d, 0x88, 0x6c, 0x0f, 0x40, 0x7d, 0x0f,
	0xa1, 0x20, 0xa5, 0x3e, 0xa7, 0xe4, 0xc4, 0xe7, 0x60, 0x6b, 0xd5, 0xab, 0x17, 0xc8, 0x21, 0xb7,
	0x3f, 0xa8, 0xa8, 0x06, 0x17, 0xd4, 0x6d, 0xb4, 0x1e, 0xc5, 0x13, 0x7f, 0x14, 0x91, 0xa6, 0x62,
	0x3c, 0x9e, 0xce, 0xac, 0x47, 0x70, 0x7d, 0x48, 0x76, 0x65, 0x42, 0x7f, 0x82, 0x34, 0x96, 0xd0,
	0xb8, 0xa9, 0x1a, 0x5b, 0xd3, 0x99, 0x55, 0x07, 0x42, 0x2f, 0xa1, 0xb1, 0xbe, 0x8b, 0xd6, 0xfc,
	0x80, 0x47, 0x13, 0xda, 0xac, 0x18, 0x0f, 0xa7, 0x33, 0xab, 0x01, 0xa9, 0x43, 0x80, 0x44, 0x32,
	0x18, 0xb1, 0x8c, 0x92, 0x66, 0xf5, 0x56, 0xf2, 0x35, 0x40, 0x86, 0xf6, 0xf1, 0xab, 0xa9, 0xdc,
	0x9a, 0xe0, 0xe7, 0x0a, 0xda, 0x84, 0xfc, 0x9b, 0x68, 0xc4, 0x69, 0x9a, 0xfd, 0x6f, 0xcb, 0x2a,
	0xfa, 0x91, 0xab, 0x53, 0x2b, 0xfb, 0xb9, 0x6f, 0x2f, 0xda, 0x9a, 0xf0, 0xa5, 0xd3, 0xbf, 0x9e,
	0x9b, 0xea, 0xcd, 0xdc, 0x54, 0x7f, 0xcc, 0x4d, 0xf5, 0xd3, 0xc2, 0x54, 0x6e, 0x16, 0xa6, 0xf2,
	0x7d, 0x61, 0x2a, 0xef, 0x5e, 0x85, 0x11, 0x3f, 0x3d, 0x1f, 0x38, 0x01, 0x1b, 0xbb, 0xb0, 0x33,
	0x2f, 0x62, 0xca, 0x2f, 0x58, 0x3a, 0x2c, 0x22, 0x3f, 0x89, 0xdc, 0x90, 0xb9, 0x31, 0x23, 0xf4,
	0xaf, 0x67, 0x67, 0xb0, 0x06, 0x6f, 0xc3, 0xc1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x36,
	0xa2, 0x46, 0x95, 0x04, 0x00, 0x00,
}

func (m *OrderID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.OrderID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *OrderFilters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderFilters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderFilters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x2a
	}
	if m.OSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovOrder(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovOrder(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovOrder(uint64(m.OSeq))
	}
	return n
}

func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OrderID.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.State != 0 {
		n += 1 + sovOrder(uint64(m.State))
	}
	l = m.Spec.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.CreatedAt != 0 {
		n += 1 + sovOrder(uint64(m.CreatedAt))
	}
	return n
}

func (m *OrderFilters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovOrder(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovOrder(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovOrder(uint64(m.OSeq))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: OrderID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Order_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowOrder
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
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *OrderFilters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: OrderFilters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderFilters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)
