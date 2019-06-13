package model

type Car struct {
	// 大写开头的为公有字段，能跨包访问
	Brand    string
	TopSpeed int
	// 小写开头的为私有字段，只能同包中访问
	horsePower int
}
