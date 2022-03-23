package repository

//NewMemoryRepository
func NewPostgresRepository() CaseRepository {
	return &repo{}
}
