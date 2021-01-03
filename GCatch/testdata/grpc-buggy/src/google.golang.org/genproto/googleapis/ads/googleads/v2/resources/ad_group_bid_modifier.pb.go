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
// source: google/ads/googleads/v2/resources/ad_group_bid_modifier.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Represents an ad group bid modifier.
type AdGroupBidModifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the ad group bid modifier.
	// Ad group bid modifier resource names have the form:
	//
	// `customers/{customer_id}/adGroupBidModifiers/{ad_group_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The ad group to which this criterion belongs.
	AdGroup *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Output only. The ID of the criterion to bid modify.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bid when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. The range is 1.0 - 6.0 for PreferredContent.
	// Use 0 to opt out of a Device type.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// Output only. The base ad group from which this draft/trial adgroup bid modifier was
	// created. If ad_group is a base ad group then this field will be equal to
	// ad_group. If the ad group was created in the draft or trial and has no
	// corresponding base ad group, then this field will be null.
	// This field is readonly.
	BaseAdGroup *wrappers.StringValue `protobuf:"bytes,9,opt,name=base_ad_group,json=baseAdGroup,proto3" json:"base_ad_group,omitempty"`
	// Output only. Bid modifier source.
	BidModifierSource enums.BidModifierSourceEnum_BidModifierSource `protobuf:"varint,10,opt,name=bid_modifier_source,json=bidModifierSource,proto3,enum=google.ads.googleads.v2.enums.BidModifierSourceEnum_BidModifierSource" json:"bid_modifier_source,omitempty"`
	// The criterion of this ad group bid modifier.
	//
	// Types that are assignable to Criterion:
	//	*AdGroupBidModifier_HotelDateSelectionType
	//	*AdGroupBidModifier_HotelAdvanceBookingWindow
	//	*AdGroupBidModifier_HotelLengthOfStay
	//	*AdGroupBidModifier_HotelCheckInDay
	//	*AdGroupBidModifier_Device
	//	*AdGroupBidModifier_PreferredContent
	Criterion isAdGroupBidModifier_Criterion `protobuf_oneof:"criterion"`
}

func (x *AdGroupBidModifier) Reset() {
	*x = AdGroupBidModifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupBidModifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupBidModifier) ProtoMessage() {}

func (x *AdGroupBidModifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupBidModifier.ProtoReflect.Descriptor instead.
func (*AdGroupBidModifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupBidModifier) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupBidModifier) GetAdGroup() *wrappers.StringValue {
	if x != nil {
		return x.AdGroup
	}
	return nil
}

func (x *AdGroupBidModifier) GetCriterionId() *wrappers.Int64Value {
	if x != nil {
		return x.CriterionId
	}
	return nil
}

func (x *AdGroupBidModifier) GetBidModifier() *wrappers.DoubleValue {
	if x != nil {
		return x.BidModifier
	}
	return nil
}

func (x *AdGroupBidModifier) GetBaseAdGroup() *wrappers.StringValue {
	if x != nil {
		return x.BaseAdGroup
	}
	return nil
}

func (x *AdGroupBidModifier) GetBidModifierSource() enums.BidModifierSourceEnum_BidModifierSource {
	if x != nil {
		return x.BidModifierSource
	}
	return enums.BidModifierSourceEnum_UNSPECIFIED
}

func (m *AdGroupBidModifier) GetCriterion() isAdGroupBidModifier_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (x *AdGroupBidModifier) GetHotelDateSelectionType() *common.HotelDateSelectionTypeInfo {
	if x, ok := x.GetCriterion().(*AdGroupBidModifier_HotelDateSelectionType); ok {
		return x.HotelDateSelectionType
	}
	return nil
}

func (x *AdGroupBidModifier) GetHotelAdvanceBookingWindow() *common.HotelAdvanceBookingWindowInfo {
	if x, ok := x.GetCriterion().(*AdGroupBidModifier_HotelAdvanceBookingWindow); ok {
		return x.HotelAdvanceBookingWindow
	}
	return nil
}

func (x *AdGroupBidModifier) GetHotelLengthOfStay() *common.HotelLengthOfStayInfo {
	if x, ok := x.GetCriterion().(*AdGroupBidModifier_HotelLengthOfStay); ok {
		return x.HotelLengthOfStay
	}
	return nil
}

