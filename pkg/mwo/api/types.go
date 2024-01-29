package api

import (
	"time"
)

type MatchResponse struct {
	MatchDetails MatchDetails  `json:"MatchDetails"`
	UserDetails  []UserDetails `json:"UserDetails"`
}

type MatchDetails struct {
	Map                string    `json:"Map"`
	ViewMode           string    `json:"ViewMode"`
	TimeOfDay          string    `json:"TimeOfDay"`
	GameMode           string    `json:"GameMode"`
	Region             string    `json:"Region"`
	MatchTimeMinutes   int       `json:"MatchTimeMinutes,string"`
	UseStockLoadout    bool      `json:"UseStockLoadout"`
	NoMechQuirks       bool      `json:"NoMechQuirks"`
	NoMechEfficiencies bool      `json:"NoMechEfficiencies"`
	WinningTeam        string    `json:"WinningTeam"`
	Team1Score         int       `json:"Team1Score"`
	Team2Score         int       `json:"Team2Score"`
	MatchDuration      int       `json:"MatchDuration,string"`
	CompleteTime       time.Time `json:"CompleteTime"`
}

type UserDetails struct {
	Username            string `json:"Username"`
	IsSpectator         bool   `json:"IsSpectator"`
	Team                string `json:"Team"`
	Lance               string `json:"Lance"`
	MechItemID          int    `json:"MechItemID"`
	MechName            string `json:"MechName"`
	SkillTier           int    `json:"SkillTier"`
	HealthPercentage    int    `json:"HealthPercentage"`
	Kills               int    `json:"Kills"`
	KillsMostDamage     int    `json:"KillsMostDamage"`
	Assists             int    `json:"Assists"`
	ComponentsDestroyed int    `json:"ComponentsDestroyed"`
	MatchScore          int    `json:"MatchScore"`
	Damage              int    `json:"Damage"`
	TeamDamage          int    `json:"TeamDamage"`
	UnitTag             string `json:"UnitTag"`
}
