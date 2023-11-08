package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type Jobs struct {
	gorm.Model
	Company          Company            `json:"-" gorm:"ForeignKey:cid"`
	Cid              uint               `json:"cid"`
	MinNP            int                `json:"minNp"`
	MaxNP            int                `json:"maxNp"`
	Budget           int                `json:"budget"`
	Locations        []Locations        `gorm:"many2many:jobs_locations;"`
	TechnologyStacks []TechnologyStacks `gorm:"many2many:jobs_technologyStacks;"`
	WorkModes        []WorkModes        `gorm:"many2many:jobs_workMode;"`
	Description      string             `json:"description"`
	MinExp           int                `json:"minExp"`
	MaxExp           int                `json:"maxExp"`
	Qualifications   []Qualifications   `gorm:"many2many:jobs_qualifications;"`
	Shifts           []Shifts           `gorm:"many2many:jobs_shifts;"`
	JobTypes         []JobTypes         `gorm:"many2many:jobs_jobTypes;"`
}

type NewJobRequest struct {
	Cid              uint   `json:"cid"`
	MinNP            int    `json:"min_np"`
	MaxNP            int    `json:"max_np"`
	Budget           int    `json:"budget"`
	Locations        []uint `json:"locations"`
	TechnologyStacks []uint `json:"technologyStacks"`
	WorkModes        []uint `json:"workModes"`
	Description      string `json:"description"`
	MinExp           int    `json:"minExp"`
	MaxExp           int    `json:"maxExp"`
	Qualifications   []uint `json:"qualifications"`
	Shifts           []uint `json:"shifts"`
	JobTypes         []uint `json:"jobTypes"`
}

type RequestJob struct {
	Jid                uint64
	Cid                uint
	MinNP              int
	MaxNP              int
	Budget             int
	LocationsIDs       []uint
	TechnologyStackIDs []uint
	WorkModeIDs        []uint
	Description        string
	MinExp             int
	MaxExp             int
	Qualification      []uint
	Shift              []uint
	JobType            []uint
}

type NewJobResponse struct {
	ID uint
}

type Locations struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type TechnologyStacks struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type WorkModes struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Qualifications struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Shifts struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type JobTypes struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
