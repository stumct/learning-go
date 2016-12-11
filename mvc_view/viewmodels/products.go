package viewmodels

type Products struct {
	Title    string
	Active   string
	Products []Product
}

func GetProducts(id int) Products {
	var result Products
	result.Active = "shop"
	var shopName string
	switch id {
	case 1:
		shopName = "Juice"
	case 2:
		shopName = "Supply"
	case 3:
		shopName = "Advertising"
	}
	results.Title = "Lemonade Stand Society - " + shopName + "Shop"

	lemonJuice := MakeLemonJuiceProduct()
	appleJuice := MakeAppleJuiceProduct()
	kwiJuice := MakeKiwiJuiceProduct()
}
