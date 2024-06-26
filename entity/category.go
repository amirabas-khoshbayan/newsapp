package entity

type Category string

const (
	Business Category = "business"
	Health   Category = "health"
	Sports   Category = "sports"
	Politics Category = "politics"
)

func CategoryList() []Category {
	return []Category{Business, Health, Sports, Politics}
}
