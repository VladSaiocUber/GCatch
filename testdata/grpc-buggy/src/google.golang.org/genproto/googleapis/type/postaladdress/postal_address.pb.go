// Copyright 2019 Google LLC.
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/type/postal_address.proto

package postaladdress

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// Represents a postal address, e.g. for postal delivery or payments addresses.
// Given a postal address, a postal service can deliver items to a premise, P.O.
// Box or similar.
// It is not intended to model geographical locations (roads, towns,
// mountains).
//
// In typical usage an address would be created via user input or from importing
// existing data, depending on the type of process.
//
// Advice on address input / editing:
//  - Use an i18n-ready address widget such as
//    https://github.com/google/libaddressinput)
// - Users should not be presented with UI elements for input or editing of
//   fields outside countries where that field is used.
//
// For more guidance on how to use this schema, please see:
// https://support.google.com/business/answer/6397478
type PostalAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The schema revision of the `PostalAddress`. This must be set to 0, which is
	// the latest revision.
	//
	// All new revisions **must** be backward compatible with old revisions.
	Revision int32 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// Required. CLDR region code of the country/region of the address. This
	// is never inferred and it is up to the user to ensure the value is
	// correct. See http://cldr.unicode.org/ and
	// http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html
	// for details. Example: "CH" for Switzerland.
	RegionCode string `protobuf:"bytes,2,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// Optional. BCP-47 language code of the contents of this address (if
	// known). This is often the UI language of the input form or is expected
	// to match one of the languages used in the address' country/region, or their
	// transliterated equivalents.
	// This can affect formatting in certain countries, but is not critical
	// to the correctness of the data and will never affect any validation or
	// other non-formatting related operations.
	//
	// If this value is not known, it should be omitted (rather than specifying a
	// possibly incorrect default).
	//
	// Examples: "zh-Hant", "ja", "ja-Latn", "en".
	LanguageCode string `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. Postal code of the address. Not all countries use or require
	// postal codes to be present, but where they are used, they may trigger
	// additional validation with other parts of the address (e.g. state/zip
	// validation in the U.S.A.).
	PostalCode string `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	// Optional. Additional, country-specific, sorting code. This is not used
	// in most regions. Where it is used, the value is either a string like
	// "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number
	// alone, representing the "sector code" (Jamaica), "delivery area indicator"
	// (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
	SortingCode string `protobuf:"bytes,5,opt,name=sorting_code,json=sortingCode,proto3" json:"sorting_code,omitempty"`
	// Optional. Highest administrative subdivision which is used for postal
	// addresses of a country or region.
	// For example, this can be a state, a province, an oblast, or a prefecture.
	// Specifically, for Spain this is the province and not the autonomous
	// community (e.g. "Barcelona" and not "Catalonia").
	// Many countries don't use an administrative area in postal addresses. E.g.
	// in Switzerland this should be left unpopulated.
	AdministrativeArea string `protobuf:"bytes,6,opt,name=administrative_area,json=administrativeArea,proto3" json:"administrative_area,omitempty"`
	// Optional. Generally refers to the city/town portion of the address.
	// Examples: US city, IT comune, UK post town.
	// In regions of the world where localities are not well defined or do not fit
	// into this structure well, leave locality empty and use address_lines.
	Locality string `protobuf:"bytes,7,opt,name=locality,proto3" json:"locality,omitempty"`
	// Optional. Sublocality of the address.
	// For example, this can be neighborhoods, boroughs, districts.
	Sublocality string `protobuf:"bytes,8,opt,name=sublocality,proto3" json:"sublocality,omitempty"`
	// Unstructured address lines describing the lower levels of an address.
	//
	// Because values in address_lines do not have type information and may
	// sometimes contain multiple values in a single field (e.g.
	// "Austin, TX"), it is important that the line order is clear. The order of
	// address lines should be "envelope order" for the country/region of the
	// address. In places where this can vary (e.g. Japan), address_language is
	// used to make it explicit (e.g. "ja" for large-to-small ordering and
	// "ja-Latn" or "en" for small-to-large). This way, the most specific line of
	// an address can be selected based on the language.
	//
	// The minimum permitted structural representation of an address consists
	// of a region_code with all remaining information placed in the
	// address_lines. It would be possible to format such an address very
	// approximately without geocoding, but no semantic reasoning could be
	// made about any of the address components until it was at least
	// partially resolved.
	//
	// Creating an address only containing a region_code and address_lines, and
	// then geocoding is the recommended way to handle completely unstructured
	// addresses (as opposed to guessing which parts of the address should be
	// localities or administrative areas).
	AddressLines []string `protobuf:"bytes,9,rep,name=address_lines,json=addressLines,proto3" json:"address_lines,omitempty"`
	// Optional. The recipient at the address.
	// This field may, under certain circumstances, contain multiline information.
	// For example, it might contain "care of" information.
	Recipients []string `protobuf:"bytes,10,rep,name=recipients,proto3" json:"recipients,omitempty"`
	// Optional. The name of the organization at the address.
	Organization string `protobuf:"bytes,11,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *PostalAddress) Reset() {
	*x = PostalAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_type_postal_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostalAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostalAddress) ProtoMessage() {}

