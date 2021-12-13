// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v2/hds.proto

package envoy_service_discovery_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Capability_Protocol int32

const (
	Capability_HTTP  Capability_Protocol = 0
	Capability_TCP   Capability_Protocol = 1
	Capability_REDIS Capability_Protocol = 2
)

var Capability_Protocol_name = map[int32]string{
	0: "HTTP",
	1: "TCP",
	2: "REDIS",
}

var Capability_Protocol_value = map[string]int32{
	"HTTP":  0,
	"TCP":   1,
	"REDIS": 2,
}

func (x Capability_Protocol) String() string {
	return proto.EnumName(Capability_Protocol_name, int32(x))
}

func (Capability_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_773ad67555497672, []int{0, 0}
}

type Capability struct {
	HealthCheckProtocols []Capability_Protocol `protobuf:"varint,1,rep,packed,name=health_check_protocols,json=healthCheckProtocols,proto3,enum=envoy.service.discovery.v2.Capability_Protocol" json:"health_check_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Capability) Reset()         { *m = Capability{} }
func (m *Capability) String() string { return proto.CompactTextString(m) }
func (*Capability) ProtoMessage()    {}
func (*Capability) Descriptor() ([]byte, []int) {
	return fileDescriptor_773ad67555497672, []int{0}
}

func (m *Capability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capability.Unmarshal(m, b)
}
func (m *Capability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capability.Marshal(b, m, deterministic)
}
func (m *Capability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capability.Merge(m, src)
}
func (m *Capability) XXX_Size() int {
	return xxx_messageInfo_Capability.Size(m)
}
func (m *Capability) XXX_DiscardUnknown() {
	xxx_messageInfo_Capability.DiscardUnknown(m)
}

var xxx_messageInfo_Capability proto.InternalMessageInfo

func (m *Capability) GetHealthCheckProtocols() []Capability_Protocol {
	if m != nil {
		return m.HealthCheckProtocols
	}
	return nil
}

type HealthCheckRequest struct {
	Node                 *core.Node  `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Capability           *Capability `protobuf:"bytes,2,opt,name=capability,proto3" json:"capability,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_773ad67555497672, []int{1}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *HealthCheckRequest) GetCapability() *Capability {
	if m != nil {
		return m.Capability
	}
	return nil
}

type EndpointHealth struct {
	Endpoint             *endpoint.Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	HealthStatus         core.HealthStatus  `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"health_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EndpointHealth) Reset()         { *m = EndpointHealth{} }
func (m *EndpointHealth) String() string { return proto.CompactTextString(m) }
func (*EndpointHealth) ProtoMessage()    {}
func (*EndpointHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_773ad67555497672, []int{2}
}

func (m *EndpointHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointHealth.Unmarshal(m, b)
}
func (m *EndpointHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointHealth.Marshal(b, m, deterministic)
}
func (m *EndpointHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointHealth.Merge(m, src)
}
func (m *EndpointHealth) XXX_Size() int {
	return xxx_messageInfo_EndpointHealth.Size(m)
}
func (m *EndpointHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointHealth.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointHealth proto.InternalMessageInfo

func (m *EndpointHealth) GetEndpoint() *endpoint.Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *EndpointHealth) GetHealthStatus() core.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

type EndpointHealthResponse struct {
	EndpointsHealth      []*EndpointHealth `protobuf:"bytes,1,rep,name=endpoints_health,json=endpointsHealth,proto3" json:"endpoints_health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EndpointHealthResponse) Reset()         { *m = EndpointHealthResponse{} }
func (m *EndpointHealthResponse) String() string { return proto.CompactTextString(m) }
func (*EndpointHealthResponse) ProtoMessage()    {}
func (*EndpointHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_773ad67555497672, []int{3}
}

func (m *EndpointHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointHealthResponse.Unmarshal(m, b)
}
func (m *EndpointHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointHealthResponse.Marshal(b, m, deterministic)
}
func (m *EndpointHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointHealthResponse.Merge(m, src)
}
func (m *EndpointHealthResponse) XXX_Size() int {
	return xxx_messageInfo_EndpointHealthResponse.Size(m)
}
func (m *EndpointHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointHealthResponse proto.InternalMessageInfo

func (m *EndpointHealthResponse) GetEndpointsHealth() []*EndpointHealth {
	if m != nil {
		return m.EndpointsHealth
	}
	return nil
}

type HealthCheckRequestOrEndpointHealthResponse struct {
	// Types that are valid to be assigned to RequestType:
	//	*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest
	//	*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse
	RequestType          isHealthCheckRequestOrEndpointHealthResponse_RequestType `protobuf_oneof:"request_type"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *HealthCheckRequestOrEndpointHealthResponse) Reset() {
	*m = HealthCheckRequestOrEndpointHealthResponse{}
}
func (m *HealthCheckRequestOrEndpointHealthResponse) String() string {
	return proto.CompactTextString(m)
}
func (*HealthCheckRequestOrEndpointHealthResponse) ProtoMessage() {}
func (*HealthCheckRequestOrEndpointHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_773ad67555497672, []int{4}
}

func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Unmarshal(m, b)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Merge(m, src)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Size(m)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse proto.InternalMessageInfo

type isHealthCheckRequestOrEndpointHealthResponse_RequestType interface {
	isHealthCheckRequestOrEndpointHealthResponse_RequestType()
}

type HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest struct {
	HealthCheckRequest *HealthCheckRequest `protobuf:"bytes,1,opt,name=health_check_request,json=healthCheckRequest,proto3,oneof"`
}

type HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse struct {
	EndpointHealthResponse *EndpointHealthResponse `protobuf:"bytes,2,opt,name=endpoint_health_response,json=endpointHealthResponse,proto3,oneof"`
}

func (*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}

func (*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetRequestType() isHealthCheckRequestOrEndpointHealthResponse_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetHealthCheckRequest() *HealthCheckRequest {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest); ok {
		return x.HealthCheckRequest
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetEndpointHealthResponse() *EndpointHealthResponse {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse); ok {
		return x.EndpointHealthResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheckRequestOrEndpointHealthResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest)(nil),
		(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse)(nil),
	}
}

type LocalityEndpoints struct {
	Locality             *core.Locality       `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	Endpoints            []*endpoint.Endpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LocalityEndpoints) Reset()         { *m = LocalityEndpoints{} }
func (m *LocalityEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityEndpoints) ProtoMessage()    {}
func (*LocalityEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_773ad67555497672, []int{5}
}

func (m *LocalityEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityEndpoints.Unmarshal(m, b)
}
func (m *LocalityEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityEndpoints.Marshal(b, m, deterministic)
}
func (m *LocalityEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityEndpoints.Merge(m, src)
}
func (m *LocalityEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityEndpoints.Size(m)
}
func (m *LocalityEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityEndpoints proto.InternalMessageInfo

func (m *LocalityEndpoints) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityEndpoints) GetEndpoints() []*endpoint.Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type ClusterHealthCheck struct {
	ClusterName          string               `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	HealthChecks         []*core.HealthCheck  `protobuf:"bytes,2,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	LocalityEndpoints    []*LocalityEndpoints `protobuf:"bytes,3,rep,name=locality_endpoints,json=localityEndpoints,proto3" json:"locality_endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClusterHealthCheck) Reset()         { *m = ClusterHealthCheck{} }
func (m *ClusterHealthCheck) String() string { return proto.CompactTextString(m) }
func (*ClusterHealthCheck) ProtoMessage()    {}
func (*ClusterHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_773ad67555497672, []int{6}
}

func (m *ClusterHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterHealthCheck.Unmarshal(m, b)
}
func (m *ClusterHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterHealthCheck.Marshal(b, m, deterministic)
}
func (m *ClusterHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterHealthCheck.Merge(m, src)
}
func (m *ClusterHealthCheck) XXX_Size() int {
	return xxx_messageInfo_ClusterHealthCheck.Size(m)
}
func (m *ClusterHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterHealthCheck proto.InternalMessageInfo

func (m *ClusterHealthCheck) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterHealthCheck) GetHealthChecks() []*core.HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func (m *ClusterHealthCheck) GetLocalityEndpoints() []*LocalityEndpoints {
	if m != nil {
		return m.LocalityEndpoints
	}
	return nil
}

type HealthCheckSpecifier struct {
	ClusterHealthChecks  []*ClusterHealthCheck `protobuf:"bytes,1,rep,name=cluster_health_checks,json=clusterHealthChecks,proto3" json:"cluster_health_checks,omitempty"`
	Interval             *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HealthCheckSpecifier) Reset()         { *m = HealthCheckSpecifier{} }
func (m *HealthCheckSpecifier) String() string { return proto.CompactTextString(m) }
func (*HealthCheckSpecifier) ProtoMessage()    {}
func (*HealthCheckSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_773ad67555497672, []int{7}
}

func (m *HealthCheckSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckSpecifier.Unmarshal(m, b)
}
func (m *HealthCheckSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckSpecifier.Marshal(b, m, deterministic)
}
func (m *HealthCheckSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckSpecifier.Merge(m, src)
}
func (m *HealthCheckSpecifier) XXX_Size() int {
	return xxx_messageInfo_HealthCheckSpecifier.Size(m)
}
func (m *HealthCheckSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckSpecifier proto.InternalMessageInfo

func (m *HealthCheckSpecifier) GetClusterHealthChecks() []*ClusterHealthCheck {
	if m != nil {
		return m.ClusterHealthChecks
	}
	return nil
}

func (m *HealthCheckSpecifier) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.service.discovery.v2.Capability_Protocol", Capability_Protocol_name, Capability_Protocol_value)
	proto.RegisterType((*Capability)(nil), "envoy.service.discovery.v2.Capability")
	proto.RegisterType((*HealthCheckRequest)(nil), "envoy.service.discovery.v2.HealthCheckRequest")
	proto.RegisterType((*EndpointHealth)(nil), "envoy.service.discovery.v2.EndpointHealth")
	proto.RegisterType((*EndpointHealthResponse)(nil), "envoy.service.discovery.v2.EndpointHealthResponse")
	proto.RegisterType((*HealthCheckRequestOrEndpointHealthResponse)(nil), "envoy.service.discovery.v2.HealthCheckRequestOrEndpointHealthResponse")
	proto.RegisterType((*LocalityEndpoints)(nil), "envoy.service.discovery.v2.LocalityEndpoints")
	proto.RegisterType((*ClusterHealthCheck)(nil), "envoy.service.discovery.v2.ClusterHealthCheck")
	proto.RegisterType((*HealthCheckSpecifier)(nil), "envoy.service.discovery.v2.HealthCheckSpecifier")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/hds.proto", fileDescriptor_773ad67555497672)
}

var fileDescriptor_773ad67555497672 = []byte{
	// 814 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0xce, 0x24, 0xbb, 0x4b, 0xfa, 0x1a, 0x42, 0x3a, 0x94, 0x6c, 0x36, 0x94, 0x65, 0xb1, 0x56,
	0x28, 0x6a, 0x85, 0x5d, 0xa5, 0x42, 0x48, 0x45, 0x5c, 0x9a, 0xb4, 0x0a, 0x12, 0x2a, 0x91, 0x53,
	0x6e, 0x48, 0x61, 0x62, 0x4f, 0x1b, 0x0b, 0xc7, 0x63, 0x3c, 0x93, 0x88, 0xdc, 0x10, 0x27, 0x10,
	0x07, 0x90, 0x7a, 0xe1, 0xc2, 0x1f, 0xc1, 0x11, 0xfe, 0x02, 0xae, 0x88, 0x1b, 0x67, 0x4e, 0x1c,
	0xf9, 0x03, 0x10, 0xca, 0xfc, 0x70, 0x9c, 0xba, 0x09, 0xed, 0x6d, 0x6f, 0xf1, 0x7b, 0xdf, 0xfb,
	0xde, 0xf7, 0xde, 0xf7, 0xec, 0xc0, 0x73, 0x1a, 0xcd, 0xd8, 0xdc, 0xe1, 0x34, 0x99, 0x05, 0x1e,
	0x75, 0xfc, 0x80, 0x7b, 0x6c, 0x46, 0x93, 0xb9, 0x33, 0x6b, 0x3b, 0x63, 0x9f, 0xdb, 0x71, 0xc2,
	0x04, 0xc3, 0x4d, 0x89, 0xb2, 0x35, 0xca, 0x4e, 0x51, 0xf6, 0xac, 0xdd, 0xdc, 0x53, 0x0c, 0x24,
	0x0e, 0x16, 0x35, 0x1e, 0x4b, 0xa8, 0x33, 0x22, 0x9c, 0xaa, 0xca, 0xe6, 0xf3, 0x7c, 0x76, 0x4c,
	0x49, 0x28, 0xc6, 0x43, 0x6f, 0x4c, 0xbd, 0xcf, 0x35, 0xca, 0x59, 0x41, 0xd1, 0xc8, 0x8f, 0x59,
	0x10, 0x89, 0xf4, 0xc7, 0xd0, 0x63, 0x93, 0x98, 0x45, 0x34, 0x12, 0x5a, 0x50, 0x73, 0xef, 0x8a,
	0xb1, 0xab, 0x90, 0xca, 0x0a, 0x12, 0x45, 0x4c, 0x10, 0x11, 0xb0, 0xc8, 0x64, 0x9f, 0xea, 0xac,
	0x7c, 0x1a, 0x4d, 0x2f, 0x1d, 0x7f, 0x9a, 0x48, 0x80, 0xc9, 0x4f, 0xfd, 0x98, 0x64, 0xeb, 0x9c,
	0x49, 0x70, 0x95, 0x10, 0x61, 0x44, 0xbf, 0x91, 0xcb, 0x73, 0x41, 0xc4, 0x54, 0xd3, 0x5b, 0x3f,
	0x21, 0x80, 0x0e, 0x89, 0xc9, 0x28, 0x08, 0x03, 0x31, 0xc7, 0x14, 0xea, 0xd9, 0x91, 0x86, 0x12,
	0xe4, 0xb1, 0x90, 0x37, 0xd0, 0xb3, 0x52, 0xab, 0xda, 0x76, 0xec, 0xf5, 0xdb, 0xb3, 0x97, 0x3c,
	0x76, 0x5f, 0xd7, 0xb9, 0xbb, 0x8a, 0xae, 0xb3, 0x60, 0x33, 0x41, 0x6e, 0xb5, 0xa0, 0x6c, 0x1e,
	0x70, 0x19, 0x1e, 0xf4, 0x2e, 0x2e, 0xfa, 0xb5, 0x02, 0x7e, 0x09, 0x4a, 0x17, 0x9d, 0x7e, 0x0d,
	0xe1, 0x2d, 0x78, 0xe8, 0x9e, 0x76, 0x3f, 0x1c, 0xd4, 0x8a, 0xd6, 0xb7, 0x08, 0x70, 0x6f, 0x49,
	0xe1, 0xd2, 0x2f, 0xa6, 0x94, 0x0b, 0x7c, 0x00, 0x0f, 0x22, 0xe6, 0xd3, 0x06, 0x7a, 0x86, 0x5a,
	0xdb, 0xed, 0xc7, 0x5a, 0x15, 0x89, 0x83, 0x85, 0x8e, 0x85, 0x33, 0xf6, 0x39, 0xf3, 0xa9, 0x2b,
	0x41, 0xf8, 0x0c, 0xc0, 0x4b, 0xa5, 0x35, 0x8a, 0xb2, 0xe4, 0xed, 0xbb, 0x0d, 0xe2, 0x66, 0x2a,
	0xad, 0x6b, 0x04, 0xd5, 0x53, 0x6d, 0xa3, 0xd2, 0x84, 0xdf, 0x87, 0xb2, 0x31, 0x56, 0x6b, 0x79,
	0x73, 0x55, 0x8b, 0xc9, 0xda, 0xa6, 0xd0, 0x4d, 0x0b, 0x70, 0x17, 0x5e, 0xd6, 0xcb, 0x56, 0x96,
	0x48, 0x69, 0xd5, 0x9b, 0x0c, 0x72, 0x1a, 0xd5, 0x6e, 0x20, 0x61, 0x6e, 0x65, 0x9c, 0x79, 0xb2,
	0x18, 0xd4, 0x57, 0x45, 0xb9, 0x94, 0xc7, 0x2c, 0xe2, 0x14, 0x7f, 0x02, 0x35, 0xd3, 0x8b, 0x0f,
	0x55, 0x8d, 0xb4, 0x71, 0xbb, 0xbd, 0xbf, 0x69, 0xfa, 0x1b, 0x6c, 0xaf, 0xa4, 0x1c, 0x2a, 0x60,
	0xfd, 0x50, 0x84, 0xfd, 0xbc, 0x25, 0x1f, 0x27, 0x6b, 0x54, 0x8c, 0x60, 0x77, 0xe5, 0xa4, 0x12,
	0x85, 0xd7, 0xeb, 0xb2, 0x37, 0x29, 0xc9, 0x77, 0xe9, 0x15, 0x5c, 0x3c, 0xce, 0x9f, 0x43, 0x04,
	0x8d, 0xf4, 0xfd, 0xd2, 0xcd, 0x12, 0xdd, 0x5f, 0xfb, 0xdd, 0xbe, 0xc7, 0xc4, 0xba, 0xb2, 0x57,
	0x70, 0xeb, 0xf4, 0xd6, 0xcc, 0x49, 0x15, 0x2a, 0x7a, 0x8c, 0xa1, 0x98, 0xc7, 0xd4, 0xfa, 0x0e,
	0xc1, 0xce, 0x47, 0xcc, 0x23, 0x8b, 0x33, 0x31, 0x64, 0x1c, 0xbf, 0x07, 0xe5, 0x50, 0x07, 0xf5,
	0xb4, 0xaf, 0xdf, 0x62, 0xad, 0xa9, 0x73, 0x53, 0x30, 0xfe, 0x00, 0xb6, 0xd2, 0xa5, 0x37, 0x8a,
	0xd2, 0xb1, 0xff, 0x3d, 0xab, 0x65, 0x85, 0xf5, 0x27, 0x02, 0xdc, 0x09, 0xa7, 0x5c, 0xd0, 0x24,
	0xb3, 0x41, 0xfc, 0x16, 0x54, 0x3c, 0x15, 0x1d, 0x46, 0x64, 0xa2, 0xde, 0x9d, 0x2d, 0x77, 0x5b,
	0xc7, 0xce, 0xc9, 0x84, 0xe2, 0x4e, 0x7a, 0x91, 0xd2, 0x2b, 0xd3, 0xfc, 0xe9, 0xda, 0x8b, 0x54,
	0x2e, 0x54, 0x32, 0x96, 0x70, 0xfc, 0x29, 0x60, 0x33, 0xc9, 0x70, 0x39, 0x46, 0x49, 0x32, 0xbd,
	0xb3, 0xc9, 0x86, 0xdc, 0x06, 0xdd, 0x9d, 0xf0, 0x66, 0xc8, 0xfa, 0x19, 0xc1, 0x6e, 0xa6, 0xf7,
	0x20, 0xa6, 0x5e, 0x70, 0x19, 0xd0, 0x04, 0x8f, 0xe0, 0x35, 0x33, 0xde, 0xea, 0x0c, 0xea, 0xe4,
	0x37, 0x1e, 0x5a, 0x7e, 0x5b, 0xee, 0xab, 0x5e, 0x2e, 0xc6, 0xf1, 0xbb, 0x50, 0x0e, 0x22, 0x41,
	0x93, 0x19, 0x09, 0xf5, 0x5d, 0x3d, 0xb1, 0xd5, 0xf7, 0xd9, 0x36, 0xdf, 0x67, 0xbb, 0xab, 0xbf,
	0xcf, 0x6e, 0x0a, 0x6d, 0xff, 0x5d, 0x84, 0xba, 0xe2, 0xe9, 0x9a, 0xae, 0x03, 0x25, 0x03, 0x5f,
	0x23, 0xd8, 0x19, 0x88, 0x84, 0x92, 0x49, 0xd6, 0xaa, 0xb3, 0xfb, 0xbd, 0x15, 0xeb, 0xde, 0xbd,
	0xe6, 0xe1, 0x1d, 0x79, 0xd2, 0x2d, 0x5a, 0x85, 0x16, 0x3a, 0x44, 0xf8, 0x17, 0x04, 0xb5, 0x33,
	0x2a, 0xbc, 0xf1, 0x8b, 0x21, 0xea, 0xe0, 0xeb, 0x3f, 0xfe, 0xba, 0x2e, 0xee, 0x59, 0xcd, 0xc5,
	0x3f, 0x6a, 0x0a, 0x3f, 0xce, 0xda, 0x2c, 0x11, 0xa5, 0x63, 0xb4, 0x7f, 0xf2, 0xd9, 0x3f, 0x3f,
	0xfe, 0xfb, 0xfd, 0xc3, 0x27, 0xf8, 0xf1, 0x6a, 0x17, 0x05, 0xb6, 0x67, 0x47, 0xbf, 0x7e, 0xf5,
	0xdb, 0xef, 0x8f, 0x8a, 0x35, 0x04, 0xad, 0x80, 0x29, 0x25, 0x71, 0xc2, 0xbe, 0x9c, 0x6f, 0x10,
	0x75, 0x52, 0xee, 0xf9, 0x5c, 0xfe, 0x5f, 0xf5, 0xd1, 0x37, 0x08, 0x8d, 0x1e, 0x49, 0xaf, 0x8f,
	0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x44, 0xf0, 0xa3, 0x70, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthDiscoveryServiceClient is the client API for HealthDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthDiscoveryServiceClient interface {
	StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error)
	FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error)
}

type healthDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewHealthDiscoveryServiceClient(cc *grpc.ClientConn) HealthDiscoveryServiceClient {
	return &healthDiscoveryServiceClient{cc}
}

func (c *healthDiscoveryServiceClient) StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HealthDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.HealthDiscoveryService/StreamHealthCheck", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthDiscoveryServiceStreamHealthCheckClient{stream}
	return x, nil
}

