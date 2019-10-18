// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/tags.proto

package datacatalog

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FieldType_PrimitiveType int32

const (
	// This is the default invalid value for a type.
	FieldType_PRIMITIVE_TYPE_UNSPECIFIED FieldType_PrimitiveType = 0
	// A double precision number.
	FieldType_DOUBLE FieldType_PrimitiveType = 1
	// An UTF-8 string.
	FieldType_STRING FieldType_PrimitiveType = 2
	// A boolean value.
	FieldType_BOOL FieldType_PrimitiveType = 3
	// A timestamp.
	FieldType_TIMESTAMP FieldType_PrimitiveType = 4
)

var FieldType_PrimitiveType_name = map[int32]string{
	0: "PRIMITIVE_TYPE_UNSPECIFIED",
	1: "DOUBLE",
	2: "STRING",
	3: "BOOL",
	4: "TIMESTAMP",
}

var FieldType_PrimitiveType_value = map[string]int32{
	"PRIMITIVE_TYPE_UNSPECIFIED": 0,
	"DOUBLE":                     1,
	"STRING":                     2,
	"BOOL":                       3,
	"TIMESTAMP":                  4,
}

func (x FieldType_PrimitiveType) String() string {
	return proto.EnumName(FieldType_PrimitiveType_name, int32(x))
}

func (FieldType_PrimitiveType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{4, 0}
}

// Tags are used to attach custom metadata to Data Catalog resources. Tags
// conform to the specifications within their tag template.
type Tag struct {
	// Required when used in
	// [UpdateTagRequest][google.cloud.datacatalog.v1beta1.UpdateTagRequest]. The
	// resource name of the tag in URL format. Example:
	//
	// * projects/{project_id}/locations/{location}/entrygroups/{entry_group_id}/entries/{entry_id}/tags/{tag_id}
	//
	// where `tag_id` is a system-generated identifier.
	// Note that this Tag may not actually be stored in the location in this name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The resource name of the tag template that this tag uses.
	// Example:
	//
	// * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id}
	//
	// This field cannot be modified after creation.
	Template string `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Output only. The display name of the tag template.
	TemplateDisplayName string `protobuf:"bytes,5,opt,name=template_display_name,json=templateDisplayName,proto3" json:"template_display_name,omitempty"`
	// Optional. The scope within the parent resource that this tag is attached
	// to. If not provided, the tag is attached to the parent resource itself.
	// Deleting the scope from the parent resource will delete all tags attached
	// to that scope. These fields cannot be updated after creation.
	//
	// Types that are valid to be assigned to Scope:
	//	*Tag_Column
	Scope isTag_Scope `protobuf_oneof:"scope"`
	// Required. This maps the ID of a tag field to the value of and additional
	// information about that field. Valid field IDs are defined by the tag's
	// template. A tag must have at least 1 field and at most 500 fields.
	Fields               map[string]*TagField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{0}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *Tag) GetTemplateDisplayName() string {
	if m != nil {
		return m.TemplateDisplayName
	}
	return ""
}

type isTag_Scope interface {
	isTag_Scope()
}

type Tag_Column struct {
	Column string `protobuf:"bytes,4,opt,name=column,proto3,oneof"`
}

func (*Tag_Column) isTag_Scope() {}

func (m *Tag) GetScope() isTag_Scope {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *Tag) GetColumn() string {
	if x, ok := m.GetScope().(*Tag_Column); ok {
		return x.Column
	}
	return ""
}

func (m *Tag) GetFields() map[string]*TagField {
	if m != nil {
		return m.Fields
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Tag) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Tag_Column)(nil),
	}
}

// Contains the value and supporting information for a field within
// a [Tag][google.cloud.datacatalog.v1beta1.Tag].
type TagField struct {
	// Output only. The display name of this field.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The value of this field.
	//
	// Types that are valid to be assigned to Kind:
	//	*TagField_DoubleValue
	//	*TagField_StringValue
	//	*TagField_BoolValue
	//	*TagField_TimestampValue
	//	*TagField_EnumValue_
	Kind                 isTagField_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TagField) Reset()         { *m = TagField{} }
func (m *TagField) String() string { return proto.CompactTextString(m) }
func (*TagField) ProtoMessage()    {}
func (*TagField) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{1}
}

func (m *TagField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagField.Unmarshal(m, b)
}
func (m *TagField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagField.Marshal(b, m, deterministic)
}
func (m *TagField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagField.Merge(m, src)
}
func (m *TagField) XXX_Size() int {
	return xxx_messageInfo_TagField.Size(m)
}
func (m *TagField) XXX_DiscardUnknown() {
	xxx_messageInfo_TagField.DiscardUnknown(m)
}

var xxx_messageInfo_TagField proto.InternalMessageInfo

func (m *TagField) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type isTagField_Kind interface {
	isTagField_Kind()
}

type TagField_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,2,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type TagField_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type TagField_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type TagField_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,5,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type TagField_EnumValue_ struct {
	EnumValue *TagField_EnumValue `protobuf:"bytes,6,opt,name=enum_value,json=enumValue,proto3,oneof"`
}

func (*TagField_DoubleValue) isTagField_Kind() {}

func (*TagField_StringValue) isTagField_Kind() {}

func (*TagField_BoolValue) isTagField_Kind() {}

func (*TagField_TimestampValue) isTagField_Kind() {}

func (*TagField_EnumValue_) isTagField_Kind() {}

func (m *TagField) GetKind() isTagField_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *TagField) GetDoubleValue() float64 {
	if x, ok := m.GetKind().(*TagField_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *TagField) GetStringValue() string {
	if x, ok := m.GetKind().(*TagField_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *TagField) GetBoolValue() bool {
	if x, ok := m.GetKind().(*TagField_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *TagField) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetKind().(*TagField_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *TagField) GetEnumValue() *TagField_EnumValue {
	if x, ok := m.GetKind().(*TagField_EnumValue_); ok {
		return x.EnumValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagField) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagField_DoubleValue)(nil),
		(*TagField_StringValue)(nil),
		(*TagField_BoolValue)(nil),
		(*TagField_TimestampValue)(nil),
		(*TagField_EnumValue_)(nil),
	}
}

// Holds an enum value.
type TagField_EnumValue struct {
	// The display name of the enum value.
	DisplayName          string   `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagField_EnumValue) Reset()         { *m = TagField_EnumValue{} }
