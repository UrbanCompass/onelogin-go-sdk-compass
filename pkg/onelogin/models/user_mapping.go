package models

// UserMappingsQuery represents available query parameters for mappings
type UserMappingsQuery struct {
	Limit            string `url:"limit,omitempty"`
	Page             string `url:"page,omitempty"`
	Cursor           string `url:"cursor,omitempty"`
	HasCondition     string `url:"has_condition,omitempty"`
	HasConditionType string `url:"has_condition_type,omitempty"`
	HasAction        string `url:"has_action,omitempty"`
	HasActionType    string `url:"has_action_type,omitempty"`
	Enabled          string `url:"enabled,omitempty"`
}

// UserMapping is the contract for User Mappings.
type UserMapping struct {
	ID         *int32                  `json:"id,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Match      *string                 `json:"match,omitempty"`
	Enabled    *bool                   `json:"enabled,omitempty"`
	Position   *int32                  `json:"position,omitempty"`
	Conditions []UserMappingConditions `json:"conditions"`
	Actions    []UserMappingActions    `json:"actions"`
}

// UserMappingConditions is the contract for User Mapping Conditions.
type UserMappingConditions struct {
	Source   *string `json:"source,omitempty"`
	Operator *string `json:"operator,omitempty"`
	Value    *string `json:"value,omitempty"`
}

// UserMappingActions is the contract for User Mapping Actions.
type UserMappingActions struct {
	Action *string  `json:"action,omitempty"`
	Value  []string `json:"value,omitempty"`
}

func (u *UserMappingsQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"limit":            ValidateString,
		"page":             ValidateString,
		"cursor":           ValidateString,
		"has_condition":    ValidateString,
		"has_condition_id": ValidateString,
		"has_action":       ValidateString,
		"has_action_id":    ValidateString,
		"enabled":          ValidateBool,
	}
}
