package types

type Config struct {
	Engine         string
	AuthToken      string
	AuthType       string
	Backend        string
	BasicAuth      bool
	ClientCaKeys   string
	ClientCert     string
	ClientKey      string
	Timeout        int64
	ClientInsecure bool
	BackendNodes   []string
	Password       string
	Scheme         string
	Table          string
	Separator      string
	Username       string
	AppID          string
	UserID         string
	RoleID         string
	SecretID       string
	YAMLFile       []string
	Filter         string
	Path           string
	Role           string
}
