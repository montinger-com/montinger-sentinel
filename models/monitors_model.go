package models

import "time"

type Monitor struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Type   string `json:"type" bson:"type"`
	Status string `json:"status" bson:"status"`

	APIKey string `json:"api_key,omitempty" bson:"api_key,omitempty"`

	LastDataOn *time.Time `json:"last_data_on,omitempty" bson:"last_data_on,omitempty"`
	LastData   LastData   `json:"last_data,omitempty" bson:"last_data,omitempty"`

	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type LastData struct {
	CPUUsage    float64 `json:"cpu_usage,omitempty" bson:"cpu_usage,omitempty"`
	MemoryUsage float64 `json:"memory_usage,omitempty" bson:"memory_usage,omitempty"`
}
