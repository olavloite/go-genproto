// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/ad_group_criterion_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible ad group criterion errors.
type AdGroupCriterionErrorEnum_AdGroupCriterionError int32

const (
	// Enum unspecified.
	AdGroupCriterionErrorEnum_UNSPECIFIED AdGroupCriterionErrorEnum_AdGroupCriterionError = 0
	// The received error code is not known in this version.
	AdGroupCriterionErrorEnum_UNKNOWN AdGroupCriterionErrorEnum_AdGroupCriterionError = 1
	// No link found between the AdGroupCriterion and the label.
	AdGroupCriterionErrorEnum_AD_GROUP_CRITERION_LABEL_DOES_NOT_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 2
	// The label has already been attached to the AdGroupCriterion.
	AdGroupCriterionErrorEnum_AD_GROUP_CRITERION_LABEL_ALREADY_EXISTS AdGroupCriterionErrorEnum_AdGroupCriterionError = 3
	// Negative AdGroupCriterion cannot have labels.
	AdGroupCriterionErrorEnum_CANNOT_ADD_LABEL_TO_NEGATIVE_CRITERION AdGroupCriterionErrorEnum_AdGroupCriterionError = 4
	// Too many operations for a single call.
	AdGroupCriterionErrorEnum_TOO_MANY_OPERATIONS AdGroupCriterionErrorEnum_AdGroupCriterionError = 5
	// Negative ad group criteria are not updateable.
	AdGroupCriterionErrorEnum_CANT_UPDATE_NEGATIVE AdGroupCriterionErrorEnum_AdGroupCriterionError = 6
	// Concrete type of criterion (keyword v.s. placement) is required for ADD
	// and SET operations.
	AdGroupCriterionErrorEnum_CONCRETE_TYPE_REQUIRED AdGroupCriterionErrorEnum_AdGroupCriterionError = 7
	// Bid is incompatible with ad group's bidding settings.
	AdGroupCriterionErrorEnum_BID_INCOMPATIBLE_WITH_ADGROUP AdGroupCriterionErrorEnum_AdGroupCriterionError = 8
	// Cannot target and exclude the same criterion at once.
	AdGroupCriterionErrorEnum_CANNOT_TARGET_AND_EXCLUDE AdGroupCriterionErrorEnum_AdGroupCriterionError = 9
	// The URL of a placement is invalid.
	AdGroupCriterionErrorEnum_ILLEGAL_URL AdGroupCriterionErrorEnum_AdGroupCriterionError = 10
	// Keyword text was invalid.
	AdGroupCriterionErrorEnum_INVALID_KEYWORD_TEXT AdGroupCriterionErrorEnum_AdGroupCriterionError = 11
	// Destination URL was invalid.
	AdGroupCriterionErrorEnum_INVALID_DESTINATION_URL AdGroupCriterionErrorEnum_AdGroupCriterionError = 12
	// The destination url must contain at least one tag (e.g. {lpurl})
	AdGroupCriterionErrorEnum_MISSING_DESTINATION_URL_TAG AdGroupCriterionErrorEnum_AdGroupCriterionError = 13
	// Keyword-level cpm bid is not supported
	AdGroupCriterionErrorEnum_KEYWORD_LEVEL_BID_NOT_SUPPORTED_FOR_MANUALCPM AdGroupCriterionErrorEnum_AdGroupCriterionError = 14
	// For example, cannot add a biddable ad group criterion that had been
	// removed.
	AdGroupCriterionErrorEnum_INVALID_USER_STATUS AdGroupCriterionErrorEnum_AdGroupCriterionError = 15
	// Criteria type cannot be targeted for the ad group. Either the account is
	// restricted to keywords only, the criteria type is incompatible with the
	// campaign's bidding strategy, or the criteria type can only be applied to
	// campaigns.
	AdGroupCriterionErrorEnum_CANNOT_ADD_CRITERIA_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 16
	// Criteria type cannot be excluded for the ad group. Refer to the
	// documentation for a specific criterion to check if it is excludable.
	AdGroupCriterionErrorEnum_CANNOT_EXCLUDE_CRITERIA_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 17
	// Partial failure is not supported for shopping campaign mutate operations.
	AdGroupCriterionErrorEnum_CAMPAIGN_TYPE_NOT_COMPATIBLE_WITH_PARTIAL_FAILURE AdGroupCriterionErrorEnum_AdGroupCriterionError = 27
	// Operations in the mutate request changes too many shopping ad groups.
	// Please split requests for multiple shopping ad groups across multiple
	// requests.
	AdGroupCriterionErrorEnum_OPERATIONS_FOR_TOO_MANY_SHOPPING_ADGROUPS AdGroupCriterionErrorEnum_AdGroupCriterionError = 28
	// Not allowed to modify url fields of an ad group criterion if there are
	// duplicate elements for that ad group criterion in the request.
	AdGroupCriterionErrorEnum_CANNOT_MODIFY_URL_FIELDS_WITH_DUPLICATE_ELEMENTS AdGroupCriterionErrorEnum_AdGroupCriterionError = 29
	// Cannot set url fields without also setting final urls.
	AdGroupCriterionErrorEnum_CANNOT_SET_WITHOUT_FINAL_URLS AdGroupCriterionErrorEnum_AdGroupCriterionError = 30
	// Cannot clear final urls if final mobile urls exist.
	AdGroupCriterionErrorEnum_CANNOT_CLEAR_FINAL_URLS_IF_FINAL_MOBILE_URLS_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 31
	// Cannot clear final urls if final app urls exist.
	AdGroupCriterionErrorEnum_CANNOT_CLEAR_FINAL_URLS_IF_FINAL_APP_URLS_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 32
	// Cannot clear final urls if tracking url template exists.
	AdGroupCriterionErrorEnum_CANNOT_CLEAR_FINAL_URLS_IF_TRACKING_URL_TEMPLATE_EXISTS AdGroupCriterionErrorEnum_AdGroupCriterionError = 33
	// Cannot clear final urls if url custom parameters exist.
	AdGroupCriterionErrorEnum_CANNOT_CLEAR_FINAL_URLS_IF_URL_CUSTOM_PARAMETERS_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 34
	// Cannot set both destination url and final urls.
	AdGroupCriterionErrorEnum_CANNOT_SET_BOTH_DESTINATION_URL_AND_FINAL_URLS AdGroupCriterionErrorEnum_AdGroupCriterionError = 35
	// Cannot set both destination url and tracking url template.
	AdGroupCriterionErrorEnum_CANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE AdGroupCriterionErrorEnum_AdGroupCriterionError = 36
	// Final urls are not supported for this criterion type.
	AdGroupCriterionErrorEnum_FINAL_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 37
	// Final mobile urls are not supported for this criterion type.
	AdGroupCriterionErrorEnum_FINAL_MOBILE_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 38
	// Ad group is invalid due to the listing groups it contains.
	AdGroupCriterionErrorEnum_INVALID_LISTING_GROUP_HIERARCHY AdGroupCriterionErrorEnum_AdGroupCriterionError = 39
	// Listing group unit cannot have children.
	AdGroupCriterionErrorEnum_LISTING_GROUP_UNIT_CANNOT_HAVE_CHILDREN AdGroupCriterionErrorEnum_AdGroupCriterionError = 40
	// Subdivided listing groups must have an "others" case.
	AdGroupCriterionErrorEnum_LISTING_GROUP_SUBDIVISION_REQUIRES_OTHERS_CASE AdGroupCriterionErrorEnum_AdGroupCriterionError = 41
	// Dimension type of listing group must be the same as that of its siblings.
	AdGroupCriterionErrorEnum_LISTING_GROUP_REQUIRES_SAME_DIMENSION_TYPE_AS_SIBLINGS AdGroupCriterionErrorEnum_AdGroupCriterionError = 42
	// Listing group cannot be added to the ad group because it already exists.
	AdGroupCriterionErrorEnum_LISTING_GROUP_ALREADY_EXISTS AdGroupCriterionErrorEnum_AdGroupCriterionError = 43
	// Listing group referenced in the operation was not found in the ad group.
	AdGroupCriterionErrorEnum_LISTING_GROUP_DOES_NOT_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 44
	// Recursive removal failed because listing group subdivision is being
	// created or modified in this request.
	AdGroupCriterionErrorEnum_LISTING_GROUP_CANNOT_BE_REMOVED AdGroupCriterionErrorEnum_AdGroupCriterionError = 45
	// Listing group type is not allowed for specified ad group criterion type.
	AdGroupCriterionErrorEnum_INVALID_LISTING_GROUP_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 46
	// Listing group in an ADD operation specifies a non temporary criterion id.
	AdGroupCriterionErrorEnum_LISTING_GROUP_ADD_MAY_ONLY_USE_TEMP_ID AdGroupCriterionErrorEnum_AdGroupCriterionError = 47
)

