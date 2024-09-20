// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.0
// source: tigerXdragon.proto

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

// kindID 2004
//场景
type TigerXdragonSceneResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp  int64       `protobuf:"varint,1,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`      //时间戳
	Inning     string      `protobuf:"bytes,2,opt,name=inning,proto3" json:"inning,omitempty"`             //牌局号
	Chips      []int32     `protobuf:"varint,3,rep,packed,name=chips,proto3" json:"chips,omitempty"`       //筹 码
	AwardAreas [][]byte    `protobuf:"bytes,4,rep,name=awardAreas,proto3" json:"awardAreas,omitempty"`     //开奖记录(路单)
	AreaBets   []int64     `protobuf:"varint,5,rep,packed,name=areaBets,proto3" json:"areaBets,omitempty"` //各下注区当前总下注额
	MyBets     []int64     `protobuf:"varint,6,rep,packed,name=myBets,proto3" json:"myBets,omitempty"`     //个人在各下注区的总下注额
	AllPlayers *PlayerList `protobuf:"bytes,7,opt,name=allPlayers,proto3" json:"allPlayers,omitempty"`     //玩家列表
}

func (x *TigerXdragonSceneResp) Reset() {
	*x = TigerXdragonSceneResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonSceneResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonSceneResp) ProtoMessage() {}

func (x *TigerXdragonSceneResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonSceneResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonSceneResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{0}
}

func (x *TigerXdragonSceneResp) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *TigerXdragonSceneResp) GetInning() string {
	if x != nil {
		return x.Inning
	}
	return ""
}

func (x *TigerXdragonSceneResp) GetChips() []int32 {
	if x != nil {
		return x.Chips
	}
	return nil
}

func (x *TigerXdragonSceneResp) GetAwardAreas() [][]byte {
	if x != nil {
		return x.AwardAreas
	}
	return nil
}

func (x *TigerXdragonSceneResp) GetAreaBets() []int64 {
	if x != nil {
		return x.AreaBets
	}
	return nil
}

func (x *TigerXdragonSceneResp) GetMyBets() []int64 {
	if x != nil {
		return x.MyBets
	}
	return nil
}

func (x *TigerXdragonSceneResp) GetAllPlayers() *PlayerList {
	if x != nil {
		return x.AllPlayers
	}
	return nil
}

//状态
// 服务端推送
//(准备)
type TigerXdragonStateStartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times  *TimeInfo `protobuf:"bytes,1,opt,name=times,proto3" json:"times,omitempty"`
	Inning string    `protobuf:"bytes,2,opt,name=inning,proto3" json:"inning,omitempty"` // 牌局号
}

func (x *TigerXdragonStateStartResp) Reset() {
	*x = TigerXdragonStateStartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonStateStartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonStateStartResp) ProtoMessage() {}

func (x *TigerXdragonStateStartResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonStateStartResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonStateStartResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{1}
}

func (x *TigerXdragonStateStartResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

func (x *TigerXdragonStateStartResp) GetInning() string {
	if x != nil {
		return x.Inning
	}
	return ""
}

//(下注)
type TigerXdragonStatePlayingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times *TimeInfo `protobuf:"bytes,1,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *TigerXdragonStatePlayingResp) Reset() {
	*x = TigerXdragonStatePlayingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonStatePlayingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonStatePlayingResp) ProtoMessage() {}

func (x *TigerXdragonStatePlayingResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonStatePlayingResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonStatePlayingResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{2}
}

func (x *TigerXdragonStatePlayingResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

//(开奖)
type TigerXdragonStateOpenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times    *TimeInfo             `protobuf:"bytes,1,opt,name=times,proto3" json:"times,omitempty"`
	OpenInfo *TigerXdragonOpenResp `protobuf:"bytes,2,opt,name=openInfo,proto3" json:"openInfo,omitempty"`
}

func (x *TigerXdragonStateOpenResp) Reset() {
	*x = TigerXdragonStateOpenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonStateOpenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonStateOpenResp) ProtoMessage() {}

func (x *TigerXdragonStateOpenResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonStateOpenResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonStateOpenResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{3}
}

func (x *TigerXdragonStateOpenResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

func (x *TigerXdragonStateOpenResp) GetOpenInfo() *TigerXdragonOpenResp {
	if x != nil {
		return x.OpenInfo
	}
	return nil
}

//(结束)
type TigerXdragonStateOverResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times *TimeInfo `protobuf:"bytes,1,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *TigerXdragonStateOverResp) Reset() {
	*x = TigerXdragonStateOverResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonStateOverResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonStateOverResp) ProtoMessage() {}

