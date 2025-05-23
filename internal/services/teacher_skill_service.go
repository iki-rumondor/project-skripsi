package services

import (
	"errors"
	"log"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type TeacherSkillService struct {
	Repo interfaces.TeacherSkillRepoInterface
}

func NewTeacherSkillService(repo interfaces.TeacherSkillRepoInterface) interfaces.TeacherSkillServiceInterface {
	return &TeacherSkillService{
		Repo: repo,
	}
}

func (s *TeacherSkillService) CreateTeacherSkill(userUuid string, req *request.TeacherSkill) error {

	teacher, err := s.Repo.FindTeacherBy("uuid", req.TeacherUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Dosen Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if teacher.Department.User.Uuid != userUuid {
		return response.NOTFOUND_ERR("Kemampuan Dosen Tidak Ditemukan")
	}

	subject, err := s.Repo.FindSubjectByUuid(req.SubjectUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Mata Kuliah tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	academicYear, err := s.Repo.FindAcademicYearByUuid(req.AcademicYearUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Tahun Ajaran tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	model := models.TeacherSkill{
		Skill:          req.Skill,
		TeacherID:      teacher.ID,
		SubjectID:      subject.ID,
		AcademicYearID: academicYear.ID,
	}

	if err := s.Repo.CreateTeacherSkill(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *TeacherSkillService) GetAllTeacherSkills(userUuid string) (*[]models.TeacherSkill, error) {

	result, err := s.Repo.FindTeacherSkills(userUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *TeacherSkillService) GetTeacherSkillsByYear(userUuid, yearUuid string) (*[]models.TeacherSkill, error) {

	year, err := s.Repo.FindAcademicYearByUuid(yearUuid)
	if err != nil{
		return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
	}
	result, err := s.Repo.FindTeacherSkillsByYear(userUuid, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *TeacherSkillService) GetTeacherSkill(userUuid, uuid string) (*models.TeacherSkill, error) {
	result, err := s.Repo.FindUserTeacherSkill(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Kemampuan Dosen Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *TeacherSkillService) UpdateTeacherSkill(userUuid, uuid string, req *request.UpdateTeacherSkill) error {

	result, err := s.GetTeacherSkill(userUuid, uuid)
	if err != nil {
		return err
	}

	teacher, err := s.Repo.FindTeacherBy("uuid", req.TeacherUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Dosen Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	subject, err := s.Repo.FindSubjectByUuid(req.SubjectUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Mata Kuliah tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	model := models.TeacherSkill{
		ID:        result.ID,
		Skill:     req.Skill,
		TeacherID: teacher.ID,
		SubjectID: subject.ID,
	}

	if err := s.Repo.UpdateTeacherSkill(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TeacherSkillService) DeleteTeacherSkill(userUuid, uuid string) error {
	result, err := s.GetTeacherSkill(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.TeacherSkill{
		ID: result.ID,
	}

	if err := s.Repo.DeleteTeacherSkill(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TeacherSkillService) GetByDepartment(departmentUuid, yearUuid string) (*[]models.TeacherSkill, error) {

	department, err := s.Repo.FindDepartment(departmentUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	year, err := s.Repo.FindAcademicYearBy("uuid", yearUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	result, err := s.Repo.FindByDepartment(department.ID, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}
	
	return result, nil
}