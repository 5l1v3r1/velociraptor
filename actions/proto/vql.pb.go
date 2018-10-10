// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vql.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "www.velocidex.com/golang/velociraptor/artifacts/proto"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VQLRequest struct {
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	VQL                  string   `protobuf:"bytes,1,opt,name=VQL,proto3" json:"VQL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VQLRequest) Reset()         { *m = VQLRequest{} }
func (m *VQLRequest) String() string { return proto.CompactTextString(m) }
func (*VQLRequest) ProtoMessage()    {}
func (*VQLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_588db8edf5d83b48, []int{0}
}
func (m *VQLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLRequest.Unmarshal(m, b)
}
func (m *VQLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLRequest.Marshal(b, m, deterministic)
}
func (dst *VQLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLRequest.Merge(dst, src)
}
func (m *VQLRequest) XXX_Size() int {
	return xxx_messageInfo_VQLRequest.Size(m)
}
func (m *VQLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VQLRequest proto.InternalMessageInfo

func (m *VQLRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VQLRequest) GetVQL() string {
	if m != nil {
		return m.VQL
	}
	return ""
}

type VQLEnv struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VQLEnv) Reset()         { *m = VQLEnv{} }
func (m *VQLEnv) String() string { return proto.CompactTextString(m) }
func (*VQLEnv) ProtoMessage()    {}
func (*VQLEnv) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_588db8edf5d83b48, []int{1}
}
func (m *VQLEnv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLEnv.Unmarshal(m, b)
}
func (m *VQLEnv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLEnv.Marshal(b, m, deterministic)
}
func (dst *VQLEnv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLEnv.Merge(dst, src)
}
func (m *VQLEnv) XXX_Size() int {
	return xxx_messageInfo_VQLEnv.Size(m)
}
func (m *VQLEnv) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLEnv.DiscardUnknown(m)
}

var xxx_messageInfo_VQLEnv proto.InternalMessageInfo

func (m *VQLEnv) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *VQLEnv) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type VQLCollectorArgs struct {
	Env                  []*VQLEnv          `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	Query                []*VQLRequest      `protobuf:"bytes,2,rep,name=Query,proto3" json:"Query,omitempty"`
	MaxRow               uint64             `protobuf:"varint,4,opt,name=max_row,json=maxRow,proto3" json:"max_row,omitempty"`
	MaxWait              uint64             `protobuf:"varint,6,opt,name=max_wait,json=maxWait,proto3" json:"max_wait,omitempty"`
	Artifacts            []*proto1.Artifact `protobuf:"bytes,5,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VQLCollectorArgs) Reset()         { *m = VQLCollectorArgs{} }
func (m *VQLCollectorArgs) String() string { return proto.CompactTextString(m) }
func (*VQLCollectorArgs) ProtoMessage()    {}
func (*VQLCollectorArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_588db8edf5d83b48, []int{2}
}
func (m *VQLCollectorArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLCollectorArgs.Unmarshal(m, b)
}
func (m *VQLCollectorArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLCollectorArgs.Marshal(b, m, deterministic)
}
func (dst *VQLCollectorArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLCollectorArgs.Merge(dst, src)
}
func (m *VQLCollectorArgs) XXX_Size() int {
	return xxx_messageInfo_VQLCollectorArgs.Size(m)
}
func (m *VQLCollectorArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLCollectorArgs.DiscardUnknown(m)
}

var xxx_messageInfo_VQLCollectorArgs proto.InternalMessageInfo

func (m *VQLCollectorArgs) GetEnv() []*VQLEnv {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *VQLCollectorArgs) GetQuery() []*VQLRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *VQLCollectorArgs) GetMaxRow() uint64 {
	if m != nil {
		return m.MaxRow
	}
	return 0
}

func (m *VQLCollectorArgs) GetMaxWait() uint64 {
	if m != nil {
		return m.MaxWait
	}
	return 0
}

func (m *VQLCollectorArgs) GetArtifacts() []*proto1.Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type VQLTypeMap struct {
	Column               string   `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VQLTypeMap) Reset()         { *m = VQLTypeMap{} }
