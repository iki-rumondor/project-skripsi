package application

import (
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type ProdiService struct {
	Repo *repository.Repositories
}

var (
	internalErr = &response.Error{
		Code:    500,
		Message: "Terjadi Kesalahan Sistem, Silahkan Hubungi Developper",
	}
)

func NewProdiService(repo *repository.Repositories) *ProdiService {
	return &ProdiService{
		Repo: repo,
	}
}

func (s *ProdiService) GetAllProdi() (*[]domain.Prodi, error) {

	jurusan, err := s.Repo.ProdiRepo.FindAllProdi()
	if err != nil {
		return nil, err
	}

	return jurusan, nil
}

func (s *ProdiService) GetProdiByID(id uint) (*domain.Prodi, error) {

	jurusan, err := s.Repo.ProdiRepo.FindProdi(id)
	if err != nil {
		return nil, err
	}

	return jurusan, nil
}

func (s *ProdiService) GetProdiByUuid(uuid string) (*domain.Prodi, error) {

	result, err := s.Repo.ProdiRepo.FindProdiByUuid(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: "Program Studi tidak ditemukan",
			}
		}
		return nil, internalErr
	}

	return result, nil
}

func (s *ProdiService) CreateProdi(prodi *domain.Prodi) error {

	if err := s.Repo.ProdiRepo.CreateProdi(prodi); err != nil {
		return err
	}

	return nil
}

func (s *ProdiService) DeleteProdi(uuid string) error {
	prodi, err := s.GetProdiByUuid(uuid)
	if err != nil {
		return err
	}

	if err := s.Repo.ProdiRepo.DeleteProdi(prodi); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    403,
				Message: "Data ini tidak dapat dihapus karena berelasi dengan data lain",
			}
		}

		return internalErr
	}

	return nil
}

func (s *ProdiService) UpdateProdi(uuid string, body *request.Prodi) error {

	prodi, err := s.GetProdiByUuid(uuid)
	if err != nil {
		return err
	}

	jurusanID, err := strconv.Atoi(body.JurusanID)
	if err != nil {
		return &response.Error{
			Code:    404,
			Message: "Data jurusan tidak valid",
		}
	}

	model := domain.Prodi{
		ID:        prodi.ID,
		Nama:      body.Nama,
		Kaprodi:   body.Kaprodi,
		JurusanID: uint(jurusanID),
	}

	if err := s.Repo.ProdiRepo.UpdateProdi(&model); err != nil {
		return internalErr
	}

	return nil
}

func (s *ProdiService) CreateSubject(userUuid string, body *request.Subject) error {
	prodi, err := s.GetProdiByUuid(userUuid)
	if err != nil {
		return err
	}

	model := domain.Subject{
		Uuid:    uuid.NewString(),
		Code:    body.Code,
		Name:    body.Name,
		ProdiID: prodi.ID,
	}

	if err := s.Repo.SubjectRepo.CreateSubject(&model); err != nil {
		return internalErr
	}

	return nil
}

func (s *ProdiService) GetAllSubject() (*[]domain.Subject, error) {

	result, err := s.Repo.SubjectRepo.FindAllSubject()
	if err != nil {
		return nil, internalErr
	}

	return result, nil
}

func (s *ProdiService) GetProdiSubjects(uuid string) (*[]domain.Subject, error) {
	prodi, err := s.GetProdiByUuid(uuid)
	if err != nil{
		return nil, err
	}

	result, err := s.Repo.SubjectRepo.FindSubjectsWhere("prodi_id", prodi.ID)
	if err != nil {
		return nil, internalErr
	}

	return result, nil
}

func (s *ProdiService) GetSubjectByUuid(uuid string) (*domain.Subject, error) {

	model, err := s.Repo.SubjectRepo.FindSubjectByUuid(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: "Mata Kuliah tidak ditemukan",
			}
		}
		return nil, internalErr
	}

	return model, nil
}

func (s *ProdiService) UpdateSubject(userUuid, subjectUuid string, body *request.Subject) error {

	subject, err := s.GetSubjectByUuid(subjectUuid)
	if err != nil {
		return err
	}

	if subject.Prodi.Uuid != userUuid{
		return &response.Error{
			Code: 403,
			Message: "Mata Kuliah Tidak Terdaftar Di Program Studi ",
		}
	}

	prodi, err := s.GetProdiByUuid(userUuid)
	if err != nil {
		return err
	}

	model := domain.Subject{
		ID:      subject.ID,
		Name:    body.Name,
		Code:    body.Code,
		ProdiID: prodi.ID,
	}

	if err := s.Repo.SubjectRepo.UpdateSubject(&model); err != nil {
		return internalErr
	}

	return nil
}

func (s *ProdiService) DeleteSubject(userUuid, subjectUuid string) error {
	model, err := s.GetSubjectByUuid(subjectUuid)
	if err != nil {
		return err
	}

	if model.Prodi.Uuid != userUuid{
		return &response.Error{
			Code: 403,
			Message: "Mata Kuliah Tidak Terdaftar Di Program Studi ",
		}
	}

	if err := s.Repo.SubjectRepo.DeleteSubject(model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    403,
				Message: "Data ini tidak dapat dihapus karena berelasi dengan data lain",
			}
		}

		return internalErr
	}

	return nil
}
