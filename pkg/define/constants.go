package define

const (
	IntZero  = 0
	IntOne   = 1
	IntTwo   = 2
	IntThree = 3
)

const (
	ErrParam = "參數錯誤"
	ErrDB    = "數據庫操作出錯"
	ErrCache = "緩存操作出錯"
	StatusOk = "successful"
)

// db连接关键字
const (
	QiPaiDB  = "qipaidb"
	StatDB   = "statdb"
	VipCache = "vipcache"
)

// 错误提示
const (
	RetryLaterError      = "稍後重試"
	OperationFrequentTip = "操作頻繁，請稍等"
	IllConditionedTip    = "未達到VIP所需條件"
	SalaryGetTip         = "非VIP無法領取工資"
	GetSalaryLaterTip    = "請稍後領取今日工資"
	WatchADLaterTip      = "請稍後觀看廣告"
	VipIsExpiredTip      = "會員已過期"
	MaxAdCountTip        = "廣告觀看數量達到上限"
	RenewVipFailed       = "續費失敗，請稍候重試"
	VipADCoolingTip      = "廣告時間冷卻中，無法觀看視頻廣告"
)

// 年月日常量
const (
	YMDLayout         = "2006-01-02"
	LayoutDayForCache = "20060102"
	OneDaySeconds     = 86400
)
