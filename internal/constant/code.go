package constant

//var (
//	ChineseNum = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
//)

var StatusText = map[int]string{
	Title001: "提示",
	Title002: "警告",
	Title003: "一般错误",
	Title004: "严重错误",
	Title005: "致命错误",
	Title006: "♛恭喜♚",
	Title007: "感谢",
	Title008: "系统",
	Title009: "匹配失败",

	Flag0001: "ok",
	Flag0002: "success",
	Flag0003: "fail",
	Flag0004: "over",

	Register01: "♚♛♝♞♜♟注册成功♔♕♗♘♖♟",
	Register02: "注册失败",
	Register03: "注册出错",
	Register04: "注册",
	Register05: "该用户名已被注册",
	Register06: "用户名长度超出",
	Register07: "用户名长度不对",
	Register08: "密码长度不对",
	Register09: "不能包含空格符",
	Register10: "用户名不能为空",
	Register11: "密码不能为空",

	Login01: "ꕥ登录成功ꕥ",
	Login02: "登录失败!",
	Login03: "登录出错!",
	Login04: "您的账号已经在异地登录了!",
	Login05: "登录",
	Login06: "请求无效,请您重新登录后重试!",
	Login07: "无法获取用户信息,请您联系客服!",
	Login08: "账号或密码错误!",
	Login09: "您所参与的游戏没有结束!",
	Login10: "您的账号正被异地强登,请及时联系客服", //"重连成功"
	Login11: "TOKEN已过期,请您退出后重新登录",
	Login12: "已在异地登录,本次登录已被禁止",
	Login13: "操作频繁,请延迟3~5秒,再重试",
	Login14: "签到",
	Login15: "TOKEN无效",
	Login16: "登出异常",

	NetworkErr01: "连接失败",
	NetworkErr02: "发送失败",
	NetworkErr03: "接收失败",
	NetworkErr04: "监听失败",
	NetworkErr05: "重连失败",
	NetworkErr06: "认证失败",
	NetworkErr07: "频繁发送",
	NetworkErr08: "无效地址",
	NetworkErr09: "协议无效",
	NetworkErr10: "连接成功",
	NetworkErr11: "重连成功",
	NetworkErr12: "发送成功",
	NetworkErr13: "接收成功",

	ClassInfo01: "获取游戏分类列表失败!",
	ClassInfo02: "游戏分类ID无效!",
	ClassInfo03: "空荡荡页面!",

	TableInfo01: "获取牌桌列表失败!",
	TableInfo02: "牌桌信息配置失败!",
	TableInfo03: "获取游戏信息失败!",
	TableInfo04: "游戏已解散,请您移步至其他游戏!",
	TableInfo05: "换桌失败!",
	TableInfo06: "❣您所参与的游戏本轮还没结束❣",
	TableInfo07: "牌桌已存在",
	TableInfo08: "创建失败:已超过房间最大人数",
	TableInfo09: "无法获取有效牌桌",
	TableInfo10: "牌桌不足,需房主补充牌桌",
	TableInfo11: "座椅上的数据出错",
	TableInfo12: "金币不足,配桌失败",
	TableInfo13: "牌桌删除失败",
	TableInfo14: "牌桌创建失败",

	Setting01: "配置信息错误!",
	Setting02: "座位数过少!",
	Setting03: "游戏类型不正确!",
	Setting04: "游戏种类不正确!",
	Setting05: "无法创建房间!",
	Setting06: "㊗房间创建成功㊗",
	Setting07: "㊖缴纳房费㊖",
	Setting08: "㊖退还房费㊖",
	Setting09: "缴纳房费失败!",
	Setting10: "退还房费失败!",
	Setting11: "您已经配置过游戏了,请勿重复配置!",
	Setting12: "名字不能为空!",
	Setting13: "名字不能包含空格!",
	Setting14: "可用筹码必须大于0",
	Setting15: "携带筹码不得低于底注!",
	Setting16: "游戏暂未开售,详情请您留意官方公告!",
	Setting17: "游戏场数不能为0!",
	Setting18: "游戏场数不得大于10000!",
	Setting19: "名字长度不能超过6个字符!",
	Setting20: "修改名字出错",

	Room01: "密码有误!",
	Room02: "登录房间成功",
	Room03: "登录房间失败",
	Room04: "登录房间出错,请您重新拉取房间信息!",
	Room05: "游戏已解散!",
	Room06: "该游戏已经满人了,请您关注其他游戏,会有不一样的惊喜哟!",
	Room07: "没有多余的桌子了,请您在此耐心等候!",
	Room08: "这是崭新的一桌,请您耐心等候其他玩家!",
	Room09: "换桌失败,玩家未处于空闲状态!",
	Room10: "换桌成功!",
	Room11: "您没有房主权限!",
	Room12: "平台无效,无法进入该游戏",
	Room13: "没有房间权限,无法进入该游戏",
	Room14: "通用房卡",
	Room15: "牌桌数已达到最大限制",
	Room16: "房间不存在",
	Room17: "添加牌桌失败:游戏未开放或已关闭",
	Room18: "房间创建失败",

	User01: "玩家",
	User02: "无效请求",
	User03: "玩家数据出错",
	User04: "玩家断线",
	User05: "玩家加入",
	User06: "玩家离开",
	User07: "玩家退出",
	User08: "玩家加码",
	User09: "玩家同意",
	User10: "玩家拒绝",
	User11: "您还不是VIP用户,暂不能使用该权限!",
	User12: "您的金币不足",
	User13: "充值成功",
	User14: "提现成功",
	User15: "㊖账户冻结,详情请联系客服☎",
	User16: "已取消上庄!",
	User17: "囍🌹🌹上庄成功🌹🌹囍",
	User18: "您已成功加入抢庄队伍！",
	User19: "您已被移出抢庄队伍！",
	User20: "您的权限不足!",
	User21: "请勿重复操作!",
	User22: "感谢您的反馈!",
	User23: "您的房卡不足,请您前往商城购买!",
	User24: "购买房卡成功,预祝您游戏愉快!",
	User25: "购买房卡失败:无效的用户信息!",
	User26: "购买房卡失败:无效的求购信息!",
	User27: "暂无新邮件",
	User28: "您今天已经签到过了",
	User29: "加入游戏队列失败",
	User30: "加入游戏队列成功,请耐心稍等",

	Mysql01: "数据库",
	Mysql02: "数据库登录成功",
	Mysql03: "数据库登录失败",
	Mysql04: "数据库登录出错",
	Mysql05: "数据库添加失败",
	Mysql06: "数据库删除失败",
	Mysql07: "数据库查询失败",
	Mysql08: "数据库修改失败",
	Mysql09: "数据库操作成功",
	Mysql10: "数据库清库成功",
	Mysql11: "结算出错!",
	Mysql12: "无法配置游戏,请您联系客服!",
	Mysql13: "领取奖励失败",
	Mysql14: "房卡置换失败",
	Mysql15: "删除邮件失败",
	Mysql16: "上一次记录查找失败",

	Redis01: "no_max_score",
	Redis02: "essence is nil",

	Game01: "房间等级不匹配!",
	Game02: "游戏内未能查找到相关玩家!",
	Game03: "⏰过了下注时间⏰",
	Game04: "您的操作失效:不符合游戏场景!",
	Game05: "下注失败: 金币不足!",
	Game06: "下注失败: 无效区域!",
	Game07: "下注失败: 无效金额!",
	Game08: "没有必要的进入规则",
	Game09: "未准备被踢出房间!",

	Game10: "꧁抢庄成功꧂",
	Game11: "抢庄失败:超过限定人数",
	Game12: "👓不符合看牌规则👓",
	Game13: "抱歉,游戏已满员༺๑๑༻",
	Game14: "游戏正在进行",
	Game15: "未满足准入条件",
	Game16: "游戏未开放",
	Game17: "游戏维护中♠♣♥♦",
	Game18: "游戏已暂停",
	Game19: "游戏已关闭",

	Game20: "游戏已结束[𝕲𝖆𝖒𝖊𝕺𝖛𝖊𝖗]",
	Game21: "☕还没轮到您,请您耐心等候☕",
	Game22: "本轮没有庄家,游戏将重新开始!",
	Game23: "游戏中,请勿离桌!",
	Game24: "☕未到下注阶段☕",
	Game25: "您已抢过庄了,请耐心等候结果!",
	Game26: "您已是当前庄家,请勿参与抢庄!",
	Game27: "下注失败:庄家不可以下注!",
	Game28: "下注阶段不能下庄!",
	Game29: "定庄阶段不能下庄!",

	Game30: "🛀您已经处于准备状态了🛀",
	Game31: "您已经取消准备了!",
	Game32: "💪您已经下过注了㊟",
	Game33: "下注失败:座位无效!",
	Game34: "❤其他玩家未准备,请您耐心等待❤",
	Game35: "跟注失败:您的金币不足!",
	Game36: "第一轮,禁止比牌!",
	Game37: "☹无效操作☹",
	Game38: "本桌游戏已经飞出✈✈✈✈银河系❂",
	Game39: "༺๑๑༻本桌已经满员,请您尝试其他桌位上的游戏༺๑๑༻",
	Game40: "🎲由于您未掷骰子,系统掷骰子🎲",
	Game41: "牌值有误,请您核对后重试!",
	Game42: "游戏出现故障,即将重置场景!",
	Game43: "第一轮,禁止过牌!",
	Game44: "您已操作过了!",
	Game45: "托管状态,禁止主动出牌!",
	Game46: "您的金币过低无法上庄",
	Game47: "进入游戏失败",
	Game48: "稍等",
	Game49: "等待匹配",
	Game50: "匹配成功",
	Game51: "匹配失败",
	Game52: "玩家未选择武将!",
	Game53: "本局已有玩家,正在重新匹配!",
	Game54: "玩家断线,正在重新匹配!",
	Game55: "敌方未选择武将,正在重新匹配!",
	Game56: "您所参与的游戏 本轮已结束!",
	Game57: "您还未选择武将,不能进行匹配",
	Game58: "该武将已在阵列,不可重复使用",
	Game59: "下注金额不得低于底注",
	Game60: "无效的场景操作",
	Game61: "设置时长过长",
	Game62: "设置时长过短",
	Game63: "操作不符合规则",

	Chat001: "申请加好友失败",
	Chat002: "获取邀请信息失败",
	Chat003: "邀请对战失败",
	Chat004: "获取群列表失败",
	Chat005: "获取好友列表失败",
	Chat006: "建群失败",
	Chat007: "处理邀请失败: 邀请不存在",
	Chat008: "处理邀请失败",
	Chat009: "删除好友失败",
	Chat010: "获取好友在线状态失败",
	Chat011: "获取好友申请列表失败",
	Chat012: "申请入群失败:您已经申请过了",
	Chat013: "解散群失败",
	Chat014: "分配积分失败: 成员不存在",
	Chat015: "分配积分失败",
	Chat016: "处理好友申请失败: 申请不存在",
	Chat017: "处理好友申请失败",
	Chat018: "邀请入群失败",
	Chat019: "获取群列表失败",
	Chat020: "发送文本消息失败",
	Chat021: "切换职务失败: 成员不存在",
	Chat022: "切换职务失败",
	Chat023: "处理入群申请失败: 申请不存在",
	Chat024: "处理入群申请失败",
	Chat025: "获取我的群失败",
	Chat026: "获取申请列表失败",
	Chat027: "已发送",
	Chat028: "消息发送失败",
	Chat029: "已发送,但对方已离线",
	Chat030: "获取聊天记录失败",
	Chat031: "申请已提交,请等待审批",
	Chat032: "审批已通过",
	Chat033: "没有审批人员",
	Chat034: "已经审批过了",
	Chat035: "抱歉,您没有审批的权限",
	Chat036: "抱歉,您没有该权限",
	Chat037: "抱歉,同级权限无法操作",
	Chat038: "抱歉,您不是群主无法操作",
	Chat039: "不能邀请自己",
	Chat040: "邀请入群失败:成员已在群里了",
	Chat041: "删除群成员失败",
	Chat042: "邀请入群失败:该成员信息无效",
	Chat043: "成功发起邀请,请等待对方通过",
	Chat044: "您的邀请已被接纳",
	Chat045: "您的邀请已被拒绝",
	Chat046: "群主无法切换级别",
	Chat047: "不能剔出群主",

	Service001: "SDK无效",
	Service002: "PID无效",
	Service003: "CID无效",
	Service004: "参数无效",
}

