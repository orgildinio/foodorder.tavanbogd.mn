package controllers

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestUsers(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Users(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Users() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
