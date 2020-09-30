package store

// User holds user's configuration.
type UserV2 struct {
	Email    string `json:"email"`
	Token    string `json:"token,omitempty"`
	KeyStore string `json:"keystore,omitempty"`
}

// ConfigRoot holds root fields of the configuration file.
type ConfigRootV2 struct {
	SchemaVersion  uint      `json:"schemaVersion"`
	Users          []*UserV2 `json:"users"`
	CurrentContext string    `json:"currentContext"`
}