func (x *AdGroupBidModifier) GetHotelCheckInDay() *common.HotelCheckInDayInfo {
	if x, ok := x.GetCriterion().(*AdGroupBidModifier_HotelCheckInDay); ok {
		return x.HotelCheckInDay
	}
	return nil
}

func (x *AdGroupBidModifier) GetDevice() *common.DeviceInfo {
	if x, ok := x.GetCriterion().(*AdGroupBidModifier_Device); ok {
		return x.Device
	}
	return nil
}

func (x *AdGroupBidModifier) GetPreferredContent() *common.PreferredContentInfo {
	if x, ok := x.GetCriterion().(*AdGroupBidModifier_PreferredContent); ok {
		return x.PreferredContent
	}
	return nil
}

type isAdGroupBidModifier_Criterion interface {
	isAdGroupBidModifier_Criterion()
}

type AdGroupBidModifier_HotelDateSelectionType struct {
	// Immutable. Criterion for hotel date selection (default dates vs. user selected).
	HotelDateSelectionType *common.HotelDateSelectionTypeInfo `protobuf:"bytes,5,opt,name=hotel_date_selection_type,json=hotelDateSelectionType,proto3,oneof"`
}

type AdGroupBidModifier_HotelAdvanceBookingWindow struct {
	// Immutable. Criterion for number of days prior to the stay the booking is being made.
	HotelAdvanceBookingWindow *common.HotelAdvanceBookingWindowInfo `protobuf:"bytes,6,opt,name=hotel_advance_booking_window,json=hotelAdvanceBookingWindow,proto3,oneof"`
}

type AdGroupBidModifier_HotelLengthOfStay struct {
	// Immutable. Criterion for length of hotel stay in nights.
	HotelLengthOfStay *common.HotelLengthOfStayInfo `protobuf:"bytes,7,opt,name=hotel_length_of_stay,json=hotelLengthOfStay,proto3,oneof"`
}

type AdGroupBidModifier_HotelCheckInDay struct {
	// Immutable. Criterion for day of the week the booking is for.
	HotelCheckInDay *common.HotelCheckInDayInfo `protobuf:"bytes,8,opt,name=hotel_check_in_day,json=hotelCheckInDay,proto3,oneof"`
}

type AdGroupBidModifier_Device struct {
	// Immutable. A device criterion.
	Device *common.DeviceInfo `protobuf:"bytes,11,opt,name=device,proto3,oneof"`
}

type AdGroupBidModifier_PreferredContent struct {
	// Immutable. A preferred content criterion.
	PreferredContent *common.PreferredContentInfo `protobuf:"bytes,12,opt,name=preferred_content,json=preferredContent,proto3,oneof"`
}

func (*AdGroupBidModifier_HotelDateSelectionType) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelAdvanceBookingWindow) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelLengthOfStay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelCheckInDay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_Device) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_PreferredContent) isAdGroupBidModifier_Criterion() {}

var File_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x69, 0x64,
	0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x0a, 0x0a, 0x12, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x61, 0x0a, 0x08, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x28, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x07, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x43, 0x0a, 0x0c, 0x63, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3f, 0x0a,
	0x0c, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0b, 0x62, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x6a,
	0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x28, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x7b, 0x0a, 0x13, 0x62, 0x69,
	0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42, 0x69,
	0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x62, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x19, 0x68, 0x6f, 0x74, 0x65, 0x6c,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x44, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x16, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x1c, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f,
	0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x6f,
	0x74, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05,
	0x48, 0x00, 0x52, 0x19, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x6d, 0x0a,
	0x14, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6f, 0x66,
	0x5f, 0x73, 0x74, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x4f, 0x66, 0x53, 0x74, 0x61, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x11, 0x68, 0x6f, 0x74, 0x65, 0x6c,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x4f, 0x66, 0x53, 0x74, 0x61, 0x79, 0x12, 0x67, 0x0a, 0x12,
	0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x64,
	0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0,
	0x41, 0x05, 0x48, 0x00, 0x52, 0x0f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x44, 0x61, 0x79, 0x12, 0x49, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x68, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x10, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x72, 0xea, 0x41, 0x6f, 0x0a,
	0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x40, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x7d, 0x42, 0x0b,
	0x0a, 0x09, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x42, 0x84, 0x02, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x17, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69,
	0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x32,
	0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDescData = file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDesc
)

