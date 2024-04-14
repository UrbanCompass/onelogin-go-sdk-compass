package onelogin

import (
	"github.com/UrbanCompass/onelogin-go-sdk-compass/v4/pkg/onelogin/api"
	olerror "github.com/UrbanCompass/onelogin-go-sdk-compass/v4/pkg/onelogin/error"
	mod "github.com/UrbanCompass/onelogin-go-sdk-compass/v4/pkg/onelogin/models"
)

type OneloginSDKAPI interface {
	GetToken() (string, error)
	GenerateInviteLink(email string) (interface{}, error)
	ListConnectors() (interface{}, error)
	SendInviteLink(email string) (interface{}, error)

	// api_authorizations
	CreateAuthServer(authServer *mod.AuthServer) (interface{}, error)
	GetAuthServers(queryParams mod.Queryable) (interface{}, error)
	GetAuthServerByID(id int, queryParams mod.Queryable) (interface{}, error)
	UpdateAuthServer(id int, authServer mod.AuthServer) (interface{}, error)
	DeleteAuthServer(id int) (interface{}, error)
	CreateAuthServerClaim(id int, claim mod.AccessTokenClaim) (interface{}, error)
	DeleteAuthClaim(id, claimID int) (interface{}, error)
	GetAuthClaims(id int, queryParams mod.Queryable) (interface{}, error)
	UpdateClaim(id, claimID int, claim mod.AccessTokenClaim) (interface{}, error)
	CreateAuthServerScope(id int, scope mod.Scope) (interface{}, error)
	DeleteAuthServerScope(id, scopeID int) (interface{}, error)
	GetAuthServerScopes(id int, queryParams mod.Queryable) (interface{}, error)
	UpdateAuthServerScope(id, scopeID int, scope mod.Scope) (interface{}, error)
	CreateClientApp(id int, clientApp mod.ClientApp) (interface{}, error)
	GetClientApps(id int) (interface{}, error)
	DeleteClientApp(id, clientID int) (interface{}, error)
	UpdateClientApp(id, clientID int, clientApp mod.ClientApp) (interface{}, error)

	// apps
	CreateApp(app mod.App) (interface{}, error)
	GetApps(queryParams mod.Queryable) (interface{}, error)
	GetAppsWithCursor(queryParams mod.Queryable) (interface{}, *string, error)
	GetAppByID(id int, queryParams mod.Queryable) (interface{}, error)
	UpdateApp(id int, app mod.App) (interface{}, error)
	DeleteApp(id int) (interface{}, error)
	CreateAppRule(id int, appRule mod.AppRule) (interface{}, error)
	GetAppRules(id int, queryParams mod.Queryable) (interface{}, error)
	GetAppRulesWithCursor(id int, queryParams mod.Queryable) (interface{}, *string, error)
	GetAppRuleByID(id, ruleID int, queryParams mod.Queryable) (interface{}, error)
	UpdateAppRule(id, ruleID int, appRule mod.AppRule, queryParams map[string]string) (interface{}, error)
	DeleteAppRule(id, ruleID int, queryParams map[string]string) (interface{}, error)
	GetAppUsers(appID int) (interface{}, error)
	GetAppUsersWithCursor(appID int, queryParams mod.Queryable) (interface{}, *string, error)

	// groups
	GetGroupByID(groupID int) (interface{}, error)
	GetGroups() (interface{}, error)

	// mfas
	GetAvailableMFAFactors(userID int) (interface{}, error)
	EnrollMFAFactor(factor mod.EnrollFactorRequest, userID int) (interface{}, error)
	VerifyMFAEnrollment(userID, registrationID, otp int) (interface{}, error)
	ActivateMFAFactor(userID int, request mod.ActivateFactorRequest) (interface{}, error)
	RemoveMFAFactor(userID, deviceID int) (interface{}, error)
	GetEnrolledMFAFactors(userID int) (interface{}, error)
	GenerateMFAToken(userID int, request mod.GenerateMFATokenRequest) (interface{}, error)

	// privileges
	ListPrivileges() (interface{}, error)
	CreatePrivilege(privilege mod.Privilege) (interface{}, error)
	GetPrivilege(privilegeID int) (interface{}, error)
	DeletePrivilege(privilegeID int) (interface{}, error)
	UpdatePrivilege(privilegeID int) (interface{}, error)
	GetPrivilegeUsers(privilegeID int) (interface{}, error)
	AssignUsersToPrivilege(privilegeID int) (interface{}, error)
	RemovePrivilegeFromUser(privilegeID int, userID int) (interface{}, error)
	GetPrivilegeRoles(privilegeID int) (interface{}, error)
	AddPrivilegeToRole(privilegeID int, roleID int) (interface{}, error)
	DeleteRoleFromPrivilege(privilegeID int, roleID int) (interface{}, error)

	// roles
	CreateRole(role *mod.Role) (interface{}, error)
	GetRoles(queryParams mod.Queryable) (interface{}, error)
	GetRolesWithCursor(queryParams mod.Queryable) (interface{}, *string, error)
	GetRoleByID(id int, queryParams mod.Queryable) (interface{}, error)
	UpdateRole(id int, role mod.Role, queryParams map[string]string) (interface{}, error)
	DeleteRole(id int, queryParams map[string]string) (interface{}, error)
	GetRoleUsers(roleID int, queryParams mod.Queryable) (interface{}, error)
	AddRoleUsers(roleID int) (interface{}, error)
	DeleteRoleUsers(roleID int, users []int) (interface{}, error)
	GetRoleAdmins(roleID int) (interface{}, error)
	AddRoleAdmins(roleID int) (interface{}, error)
	DeleteRoleAdmins(roleID int, admins []int) (interface{}, error)
	GetRoleApps(roleID int) (interface{}, error)
	UpdateRoleApps(roleID int, apps []int) (interface{}, error)

	// saml
	VerifyFactorSAML(request mod.VerifyMFATokenRequest) (interface{}, error)
	GenerateSAMLAssertion(request mod.GenerateSAMLTokenRequest) (interface{}, error)

	// smart_hooks
	CreateHook(hook mod.SmartHook) (interface{}, error)
	DeleteHook(hookID int) (interface{}, error)
	GetHook(hookID int, query mod.Queryable) (interface{}, error)
	ListHooks(query mod.Queryable) (interface{}, error)
	UpdateSmartHook(hookID int, hook mod.SmartHook) (interface{}, error)
	ListEnvironmentVariables() (interface{}, error)
	CreateEnvironmentVariable(name, value string) (interface{}, error)
	GetEnvironmentVariable(envVarID int) (interface{}, error)
	UpdateEnvironmentVariable(envVarID int, name, value string) (interface{}, error)
	DeleteEnvironmentVariable(envVarID int) (interface{}, error)
	GetHookLogs(hookID int, query mod.Queryable) (interface{}, error)

	// user_mappings
	ListMappings() (interface{}, error)
	ListMappingsWithCursor(queryParams mod.Queryable) (interface{}, *string, error)
	CreateMapping(mapping mod.UserMapping) (interface{}, error)
	DeleteMapping(mappingID int) (interface{}, error)
	GetMapping(mappingID int) (interface{}, error)
	ListActions() (interface{}, error)
	UpdateMapping(mappingID int) (interface{}, error)
	BulkSortMappings(mappingIDs []int) (interface{}, error)
	ListActionValues(actionValue string) (interface{}, error)
	ListConditionValues(conditionValue string) (interface{}, error)
	ListConditionOperators(conditionValue string) (interface{}, error)
	DryrunMapping(mappingID int) (interface{}, error)
	ListConditions() (interface{}, error)

	// users
	CreateUser(user mod.User) (interface{}, error)
	GetUsers(query mod.Queryable) (interface{}, error)
	GetUsersWithCursor(query mod.Queryable) (interface{}, *string, error)
	GetUsersV1(query mod.Queryable) (interface{}, error)
	GetUserByID(id int, queryParams mod.Queryable) (interface{}, error)
	GetUserByIDV1(id int, queryParams mod.Queryable) (interface{}, error)
	GetUserApps(id int, queryParams mod.Queryable) (interface{}, error)
	UpdateUser(id int, user mod.User) (interface{}, error)
	DeleteUser(id int) (interface{}, error)
	UpdatePasswordSecure(id int) (interface{}, error)
	UpdatePasswordInsecure(id int) (interface{}, error)
	LockUserAccount(id int) (interface{}, error)
	GetUserRoles(id int) (interface{}, error)
	LogOutUser(userID int) (interface{}, error)
	AssignRolesToUser(userID int, roles []int) (interface{}, error)
	SetUserState(userID, state int) (interface{}, error)
	RemoveUserRole(userID int) (interface{}, error)
	GetCustomAttributes() (interface{}, error)
	SetCustomAttributes(userID int, attr interface{}) (interface{}, error)
}

