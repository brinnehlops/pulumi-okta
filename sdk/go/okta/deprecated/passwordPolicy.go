// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package deprecated

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type PasswordPolicy struct {
	pulumi.CustomResourceState

	// Authentication Provider: OKTA or ACTIVE_DIRECTORY.
	AuthProvider pulumi.StringPtrOutput `pulumi:"authProvider"`
	// Policy Description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Enable or disable email password recovery: ACTIVE or INACTIVE.
	EmailRecovery pulumi.StringPtrOutput `pulumi:"emailRecovery"`
	// List of Group IDs to Include
	GroupsIncludeds pulumi.StringArrayOutput `pulumi:"groupsIncludeds"`
	// Policy Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of minutes before a locked account is unlocked: 0 = no limit.
	PasswordAutoUnlockMinutes pulumi.IntPtrOutput `pulumi:"passwordAutoUnlockMinutes"`
	// Check Passwords Against Common Password Dictionary.
	PasswordDictionaryLookup pulumi.BoolPtrOutput `pulumi:"passwordDictionaryLookup"`
	// User firstName attribute must be excluded from the password
	PasswordExcludeFirstName pulumi.BoolPtrOutput `pulumi:"passwordExcludeFirstName"`
	// User lastName attribute must be excluded from the password
	PasswordExcludeLastName pulumi.BoolPtrOutput `pulumi:"passwordExcludeLastName"`
	// If the user name must be excluded from the password.
	PasswordExcludeUsername pulumi.BoolPtrOutput `pulumi:"passwordExcludeUsername"`
	// Length in days a user will be warned before password expiry: 0 = no warning.
	PasswordExpireWarnDays pulumi.IntPtrOutput `pulumi:"passwordExpireWarnDays"`
	// Number of distinct passwords that must be created before they can be reused: 0 = none.
	PasswordHistoryCount pulumi.IntPtrOutput `pulumi:"passwordHistoryCount"`
	// Length in days a password is valid before expiry: 0 = no limit.
	PasswordMaxAgeDays pulumi.IntPtrOutput `pulumi:"passwordMaxAgeDays"`
	// Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
	PasswordMaxLockoutAttempts pulumi.IntPtrOutput `pulumi:"passwordMaxLockoutAttempts"`
	// Minimum time interval in minutes between password changes: 0 = no limit.
	PasswordMinAgeMinutes pulumi.IntPtrOutput `pulumi:"passwordMinAgeMinutes"`
	// Minimum password length.
	PasswordMinLength pulumi.IntPtrOutput `pulumi:"passwordMinLength"`
	// If a password must contain at least one lower case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinLowercase pulumi.IntPtrOutput `pulumi:"passwordMinLowercase"`
	// If a password must contain at least one number: 0 = no, 1 = yes. Default = 1
	PasswordMinNumber pulumi.IntPtrOutput `pulumi:"passwordMinNumber"`
	// If a password must contain at least one symbol (!@#$%^&*): 0 = no, 1 = yes. Default = 1
	PasswordMinSymbol pulumi.IntPtrOutput `pulumi:"passwordMinSymbol"`
	// If a password must contain at least one upper case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinUppercase pulumi.IntPtrOutput `pulumi:"passwordMinUppercase"`
	// If a user should be informed when their account is locked.
	PasswordShowLockoutFailures pulumi.BoolPtrOutput `pulumi:"passwordShowLockoutFailures"`
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Min length of the password recovery question answer.
	QuestionMinLength pulumi.IntPtrOutput `pulumi:"questionMinLength"`
	// Enable or disable security question password recovery: ACTIVE or INACTIVE.
	QuestionRecovery pulumi.StringPtrOutput `pulumi:"questionRecovery"`
	// Lifetime in minutes of the recovery email token.
	RecoveryEmailToken pulumi.IntPtrOutput `pulumi:"recoveryEmailToken"`
	// When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's
	// Windows account.
	SkipUnlock pulumi.BoolPtrOutput `pulumi:"skipUnlock"`
	// Enable or disable SMS password recovery: ACTIVE or INACTIVE.
	SmsRecovery pulumi.StringPtrOutput `pulumi:"smsRecovery"`
	// Policy Status: ACTIVE or INACTIVE.
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewPasswordPolicy registers a new resource with the given unique name, arguments, and options.
func NewPasswordPolicy(ctx *pulumi.Context,
	name string, args *PasswordPolicyArgs, opts ...pulumi.ResourceOption) (*PasswordPolicy, error) {
	if args == nil {
		args = &PasswordPolicyArgs{}
	}
	var resource PasswordPolicy
	err := ctx.RegisterResource("okta:deprecated/passwordPolicy:PasswordPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPasswordPolicy gets an existing PasswordPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPasswordPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PasswordPolicyState, opts ...pulumi.ResourceOption) (*PasswordPolicy, error) {
	var resource PasswordPolicy
	err := ctx.ReadResource("okta:deprecated/passwordPolicy:PasswordPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PasswordPolicy resources.
type passwordPolicyState struct {
	// Authentication Provider: OKTA or ACTIVE_DIRECTORY.
	AuthProvider *string `pulumi:"authProvider"`
	// Policy Description
	Description *string `pulumi:"description"`
	// Enable or disable email password recovery: ACTIVE or INACTIVE.
	EmailRecovery *string `pulumi:"emailRecovery"`
	// List of Group IDs to Include
	GroupsIncludeds []string `pulumi:"groupsIncludeds"`
	// Policy Name
	Name *string `pulumi:"name"`
	// Number of minutes before a locked account is unlocked: 0 = no limit.
	PasswordAutoUnlockMinutes *int `pulumi:"passwordAutoUnlockMinutes"`
	// Check Passwords Against Common Password Dictionary.
	PasswordDictionaryLookup *bool `pulumi:"passwordDictionaryLookup"`
	// User firstName attribute must be excluded from the password
	PasswordExcludeFirstName *bool `pulumi:"passwordExcludeFirstName"`
	// User lastName attribute must be excluded from the password
	PasswordExcludeLastName *bool `pulumi:"passwordExcludeLastName"`
	// If the user name must be excluded from the password.
	PasswordExcludeUsername *bool `pulumi:"passwordExcludeUsername"`
	// Length in days a user will be warned before password expiry: 0 = no warning.
	PasswordExpireWarnDays *int `pulumi:"passwordExpireWarnDays"`
	// Number of distinct passwords that must be created before they can be reused: 0 = none.
	PasswordHistoryCount *int `pulumi:"passwordHistoryCount"`
	// Length in days a password is valid before expiry: 0 = no limit.
	PasswordMaxAgeDays *int `pulumi:"passwordMaxAgeDays"`
	// Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
	PasswordMaxLockoutAttempts *int `pulumi:"passwordMaxLockoutAttempts"`
	// Minimum time interval in minutes between password changes: 0 = no limit.
	PasswordMinAgeMinutes *int `pulumi:"passwordMinAgeMinutes"`
	// Minimum password length.
	PasswordMinLength *int `pulumi:"passwordMinLength"`
	// If a password must contain at least one lower case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinLowercase *int `pulumi:"passwordMinLowercase"`
	// If a password must contain at least one number: 0 = no, 1 = yes. Default = 1
	PasswordMinNumber *int `pulumi:"passwordMinNumber"`
	// If a password must contain at least one symbol (!@#$%^&*): 0 = no, 1 = yes. Default = 1
	PasswordMinSymbol *int `pulumi:"passwordMinSymbol"`
	// If a password must contain at least one upper case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinUppercase *int `pulumi:"passwordMinUppercase"`
	// If a user should be informed when their account is locked.
	PasswordShowLockoutFailures *bool `pulumi:"passwordShowLockoutFailures"`
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority *int `pulumi:"priority"`
	// Min length of the password recovery question answer.
	QuestionMinLength *int `pulumi:"questionMinLength"`
	// Enable or disable security question password recovery: ACTIVE or INACTIVE.
	QuestionRecovery *string `pulumi:"questionRecovery"`
	// Lifetime in minutes of the recovery email token.
	RecoveryEmailToken *int `pulumi:"recoveryEmailToken"`
	// When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's
	// Windows account.
	SkipUnlock *bool `pulumi:"skipUnlock"`
	// Enable or disable SMS password recovery: ACTIVE or INACTIVE.
	SmsRecovery *string `pulumi:"smsRecovery"`
	// Policy Status: ACTIVE or INACTIVE.
	Status *string `pulumi:"status"`
}

type PasswordPolicyState struct {
	// Authentication Provider: OKTA or ACTIVE_DIRECTORY.
	AuthProvider pulumi.StringPtrInput
	// Policy Description
	Description pulumi.StringPtrInput
	// Enable or disable email password recovery: ACTIVE or INACTIVE.
	EmailRecovery pulumi.StringPtrInput
	// List of Group IDs to Include
	GroupsIncludeds pulumi.StringArrayInput
	// Policy Name
	Name pulumi.StringPtrInput
	// Number of minutes before a locked account is unlocked: 0 = no limit.
	PasswordAutoUnlockMinutes pulumi.IntPtrInput
	// Check Passwords Against Common Password Dictionary.
	PasswordDictionaryLookup pulumi.BoolPtrInput
	// User firstName attribute must be excluded from the password
	PasswordExcludeFirstName pulumi.BoolPtrInput
	// User lastName attribute must be excluded from the password
	PasswordExcludeLastName pulumi.BoolPtrInput
	// If the user name must be excluded from the password.
	PasswordExcludeUsername pulumi.BoolPtrInput
	// Length in days a user will be warned before password expiry: 0 = no warning.
	PasswordExpireWarnDays pulumi.IntPtrInput
	// Number of distinct passwords that must be created before they can be reused: 0 = none.
	PasswordHistoryCount pulumi.IntPtrInput
	// Length in days a password is valid before expiry: 0 = no limit.
	PasswordMaxAgeDays pulumi.IntPtrInput
	// Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
	PasswordMaxLockoutAttempts pulumi.IntPtrInput
	// Minimum time interval in minutes between password changes: 0 = no limit.
	PasswordMinAgeMinutes pulumi.IntPtrInput
	// Minimum password length.
	PasswordMinLength pulumi.IntPtrInput
	// If a password must contain at least one lower case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinLowercase pulumi.IntPtrInput
	// If a password must contain at least one number: 0 = no, 1 = yes. Default = 1
	PasswordMinNumber pulumi.IntPtrInput
	// If a password must contain at least one symbol (!@#$%^&*): 0 = no, 1 = yes. Default = 1
	PasswordMinSymbol pulumi.IntPtrInput
	// If a password must contain at least one upper case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinUppercase pulumi.IntPtrInput
	// If a user should be informed when their account is locked.
	PasswordShowLockoutFailures pulumi.BoolPtrInput
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrInput
	// Min length of the password recovery question answer.
	QuestionMinLength pulumi.IntPtrInput
	// Enable or disable security question password recovery: ACTIVE or INACTIVE.
	QuestionRecovery pulumi.StringPtrInput
	// Lifetime in minutes of the recovery email token.
	RecoveryEmailToken pulumi.IntPtrInput
	// When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's
	// Windows account.
	SkipUnlock pulumi.BoolPtrInput
	// Enable or disable SMS password recovery: ACTIVE or INACTIVE.
	SmsRecovery pulumi.StringPtrInput
	// Policy Status: ACTIVE or INACTIVE.
	Status pulumi.StringPtrInput
}

func (PasswordPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordPolicyState)(nil)).Elem()
}

type passwordPolicyArgs struct {
	// Authentication Provider: OKTA or ACTIVE_DIRECTORY.
	AuthProvider *string `pulumi:"authProvider"`
	// Policy Description
	Description *string `pulumi:"description"`
	// Enable or disable email password recovery: ACTIVE or INACTIVE.
	EmailRecovery *string `pulumi:"emailRecovery"`
	// List of Group IDs to Include
	GroupsIncludeds []string `pulumi:"groupsIncludeds"`
	// Policy Name
	Name *string `pulumi:"name"`
	// Number of minutes before a locked account is unlocked: 0 = no limit.
	PasswordAutoUnlockMinutes *int `pulumi:"passwordAutoUnlockMinutes"`
	// Check Passwords Against Common Password Dictionary.
	PasswordDictionaryLookup *bool `pulumi:"passwordDictionaryLookup"`
	// User firstName attribute must be excluded from the password
	PasswordExcludeFirstName *bool `pulumi:"passwordExcludeFirstName"`
	// User lastName attribute must be excluded from the password
	PasswordExcludeLastName *bool `pulumi:"passwordExcludeLastName"`
	// If the user name must be excluded from the password.
	PasswordExcludeUsername *bool `pulumi:"passwordExcludeUsername"`
	// Length in days a user will be warned before password expiry: 0 = no warning.
	PasswordExpireWarnDays *int `pulumi:"passwordExpireWarnDays"`
	// Number of distinct passwords that must be created before they can be reused: 0 = none.
	PasswordHistoryCount *int `pulumi:"passwordHistoryCount"`
	// Length in days a password is valid before expiry: 0 = no limit.
	PasswordMaxAgeDays *int `pulumi:"passwordMaxAgeDays"`
	// Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
	PasswordMaxLockoutAttempts *int `pulumi:"passwordMaxLockoutAttempts"`
	// Minimum time interval in minutes between password changes: 0 = no limit.
	PasswordMinAgeMinutes *int `pulumi:"passwordMinAgeMinutes"`
	// Minimum password length.
	PasswordMinLength *int `pulumi:"passwordMinLength"`
	// If a password must contain at least one lower case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinLowercase *int `pulumi:"passwordMinLowercase"`
	// If a password must contain at least one number: 0 = no, 1 = yes. Default = 1
	PasswordMinNumber *int `pulumi:"passwordMinNumber"`
	// If a password must contain at least one symbol (!@#$%^&*): 0 = no, 1 = yes. Default = 1
	PasswordMinSymbol *int `pulumi:"passwordMinSymbol"`
	// If a password must contain at least one upper case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinUppercase *int `pulumi:"passwordMinUppercase"`
	// If a user should be informed when their account is locked.
	PasswordShowLockoutFailures *bool `pulumi:"passwordShowLockoutFailures"`
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority *int `pulumi:"priority"`
	// Min length of the password recovery question answer.
	QuestionMinLength *int `pulumi:"questionMinLength"`
	// Enable or disable security question password recovery: ACTIVE or INACTIVE.
	QuestionRecovery *string `pulumi:"questionRecovery"`
	// Lifetime in minutes of the recovery email token.
	RecoveryEmailToken *int `pulumi:"recoveryEmailToken"`
	// When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's
	// Windows account.
	SkipUnlock *bool `pulumi:"skipUnlock"`
	// Enable or disable SMS password recovery: ACTIVE or INACTIVE.
	SmsRecovery *string `pulumi:"smsRecovery"`
	// Policy Status: ACTIVE or INACTIVE.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a PasswordPolicy resource.
type PasswordPolicyArgs struct {
	// Authentication Provider: OKTA or ACTIVE_DIRECTORY.
	AuthProvider pulumi.StringPtrInput
	// Policy Description
	Description pulumi.StringPtrInput
	// Enable or disable email password recovery: ACTIVE or INACTIVE.
	EmailRecovery pulumi.StringPtrInput
	// List of Group IDs to Include
	GroupsIncludeds pulumi.StringArrayInput
	// Policy Name
	Name pulumi.StringPtrInput
	// Number of minutes before a locked account is unlocked: 0 = no limit.
	PasswordAutoUnlockMinutes pulumi.IntPtrInput
	// Check Passwords Against Common Password Dictionary.
	PasswordDictionaryLookup pulumi.BoolPtrInput
	// User firstName attribute must be excluded from the password
	PasswordExcludeFirstName pulumi.BoolPtrInput
	// User lastName attribute must be excluded from the password
	PasswordExcludeLastName pulumi.BoolPtrInput
	// If the user name must be excluded from the password.
	PasswordExcludeUsername pulumi.BoolPtrInput
	// Length in days a user will be warned before password expiry: 0 = no warning.
	PasswordExpireWarnDays pulumi.IntPtrInput
	// Number of distinct passwords that must be created before they can be reused: 0 = none.
	PasswordHistoryCount pulumi.IntPtrInput
	// Length in days a password is valid before expiry: 0 = no limit.
	PasswordMaxAgeDays pulumi.IntPtrInput
	// Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
	PasswordMaxLockoutAttempts pulumi.IntPtrInput
	// Minimum time interval in minutes between password changes: 0 = no limit.
	PasswordMinAgeMinutes pulumi.IntPtrInput
	// Minimum password length.
	PasswordMinLength pulumi.IntPtrInput
	// If a password must contain at least one lower case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinLowercase pulumi.IntPtrInput
	// If a password must contain at least one number: 0 = no, 1 = yes. Default = 1
	PasswordMinNumber pulumi.IntPtrInput
	// If a password must contain at least one symbol (!@#$%^&*): 0 = no, 1 = yes. Default = 1
	PasswordMinSymbol pulumi.IntPtrInput
	// If a password must contain at least one upper case letter: 0 = no, 1 = yes. Default = 1
	PasswordMinUppercase pulumi.IntPtrInput
	// If a user should be informed when their account is locked.
	PasswordShowLockoutFailures pulumi.BoolPtrInput
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrInput
	// Min length of the password recovery question answer.
	QuestionMinLength pulumi.IntPtrInput
	// Enable or disable security question password recovery: ACTIVE or INACTIVE.
	QuestionRecovery pulumi.StringPtrInput
	// Lifetime in minutes of the recovery email token.
	RecoveryEmailToken pulumi.IntPtrInput
	// When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's
	// Windows account.
	SkipUnlock pulumi.BoolPtrInput
	// Enable or disable SMS password recovery: ACTIVE or INACTIVE.
	SmsRecovery pulumi.StringPtrInput
	// Policy Status: ACTIVE or INACTIVE.
	Status pulumi.StringPtrInput
}

func (PasswordPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordPolicyArgs)(nil)).Elem()
}

