package counters

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-13 08:36
 *      email  : 398757724@qq.com
 *      Desc   : 包含计数器功能
 ***************************************/
// alertCounter 是一个未公开的类型，这个类型用于保存告警计数
type alertCounter int

// new返回一个未公开的alertCounter类型的值
func New(value int)alertCounter  {
	return alertCounter(value)
}