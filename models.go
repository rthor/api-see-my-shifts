package main

import (
    "time"
	"github.com/jinzhu/gorm"
)

type TeamPlan struct {
    gorm.Model
	Name string `json:"name"`
}

type Team struct {
    gorm.Model
    Plan TeamPlan `gorm:"foreignKeyPlanID" json:"plan"`
	Name string   `json:"name"`
}

type User struct {
    gorm.Model
	Name       string `json:"name"`
	ShortName  string `json:"shortName"`
	FacebookID string `json:"facebookId"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type UserTeam struct {
    gorm.Model
	Team Team   `json:"team"`
	User User   `json:"user"`
}

type Shift struct {
    gorm.Model
	User      User      `json:"user"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Fulfilled bool      `json:"isFulfilled"`
}

type Day struct {
    gorm.Model
	Date   time.Time `json:"date"`
	Shifts []Shift   `json:"shifts"`
}
