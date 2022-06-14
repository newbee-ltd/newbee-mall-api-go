package enum

type GoodsCategoryLevel int8

const (
	Default    GoodsCategoryLevel = 0
	LevelOne   GoodsCategoryLevel = 1
	LevelTwo   GoodsCategoryLevel = 2
	LevelThree GoodsCategoryLevel = 3
)

func (g GoodsCategoryLevel) Info() (int, string) {
	switch g {
	case LevelOne:
		return 1, "一级分类"
	case LevelTwo:
		return 2, "二级分类"
	case LevelThree:
		return 3, "三级分类"
	default:
		return 0, "error"
	}
}

func (g GoodsCategoryLevel) Code() int {
	switch g {
	case LevelOne:
		return 1
	case LevelTwo:
		return 2
	case LevelThree:
		return 3
	default:
		return 0
	}
}
