package client

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *CooldownSchema) UnmarshalJSON(data []byte) error {
	// Alias to avoid infinite recursion
	type Alias CooldownSchema
	aux := &struct {
		Expiration string `json:"expiration"`
		StartedAt  string `json:"started_at"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	var err error
	if c.Expiration, err = time.Parse(time.RFC3339Nano, aux.Expiration); err != nil {
		return fmt.Errorf("error parsing expiration: %w", err)
	}

	if c.StartedAt, err = time.Parse(time.RFC3339Nano, aux.StartedAt); err != nil {
		return fmt.Errorf("error parsing started_at: %w", err)
	}

	return nil
}
