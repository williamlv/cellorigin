// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// client -> login
type LoginREQ struct {
	PlatformToken string `protobuf:"bytes,1,opt,name=PlatformToken,json=platformToken" json:"PlatformToken,omitempty"`
	PlatformUID   string `protobuf:"bytes,2,opt,name=PlatformUID,json=platformUID" json:"PlatformUID,omitempty"`
	PlatformName  string `protobuf:"bytes,3,opt,name=PlatformName,json=platformName" json:"PlatformName,omitempty"`
	ClientVersion string `protobuf:"bytes,4,opt,name=ClientVersion,json=clientVersion" json:"ClientVersion,omitempty"`
}

func (m *LoginREQ) Reset()                    { *m = LoginREQ{} }
func (m *LoginREQ) String() string            { return proto.CompactTextString(m) }
func (*LoginREQ) ProtoMessage()               {}
func (*LoginREQ) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// 服务器信息
type ServerInfo struct {
	Name        string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=DisplayName,json=displayName" json:"DisplayName,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=Address,json=address" json:"Address,omitempty"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// login -> client
type LoginACK struct {
	Result     string        `protobuf:"bytes,1,opt,name=Result,json=result" json:"Result,omitempty"`
	Token      string        `protobuf:"bytes,2,opt,name=Token,json=token" json:"Token,omitempty"`
	ServerList []*ServerInfo `protobuf:"bytes,3,rep,name=ServerList,json=serverList" json:"ServerList,omitempty"`
}

func (m *LoginACK) Reset()                    { *m = LoginACK{} }
func (m *LoginACK) String() string            { return proto.CompactTextString(m) }
func (*LoginACK) ProtoMessage()               {}
func (*LoginACK) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *LoginACK) GetServerList() []*ServerInfo {
	if m != nil {
		return m.ServerList
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginREQ)(nil), "gamedef.LoginREQ")
	proto.RegisterType((*ServerInfo)(nil), "gamedef.ServerInfo")
	proto.RegisterType((*LoginACK)(nil), "gamedef.LoginACK")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x95, 0xaf, 0x69, 0xf2, 0x71, 0xd3, 0x2e, 0x17, 0x84, 0x3c, 0x56, 0x11, 0x43, 0xa7,
	0x0c, 0xf4, 0x09, 0xaa, 0x96, 0xa1, 0xa2, 0x42, 0x10, 0xfe, 0x4c, 0x2c, 0x81, 0xdc, 0x54, 0x16,
	0x89, 0x1d, 0xd9, 0x06, 0x89, 0x77, 0xe1, 0x61, 0x49, 0x6c, 0x87, 0x64, 0xf4, 0xb9, 0x47, 0xe7,
	0xfc, 0x7c, 0x20, 0xa9, 0xe5, 0x89, 0x8b, 0xac, 0x55, 0xd2, 0x48, 0x8c, 0x4f, 0x45, 0x43, 0x25,
	0x55, 0xe9, 0x4f, 0x00, 0xff, 0x8f, 0xfd, 0x21, 0xbf, 0x79, 0xc0, 0x2b, 0x58, 0xde, 0xd7, 0x85,
	0xa9, 0xa4, 0x6a, 0x9e, 0xe4, 0x07, 0x09, 0x16, 0xac, 0x82, 0xf5, 0x59, 0xbe, 0x6c, 0xa7, 0x22,
	0xae, 0x20, 0x19, 0x5c, 0xcf, 0x87, 0x3d, 0xfb, 0x67, 0x3d, 0x49, 0x3b, 0x4a, 0x98, 0xc2, 0x62,
	0x70, 0xdc, 0x75, 0x3d, 0x6c, 0x66, 0x2d, 0x8b, 0x76, 0xa2, 0xf5, 0x5d, 0xbb, 0x9a, 0x93, 0x30,
	0x2f, 0xa4, 0x34, 0x97, 0x82, 0x85, 0xae, 0xeb, 0x7d, 0x2a, 0xa6, 0xaf, 0x00, 0x8f, 0xa4, 0xbe,
	0x48, 0x1d, 0x44, 0x25, 0x11, 0x21, 0xb4, 0x79, 0x0e, 0x2b, 0x14, 0x7d, 0x4e, 0x47, 0xb3, 0xe7,
	0xba, 0x8b, 0xfe, 0xb6, 0x27, 0x4f, 0x53, 0x8e, 0x12, 0x32, 0x88, 0xb7, 0x65, 0xa9, 0x48, 0x6b,
	0x0f, 0x12, 0x17, 0xee, 0x99, 0x36, 0xfe, 0xef, 0xdb, 0xdd, 0x2d, 0x5e, 0x42, 0x94, 0x93, 0xfe,
	0xac, 0x8d, 0x4f, 0x8f, 0x94, 0x7d, 0xe1, 0x05, 0xcc, 0xdd, 0x16, 0x2e, 0x79, 0x6e, 0xec, 0x06,
	0x9b, 0x81, 0xeb, 0xc8, 0xb5, 0xe9, 0x62, 0x67, 0xeb, 0xe4, 0xfa, 0x3c, 0xf3, 0xa3, 0x66, 0x23,
	0x72, 0x0e, 0xfa, 0xcf, 0xf6, 0x16, 0xd9, 0xed, 0x37, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b,
	0x15, 0xdc, 0xc4, 0x8a, 0x01, 0x00, 0x00,
}
