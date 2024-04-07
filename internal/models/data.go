package models

import (
	"time"

	"github.com/FM1337/ASB/internal/ent/serverconfig"
)

type ServerConfig struct {
	/** Bools */
	RemoveRoles      bool
	GiveRole         bool
	Timeout          bool
	Kick             bool
	Ban              bool
	CheckInvites     bool
	CheckLinks       bool
	EnforceRatelimit bool
	SendAlerts       bool
	FlagLinks        bool

	/** discord ids */
	LogChannel       string
	GivenRole        string
	ExcludedChannels []string
	ExcludedRoles    []string
	ExcluedUsers     []string

	/** pre defined values */
	RateLimitTime        serverconfig.RatelimitTime
	TimeoutTime          serverconfig.TimeoutTime
	BanMessageDeleteTime serverconfig.BanDeleteMessageTime
}

type Cooldown struct {
	UserId  string
	HashId  string
	Count   int
	ResetAt time.Time
}
