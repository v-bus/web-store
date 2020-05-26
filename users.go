package main

import (
	"fmt"
	dbproduct "web-store/db"

	log "github.com/sirupsen/logrus"
)

//getUsersByRole returns user string list by role
func getUsersByRole(role string) []string {
	var u = []dbproduct.Users{}
	var res []string
	if db.HasTable(dbproduct.Users{}) {
		if err := db.Find(&u, fmt.Sprintf("role = '%s'", role)).Error; err != nil {
			log.Warn(err)
		} else {
			for _, usr := range u {
				res = append(res, usr.User)
			}
		}
	} else {
		log.Warn("No users found with role ", role, "in users")
	}
	return res
}
