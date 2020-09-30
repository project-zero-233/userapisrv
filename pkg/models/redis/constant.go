package redis

const (
	// cache name
	cacheNameUserVipAD      = "user:vip:%v:%v:ad"         // user:vip:$game_id:$user_id:ad  广告缓存
	cacheNameUserVipCache   = "user:vip:%v:%v:cache"      // user:vip:$game_id:$user_id:cache 作为其他额外信息存储
	cacheNameUserVipInfo    = "user:vip:%v:%v:info"       // user:vip:$game_id:$user_id:info	用户VIP信息缓存
	cacheNameUserVipSalary  = "user:vip:%v:%v:salary:%v"  // user:vip:$game_id:$user_id:salary:$logDate	用户每日salary缓存
	cacheNameUserVipCooling = "user:vip:%v:%v:cooling:%v" // user:vip:$game_id:$user_id:cooling:$type	用户观看VIP视频冷却

	// expire time
	expireAtNonVip = 15 // 非vip保留15天数据
	expireAtSalary = 1  // salary 保留1天数据
	expireAtCache  = 15 // 保留15天数据
)

const (
	// user vip members
	userVipADCount        = "ad_count"
	userVipRenewADCount   = "renew_ad_count"
	userVipCacheAD        = "cache_ad"
	userVipOldVipID       = "old_vip_id"
	userVipID             = "vip_id"
	userVipName           = "vip_name"
	userVipADTarget       = "vip_ad_target"
	userVipFrom           = "vip_from"
	userVipFillAD         = "fill_ad"
	userVipLastRemainTime = "last_remain_time" // 记录VIP时间变化的最后一次变化，也就是当时间发生变化时记录了上一次的剩余时间

	// salary标识
	ReceiveSalary   = 1 // 领取了salary就是1
	UnReceiveSalary = 0 // 未领取就是0

	// 广告类型
	TypeADNormal = 1 // 普通广告
	TypeADRenew  = 2 // 续费广告
	TypeADFill   = 3 // 补充广告，不作为数据统计，仅仅作为补充使用

	// flag标识
	UerVipStatusNormal      = -1 // user vip 正常是-1
	UerVipStatusNotInit     = 0  // user vip 未初始化是0
	UerVipStatusEmptySalary = 1  // user vip 的salary为空是1
	UerVipStatusErr         = 2  // user vip 发生错误是2

	// vip 来源
	VipFromAD   = 1 // 来自广告
	VipFromProp = 2 // 来自道具

	// 非VIP ID，作为消除魔法值标识用户的VIP ID是0
	VipIDZero = 0

	// VIP观看广告冷却时间，每观看2次冷却220s
	VipCoolingCount    = 2   // 每观看2次开启冷却
	VipCoolingSecond   = 210 // 210s 冷却
	VipCoolingGetVip   = "getVip"
	VipCoolingRenewVip = "renewVip"
)

/**
 * 任务统计---------------------
**/
const (
	// 任务缓存
	FieldNameUserStatistics              = "%v:%v:%v:%v"                        // $module:$gameID:$gameAreaID:$target   用户统计信息field
	cacheNameUserStatistics              = "userstatistics:%v:%v"               // userstatistics:$user_id:$log_date  用户统计信息key
	cacheNameUserStatisticsUpdatedCommon = "userstatistics:update:common:%v:%v" // userstatistics:update:common:$user_id:$log_date 用户信息是否被更新
)

const (
	expireAtUserTask = 1 // 用户每日任务
)

const (
	UserStatisticChanged = 1 // 用户统计信息变化
)
