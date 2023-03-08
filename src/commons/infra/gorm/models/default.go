package commonsInfraGorm

import "time"

type Default struct {
	// Id        string `json:"id" gorm:"primary_key;type:string"`
	Created time.Time `json:"created" gorm:"autoUpdateTime"`
	Updated time.Time `json:"updated" gorm:"autoUpdateTime"`
}
