package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type UtilRepoImplementation struct {
	db *gorm.DB
}

func NewUtilRepository(db *gorm.DB) UtilRepository {
	return &UtilRepoImplementation{
		db: db,
	}
}

func (r *UtilRepoImplementation) FindAllJurusan() (*[]domain.Jurusan, error){
	var jurusan []domain.Jurusan
	if err := r.db.Find(&jurusan).Error; err != nil{
		return nil, err
	} 

	return &jurusan, nil
}