func (m *VQLTypeMap) String() string { return proto.CompactTextString(m) }
func (*VQLTypeMap) ProtoMessage()    {}
func (*VQLTypeMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_588db8edf5d83b48, []int{3}
}
func (m *VQLTypeMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLTypeMap.Unmarshal(m, b)
}
func (m *VQLTypeMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLTypeMap.Marshal(b, m, deterministic)
}
func (dst *VQLTypeMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLTypeMap.Merge(dst, src)
}
func (m *VQLTypeMap) XXX_Size() int {
	return xxx_messageInfo_VQLTypeMap.Size(m)
}
func (m *VQLTypeMap) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLTypeMap.DiscardUnknown(m)
}

var xxx_messageInfo_VQLTypeMap proto.InternalMessageInfo

func (m *VQLTypeMap) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *VQLTypeMap) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type VQLResponse struct {
	Response             string        `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	Columns              []string      `protobuf:"bytes,2,rep,name=Columns,proto3" json:"Columns,omitempty"`
	Types                []*VQLTypeMap `protobuf:"bytes,8,rep,name=types,proto3" json:"types,omitempty"`
	QueryId              uint64        `protobuf:"varint,5,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	Part                 uint64        `protobuf:"varint,6,opt,name=part,proto3" json:"part,omitempty"`
	Query                *VQLRequest   `protobuf:"bytes,3,opt,name=Query,proto3" json:"Query,omitempty"`
	Timestamp            uint64        `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TotalRows            uint64        `protobuf:"varint,7,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VQLResponse) Reset()         { *m = VQLResponse{} }
func (m *VQLResponse) String() string { return proto.CompactTextString(m) }
func (*VQLResponse) ProtoMessage()    {}
func (*VQLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_588db8edf5d83b48, []int{4}
}
func (m *VQLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLResponse.Unmarshal(m, b)
}
func (m *VQLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLResponse.Marshal(b, m, deterministic)
}
func (dst *VQLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLResponse.Merge(dst, src)
}
func (m *VQLResponse) XXX_Size() int {
	return xxx_messageInfo_VQLResponse.Size(m)
}
func (m *VQLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VQLResponse proto.InternalMessageInfo

func (m *VQLResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *VQLResponse) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *VQLResponse) GetTypes() []*VQLTypeMap {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *VQLResponse) GetQueryId() uint64 {
	if m != nil {
		return m.QueryId
	}
	return 0
}

func (m *VQLResponse) GetPart() uint64 {
	if m != nil {
		return m.Part
	}
	return 0
}

func (m *VQLResponse) GetQuery() *VQLRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *VQLResponse) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *VQLResponse) GetTotalRows() uint64 {
	if m != nil {
		return m.TotalRows
	}
	return 0
}

// FIXME: We replicate a small subset of GRR's elaborate knowledgebase
// protos here because the GUI API plugins use this to construct the
// GRR APIs. When we re-implement the API plugins, refactor this into
// a more sane structure.
type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_588db8edf5d83b48, []int{5}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Knowledgebase struct {
	Users                []*User  `protobuf:"bytes,32,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Knowledgebase) Reset()         { *m = Knowledgebase{} }
func (m *Knowledgebase) String() string { return proto.CompactTextString(m) }
func (*Knowledgebase) ProtoMessage()    {}
func (*Knowledgebase) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_588db8edf5d83b48, []int{6}
}
func (m *Knowledgebase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Knowledgebase.Unmarshal(m, b)
}
func (m *Knowledgebase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Knowledgebase.Marshal(b, m, deterministic)
}
func (dst *Knowledgebase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Knowledgebase.Merge(dst, src)
}
func (m *Knowledgebase) XXX_Size() int {
	return xxx_messageInfo_Knowledgebase.Size(m)
}
func (m *Knowledgebase) XXX_DiscardUnknown() {
	xxx_messageInfo_Knowledgebase.DiscardUnknown(m)
}

var xxx_messageInfo_Knowledgebase proto.InternalMessageInfo

func (m *Knowledgebase) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ClientInfo struct {
	Info []*VQLResponse `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	// GRR Keeps these separated so they can be searched on.
	LastTimestamp        uint64         `protobuf:"varint,2,opt,name=last_timestamp,json=lastTimestamp,proto3" json:"last_timestamp,omitempty"`
	Hostname             string         `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Fqdn                 string         `protobuf:"bytes,4,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	System               string         `protobuf:"bytes,5,opt,name=system,proto3" json:"system,omitempty"`
	Release              string         `protobuf:"bytes,6,opt,name=release,proto3" json:"release,omitempty"`
	Architecture         string         `protobuf:"bytes,7,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Usernames            string         `protobuf:"bytes,8,opt,name=usernames,proto3" json:"usernames,omitempty"`
	MacAddress           string         `protobuf:"bytes,9,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	IpAddress            string         `protobuf:"bytes,10,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Ping                 uint64         `protobuf:"varint,11,opt,name=ping,proto3" json:"ping,omitempty"`
	ClientVersion        string         `protobuf:"bytes,12,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	ClientName           string         `protobuf:"bytes,13,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	Knowledgebase        *Knowledgebase `protobuf:"bytes,14,opt,name=knowledgebase,proto3" json:"knowledgebase,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_588db8edf5d83b48, []int{7}
}
func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (dst *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(dst, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetInfo() []*VQLResponse {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ClientInfo) GetLastTimestamp() uint64 {
	if m != nil {
		return m.LastTimestamp
	}
	return 0
}

func (m *ClientInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ClientInfo) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *ClientInfo) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *ClientInfo) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *ClientInfo) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *ClientInfo) GetUsernames() string {
	if m != nil {
		return m.Usernames
	}
	return ""
}

