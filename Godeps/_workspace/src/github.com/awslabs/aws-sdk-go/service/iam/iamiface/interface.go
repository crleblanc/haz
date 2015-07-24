package iamiface

import (
	"github.com/awslabs/aws-sdk-go/service/iam"
)

type IAMAPI interface {
	AddClientIDToOpenIDConnectProvider(*iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error)

	AddRoleToInstanceProfile(*iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error)

	AddUserToGroup(*iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error)

	AttachGroupPolicy(*iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error)

	AttachRolePolicy(*iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error)

	AttachUserPolicy(*iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error)

	ChangePassword(*iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error)

	CreateAccessKey(*iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error)

	CreateAccountAlias(*iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error)

	CreateGroup(*iam.CreateGroupInput) (*iam.CreateGroupOutput, error)

	CreateInstanceProfile(*iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error)

	CreateLoginProfile(*iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error)

	CreateOpenIDConnectProvider(*iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error)

	CreatePolicy(*iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error)

	CreatePolicyVersion(*iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error)

	CreateRole(*iam.CreateRoleInput) (*iam.CreateRoleOutput, error)

	CreateSAMLProvider(*iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error)

	CreateUser(*iam.CreateUserInput) (*iam.CreateUserOutput, error)

	CreateVirtualMFADevice(*iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error)

	DeactivateMFADevice(*iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error)

	DeleteAccessKey(*iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error)

	DeleteAccountAlias(*iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error)

	DeleteAccountPasswordPolicy(*iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error)

	DeleteGroup(*iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error)

	DeleteGroupPolicy(*iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error)

	DeleteInstanceProfile(*iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error)

	DeleteLoginProfile(*iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error)

	DeleteOpenIDConnectProvider(*iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error)

	DeletePolicy(*iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error)

	DeletePolicyVersion(*iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error)

	DeleteRole(*iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error)

	DeleteRolePolicy(*iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error)

	DeleteSAMLProvider(*iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error)

	DeleteServerCertificate(*iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error)

	DeleteSigningCertificate(*iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error)

	DeleteUser(*iam.DeleteUserInput) (*iam.DeleteUserOutput, error)

	DeleteUserPolicy(*iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error)

	DeleteVirtualMFADevice(*iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error)

	DetachGroupPolicy(*iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error)

	DetachRolePolicy(*iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error)

	DetachUserPolicy(*iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error)

	EnableMFADevice(*iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error)

	GenerateCredentialReport(*iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error)

	GetAccessKeyLastUsed(*iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error)

	GetAccountAuthorizationDetails(*iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error)

	GetAccountPasswordPolicy(*iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error)

	GetAccountSummary(*iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error)

	GetCredentialReport(*iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error)

	GetGroup(*iam.GetGroupInput) (*iam.GetGroupOutput, error)

	GetGroupPolicy(*iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error)

	GetInstanceProfile(*iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error)

	GetLoginProfile(*iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error)

	GetOpenIDConnectProvider(*iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error)

	GetPolicy(*iam.GetPolicyInput) (*iam.GetPolicyOutput, error)

	GetPolicyVersion(*iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error)

	GetRole(*iam.GetRoleInput) (*iam.GetRoleOutput, error)

	GetRolePolicy(*iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error)

	GetSAMLProvider(*iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error)

	GetServerCertificate(*iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error)

	GetUser(*iam.GetUserInput) (*iam.GetUserOutput, error)

	GetUserPolicy(*iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error)

	ListAccessKeys(*iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error)

	ListAccountAliases(*iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error)

	ListAttachedGroupPolicies(*iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error)

	ListAttachedRolePolicies(*iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error)

	ListAttachedUserPolicies(*iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error)

	ListEntitiesForPolicy(*iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error)

	ListGroupPolicies(*iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error)

	ListGroups(*iam.ListGroupsInput) (*iam.ListGroupsOutput, error)

	ListGroupsForUser(*iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error)

	ListInstanceProfiles(*iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error)

	ListInstanceProfilesForRole(*iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error)

	ListMFADevices(*iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error)

	ListOpenIDConnectProviders(*iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error)

	ListPolicies(*iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error)

	ListPolicyVersions(*iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error)

	ListRolePolicies(*iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error)

	ListRoles(*iam.ListRolesInput) (*iam.ListRolesOutput, error)

	ListSAMLProviders(*iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error)

	ListServerCertificates(*iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error)

	ListSigningCertificates(*iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error)

	ListUserPolicies(*iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error)

	ListUsers(*iam.ListUsersInput) (*iam.ListUsersOutput, error)

	ListVirtualMFADevices(*iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error)

	PutGroupPolicy(*iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error)

	PutRolePolicy(*iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error)

	PutUserPolicy(*iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error)

	RemoveClientIDFromOpenIDConnectProvider(*iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error)

	RemoveRoleFromInstanceProfile(*iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error)

	RemoveUserFromGroup(*iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error)

	ResyncMFADevice(*iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error)

	SetDefaultPolicyVersion(*iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error)

	UpdateAccessKey(*iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error)

	UpdateAccountPasswordPolicy(*iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error)

	UpdateAssumeRolePolicy(*iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error)

	UpdateGroup(*iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error)

	UpdateLoginProfile(*iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error)

	UpdateOpenIDConnectProviderThumbprint(*iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error)

	UpdateSAMLProvider(*iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error)

	UpdateServerCertificate(*iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error)

	UpdateSigningCertificate(*iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error)

	UpdateUser(*iam.UpdateUserInput) (*iam.UpdateUserOutput, error)

	UploadServerCertificate(*iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error)

	UploadSigningCertificate(*iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error)
}
