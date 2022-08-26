// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: mesh/v1alpha1/zone_ingress.proto

package v1alpha1

import (
	_ "github.com/kumahq/kuma/api/mesh"
	_ "github.com/kumahq/protoc-gen-kumadoc/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ZoneIngress allows us to configure dataplane in the Ingress mode. In this
// mode, dataplane has only inbound interfaces. Every inbound interface matches
// with services that reside in that cluster.
type ZoneIngress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Zone field contains Zone name where ingress is serving, field will be
	// automatically set by Global Kuma CP
	Zone string `protobuf:"bytes,1,opt,name=zone,proto3" json:"zone,omitempty"`
	// Networking defines the address and port of the Ingress to listen on.
	// Additionally publicly advertised address and port could be specified.
	Networking *ZoneIngress_Networking `protobuf:"bytes,2,opt,name=networking,proto3" json:"networking,omitempty"`
	// AvailableService contains tags that represent unique subset of
	// endpoints
	AvailableServices []*ZoneIngress_AvailableService `protobuf:"bytes,3,rep,name=availableServices,proto3" json:"availableServices,omitempty"`
}

func (x *ZoneIngress) Reset() {
	*x = ZoneIngress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_zone_ingress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneIngress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneIngress) ProtoMessage() {}

func (x *ZoneIngress) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_zone_ingress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneIngress.ProtoReflect.Descriptor instead.
func (*ZoneIngress) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_zone_ingress_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneIngress) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *ZoneIngress) GetNetworking() *ZoneIngress_Networking {
	if x != nil {
		return x.Networking
	}
	return nil
}

func (x *ZoneIngress) GetAvailableServices() []*ZoneIngress_AvailableService {
	if x != nil {
		return x.AvailableServices
	}
	return nil
}

