// @Title pre_detect_database.go
// @Description
// @Author Hunter 2024/9/4 11:26

package database

import (
	"fmt"
)

func PreDetectDatabase() error {
	// if err := preDetectPostgres(); err != nil {
	// 	return err
	// }

	if err := preDetectRedis(); err != nil {
		return err
	}

	fmt.Println("All database connections are successful.")
	return nil
}
