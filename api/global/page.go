package global

func (o *PageInfo) GetOffset() int64 {
	return (o.PageNum - 1) * o.PageSize
}