func file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDescData
}

var file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_goTypes = []interface{}{
	(*AdGroupBidModifier)(nil),                         // 0: google.ads.googleads.v2.resources.AdGroupBidModifier
	(*wrappers.StringValue)(nil),                       // 1: google.protobuf.StringValue
	(*wrappers.Int64Value)(nil),                        // 2: google.protobuf.Int64Value
	(*wrappers.DoubleValue)(nil),                       // 3: google.protobuf.DoubleValue
	(enums.BidModifierSourceEnum_BidModifierSource)(0), // 4: google.ads.googleads.v2.enums.BidModifierSourceEnum.BidModifierSource
	(*common.HotelDateSelectionTypeInfo)(nil),          // 5: google.ads.googleads.v2.common.HotelDateSelectionTypeInfo
	(*common.HotelAdvanceBookingWindowInfo)(nil),       // 6: google.ads.googleads.v2.common.HotelAdvanceBookingWindowInfo
	(*common.HotelLengthOfStayInfo)(nil),               // 7: google.ads.googleads.v2.common.HotelLengthOfStayInfo
	(*common.HotelCheckInDayInfo)(nil),                 // 8: google.ads.googleads.v2.common.HotelCheckInDayInfo
	(*common.DeviceInfo)(nil),                          // 9: google.ads.googleads.v2.common.DeviceInfo
	(*common.PreferredContentInfo)(nil),                // 10: google.ads.googleads.v2.common.PreferredContentInfo
}
var file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_depIdxs = []int32{
	1,  // 0: google.ads.googleads.v2.resources.AdGroupBidModifier.ad_group:type_name -> google.protobuf.StringValue
	2,  // 1: google.ads.googleads.v2.resources.AdGroupBidModifier.criterion_id:type_name -> google.protobuf.Int64Value
	3,  // 2: google.ads.googleads.v2.resources.AdGroupBidModifier.bid_modifier:type_name -> google.protobuf.DoubleValue
	1,  // 3: google.ads.googleads.v2.resources.AdGroupBidModifier.base_ad_group:type_name -> google.protobuf.StringValue
	4,  // 4: google.ads.googleads.v2.resources.AdGroupBidModifier.bid_modifier_source:type_name -> google.ads.googleads.v2.enums.BidModifierSourceEnum.BidModifierSource
	5,  // 5: google.ads.googleads.v2.resources.AdGroupBidModifier.hotel_date_selection_type:type_name -> google.ads.googleads.v2.common.HotelDateSelectionTypeInfo
	6,  // 6: google.ads.googleads.v2.resources.AdGroupBidModifier.hotel_advance_booking_window:type_name -> google.ads.googleads.v2.common.HotelAdvanceBookingWindowInfo
	7,  // 7: google.ads.googleads.v2.resources.AdGroupBidModifier.hotel_length_of_stay:type_name -> google.ads.googleads.v2.common.HotelLengthOfStayInfo
	8,  // 8: google.ads.googleads.v2.resources.AdGroupBidModifier.hotel_check_in_day:type_name -> google.ads.googleads.v2.common.HotelCheckInDayInfo
	9,  // 9: google.ads.googleads.v2.resources.AdGroupBidModifier.device:type_name -> google.ads.googleads.v2.common.DeviceInfo
	10, // 10: google.ads.googleads.v2.resources.AdGroupBidModifier.preferred_content:type_name -> google.ads.googleads.v2.common.PreferredContentInfo
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_init() }
func file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_init() {
	if File_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupBidModifier); i {
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
	file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AdGroupBidModifier_HotelDateSelectionType)(nil),
		(*AdGroupBidModifier_HotelAdvanceBookingWindow)(nil),
		(*AdGroupBidModifier_HotelLengthOfStay)(nil),
		(*AdGroupBidModifier_HotelCheckInDay)(nil),
		(*AdGroupBidModifier_Device)(nil),
		(*AdGroupBidModifier_PreferredContent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto = out.File
	file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_rawDesc = nil
	file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_goTypes = nil
	file_google_ads_googleads_v2_resources_ad_group_bid_modifier_proto_depIdxs = nil
}