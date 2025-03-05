package service

type ICategoryService interface {
	GetAllCategories() string
	GetCategoryById() string
	CreateCategory() string
	UpdateCategory() string
	DeleteCategory() string
}

type CategoryService struct {
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (c *CategoryService) GetAllCategories() string {
	return "Get all categories from service"
}

func (c *CategoryService) GetCategoryById() string {
	return "Get by id service"
}

func (c *CategoryService) CreateCategory() string {
	return "Create in service"
}

func (c *CategoryService) UpdateCategory() string {
	return "Update in service"
}

func (c *CategoryService) DeleteCategory() string {
	return "Delete in service"
}
