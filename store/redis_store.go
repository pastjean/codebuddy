package store

import (
	"encoding/json"

	"github.com/pastjean/codebuddy/models"
	"gopkg.in/redis.v3"
)

type redisStore struct {
	uri    string
	prefix string
	client *redis.Client
}

// NewRedisStore returns a new configured store which store information in a
// redis server
func NewRedisStore(uri string) Store {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &redisStore{
		client: client,
		prefix: "buddy-",
	}
}

func (s *redisStore) Provision() error {
	teamsKey := s.prefix + "teams"
	return s.client.Del(teamsKey).Err()
}

func (s *redisStore) GetTeams() (models.TeamsMap, error) {
	teamsCmdResult := s.client.HGetAll(s.prefix + "teams")

	if err := teamsCmdResult.Err(); err != nil {
		return nil, err
	}

	teamsMap := models.TeamsMap{}
	key := ""
	for i, val := range teamsCmdResult.Val() {
		if i%2 == 0 {
			key = val
		} else {
			members := models.TeamMembers{}
			err := json.Unmarshal([]byte(val), &members)
			if err != nil {
				return nil, err
			}
			teamsMap[key] = members
		}
	}
	return teamsMap, nil
}

func (s *redisStore) GetTeamMembers(team string) (models.TeamMembers, error) {
	members := models.TeamMembers{}

	marshaledMembers, err := s.client.HGet(s.prefix+"teams", team).Result()
	if err != nil {
		return members, err
	}

	if err := json.Unmarshal([]byte(marshaledMembers), &members); err != nil {
		return members, err
	}
	return members, nil
}

func (s *redisStore) UpdateTeam(team string, members models.TeamMembers) error {
	marshaledMembers, err := json.Marshal(members)
	if err != nil {
		return err
	}

	return s.client.HSet(s.prefix+"teams", team, string(marshaledMembers)).Err()
}
