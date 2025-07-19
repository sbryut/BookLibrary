package dto

type CreateGenreDTO struct {
	Name string `json:"name"`
}

type UpdateGenreDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
