package controller


type Category struct {
	Architecture int
	Athletics int
	Foodie int
	Nature int
	People int
	Travel int
}

func New() *Category {
	return &Category{1,2,3,4,5,6}
}

