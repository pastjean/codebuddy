package store

import "github.com/pastjean/codebuddy/models"

// Store store is the representation of a type of repository which is abstracting
// the way we retrieve data
type Store interface {
	GetTeams() (models.TeamsMap, error)
	GetTeamMembers(team string) (models.TeamMembers, error)
	UpdateTeam(team string, members models.TeamMembers) error
	Provision() error
}
