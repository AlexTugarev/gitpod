// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wssync.proto

package api

import (
	context "context"
	fmt "fmt"
	api "github.com/gitpod-io/gitpod/content-service/api"
	proto "github.com/golang/protobuf/proto"
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

// WorkspaceContentState describes the availability and reliability of the workspace content
type WorkspaceContentState int32

const (
	// NONE means that there currently is no workspace content and no work is underway to change that.
	WorkspaceContentState_NONE WorkspaceContentState = 0
	// SETTING_UP indicates that the workspace content is currently being produced/checked out/unarchived and is
	// very likely to change. In this state one must not modify or rely on the workspace content.
	WorkspaceContentState_SETTING_UP WorkspaceContentState = 1
	// AVAILABLE indicates that the workspace content is fully present and ready for use.
	WorkspaceContentState_AVAILABLE WorkspaceContentState = 2
	// WRAPPING_UP means that the workspace is being torn down, i.e. a final backup is being produced and the content
	// is deleted locally. In this state one must not modify or rely on the workspace content.
	WorkspaceContentState_WRAPPING_UP WorkspaceContentState = 3
)

var WorkspaceContentState_name = map[int32]string{
	0: "NONE",
	1: "SETTING_UP",
	2: "AVAILABLE",
	3: "WRAPPING_UP",
}

var WorkspaceContentState_value = map[string]int32{
	"NONE":        0,
	"SETTING_UP":  1,
	"AVAILABLE":   2,
	"WRAPPING_UP": 3,
}

func (x WorkspaceContentState) String() string {
	return proto.EnumName(WorkspaceContentState_name, int32(x))
}

func (WorkspaceContentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{0}
}

// InitWorkspaceRequest intialises a new workspace folder in the working area
type InitWorkspaceRequest struct {
	// ID is a unique identifier of this workspace. No other workspace with the same name must exist in the realm of this daemon
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Metadata is data associated with this workspace that's required for other parts of Gitpod to function
	Metadata *WorkspaceMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Initializer specifies how the workspace is to be initialized
	Initializer *api.WorkspaceInitializer `protobuf:"bytes,3,opt,name=initializer,proto3" json:"initializer,omitempty"`
	// full_workspace_backup means we ignore the initializer and wait for a workspace pod with the given instance ID to
	// appear at our local containerd.
	FullWorkspaceBackup bool `protobuf:"varint,4,opt,name=full_workspace_backup,json=fullWorkspaceBackup,proto3" json:"fullWorkspaceBackup,omitempty"`
	// content_manifest describes the layers that comprise the workspace image content.
	// This manifest is not used to actually download content, but to produce a new manifest for snapshots and backups.
	// This field is ignored if full_workspace_backup is false.
	ContentManifest      []byte   `protobuf:"bytes,5,opt,name=content_manifest,json=contentManifest,proto3" json:"contentManifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitWorkspaceRequest) Reset()         { *m = InitWorkspaceRequest{} }
func (m *InitWorkspaceRequest) String() string { return proto.CompactTextString(m) }
func (*InitWorkspaceRequest) ProtoMessage()    {}
func (*InitWorkspaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{0}
}

func (m *InitWorkspaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitWorkspaceRequest.Unmarshal(m, b)
}
func (m *InitWorkspaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitWorkspaceRequest.Marshal(b, m, deterministic)
}
func (m *InitWorkspaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitWorkspaceRequest.Merge(m, src)
}
func (m *InitWorkspaceRequest) XXX_Size() int {
	return xxx_messageInfo_InitWorkspaceRequest.Size(m)
}
func (m *InitWorkspaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitWorkspaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitWorkspaceRequest proto.InternalMessageInfo

func (m *InitWorkspaceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InitWorkspaceRequest) GetMetadata() *WorkspaceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *InitWorkspaceRequest) GetInitializer() *api.WorkspaceInitializer {
	if m != nil {
		return m.Initializer
	}
	return nil
}

func (m *InitWorkspaceRequest) GetFullWorkspaceBackup() bool {
	if m != nil {
		return m.FullWorkspaceBackup
	}
	return false
}

func (m *InitWorkspaceRequest) GetContentManifest() []byte {
	if m != nil {
		return m.ContentManifest
	}
	return nil
}

// WorkspaceMetadata is data associated with a workspace that's required for other parts of the system to function
type WorkspaceMetadata struct {
	// owner is the ID of the Gitpod user to whom we'll bill this workspace and who we consider responsible for its content
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// meta_id is the workspace ID of this currently running workspace instance on the "meta pool" side
	MetaId               string   `protobuf:"bytes,2,opt,name=meta_id,json=metaId,proto3" json:"metaId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkspaceMetadata) Reset()         { *m = WorkspaceMetadata{} }
func (m *WorkspaceMetadata) String() string { return proto.CompactTextString(m) }
func (*WorkspaceMetadata) ProtoMessage()    {}
func (*WorkspaceMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{1}
}

func (m *WorkspaceMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkspaceMetadata.Unmarshal(m, b)
}
func (m *WorkspaceMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkspaceMetadata.Marshal(b, m, deterministic)
}
func (m *WorkspaceMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkspaceMetadata.Merge(m, src)
}
func (m *WorkspaceMetadata) XXX_Size() int {
	return xxx_messageInfo_WorkspaceMetadata.Size(m)
}
func (m *WorkspaceMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkspaceMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkspaceMetadata proto.InternalMessageInfo

func (m *WorkspaceMetadata) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *WorkspaceMetadata) GetMetaId() string {
	if m != nil {
		return m.MetaId
	}
	return ""
}

type InitWorkspaceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitWorkspaceResponse) Reset()         { *m = InitWorkspaceResponse{} }
func (m *InitWorkspaceResponse) String() string { return proto.CompactTextString(m) }
func (*InitWorkspaceResponse) ProtoMessage()    {}
func (*InitWorkspaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{2}
}

func (m *InitWorkspaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitWorkspaceResponse.Unmarshal(m, b)
}
func (m *InitWorkspaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitWorkspaceResponse.Marshal(b, m, deterministic)
}
func (m *InitWorkspaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitWorkspaceResponse.Merge(m, src)
}
func (m *InitWorkspaceResponse) XXX_Size() int {
	return xxx_messageInfo_InitWorkspaceResponse.Size(m)
}
func (m *InitWorkspaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitWorkspaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitWorkspaceResponse proto.InternalMessageInfo

// WaitForInitRequest waits for a workspace to be initialized
type WaitForInitRequest struct {
	// ID is a unique identifier of the workspace
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitForInitRequest) Reset()         { *m = WaitForInitRequest{} }
func (m *WaitForInitRequest) String() string { return proto.CompactTextString(m) }
func (*WaitForInitRequest) ProtoMessage()    {}
func (*WaitForInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{3}
}

func (m *WaitForInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitForInitRequest.Unmarshal(m, b)
}
func (m *WaitForInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitForInitRequest.Marshal(b, m, deterministic)
}
func (m *WaitForInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitForInitRequest.Merge(m, src)
}
func (m *WaitForInitRequest) XXX_Size() int {
	return xxx_messageInfo_WaitForInitRequest.Size(m)
}
func (m *WaitForInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitForInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WaitForInitRequest proto.InternalMessageInfo

func (m *WaitForInitRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type WaitForInitResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitForInitResponse) Reset()         { *m = WaitForInitResponse{} }
func (m *WaitForInitResponse) String() string { return proto.CompactTextString(m) }
func (*WaitForInitResponse) ProtoMessage()    {}
func (*WaitForInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{4}
}

func (m *WaitForInitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitForInitResponse.Unmarshal(m, b)
}
func (m *WaitForInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitForInitResponse.Marshal(b, m, deterministic)
}
func (m *WaitForInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitForInitResponse.Merge(m, src)
}
func (m *WaitForInitResponse) XXX_Size() int {
	return xxx_messageInfo_WaitForInitResponse.Size(m)
}
func (m *WaitForInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitForInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WaitForInitResponse proto.InternalMessageInfo

// TakeSnapshotRequest creates a backup/snapshot of a workspace
type TakeSnapshotRequest struct {
	// ID is the identifier of the workspace of which we want to create a snapshot of
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TakeSnapshotRequest) Reset()         { *m = TakeSnapshotRequest{} }
func (m *TakeSnapshotRequest) String() string { return proto.CompactTextString(m) }
func (*TakeSnapshotRequest) ProtoMessage()    {}
func (*TakeSnapshotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{5}
}

func (m *TakeSnapshotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeSnapshotRequest.Unmarshal(m, b)
}
func (m *TakeSnapshotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeSnapshotRequest.Marshal(b, m, deterministic)
}
func (m *TakeSnapshotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeSnapshotRequest.Merge(m, src)
}
func (m *TakeSnapshotRequest) XXX_Size() int {
	return xxx_messageInfo_TakeSnapshotRequest.Size(m)
}
func (m *TakeSnapshotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeSnapshotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TakeSnapshotRequest proto.InternalMessageInfo

func (m *TakeSnapshotRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TakeSnapshotResponse struct {
	// url is the name of the resulting snapshot
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TakeSnapshotResponse) Reset()         { *m = TakeSnapshotResponse{} }
func (m *TakeSnapshotResponse) String() string { return proto.CompactTextString(m) }
func (*TakeSnapshotResponse) ProtoMessage()    {}
func (*TakeSnapshotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{6}
}

func (m *TakeSnapshotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeSnapshotResponse.Unmarshal(m, b)
}
func (m *TakeSnapshotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeSnapshotResponse.Marshal(b, m, deterministic)
}
func (m *TakeSnapshotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeSnapshotResponse.Merge(m, src)
}
func (m *TakeSnapshotResponse) XXX_Size() int {
	return xxx_messageInfo_TakeSnapshotResponse.Size(m)
}
func (m *TakeSnapshotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeSnapshotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TakeSnapshotResponse proto.InternalMessageInfo

func (m *TakeSnapshotResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type DisposeWorkspaceRequest struct {
	// ID is a unique identifier of the workspace to dispose of
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Backup triggers a final backup prior to disposal
	Backup               bool     `protobuf:"varint,2,opt,name=backup,proto3" json:"backup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisposeWorkspaceRequest) Reset()         { *m = DisposeWorkspaceRequest{} }
func (m *DisposeWorkspaceRequest) String() string { return proto.CompactTextString(m) }
func (*DisposeWorkspaceRequest) ProtoMessage()    {}
func (*DisposeWorkspaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{7}
}

func (m *DisposeWorkspaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisposeWorkspaceRequest.Unmarshal(m, b)
}
func (m *DisposeWorkspaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisposeWorkspaceRequest.Marshal(b, m, deterministic)
}
func (m *DisposeWorkspaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisposeWorkspaceRequest.Merge(m, src)
}
func (m *DisposeWorkspaceRequest) XXX_Size() int {
	return xxx_messageInfo_DisposeWorkspaceRequest.Size(m)
}
func (m *DisposeWorkspaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisposeWorkspaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisposeWorkspaceRequest proto.InternalMessageInfo

func (m *DisposeWorkspaceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DisposeWorkspaceRequest) GetBackup() bool {
	if m != nil {
		return m.Backup
	}
	return false
}

type DisposeWorkspaceResponse struct {
	// git_status is the current state of the Git repo in this workspace prior to disposal.
	// If the workspace has no Git repo at its checkout location, this is nil.
	GitStatus            *api.GitStatus `protobuf:"bytes,1,opt,name=git_status,json=gitStatus,proto3" json:"gitStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DisposeWorkspaceResponse) Reset()         { *m = DisposeWorkspaceResponse{} }
func (m *DisposeWorkspaceResponse) String() string { return proto.CompactTextString(m) }
func (*DisposeWorkspaceResponse) ProtoMessage()    {}
func (*DisposeWorkspaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c7b25cc593b9e14, []int{8}
}

func (m *DisposeWorkspaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisposeWorkspaceResponse.Unmarshal(m, b)
}
func (m *DisposeWorkspaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisposeWorkspaceResponse.Marshal(b, m, deterministic)
}
func (m *DisposeWorkspaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisposeWorkspaceResponse.Merge(m, src)
}
func (m *DisposeWorkspaceResponse) XXX_Size() int {
	return xxx_messageInfo_DisposeWorkspaceResponse.Size(m)
}
func (m *DisposeWorkspaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisposeWorkspaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisposeWorkspaceResponse proto.InternalMessageInfo

func (m *DisposeWorkspaceResponse) GetGitStatus() *api.GitStatus {
	if m != nil {
		return m.GitStatus
	}
	return nil
}

func init() {
	proto.RegisterEnum("wssync.WorkspaceContentState", WorkspaceContentState_name, WorkspaceContentState_value)
	proto.RegisterType((*InitWorkspaceRequest)(nil), "wssync.InitWorkspaceRequest")
	proto.RegisterType((*WorkspaceMetadata)(nil), "wssync.WorkspaceMetadata")
	proto.RegisterType((*InitWorkspaceResponse)(nil), "wssync.InitWorkspaceResponse")
	proto.RegisterType((*WaitForInitRequest)(nil), "wssync.WaitForInitRequest")
	proto.RegisterType((*WaitForInitResponse)(nil), "wssync.WaitForInitResponse")
	proto.RegisterType((*TakeSnapshotRequest)(nil), "wssync.TakeSnapshotRequest")
	proto.RegisterType((*TakeSnapshotResponse)(nil), "wssync.TakeSnapshotResponse")
	proto.RegisterType((*DisposeWorkspaceRequest)(nil), "wssync.DisposeWorkspaceRequest")
	proto.RegisterType((*DisposeWorkspaceResponse)(nil), "wssync.DisposeWorkspaceResponse")
}

func init() {
	proto.RegisterFile("wssync.proto", fileDescriptor_9c7b25cc593b9e14)
}

var fileDescriptor_9c7b25cc593b9e14 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x8f, 0xd2, 0x4c,
	0x14, 0xde, 0x76, 0x77, 0x79, 0x97, 0xc3, 0x7e, 0xf4, 0x1d, 0x40, 0x6a, 0x5d, 0x23, 0x69, 0x76,
	0x23, 0x6b, 0x02, 0x24, 0x18, 0x13, 0x6f, 0x41, 0xd9, 0x95, 0xb8, 0x8b, 0x58, 0x50, 0x12, 0x6f,
	0x9a, 0xa1, 0x9d, 0x65, 0x27, 0x40, 0xa7, 0x76, 0xa6, 0x12, 0xbd, 0xf6, 0x4f, 0xfa, 0x6f, 0x4c,
	0xdb, 0xa1, 0xcb, 0x67, 0xbc, 0x9b, 0xd3, 0xe7, 0xe3, 0x4c, 0xcf, 0x73, 0x5a, 0x38, 0x9e, 0x73,
	0xfe, 0xd3, 0x73, 0x6a, 0x7e, 0xc0, 0x04, 0x43, 0x99, 0xa4, 0x32, 0x2e, 0x1d, 0xe6, 0x09, 0xe2,
	0x89, 0x2a, 0x27, 0xc1, 0x0f, 0xea, 0x90, 0x2a, 0xf6, 0x69, 0x9d, 0x7a, 0x54, 0x50, 0x3c, 0xa5,
	0xbf, 0x48, 0x90, 0xd0, 0xcd, 0xdf, 0x2a, 0x14, 0x3a, 0x1e, 0x15, 0x43, 0x16, 0x4c, 0xb8, 0x8f,
	0x1d, 0x62, 0x91, 0xef, 0x21, 0xe1, 0x02, 0x9d, 0x82, 0x4a, 0x5d, 0x5d, 0x29, 0x2b, 0x95, 0xac,
	0xa5, 0x52, 0x17, 0xbd, 0x81, 0xa3, 0x19, 0x11, 0xd8, 0xc5, 0x02, 0xeb, 0x6a, 0x59, 0xa9, 0xe4,
	0x1a, 0x4f, 0x6b, 0xb2, 0x71, 0xaa, 0xbd, 0x93, 0x04, 0x2b, 0xa5, 0xa2, 0x6b, 0xc8, 0x2d, 0x35,
	0xd5, 0xf7, 0x63, 0xe5, 0x45, 0x4d, 0x5e, 0x4e, 0xde, 0xed, 0xd1, 0xa1, 0xf3, 0xc8, 0xb5, 0x96,
	0x85, 0xa8, 0x01, 0xc5, 0xfb, 0x70, 0x3a, 0xb5, 0xe7, 0x0b, 0xa6, 0x3d, 0xc2, 0xce, 0x24, 0xf4,
	0xf5, 0x83, 0xb2, 0x52, 0x39, 0xb2, 0xf2, 0x11, 0x98, 0xba, 0xb4, 0x62, 0x08, 0x5d, 0x81, 0x26,
	0xfb, 0xd8, 0x33, 0xec, 0xd1, 0x7b, 0xc2, 0x85, 0x7e, 0x58, 0x56, 0x2a, 0xc7, 0xd6, 0x99, 0x7c,
	0x7e, 0x27, 0x1f, 0x9b, 0x2d, 0xf8, 0x7f, 0xe3, 0x2d, 0x50, 0x01, 0x0e, 0xd9, 0xdc, 0x23, 0x81,
	0x9c, 0x42, 0x52, 0xa0, 0x12, 0xfc, 0x17, 0xbd, 0x9d, 0x4d, 0xdd, 0x78, 0x0e, 0x59, 0x2b, 0x13,
	0x95, 0x1d, 0xd7, 0x2c, 0x41, 0x71, 0x6d, 0x92, 0xdc, 0x67, 0x1e, 0x27, 0xe6, 0x05, 0xa0, 0x21,
	0xa6, 0xe2, 0x9a, 0x05, 0x11, 0xbe, 0x63, 0xc0, 0x66, 0x11, 0xf2, 0x2b, 0x2c, 0x29, 0xbe, 0x84,
	0xfc, 0x00, 0x4f, 0x48, 0xdf, 0xc3, 0x3e, 0x7f, 0x60, 0x3b, 0xd5, 0x15, 0x28, 0xac, 0xd2, 0x12,
	0x39, 0xd2, 0x60, 0x3f, 0x0c, 0xa6, 0x92, 0x18, 0x1d, 0xcd, 0x26, 0x94, 0xde, 0x53, 0xee, 0x33,
	0x4e, 0xfe, 0x99, 0xf9, 0x13, 0xc8, 0xc8, 0x29, 0xab, 0xf1, 0x94, 0x65, 0x65, 0x0e, 0x40, 0xdf,
	0xb4, 0x90, 0x0d, 0xdf, 0x02, 0x8c, 0xa9, 0xb0, 0xb9, 0xc0, 0x22, 0xe4, 0xb1, 0x57, 0xb4, 0x29,
	0x6b, 0x79, 0xdf, 0x50, 0xd1, 0x8f, 0x09, 0x56, 0x76, 0xbc, 0x38, 0xbe, 0xfa, 0x0c, 0xc5, 0xd4,
	0xee, 0x5d, 0xc2, 0x8f, 0x10, 0x82, 0x8e, 0xe0, 0xa0, 0xfb, 0xa9, 0xdb, 0xd6, 0xf6, 0xd0, 0x29,
	0x40, 0xbf, 0x3d, 0x18, 0x74, 0xba, 0x37, 0xf6, 0x97, 0x9e, 0xa6, 0xa0, 0x13, 0xc8, 0x36, 0xbf,
	0x36, 0x3b, 0xb7, 0xcd, 0xd6, 0x6d, 0x5b, 0x53, 0xd1, 0x19, 0xe4, 0x86, 0x56, 0xb3, 0xd7, 0x93,
	0xf8, 0x7e, 0xe3, 0x8f, 0x0a, 0xa5, 0x0d, 0xcf, 0xe4, 0x0e, 0xa8, 0x0b, 0x27, 0x2b, 0x71, 0xa1,
	0xf3, 0xc5, 0x3e, 0x6f, 0xfb, 0x1e, 0x8c, 0xe7, 0x3b, 0x50, 0x19, 0xd3, 0x1e, 0xfa, 0x00, 0xb9,
	0xa5, 0xfc, 0x90, 0x91, 0x7e, 0x1d, 0x1b, 0xd1, 0x1b, 0xcf, 0xb6, 0x62, 0xa9, 0xd3, 0x47, 0x38,
	0x5e, 0xce, 0x12, 0xa5, 0xf4, 0x2d, 0x8b, 0x60, 0x9c, 0x6f, 0x07, 0x53, 0xb3, 0x21, 0x68, 0xeb,
	0x59, 0xa1, 0x17, 0x0b, 0xcd, 0x8e, 0x45, 0x30, 0xca, 0xbb, 0x09, 0x0b, 0xe3, 0xd6, 0xd5, 0xb7,
	0x97, 0x63, 0x2a, 0x1e, 0xc2, 0x51, 0xcd, 0x61, 0xb3, 0xfa, 0x98, 0x0a, 0x9f, 0xb9, 0x55, 0xca,
	0xe4, 0xa9, 0x3e, 0xe7, 0xd5, 0xc8, 0xa1, 0x8e, 0x7d, 0x3a, 0xca, 0xc4, 0xff, 0x9a, 0xd7, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x45, 0x30, 0xdf, 0xaa, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkspaceContentServiceClient is the client API for WorkspaceContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkspaceContentServiceClient interface {
	// initWorkspace intialises a new workspace folder in the working area
	InitWorkspace(ctx context.Context, in *InitWorkspaceRequest, opts ...grpc.CallOption) (*InitWorkspaceResponse, error)
	// WaitForInit waits until a workspace is fully initialized.
	// If the workspace is already initialized, this function returns immediately.
	// If there is no initialization is going on, an error is returned.
	WaitForInit(ctx context.Context, in *WaitForInitRequest, opts ...grpc.CallOption) (*WaitForInitResponse, error)
	// TakeSnapshot creates a backup/snapshot of a workspace
	TakeSnapshot(ctx context.Context, in *TakeSnapshotRequest, opts ...grpc.CallOption) (*TakeSnapshotResponse, error)
	// disposeWorkspace cleans up a workspace, possibly after taking a final backup
	DisposeWorkspace(ctx context.Context, in *DisposeWorkspaceRequest, opts ...grpc.CallOption) (*DisposeWorkspaceResponse, error)
}

type workspaceContentServiceClient struct {
	cc grpc.ClientConnInterface `json:"cc,omitempty"`
}

func NewWorkspaceContentServiceClient(cc grpc.ClientConnInterface) WorkspaceContentServiceClient {
	return &workspaceContentServiceClient{cc}
}

func (c *workspaceContentServiceClient) InitWorkspace(ctx context.Context, in *InitWorkspaceRequest, opts ...grpc.CallOption) (*InitWorkspaceResponse, error) {
	out := new(InitWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/wssync.WorkspaceContentService/InitWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceContentServiceClient) WaitForInit(ctx context.Context, in *WaitForInitRequest, opts ...grpc.CallOption) (*WaitForInitResponse, error) {
	out := new(WaitForInitResponse)
	err := c.cc.Invoke(ctx, "/wssync.WorkspaceContentService/WaitForInit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceContentServiceClient) TakeSnapshot(ctx context.Context, in *TakeSnapshotRequest, opts ...grpc.CallOption) (*TakeSnapshotResponse, error) {
	out := new(TakeSnapshotResponse)
	err := c.cc.Invoke(ctx, "/wssync.WorkspaceContentService/TakeSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceContentServiceClient) DisposeWorkspace(ctx context.Context, in *DisposeWorkspaceRequest, opts ...grpc.CallOption) (*DisposeWorkspaceResponse, error) {
	out := new(DisposeWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/wssync.WorkspaceContentService/DisposeWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceContentServiceServer is the server API for WorkspaceContentService service.
type WorkspaceContentServiceServer interface {
	// initWorkspace intialises a new workspace folder in the working area
	InitWorkspace(context.Context, *InitWorkspaceRequest) (*InitWorkspaceResponse, error)
	// WaitForInit waits until a workspace is fully initialized.
	// If the workspace is already initialized, this function returns immediately.
	// If there is no initialization is going on, an error is returned.
	WaitForInit(context.Context, *WaitForInitRequest) (*WaitForInitResponse, error)
	// TakeSnapshot creates a backup/snapshot of a workspace
	TakeSnapshot(context.Context, *TakeSnapshotRequest) (*TakeSnapshotResponse, error)
	// disposeWorkspace cleans up a workspace, possibly after taking a final backup
	DisposeWorkspace(context.Context, *DisposeWorkspaceRequest) (*DisposeWorkspaceResponse, error)
}

// UnimplementedWorkspaceContentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkspaceContentServiceServer struct {
}

func (*UnimplementedWorkspaceContentServiceServer) InitWorkspace(ctx context.Context, req *InitWorkspaceRequest) (*InitWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitWorkspace not implemented")
}
func (*UnimplementedWorkspaceContentServiceServer) WaitForInit(ctx context.Context, req *WaitForInitRequest) (*WaitForInitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitForInit not implemented")
}
func (*UnimplementedWorkspaceContentServiceServer) TakeSnapshot(ctx context.Context, req *TakeSnapshotRequest) (*TakeSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeSnapshot not implemented")
}
func (*UnimplementedWorkspaceContentServiceServer) DisposeWorkspace(ctx context.Context, req *DisposeWorkspaceRequest) (*DisposeWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisposeWorkspace not implemented")
}

func RegisterWorkspaceContentServiceServer(s *grpc.Server, srv WorkspaceContentServiceServer) {
	s.RegisterService(&_WorkspaceContentService_serviceDesc, srv)
}

func _WorkspaceContentService_InitWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceContentServiceServer).InitWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wssync.WorkspaceContentService/InitWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceContentServiceServer).InitWorkspace(ctx, req.(*InitWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceContentService_WaitForInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitForInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceContentServiceServer).WaitForInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wssync.WorkspaceContentService/WaitForInit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceContentServiceServer).WaitForInit(ctx, req.(*WaitForInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceContentService_TakeSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceContentServiceServer).TakeSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wssync.WorkspaceContentService/TakeSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceContentServiceServer).TakeSnapshot(ctx, req.(*TakeSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceContentService_DisposeWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisposeWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceContentServiceServer).DisposeWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wssync.WorkspaceContentService/DisposeWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceContentServiceServer).DisposeWorkspace(ctx, req.(*DisposeWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkspaceContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wssync.WorkspaceContentService",
	HandlerType: (*WorkspaceContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitWorkspace",
			Handler:    _WorkspaceContentService_InitWorkspace_Handler,
		},
		{
			MethodName: "WaitForInit",
			Handler:    _WorkspaceContentService_WaitForInit_Handler,
		},
		{
			MethodName: "TakeSnapshot",
			Handler:    _WorkspaceContentService_TakeSnapshot_Handler,
		},
		{
			MethodName: "DisposeWorkspace",
			Handler:    _WorkspaceContentService_DisposeWorkspace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wssync.proto",
}