func (x *TigerXdragonStateOverResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonStateOverResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonStateOverResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{4}
}

func (x *TigerXdragonStateOverResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

//游戏消息
//下注
type TigerXdragonBetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BetArea  int32 `protobuf:"varint,1,opt,name=betArea,proto3" json:"betArea,omitempty"`   //下注区域
	BetScore int64 `protobuf:"varint,2,opt,name=betScore,proto3" json:"betScore,omitempty"` //下注金额
}

func (x *TigerXdragonBetReq) Reset() {
	*x = TigerXdragonBetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonBetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonBetReq) ProtoMessage() {}

func (x *TigerXdragonBetReq) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonBetReq.ProtoReflect.Descriptor instead.
func (*TigerXdragonBetReq) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{5}
}

func (x *TigerXdragonBetReq) GetBetArea() int32 {
	if x != nil {
		return x.BetArea
	}
	return 0
}

func (x *TigerXdragonBetReq) GetBetScore() int64 {
	if x != nil {
		return x.BetScore
	}
	return 0
}

type TigerXdragonBetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	BetArea  int32 `protobuf:"varint,2,opt,name=betArea,proto3" json:"betArea,omitempty"`   //下注区域
	BetScore int64 `protobuf:"varint,3,opt,name=betScore,proto3" json:"betScore,omitempty"` //下注金额
}

func (x *TigerXdragonBetResp) Reset() {
	*x = TigerXdragonBetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonBetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonBetResp) ProtoMessage() {}

func (x *TigerXdragonBetResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonBetResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonBetResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{6}
}

func (x *TigerXdragonBetResp) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TigerXdragonBetResp) GetBetArea() int32 {
	if x != nil {
		return x.BetArea
	}
	return 0
}

func (x *TigerXdragonBetResp) GetBetScore() int64 {
	if x != nil {
		return x.BetScore
	}
	return 0
}

//结束
type TigerXdragonOpenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards     []byte `protobuf:"bytes,1,opt,name=cards,proto3" json:"cards,omitempty"`         //开奖数据
	AwardArea []byte `protobuf:"bytes,2,opt,name=awardArea,proto3" json:"awardArea,omitempty"` //中奖区域 1~13
}

func (x *TigerXdragonOpenResp) Reset() {
	*x = TigerXdragonOpenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonOpenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonOpenResp) ProtoMessage() {}

func (x *TigerXdragonOpenResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonOpenResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonOpenResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{7}
}

func (x *TigerXdragonOpenResp) GetCards() []byte {
	if x != nil {
		return x.Cards
	}
	return nil
}

func (x *TigerXdragonOpenResp) GetAwardArea() []byte {
	if x != nil {
		return x.AwardArea
	}
	return nil
}

//结算
type TigerXdragonCheckoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyAcquire int64   `protobuf:"varint,1,opt,name=myAcquire,proto3" json:"myAcquire,omitempty"`      //获得金币(结算)
	Acquires  []int64 `protobuf:"varint,2,rep,packed,name=acquires,proto3" json:"acquires,omitempty"` //各个区域输赢情况
}

func (x *TigerXdragonCheckoutResp) Reset() {
	*x = TigerXdragonCheckoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonCheckoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonCheckoutResp) ProtoMessage() {}

func (x *TigerXdragonCheckoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonCheckoutResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonCheckoutResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{8}
}

func (x *TigerXdragonCheckoutResp) GetMyAcquire() int64 {
	if x != nil {
		return x.MyAcquire
	}
	return 0
}

func (x *TigerXdragonCheckoutResp) GetAcquires() []int64 {
	if x != nil {
		return x.Acquires
	}
	return nil
}

/////////////预留协议///////////////////////////
//抢庄
type TigerXdragonHostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsWant bool `protobuf:"varint,1,opt,name=isWant,proto3" json:"isWant,omitempty"` //true上庄 false取消上庄
}

func (x *TigerXdragonHostReq) Reset() {
	*x = TigerXdragonHostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonHostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonHostReq) ProtoMessage() {}

func (x *TigerXdragonHostReq) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonHostReq.ProtoReflect.Descriptor instead.
func (*TigerXdragonHostReq) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{9}
}

func (x *TigerXdragonHostReq) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

type TigerXdragonHostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	IsWant bool  `protobuf:"varint,2,opt,name=isWant,proto3" json:"isWant,omitempty"` //true上庄 false取消上庄
}