type ZoneIngress_Networking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address on which inbound listener will be exposed
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// AdvertisedAddress defines IP or DNS name on which ZoneIngress is
	// accessible to other Kuma clusters.
	AdvertisedAddress string `protobuf:"bytes,2,opt,name=advertisedAddress,proto3" json:"advertisedAddress,omitempty"`
	// Port of the inbound interface that will forward requests to the service.
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// AdvertisedPort defines port on which ZoneIngress is accessible to other
	// Kuma clusters.
	AdvertisedPort uint32 `protobuf:"varint,4,opt,name=advertisedPort,proto3" json:"advertisedPort,omitempty"`
	// Admin contains configuration related to Envoy Admin API
	Admin *EnvoyAdmin `protobuf:"bytes,5,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *ZoneIngress_Networking) Reset() {
	*x = ZoneIngress_Networking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_zone_ingress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneIngress_Networking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneIngress_Networking) ProtoMessage() {}

func (x *ZoneIngress_Networking) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_zone_ingress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneIngress_Networking.ProtoReflect.Descriptor instead.
func (*ZoneIngress_Networking) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_zone_ingress_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ZoneIngress_Networking) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ZoneIngress_Networking) GetAdvertisedAddress() string {
	if x != nil {
		return x.AdvertisedAddress
	}
	return ""
}

func (x *ZoneIngress_Networking) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ZoneIngress_Networking) GetAdvertisedPort() uint32 {
	if x != nil {
		return x.AdvertisedPort
	}
	return 0
}

func (x *ZoneIngress_Networking) GetAdmin() *EnvoyAdmin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type ZoneIngress_AvailableService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tags of the service
	Tags map[string]string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// number of instances available for given tags
	Instances uint32 `protobuf:"varint,2,opt,name=instances,proto3" json:"instances,omitempty"`
	// mesh of the instances available for given tags
	Mesh string `protobuf:"bytes,3,opt,name=mesh,proto3" json:"mesh,omitempty"`
	// instance of external service available from the zone
	ExternalService bool `protobuf:"varint,4,opt,name=externalService,proto3" json:"externalService,omitempty"`
}

func (x *ZoneIngress_AvailableService) Reset() {
	*x = ZoneIngress_AvailableService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_zone_ingress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneIngress_AvailableService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneIngress_AvailableService) ProtoMessage() {}

func (x *ZoneIngress_AvailableService) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_zone_ingress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneIngress_AvailableService.ProtoReflect.Descriptor instead.
func (*ZoneIngress_AvailableService) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_zone_ingress_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ZoneIngress_AvailableService) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ZoneIngress_AvailableService) GetInstances() uint32 {
	if x != nil {
		return x.Instances
	}
	return 0
}

func (x *ZoneIngress_AvailableService) GetMesh() string {
	if x != nil {
		return x.Mesh
	}
	return ""
}

func (x *ZoneIngress_AvailableService) GetExternalService() bool {
	if x != nil {
		return x.ExternalService
	}
	return false
}

var File_mesh_v1alpha1_zone_ingress_proto protoreflect.FileDescriptor

var file_mesh_v1alpha1_zone_ingress_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x12, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x06, 0x0a, 0x0b, 0x5a, 0x6f,
	0x6e, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x50, 0x0a,
	0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x04, 0x88,
	0xb5, 0x18, 0x01, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x5e, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x75, 0x6d,
	0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x11, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0xde, 0x01, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1e,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32,
	0x0a, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52,
	0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x0e,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x0e, 0x61, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x75, 0x6d, 0x61,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45,
	0x6e, 0x76, 0x6f, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x1a, 0xf7, 0x01, 0x0a, 0x10, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x8c, 0x01, 0xaa, 0x8c, 0x89,
	0xa6, 0x01, 0x15, 0x0a, 0x13, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x0d, 0x12, 0x0b,
	0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0xaa, 0x8c, 0x89, 0xa6, 0x01,
	0x02, 0x18, 0x01, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x06, 0x22, 0x04, 0x6d, 0x65, 0x73, 0x68, 0xaa,
	0x8c, 0x89, 0xa6, 0x01, 0x04, 0x52, 0x02, 0x08, 0x01, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x04, 0x52,
	0x02, 0x10, 0x01, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x10, 0x3a, 0x0e, 0x0a, 0x0c, 0x7a, 0x6f, 0x6e,
	0x65, 0x2d, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x12, 0x3a,
	0x10, 0x12, 0x0e, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x02, 0x58, 0x01, 0x42, 0x4c, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x68, 0x71, 0x2f, 0x6b,
	0x75, 0x6d, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x8a, 0xb5, 0x18, 0x1e, 0x50, 0x02, 0xa2, 0x01, 0x0b, 0x5a, 0x6f,
	0x6e, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0xf2, 0x01, 0x0b, 0x7a, 0x6f, 0x6e, 0x65,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mesh_v1alpha1_zone_ingress_proto_rawDescOnce sync.Once
	file_mesh_v1alpha1_zone_ingress_proto_rawDescData = file_mesh_v1alpha1_zone_ingress_proto_rawDesc
)

func file_mesh_v1alpha1_zone_ingress_proto_rawDescGZIP() []byte {
	file_mesh_v1alpha1_zone_ingress_proto_rawDescOnce.Do(func() {
		file_mesh_v1alpha1_zone_ingress_proto_rawDescData = protoimpl.X.CompressGZIP(file_mesh_v1alpha1_zone_ingress_proto_rawDescData)
	})
	return file_mesh_v1alpha1_zone_ingress_proto_rawDescData
}

var file_mesh_v1alpha1_zone_ingress_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mesh_v1alpha1_zone_ingress_proto_goTypes = []interface{}{
	(*ZoneIngress)(nil),                  // 0: kuma.mesh.v1alpha1.ZoneIngress
	(*ZoneIngress_Networking)(nil),       // 1: kuma.mesh.v1alpha1.ZoneIngress.Networking
	(*ZoneIngress_AvailableService)(nil), // 2: kuma.mesh.v1alpha1.ZoneIngress.AvailableService
	nil,                                  // 3: kuma.mesh.v1alpha1.ZoneIngress.AvailableService.TagsEntry
	(*EnvoyAdmin)(nil),                   // 4: kuma.mesh.v1alpha1.EnvoyAdmin
}
var file_mesh_v1alpha1_zone_ingress_proto_depIdxs = []int32{
	1, // 0: kuma.mesh.v1alpha1.ZoneIngress.networking:type_name -> kuma.mesh.v1alpha1.ZoneIngress.Networking
	2, // 1: kuma.mesh.v1alpha1.ZoneIngress.availableServices:type_name -> kuma.mesh.v1alpha1.ZoneIngress.AvailableService
	4, // 2: kuma.mesh.v1alpha1.ZoneIngress.Networking.admin:type_name -> kuma.mesh.v1alpha1.EnvoyAdmin
	3, // 3: kuma.mesh.v1alpha1.ZoneIngress.AvailableService.tags:type_name -> kuma.mesh.v1alpha1.ZoneIngress.AvailableService.TagsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mesh_v1alpha1_zone_ingress_proto_init() }
func file_mesh_v1alpha1_zone_ingress_proto_init() {
	if File_mesh_v1alpha1_zone_ingress_proto != nil {
		return
	}
	file_mesh_v1alpha1_envoy_admin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mesh_v1alpha1_zone_ingress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneIngress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mesh_v1alpha1_zone_ingress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneIngress_Networking); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mesh_v1alpha1_zone_ingress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneIngress_AvailableService); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mesh_v1alpha1_zone_ingress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mesh_v1alpha1_zone_ingress_proto_goTypes,
		DependencyIndexes: file_mesh_v1alpha1_zone_ingress_proto_depIdxs,
		MessageInfos:      file_mesh_v1alpha1_zone_ingress_proto_msgTypes,
	}.Build()
	File_mesh_v1alpha1_zone_ingress_proto = out.File
	file_mesh_v1alpha1_zone_ingress_proto_rawDesc = nil
	file_mesh_v1alpha1_zone_ingress_proto_goTypes = nil
	file_mesh_v1alpha1_zone_ingress_proto_depIdxs = nil
}