package manger

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	commMsg "superman/internal/protocol/go_file/common"
	gameMsg "superman/internal/protocol/go_file/game"
	gateMsg "superman/internal/protocol/go_file/gate"
	sqlmodel "superman/internal/sql_model/minigame"
	"superman/nodes/game/db"

	"superman/internal/redis_cluster"
	"superman/internal/utils"
	"sync"
	"sync/atomic"
	"time"
)

// 废弃
var dbCmpt *db.Component

func GetDBCmpt() *db.Component {
	if dbCmpt == nil {
		var ok bool
		if dbCmpt, ok = GetClientMgr().GetApp().Find(GameDb).(*db.Component); !ok {
			panic("DB NOT init!!!")
		}
	}
	return dbCmpt
}

// Table 桌牌
type Table struct {
	*commMsg.TableInfo
	GameHandle IGameOperate
	sitters    sync.Map // 座位上的玩家(根据玩家状态区分，旁观者和正在完的玩家)
	maxChairID int32    // 便于新增椅子号
	sitCount   int32
	isStart    bool
}

// ///////////////////桌牌功能//////////////////////////////////////////////
//var settingFixTime int64

// Init 加载配置 [废弃:不加载公用配置,拆分至子游戏内部实现]
func (tb *Table) Init() {
	//解析游戏
	////获取文件的
	//fInfo, _ := os.Stat(settingFile)
	//lastFixTime := fInfo.ModTime().Unix()
	//// 在文件被修改或配置内容为空时，则读取配置信息
	//if lastFixTime != settingFixTime  {
	//	//读取配置信息
	//	yamlFile, err := os.ReadFile(settingFile)
	//	//log.Debug("yamlFile:%v", yamlFile)
	//	if err != nil {
	//		log.Fatal("yamlFile.Get err #%v ", err)
	//	}
	//
	//	err = yaml.Unmarshal(yamlFile, self.Config)
	//	// err = yaml.Unmarshal(yamlFile, &resultMap)
	//	if err != nil {
	//		log.Error("Unmarshal: %v", err)
	//	}
	//}
	//settingFixTime = lastFixTime

	//log.Debug("conf:%v", self.Config)
}

// Reset 重置座位上的玩家
func (tb *Table) Reset() int64 {
	// 遍历过程中存在slice的数据删除

	tb.Init()
	return tb.Id
}

