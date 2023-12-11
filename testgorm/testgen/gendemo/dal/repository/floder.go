package repository

import (
	"gorm.io/gen"
)

type FloderRepositoryG interface {
	// FIndByUserId 查出属于某一个用户的id

	// sql(select * from @@table where user_id = @UserID)
	FindByUserId(UserID int64) ([]gen.T, error)
	//sql(delete from @@table where id = @id)
	DelOneByID(id int64) error
}

//
//type FloderRepository interface {
//	FirstById(id int64) (*model.Floder, error)
//}
//
//func NewFloderRepository(repository *Repository) FloderRepository {
//	return &floderRepository{
//		Repository: repository,
//	}
//}
//
//type floderRepository struct {
//	*Repository
//}
//
//func (r *floderRepository) FirstById(id int64) (*model.Floder, error) {
//	var floder model.Floder
//	// TODO: query db
//	return &floder, nil
//}
