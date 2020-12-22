// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v4/services/customer_label_service.proto

package services

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v4/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for [CustomerLabelService.GetCustomerLabel][google.ads.googleads.v4.services.CustomerLabelService.GetCustomerLabel].
type GetCustomerLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the customer-label relationship to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetCustomerLabelRequest) Reset() {
	*x = GetCustomerLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerLabelRequest) ProtoMessage() {}

func (x *GetCustomerLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerLabelRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerLabelRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerLabelRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for [CustomerLabelService.MutateCustomerLabels][google.ads.googleads.v4.services.CustomerLabelService.MutateCustomerLabels].
type MutateCustomerLabelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the customer whose customer-label relationships are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on customer-label relationships.
	Operations []*CustomerLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateCustomerLabelsRequest) Reset() {
	*x = MutateCustomerLabelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerLabelsRequest) ProtoMessage() {}

func (x *MutateCustomerLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerLabelsRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerLabelsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateCustomerLabelsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerLabelsRequest) GetOperations() []*CustomerLabelOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateCustomerLabelsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateCustomerLabelsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on a customer-label relationship.
type CustomerLabelOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*CustomerLabelOperation_Create
	//	*CustomerLabelOperation_Remove
	Operation isCustomerLabelOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CustomerLabelOperation) Reset() {
	*x = CustomerLabelOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerLabelOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerLabelOperation) ProtoMessage() {}