// AddChair 添加椅子
func (tb *Table) AddChair(player *Player) (errCode int) {
	// 是否满员了
	if tb.MaxSitter != Unlimited && tb.MaxSitter < tb.SitCount()+1 {
		return Game39
	}
	if player.PlayerInfo.Coin < tb.PlayScore {
		GetClientMgr().SendPopResultX(player.UserID, FAILED, StatusText[Title009], StatusText[TableInfo12])
		return TableInfo12
	}

	playerList := make([]int64, 0)
	player.PlayerInfo.State = commMsg.PlayerState_PlayerSitDown
	resp := &gameMsg.JoinGameReadyQueueResp{
		RoomID:     tb.Rid,
		TableID:    tb.Id,
		PlayerList: make([]*commMsg.PlayerSimpleInfo, 0),
	}
	// 检测是否存在了
	sit, ok := tb.sitters.Load(player.UserID)
	isNew := false
	if !ok {
		isNew = true
		sit = NewChair(atomic.AddInt32(&tb.maxChairID, 1), tb.Gid, player)
		atomic.AddInt32(&tb.sitCount, 1)
		tb.sitters.Store(player.UserID, sit)
	}
	if sit == nil {
		return User03
	}
	sitter, ok1 := sit.(*Chair)
	if !ok1 {
		return TableInfo11
	}

	// 获取玩家简要信息
	tb.sitters.Range(func(key, value any) bool {
		alreadyChair := value.(*Chair)
		info := alreadyChair.PlayerInfo
		resp.PlayerList = append(resp.PlayerList, alreadyChair.ToSimpleInfo())
		playerList = append(playerList, info.UserID)
		return true
	})

	// 发送给刚坐下的玩家，关于其他玩家的信息
	player.InRooId = tb.Rid
	player.InTableId = tb.Id
	player.InChairId = sitter.ID
	player.GameHandle = tb.GameHandle
	GetClientMgr().SendTo(player.UserID, resp)

	// 发给其他玩家，当前刚坐下的玩家
	resp.PlayerList = []*commMsg.PlayerSimpleInfo{sitter.ToSimpleInfo()}
	GetClientMgr().NotifyButOne(playerList, player.UserID, resp)

	// 游戏准备 添加刚坐下的玩家
	tb.GameHandle.UpdateInfo([]interface{}{player.State, player.UserID})
	if isNew {
		if err := tb.RedisSitCount(Fault); err != nil {
			return TableInfo09
		}
		// 玩家积分转换
		if tb.PlayScore == Unlimited {
			sitter.PlayerInfo.Gold = sitter.PlayerInfo.Coin + 100*sitter.PlayerInfo.YuanBao
		} else {
			sitter.PlayerInfo.Gold = tb.PlayScore
		}

		log.Infof("房间[%v] 牌桌[%v] 游戏[%v] 新增玩家[%v] 座椅[%v]", tb.Rid, tb.Id, tb.Gid, player.UserID, sitter.ID)
		tb.GameHandle.Ready([]interface{}{player})
		// 满员后,开始游戏
		if !tb.isStart && ((tb.MaxSitter == Unlimited && tb.sitCount == Default) || tb.MaxSitter == tb.sitCount) {
			tb.isStart = true
			now := time.Now().Unix()
			wait := tb.OpenTime - now
			if wait < INVALID {
				wait = 1
			}
			time.AfterFunc(time.Duration(wait)*time.Second, func() {
				if tb != nil {
					if tb.GameHandle != nil {
						tb.GameHandle.Start(nil)
					} else {
						log.Errorf("牌桌[%v] 游戏[%v]启动失败!!! ", tb.Id, tb.Gid)
					}
				}
			})
		}
	}
	tb.GameHandle.Scene([]interface{}{player})
	return SUCCESS
}

// GetChairData 获取牌桌信息
func (tb *Table) GetChairData(uid int64) (value any, ok bool) {
	return tb.sitters.Load(uid)
}

// GetChair 获取牌桌信息
func (tb *Table) GetChair(uid int64) *Chair {
	val, ok := tb.sitters.Load(uid)
	if ok {
		chair, ok1 := val.(*Chair)
		if !ok1 {
			return nil
		}
		return chair
	}
	return nil
}

func (tb *Table) GetChairInfos() []*commMsg.PlayerInfo {
	list := make([]*commMsg.PlayerInfo, 0)
	tb.sitters.Range(func(key, value any) bool {
		chair, ok1 := value.(*Chair)
		if !ok1 {
			return true
		}
		list = append(list, chair.PlayerInfo)
		return true
	})
	return list
}

func (tb *Table) RemoveChair(uid int64) {
	tb.sitters.Range(func(key, value any) bool {
		if tempUid, ok := key.(int64); ok && tempUid == uid {
			person := GetPlayerMgr().Get(uid)
			person.InChairId = INVALID
			person.InTableId = INVALID
			person.GameHandle = nil
			if chair, ok1 := value.(*Chair); ok1 {
				if chair.ID == tb.maxChairID {
					tb.maxChairID--
					if tb.maxChairID < 1 {
						tb.maxChairID = 1
					}
				}
				chair = nil
			}
		}
		return true
	})
	tb.sitters.Delete(uid)
	tb.RedisSitCount(Default)
	atomic.SwapInt32(&tb.sitCount, -1)
}

func (tb *Table) AddRedisSit(info redis_cluster.RedisMatchInfo) error {
	key := GetMatchKey(tb.Gid, tb.Rid)
	list, err := redis_cluster.SingleRedis().GetTopCount(context.Background(), key, INVALID)
	if err != nil {
		return err
	}
	for _, item := range list {
		val, ok := item.Member.(string)
		if !ok {
			continue
		}
		info1 := redis_cluster.RedisMatchInfo{}
		err = json.Unmarshal([]byte(val), &info1)
		if err != nil {
			continue
		}
		if info1.Tid == info.Tid {
			return nil
		}
	}
	data, _ := json.Marshal(&info)
	err = redis_cluster.SingleRedis().AddRank(context.Background(), key, FullScore, data)
	return err
}

