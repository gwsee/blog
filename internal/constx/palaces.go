package constx

const (
	PalacesMemoInit = iota //初始化
	PalacesMemoDone        //已完成
)
const (
	_                = iota
	PalacesTodoNum   //按次数
	PalacesTodoShort //短期 即有截止日期
	PalacesTodoLong  //长期 有开始时间 没有截止时间
)
const PalacesTodoNumMin = 1 //按次数最低数量
const (
	PlacesTodoInit = iota //初始化
	PlacesTodoDone        //已完成
)