func (m *TagField_EnumValue) String() string { return proto.CompactTextString(m) }
func (*TagField_EnumValue) ProtoMessage()    {}
func (*TagField_EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{1, 0}
}

func (m *TagField_EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagField_EnumValue.Unmarshal(m, b)
}
func (m *TagField_EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagField_EnumValue.Marshal(b, m, deterministic)
}
func (m *TagField_EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagField_EnumValue.Merge(m, src)
}
func (m *TagField_EnumValue) XXX_Size() int {
	return xxx_messageInfo_TagField_EnumValue.Size(m)
}
func (m *TagField_EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TagField_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_TagField_EnumValue proto.InternalMessageInfo

func (m *TagField_EnumValue) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// A tag template defines the schema of the tags used to attach to Data Catalog
// resources. It defines the mapping of accepted field names and types that can
// be used within the tag. The tag template also controls the access to the tag.
type TagTemplate struct {
	// Required when used in
	// [UpdateTagTemplateRequest][google.cloud.datacatalog.v1beta1.UpdateTagTemplateRequest].
	// The resource name of the tag template in URL format. Example:
	//
	// * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id}
	//
	// Note that this TagTemplate and its child resources may not actually be
	// stored in the location in this name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The display name for this template. Defaults to an empty string.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. Map of tag template field IDs to the settings for the field.
	// This map is an exhaustive list of the allowed fields. This map must contain
	// at least one field and at most 500 fields.
	//
	// The keys to this map are tag template field IDs. Field IDs can contain
	// letters (both uppercase and lowercase), numbers (0-9) and underscores (_).
	// Field IDs must be at least 1 character long and at most
	// 64 characters long. Field IDs must start with a letter or underscore.
	Fields               map[string]*TagTemplateField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *TagTemplate) Reset()         { *m = TagTemplate{} }
func (m *TagTemplate) String() string { return proto.CompactTextString(m) }
func (*TagTemplate) ProtoMessage()    {}
func (*TagTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{2}
}

func (m *TagTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagTemplate.Unmarshal(m, b)
}
func (m *TagTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagTemplate.Marshal(b, m, deterministic)
}
func (m *TagTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagTemplate.Merge(m, src)
}
func (m *TagTemplate) XXX_Size() int {
	return xxx_messageInfo_TagTemplate.Size(m)
}
func (m *TagTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_TagTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_TagTemplate proto.InternalMessageInfo

func (m *TagTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TagTemplate) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TagTemplate) GetFields() map[string]*TagTemplateField {
	if m != nil {
		return m.Fields
	}
	return nil
}

// The template for an individual field within a tag template.
type TagTemplateField struct {
	// Output only. The resource name of the tag template field in URL format.
	// Example:
	//
	// * projects/{project_id}/locations/{location}/tagTemplates/{tag_template}/fields/{field}
	//
	// Note that this TagTemplateField may not actually be stored in the location
	// in this name.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The display name for this field. Defaults to an empty string.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The type of value this tag field can contain.
	Type                 *FieldType `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TagTemplateField) Reset()         { *m = TagTemplateField{} }
func (m *TagTemplateField) String() string { return proto.CompactTextString(m) }
func (*TagTemplateField) ProtoMessage()    {}
func (*TagTemplateField) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{3}
}

func (m *TagTemplateField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagTemplateField.Unmarshal(m, b)
}
func (m *TagTemplateField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagTemplateField.Marshal(b, m, deterministic)
}
func (m *TagTemplateField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagTemplateField.Merge(m, src)
}
func (m *TagTemplateField) XXX_Size() int {
	return xxx_messageInfo_TagTemplateField.Size(m)
}
func (m *TagTemplateField) XXX_DiscardUnknown() {
	xxx_messageInfo_TagTemplateField.DiscardUnknown(m)
}

var xxx_messageInfo_TagTemplateField proto.InternalMessageInfo

func (m *TagTemplateField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TagTemplateField) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TagTemplateField) GetType() *FieldType {
	if m != nil {
		return m.Type
	}
	return nil
}

type FieldType struct {
	// Required.
	//
	// Types that are valid to be assigned to TypeDecl:
	//	*FieldType_PrimitiveType_
	//	*FieldType_EnumType_
	TypeDecl             isFieldType_TypeDecl `protobuf_oneof:"type_decl"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FieldType) Reset()         { *m = FieldType{} }
func (m *FieldType) String() string { return proto.CompactTextString(m) }
func (*FieldType) ProtoMessage()    {}
func (*FieldType) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{4}
}

