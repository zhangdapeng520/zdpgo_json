package query

const (
	Null   Type = iota // json中的null类型
	False              // json中的false类型
	Number             // json中的number类型
	String             // json中的string类型
	True               // json中的true类型
	JSON               // json中的嵌套json类型
)
