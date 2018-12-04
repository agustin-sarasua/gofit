package main

import (
	"errors"
	"time"

	"github.com/agustin-sarasua/gofit/model"
)

func validateCreateClass(c *model.Class) error {
	t, err := time.Parse(model.TimeLayout, c.StartTime)
	if err != nil {
		return err
	}
	if time.Now().After(t) {
		return errors.New("StartTime < Now")
	}
	return nil
}
