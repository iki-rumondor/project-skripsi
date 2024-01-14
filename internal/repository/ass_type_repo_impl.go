package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type AssTypeRepoImplementation struct {
	db *gorm.DB
}

func NewAssTypeRepository(db *gorm.DB) AssTypeRepository {
	return &AssTypeRepoImplementation{
		db: db,
	}
}

func (r *AssTypeRepoImplementation) CreateAssType(model *domain.AssessmentType) error {
	return r.db.Create(&model).Error
}

// func (r *AssTypeRepoImplementation) PaginationAssType(pagination *domain.Pagination) (*domain.Pagination, error) {
// 	var results []domain.AssessmentType
// 	var totalRows int64 = 0

// 	if err := r.db.Model(&domain.AssessmentType{}).Count(&totalRows).Error; err != nil {
// 		return nil, err
// 	}

// 	if pagination.Limit == 0 {
// 		pagination.Limit = int(totalRows)
// 	}

// 	offset := pagination.Page * pagination.Limit

// 	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("Teacher").Find(&classes).Error; err != nil {
// 		return nil, err
// 	}

// 	var res = []response.ClassResponse{}
// 	for _, class := range classes {
// 		res = append(res, response.ClassResponse{
// 			Uuid: class.Uuid,
// 			Name: class.Name,
// 			Teacher: &response.TeacherData{
// 				Uuid:          class.Teacher.Uuid,
// 				JK:            class.Teacher.JK,
// 				Nip:           class.Teacher.Nip,
// 				Nuptk:         class.Teacher.Nuptk,
// 				StatusPegawai: class.Teacher.StatusPegawai,
// 				TempatLahir:   class.Teacher.TempatLahir,
// 				TanggalLahir:  class.Teacher.TanggalLahir,
// 				NoHp:          class.Teacher.NoHp,
// 				Jabatan:       class.Teacher.Jabatan,
// 				TotalJtm:      class.Teacher.TotalJtm,
// 				Alamat:        class.Teacher.Alamat,
// 				CreatedAt:     class.Teacher.CreatedAt,
// 				UpdatedAt:     class.Teacher.UpdatedAt,
// 			},
// 			CreatedAt: class.CreatedAt,
// 			UpdatedAt: class.UpdatedAt,
// 		})
// 	}

// 	pagination.Rows = res

// 	pagination.TotalRows = int(totalRows)

// 	return pagination, nil
// }

func (r *AssTypeRepoImplementation) FindAllAssType() (*[]domain.AssessmentType, error) {
	var result []domain.AssessmentType
	if err := r.db.Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AssTypeRepoImplementation) FindAssTypeByUuid(uuid string) (*domain.AssessmentType, error) {
	var result domain.AssessmentType
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AssTypeRepoImplementation) UpdateAssType(model *domain.AssessmentType) error {
	return r.db.Model(&model).Updates(&model).Error
}

func (r *AssTypeRepoImplementation) DeleteAssType(model *domain.AssessmentType) error {
	return r.db.Delete(&model).Error
}
