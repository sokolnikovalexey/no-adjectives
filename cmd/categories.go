package cmd

import "errors"

func GetAvailableCategories() []CategoryItem {
	var categories []CategoryItem
	categories = append(categories, CategoryItem{"Yoshkar-Ola/index", "Йошкар-Ола"})
	categories = append(categories, CategoryItem{"science", "Наука"})
	categories = append(categories, CategoryItem{"computers", "IT"})
	categories = append(categories, CategoryItem{"finances", "Финансы"})
	categories = append(categories, CategoryItem{"auto", "Авто"})
	categories = append(categories, CategoryItem{"politics", "Политика"})
	categories = append(categories, CategoryItem{"sport", "Спорт"})
	return categories
}

func GetCategoryById(id string) (item CategoryItem , err error) {
	for _, category := range GetAvailableCategories() {
		if category.Id == id {
			item = category
			return
		}
	}
	return item, UnknownCategoryError
}

var UnknownCategoryError = errors.New("unknown category")

type CategoryItem struct {
	Id 		string
	Title 	string
}
