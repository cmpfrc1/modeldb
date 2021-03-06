// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/public/modeldb/versioning/Code.proto

package versioning

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CodeBlob struct {
	// Types that are valid to be assigned to Content:
	//	*CodeBlob_Git
	//	*CodeBlob_Notebook
	Content              isCodeBlob_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CodeBlob) Reset()         { *m = CodeBlob{} }
func (m *CodeBlob) String() string { return proto.CompactTextString(m) }
func (*CodeBlob) ProtoMessage()    {}
func (*CodeBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d8129ea3c2050d, []int{0}
}

func (m *CodeBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeBlob.Unmarshal(m, b)
}
func (m *CodeBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeBlob.Marshal(b, m, deterministic)
}
func (m *CodeBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeBlob.Merge(m, src)
}
func (m *CodeBlob) XXX_Size() int {
	return xxx_messageInfo_CodeBlob.Size(m)
}
func (m *CodeBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeBlob.DiscardUnknown(m)
}

var xxx_messageInfo_CodeBlob proto.InternalMessageInfo

type isCodeBlob_Content interface {
	isCodeBlob_Content()
}

type CodeBlob_Git struct {
	Git *GitCodeBlob `protobuf:"bytes,1,opt,name=git,proto3,oneof"`
}

type CodeBlob_Notebook struct {
	Notebook *NotebookCodeBlob `protobuf:"bytes,2,opt,name=notebook,proto3,oneof"`
}

func (*CodeBlob_Git) isCodeBlob_Content() {}

func (*CodeBlob_Notebook) isCodeBlob_Content() {}

func (m *CodeBlob) GetContent() isCodeBlob_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *CodeBlob) GetGit() *GitCodeBlob {
	if x, ok := m.GetContent().(*CodeBlob_Git); ok {
		return x.Git
	}
	return nil
}

func (m *CodeBlob) GetNotebook() *NotebookCodeBlob {
	if x, ok := m.GetContent().(*CodeBlob_Notebook); ok {
		return x.Notebook
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CodeBlob) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CodeBlob_Git)(nil),
		(*CodeBlob_Notebook)(nil),
	}
}

type GitCodeBlob struct {
	Repo                 string   `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Branch               string   `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	Tag                  string   `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	IsDirty              bool     `protobuf:"varint,5,opt,name=is_dirty,json=isDirty,proto3" json:"is_dirty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitCodeBlob) Reset()         { *m = GitCodeBlob{} }
func (m *GitCodeBlob) String() string { return proto.CompactTextString(m) }
func (*GitCodeBlob) ProtoMessage()    {}
func (*GitCodeBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d8129ea3c2050d, []int{1}
}

func (m *GitCodeBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitCodeBlob.Unmarshal(m, b)
}
func (m *GitCodeBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitCodeBlob.Marshal(b, m, deterministic)
}
func (m *GitCodeBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitCodeBlob.Merge(m, src)
}
func (m *GitCodeBlob) XXX_Size() int {
	return xxx_messageInfo_GitCodeBlob.Size(m)
}
func (m *GitCodeBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_GitCodeBlob.DiscardUnknown(m)
}

var xxx_messageInfo_GitCodeBlob proto.InternalMessageInfo

func (m *GitCodeBlob) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *GitCodeBlob) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *GitCodeBlob) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *GitCodeBlob) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *GitCodeBlob) GetIsDirty() bool {
	if m != nil {
		return m.IsDirty
	}
	return false
}

type NotebookCodeBlob struct {
	Path                 *PathDatasetBlob `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	GitRepo              *GitCodeBlob     `protobuf:"bytes,2,opt,name=git_repo,json=gitRepo,proto3" json:"git_repo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NotebookCodeBlob) Reset()         { *m = NotebookCodeBlob{} }
func (m *NotebookCodeBlob) String() string { return proto.CompactTextString(m) }
func (*NotebookCodeBlob) ProtoMessage()    {}
func (*NotebookCodeBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d8129ea3c2050d, []int{2}
}

func (m *NotebookCodeBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookCodeBlob.Unmarshal(m, b)
}
func (m *NotebookCodeBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookCodeBlob.Marshal(b, m, deterministic)
}
func (m *NotebookCodeBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookCodeBlob.Merge(m, src)
}
func (m *NotebookCodeBlob) XXX_Size() int {
	return xxx_messageInfo_NotebookCodeBlob.Size(m)
}
func (m *NotebookCodeBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookCodeBlob.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookCodeBlob proto.InternalMessageInfo

func (m *NotebookCodeBlob) GetPath() *PathDatasetBlob {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *NotebookCodeBlob) GetGitRepo() *GitCodeBlob {
	if m != nil {
		return m.GitRepo
	}
	return nil
}

type CodeDiff struct {
	// Types that are valid to be assigned to Content:
	//	*CodeDiff_Git
	//	*CodeDiff_Notebook
	Content              isCodeDiff_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CodeDiff) Reset()         { *m = CodeDiff{} }
func (m *CodeDiff) String() string { return proto.CompactTextString(m) }
func (*CodeDiff) ProtoMessage()    {}
func (*CodeDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d8129ea3c2050d, []int{3}
}

func (m *CodeDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeDiff.Unmarshal(m, b)
}
func (m *CodeDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeDiff.Marshal(b, m, deterministic)
}
func (m *CodeDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeDiff.Merge(m, src)
}
func (m *CodeDiff) XXX_Size() int {
	return xxx_messageInfo_CodeDiff.Size(m)
}
func (m *CodeDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeDiff.DiscardUnknown(m)
}

var xxx_messageInfo_CodeDiff proto.InternalMessageInfo

type isCodeDiff_Content interface {
	isCodeDiff_Content()
}

type CodeDiff_Git struct {
	Git *GitCodeDiff `protobuf:"bytes,1,opt,name=git,proto3,oneof"`
}

type CodeDiff_Notebook struct {
	Notebook *NotebookCodeDiff `protobuf:"bytes,2,opt,name=notebook,proto3,oneof"`
}

func (*CodeDiff_Git) isCodeDiff_Content() {}

func (*CodeDiff_Notebook) isCodeDiff_Content() {}

func (m *CodeDiff) GetContent() isCodeDiff_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *CodeDiff) GetGit() *GitCodeDiff {
	if x, ok := m.GetContent().(*CodeDiff_Git); ok {
		return x.Git
	}
	return nil
}

func (m *CodeDiff) GetNotebook() *NotebookCodeDiff {
	if x, ok := m.GetContent().(*CodeDiff_Notebook); ok {
		return x.Notebook
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CodeDiff) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CodeDiff_Git)(nil),
		(*CodeDiff_Notebook)(nil),
	}
}

