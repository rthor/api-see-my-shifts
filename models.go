package main

import (
    "time"
)

type TeamPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Team struct {
	ID   string   `json:"id"`
	Plan TeamPlan `json:"plan"`
	Name string   `json:"name"`
}

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ShortName  string `json:"shortName"`
	FacebookID string `json:"facebookId"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type UserTeam struct {
	ID   string `json:"id"`
	Team Team   `json:"team"`
	User User   `json:"user"`
}

type Shift struct {
	ID        string    `json:"id"`
	User      User      `json:"user"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Fulfilled bool      `json:"isFulfilled"`
}

type Day struct {
	ID     string    `json:"id"`
	Date   time.Time `json:"date"`
	Shifts []Shift   `json:"shifts"`
}