func (m *FieldType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldType.Unmarshal(m, b)
}
func (m *FieldType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldType.Marshal(b, m, deterministic)
}
func (m *FieldType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldType.Merge(m, src)
}
func (m *FieldType) XXX_Size() int {
	return xxx_messageInfo_FieldType.Size(m)
}
func (m *FieldType) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldType.DiscardUnknown(m)
}

var xxx_messageInfo_FieldType proto.InternalMessageInfo

type isFieldType_TypeDecl interface {
	isFieldType_TypeDecl()
}

type FieldType_PrimitiveType_ struct {
	PrimitiveType FieldType_PrimitiveType `protobuf:"varint,1,opt,name=primitive_type,json=primitiveType,proto3,enum=google.cloud.datacatalog.v1beta1.FieldType_PrimitiveType,oneof"`
}

type FieldType_EnumType_ struct {
	EnumType *FieldType_EnumType `protobuf:"bytes,2,opt,name=enum_type,json=enumType,proto3,oneof"`
}

func (*FieldType_PrimitiveType_) isFieldType_TypeDecl() {}

func (*FieldType_EnumType_) isFieldType_TypeDecl() {}

func (m *FieldType) GetTypeDecl() isFieldType_TypeDecl {
	if m != nil {
		return m.TypeDecl
	}
	return nil
}

func (m *FieldType) GetPrimitiveType() FieldType_PrimitiveType {
	if x, ok := m.GetTypeDecl().(*FieldType_PrimitiveType_); ok {
		return x.PrimitiveType
	}
	return FieldType_PRIMITIVE_TYPE_UNSPECIFIED
}

