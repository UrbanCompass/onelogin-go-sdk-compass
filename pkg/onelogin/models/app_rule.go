package models

type Condition struct {
	Source   string `json:"source"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Action struct {
	Action     string   `json:"action"`
	Value      []string `json:"value,omitempty"`
	Expression string   `json:"expression,omitempty"`
	Scriplet   string   `json:"scriplet,omitempty"`
	Macro      string   `json:"macro,omitempty"`
}

type AppRule struct {
	AppID      int         `json:"app_id"`
	Name       string      `json:"name"`
	Enabled    bool        `json:"enabled"`
	Match      string      `json:"match"`
	Position   int         `json:"position,omitempty"`
	Conditions []Condition `json:"conditions"`
	Actions    []Action    `json:"actions"`
}

type AppRuleQuery struct {
	Limit            string  `url:"limit,omitempty"`
	Page             string  `url:"page,omitempty"`
	Cursor           string  `url:"cursor,omitempty"`
	Enabled          bool    `url:"enabled,omitempty"`
	HasCondition     *string `url:"has_condition,omitempty"`
	HasConditionType *string `url:"has_condition_type,omitempty"`
	HasAction        *string `url:"has_action,omitempty"`
	HasActionType    *string `url:"has_action_type,omitempty"`
}

func (q *AppRuleQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"limit":              ValidateString,
		"page":               ValidateString,
		"cursor":             ValidateString,
		"enabled":            ValidateBool,
		"has_condition":      ValidateString,
		"has_condition_type": ValidateString,
		"has_action":         ValidateString,
		"has_action_type":    ValidateString,
	}
}