func (x *CustomerLabelOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerLabelOperation.ProtoReflect.Descriptor instead.
func (*CustomerLabelOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescGZIP(), []int{2}
}

func (m *CustomerLabelOperation) GetOperation() isCustomerLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CustomerLabelOperation) GetCreate() *resources.CustomerLabel {
	if x, ok := x.GetOperation().(*CustomerLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *CustomerLabelOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*CustomerLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isCustomerLabelOperation_Operation interface {
	isCustomerLabelOperation_Operation()
}

type CustomerLabelOperation_Create struct {
	// Create operation: No resource name is expected for the new customer-label
	// relationship.
	Create *resources.CustomerLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerLabelOperation_Remove struct {
	// Remove operation: A resource name for the customer-label relationship
	// being removed, in this format:
	//
	// `customers/{customer_id}/customerLabels/{label_id}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CustomerLabelOperation_Create) isCustomerLabelOperation_Operation() {}

func (*CustomerLabelOperation_Remove) isCustomerLabelOperation_Operation() {}

// Response message for a customer labels mutate.
type MutateCustomerLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateCustomerLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateCustomerLabelsResponse) Reset() {
	*x = MutateCustomerLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerLabelsResponse) ProtoMessage() {}

func (x *MutateCustomerLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerLabelsResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerLabelsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerLabelsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateCustomerLabelsResponse) GetResults() []*MutateCustomerLabelResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for a customer label mutate.
type MutateCustomerLabelResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateCustomerLabelResult) Reset() {
	*x = MutateCustomerLabelResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerLabelResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerLabelResult) ProtoMessage() {}

func (x *MutateCustomerLabelResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerLabelResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerLabelResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateCustomerLabelResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v4_services_customer_label_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v4_services_customer_label_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x53, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x28, 0x0a,
	0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf0, 0x01, 0x0a, 0x1b, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x5d, 0x0a, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x8b, 0x01, 0x0a, 0x16, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbd, 0x01, 0x0a, 0x1c, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x55, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x40, 0x0a, 0x19, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xf0, 0x03, 0x0a, 0x14, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xc9, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x22, 0x48, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x76, 0x34,
	0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xee, 0x01,
	0x0a, 0x14, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x22, 0x33, 0x2f,
	0x76, 0x34, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1b,
	0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x80, 0x02, 0x0a, 0x24,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x42, 0x19, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x34, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x34, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x34, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescData = file_google_ads_googleads_v4_services_customer_label_service_proto_rawDesc
)

func file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v4_services_customer_label_service_proto_rawDescData
}

var file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v4_services_customer_label_service_proto_goTypes = []interface{}{
	(*GetCustomerLabelRequest)(nil),      // 0: google.ads.googleads.v4.services.GetCustomerLabelRequest
	(*MutateCustomerLabelsRequest)(nil),  // 1: google.ads.googleads.v4.services.MutateCustomerLabelsRequest
	(*CustomerLabelOperation)(nil),       // 2: google.ads.googleads.v4.services.CustomerLabelOperation
	(*MutateCustomerLabelsResponse)(nil), // 3: google.ads.googleads.v4.services.MutateCustomerLabelsResponse
	(*MutateCustomerLabelResult)(nil),    // 4: google.ads.googleads.v4.services.MutateCustomerLabelResult
	(*resources.CustomerLabel)(nil),      // 5: google.ads.googleads.v4.resources.CustomerLabel
	(*status.Status)(nil),                // 6: google.rpc.Status
}
var file_google_ads_googleads_v4_services_customer_label_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v4.services.MutateCustomerLabelsRequest.operations:type_name -> google.ads.googleads.v4.services.CustomerLabelOperation
	5, // 1: google.ads.googleads.v4.services.CustomerLabelOperation.create:type_name -> google.ads.googleads.v4.resources.CustomerLabel
	6, // 2: google.ads.googleads.v4.services.MutateCustomerLabelsResponse.partial_failure_error:type_name -> google.rpc.Status
	4, // 3: google.ads.googleads.v4.services.MutateCustomerLabelsResponse.results:type_name -> google.ads.googleads.v4.services.MutateCustomerLabelResult
	0, // 4: google.ads.googleads.v4.services.CustomerLabelService.GetCustomerLabel:input_type -> google.ads.googleads.v4.services.GetCustomerLabelRequest
	1, // 5: google.ads.googleads.v4.services.CustomerLabelService.MutateCustomerLabels:input_type -> google.ads.googleads.v4.services.MutateCustomerLabelsRequest
	5, // 6: google.ads.googleads.v4.services.CustomerLabelService.GetCustomerLabel:output_type -> google.ads.googleads.v4.resources.CustomerLabel
	3, // 7: google.ads.googleads.v4.services.CustomerLabelService.MutateCustomerLabels:output_type -> google.ads.googleads.v4.services.MutateCustomerLabelsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v4_services_customer_label_service_proto_init() }
func file_google_ads_googleads_v4_services_customer_label_service_proto_init() {
	if File_google_ads_googleads_v4_services_customer_label_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerLabelRequest); i {
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
		file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerLabelsRequest); i {
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
		file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerLabelOperation); i {
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
		file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerLabelsResponse); i {
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
		file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerLabelResult); i {
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
	file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CustomerLabelOperation_Create)(nil),
		(*CustomerLabelOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v4_services_customer_label_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v4_services_customer_label_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v4_services_customer_label_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v4_services_customer_label_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v4_services_customer_label_service_proto = out.File
	file_google_ads_googleads_v4_services_customer_label_service_proto_rawDesc = nil
	file_google_ads_googleads_v4_services_customer_label_service_proto_goTypes = nil
	file_google_ads_googleads_v4_services_customer_label_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerLabelServiceClient is the client API for CustomerLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerLabelServiceClient interface {
	// Returns the requested customer-label relationship in full detail.
	GetCustomerLabel(ctx context.Context, in *GetCustomerLabelRequest, opts ...grpc.CallOption) (*resources.CustomerLabel, error)
	// Creates and removes customer-label relationships.
	// Operation statuses are returned.
	MutateCustomerLabels(ctx context.Context, in *MutateCustomerLabelsRequest, opts ...grpc.CallOption) (*MutateCustomerLabelsResponse, error)
}

type customerLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerLabelServiceClient(cc grpc.ClientConnInterface) CustomerLabelServiceClient {
	return &customerLabelServiceClient{cc}
}

func (c *customerLabelServiceClient) GetCustomerLabel(ctx context.Context, in *GetCustomerLabelRequest, opts ...grpc.CallOption) (*resources.CustomerLabel, error) {
	out := new(resources.CustomerLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v4.services.CustomerLabelService/GetCustomerLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerLabelServiceClient) MutateCustomerLabels(ctx context.Context, in *MutateCustomerLabelsRequest, opts ...grpc.CallOption) (*MutateCustomerLabelsResponse, error) {
	out := new(MutateCustomerLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v4.services.CustomerLabelService/MutateCustomerLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerLabelServiceServer is the server API for CustomerLabelService service.
type CustomerLabelServiceServer interface {
	// Returns the requested customer-label relationship in full detail.
	GetCustomerLabel(context.Context, *GetCustomerLabelRequest) (*resources.CustomerLabel, error)
	// Creates and removes customer-label relationships.
	// Operation statuses are returned.
	MutateCustomerLabels(context.Context, *MutateCustomerLabelsRequest) (*MutateCustomerLabelsResponse, error)
}

// UnimplementedCustomerLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerLabelServiceServer struct {
}

func (*UnimplementedCustomerLabelServiceServer) GetCustomerLabel(context.Context, *GetCustomerLabelRequest) (*resources.CustomerLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCustomerLabel not implemented")
}
func (*UnimplementedCustomerLabelServiceServer) MutateCustomerLabels(context.Context, *MutateCustomerLabelsRequest) (*MutateCustomerLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCustomerLabels not implemented")
}

func RegisterCustomerLabelServiceServer(s *grpc.Server, srv CustomerLabelServiceServer) {
	s.RegisterService(&_CustomerLabelService_serviceDesc, srv)
}

func _CustomerLabelService_GetCustomerLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerLabelServiceServer).GetCustomerLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v4.services.CustomerLabelService/GetCustomerLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerLabelServiceServer).GetCustomerLabel(ctx, req.(*GetCustomerLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerLabelService_MutateCustomerLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerLabelServiceServer).MutateCustomerLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v4.services.CustomerLabelService/MutateCustomerLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerLabelServiceServer).MutateCustomerLabels(ctx, req.(*MutateCustomerLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v4.services.CustomerLabelService",
	HandlerType: (*CustomerLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerLabel",
			Handler:    _CustomerLabelService_GetCustomerLabel_Handler,
		},
		{
			MethodName: "MutateCustomerLabels",
			Handler:    _CustomerLabelService_MutateCustomerLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v4/services/customer_label_service.proto",
}