func (m *FieldType) GetEnumType() *FieldType_EnumType {
	if x, ok := m.GetTypeDecl().(*FieldType_EnumType_); ok {
		return x.EnumType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FieldType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FieldType_PrimitiveType_)(nil),
		(*FieldType_EnumType_)(nil),
	}
}

type FieldType_EnumType struct {
	// Required on create; optional on update. The set of allowed values for
	// this enum. This set must not be empty, the display names of the values in
	// this set must not be empty and the display names of the values must be
	// case-insensitively unique within this set. Currently, enum values can
	// only be added to the list of allowed values. Deletion and renaming of
	// enum values are not supported. Can have up to 500 allowed values.
	AllowedValues        []*FieldType_EnumType_EnumValue `protobuf:"bytes,1,rep,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FieldType_EnumType) Reset()         { *m = FieldType_EnumType{} }
func (m *FieldType_EnumType) String() string { return proto.CompactTextString(m) }
func (*FieldType_EnumType) ProtoMessage()    {}
func (*FieldType_EnumType) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{4, 0}
}

func (m *FieldType_EnumType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldType_EnumType.Unmarshal(m, b)
}
func (m *FieldType_EnumType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldType_EnumType.Marshal(b, m, deterministic)
}
func (m *FieldType_EnumType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldType_EnumType.Merge(m, src)
}
func (m *FieldType_EnumType) XXX_Size() int {
	return xxx_messageInfo_FieldType_EnumType.Size(m)
}
func (m *FieldType_EnumType) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldType_EnumType.DiscardUnknown(m)
}

var xxx_messageInfo_FieldType_EnumType proto.InternalMessageInfo

func (m *FieldType_EnumType) GetAllowedValues() []*FieldType_EnumType_EnumValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

type FieldType_EnumType_EnumValue struct {
	// Required. The display name of the enum value. Must not be an empty
	// string.
	DisplayName          string   `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldType_EnumType_EnumValue) Reset()         { *m = FieldType_EnumType_EnumValue{} }
func (m *FieldType_EnumType_EnumValue) String() string { return proto.CompactTextString(m) }
func (*FieldType_EnumType_EnumValue) ProtoMessage()    {}
func (*FieldType_EnumType_EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd303c40f581309, []int{4, 0, 0}
}

func (m *FieldType_EnumType_EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldType_EnumType_EnumValue.Unmarshal(m, b)
}
func (m *FieldType_EnumType_EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldType_EnumType_EnumValue.Marshal(b, m, deterministic)
}
func (m *FieldType_EnumType_EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldType_EnumType_EnumValue.Merge(m, src)
}
func (m *FieldType_EnumType_EnumValue) XXX_Size() int {
	return xxx_messageInfo_FieldType_EnumType_EnumValue.Size(m)
}
func (m *FieldType_EnumType_EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldType_EnumType_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_FieldType_EnumType_EnumValue proto.InternalMessageInfo

func (m *FieldType_EnumType_EnumValue) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.datacatalog.v1beta1.FieldType_PrimitiveType", FieldType_PrimitiveType_name, FieldType_PrimitiveType_value)
	proto.RegisterType((*Tag)(nil), "google.cloud.datacatalog.v1beta1.Tag")
	proto.RegisterMapType((map[string]*TagField)(nil), "google.cloud.datacatalog.v1beta1.Tag.FieldsEntry")
	proto.RegisterType((*TagField)(nil), "google.cloud.datacatalog.v1beta1.TagField")
	proto.RegisterType((*TagField_EnumValue)(nil), "google.cloud.datacatalog.v1beta1.TagField.EnumValue")
	proto.RegisterType((*TagTemplate)(nil), "google.cloud.datacatalog.v1beta1.TagTemplate")
	proto.RegisterMapType((map[string]*TagTemplateField)(nil), "google.cloud.datacatalog.v1beta1.TagTemplate.FieldsEntry")
	proto.RegisterType((*TagTemplateField)(nil), "google.cloud.datacatalog.v1beta1.TagTemplateField")
	proto.RegisterType((*FieldType)(nil), "google.cloud.datacatalog.v1beta1.FieldType")
	proto.RegisterType((*FieldType_EnumType)(nil), "google.cloud.datacatalog.v1beta1.FieldType.EnumType")
	proto.RegisterType((*FieldType_EnumType_EnumValue)(nil), "google.cloud.datacatalog.v1beta1.FieldType.EnumType.EnumValue")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/tags.proto", fileDescriptor_6fd303c40f581309)
}