var AdGroupCriterionErrorEnum_AdGroupCriterionError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "AD_GROUP_CRITERION_LABEL_DOES_NOT_EXIST",
	3:  "AD_GROUP_CRITERION_LABEL_ALREADY_EXISTS",
	4:  "CANNOT_ADD_LABEL_TO_NEGATIVE_CRITERION",
	5:  "TOO_MANY_OPERATIONS",
	6:  "CANT_UPDATE_NEGATIVE",
	7:  "CONCRETE_TYPE_REQUIRED",
	8:  "BID_INCOMPATIBLE_WITH_ADGROUP",
	9:  "CANNOT_TARGET_AND_EXCLUDE",
	10: "ILLEGAL_URL",
	11: "INVALID_KEYWORD_TEXT",
	12: "INVALID_DESTINATION_URL",
	13: "MISSING_DESTINATION_URL_TAG",
	14: "KEYWORD_LEVEL_BID_NOT_SUPPORTED_FOR_MANUALCPM",
	15: "INVALID_USER_STATUS",
	16: "CANNOT_ADD_CRITERIA_TYPE",
	17: "CANNOT_EXCLUDE_CRITERIA_TYPE",
	27: "CAMPAIGN_TYPE_NOT_COMPATIBLE_WITH_PARTIAL_FAILURE",
	28: "OPERATIONS_FOR_TOO_MANY_SHOPPING_ADGROUPS",
	29: "CANNOT_MODIFY_URL_FIELDS_WITH_DUPLICATE_ELEMENTS",
	30: "CANNOT_SET_WITHOUT_FINAL_URLS",
	31: "CANNOT_CLEAR_FINAL_URLS_IF_FINAL_MOBILE_URLS_EXIST",
	32: "CANNOT_CLEAR_FINAL_URLS_IF_FINAL_APP_URLS_EXIST",
	33: "CANNOT_CLEAR_FINAL_URLS_IF_TRACKING_URL_TEMPLATE_EXISTS",
	34: "CANNOT_CLEAR_FINAL_URLS_IF_URL_CUSTOM_PARAMETERS_EXIST",
	35: "CANNOT_SET_BOTH_DESTINATION_URL_AND_FINAL_URLS",
	36: "CANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE",
	37: "FINAL_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE",
	38: "FINAL_MOBILE_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE",
	39: "INVALID_LISTING_GROUP_HIERARCHY",
	40: "LISTING_GROUP_UNIT_CANNOT_HAVE_CHILDREN",
	41: "LISTING_GROUP_SUBDIVISION_REQUIRES_OTHERS_CASE",
	42: "LISTING_GROUP_REQUIRES_SAME_DIMENSION_TYPE_AS_SIBLINGS",
	43: "LISTING_GROUP_ALREADY_EXISTS",
	44: "LISTING_GROUP_DOES_NOT_EXIST",
	45: "LISTING_GROUP_CANNOT_BE_REMOVED",
	46: "INVALID_LISTING_GROUP_TYPE",
	47: "LISTING_GROUP_ADD_MAY_ONLY_USE_TEMP_ID",
}
var AdGroupCriterionErrorEnum_AdGroupCriterionError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"AD_GROUP_CRITERION_LABEL_DOES_NOT_EXIST":                   2,
	"AD_GROUP_CRITERION_LABEL_ALREADY_EXISTS":                   3,
	"CANNOT_ADD_LABEL_TO_NEGATIVE_CRITERION":                    4,
	"TOO_MANY_OPERATIONS":                                       5,
	"CANT_UPDATE_NEGATIVE":                                      6,
	"CONCRETE_TYPE_REQUIRED":                                    7,
	"BID_INCOMPATIBLE_WITH_ADGROUP":                             8,
	"CANNOT_TARGET_AND_EXCLUDE":                                 9,
	"ILLEGAL_URL":                                               10,
	"INVALID_KEYWORD_TEXT":                                      11,
	"INVALID_DESTINATION_URL":                                   12,
	"MISSING_DESTINATION_URL_TAG":                               13,
	"KEYWORD_LEVEL_BID_NOT_SUPPORTED_FOR_MANUALCPM":             14,
	"INVALID_USER_STATUS":                                       15,
	"CANNOT_ADD_CRITERIA_TYPE":                                  16,
	"CANNOT_EXCLUDE_CRITERIA_TYPE":                              17,
	"CAMPAIGN_TYPE_NOT_COMPATIBLE_WITH_PARTIAL_FAILURE":         27,
	"OPERATIONS_FOR_TOO_MANY_SHOPPING_ADGROUPS":                 28,
	"CANNOT_MODIFY_URL_FIELDS_WITH_DUPLICATE_ELEMENTS":          29,
	"CANNOT_SET_WITHOUT_FINAL_URLS":                             30,
	"CANNOT_CLEAR_FINAL_URLS_IF_FINAL_MOBILE_URLS_EXIST":        31,
	"CANNOT_CLEAR_FINAL_URLS_IF_FINAL_APP_URLS_EXIST":           32,
	"CANNOT_CLEAR_FINAL_URLS_IF_TRACKING_URL_TEMPLATE_EXISTS":   33,
	"CANNOT_CLEAR_FINAL_URLS_IF_URL_CUSTOM_PARAMETERS_EXIST":    34,
	"CANNOT_SET_BOTH_DESTINATION_URL_AND_FINAL_URLS":            35,
	"CANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE": 36,
	"FINAL_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE":               37,
	"FINAL_MOBILE_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE":        38,
	"INVALID_LISTING_GROUP_HIERARCHY":                           39,
	"LISTING_GROUP_UNIT_CANNOT_HAVE_CHILDREN":                   40,
	"LISTING_GROUP_SUBDIVISION_REQUIRES_OTHERS_CASE":            41,
	"LISTING_GROUP_REQUIRES_SAME_DIMENSION_TYPE_AS_SIBLINGS":    42,
	"LISTING_GROUP_ALREADY_EXISTS":                              43,
	"LISTING_GROUP_DOES_NOT_EXIST":                              44,
	"LISTING_GROUP_CANNOT_BE_REMOVED":                           45,
	"INVALID_LISTING_GROUP_TYPE":                                46,
	"LISTING_GROUP_ADD_MAY_ONLY_USE_TEMP_ID":                    47,
}

