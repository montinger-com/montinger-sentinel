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
	CPU    *CPU    `json:"cpu,omitempty" bson:"cpu,omitempty"`
	Memory *Memory `json:"memory,omitempty" bson:"memory,omitempty"`
	OS     *OS     `json:"os,omitempty" bson:"os,omitempty"`
	Uptime uint64  `json:"uptime,omitempty" bson:"uptime,omitempty"`
}

type CPU struct {
	UsedPercent float64 `json:"used_percent,omitempty" bson:"used_percent,omitempty"`
}

type Memory struct {
	Total       uint64  `json:"total,omitempty" bson:"total,omitempty"`
	Available   uint64  `json:"available,omitempty" bson:"available,omitempty"`
	Used        uint64  `json:"used,omitempty" bson:"used,omitempty"`
	UsedPercent float64 `json:"used_percent,omitempty" bson:"used_percent,omitempty"`
}

type OS struct {
	Type            string `json:"type,omitempty" bson:"type,omitempty"`
	Platform        string `json:"platform,omitempty" bson:"platform,omitempty"`
	PlatformFamily  string `json:"platform_family,omitempty" bson:"platform_family,omitempty"`
	PlatformVersion string `json:"platform_version,omitempty" bson:"platform_version,omitempty"`
	KernelVersion   string `json:"kernel_version,omitempty" bson:"kernel_version,omitempty"`
	KernelArch      string `json:"kernel_arch,omitempty" bson:"kernel_arch,omitempty"`
}
