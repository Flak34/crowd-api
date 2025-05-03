package project_service

type ListProjectsDTO struct {
	Filter ProjectsFilter
	Page   *Page
	Sort   *Sort
}

type ProjectsFilter struct {
	CreatorID int
	Status    string
}

type Page struct {
	PerPage int
	PageNum int
}

type Sort struct {
	SortField string
	Desc      bool
}