type GitCodeDiff struct {
	A                    *GitCodeBlob `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    *GitCodeBlob `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GitCodeDiff) Reset()         { *m = GitCodeDiff{} }
func (m *GitCodeDiff) String() string { return proto.CompactTextString(m) }
func (*GitCodeDiff) ProtoMessage()    {}
func (*GitCodeDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d8129ea3c2050d, []int{4}
}

func (m *GitCodeDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitCodeDiff.Unmarshal(m, b)
}
func (m *GitCodeDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitCodeDiff.Marshal(b, m, deterministic)
}
func (m *GitCodeDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitCodeDiff.Merge(m, src)
}
func (m *GitCodeDiff) XXX_Size() int {
	return xxx_messageInfo_GitCodeDiff.Size(m)
}
func (m *GitCodeDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_GitCodeDiff.DiscardUnknown(m)
}

var xxx_messageInfo_GitCodeDiff proto.InternalMessageInfo

func (m *GitCodeDiff) GetA() *GitCodeBlob {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *GitCodeDiff) GetB() *GitCodeBlob {
	if m != nil {
		return m.B
	}
	return nil
}

type NotebookCodeDiff struct {
	A                    *NotebookCodeBlob `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    *NotebookCodeBlob `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NotebookCodeDiff) Reset()         { *m = NotebookCodeDiff{} }
func (m *NotebookCodeDiff) String() string { return proto.CompactTextString(m) }
func (*NotebookCodeDiff) ProtoMessage()    {}
func (*NotebookCodeDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d8129ea3c2050d, []int{5}
}

