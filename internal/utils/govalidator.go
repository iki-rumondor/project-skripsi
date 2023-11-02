package utils

import (
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

func NewCustomValidator(gormDB *gorm.DB) {

	// Get data from table param[0] where param[1] = str

	govalidator.ParamTagMap["unique"] = govalidator.ParamValidator(func(str string, params ...string) bool {

		var result int
		query := `SELECT COUNT(*) as total FROM %s WHERE %s = ?`
		gormDB.Raw(fmt.Sprintf(query, params[0], params[1]), str).Scan(&result)

		return result == 0
	})

	govalidator.ParamTagRegexMap["unique"] = regexp.MustCompile("^unique\\((\\w+)\\.(\\w+)\\)$")
}
