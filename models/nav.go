package models

type Nav struct {
	Id        int
	Title     string
	Link      string
	Position  int    //导航所在页面位置  0顶部 1中部 2 底部
	IsOpennew int    //是否新窗口打开
	Relation  string //关联商品
	Sort      int
	Status    int
	AddTime   int
}

func (Nav) TableName() string {
	return "nav"
}