func (m *ClientInfo) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *ClientInfo) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ClientInfo) GetPing() uint64 {
	if m != nil {
		return m.Ping
	}
	return 0
}

func (m *ClientInfo) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *ClientInfo) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *ClientInfo) GetKnowledgebase() *Knowledgebase {
	if m != nil {
		return m.Knowledgebase
	}
	return nil
}

func init() {
	proto.RegisterType((*VQLRequest)(nil), "proto.VQLRequest")
	proto.RegisterType((*VQLEnv)(nil), "proto.VQLEnv")
	proto.RegisterType((*VQLCollectorArgs)(nil), "proto.VQLCollectorArgs")
	proto.RegisterType((*VQLTypeMap)(nil), "proto.VQLTypeMap")
	proto.RegisterType((*VQLResponse)(nil), "proto.VQLResponse")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*Knowledgebase)(nil), "proto.Knowledgebase")
	proto.RegisterType((*ClientInfo)(nil), "proto.ClientInfo")
}

func init() { proto.RegisterFile("vql.proto", fileDescriptor_vql_588db8edf5d83b48) }

var fileDescriptor_vql_588db8edf5d83b48 = []byte{
	// 1359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x4d, 0x6f, 0x1c, 0x45,
	0x13, 0xd6, 0x66, 0x77, 0x6d, 0x6f, 0x3b, 0xce, 0x9b, 0xb4, 0xa2, 0x68, 0xde, 0x10, 0x42, 0x67,
	0xf9, 0x4a, 0x14, 0x65, 0xfd, 0x11, 0x45, 0x09, 0x46, 0x10, 0x79, 0x9d, 0x04, 0x05, 0xec, 0x20,
	0x0f, 0xc6, 0xa0, 0x28, 0xc2, 0xf4, 0xce, 0xd4, 0xee, 0xb4, 0xd2, 0xd3, 0x3d, 0xee, 0xee, 0xd9,
	0xf1, 0x1e, 0x72, 0xe2, 0x00, 0x7f, 0x82, 0x3b, 0x37, 0xce, 0xdc, 0x39, 0x20, 0xf1, 0x03, 0x90,
	0x38, 0xc2, 0x8d, 0xdf, 0xc0, 0x01, 0x75, 0xcd, 0xc7, 0xda, 0x11, 0x48, 0x70, 0xf2, 0x74, 0x75,
	0xd7, 0x53, 0x4f, 0x7d, 0x3c, 0xb5, 0x26, 0xbd, 0xe9, 0x91, 0x1c, 0x64, 0x46, 0x3b, 0x4d, 0xbb,
	0xf8, 0xe7, 0xf2, 0x66, 0x51, 0x14, 0x83, 0x29, 0x48, 0x1d, 0x89, 0x18, 0x8e, 0x07, 0x91, 0x4e,
	0x57, 0x27, 0x5a, 0x72, 0x35, 0x59, 0x2d, 0x8d, 0x86, 0x67, 0x4e, 0x9b, 0x55, 0x7c, 0xbc, 0x6a,
	0x21, 0xe5, 0xca, 0x89, 0xa8, 0x84, 0xb8, 0x7c, 0xff, 0xbf, 0xf8, 0x8e, 0xa5, 0x2e, 0x0e, 0x53,
	0x70, 0x3c, 0xe6, 0x8e, 0x57, 0x00, 0x0f, 0xfe, 0x1d, 0x00, 0x37, 0x4e, 0x8c, 0x79, 0xe4, 0x6c,
	0x05, 0x55, 0x9f, 0x4b, 0x94, 0xfe, 0xaf, 0x2d, 0x42, 0x0e, 0xf6, 0x76, 0x42, 0x38, 0xca, 0xc1,
	0x3a, 0xfa, 0x4d, 0x8b, 0x74, 0x9e, 0xf0, 0x14, 0x82, 0x33, 0xac, 0x75, 0xbd, 0x37, 0x74, 0xbf,
	0xfd, 0xf9, 0xfb, 0x8f, 0x2d, 0x45, 0xe5, 0x7e, 0x02, 0x4c, 0xf1, 0x14, 0x98, 0x1e, 0x33, 0x97,
	0x08, 0xcb, 0x8e, 0x72, 0x30, 0xb3, 0x01, 0xdb, 0xf7, 0xdf, 0x36, 0xd1, 0xb9, 0x8c, 0xd9, 0x08,
	0x58, 0x0c, 0x36, 0x32, 0x22, 0x73, 0x62, 0x0a, 0xcc, 0x69, 0x26, 0x54, 0x2c, 0x22, 0xee, 0x80,
	0x15, 0x09, 0x77, 0xcc, 0xcd, 0x32, 0xf4, 0x17, 0x6a, 0xac, 0x4d, 0xca, 0x85, 0xd3, 0x8a, 0xb9,
	0x04, 0x4a, 0x28, 0x66, 0xc0, 0x19, 0x01, 0x53, 0xb0, 0x83, 0x10, 0x19, 0xd0, 0x2d, 0xd2, 0x3e,
	0xd8, 0xdb, 0x09, 0x5a, 0x48, 0x64, 0x15, 0x89, 0xdc, 0xa0, 0x6f, 0x7b, 0x22, 0x07, 0x7b, 0x3b,
	0x95, 0x93, 0xd3, 0x0c, 0x8e, 0x21, 0xca, 0x1d, 0xb0, 0x0a, 0x2c, 0x92, 0x02, 0x94, 0x1b, 0x84,
	0xde, 0xb7, 0xbf, 0x46, 0x16, 0x0e, 0xf6, 0x76, 0x1e, 0xaa, 0x29, 0x3d, 0x4f, 0xda, 0xcf, 0x61,
	0x56, 0x82, 0x85, 0xfe, 0x93, 0x5e, 0x24, 0xdd, 0x29, 0x97, 0x79, 0x95, 0x69, 0x58, 0x1e, 0xfa,
	0x3f, 0x74, 0xc9, 0xf9, 0x83, 0xbd, 0x9d, 0x6d, 0x2d, 0x25, 0x44, 0x4e, 0x9b, 0x2d, 0x33, 0xb1,
	0xf4, 0x19, 0x69, 0x83, 0x9a, 0x06, 0x6d, 0xd6, 0xbe, 0xbe, 0xbc, 0xb1, 0x52, 0x16, 0x6e, 0x50,
	0x02, 0x0f, 0xdf, 0x45, 0x62, 0x77, 0xe8, 0xed, 0x87, 0x6a, 0x2a, 0x8c, 0x56, 0x29, 0x28, 0xc7,
	0xa6, 0xdc, 0x08, 0x3e, 0x92, 0x60, 0x3d, 0xc1, 0x11, 0xb0, 0xcc, 0xe8, 0xa9, 0x88, 0x21, 0x66,
	0x63, 0x6d, 0xe6, 0x29, 0x0f, 0x42, 0x0f, 0x4b, 0x9f, 0x92, 0xee, 0x9e, 0x3f, 0x06, 0x67, 0x10,
	0xff, 0xc2, 0x1c, 0xbf, 0x6a, 0xca, 0x70, 0x1d, 0x63, 0xdc, 0xa4, 0x37, 0x4e, 0x26, 0x2f, 0x4a,
	0xf4, 0x7f, 0x48, 0xbf, 0x84, 0xa4, 0xdf, 0xb6, 0xc8, 0x62, 0xca, 0x8f, 0x0f, 0x8d, 0x2e, 0x82,
	0x0e, 0x6b, 0x5d, 0xef, 0x0c, 0xbf, 0x6a, 0x21, 0xd8, 0x0b, 0x3a, 0xf2, 0x60, 0x29, 0x3f, 0x16,
	0x69, 0x9e, 0x32, 0xa3, 0x0b, 0xcb, 0x32, 0x30, 0xcc, 0x80, 0xcd, 0xb4, 0xb2, 0x30, 0x60, 0x61,
	0xf5, 0x65, 0x99, 0xe4, 0x66, 0x02, 0x9e, 0x37, 0x57, 0x65, 0xef, 0x0b, 0x21, 0xa5, 0x4f, 0xcb,
	0x66, 0x52, 0x38, 0xc6, 0x23, 0xa3, 0xad, 0x65, 0x69, 0x2e, 0x9d, 0xc8, 0x24, 0x34, 0x10, 0x76,
	0xd0, 0xbf, 0xb0, 0xcb, 0x8f, 0xe7, 0xd8, 0x19, 0x37, 0x6e, 0xa3, 0xb3, 0xbe, 0xb6, 0xb6, 0x16,
	0x2e, 0xa4, 0xfc, 0x38, 0xd4, 0x05, 0xfd, 0xb9, 0x45, 0x96, 0x3c, 0xbf, 0x82, 0x0b, 0x17, 0x2c,
	0x20, 0xc1, 0xef, 0x4b, 0x82, 0xdf, 0xb5, 0xe8, 0x8b, 0x47, 0xda, 0x30, 0xa9, 0xd5, 0xa4, 0xc9,
	0xb7, 0xf0, 0xe8, 0x2e, 0x37, 0x0a, 0xb1, 0x04, 0x97, 0x3e, 0x58, 0x2e, 0x9d, 0x65, 0x7c, 0xec,
	0x90, 0xa2, 0xb0, 0xe8, 0x51, 0x0d, 0xa6, 0xb0, 0xcc, 0xc0, 0x51, 0x2e, 0x4c, 0x55, 0x7a, 0x98,
	0xfa, 0xee, 0x48, 0x61, 0x1d, 0x28, 0x30, 0x96, 0x15, 0x89, 0x88, 0x12, 0xa6, 0x60, 0x8a, 0x99,
	0x73, 0x29, 0x67, 0x2c, 0xd2, 0x69, 0x26, 0xc1, 0xc1, 0xa0, 0xff, 0xfa, 0x90, 0xbb, 0x28, 0x41,
	0x00, 0xb0, 0xce, 0xfa, 0xfc, 0x1d, 0x73, 0xfc, 0x39, 0x9c, 0x08, 0xb3, 0x71, 0x66, 0x7d, 0x2d,
	0xf4, 0x15, 0xfe, 0x8c, 0x0b, 0x47, 0x15, 0xe9, 0x35, 0x6a, 0x0b, 0xba, 0xd8, 0xcd, 0xff, 0x55,
	0xdd, 0xdc, 0xaa, 0xec, 0xc3, 0xfb, 0x98, 0xdd, 0x3b, 0xf4, 0x6e, 0x6d, 0xb1, 0xcc, 0x7a, 0x52,
	0x63, 0xa3, 0x53, 0xec, 0xa0, 0x05, 0xe3, 0xe9, 0x38, 0xcd, 0x12, 0x90, 0x19, 0x2b, 0x84, 0x4b,
	0x4e, 0x2a, 0x2e, 0x9c, 0x87, 0xd8, 0x5c, 0xf9, 0xe3, 0x6b, 0xd6, 0x23, 0x8b, 0x1f, 0xf8, 0x64,
	0x44, 0xd4, 0xbf, 0x87, 0x4a, 0xde, 0x9f, 0x65, 0xb0, 0xcb, 0x33, 0x7a, 0x89, 0x2c, 0x44, 0x5a,
	0xe6, 0xa9, 0xaa, 0x86, 0xbe, 0x3a, 0x51, 0x4a, 0x3a, 0x5e, 0x88, 0xd5, 0xd8, 0xe3, 0x77, 0xff,
	0xa7, 0x2e, 0x59, 0xc6, 0x79, 0x2b, 0x3b, 0x47, 0x37, 0xc9, 0x52, 0xfd, 0x5d, 0xe9, 0xef, 0x2a,
	0xd2, 0x0e, 0xe8, 0xa5, 0x0f, 0x3f, 0xf9, 0xf8, 0x09, 0x03, 0x15, 0x69, 0x3f, 0xce, 0xcd, 0xb4,
	0x84, 0xcd, 0x7b, 0x1a, 0x92, 0xc5, 0x6d, 0x8c, 0x64, 0x71, 0xa0, 0x7b, 0xc3, 0x7b, 0xe8, 0xba,
	0x41, 0xd7, 0xb6, 0xb0, 0xf2, 0x7e, 0x03, 0x94, 0x54, 0x58, 0x02, 0x3c, 0x16, 0x6a, 0x62, 0xbd,
	0x3e, 0xe2, 0x3c, 0x82, 0x98, 0x8d, 0x66, 0x27, 0xe5, 0x51, 0x03, 0xd1, 0x2f, 0x48, 0xd7, 0xf3,
	0xb4, 0xc1, 0xd2, 0xcb, 0x12, 0xa9, 0xb2, 0x1d, 0xde, 0xc1, 0x20, 0xab, 0xf4, 0xd6, 0x2e, 0xcf,
	0x32, 0xa1, 0x26, 0x6c, 0x04, 0xae, 0x00, 0x50, 0x75, 0x28, 0xbf, 0xbb, 0x2c, 0xe3, 0x2a, 0xf6,
	0xf8, 0xc2, 0xe0, 0x2e, 0xb2, 0x83, 0xb0, 0x84, 0xa5, 0x21, 0x59, 0xc2, 0x90, 0x87, 0x22, 0x0e,
	0xba, 0x38, 0x85, 0x77, 0x11, 0x6f, 0x9d, 0xae, 0x6e, 0x27, 0x46, 0x2b, 0x2d, 0xf5, 0x44, 0x44,
	0x5c, 0x32, 0x6d, 0x62, 0x30, 0xe5, 0x0e, 0xac, 0xf7, 0x56, 0x51, 0x4f, 0x7c, 0xcc, 0x9c, 0x1e,
	0x84, 0x8b, 0x68, 0x7d, 0x1c, 0x53, 0x47, 0x3a, 0x7e, 0x48, 0xab, 0xa9, 0xfe, 0x12, 0xf1, 0x9e,
	0xd2, 0xcf, 0x77, 0xbc, 0x98, 0x50, 0xc4, 0x8d, 0x4e, 0x18, 0x37, 0x2f, 0x6b, 0x89, 0xab, 0x19,
	0x8e, 0xb8, 0xad, 0x06, 0x39, 0xe2, 0xa6, 0x14, 0x7c, 0x02, 0x68, 0x6f, 0xd6, 0xf0, 0xbc, 0x03,
	0x18, 0x8d, 0xee, 0xd6, 0xcb, 0xa4, 0xcd, 0x5a, 0x7f, 0xbf, 0x4c, 0xde, 0x40, 0x26, 0x57, 0xe9,
	0x95, 0xfd, 0x26, 0x05, 0x9c, 0xec, 0x82, 0xdb, 0x7a, 0x99, 0xc4, 0xcd, 0xfe, 0x78, 0x46, 0x7a,
	0x4e, 0xa4, 0x60, 0x1d, 0x4f, 0xb3, 0x6a, 0x81, 0xbc, 0x8f, 0xfe, 0xf7, 0xc8, 0x72, 0xf8, 0xe0,
	0xd1, 0x03, 0xee, 0xc0, 0xdf, 0x97, 0x9b, 0xa9, 0x79, 0x89, 0x34, 0x6b, 0x66, 0x08, 0x3c, 0xf1,
	0x23, 0xca, 0x11, 0x79, 0x0e, 0x48, 0x43, 0x42, 0x9c, 0x76, 0x5c, 0xfa, 0xf5, 0x64, 0x83, 0x45,
	0x84, 0xbf, 0x8d, 0xf0, 0xb7, 0xe8, 0xcd, 0x7d, 0x7f, 0xc3, 0x54, 0x9e, 0x8e, 0xca, 0x8a, 0xe3,
	0x1a, 0x11, 0xea, 0x74, 0xda, 0x58, 0x0c, 0x8f, 0xe9, 0x1f, 0x87, 0xba, 0xb0, 0xfd, 0x87, 0xa4,
	0xf3, 0xa9, 0x05, 0x43, 0xdf, 0x23, 0x4b, 0xb9, 0x05, 0xe3, 0xfb, 0x5e, 0x8d, 0xf0, 0x35, 0x44,
	0x7e, 0x85, 0xfe, 0xdf, 0x73, 0xad, 0xef, 0xea, 0x5e, 0xfa, 0xf3, 0x20, 0x6c, 0x5c, 0xfa, 0x1b,
	0x64, 0xe5, 0x23, 0xa5, 0x0b, 0x09, 0xf1, 0x04, 0x46, 0xdc, 0x02, 0xbd, 0x46, 0xba, 0xfe, 0xd2,
	0x06, 0x0c, 0x47, 0x70, 0xb9, 0x2a, 0xac, 0x8f, 0x15, 0x96, 0x37, 0xfd, 0x5f, 0xda, 0x84, 0x6c,
	0xe3, 0xfe, 0x7d, 0xac, 0xc6, 0x9a, 0xbe, 0x45, 0x3a, 0xfe, 0x97, 0x2e, 0x68, 0xa1, 0x03, 0x3d,
	0xd9, 0x89, 0x92, 0x7c, 0x88, 0xf7, 0xf4, 0x4d, 0x72, 0x4e, 0x72, 0xeb, 0x0e, 0xe7, 0x85, 0xf6,
	0xd2, 0xec, 0x84, 0x2b, 0xde, 0xba, 0xdf, 0x14, 0xeb, 0x32, 0x59, 0x4a, 0xb4, 0x75, 0x98, 0x50,
	0x1b, 0xb5, 0xdb, 0x9c, 0xbd, 0xa6, 0xc7, 0x47, 0xb1, 0xc2, 0x0e, 0xf5, 0x42, 0xfc, 0xf6, 0xfa,
	0xb7, 0x33, 0xeb, 0x20, 0xc5, 0x89, 0xee, 0x85, 0xd5, 0x89, 0x06, 0x64, 0xd1, 0x80, 0x04, 0x6e,
	0x01, 0x47, 0xb3, 0x17, 0xd6, 0x47, 0xda, 0x27, 0x67, 0xb9, 0x89, 0x12, 0xe1, 0x20, 0x72, 0xb9,
	0x01, 0x6c, 0x48, 0x2f, 0x3c, 0x65, 0xa3, 0x57, 0x48, 0xaf, 0xae, 0x91, 0x57, 0xa3, 0x7f, 0x30,
	0x37, 0xd0, 0xd7, 0xc8, 0x72, 0xca, 0xa3, 0x43, 0x1e, 0xc7, 0x06, 0xac, 0x0d, 0x7a, 0x78, 0x4f,
	0x52, 0x1e, 0x6d, 0x95, 0x16, 0xfa, 0x2a, 0x21, 0x22, 0x6b, 0xee, 0x49, 0xe9, 0x2f, 0xb2, 0xfa,
	0x9a, 0x92, 0x8e, 0x17, 0x6f, 0xb0, 0x8c, 0x05, 0xc0, 0x6f, 0x5f, 0x9e, 0xf2, 0x47, 0xed, 0x70,
	0x0a, 0xc6, 0x0a, 0xad, 0x82, 0xb3, 0xe8, 0xb6, 0x52, 0x5a, 0x0f, 0x4a, 0xa3, 0x0f, 0x5d, 0x3d,
	0xc3, 0x0a, 0xad, 0x94, 0xa1, 0x4b, 0x13, 0xfe, 0x3b, 0xb1, 0x49, 0x56, 0x9e, 0x9f, 0xec, 0x68,
	0x70, 0x0e, 0x15, 0x72, 0xb1, 0xea, 0xcb, 0xa9, 0x6e, 0x87, 0xa7, 0x9f, 0x8e, 0x16, 0xf0, 0xcd,
	0xed, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x72, 0xb0, 0x45, 0x02, 0x0a, 0x00, 0x00,
}
