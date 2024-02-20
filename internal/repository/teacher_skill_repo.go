package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type TeacherSkillRepository struct {
	db *gorm.DB
}

func NewTeacherSkillRepository(db *gorm.DB) interfaces.TeacherSkillRepoInterface {
	return &TeacherSkillRepository{
		db: db,
	}
}

func (r *TeacherSkillRepository) FindTeacherSkills(userUuid string) (*[]models.TeacherSkill, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.TeacherSkill
	if err := r.db.Joins("Teacher").Preload("Subject").Find(&result, "teacher.department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherSkillRepository) FindUserTeacherSkill(userUuid, uuid string) (*models.TeacherSkill, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.TeacherSkill
	if err := r.db.Joins("Teacher").Preload("Subject").First(&result, "teacher_skills.uuid = ? AND teacher.department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherSkillRepository) FindTeacherSkillsByYear(userUuid string, yearID uint) (*[]models.TeacherSkill, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.TeacherSkill
	if err := r.db.Joins("Teacher").Preload("Subject").Find(&result, "teacher.department_id = ? AND academic_year_id = ?", user.Department.ID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherSkillRepository) FindTeacherBy(column string, value interface{}) (*models.Teacher, error) {
	var result models.Teacher
	if err := r.db.Preload("Department.User").First(&result, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherSkillRepository) FindSubjectByUuid(uuid string) (*models.Subject, error) {
	var result models.Subject
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherSkillRepository) FindAcademicYearByUuid(uuid string) (*models.AcademicYear, error) {
	var result models.AcademicYear
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherSkillRepository) CreateTeacherSkill(model *models.TeacherSkill) error {
	return r.db.Create(model).Error
}

func (r *TeacherSkillRepository) UpdateTeacherSkill(model *models.TeacherSkill) error {
	return r.db.Updates(model).Error
}

func (r *TeacherSkillRepository) DeleteTeacherSkill(model *models.TeacherSkill) error {
	return r.db.Delete(model).Error
}
