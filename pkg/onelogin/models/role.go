package models

// RoleQuery represents available query parameters
type RoleQuery struct {
	Fields      string `url:"fields,omitempty"`
	Limit       string `url:"limit,omitempty"`
	Page        string `url:"page,omitempty"`
	AfterCursor string `url:"after_cursor,omitempty"`
	Cursor      string `url:"cursor,omitempty"`
}

// Role represents the Role resource in OneLogin
type Role struct {
	ID     *int32  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Admins []int32 `json:"admins,omitempty"`
	Apps   []int32 `json:"apps,omitempty"`
	Users  []int32 `json:"users,omitempty"`
}

func (r *RoleQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"fields":  ValidateString,
		"limit":   ValidateString,
		"page":    ValidateString,
		"cursor":  ValidateString,
	}
}