type HealthDiscoveryService_StreamHealthCheckClient interface {
	Send(*HealthCheckRequestOrEndpointHealthResponse) error
	Recv() (*HealthCheckSpecifier, error)
	grpc.ClientStream
}

type healthDiscoveryServiceStreamHealthCheckClient struct {
	grpc.ClientStream
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Send(m *HealthCheckRequestOrEndpointHealthResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Recv() (*HealthCheckSpecifier, error) {
	m := new(HealthCheckSpecifier)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *healthDiscoveryServiceClient) FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error) {
	out := new(HealthCheckSpecifier)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v2.HealthDiscoveryService/FetchHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthDiscoveryServiceServer is the server API for HealthDiscoveryService service.
type HealthDiscoveryServiceServer interface {
	StreamHealthCheck(HealthDiscoveryService_StreamHealthCheckServer) error
	FetchHealthCheck(context.Context, *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error)
}

// UnimplementedHealthDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthDiscoveryServiceServer struct {
}

func (*UnimplementedHealthDiscoveryServiceServer) StreamHealthCheck(srv HealthDiscoveryService_StreamHealthCheckServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamHealthCheck not implemented")
}
func (*UnimplementedHealthDiscoveryServiceServer) FetchHealthCheck(ctx context.Context, req *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchHealthCheck not implemented")
}

func RegisterHealthDiscoveryServiceServer(s *grpc.Server, srv HealthDiscoveryServiceServer) {
	s.RegisterService(&_HealthDiscoveryService_serviceDesc, srv)
}

func _HealthDiscoveryService_StreamHealthCheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HealthDiscoveryServiceServer).StreamHealthCheck(&healthDiscoveryServiceStreamHealthCheckServer{stream})
}

type HealthDiscoveryService_StreamHealthCheckServer interface {
	Send(*HealthCheckSpecifier) error
	Recv() (*HealthCheckRequestOrEndpointHealthResponse, error)
	grpc.ServerStream
}

type healthDiscoveryServiceStreamHealthCheckServer struct {
	grpc.ServerStream
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Send(m *HealthCheckSpecifier) error {
	return x.ServerStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Recv() (*HealthCheckRequestOrEndpointHealthResponse, error) {
	m := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HealthDiscoveryService_FetchHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v2.HealthDiscoveryService/FetchHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, req.(*HealthCheckRequestOrEndpointHealthResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.HealthDiscoveryService",
	HandlerType: (*HealthDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchHealthCheck",
			Handler:    _HealthDiscoveryService_FetchHealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHealthCheck",
			Handler:       _HealthDiscoveryService_StreamHealthCheck_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/hds.proto",
}