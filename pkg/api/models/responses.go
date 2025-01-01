package models

import (
	"time"
)

type SearchResultMedia struct {
	System System `json:"system"`
	Name   string `json:"name"`
	Path   string `json:"path"`
}

type SearchResults struct {
	Results []SearchResultMedia `json:"results"`
	Total   int                 `json:"total"`
}

type IndexStatusResponse struct {
	Exists      bool   `json:"exists"`
	Indexing    bool   `json:"indexing"`
	TotalSteps  int    `json:"totalSteps"`
	CurrentStep int    `json:"currentStep"`
	CurrentDesc string `json:"currentDesc"`
	TotalFiles  int    `json:"totalFiles"`
}

type SettingsResponse struct {
	RunZapScript            bool     `json:"runZapScript"`
	DebugLogging            bool     `json:"debugLogging"`
	AudioScanFeedback       bool     `json:"audioScanFeedback"`
	ReadersAutoDetect       bool     `json:"readersAutoDetect"`
	ReadersScanMode         string   `json:"readersScanMode"`
	ReadersScanExitDelay    float32  `json:"readersScanExitDelay"`
	ReadersScanIgnoreSystem []string `json:"readersScanIgnoreSystems"`
}

type System struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type SystemsResponse struct {
	Systems []System `json:"systems"`
}

type HistoryReponseEntry struct {
	Time    time.Time `json:"time"`
	Type    string    `json:"type"`
	UID     string    `json:"uid"`
	Text    string    `json:"text"`
	Data    string    `json:"data"`
	Success bool      `json:"success"`
}

type HistoryResponse struct {
	Entries []HistoryReponseEntry `json:"entries"`
}

type AllMappingsResponse struct {
	Mappings []MappingResponse `json:"mappings"`
}

type MappingResponse struct {
	Id       string `json:"id"`
	Added    string `json:"added"`
	Label    string `json:"label"`
	Enabled  bool   `json:"enabled"`
	Type     string `json:"type"`
	Match    string `json:"match"`
	Pattern  string `json:"pattern"`
	Override string `json:"override"`
}

type TokenResponse struct {
	Type     string    `json:"type"`
	UID      string    `json:"uid"`
	Text     string    `json:"text"`
	Data     string    `json:"data"`
	ScanTime time.Time `json:"scanTime"`
}

type IndexResponse struct {
	Exists             bool    `json:"exists"`
	Indexing           bool    `json:"indexing"`
	TotalSteps         *int    `json:"totalSteps,omitempty"`
	CurrentStep        *int    `json:"currentStep,omitempty"`
	CurrentStepDisplay *string `json:"currentStepDisplay,omitempty"`
	TotalFiles         *int    `json:"totalFiles,omitempty"`
}

// TODO: legacy, remove in v2
type ReaderStatusResponse struct {
	Connected bool   `json:"connected"`
	Type      string `json:"type"`
}

type ReaderResponse struct {
	// TODO: type
	Connected bool   `json:"connected"`
	Device    string `json:"device"`
	Info      string `json:"info"`
}

type PlayingResponse struct {
	SystemId   string `json:"systemId"`
	SystemName string `json:"systemName"`
	MediaPath  string `json:"mediaPath"`
	MediaName  string `json:"mediaName"`
}

type StatusResponse struct {
	Reader      ReaderStatusResponse `json:"reader"` // TODO: remove in v2
	Readers     []ReaderResponse     `json:"readers"`
	ActiveToken TokenResponse        `json:"activeToken"`
	LastToken   TokenResponse        `json:"lastToken"`
	Launching   bool                 `json:"launching"`
	GamesIndex  IndexResponse        `json:"gamesIndex"`
	Playing     PlayingResponse      `json:"playing"`
}

type VersionResponse struct {
	Version  string `json:"version"`
	Platform string `json:"platform"`
}

type MediaResponse struct {
	Database IndexResponse     `json:"database"`
	Active   []PlayingResponse `json:"active"`
}