// OneloginSDK represents the Onelogin SDK.
type OneloginSDK struct {
	Client *api.Client
}

// Assert that the implementation satisfies the interface.
var _ OneloginSDKAPI = &OneloginSDK{}

// NewOneloginSDK creates a new instance of the Onelogin SDK.
func NewOneloginSDK() (*OneloginSDK, error) {
	client, err := api.NewClient()
	if err != nil {
		return nil, err
	}
	return &OneloginSDK{Client: client}, nil
}

// GetToken performs the authentication process using the env credentials.
func (sdk *OneloginSDK) GetToken() (string, error) {
	// Call the authenticator to perform the authentication process
	accessTk, err := sdk.Client.Auth.GetToken()
	if err != nil {
		return "", olerror.NewSDKError("Access Token retrieval unsuccessful")
	}
	return accessTk, nil
}

func (sdk *OneloginSDK) GenerateInviteLink(email string) (interface{}, error) {
	p := "api/1/invites/get_invite_link"
	resp, err := sdk.Client.Post(&p, email)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (sdk *OneloginSDK) ListConnectors() (interface{}, error) {
	p := "api/2/connectors"
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (sdk *OneloginSDK) SendInviteLink(email string) (interface{}, error) {
	p := "api/1/invites/send_invite_link"
	resp, err := sdk.Client.Post(&p, email)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
