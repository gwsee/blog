package global

func (o *PageInfo) GetOffset() int64 {
	return (o.PageNum - 1) * o.PageSize
}
func (o *PageInfo) GetLimit() int64 {
	return o.PageSize
}
