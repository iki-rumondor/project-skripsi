package testing

import (
	"errors"
	"testing"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
	"github.com/iki-rumondor/go-monev/internal/services"
	"github.com/iki-rumondor/go-monev/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateAcademicPlan_Success(t *testing.T) {
	mockRepo := &mocks.AcademicPlanRepoInterface{}
	mockRepo.On("FindSubjectBy", mock.Anything, mock.Anything).Return(&models.Subject{}, nil)
	mockRepo.On("FindAcademicYearBy", mock.Anything, mock.Anything).Return(&models.AcademicYear{}, nil)
	mockRepo.On("CreateAcademicPlan", mock.Anything).Return(nil)

	academicPlanService := services.NewAcademicPlanService(mockRepo)

	available := true

	req := request.AcademicPlan{
		SubjectUuid:      "subject_uuid",
		AcademicYearUuid: "academic_year_uuid",
		Available:        &available,
		Note:             "test note",
	}

	err := academicPlanService.CreateAcademicPlan("user_uuid", &req)
	assert.NoError(t, err)
}
func TestCreateAcademicPlan_Failed(t *testing.T) {
	mockRepo := &mocks.AcademicPlanRepoInterface{}
	mockRepo.On("FindSubjectBy", mock.Anything, mock.Anything).Return(&models.Subject{}, nil)
	mockRepo.On("FindAcademicYearBy", mock.Anything, mock.Anything).Return(&models.AcademicYear{}, nil)
	mockRepo.On("CreateAcademicPlan", mock.Anything).Return(errors.New("something error"))

	academicPlanService := services.NewAcademicPlanService(mockRepo)

	available := true

	req := request.AcademicPlan{
		SubjectUuid:      "subject_uuid",
		AcademicYearUuid: "academic_year_uuid",
		Available:        &available,
		Note:             "test note",
	}

	err := academicPlanService.CreateAcademicPlan("user_uuid", &req)
	assert.Error(t, err)
}

func TestCreateAcademicPlan_SubjectNotFound(t *testing.T) {
	mockRepo := &mocks.AcademicPlanRepoInterface{}
	mockRepo.On("FindSubjectBy", mock.Anything, mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	academicPlanService := services.NewAcademicPlanService(mockRepo)

	available := true

	req := &request.AcademicPlan{
		SubjectUuid:      "subject_uuid",
		AcademicYearUuid: "academic_year_uuid",
		Available:        &available,
		Note:             "test note",
	}

	err := academicPlanService.CreateAcademicPlan("user_uuid", req)

	assert.Error(t, err)
	assert.Equal(t, "404: Mata Kuliah Tidak Ditemukan", err.Error())
}
func TestCreateAcademicPlan_AcademicYearNotFound(t *testing.T) {
	mockRepo := &mocks.AcademicPlanRepoInterface{}
	mockRepo.On("FindSubjectBy", mock.Anything, mock.Anything).Return(&models.Subject{}, nil)
	mockRepo.On("FindAcademicYearBy", mock.Anything, mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	academicPlanService := services.NewAcademicPlanService(mockRepo)

	available := true

	req := &request.AcademicPlan{
		SubjectUuid:      "subject_uuid",
		AcademicYearUuid: "academic_year_uuid",
		Available:        &available,
		Note:             "test note",
	}

	err := academicPlanService.CreateAcademicPlan("user_uuid", req)

	assert.Error(t, err)
	assert.Equal(t, "404: Tahun Ajaran Tidak Ditemukan", err.Error())
}
