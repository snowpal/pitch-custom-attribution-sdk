package lib

const (
	XApiKey      = "REPLACE_WITH_YOUR_KEY"
	XProductCode = "REPLACE_WITH_PRODUCT_CODE"
)

const GatewayHost = "https://gateway.snowpal.com/"

const (
	RouteRegistrationRegisterNewUserByEmail = "app/users/sign-up"
	RouteRegistrationSignInByEmail          = "app/users/sign-in"
	RouteRegistrationResetPassword          = "app/users/reset-password"
	RouteRegistrationActivateUser           = "app/user-verified/%s"
)

const (
	RouteUserGetUsers              = "users"
	RouteUserGetUserByUuid         = "users/uuid/%s"
	RouteUserDeactivateUserAccount = "users/deactivate-account"
	RouteUserDeleteUserAccount     = "users/delete-account"
)

const (
	RouteProfilesAddProfile        = "profiles"
	RouteProfilesGetProfileById    = "profiles/by-id/%s"
	RouteProfilesGetProfileByEmail = "profiles/by-email/%s"
	RouteProfilesGetMyProfile      = "profiles"
	RouteProfilesUpdateProfile     = "profiles"
)
