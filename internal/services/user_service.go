package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"github.com/iki-rumondor/go-monev/internal/utils"
	"gorm.io/gorm"
)

type UserService struct {
	Repo interfaces.UserRepoInterface
}

func NewUserService(repo interfaces.UserRepoInterface) interfaces.UserServiceInterface {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetUser(column string, value interface{}) (*models.User, error) {
	user, err := s.Repo.FindUserBy(column, value)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NOTFOUND_ERR("User Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return user, nil
}

func (s *UserService) GetAllUser() (*[]models.User, error) {
	user, err := s.Repo.FindUsers()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return user, nil
}

func (s *UserService) VerifyUser(req *request.SignIn) (string, error) {

	user, err := s.Repo.FindUserBy("username", req.Username)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", &response.Error{
				Code:    401,
				Message: "Username atau Password Salah",
			}
		}
		return "", response.SERVICE_INTERR
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return "", &response.Error{
			Code:    401,
			Message: "Username atau password salah",
		}
	}

	jwt, err := utils.GenerateToken(user.Uuid, user.Role.Name)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *UserService) GetDashboardAdmin() (map[string]interface{}, error) {
	subjects, err := s.Repo.FindSubjects()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	practicalSubjects, err := s.Repo.FindPracticalSubjects()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	teachers, err := s.Repo.GetAll("teachers")
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	years, err := s.Repo.GetAll("academic_years")
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	departments, err := s.Repo.GetAll("departments")
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	majors, err := s.Repo.GetAll("majors")
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	setting, err := s.Repo.GetOne("settings", "name", "step_monev")
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	resp := map[string]interface{}{
		"g_subject":   len(*subjects),
		"p_subject":   len(*practicalSubjects),
		"teacher":     len(teachers),
		"year":        len(years),
		"departments": len(departments),
		"majors":      len(majors),
		"step":        setting["value"],
	}

	return resp, nil
}

func (s *UserService) CountMonevByYear(userUuid, yearUuid string) (map[string]int, error) {

	user, err := s.GetUser("uuid", userUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.Repo.GetOne("academic_years", "uuid", yearUuid)
	if err != nil {
		return nil, err
	}

	yearID := year["id"].(uint64)

	result, err := s.Repo.CountMonevByYear(user.Department.ID, uint(yearID))
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *UserService) CountDepartmentMonev(departmentUuid, yearUuid string) (map[string]interface{}, error) {

	department, err := s.Repo.GetOne("departments", "uuid", departmentUuid)
	if err != nil {
		return nil, err
	}

	departmentID := department["id"].(uint64)

	year, err := s.Repo.GetOne("academic_years", "uuid", yearUuid)
	if err != nil {
		return nil, err
	}

	yearID := year["id"].(uint64)

	result, err := s.Repo.CountMonevByYear(uint(departmentID), uint(yearID))
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var subjects []models.Subject
	condition := fmt.Sprintf("department_id = '%d'", departmentID)
	if err := s.Repo.Find(&subjects, condition); err != nil {
		return nil, response.SERVICE_INTERR
	}

	var teacherSkill []models.TeacherSkill
	if err := s.Repo.FindTeacherSkills(&teacherSkill, uint(departmentID), uint(yearID)); err != nil {
		return nil, response.SERVICE_INTERR
	}

	var skills string
	for _, item := range teacherSkill {
		skills += item.Skill + " "
	}

	resp := map[string]interface{}{
		"first_monev": []int{
			result["plans"],
			result["modules"],
			result["tools"],
			result["skills"],
			result["facilities"],
		},
		"middle_monev": []int{
			result["t_att"],
			result["s_att"],
			result["mid_plans"],
		},
		"middle_last_monev": []int{
			result["lt_att"],
			result["ls_att"],
			result["last_plans"],
		},
		"last_monev": []int{
			result["final"],
			result["passed"],
			result["grade"],
		},
		"subjects":         len(subjects),
		"practical":        len(subjects),
		"plansAvailable":   result["plansAvailable"],
		"modulesAvailable": result["modulesAvailable"],
		"skills":           skills,
	}

	return resp, nil
}

func (s *UserService) Update(id uint, tableName, column string, value interface{}) error {

	if err := s.Repo.Update(id, tableName, column, value); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *UserService) GetAll(tableName string) ([]map[string]interface{}, error) {
	result, err := s.Repo.GetAll(tableName)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *UserService) CreateUser(req *request.CreateUser) error {
	role, err := s.Repo.GetOne("roles", "uuid", req.RoleUuid)
	if err != nil {
		return response.NOTFOUND_ERR("Role Tidak Ditemukan")
	}

	model := models.User{
		Username: req.Username,
		Password: req.Password,
		RoleID:   role["id"].(uint),
	}

	if err := s.Repo.CreateUser(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *UserService) GetDepartmentsChart(yearUuid string) (map[string]interface{}, error) {
	var model models.AcademicYear
	condition := fmt.Sprintf("uuid = '%s'", yearUuid)
	if err := s.Repo.First(&model, condition); err != nil {
		return nil, response.BADREQ_ERR("Tahun Ajaran Tidak Ditemukan")
	}

	var departments []models.Department
	if err := s.Repo.FindDepartments(&departments); err != nil {
		return nil, response.SERVICE_INTERR
	}

	var depNames []string
	var subjects []int
	var rps []int
	for _, item := range departments {
		var rpsAmount int
		depNames = append(depNames, item.Name)
		subjects = append(subjects, len(*item.Subjects))
		for _, item := range *item.Subjects {
			if item.AcademicPlan != nil && *item.AcademicPlan.Available && item.AcademicPlan.AcademicYearID == model.ID {
				rpsAmount++
			}
		}
		rps = append(rps, rpsAmount)
	}

	resp := map[string]interface{}{
		"departments": depNames,
		"subjects":    subjects,
		"rps":         rps,
	}
	return resp, nil
}

func (s *UserService) GetCurrentAcademicYear() (*response.AcademicYear, error) {
	var year models.AcademicYear
	if err := s.Repo.FirstWithOrder(&year, "", "year desc, semester desc"); err != nil {
		return nil, response.SERVICE_INTERR
	}

	first := utils.AddDate(year.FirstDate, year.FirstDays)
	middle := utils.AddDate(year.MiddleDate, year.MiddleDays)
	middle_last := utils.AddDate(year.MiddleLastDate, year.MiddleLastDays)
	last := utils.AddDate(year.LastDate, year.LastDays)

	resp := response.AcademicYear{
		Uuid:            year.Uuid,
		Semester:        year.Semester,
		Year:            year.Year,
		FirstDate:       year.FirstDate,
		MiddleDate:      year.MiddleDate,
		MiddleLastDate:  year.MiddleLastDate,
		LastDate:        year.LastDate,
		FirstDays:       year.FirstDays,
		MiddleDays:      year.MiddleDays,
		MiddleLastDays:  year.MiddleLastDays,
		LastDays:        year.LastDays,
		FirstRange:      first,
		MiddleRange:     middle,
		MiddleLastRange: middle_last,
		LastRange:       last,
		CreatedAt:       year.CreatedAt,
		UpdatedAt:       year.UpdatedAt,
	}

	return &resp, nil
}
