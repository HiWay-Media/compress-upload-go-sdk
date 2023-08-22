package models

import (
	"gorm.io/datatypes"
	"time"
)


type Asset struct {
	UUID                    string         `gorm:"column:uuid;primaryKey" json:"uuid,omitempty"`
	CustomerSiteStorageUUID string         `gorm:"column:customer_site_storage_uuid" json:"customer_site_storage_uuid,omitempty"`
	Location                string         `gorm:"column:location" json:"location,omitempty"`
	Path                    string         `gorm:"column:path" json:"path,omitempty"`
	Filename                string         `gorm:"column:filename" json:"filename,omitempty"`
	FileTipology            string         `gorm:"column:file_tipology" json:"file_tipology,omitempty"`
	MasterFileUUID          string         `gorm:"column:master_file_uuid" json:"master_file_uuid,omitempty"`
	IsOriginal              bool           `gorm:"column:is_original" json:"is_original,omitempty"`
	Probe                   datatypes.JSON `gorm:"column:probe" json:"probe,omitempty"`
	Info                    datatypes.JSON `gorm:"column:info" json:"info,omitempty"`
	QC                      datatypes.JSON `gorm:"column:qc" json:"qc,omitempty"`
	SD                      datatypes.JSON `gorm:"column:sd" json:"sd,omitempty"`
	Sprites                 datatypes.JSON `gorm:"column:sprites" json:"sprites,omitempty"`
	Fingerprint             string         `gorm:"column:fingerprint" json:"fingerprint,omitempty"`
	Size                    *int64         `gorm:"column:size" json:"size,omitempty"`
	Published               bool           `gorm:"column:published" json:"published,omitempty"`
	FileCreatedAt           *time.Time     `gorm:"column:file_created_at" json:"file_created_at,omitempty"`
	FileUpdatedAt           *time.Time     `gorm:"column:file_updated_at" json:"file_updated_at,omitempty"`
	JobID                   *int           `gorm:"column:jobid" json:"job_id,omitempty"`
	FileURL                 string         `gorm:"column:file_url" json:"file_url,omitempty"`
	Title                   string         `gorm:"column:title" json:"title,omitempty"`
	JobIDCreatedAt          *time.Time     `gorm:"column:jobid_created_at" json:"job_id_created_at,omitempty"`
	State                   *int           `gorm:"column:state" json:"state,omitempty"`
	Folder                  string         `gorm:"column:folder" json:"folder,omitempty"`
	Token                   string         `gorm:"column:token" json:"token,omitempty"`
	TotalSize               *int64         `gorm:"column:total_size" json:"total_size,omitempty"`
	CategoryID              *int           `gorm:"column:category_id" json:"category_id,omitempty"`
	URLDomain               string         `gorm:"column:url_domain" json:"url_domain,omitempty"`
	IsMultimedia            bool           `gorm:"column:is_multimedia" json:"is_multimedia,omitempty"`
	CreatedAt               *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt               *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt               *time.Time     `gorm:"column:deleted_at;index" json:"-"`
}

type AssetTag struct {
	UUID      string     `gorm:"primaryKey" json:"uuid,omitempty"`
	AssetUUID string     `json:"asset_uuid,omitempty"`
	TagName   string     `json:"tag_name,omitempty"`
	TagValue  string     `json:"tag_value,omitempty"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index" json:"-"`
}

type AssetAttribute struct {
	UUID           string     `gorm:"primaryKey" json:"uuid,omitempty"`
	AssetUUID      string     `json:"asset_uuid,omitempty"`
	AttributeName  string     `json:"attribute_name,omitempty"`
	AttributeType  string     `json:"attribute_type,omitempty"`
	AttributeValue string     `json:"attribute_value,omitempty"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at;index" json:"-"`
}