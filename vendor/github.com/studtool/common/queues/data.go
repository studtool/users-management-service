package queues

//go:generate

//easyjson:json
type ProfileToCreateData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type ProfileToDeleteData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type DocumentUserToCreateData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type DocumentUserToDeleteData struct {
	UserID string `json:"userId"`
}