func (tb *Table) RedisSitCount(mus int) error {
	key := GetMatchKey(tb.Gid, tb.Rid)
	list, err := redis_cluster.SingleRedis().GetTopCount(context.Background(), key, INVALID)
	if err != nil {
		return err
	}
	if 0 == len(list) {
		return fmt.Errorf("no have list. tid:%v", tb.Id)
	}
	info := redis_cluster.RedisMatchInfo{}
	index := 0
	score := float64(0)
	for i, item := range list {
		val, ok := item.Member.(string)
		if !ok {
			continue
		}
		err = json.Unmarshal([]byte(val), &info)
		if err != nil {
			continue
		}
		if info.Tid == tb.Id {
			score = item.Score
			index = i
			break
		}
	}
	if info.Tid != tb.Id {
		return fmt.Errorf("no have tid:%v", tb.Id)
	}
	old := list[index]
	info.SitCount = tb.SitCount()
	data, _ := json.Marshal(info)
	score += float64(mus)
	err = redis_cluster.SingleRedis().ModifyAndAddToSetWithLua(context.Background(), key, old.Member, data, score)
	return err
}

func (tb *Table) ClearChairs() {
	tb.sitters.Range(func(key, value any) bool {
		if uid, ok := key.(int64); ok {
			person := GetPlayerMgr().Get(uid)
			person.InChairId = INVALID
			person.InTableId = INVALID
			person.GameHandle = nil
			person.State = commMsg.PlayerState_PlayerStandUp
			value = nil
		}
		return true
	})
	tb.maxChairID = 0
	atomic.SwapInt32(&tb.sitCount, 0)
}

func (tb *Table) Close() {
	tb.ClearChairs()
	tb.sitters = sync.Map{}
	tb.GameHandle = nil
	tb.TableInfo = nil
}

// NextChair 轮换规则 根据座椅ID大小来轮换
func (tb *Table) NextChair(curUid int64) *Chair {
	//
	val, ok := tb.sitters.Load(curUid)
	if !ok {
		return nil
	}
	sitter, ok1 := val.(*Chair)
	if !ok1 {
		return nil
	}
	curChairID := sitter.ID
	nextChairID := curChairID + 1
	if tb.maxChairID < nextChairID {
		nextChairID = 1
	}
	var nextChair *Chair
	for i := int32(0); i < tb.maxChairID; i++ {
		nextChairID += i
		if tb.maxChairID < nextChairID {
			nextChairID = nextChairID - tb.maxChairID
			if nextChairID < 1 {
				nextChairID = 1
			}
		}
		tb.sitters.Range(func(key, value any) bool {
			if chairInfo, ok2 := value.(*Chair); ok2 && chairInfo.ID == nextChairID {
				nextChair = chairInfo
				return false
			}
			return true
		})
		if nextChair != nil {
			return nextChair
		}
	}

	return nextChair
}
func (tb *Table) NextChairUID(curUid int64) int64 {
	chair := tb.NextChair(curUid)
	if chair == nil {
		return 0
	}
	return chair.UserID
}

