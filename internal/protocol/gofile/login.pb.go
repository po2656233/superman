// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.0
// source: login.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

/////////////[优秀如你]-->Req:请求 Resp:反馈<--[交互专用]///////////////////////////////////
// 走web服时,该项忽略
//注册
type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                     //用户
	Password       string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`             //密码
	SecurityCode   string `protobuf:"bytes,3,opt,name=securityCode,proto3" json:"securityCode,omitempty"`     //验证码
	MachineCode    string `protobuf:"bytes,4,opt,name=machineCode,proto3" json:"machineCode,omitempty"`       //机器码
	InvitationCode string `protobuf:"bytes,5,opt,name=invitationCode,proto3" json:"invitationCode,omitempty"` //邀请码
	//选填
	Gender     int32  `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`        //性别
	Age        int32  `protobuf:"varint,7,opt,name=age,proto3" json:"age,omitempty"`              //年龄
	FaceID     int32  `protobuf:"varint,8,opt,name=faceID,proto3" json:"faceID,omitempty"`        //头像
	PassPortID string `protobuf:"bytes,9,opt,name=passPortID,proto3" json:"passPortID,omitempty"` //证件号
	RealName   string `protobuf:"bytes,10,opt,name=realName,proto3" json:"realName,omitempty"`    //真实名字
	PhoneNum   string `protobuf:"bytes,11,opt,name=phoneNum,proto3" json:"phoneNum,omitempty"`    //手机
	Email      string `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`          //邮箱
	Address    string `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"`      //住址
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetSecurityCode() string {
	if x != nil {
		return x.SecurityCode
	}
	return ""
}

func (x *RegisterReq) GetMachineCode() string {
	if x != nil {
		return x.MachineCode
	}
	return ""
}

func (x *RegisterReq) GetInvitationCode() string {
	if x != nil {
		return x.InvitationCode
	}
	return ""
}

func (x *RegisterReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *RegisterReq) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *RegisterReq) GetFaceID() int32 {
	if x != nil {
		return x.FaceID
	}
	return 0
}

func (x *RegisterReq) GetPassPortID() string {
	if x != nil {
		return x.PassPortID
	}
	return ""
}

func (x *RegisterReq) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *RegisterReq) GetPhoneNum() string {
	if x != nil {
		return x.PhoneNum
	}
	return ""
}

func (x *RegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  UserInfo info = 1;
	SdkId    int32  `protobuf:"varint,2,opt,name=sdkId,proto3" json:"sdkId,omitempty"`       // sdk id
	Pid      int32  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`           // 包id
	OpenId   string `protobuf:"bytes,4,opt,name=openId,proto3" json:"openId,omitempty"`      // sdk的openid 即uid
	ServerId int32  `protobuf:"varint,5,opt,name=serverId,proto3" json:"serverId,omitempty"` // 所在游戏服id
	Ip       string `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`              // 请求ip
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResp) GetSdkId() int32 {
	if x != nil {
		return x.SdkId
	}
	return 0
}

func (x *RegisterResp) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *RegisterResp) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *RegisterResp) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *RegisterResp) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// 登陆请求(建立连接后的第一条消息，验证通过后则进行后续流程)
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int32            `protobuf:"varint,1,opt,name=serverId,proto3" json:"serverId,omitempty"`                                                                                     // 当前登陆的服务器id
	Token    string           `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                                                                                            // 登陆token(web login api生成的base64字符串)
	Params   map[int32]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 登陆时上传的参数 key: LoginParams
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *LoginRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginRequest) GetParams() map[int32]string {
	if x != nil {
		return x.Params
	}
	return nil
}

// 登陆响应
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64            `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                                                                                               // 游戏内的用户唯一id
	Pid    int32            `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`                                                                                               // 平台id
	OpenId string           `protobuf:"bytes,3,opt,name=openId,proto3" json:"openId,omitempty"`                                                                                          // 平台openId(平台的帐号唯一id)
	Params map[int32]string `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 登陆后的扩展参数，按需增加
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LoginResponse) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *LoginResponse) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *LoginResponse) GetParams() map[int32]string {
	if x != nil {
		return x.Params
	}
	return nil
}

//登录
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account      string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`           //账号
	Password     string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`         //密码
	SecurityCode string `protobuf:"bytes,3,opt,name=securityCode,proto3" json:"securityCode,omitempty"` //验证码
	MachineCode  string `protobuf:"bytes,4,opt,name=machineCode,proto3" json:"machineCode,omitempty"`   //机器码
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginReq) GetSecurityCode() string {
	if x != nil {
		return x.SecurityCode
	}
	return ""
}

