package domain

// 生成订单序列，数据库层没有使用自增
type Sequence struct {
	Name   string
	NextId int
}