func (x *TigerXdragonHostResp) Reset() {
	*x = TigerXdragonHostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonHostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonHostResp) ProtoMessage() {}

func (x *TigerXdragonHostResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonHostResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonHostResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{10}
}

func (x *TigerXdragonHostResp) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TigerXdragonHostResp) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

//超级抢庄
type TigerXdragonSuperHostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsWant bool `protobuf:"varint,1,opt,name=isWant,proto3" json:"isWant,omitempty"` //true上庄 false取消
}

func (x *TigerXdragonSuperHostReq) Reset() {
	*x = TigerXdragonSuperHostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonSuperHostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonSuperHostReq) ProtoMessage() {}

func (x *TigerXdragonSuperHostReq) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonSuperHostReq.ProtoReflect.Descriptor instead.
func (*TigerXdragonSuperHostReq) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{11}
}

func (x *TigerXdragonSuperHostReq) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

type TigerXdragonSuperHostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	IsWant bool  `protobuf:"varint,2,opt,name=isWant,proto3" json:"isWant,omitempty"` //true上庄 false取消
}

func (x *TigerXdragonSuperHostResp) Reset() {
	*x = TigerXdragonSuperHostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tigerXdragon_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TigerXdragonSuperHostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TigerXdragonSuperHostResp) ProtoMessage() {}

func (x *TigerXdragonSuperHostResp) ProtoReflect() protoreflect.Message {
	mi := &file_tigerXdragon_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TigerXdragonSuperHostResp.ProtoReflect.Descriptor instead.
func (*TigerXdragonSuperHostResp) Descriptor() ([]byte, []int) {
	return file_tigerXdragon_proto_rawDescGZIP(), []int{12}
}

func (x *TigerXdragonSuperHostResp) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TigerXdragonSuperHostResp) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

var File_tigerXdragon_proto protoreflect.FileDescriptor

var file_tigerXdragon_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x69, 0x67, 0x65, 0x72, 0x58, 0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x15, 0x54, 0x69, 0x67,
	0x65, 0x72, 0x58, 0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x70,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63, 0x68, 0x69, 0x70, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x77, 0x61, 0x72, 0x64, 0x41, 0x72, 0x65, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x0a, 0x61, 0x77, 0x61, 0x72, 0x64, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x42, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x42, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x79,
	0x42, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x79, 0x42, 0x65,
	0x74, 0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x22, 0x58, 0x0a, 0x1a, 0x54, 0x69, 0x67, 0x65, 0x72, 0x58, 0x64, 0x72, 0x61, 0x67,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x22, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x42, 0x0a, 0x1c,
	0x54, 0x69, 0x67, 0x65, 0x72, 0x58, 0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x22, 0x75, 0x0a, 0x19, 0x54, 0x69, 0x67, 0x65, 0x72, 0x58, 0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a,
	0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x12, 0x34, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x67, 0x65, 0x72, 0x58, 0x64,
	0x72, 0x61, 0x67, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x6f,
	0x70, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3f, 0x0a, 0x19, 0x54, 0x69, 0x67, 0x65, 0x72,
	0x58, 0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x12, 0x54, 0x69, 0x67, 0x65,
	0x72, 0x58, 0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x42, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x62, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x74, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x65, 0x74, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x22, 0x63, 0x0a, 0x13, 0x54, 0x69, 0x67, 0x65, 0x72, 0x58, 0x64, 0x72,
	0x61, 0x67, 0x6f, 0x6e, 0x42, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x62, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x4a, 0x0a, 0x14, 0x54, 0x69, 0x67,
	0x65, 0x72, 0x58, 0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x77, 0x61, 0x72, 0x64,
	0x41, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x77, 0x61, 0x72,
	0x64, 0x41, 0x72, 0x65, 0x61, 0x22, 0x54, 0x0a, 0x18, 0x54, 0x69, 0x67, 0x65, 0x72, 0x58, 0x64,
	0x72, 0x61, 0x67, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x79, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x79, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x08, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x22, 0x2d, 0x0a, 0x13, 0x54,
	0x69, 0x67, 0x65, 0x72, 0x58, 0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x14, 0x54, 0x69,
	0x67, 0x65, 0x72, 0x58, 0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x57, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x57, 0x61,
	0x6e, 0x74, 0x22, 0x32, 0x0a, 0x18, 0x54, 0x69, 0x67, 0x65, 0x72, 0x58, 0x64, 0x72, 0x61, 0x67,
	0x6f, 0x6e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x19, 0x54, 0x69, 0x67, 0x65, 0x72, 0x58,
	0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x73, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x57,
	0x61, 0x6e, 0x74, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tigerXdragon_proto_rawDescOnce sync.Once
	file_tigerXdragon_proto_rawDescData = file_tigerXdragon_proto_rawDesc
)