func (m *NotebookCodeDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookCodeDiff.Unmarshal(m, b)
}
func (m *NotebookCodeDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookCodeDiff.Marshal(b, m, deterministic)
}
func (m *NotebookCodeDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookCodeDiff.Merge(m, src)
}
func (m *NotebookCodeDiff) XXX_Size() int {
	return xxx_messageInfo_NotebookCodeDiff.Size(m)
}
func (m *NotebookCodeDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookCodeDiff.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookCodeDiff proto.InternalMessageInfo

func (m *NotebookCodeDiff) GetA() *NotebookCodeBlob {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *NotebookCodeDiff) GetB() *NotebookCodeBlob {
	if m != nil {
		return m.B
	}
	return nil
}

func init() {
	proto.RegisterType((*CodeBlob)(nil), "ai.verta.modeldb.versioning.CodeBlob")
	proto.RegisterType((*GitCodeBlob)(nil), "ai.verta.modeldb.versioning.GitCodeBlob")
	proto.RegisterType((*NotebookCodeBlob)(nil), "ai.verta.modeldb.versioning.NotebookCodeBlob")
	proto.RegisterType((*CodeDiff)(nil), "ai.verta.modeldb.versioning.CodeDiff")
	proto.RegisterType((*GitCodeDiff)(nil), "ai.verta.modeldb.versioning.GitCodeDiff")
	proto.RegisterType((*NotebookCodeDiff)(nil), "ai.verta.modeldb.versioning.NotebookCodeDiff")
}

func init() {
	proto.RegisterFile("protos/public/modeldb/versioning/Code.proto", fileDescriptor_57d8129ea3c2050d)
}

var fileDescriptor_57d8129ea3c2050d = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0x9e, 0x92, 0x2c, 0x71, 0x94, 0x4b, 0xd0, 0x61, 0x78, 0xdb, 0x25, 0xf8, 0x14, 0xd8, 0x26,
	0xc3, 0x06, 0xbb, 0x64, 0x87, 0xc5, 0x09, 0x6c, 0xa1, 0xb4, 0x04, 0x1d, 0x7a, 0xe8, 0x25, 0xc8,
	0xb6, 0x22, 0x8b, 0x26, 0x96, 0xb1, 0x5f, 0x02, 0x2d, 0xf4, 0x1f, 0xf4, 0x0f, 0xf4, 0xd6, 0x9f,
	0x5a, 0xa4, 0xba, 0x71, 0x08, 0x25, 0xc1, 0xa5, 0x37, 0xe9, 0xf1, 0x7d, 0xdf, 0xfb, 0xf4, 0x3e,
	0x3d, 0xfc, 0x2d, 0xcb, 0x35, 0xe8, 0xc2, 0xcf, 0x36, 0xe1, 0x4a, 0x45, 0xfe, 0x5a, 0xc7, 0x62,
	0x15, 0x87, 0xfe, 0x56, 0xe4, 0x85, 0xd2, 0xa9, 0x4a, 0xa5, 0x3f, 0xd1, 0xb1, 0xa0, 0x16, 0x45,
	0xbe, 0x72, 0x45, 0xb7, 0x22, 0x07, 0x4e, 0x4b, 0x1c, 0xad, 0x70, 0x5f, 0xe8, 0x49, 0xa5, 0x29,
	0x07, 0x5e, 0x08, 0x78, 0x06, 0x7a, 0x8f, 0x08, 0x3b, 0x46, 0x3b, 0x58, 0xe9, 0x90, 0xfc, 0xc1,
	0x4d, 0xa9, 0xc0, 0x45, 0x03, 0x34, 0xec, 0xfd, 0x1c, 0xd2, 0x23, 0x7d, 0xe8, 0x3f, 0x05, 0x2f,
	0xb4, 0xff, 0x1f, 0x98, 0xa1, 0x91, 0x33, 0xec, 0xa4, 0x1a, 0x44, 0xa8, 0xf5, 0xb5, 0xdb, 0xb0,
	0x12, 0x3f, 0x8e, 0x4a, 0x5c, 0x94, 0xe0, 0x3d, 0x9d, 0x9d, 0x40, 0xd0, 0xc5, 0x9d, 0x48, 0xa7,
	0x20, 0x52, 0xf0, 0x6e, 0x71, 0x6f, 0xaf, 0x1b, 0x21, 0xb8, 0x95, 0x8b, 0x4c, 0x5b, 0x97, 0x5d,
	0x66, 0xcf, 0xa6, 0x96, 0xf0, 0x22, 0xb1, 0x6d, 0xbb, 0xcc, 0x9e, 0xc9, 0x27, 0xdc, 0x0e, 0x73,
	0x9e, 0x46, 0x89, 0xdb, 0xb4, 0xd5, 0xf2, 0x46, 0xfa, 0xb8, 0x09, 0x5c, 0xba, 0x2d, 0x5b, 0x34,
	0x47, 0xf2, 0x19, 0x3b, 0xaa, 0x58, 0xc4, 0x2a, 0x87, 0x1b, 0xf7, 0xe3, 0x00, 0x0d, 0x1d, 0xd6,
	0x51, 0xc5, 0xd4, 0x5c, 0xbd, 0x07, 0x84, 0xfb, 0x87, 0x3e, 0xc9, 0x5f, 0xdc, 0xca, 0x38, 0x24,
	0xe5, 0x9c, 0xbe, 0x1f, 0x7d, 0xe4, 0x9c, 0x43, 0x52, 0x4e, 0xdc, 0x70, 0x99, 0x65, 0x92, 0x09,
	0x76, 0xa4, 0x82, 0x85, 0x7d, 0x47, 0xa3, 0xde, 0xb4, 0x59, 0x47, 0x2a, 0x60, 0x22, 0xab, 0xa2,
	0x9b, 0xaa, 0xe5, 0xf2, 0x0d, 0xd1, 0x19, 0xda, 0x7b, 0x44, 0x57, 0xea, 0xbc, 0x1a, 0xdd, 0xdd,
	0x2e, 0x3a, 0x6b, 0xf2, 0x37, 0x46, 0xe3, 0xba, 0xbf, 0x8b, 0xa1, 0xb1, 0xe1, 0x05, 0xb5, 0xe7,
	0x84, 0x02, 0xef, 0xfe, 0x20, 0x3d, 0x6b, 0x62, 0x54, 0x99, 0xa8, 0xf7, 0x3f, 0x8d, 0x93, 0x51,
	0xe5, 0xa4, 0x2e, 0x39, 0x08, 0xce, 0xe7, 0xe8, 0x6a, 0x26, 0x15, 0x24, 0x9b, 0x90, 0x46, 0x7a,
	0xed, 0x5f, 0x1a, 0xf2, 0x78, 0xb6, 0x5b, 0xd2, 0x72, 0x75, 0xa5, 0x48, 0x7d, 0xa9, 0xfd, 0x53,
	0x8b, 0x1c, 0xb6, 0x2d, 0xe2, 0xd7, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x9b, 0x62, 0x5a,
	0x3d, 0x04, 0x00, 0x00,
}