// 标题
const (
	_ = 1000 + iota //标题
	Title001
	Title002
	Title003
	Title004
	Title005
	Title006
	Title007
	Title008
	Title009
	//Title010
)

// 标识
const (
	_ = 2000 + iota //标识
	Flag0001
	Flag0002
	Flag0003
	Flag0004
)

// 注册
const (
	_ = 3000 + iota //注册
	Register01
	Register02
	Register03
	Register04
	Register05
	Register06
	Register07
	Register08
	Register09
	Register10
	Register11
)

// 登录
const (
	_ = 3100 + iota //登录
	Login01
	Login02
	Login03
	Login04
	Login05
	Login06
	Login07
	Login08
	Login09
	Login10
	Login11
	Login12
	Login13
	Login14
	Login15
	Login16
)

// 重连
const (
	_ = 3200 + iota //重连
	NetworkErr01
	NetworkErr02
	NetworkErr03
	NetworkErr04
	NetworkErr05
	NetworkErr06
	NetworkErr07
	NetworkErr08
	NetworkErr09
	NetworkErr10
	NetworkErr11
	NetworkErr12
	NetworkErr13
	NetworkErr14
	NetworkErr15
	NetworkErr16
	NetworkErr17
	NetworkErr18
	NetworkErr19
)

const (
	_ = 4000 + iota //配置信息
	Setting01
	Setting02
	Setting03
	Setting04
	Setting05
	Setting06
	Setting07
	Setting08
	Setting09
	Setting10
	Setting11
	Setting12
	Setting13
	Setting14
	Setting15
	Setting16
	Setting17
	Setting18
	Setting19
	Setting20
)

