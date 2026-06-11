package models

import "time"

type PeerInfo struct {
	PeerID    string    `json:"peer_id"`
	Addrs     []string  `json:"addrs,omitempty"`
	Hosts     []string  `json:"hosts,omitempty"`
	Timestamp int64     `json:"timestamp"`
	LastResp  time.Time `json:"-"`
	LastSeen  time.Time `json:"-"`
}