var fileDescriptor_6fd303c40f581309 = []byte{
	// 901 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x8f, 0xed, 0x34, 0x24, 0x2f, 0xdb, 0x12, 0x0d, 0x42, 0x0a, 0x11, 0xda, 0x96, 0x20, 0xad,
	0x56, 0x54, 0xb2, 0xd5, 0xee, 0x4a, 0x40, 0x91, 0x10, 0xcd, 0xd6, 0xbb, 0x89, 0xd8, 0xb6, 0x91,
	0xeb, 0x56, 0x02, 0x21, 0x99, 0x89, 0x3d, 0x6b, 0xcc, 0xda, 0x1e, 0xcb, 0x7f, 0x8a, 0xa2, 0x28,
	0x48, 0x7b, 0xe0, 0x33, 0xf0, 0x0d, 0xf8, 0x3e, 0xdc, 0x39, 0x70, 0xee, 0x95, 0x0b, 0x17, 0x24,
	0x34, 0x33, 0x9e, 0xc4, 0x69, 0xb7, 0xdb, 0x2c, 0xe2, 0x94, 0x79, 0x6f, 0x7e, 0xef, 0x37, 0x6f,
	0x7e, 0xef, 0xbd, 0x89, 0x61, 0xd7, 0xa7, 0xd4, 0x0f, 0x89, 0xe1, 0x86, 0xb4, 0xf0, 0x0c, 0x0f,
	0xe7, 0xd8, 0xc5, 0x39, 0x0e, 0xa9, 0x6f, 0x5c, 0xee, 0x4d, 0x48, 0x8e, 0xf7, 0x8c, 0x1c, 0xfb,
	0x99, 0x9e, 0xa4, 0x34, 0xa7, 0x68, 0x47, 0x80, 0x75, 0x0e, 0xd6, 0x2b, 0x60, 0xbd, 0x04, 0xf7,
	0xb6, 0x4b, 0x3a, 0x9c, 0x04, 0xc6, 0x8b, 0x80, 0x84, 0x9e, 0x33, 0x21, 0x3f, 0xe0, 0xcb, 0x80,
	0xa6, 0x82, 0xa2, 0xf7, 0x41, 0x05, 0x90, 0x92, 0x8c, 0x16, 0xa9, 0x4b, 0xca, 0x2d, 0x19, 0xcb,
	0xad, 0x49, 0xf1, 0xc2, 0xc8, 0x83, 0x88, 0x64, 0x39, 0x8e, 0x12, 0x01, 0xe8, 0xff, 0xae, 0x81,
	0x66, 0x63, 0x1f, 0x21, 0xa8, 0xc7, 0x38, 0x22, 0x5d, 0x65, 0x47, 0x79, 0xd8, 0xb2, 0xf8, 0x1a,
	0x6d, 0x43, 0x33, 0x27, 0x51, 0x12, 0xe2, 0x9c, 0x74, 0x55, 0xe6, 0x1f, 0x68, 0x7f, 0x1e, 0xaa,
	0xd6, 0xc2, 0x89, 0x3e, 0x85, 0xf7, 0xe5, 0xda, 0xf1, 0x82, 0x2c, 0x09, 0xf1, 0xd4, 0xe1, 0x2c,
	0x1b, 0x12, 0xad, 0x59, 0xef, 0x49, 0xc4, 0x91, 0x00, 0x9c, 0x30, 0xe6, 0x2e, 0x34, 0x5c, 0x1a,
	0x16, 0x51, 0xdc, 0xad, 0x33, 0xe4, 0xb0, 0x66, 0x95, 0x36, 0x3a, 0x81, 0x06, 0xbf, 0x63, 0xd6,
	0xd5, 0x76, 0xb4, 0x87, 0xed, 0xfd, 0x3d, 0xfd, 0x2e, 0x7d, 0x74, 0x1b, 0xfb, 0xfa, 0x53, 0x1e,
	0x63, 0xc6, 0x79, 0x3a, 0x15, 0x49, 0x96, 0x2c, 0x3d, 0x02, 0xed, 0xca, 0x1e, 0xea, 0x80, 0xf6,
	0x92, 0x4c, 0xcb, 0x5b, 0xb2, 0x25, 0xfa, 0x0a, 0x36, 0x2e, 0x71, 0x58, 0x88, 0x1b, 0xb6, 0xf7,
	0x3f, 0x59, 0xeb, 0x3c, 0x4e, 0x69, 0x89, 0xc0, 0x03, 0xf5, 0x33, 0xe5, 0xe0, 0x95, 0x72, 0x75,
	0xf8, 0x33, 0xdc, 0xaf, 0xc2, 0x05, 0x0f, 0x4e, 0x82, 0x4c, 0x77, 0x69, 0x64, 0x30, 0x91, 0xbf,
	0x4b, 0x52, 0xfa, 0x23, 0x71, 0xf3, 0xcc, 0x98, 0x95, 0xab, 0xb9, 0x11, 0x52, 0x17, 0xe7, 0x01,
	0x8d, 0x33, 0x63, 0x26, 0x97, 0x73, 0x83, 0xb0, 0x64, 0x9f, 0xa5, 0xb4, 0x48, 0x32, 0x63, 0xc6,
	0x0d, 0xc7, 0x67, 0x96, 0xd8, 0x09, 0x88, 0xf4, 0xce, 0x79, 0x27, 0x19, 0xb3, 0x1c, 0xfb, 0xf3,
	0xc1, 0x3b, 0xb0, 0x91, 0xb9, 0x34, 0x21, 0xfd, 0xbf, 0x54, 0x68, 0xca, 0x24, 0xd1, 0x03, 0xb8,
	0xb7, 0x52, 0x1a, 0x65, 0x59, 0x9a, 0xb6, 0x57, 0x29, 0xc9, 0xc7, 0x70, 0xcf, 0xa3, 0xc5, 0x24,
	0x24, 0xce, 0x52, 0x0e, 0x65, 0x58, 0xb3, 0xda, 0xc2, 0x7b, 0xc1, 0x9c, 0x0c, 0x94, 0xe5, 0x69,
	0x10, 0xfb, 0x25, 0x48, 0x2b, 0xab, 0xd7, 0x16, 0x5e, 0x01, 0xda, 0x06, 0x98, 0x50, 0x1a, 0x96,
	0x10, 0x56, 0xe0, 0xe6, 0xb0, 0x66, 0xb5, 0x98, 0x4f, 0x00, 0x4c, 0x78, 0x77, 0xd1, 0x86, 0x25,
	0x6a, 0x83, 0x8b, 0xdf, 0x93, 0xe2, 0xcb, 0x76, 0xd5, 0x6d, 0x89, 0x1b, 0xd6, 0xac, 0xad, 0x45,
	0x90, 0xa0, 0x39, 0x07, 0x20, 0x71, 0x11, 0x95, 0x0c, 0x0d, 0xce, 0xf0, 0x78, 0xfd, 0xf2, 0xe9,
	0x66, 0x5c, 0x44, 0x9c, 0x89, 0x65, 0x47, 0xa4, 0xd1, 0xd3, 0xa1, 0xb5, 0xd8, 0x41, 0x1f, 0xbd,
	0x4e, 0xbd, 0x15, 0xe1, 0x06, 0x0d, 0xa8, 0xbf, 0x0c, 0x62, 0xaf, 0xff, 0x8f, 0x0a, 0x6d, 0x1b,
	0xfb, 0xb6, 0x1c, 0x8e, 0xd7, 0x4d, 0xd4, 0x75, 0x3a, 0xf5, 0x06, 0x1d, 0xba, 0xb8, 0x36, 0x00,
	0x9f, 0xaf, 0x75, 0x23, 0x79, 0xea, 0xed, 0x83, 0x10, 0xdd, 0x35, 0x08, 0xc3, 0xd5, 0x41, 0xd8,
	0x7f, 0xab, 0x73, 0x6f, 0x0c, 0x44, 0x72, 0x75, 0x18, 0xc1, 0x83, 0x37, 0xcf, 0xc3, 0x42, 0xaa,
	0x27, 0xeb, 0xce, 0x45, 0xbe, 0x0c, 0x12, 0x5d, 0xef, 0xc8, 0xd7, 0x65, 0xde, 0xff, 0x55, 0x85,
	0xce, 0xf5, 0x8c, 0x16, 0x45, 0x68, 0xbc, 0xa1, 0x08, 0x37, 0x6b, 0x8a, 0x8e, 0xa0, 0x9e, 0x4f,
	0x13, 0x29, 0xc5, 0xee, 0xdd, 0x52, 0xf0, 0xd3, 0xec, 0x69, 0x42, 0x84, 0xe8, 0x3c, 0xfa, 0xe0,
	0x17, 0xe5, 0xea, 0xf0, 0x95, 0x02, 0xbb, 0xeb, 0xa9, 0x20, 0x12, 0xb6, 0xfe, 0x07, 0x29, 0xc4,
	0xbf, 0x44, 0x66, 0xcc, 0xf8, 0xef, 0xbc, 0xff, 0x87, 0x06, 0xad, 0x45, 0x82, 0x68, 0x02, 0x5b,
	0x49, 0x1a, 0x44, 0x41, 0x1e, 0x5c, 0x12, 0x87, 0xdf, 0x92, 0x09, 0xb0, 0xb5, 0x4e, 0xa3, 0x2d,
	0x48, 0xf4, 0xb1, 0x64, 0x60, 0xd6, 0xb0, 0x66, 0x6d, 0x26, 0x55, 0x07, 0x3a, 0x03, 0x3e, 0x50,
	0x4e, 0x45, 0xc4, 0xc7, 0x6f, 0x43, 0xcf, 0x06, 0xb0, 0x64, 0x6e, 0x92, 0x72, 0xdd, 0xfb, 0x4d,
	0x81, 0xa6, 0xdc, 0x40, 0x04, 0xb6, 0x70, 0x18, 0xd2, 0x9f, 0x88, 0x27, 0xe6, 0x3f, 0xeb, 0x2a,
	0x7c, 0x5c, 0xbe, 0xfc, 0x2f, 0xc7, 0x2c, 0x9f, 0x02, 0x6b, 0xb3, 0x64, 0xe5, 0x56, 0xd6, 0x7b,
	0x54, 0x7d, 0x0c, 0x6e, 0x7d, 0x4a, 0xd5, 0x95, 0xee, 0xe9, 0x7f, 0x0f, 0x9b, 0x2b, 0xfa, 0xa0,
	0xfb, 0xd0, 0x1b, 0x5b, 0xa3, 0xe3, 0x91, 0x3d, 0xba, 0x30, 0x1d, 0xfb, 0x9b, 0xb1, 0xe9, 0x9c,
	0x9f, 0x9c, 0x8d, 0xcd, 0x27, 0xa3, 0xa7, 0x23, 0xf3, 0xa8, 0x53, 0x43, 0x00, 0x8d, 0xa3, 0xd3,
	0xf3, 0xc1, 0x73, 0xb3, 0xa3, 0xb0, 0xf5, 0x99, 0x6d, 0x8d, 0x4e, 0x9e, 0x75, 0x54, 0xd4, 0x84,
	0xfa, 0xe0, 0xf4, 0xf4, 0x79, 0x47, 0x43, 0x9b, 0xd0, 0xb2, 0x47, 0xc7, 0xe6, 0x99, 0x7d, 0x78,
	0x3c, 0xee, 0xd4, 0x07, 0x6d, 0x68, 0x31, 0x69, 0x1d, 0x8f, 0xb8, 0xe1, 0x20, 0x81, 0x0f, 0x5d,
	0x1a, 0xdd, 0x7a, 0xef, 0xb1, 0xf2, 0xed, 0xd7, 0xe5, 0x9e, 0x4f, 0x43, 0x1c, 0xfb, 0x3a, 0x4d,
	0x7d, 0xc3, 0x27, 0x31, 0x7f, 0x64, 0x8d, 0x65, 0x43, 0xde, 0xfe, 0xbd, 0xf2, 0x45, 0xc5, 0xf7,
	0xb7, 0xa2, 0x4c, 0x1a, 0x3c, 0xf4, 0xd1, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0x61, 0xbb,
	0x1d, 0xe9, 0x08, 0x00, 0x00,
}