const (
	_ = 4100 + iota //配置信息-类型
	ClassInfo01
	ClassInfo02
	ClassInfo03
)

const (
	_ = 4200 + iota //桌台信息
	TableInfo01
	TableInfo02
	TableInfo03
	TableInfo04
	TableInfo05
	TableInfo06
	TableInfo07
	TableInfo08
	TableInfo09
	TableInfo10
	TableInfo11
	TableInfo12
	TableInfo13
	TableInfo14
)

const (
	_ = 5000 + iota //房 间 信息
	Room01
	Room02
	Room03
	Room04
	Room05
	Room06
	Room07
	Room08
	Room09
	Room10
	Room11
	Room12
	Room13
	Room14
	Room15
	Room16
	Room17
	Room18
)

const (
	_ = 6000 + iota //数据库
	Mysql01
	Mysql02
	Mysql03
	Mysql04
	Mysql05
	Mysql06
	Mysql07
	Mysql08
	Mysql09
	Mysql10
	Mysql11
	Mysql12
	Mysql13
	Mysql14
	Mysql15
	Mysql16

	Redis01
	Redis02
)

const (
	_ = 10000 + iota //游戏清场
	User01
	User02
	User03
	User04
	User05
	User06
	User07
	User08
	User09
	User10
	User11
	User12
	User13
	User14
	User15
	User16
	User17
	User18
	User19
	User20
	User21
	User22
	User23
	User24
	User25
	User26
	User27
	User28
	User29
	User30
)
const (
	_ = 20000 + iota //游戏清场
	Game01
	Game02
	Game03
	Game04
	Game05
	Game06
	Game07
	Game08
	Game09

	Game10
	Game11
	Game12
	Game13
	Game14
	Game15
	Game16
	Game17
	Game18
	Game19
	Game20

	Game21
	Game22
	Game23
	Game24
	Game25
	Game26
	Game27
	Game28
	Game29
	Game30
	Game31
	Game32
	Game33
	Game34
	Game35
	Game36
	Game37
	Game38
	Game39
	Game40

	Game41
	Game42
	Game43
	Game44
	Game45
	Game46
	Game47
	Game48
	Game49
	Game50
	Game51
	Game52
	Game53
	Game54
	Game55
	Game56
	Game57
	Game58
	Game59
	Game60
	Game61
	Game62
	Game63
)

