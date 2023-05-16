package service

import (
	"GraduationDesign/dao"
	"GraduationDesign/models"
)

type CategoryService struct {
}

var categoryDao dao.CategoryDao

func (CategoryService) AddCategory(category models.CategoryParam) error {
	err := categoryDao.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (CategoryService) GetAllCategory() ([]models.Category, error) {
	var categorys []models.Category
	result, err := categoryDao.GetAllCategorys()
	if err != nil {
		return nil, err
	}
	for key, value := range result {
		categorys = append(categorys, models.Category{key, value})
	}
	return categorys, err
}
func (CategoryService) GetAllAdviceCategory() ([]models.Category, error) {
	var categorys []models.Category
	result, err := categoryDao.GetAllAdviceCategorys()
	if err != nil {
		return nil, err
	}
	for key, value := range result {
		categorys = append(categorys, models.Category{key, value})
	}

	return categorys, err

}

func (CategoryService) DeleteCategory(category models.CategoryParam) error {
	err := categoryDao.DeleteCategory(category)
	if err != nil {
		return err
	}
	return nil
}
