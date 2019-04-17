package counters
//alertCounter是一个未公开的类型，这个类型用于保存告警计数
type alertCount int

func New(value int) alertCount{
	return alertCount(value)
}