const (
	_ = 30000 + iota //游戏清场
	Service001
	Service002
	Service003
	Service004
)
const (
	_ = 40000 + iota //游戏清场
	Chat001
	Chat002
	Chat003
	Chat004
	Chat005
	Chat006
	Chat007
	Chat008
	Chat009
	Chat010
	Chat011
	Chat012
	Chat013
	Chat014
	Chat015
	Chat016
	Chat017
	Chat018
	Chat019
	Chat020
	Chat021
	Chat022
	Chat023
	Chat024
	Chat025
	Chat026
	Chat027
	Chat028
	Chat029
	Chat030
	Chat031
	Chat032
	Chat033
	Chat034
	Chat035
	Chat036
	Chat037
	Chat038
	Chat039
	Chat040
	Chat041
	Chat042
	Chat043
	Chat044
	Chat045
	Chat046
	Chat047
	Chat048
	Chat049
	Chat050
)

/* 主要用于求取某位上是否 是1值
const (
	Exponent1 int = iota
	Exponent2 int = 1 << (1 * iota)
	Exponent3
	Exponent4
	Exponent5
	Exponent6
	Exponent7
	Exponent8
	Exponent9
	Exponent10
	Exponent11
	Exponent12
	Exponent13
	Exponent14
	Exponent15
	Exponent16
	Exponent17
	Exponent18
	Exponent19
	Exponent20
	Exponent21
	Exponent22
	Exponent23
	Exponent24
	Exponent25
	Exponent26
	Exponent27
	Exponent28
	Exponent29
	Exponent30
	Exponent31
	Exponent32
)*/
