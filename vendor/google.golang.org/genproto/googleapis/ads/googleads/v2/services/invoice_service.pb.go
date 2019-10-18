// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/invoice_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for fetching the invoices of a given billing setup that were
// issued during a given month.
type ListInvoicesRequest struct {
	// The ID of the customer to fetch invoices for.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The billing setup resource name of the requested invoices.
	//
	// `customers/{customer_id}/billingSetups/{billing_setup_id}`
	BillingSetup string `protobuf:"bytes,2,opt,name=billing_setup,json=billingSetup,proto3" json:"billing_setup,omitempty"`
	// Required. The issue year to retrieve invoices, in yyyy format. Only
	// invoices issued in 2019 or later can be retrieved.
	IssueYear string `protobuf:"bytes,3,opt,name=issue_year,json=issueYear,proto3" json:"issue_year,omitempty"`
	// Required. The issue month to retrieve invoices.
	IssueMonth           enums.MonthOfYearEnum_MonthOfYear `protobuf:"varint,4,opt,name=issue_month,json=issueMonth,proto3,enum=google.ads.googleads.v2.enums.MonthOfYearEnum_MonthOfYear" json:"issue_month,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ListInvoicesRequest) Reset()         { *m = ListInvoicesRequest{} }
func (m *ListInvoicesRequest) String() string { return proto.CompactTextString(m) }
func (*ListInvoicesRequest) ProtoMessage()    {}
func (*ListInvoicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f8d00dc161add1, []int{0}
}

func (m *ListInvoicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInvoicesRequest.Unmarshal(m, b)
}
func (m *ListInvoicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInvoicesRequest.Marshal(b, m, deterministic)
}
func (m *ListInvoicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInvoicesRequest.Merge(m, src)
}
func (m *ListInvoicesRequest) XXX_Size() int {
	return xxx_messageInfo_ListInvoicesRequest.Size(m)
}
func (m *ListInvoicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInvoicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListInvoicesRequest proto.InternalMessageInfo

func (m *ListInvoicesRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *ListInvoicesRequest) GetBillingSetup() string {
	if m != nil {
		return m.BillingSetup
	}
	return ""
}

func (m *ListInvoicesRequest) GetIssueYear() string {
	if m != nil {
		return m.IssueYear
	}
	return ""
}

func (m *ListInvoicesRequest) GetIssueMonth() enums.MonthOfYearEnum_MonthOfYear {
	if m != nil {
		return m.IssueMonth
	}
	return enums.MonthOfYearEnum_UNSPECIFIED
}

// Response message for [InvoiceService.ListInvoices][google.ads.googleads.v2.services.InvoiceService.ListInvoices].
type ListInvoicesResponse struct {
	// The list of invoices that match the billing setup and time period.
	Invoices             []*resources.Invoice `protobuf:"bytes,1,rep,name=invoices,proto3" json:"invoices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListInvoicesResponse) Reset()         { *m = ListInvoicesResponse{} }
func (m *ListInvoicesResponse) String() string { return proto.CompactTextString(m) }
func (*ListInvoicesResponse) ProtoMessage()    {}
func (*ListInvoicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f8d00dc161add1, []int{1}
}

func (m *ListInvoicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInvoicesResponse.Unmarshal(m, b)
}
func (m *ListInvoicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInvoicesResponse.Marshal(b, m, deterministic)
}
func (m *ListInvoicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInvoicesResponse.Merge(m, src)
}
func (m *ListInvoicesResponse) XXX_Size() int {
	return xxx_messageInfo_ListInvoicesResponse.Size(m)
}
func (m *ListInvoicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInvoicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListInvoicesResponse proto.InternalMessageInfo

func (m *ListInvoicesResponse) GetInvoices() []*resources.Invoice {
	if m != nil {
		return m.Invoices
	}
	return nil
}

func init() {
	proto.RegisterType((*ListInvoicesRequest)(nil), "google.ads.googleads.v2.services.ListInvoicesRequest")
	proto.RegisterType((*ListInvoicesResponse)(nil), "google.ads.googleads.v2.services.ListInvoicesResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/invoice_service.proto", fileDescriptor_75f8d00dc161add1)
}

var fileDescriptor_75f8d00dc161add1 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0x26, 0x59, 0x11, 0x3b, 0x5b, 0x7b, 0x48, 0x05, 0xc3, 0xaa, 0x18, 0xd6, 0x22, 0x4b, 0x0f,
	0x33, 0x18, 0xb1, 0x87, 0x88, 0x87, 0x14, 0xb4, 0x16, 0x14, 0xcb, 0x16, 0x16, 0xd4, 0xc5, 0x90,
	0x26, 0xd3, 0x38, 0x90, 0xcc, 0xc4, 0xbc, 0x64, 0x41, 0xc4, 0x8b, 0xff, 0x82, 0xff, 0x81, 0x47,
	0xc1, 0x7f, 0xa4, 0x57, 0x0f, 0x9e, 0xbc, 0x79, 0xf2, 0x4f, 0xf0, 0x24, 0x99, 0x1f, 0x69, 0x0a,
	0x86, 0xc5, 0xdb, 0xcb, 0xf7, 0xbe, 0xef, 0x7b, 0x33, 0xdf, 0xbc, 0xa0, 0xbd, 0x4c, 0x88, 0x2c,
	0xa7, 0x24, 0x4e, 0x81, 0xa8, 0xb2, 0xad, 0x56, 0x3e, 0x01, 0x5a, 0xad, 0x58, 0x42, 0x81, 0x30,
	0xbe, 0x12, 0x2c, 0xa1, 0x91, 0x06, 0x70, 0x59, 0x89, 0x5a, 0x38, 0x9e, 0x22, 0xe3, 0x38, 0x05,
	0xdc, 0xe9, 0xf0, 0xca, 0xc7, 0x46, 0x37, 0xb9, 0x37, 0xe4, 0x4c, 0x79, 0x53, 0x00, 0x29, 0x04,
	0xaf, 0xdf, 0x46, 0xe2, 0x34, 0x7a, 0x4f, 0xe3, 0x4a, 0x99, 0x4e, 0xc8, 0x90, 0xa4, 0xa2, 0x20,
	0x9a, 0xaa, 0x77, 0x1a, 0x2d, 0xb8, 0x69, 0x04, 0x25, 0x23, 0x31, 0xe7, 0xa2, 0x8e, 0x6b, 0x26,
	0x38, 0xe8, 0xee, 0xf5, 0x5e, 0x37, 0xc9, 0x19, 0xe5, 0xb5, 0x6a, 0x4c, 0x7f, 0x58, 0x68, 0xfb,
	0x19, 0x83, 0xfa, 0x50, 0x99, 0xc1, 0x9c, 0xbe, 0x6b, 0x28, 0xd4, 0xce, 0x6d, 0x34, 0x4e, 0x1a,
	0xa8, 0x45, 0x41, 0xab, 0x88, 0xa5, 0xae, 0xe5, 0x59, 0xb3, 0x8d, 0x39, 0x32, 0xd0, 0x61, 0xea,
	0xdc, 0x41, 0x57, 0x4f, 0x58, 0x9e, 0x33, 0x9e, 0x45, 0x40, 0xeb, 0xa6, 0x74, 0x6d, 0x49, 0xd9,
	0xd4, 0xe0, 0x71, 0x8b, 0x39, 0xb7, 0x10, 0x62, 0x00, 0x0d, 0x95, 0x37, 0x73, 0x47, 0x92, 0xb1,
	0x21, 0x91, 0x97, 0x34, 0xae, 0x9c, 0xd7, 0x68, 0xac, 0xda, 0x32, 0x01, 0xf7, 0x92, 0x67, 0xcd,
	0xb6, 0xfc, 0x00, 0x0f, 0xe5, 0x29, 0xd3, 0xc2, 0xcf, 0x5b, 0xee, 0x8b, 0xd3, 0xd6, 0xe0, 0x31,
	0x6f, 0x8a, 0xfe, 0xf7, 0x5c, 0x4d, 0x93, 0xc8, 0xf4, 0x0d, 0xba, 0x76, 0xf1, 0x62, 0x50, 0x0a,
	0x0e, 0xd4, 0x79, 0x82, 0xae, 0xe8, 0xe4, 0xc0, 0xb5, 0xbc, 0xd1, 0x6c, 0xec, 0xef, 0x0e, 0x4e,
	0xec, 0xc2, 0xc6, 0xda, 0x66, 0xde, 0x69, 0xfd, 0x9f, 0x16, 0xda, 0xd2, 0xe8, 0xb1, 0x7a, 0x68,
	0xe7, 0x9b, 0x85, 0x36, 0xfb, 0x33, 0x9d, 0x07, 0x78, 0xdd, 0x6e, 0xe0, 0x7f, 0x84, 0x3f, 0xd9,
	0xfb, 0x5f, 0x99, 0xba, 0xda, 0x14, 0x7f, 0xfa, 0xfe, 0xeb, 0xb3, 0x3d, 0x73, 0xee, 0xb6, 0x8b,
	0x62, 0xde, 0x0a, 0xc8, 0x87, 0xde, 0x4b, 0x3e, 0xda, 0xfd, 0x68, 0x16, 0x07, 0x26, 0x37, 0xce,
	0x42, 0xf7, 0xdc, 0x5e, 0x57, 0x25, 0x03, 0x9c, 0x88, 0x62, 0xff, 0x8f, 0x85, 0x76, 0x12, 0x51,
	0xac, 0x3d, 0xca, 0xfe, 0xf6, 0xc5, 0x14, 0x8e, 0xda, 0xbd, 0x3a, 0xb2, 0x5e, 0x3d, 0xd5, 0xc2,
	0x4c, 0xe4, 0x31, 0xcf, 0xb0, 0xa8, 0x32, 0x92, 0x51, 0x2e, 0xb7, 0x8e, 0x9c, 0x8f, 0x1a, 0xfe,
	0xdb, 0x1e, 0x9a, 0xe2, 0x8b, 0x3d, 0x3a, 0x08, 0xc3, 0xaf, 0xb6, 0x77, 0xa0, 0x0c, 0xc3, 0x14,
	0xb0, 0x2a, 0xdb, 0x6a, 0xe1, 0x63, 0x3d, 0x18, 0xce, 0x0c, 0x65, 0x19, 0xa6, 0xb0, 0xec, 0x28,
	0xcb, 0x85, 0xbf, 0x34, 0x94, 0xdf, 0xf6, 0x8e, 0xc2, 0x83, 0x20, 0x4c, 0x21, 0x08, 0x3a, 0x52,
	0x10, 0x2c, 0xfc, 0x20, 0x30, 0xb4, 0x93, 0xcb, 0xf2, 0x9c, 0xf7, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0x87, 0x53, 0xe2, 0x14, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InvoiceServiceClient is the client API for InvoiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InvoiceServiceClient interface {
	// Returns all invoices associated with a billing setup, for a given month.
	ListInvoices(ctx context.Context, in *ListInvoicesRequest, opts ...grpc.CallOption) (*ListInvoicesResponse, error)
}

type invoiceServiceClient struct {
	cc *grpc.ClientConn
}

func NewInvoiceServiceClient(cc *grpc.ClientConn) InvoiceServiceClient {
	return &invoiceServiceClient{cc}
}

func (c *invoiceServiceClient) ListInvoices(ctx context.Context, in *ListInvoicesRequest, opts ...grpc.CallOption) (*ListInvoicesResponse, error) {
	out := new(ListInvoicesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.InvoiceService/ListInvoices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServiceServer is the server API for InvoiceService service.
type InvoiceServiceServer interface {
	// Returns all invoices associated with a billing setup, for a given month.
	ListInvoices(context.Context, *ListInvoicesRequest) (*ListInvoicesResponse, error)
}

// UnimplementedInvoiceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInvoiceServiceServer struct {
}

func (*UnimplementedInvoiceServiceServer) ListInvoices(ctx context.Context, req *ListInvoicesRequest) (*ListInvoicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInvoices not implemented")
}

func RegisterInvoiceServiceServer(s *grpc.Server, srv InvoiceServiceServer) {
	s.RegisterService(&_InvoiceService_serviceDesc, srv)
}

func _InvoiceService_ListInvoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInvoicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).ListInvoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.InvoiceService/ListInvoices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).ListInvoices(ctx, req.(*ListInvoicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InvoiceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.InvoiceService",
	HandlerType: (*InvoiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInvoices",
			Handler:    _InvoiceService_ListInvoices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/invoice_service.proto",
}
