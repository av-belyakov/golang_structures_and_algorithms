package mispapi

import "net/url"

type ModuleRequest struct {
	authKey string
	host    string
	port    int
}

type ClientMISP struct {
	BaseURL  *url.URL
	AuthHash string
	Host     string
	Port     int
	Verify   bool
}

type MispElement struct {
	Event EventElement `json:"event"`
}

type EventElement struct {
	Org                OrgElement `json:"org"`
	EventCreatorEmail  any        `json:"event_creator_email"`
	Id                 string     `json:"id"`
	UUID               string     `json:"uuid"`
	OrgId              string     `json:"org_id"`
	OrgcId             string     `json:"orgc_id"`
	Distribution       string     `json:"distribution"`
	Info               string     `json:"info"`
	Date               string     `json:"date"`
	Analysis           string     `json:"analysis"`
	AttributeCount     string     `json:"attribute_count"`
	Timestamp          string     `json:"timestamp"`
	SharingGroupId     string     `json:"sharing_group_id"`
	ThreatLevelId      string     `json:"threat_level_id"`
	PublishTimestamp   string     `json:"publish_timestamp"`
	SightingTimestamp  string     `json:"sighting_timestamp"`
	ExtendsUuid        string     `json:"extends_uuid"`
	Locked             bool       `json:"locked"`
	Published          bool       `json:"published"`
	ProposalEmailLock  bool       `json:"proposal_email_lock"`
	DisableCorrelation bool       `json:"disable_correlation"`
}

type OrgElement struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
