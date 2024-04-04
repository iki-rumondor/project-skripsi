package services

import (
	"errors"
	"log"
	"strconv"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"github.com/iki-rumondor/go-monev/internal/utils"
	"gorm.io/gorm"
)

type MiddleMonevService struct {
	Repo interfaces.MiddleMonevRepoInterface
}

func NewMiddleMonevService(repo interfaces.MiddleMonevRepoInterface) interfaces.MiddleMonevServiceInterface {
	return &MiddleMonevService{
		Repo: repo,
	}
}

func (s *MiddleMonevService) CreateTeacherAttendence(userUuid string, req *request.CreateTeacherAttendence) error {
	subject, err := s.Repo.FindUserSubject(userUuid, req.SubjectUuid)
	if err != nil {
		log.Println(err.Error())
		return response.NOTFOUND_ERR("Mata Kuliah Tidak Ditemukan")
	}

	academicYear, err := s.Repo.FindAcademicYear(req.AcademicYearUuid)
	if err != nil {
		log.Println(err.Error())
		return response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
	}

	teacher, err := s.Repo.FindTeacher(req.TeacherUuid)
	if err != nil {
		log.Println(err.Error())
		return response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
	}

	middle, _ := strconv.Atoi(req.Middle)

	model := models.TeacherAttendence{
		SubjectID:      subject.ID,
		AcademicYearID: academicYear.ID,
		TeacherID:      teacher.ID,
		Middle:         uint(middle),
		Class:          req.Class,
	}

	if err := s.Repo.CreateTeacherAttendence(&model); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MiddleMonevService) CreateStudentAttendence(userUuid string, req *request.CreateStudentAttendence) error {
	subject, err := s.Repo.FindUserSubject(userUuid, req.SubjectUuid)
	if err != nil {
		log.Println(err.Error())
		return response.NOTFOUND_ERR("Mata Kuliah Tidak Ditemukan")
	}

	academicYear, err := s.Repo.FindAcademicYear(req.AcademicYearUuid)
	if err != nil {
		log.Println(err.Error())
		return response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
	}

	middle, _ := strconv.Atoi(req.Middle)
	students, _ := strconv.Atoi(req.StudentAmount)

	if middle > students {
		return response.BADREQ_ERR("Jumlah Kehadiran Mahasiswa Tidak Boleh Lebih Dari Jumlah Mahasiswa")
	}

	model := models.StudentAttendence{
		SubjectID:      subject.ID,
		AcademicYearID: academicYear.ID,
		Middle:         uint(middle),
		StudentAmount:  uint(students),
		Class:          req.Class,
	}

	if err := s.Repo.CreateStudentAttendence(&model); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MiddleMonevService) GetTeacherAttendences(userUuid, yearUuid string) (*[]models.TeacherAttendence, error) {
	year, err := s.Repo.FindAcademicYear(yearUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
	}

	result, err := s.Repo.FindTeacherAttendences(userUuid, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MiddleMonevService) GetStudentAttendences(userUuid, yearUuid string) (*[]models.StudentAttendence, error) {
	year, err := s.Repo.FindAcademicYear(yearUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
	}

	result, err := s.Repo.FindStudentAttendences(userUuid, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MiddleMonevService) GetTeacherAttendence(userUuid, uuid string) (*models.TeacherAttendence, error) {
	result, err := s.Repo.FindTeacherAttendence(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Absensi Dosen Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MiddleMonevService) GetStudentAttendence(userUuid, uuid string) (*models.StudentAttendence, error) {
	result, err := s.Repo.FindStudentAttendence(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Absensi Dosen Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MiddleMonevService) CountTotalMonev(userUuid, yearUuid string) (map[string]int, error) {

	year, err := s.Repo.FindAcademicYear(yearUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
	}

	result, err := s.Repo.CountTotalMonev(userUuid, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MiddleMonevService) UpdateTeacherAttendence(userUuid, uuid string, model *models.TeacherAttendence) error {

	result, err := s.GetTeacherAttendence(userUuid, uuid)
	if err != nil {
		return err
	}

	model.ID = result.ID

	if err := s.Repo.UpdateTeacherAttendence(model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MiddleMonevService) UpdateStudentAttendence(userUuid, uuid string, model *models.StudentAttendence) error {

	result, err := s.GetStudentAttendence(userUuid, uuid)
	if err != nil {
		return err
	}

	if result.StudentAmount < model.Last {
		return response.BADREQ_ERR("Jumlah Mahasiswa Hadir Tidak Boleh Lebih Dari Jumlah Seluruh Mahasiswa")
	}

	model.ID = result.ID

	if err := s.Repo.UpdateStudentAttendence(model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MiddleMonevService) DeleteTeacherAttendence(userUuid, uuid string) error {

	result, err := s.GetTeacherAttendence(userUuid, uuid)
	if err != nil {
		return err
	}

	if err := s.Repo.DeleteTeacherAttendence(result); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MiddleMonevService) DeleteStudentAttendence(userUuid, uuid string) error {

	result, err := s.GetStudentAttendence(userUuid, uuid)
	if err != nil {
		return err
	}

	if err := s.Repo.DeleteStudentAttendence(result); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MiddleMonevService) GetTeacherAttendencesByDepartment(departmentUuid, yearUuid string) (*[]models.TeacherAttendence, error) {

	department, err := s.Repo.FindDepartment(departmentUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	year, err := s.Repo.FindAcademicYear(yearUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
	}

	result, err := s.Repo.FindTeacherAttendences(department.User.Uuid, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MiddleMonevService) GetStudentAttendencesByDepartment(departmentUuid, yearUuid string) (*[]models.StudentAttendence, error) {

	department, err := s.Repo.FindDepartment(departmentUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	year, err := s.Repo.FindAcademicYear(yearUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
	}

	result, err := s.Repo.FindStudentAttendences(department.User.Uuid, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}