func file_tigerXdragon_proto_rawDescGZIP() []byte {
	file_tigerXdragon_proto_rawDescOnce.Do(func() {
		file_tigerXdragon_proto_rawDescData = protoimpl.X.CompressGZIP(file_tigerXdragon_proto_rawDescData)
	})
	return file_tigerXdragon_proto_rawDescData
}

var file_tigerXdragon_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_tigerXdragon_proto_goTypes = []interface{}{
	(*TigerXdragonSceneResp)(nil),        // 0: pb.TigerXdragonSceneResp
	(*TigerXdragonStateStartResp)(nil),   // 1: pb.TigerXdragonStateStartResp
	(*TigerXdragonStatePlayingResp)(nil), // 2: pb.TigerXdragonStatePlayingResp
	(*TigerXdragonStateOpenResp)(nil),    // 3: pb.TigerXdragonStateOpenResp
	(*TigerXdragonStateOverResp)(nil),    // 4: pb.TigerXdragonStateOverResp
	(*TigerXdragonBetReq)(nil),           // 5: pb.TigerXdragonBetReq
	(*TigerXdragonBetResp)(nil),          // 6: pb.TigerXdragonBetResp
	(*TigerXdragonOpenResp)(nil),         // 7: pb.TigerXdragonOpenResp
	(*TigerXdragonCheckoutResp)(nil),     // 8: pb.TigerXdragonCheckoutResp
	(*TigerXdragonHostReq)(nil),          // 9: pb.TigerXdragonHostReq
	(*TigerXdragonHostResp)(nil),         // 10: pb.TigerXdragonHostResp
	(*TigerXdragonSuperHostReq)(nil),     // 11: pb.TigerXdragonSuperHostReq
	(*TigerXdragonSuperHostResp)(nil),    // 12: pb.TigerXdragonSuperHostResp
	(*PlayerList)(nil),                   // 13: pb.PlayerList
	(*TimeInfo)(nil),                     // 14: pb.TimeInfo
}
var file_tigerXdragon_proto_depIdxs = []int32{
	13, // 0: pb.TigerXdragonSceneResp.allPlayers:type_name -> pb.PlayerList
	14, // 1: pb.TigerXdragonStateStartResp.times:type_name -> pb.TimeInfo
	14, // 2: pb.TigerXdragonStatePlayingResp.times:type_name -> pb.TimeInfo
	14, // 3: pb.TigerXdragonStateOpenResp.times:type_name -> pb.TimeInfo
	7,  // 4: pb.TigerXdragonStateOpenResp.openInfo:type_name -> pb.TigerXdragonOpenResp
	14, // 5: pb.TigerXdragonStateOverResp.times:type_name -> pb.TimeInfo
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_tigerXdragon_proto_init() }
func file_tigerXdragon_proto_init() {
	if File_tigerXdragon_proto != nil {
		return
	}
	file_baseinfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tigerXdragon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonSceneResp); i {
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
		file_tigerXdragon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonStateStartResp); i {
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
		file_tigerXdragon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonStatePlayingResp); i {
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
		file_tigerXdragon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonStateOpenResp); i {
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
		file_tigerXdragon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonStateOverResp); i {
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
		file_tigerXdragon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonBetReq); i {
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
		file_tigerXdragon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonBetResp); i {
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
		file_tigerXdragon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonOpenResp); i {
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
		file_tigerXdragon_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonCheckoutResp); i {
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
		file_tigerXdragon_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonHostReq); i {
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
		file_tigerXdragon_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonHostResp); i {
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
		file_tigerXdragon_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonSuperHostReq); i {
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
		file_tigerXdragon_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TigerXdragonSuperHostResp); i {
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
			RawDescriptor: file_tigerXdragon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tigerXdragon_proto_goTypes,
		DependencyIndexes: file_tigerXdragon_proto_depIdxs,
		MessageInfos:      file_tigerXdragon_proto_msgTypes,
	}.Build()
	File_tigerXdragon_proto = out.File
	file_tigerXdragon_proto_rawDesc = nil
	file_tigerXdragon_proto_goTypes = nil
	file_tigerXdragon_proto_depIdxs = nil
}
