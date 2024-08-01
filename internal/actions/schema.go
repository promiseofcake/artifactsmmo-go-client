package actions

import (
	"github.com/promiseofcake/artifactsmmo-cli/client"
)

// Coords defines X, Y coords as signed integers
type Coords struct {
	X int
	Y int
}

// Response is a generic return value for most actions
type Response struct {
	CharacterResponse CharacterResponse
	CooldownSchema    client.CooldownSchema
	StatusCode        int
	StatusText        string
}

// GetRemainingCooldown returns the remaining cooldown based upon the last action
// this is useful for wait loops
func (g *Response) GetRemainingCooldown() int64 {
	return int64(g.CooldownSchema.RemainingSeconds)
}

// CharacterResponse is a specific wrapper around generic return info
type CharacterResponse struct {
	client.CharacterSchema
}

// GetPosition returns the given Coords for the Character in question
func (c *CharacterResponse) GetPosition() Coords {
	return Coords{
		X: c.X,
		Y: c.Y,
	}
}

// MyCharacterResponse is a specific return for your character info because OpenAPI isn't Dry.
type MyCharacterResponse struct {
	client.MyCharacterSchema
}

// GetPosition returns the given Coords for your given Character
func (c *MyCharacterResponse) GetPosition() Coords {
	return Coords{
		X: c.X,
		Y: c.Y,
	}
}

// FightResponse wraps a generic Response with Fight related data
type FightResponse struct {
	Response
	FightResponse client.FightSchema
}

// GatherResponse wraps a generic Response with Skill related data
type GatherResponse struct {
	Response
	SkillInfo client.SkillInfoSchema
}
