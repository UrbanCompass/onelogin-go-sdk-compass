package onelogin

import (
	mod "github.com/UrbanCompass/onelogin-go-sdk-compass/v4/pkg/onelogin/models"
	utl "github.com/UrbanCompass/onelogin-go-sdk-compass/v4/pkg/onelogin/utilities"
)

var (
	RolePathV1 string = "api/1/roles"
	RolePathV2 string = "api/2/roles"
)

func (sdk *OneloginSDK) CreateRole(role *mod.Role) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, role)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// was ListRoles
func (sdk *OneloginSDK) GetRoles(queryParams mod.Queryable) (interface{}, error) {
	p := RolePathV2
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// was ListRoles
func (sdk *OneloginSDK) GetRolesV1(queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV1)
	if err != nil {
		return nil, err
	}

	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetRolesWithCursor(queryParams mod.Queryable) (interface{}, *string, error) {
	p, err := utl.BuildAPIPath(RolePathV2)
	if err != nil {
		return nil, nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, nil, err
	}
	return utl.CheckHTTPResponseWithCursor(resp)
}

func (sdk *OneloginSDK) GetRoleByID(id int, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateRole(id int, role mod.Role, queryParams map[string]string) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, role)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteRole(id int, queryParams map[string]string) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// was ListRoleUsers
func (sdk *OneloginSDK) GetRoleUsers(roleID int, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, roleID, "users")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) AddRoleUsers(roleID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, roleID, "users")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// was removeRoleUsers
func (sdk *OneloginSDK) DeleteRoleUsers(roleID int, users []int) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, roleID, "users")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.DeleteWithBody(&p, users)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetRoleAdmins(roleID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, roleID, "admins")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) AddRoleAdmins(roleID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, roleID, "admins")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// was removeRoleAdmins
func (sdk *OneloginSDK) DeleteRoleAdmins(roleID int, admins []int) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, roleID, "admins")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.DeleteWithBody(&p, admins)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetRoleApps(roleID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, roleID, "apps")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// was setRoleApps
func (sdk *OneloginSDK) UpdateRoleApps(roleID int, apps []int) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePathV2, roleID, "apps")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, apps)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
