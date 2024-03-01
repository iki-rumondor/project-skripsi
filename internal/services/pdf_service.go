package services

import (
	"errors"
	"io"
	"log"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type PdfService struct {
	Repo interfaces.PdfRepoInterface
}

func NewPdfService(repo interfaces.PdfRepoInterface) interfaces.PdfServiceInterface {
	return &PdfService{
		Repo: repo,
	}
}

func (s *PdfService) GetDepartment(uuid string) (*models.Department, error) {
	department, err := s.Repo.FindDepartmentBy("uuid", uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Program Studi Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return department, nil
}

func (s *PdfService) GetAcademicYear(uuid string) (*models.AcademicYear, error) {
	result, err := s.Repo.FindAcademicYearBy("uuid", uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Lab Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *PdfService) RpsReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	plans, err := s.Repo.FindAcademicPlans(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var pdfPlans []request.PdfPlans
	for _, item := range *plans {
		status := "Tidak Tersedia"
		if *item.Available {
			status = "Tersedia"
		}

		pdfPlans = append(pdfPlans, request.PdfPlans{
			Subject: item.Subject.Name,
			Status:  status,
			Note:    item.Note,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       pdfPlans,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "plans")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) ModuleReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	modules, err := s.Repo.FindModules(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *modules {
		status := "Tidak Tersedia"
		if *item.Available {
			status = "Tersedia"
		}

		data = append(data, map[string]interface{}{
			"subject": item.Subject.Name,
			"status":  status,
			"note":    item.Note,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "modules")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) ToolReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	tools, err := s.Repo.FindTools(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *tools {
		status := "Tidak Tersedia"
		if *item.Available {
			status = "Tersedia"
		}

		data = append(data, map[string]interface{}{
			"subject":   item.Subject.Name,
			"status":    status,
			"condition": item.Condition,
			"note":      item.Note,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "tools")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) SkillReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	skills, err := s.Repo.FindSkills(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *skills {

		data = append(data, map[string]interface{}{
			"teacher": item.Teacher.Name,
			"subject": item.Subject.Name,
			"skill":   item.Skill,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "skills")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) FacilitiesReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	facilities, err := s.Repo.FindFacilities(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *facilities {

		data = append(data, map[string]interface{}{
			"facility": item.Facility.Name,
			"amount":   item.Amount,
			"unit":     item.Unit,
			"note":     item.Note,
			"deactive": item.Deactive,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "facilities")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) StudentAttendencesReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.FindStudentAttendences(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *result {
		data = append(data, map[string]interface{}{
			"subject": item.Subject.Name,
			"amount":  item.StudentAmount,
			"middle":  item.Middle,
			"last":    item.Last,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "student-attendences")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) TeacherAttendencesReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.FindTeacherAttendences(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *result {
		data = append(data, map[string]interface{}{
			"subject": item.Subject.Name,
			"middle":  item.Middle,
			"last":    item.Last,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "teacher-attendences")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) PlanMatchReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.FindAcademicPlans(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *result {
		data = append(data, map[string]interface{}{
			"subject": item.Subject.Name,
			"middle":  item.Middle,
			"last":    item.Last,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "plans-matching")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) FinalStudentReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.FindStudentAttendences(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *result {
		data = append(data, map[string]interface{}{
			"subject": item.Subject.Name,
			"amount":  item.StudentAmount,
			"final":   item.FinalAmount,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "final")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) PassedStudentReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.FindStudentAttendences(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *result {
		data = append(data, map[string]interface{}{
			"subject": item.Subject.Name,
			"amount":  item.StudentAmount,
			"passed":  item.PassedAmount,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "passed")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}

func (s *PdfService) GradeReport(departmentUuid, yearUuid string) ([]byte, error) {
	department, err := s.GetDepartment(departmentUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.FindTeacherAttendences(department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var data []map[string]interface{}
	for _, item := range *result {
		data = append(data, map[string]interface{}{
			"subject": item.Subject.Name,
			"teacher": item.Teacher.Name,
			"grade":   item.GradeOnTime,
		})
	}

	reportPdf := request.PdfReport{
		Department: department.Name,
		Semester:   year.Semester,
		Year:       year.Year,
		Data:       data,
	}

	resp, err := s.Repo.LaravelPdf(reportPdf, "grade")
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return pdfData, nil
}
