package models

import (
	"gorm.io/datatypes"
	"time"
)


type Asset struct {
	UUID                    string         `json:"uuid,omitempty"`
	CustomerSiteStorageUUID string         `json:"customer_site_storage_uuid,omitempty"`
	Location                string         `json:"location,omitempty"`
	Path                    string         `json:"path,omitempty"`
	Filename                string         `json:"filename,omitempty"`
	FileTipology            string         `json:"file_tipology,omitempty"`
	MasterFileUUID          string         `json:"master_file_uuid,omitempty"`
	IsOriginal              *int           `json:"is_original,omitempty"`
	Probe                   datatypes.JSON `json:"probe,omitempty"`
	Info                    datatypes.JSON `json:"info,omitempty"`
	QC                      datatypes.JSON `json:"qc,omitempty"`
	SD                      datatypes.JSON `json:"sd,omitempty"`
	Sprites                 datatypes.JSON `json:"sprites,omitempty"`
	Fingerprint             string         `json:"fingerprint,omitempty"`
	Size                    *int64         `json:"size,omitempty"`
	Published               *int           `json:"published,omitempty"`
	FileCreatedAt           *time.Time     `json:"file_created_at,omitempty"`
	FileUpdatedAt           *time.Time     `json:"file_updated_at,omitempty"`
	JobID                   *int           `json:"job_id,omitempty"`
	FileURL                 string         `json:"file_url,omitempty"`
	Title                   string         `json:"title,omitempty"`
	JobIDCreatedAt          *time.Time     `json:"job_id_created_at,omitempty"`
	State                   *int           `json:"state,omitempty"`
	Folder                  string         `json:"folder,omitempty"`
	Token                   string         `json:"token,omitempty"`
	TotalSize               *int64         `json:"total_size,omitempty"`
	AssetUUID               string         `json:"asset_uuid,omitempty"`
	CategoryID              *int           `json:"category_id,omitempty"`
	URLDomain               string         `json:"url_domain,omitempty"`
	IsMultimedia            *int           `json:"is_multimedia,omitempty"`
	Entity
}

type Entity struct {
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}


type AssetTag struct {
	UUID      string `json:"uuid,omitempty"`
	AssetUUID string `json:"asset_uuid,omitempty"`
	TagName   string `json:"tag_name,omitempty"`
	TagValue  string `json:"tag_value,omitempty"`
	Entity
}

type AssetAttribute struct {
	UUID           string `gorm:"primaryKey" json:"uuid,omitempty"`
	AssetUUID      string `json:"asset_uuid,omitempty"`
	AttributeName  string `json:"attribute_name,omitempty"`
	AttributeValue string `json:"attribute_value,omitempty"`
	Entity
}
