package main

import ()

// Value structure required by fitbit api alarm schema.
type alarm struct{
	DeviceId     string `json:"deviceId"`
	Id           string `json:"id"`
	Time         string `json:"time"`
	Enabled      bool `json:"enabled"`
	Recurring    bool `json:"recurring"`
	WeekDays     string `json:"weekdays"`
	Label        string `json:"label"`
	SnoozeLength int `json:"snoozeLength"`
	SnoozeCount  int `json:"snoozeCount"`
	Vibe         string `json:"vibe"`
}