func (tb *Table) ChairWork(work func(chair *Chair)) {
	tb.sitters.Range(func(key, value any) bool {
		if sitter, ok := value.(*Chair); ok {
			work(sitter)
		}
		return true
	})
}
func (tb *Table) ChairSettle(name, inning, result string) {
	tb.sitters.Range(func(key, value any) bool {
		chair, ok := value.(*Chair)
		if !ok {
			return true
		}
		chair.PlayerInfo.State = commMsg.PlayerState_PlayerStandUp
		if chair.Total == INVALID || chair.Gain == INVALID {
			return true
		}
		if INVALID < chair.Gain && INVALID < tb.Commission {
			// 收取税收 千分之一
			chair.Gain = int64(1000-tb.Commission) * chair.Gain / 1000
		}
		//通知金币变化
		err := GetDBCmpt().AddRecord(&sqlmodel.Record{
			UID:     chair.UserID,
			Tid:     tb.Id,
			Payment: chair.Gain,
			Code:    CodeSettle,
			Order:   inning,
			Result:  result,
			Remark:  name,
		})
		if err != nil {
			log.Warnf("[%v:%v] inning:[%v] AddRecordReq is failed. err:%v", name, tb.Id, inning, err)
			return true
		}
		changeGold := &gateMsg.NotifyBalanceChangeResp{
			UserID:    chair.UserID,
			Code:      CodeSettle,
			AlterCoin: chair.Gain,
			Coin:      chair.Coin + chair.Gain,
			Reason:    inning,
		}
		chair.Coin = changeGold.Coin
		GetClientMgr().SendTo(chair.UserID, changeGold)
		return true
	})

}

// SitCount 座位总数
func (tb *Table) SitCount() int32 {
	return atomic.LoadInt32(&tb.sitCount)
}

// GetPlayList 获取玩家列表
func (tb *Table) GetPlayList() []int64 {
	list := make([]int64, 0)
	tb.sitters.Range(func(key, value any) bool {
		if uid, ok := key.(int64); ok {
			list = append(list, uid)
		}
		return true
	})

	return utils.Unique(list)
}

// GetLookList 获取所有旁观玩家列表
func (tb *Table) GetLookList() []int64 {
	list := make([]int64, 0)
	tb.sitters.Range(func(key, value any) bool {
		if chair, ok := value.(*Chair); ok {
			if chair.PlayerInfo.State == commMsg.PlayerState_PlayerLookOn {
				list = append(list, chair.PlayerInfo.UserID)
			}
		}
		return true
	})
	return utils.Unique(list)
}

// GetAllList 获取所有玩家列表
func (tb *Table) GetAllList() []int64 {
	list := make([]int64, 0)
	tb.sitters.Range(func(key, value any) bool {
		if uid, ok := key.(int64); ok {
			list = append(list, uid)
		}
		return true
	})
	return utils.Unique(list)
}

// GetOrderList 获取排序的玩家队列
func (tb *Table) GetOrderList() []int64 {
	list := make([]int64, 0)
	for i := int32(1); i <= tb.maxChairID; i++ {
		tb.sitters.Range(func(key, value any) bool {
			if chairInfo, ok := value.(*Chair); ok && chairInfo.ID == i {
				list = append(list, chairInfo.UserID)
				return false
			}
			return true
		})
	}
	return utils.Unique(list)
}

// CalibratingRemain 校准剩余次数
func (tb *Table) CalibratingRemain(delCount int32) {
	if tb.Remain <= INVALID {
		return
	}
	remain, err := GetDBCmpt().EraseRemain(tb.Id, delCount)
	if err != nil {
		log.Errorf("[Table] CalibratingRemain DecreaseGameRun err:%v", err)
		return
	}
	tb.Remain = remain
}

// CheckChairsNoLive 检测桌椅上已经离开的玩家
func (tb *Table) CheckChairsNoLive() (noLiveIDs []int64) {
	noLiveIDs = make([]int64, 0)
	tb.sitters.Range(func(key, value any) bool {
		if chair, ok := value.(*Chair); ok && INVALID != chair.LeaveTime {
			noLiveIDs = append(noLiveIDs, chair.UserID)
			if chair.ID == tb.maxChairID {
				tb.maxChairID--
				if tb.maxChairID < 1 {
					tb.maxChairID = 1
				}
			}
			chair = nil
		}
		return true
	})
	for _, uid := range noLiveIDs {
		tb.sitters.Delete(uid)
	}
	atomic.SwapInt32(&tb.sitCount, -int32(len(noLiveIDs)))
	if tb.SitCount() == INVALID {
		tb.isStart = false
	}
	return
}

/////////////////////////////////////////////////////////////////////////
