package account

import (
	"github.com/po2656233/superplace/const/code"
	cactor "github.com/po2656233/superplace/net/actor"
	"sanguoxiao/internal/hints"
	"sanguoxiao/internal/protocol/gofile"
	db2 "sanguoxiao/nodes/center/db"
	"strings"
)

type (
	ActorAccount struct {
		cactor.Base
	}
)

func (p *ActorAccount) AliasID() string {
	return "account"
}

// OnInit center为后端节点，不直接与客户端通信，所以了一些remote函数，供RPC调用
func (p *ActorAccount) OnInit() {
	p.Remote().Register(p.RegisterReq)
	p.Remote().Register(p.LoginReq)
	p.Remote().Register(p.GetUID)
}

// RegisterReq 注册开发者帐号
func (p *ActorAccount) RegisterReq(req *pb.RegisterReq) (*pb.RegisterResp, int32) {
	accountName := req.Name
	password := req.Password

	if strings.TrimSpace(accountName) == "" || strings.TrimSpace(password) == "" {
		return nil, hints.Login02
	}

	if len(accountName) < 3 || len(accountName) > 18 {
		return nil, hints.Login02
	}

	if len(password) < 3 || len(password) > 18 {
		return nil, hints.Login02
	}

	db2.DevAccountRegister(accountName, password, req.SecurityCode)

	return &pb.RegisterResp{
		Info: &pb.UserInfo{
			UserID: 10,
			Name:   accountName,
		},
	}, code.OK
}

// LoginReq 根据帐号名获取开发者帐号表
func (p *ActorAccount) LoginReq(req *pb.LoginReq) (*pb.LoginResp, int32) {
	accountName := req.Account
	password := req.Password

	devAccount, _ := db2.DevAccountWithName(accountName)
	if devAccount == nil || devAccount.Password != password {
		return nil, hints.Login07
	}

	return &pb.LoginResp{
		MainInfo: &pb.MasterInfo{
			UserInfo: &pb.UserInfo{UserID: devAccount.AccountId},
		},
	}, code.OK
}

// GetUID 获取uid
func (p *ActorAccount) GetUID(req *pb.User) (*pb.Int64, int32) {
	uid, ok := db2.BindUID(req.SdkId, req.Pid, req.OpenId)
	if uid == 0 || ok == false {
		return nil, hints.Login07
	}

	return &pb.Int64{Value: uid}, code.OK
}