func (x *LoginReq) GetMachineCode() string {
	if x != nil {
		return x.MachineCode
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainInfo   *MasterInfo `protobuf:"bytes,1,opt,name=mainInfo,proto3" json:"mainInfo,omitempty"`
	InRoomID   int64       `protobuf:"varint,2,opt,name=inRoomID,proto3" json:"inRoomID,omitempty"`     //所在房间ID(=0时,为系统房间)
	InTableNum int64       `protobuf:"varint,3,opt,name=inTableNum,proto3" json:"inTableNum,omitempty"` //所在游戏ID(=0时,不在任何游戏中)
	Token      string      `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResp) GetMainInfo() *MasterInfo {
	if x != nil {
		return x.MainInfo
	}
	return nil
}

func (x *LoginResp) GetInRoomID() int64 {
	if x != nil {
		return x.InRoomID
	}
	return 0
}

func (x *LoginResp) GetInTableNum() int64 {
	if x != nil {
		return x.InTableNum
	}
	return 0
}

func (x *LoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

//异地
type AllopatricResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *AllopatricResp) Reset() {
	*x = AllopatricResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllopatricResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllopatricResp) ProtoMessage() {}

func (x *AllopatricResp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllopatricResp.ProtoReflect.Descriptor instead.
func (*AllopatricResp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{6}
}

func (x *AllopatricResp) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

//重连
type ReconnectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account     string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`         //账号
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`       //密码
	MachineCode string `protobuf:"bytes,3,opt,name=machineCode,proto3" json:"machineCode,omitempty"` //机器码
}

func (x *ReconnectReq) Reset() {
	*x = ReconnectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconnectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconnectReq) ProtoMessage() {}

func (x *ReconnectReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconnectReq.ProtoReflect.Descriptor instead.
func (*ReconnectReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{7}
}

func (x *ReconnectReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *ReconnectReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ReconnectReq) GetMachineCode() string {
	if x != nil {
		return x.MachineCode
	}
	return ""
}

type ReconnectResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainInfo   *MasterInfo `protobuf:"bytes,1,opt,name=mainInfo,proto3" json:"mainInfo,omitempty"`
	InRoomID   int64       `protobuf:"varint,2,opt,name=inRoomID,proto3" json:"inRoomID,omitempty"`     //所在房间ID(=0时,系统房)
	InTableNum int64       `protobuf:"varint,3,opt,name=inTableNum,proto3" json:"inTableNum,omitempty"` //所在桌子编号(=0时,没有进桌)
	Token      string      `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ReconnectResp) Reset() {
	*x = ReconnectResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconnectResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconnectResp) ProtoMessage() {}

func (x *ReconnectResp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconnectResp.ProtoReflect.Descriptor instead.
func (*ReconnectResp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{8}
}

func (x *ReconnectResp) GetMainInfo() *MasterInfo {
	if x != nil {
		return x.MainInfo
	}
	return nil
}

func (x *ReconnectResp) GetInRoomID() int64 {
	if x != nil {
		return x.InRoomID
	}
	return 0
}

func (x *ReconnectResp) GetInTableNum() int64 {
	if x != nil {
		return x.InTableNum
	}
	return 0
}

func (x *ReconnectResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

//登出
type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{9}
}

func (x *LogoutReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type LogoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LogoutResp) Reset() {
	*x = LogoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResp) ProtoMessage() {}

func (x *LogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResp.ProtoReflect.Descriptor instead.
func (*LogoutResp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{10}
}

func (x *LogoutResp) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LogoutResp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf5, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x61, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x7a, 0x0a, 0x0c, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x64, 0x6b,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x64, 0x6b, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0xb1, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39,
	0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbd, 0x01, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39,
	0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x2a, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x28,
	0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x6f, 0x70, 0x61, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x66, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2a, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x69, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x1d, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x3c, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x5a,
	0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_login_proto_goTypes = []interface{}{
	(*RegisterReq)(nil),    // 0: pb.RegisterReq
	(*RegisterResp)(nil),   // 1: pb.RegisterResp
	(*LoginRequest)(nil),   // 2: pb.LoginRequest
	(*LoginResponse)(nil),  // 3: pb.LoginResponse
	(*LoginReq)(nil),       // 4: pb.LoginReq
	(*LoginResp)(nil),      // 5: pb.LoginResp
	(*AllopatricResp)(nil), // 6: pb.AllopatricResp
	(*ReconnectReq)(nil),   // 7: pb.ReconnectReq
	(*ReconnectResp)(nil),  // 8: pb.ReconnectResp
	(*LogoutReq)(nil),      // 9: pb.LogoutReq
	(*LogoutResp)(nil),     // 10: pb.LogoutResp
	nil,                    // 11: pb.LoginRequest.ParamsEntry
	nil,                    // 12: pb.LoginResponse.ParamsEntry
	(*MasterInfo)(nil),     // 13: pb.MasterInfo
}
var file_login_proto_depIdxs = []int32{
	11, // 0: pb.LoginRequest.params:type_name -> pb.LoginRequest.ParamsEntry
	12, // 1: pb.LoginResponse.params:type_name -> pb.LoginResponse.ParamsEntry
	13, // 2: pb.LoginResp.mainInfo:type_name -> pb.MasterInfo
	13, // 3: pb.ReconnectResp.mainInfo:type_name -> pb.MasterInfo
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	file_baseinfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_login_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllopatricResp); i {
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
		file_login_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconnectReq); i {
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
		file_login_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconnectResp); i {
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
		file_login_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReq); i {
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
		file_login_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResp); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
