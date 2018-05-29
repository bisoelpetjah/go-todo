package models

import (
    "time"

    "github.com/guregu/null"
)

type Todo struct {
    Id string `json:"id"`
    Message null.String `json:"message"`
    IsDone null.Bool `json:"is_done"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
