package lib

const (
	XApiKey      = "REPLACE_WITH_YOUR_KEY"
	XProductCode = "REPLACE_WITH_PRODUCT_CODE"
)

const GatewayHost = "https://gateway.snowpal.com/"

const (
	RouteRegistrationRegisterNewUserByEmail = "caa/app/users/sign-up"
	RouteRegistrationSignInByEmail          = "caa/app/users/sign-in"
	RouteRegistrationResetPassword          = "caa/app/users/reset-password"
	RouteRegistrationActivateUser           = "caa/app/user-verified/%s"
)

const (
	RouteUserGetUsers              = "caa/users"
	RouteUserGetUserByUuid         = "caa/users/uuid/%s"
	RouteUserDeactivateUserAccount = "caa/users/deactivate-account"
	RouteUserDeleteUserAccount     = "caa/users/delete-account"
)

const (
	RouteProfilesAddProfile        = "caa/profiles"
	RouteProfilesGetProfileById    = "caa/profiles/by-id/%s"
	RouteProfilesGetProfileByEmail = "caa/profiles/by-email/%s"
	RouteProfilesGetMyProfile      = "caa/profiles"
	RouteProfilesUpdateProfile     = "caa/profiles"
)

const (
	RouteFieldTypesGetFieldTypes = "caa/field-types"
)

const (
	RouteAttributesCreateTextAttribute               = "caa/attributes/text"
	RouteAttributesCreateNumberAttribute             = "caa/attributes/number"
	RouteAttributesCreateDateAttribute               = "caa/attributes/date"
	RouteAttributesCreateFileAttribute               = "caa/attributes/file"
	RouteAttributesCreateSingleSelectAttribute       = "caa/attributes/single-select"
	RouteAttributesCreateMultiselectAttribute        = "caa/attributes/multiselect"
	RouteAttributesCreateNestedSingleSelectAttribute = "caa/attributes/nested-single-select"
	RouteAttributesGetAttributes                     = "caa/attributes"
	RouteAttributesGetAttributeByID                  = "caa/attributes/%s"
	RouteAttributesGetAttributeByName                = "caa/attributes/by-name/:name"
	RouteAttributesUpdateAttribute                   = "caa/attributes/%s"
	RouteAttributesDeleteAttribute                   = "caa/attributes/%s"
)

const (
	RouteAttributeBagsCreateAttributeBag    = "caa/attribute-bags"
	RouteAttributeBagsGetAttributeBags      = "caa/attribute-bags"
	RouteAttributeBagsGetAttributeBagByID   = "caa/attribute-bags/%s"
	RouteAttributeBagsGetAttributeBagByName = "caa/attribute-bags/by-name?name=%s"
	RouteAttributeBagsUpdateAttributeBag    = "caa/attribute-bags/%s"
	RouteAttributeBagsDeleteAttributeBag    = "caa/attribute-bags/%s"
)

const (
	RouteResourcesAssociateAttributeBagsToResource = "caa/resources/%s/attribute-bags"
	RouteResourcesAddAttributeValuesToResource     = "caa/resources/%s/attribute-bags/%s"
	RouteResourcesGetResourceAttributeBags         = "caa/resources/%s/attribute-bags"
	RouteResourcesGetResourceAttributeBagValues    = "caa/resources/%s/attribute-bags/%s/values"
)
