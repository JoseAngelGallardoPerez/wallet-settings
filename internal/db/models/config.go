package models

const (
	ScopePrivate = "private"
	ScopePublic  = "public"

	PathSiteUrl = Path("regional/general/site_url")

	PathAutoLogoutStatus            = Path("profile/autologout/status")
	PathAutoLogoutTimeout           = Path("profile/autologout/timeout")
	PathAutoLogoutPadding           = Path("profile/autologout/padding")
	PathAutoLogoutInactivityMessage = Path("profile/autologout/inactivity_message")
	PathAutoLogoutMessage           = Path("profile/autologout/message")
)

type Path string

// Config is the abstract config model
type Config struct {
	ID       uint64 `gorm:"primary_key:yes;column:id;unique_index" json:"-"`
	Path     Path   `gorm:"column:path;unique_index;not null;" json:"path" binding:"required,max=255"`
	Value    string `gorm:"column:value;not null;" json:"value"`
	Scope    string `json:"scope"`
	RootOnly bool   `json:"-"` // "root only" can be modified only by root
}
