package models

// TeamMembers represent a list of usernames which are part of a team
type TeamMembers []string

// TeamsMap is a map with the key of the teamName and value of the team users
type TeamsMap map[string]TeamMembers
