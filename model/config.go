package model

import (
	"time"
)

// AgentConfig ...
type AgentConfig struct {
	MasterAddr     string
	NodeName       string
	DownloadDir    string
	ListenAddr     string
	ReportInterval time.Duration
}

// Task ... controller ->agent
type Task struct {
	ID          string
	LayerName   string
	Status      string
	Torrent     []byte
	TorrentPath string
}

// AgentStatus agent -> controller
type AgentStatus struct {
	Name  string
	Addr  string // ip:port
	Tasks []Task
}