func (x AdGroupCriterionErrorEnum_AdGroupCriterionError) String() string {
	return proto.EnumName(AdGroupCriterionErrorEnum_AdGroupCriterionError_name, int32(x))
}
func (AdGroupCriterionErrorEnum_AdGroupCriterionError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_criterion_error_697f31de276b0971, []int{0, 0}
}

// Container for enum describing possible ad group criterion errors.
type AdGroupCriterionErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupCriterionErrorEnum) Reset()         { *m = AdGroupCriterionErrorEnum{} }
func (m *AdGroupCriterionErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionErrorEnum) ProtoMessage()    {}
func (*AdGroupCriterionErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_criterion_error_697f31de276b0971, []int{0}
}
func (m *AdGroupCriterionErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionErrorEnum.Unmarshal(m, b)
}
func (m *AdGroupCriterionErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionErrorEnum.Marshal(b, m, deterministic)
}
func (dst *AdGroupCriterionErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionErrorEnum.Merge(dst, src)
}
func (m *AdGroupCriterionErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionErrorEnum.Size(m)
}
func (m *AdGroupCriterionErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupCriterionErrorEnum)(nil), "google.ads.googleads.v1.errors.AdGroupCriterionErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.AdGroupCriterionErrorEnum_AdGroupCriterionError", AdGroupCriterionErrorEnum_AdGroupCriterionError_name, AdGroupCriterionErrorEnum_AdGroupCriterionError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/ad_group_criterion_error.proto", fileDescriptor_ad_group_criterion_error_697f31de276b0971)
}

