package project_repository

type CreateProjectAnnotatorDTO struct {
	ProjectID   int
	AnnotatorID int
	TaskIDs     []int
}
