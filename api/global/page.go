package global

func (o *PageInfo) GetOffset() int64 {
	if o == nil {
		return 0
	}
	return (o.PageNum - 1) * o.PageSize
}