var fileDescriptor_ad_group_criterion_error_697f31de276b0971 = []byte{
	// 976 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x72, 0x1b, 0x35,
	0x14, 0xa6, 0x29, 0xb4, 0xa0, 0x06, 0x2a, 0xc4, 0x4f, 0xdb, 0xfc, 0xb6, 0x29, 0xb4, 0xb4, 0x21,
	0x36, 0x6e, 0xa1, 0x0c, 0xee, 0xf4, 0xe2, 0x78, 0x75, 0xbc, 0xd6, 0x44, 0x2b, 0x2d, 0x92, 0xd6,
	0x89, 0x99, 0xcc, 0x68, 0x4c, 0x9d, 0xf1, 0x64, 0xa6, 0xf5, 0x66, 0xec, 0xb4, 0x0f, 0xc4, 0x25,
	0x6f, 0xc0, 0x2b, 0xf0, 0x02, 0xbc, 0x03, 0xf7, 0xdc, 0x33, 0xda, 0x5d, 0x3b, 0x8e, 0x13, 0x52,
	0xae, 0xa2, 0xf8, 0x7c, 0x9f, 0xf4, 0x9d, 0xef, 0x9c, 0x3d, 0x87, 0xbc, 0x18, 0xe6, 0xf9, 0xf0,
	0xd5, 0x61, 0xbd, 0x3f, 0x98, 0xd4, 0xcb, 0x63, 0x38, 0xbd, 0x6d, 0xd4, 0x0f, 0xc7, 0xe3, 0x7c,
	0x3c, 0xa9, 0xf7, 0x07, 0x7e, 0x38, 0xce, 0xdf, 0x1c, 0xfb, 0x97, 0xe3, 0xa3, 0x93, 0xc3, 0xf1,
	0x51, 0x3e, 0xf2, 0x45, 0xa4, 0x76, 0x3c, 0xce, 0x4f, 0x72, 0xb6, 0x51, 0x72, 0x6a, 0xfd, 0xc1,
	0xa4, 0x36, 0xa3, 0xd7, 0xde, 0x36, 0x6a, 0x25, 0x7d, 0x65, 0x6d, 0x7a, 0xfd, 0xf1, 0x51, 0xbd,
	0x3f, 0x1a, 0xe5, 0x27, 0xfd, 0x93, 0xa3, 0x7c, 0x34, 0x29, 0xd9, 0x5b, 0x7f, 0x2d, 0x93, 0x3b,
	0x30, 0x88, 0xc3, 0xfd, 0xd1, 0xf4, 0x7a, 0x0c, 0x44, 0x1c, 0xbd, 0x79, 0xbd, 0xf5, 0xc7, 0x32,
	0xf9, 0xe2, 0xc2, 0x28, 0xbb, 0x49, 0x6e, 0x64, 0xca, 0xa6, 0x18, 0x89, 0xb6, 0x40, 0x4e, 0xdf,
	0x63, 0x37, 0xc8, 0xf5, 0x4c, 0xed, 0x2a, 0xbd, 0xa7, 0xe8, 0x15, 0xb6, 0x4d, 0x1e, 0x02, 0xf7,
	0xb1, 0xd1, 0x59, 0xea, 0x23, 0x23, 0x1c, 0x1a, 0xa1, 0x95, 0x97, 0xd0, 0x42, 0xe9, 0xb9, 0x46,
	0xeb, 0x95, 0x76, 0x1e, 0xf7, 0x85, 0x75, 0x74, 0xe9, 0x52, 0x30, 0x48, 0x83, 0xc0, 0x7b, 0x25,
	0xd6, 0xd2, 0xab, 0xec, 0x31, 0x79, 0x10, 0x81, 0x0a, 0x74, 0xe0, 0xbc, 0x02, 0x39, 0xed, 0x15,
	0xc6, 0xe0, 0x44, 0x17, 0x4f, 0x2f, 0xa0, 0xef, 0xb3, 0x5b, 0xe4, 0x33, 0xa7, 0xb5, 0x4f, 0x40,
	0xf5, 0xbc, 0x4e, 0xd1, 0x80, 0x13, 0x5a, 0x59, 0xfa, 0x01, 0xbb, 0x4d, 0x3e, 0x8f, 0x40, 0x39,
	0x9f, 0xa5, 0x1c, 0x1c, 0xce, 0xc8, 0xf4, 0x1a, 0x5b, 0x21, 0x5f, 0x46, 0x5a, 0x45, 0x06, 0x1d,
	0x7a, 0xd7, 0x4b, 0xd1, 0x1b, 0xfc, 0x39, 0x13, 0x06, 0x39, 0xbd, 0xce, 0xee, 0x91, 0xf5, 0x96,
	0xe0, 0x5e, 0xa8, 0x48, 0x27, 0x29, 0x38, 0xd1, 0x92, 0xe8, 0xf7, 0x84, 0xeb, 0x78, 0xe0, 0x85,
	0x78, 0xfa, 0x21, 0x5b, 0x27, 0x77, 0x2a, 0x75, 0x0e, 0x4c, 0x8c, 0xce, 0x83, 0xe2, 0x1e, 0xf7,
	0x23, 0x99, 0x71, 0xa4, 0x1f, 0x05, 0xd3, 0x84, 0x94, 0x18, 0x83, 0xf4, 0x99, 0x91, 0x94, 0x04,
	0x21, 0x42, 0x75, 0x41, 0x0a, 0xee, 0x77, 0xb1, 0xb7, 0xa7, 0x0d, 0xf7, 0x0e, 0xf7, 0x1d, 0xbd,
	0xc1, 0x56, 0xc9, 0xad, 0x69, 0x84, 0xa3, 0x75, 0x42, 0x15, 0xe2, 0x0b, 0xda, 0x32, 0xdb, 0x24,
	0xab, 0x89, 0xb0, 0x56, 0xa8, 0x78, 0x31, 0xe8, 0x1d, 0xc4, 0xf4, 0x63, 0xd6, 0x20, 0x3b, 0xd3,
	0xfb, 0x24, 0x76, 0x51, 0xfa, 0x20, 0x3c, 0xc8, 0xb2, 0x59, 0x9a, 0x6a, 0xe3, 0x90, 0xfb, 0xb6,
	0x36, 0xc1, 0x99, 0x0c, 0x64, 0x94, 0x26, 0xf4, 0x93, 0x60, 0xd6, 0xf4, 0xc1, 0xcc, 0xa2, 0xf1,
	0xd6, 0x81, 0xcb, 0x2c, 0xbd, 0xc9, 0xd6, 0xc8, 0xed, 0x39, 0xc7, 0x2b, 0x7f, 0xa1, 0x70, 0x87,
	0x52, 0x76, 0x97, 0xac, 0x55, 0xd1, 0x2a, 0xcd, 0x05, 0xc4, 0xa7, 0xec, 0x07, 0xd2, 0x88, 0x20,
	0x49, 0x41, 0xc4, 0xaa, 0xb4, 0x34, 0x80, 0x17, 0x2d, 0x4c, 0xc1, 0x38, 0x01, 0xd2, 0xb7, 0x41,
	0xc8, 0xcc, 0x20, 0x5d, 0x65, 0x3b, 0xe4, 0xd1, 0x69, 0xcd, 0x0a, 0xb5, 0xb3, 0x5a, 0xda, 0x8e,
	0x4e, 0xd3, 0x90, 0x7c, 0x65, 0xbc, 0xa5, 0x6b, 0xec, 0x7b, 0xf2, 0x5d, 0xa5, 0x23, 0xd1, 0x5c,
	0xb4, 0x7b, 0x85, 0x19, 0x6d, 0x81, 0x92, 0xdb, 0xf2, 0x05, 0x9e, 0xa5, 0x52, 0x44, 0xa1, 0xdc,
	0x28, 0x31, 0x41, 0xe5, 0x2c, 0x5d, 0x0f, 0x25, 0xad, 0x58, 0x16, 0x5d, 0x81, 0xd3, 0x99, 0xf3,
	0x6d, 0xa1, 0xca, 0x0a, 0x59, 0xba, 0xc1, 0x9e, 0x91, 0x27, 0x15, 0x24, 0x92, 0x08, 0x66, 0x2e,
	0xe8, 0x45, 0xbb, 0xfa, 0x2f, 0xd1, 0x2d, 0x21, 0xb1, 0xfc, 0xb1, 0xec, 0xea, 0x4d, 0xf6, 0x94,
	0xd4, 0xdf, 0xc9, 0x83, 0x34, 0x9d, 0x27, 0xdd, 0x65, 0xcf, 0xc9, 0x8f, 0x97, 0x90, 0x9c, 0x81,
	0x68, 0x37, 0xe4, 0x5d, 0x14, 0x1a, 0x93, 0x54, 0x16, 0xe9, 0x94, 0x9f, 0xc6, 0x3d, 0xd6, 0x24,
	0xcf, 0x2e, 0x21, 0x07, 0x4e, 0x94, 0x59, 0xa7, 0x93, 0x60, 0x36, 0x24, 0xe8, 0xd0, 0x4c, 0x1f,
	0xde, 0x62, 0x4f, 0x48, 0x6d, 0xce, 0x88, 0x96, 0x0e, 0x86, 0x2d, 0x74, 0x56, 0x68, 0xe5, 0x39,
	0x67, 0xee, 0xb3, 0x17, 0xe4, 0xa7, 0xff, 0xc3, 0xb9, 0x50, 0x35, 0xfd, 0x8a, 0xd5, 0xc9, 0xf6,
	0x9c, 0xc2, 0xf3, 0xcd, 0x79, 0x3a, 0x0a, 0x8a, 0x46, 0xfa, 0x3a, 0x54, 0xe2, 0xbc, 0xdd, 0xef,
	0xe4, 0x3d, 0x60, 0xf7, 0xc9, 0xe6, 0xb4, 0xb3, 0xa5, 0x08, 0xfa, 0xe2, 0x6a, 0xd8, 0x74, 0x04,
	0x1a, 0x30, 0x51, 0xa7, 0x47, 0x1f, 0x86, 0x21, 0x74, 0x36, 0x98, 0x29, 0xe1, 0x7c, 0x95, 0x5f,
	0x07, 0xc2, 0x64, 0xe9, 0x08, 0xc9, 0x0d, 0x2a, 0xfa, 0x4d, 0x70, 0xeb, 0x2c, 0xd8, 0x66, 0x2d,
	0x2e, 0xba, 0xc2, 0x86, 0x57, 0xab, 0x89, 0x61, 0xbd, 0x76, 0x9d, 0x60, 0x71, 0x04, 0x16, 0xe9,
	0xa3, 0x50, 0x9d, 0xb3, 0x9c, 0x19, 0xce, 0x42, 0x82, 0x9e, 0x8b, 0x04, 0x95, 0x9d, 0xaa, 0xf6,
	0x60, 0xbd, 0x15, 0x2d, 0x29, 0x54, 0x6c, 0xe9, 0xe3, 0xf0, 0x91, 0x9d, 0xe5, 0x2e, 0x8c, 0xc5,
	0xed, 0xf3, 0x88, 0x85, 0x29, 0xfb, 0x6d, 0x70, 0xe1, 0x2c, 0xa2, 0xca, 0xad, 0x15, 0x66, 0x5c,
	0xa2, 0xbb, 0xc8, 0xe9, 0x0e, 0xdb, 0x20, 0x2b, 0x17, 0x5b, 0x55, 0x58, 0x59, 0x0b, 0xd3, 0x77,
	0x41, 0x08, 0xe7, 0x3e, 0x81, 0x9e, 0xd7, 0x4a, 0xf6, 0xc2, 0xdc, 0x28, 0x0a, 0xec, 0x05, 0xa7,
	0xf5, 0xd6, 0x3f, 0x57, 0xc8, 0xd6, 0xcb, 0xfc, 0x75, 0xed, 0xf2, 0xf5, 0xd4, 0x5a, 0xb9, 0x70,
	0xbf, 0xa4, 0x61, 0x39, 0xa5, 0x57, 0x7e, 0xe1, 0x15, 0x7b, 0x98, 0xbf, 0xea, 0x8f, 0x86, 0xb5,
	0x7c, 0x3c, 0xac, 0x0f, 0x0f, 0x47, 0xc5, 0xea, 0x9a, 0xee, 0xca, 0xe3, 0xa3, 0xc9, 0x7f, 0xad,
	0xce, 0xe7, 0xe5, 0x9f, 0xdf, 0x96, 0xae, 0xc6, 0x00, 0xbf, 0x2f, 0x6d, 0xc4, 0xe5, 0x65, 0x30,
	0x98, 0xd4, 0xca, 0x63, 0x38, 0x75, 0x1b, 0xb5, 0xe2, 0xc9, 0xc9, 0x9f, 0x53, 0xc0, 0x01, 0x0c,
	0x26, 0x07, 0x33, 0xc0, 0x41, 0xb7, 0x71, 0x50, 0x02, 0xfe, 0x5e, 0xda, 0x2a, 0x7f, 0x6d, 0x36,
	0x61, 0x30, 0x69, 0x36, 0x67, 0x90, 0x66, 0xb3, 0xdb, 0x68, 0x36, 0x4b, 0xd0, 0xaf, 0xd7, 0x0a,
	0x75, 0x4f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xf6, 0x6e, 0xef, 0xd7, 0x07, 0x00, 0x00,
}
