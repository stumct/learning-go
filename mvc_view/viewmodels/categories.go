package viewmodels

type Categories struct {
	Title      string
	Active     string
	Categories []Category
}

type Category struct {
	ImageUrl      string
	Title         string
	Description   string
	IsOrientRight bool
	Id            int
}

func GetCategories() Categories {
	result := Categories{
		Title:  "Lemonade Stand Society - Shop",
		Active: "shop",
	}

	juiceCategory := Category{
		Id:            1,
		ImageUrl:      "juice.png",
		Title:         "Juicy Juice",
		Description:   "Have some juicy juice",
		IsOrientRight: false,
	}
	foodCategory := Category{
		Id:            2,
		ImageUrl:      "food.png",
		Title:         "Foody FOod",
		Description:   "Have some foooooood",
		IsOrientRight: true,
	}
	stuffCategory := Category{
		Id:            3,
		ImageUrl:      "stuff.png",
		Title:         "Stuffy Stuff",
		Description:   "Have some stuff",
		IsOrientRight: false,
	}

	result.Categories = []Category{
		juiceCategory,
		foodCategory,
		stuffCategory,
	}
	return result
}
