package models

import "github.com/gubesch/go-quiz/migration"

type Category struct {
	ID 				int		`json:"id,omitempty"`
	CategoryName 	string 	`json:"category_name,omitempty"`
}

func ShowCategories() (allcategories []Category, err error) {
	db := migration.GetDbInstance()
	//defer db.Close()
	categoryiesQuerie,err := db.Query("SELECT * FROM categories;")
	defer categoryiesQuerie.Close()
	if err != nil{
		return nil,err
	}
	for categoryiesQuerie.Next(){
		var category Category
		err = categoryiesQuerie.Scan(&category.ID,&category.CategoryName)
		if err != nil{
			return nil,err
		}
		allcategories = append(allcategories, category)
	}
	return
}

func (c *Category) CreateNewCategory() (err error){
	db := migration.GetDbInstance()
	//defer db.Close()
	stmtInsertCategory,err := db.Prepare("INSERT INTO `categories` (`category_name`) VALUES (?);")
	defer stmtInsertCategory.Close()
	if err != nil {
		return
	}
	_,err = stmtInsertCategory.Exec(c.CategoryName)
	if err != nil{
		return
	}
	return nil
}

func (c *Category) EditCategory() (err error){
	db := migration.GetDbInstance()
	//defer db.Close()
	stmtEditCategory,err := db.Prepare("UPDATE categories SET category_name = ? WHERE categories.id = ?;")
	defer stmtEditCategory.Close()
	if err != nil{
		return
	}
	_,err = stmtEditCategory.Exec(c.CategoryName,c.ID)
	if err != nil{
		return
	}
	return
}

func (c *Category) DeleteCategory() (err error){
	db := migration.GetDbInstance()
	//defer db.Close()
	stmtDeleteCategory,err := db.Prepare("DELETE FROM categories WHERE id=?;")
	defer stmtDeleteCategory.Close()
	if err != nil{
		return
	}
	_,err=stmtDeleteCategory.Exec(c.ID)
	if err != nil{
		return
	}
	return nil
}