func (x *PostalAddress) ProtoReflect() protoreflect.Message {
	mi := &file_google_type_postal_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostalAddress.ProtoReflect.Descriptor instead.
func (*PostalAddress) Descriptor() ([]byte, []int) {
	return file_google_type_postal_address_proto_rawDescGZIP(), []int{0}
}

func (x *PostalAddress) GetRevision() int32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *PostalAddress) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *PostalAddress) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *PostalAddress) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *PostalAddress) GetSortingCode() string {
	if x != nil {
		return x.SortingCode
	}
	return ""
}

func (x *PostalAddress) GetAdministrativeArea() string {
	if x != nil {
		return x.AdministrativeArea
	}
	return ""
}

func (x *PostalAddress) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *PostalAddress) GetSublocality() string {
	if x != nil {
		return x.Sublocality
	}
	return ""
}

func (x *PostalAddress) GetAddressLines() []string {
	if x != nil {
		return x.AddressLines
	}
	return nil
}

func (x *PostalAddress) GetRecipients() []string {
	if x != nil {
		return x.Recipients
	}
	return nil
}

func (x *PostalAddress) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

var File_google_type_postal_address_proto protoreflect.FileDescriptor

var file_google_type_postal_address_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x8d, 0x03, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x78, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x47, 0x54, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_type_postal_address_proto_rawDescOnce sync.Once
	file_google_type_postal_address_proto_rawDescData = file_google_type_postal_address_proto_rawDesc
)

func file_google_type_postal_address_proto_rawDescGZIP() []byte {
	file_google_type_postal_address_proto_rawDescOnce.Do(func() {
		file_google_type_postal_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_type_postal_address_proto_rawDescData)
	})
	return file_google_type_postal_address_proto_rawDescData
}

var file_google_type_postal_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_type_postal_address_proto_goTypes = []interface{}{
	(*PostalAddress)(nil), // 0: google.type.PostalAddress
}
var file_google_type_postal_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_type_postal_address_proto_init() }
func file_google_type_postal_address_proto_init() {
	if File_google_type_postal_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_type_postal_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostalAddress); i {
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
			RawDescriptor: file_google_type_postal_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_type_postal_address_proto_goTypes,
		DependencyIndexes: file_google_type_postal_address_proto_depIdxs,
		MessageInfos:      file_google_type_postal_address_proto_msgTypes,
	}.Build()
	File_google_type_postal_address_proto = out.File
	file_google_type_postal_address_proto_rawDesc = nil
	file_google_type_postal_address_proto_goTypes = nil
	file_google_type_postal_address_proto_depIdxs = nil
}
