package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListAccountAssignmentsRequest represents the ListAccountAssignmentsRequest schema from the OpenAPI specification
type ListAccountAssignmentsRequest struct {
	Accountid interface{} `json:"AccountId"`
	Instancearn interface{} `json:"InstanceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// ListInstancesResponse represents the ListInstancesResponse schema from the OpenAPI specification
type ListInstancesResponse struct {
	Instances interface{} `json:"Instances,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteInstanceAccessControlAttributeConfigurationResponse represents the DeleteInstanceAccessControlAttributeConfigurationResponse schema from the OpenAPI specification
type DeleteInstanceAccessControlAttributeConfigurationResponse struct {
}

// DescribeAccountAssignmentCreationStatusResponse represents the DescribeAccountAssignmentCreationStatusResponse schema from the OpenAPI specification
type DescribeAccountAssignmentCreationStatusResponse struct {
	Accountassignmentcreationstatus interface{} `json:"AccountAssignmentCreationStatus,omitempty"`
}

// PutPermissionsBoundaryToPermissionSetRequest represents the PutPermissionsBoundaryToPermissionSetRequest schema from the OpenAPI specification
type PutPermissionsBoundaryToPermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
	Permissionsboundary interface{} `json:"PermissionsBoundary"`
}

// AccessControlAttributeValue represents the AccessControlAttributeValue schema from the OpenAPI specification
type AccessControlAttributeValue struct {
	Source interface{} `json:"Source"`
}

// AccountAssignmentOperationStatusMetadata represents the AccountAssignmentOperationStatusMetadata schema from the OpenAPI specification
type AccountAssignmentOperationStatusMetadata struct {
	Requestid interface{} `json:"RequestId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Resourcearn interface{} `json:"ResourceArn"`
	Tagkeys interface{} `json:"TagKeys"`
}

// DeleteInlinePolicyFromPermissionSetRequest represents the DeleteInlinePolicyFromPermissionSetRequest schema from the OpenAPI specification
type DeleteInlinePolicyFromPermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// CreatePermissionSetResponse represents the CreatePermissionSetResponse schema from the OpenAPI specification
type CreatePermissionSetResponse struct {
	Permissionset interface{} `json:"PermissionSet,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// AttachedManagedPolicy represents the AttachedManagedPolicy schema from the OpenAPI specification
type AttachedManagedPolicy struct {
	Arn interface{} `json:"Arn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// InstanceAccessControlAttributeConfiguration represents the InstanceAccessControlAttributeConfiguration schema from the OpenAPI specification
type InstanceAccessControlAttributeConfiguration struct {
	Accesscontrolattributes interface{} `json:"AccessControlAttributes"`
}

// AccountAssignment represents the AccountAssignment schema from the OpenAPI specification
type AccountAssignment struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Permissionsetarn interface{} `json:"PermissionSetArn,omitempty"`
	Principalid interface{} `json:"PrincipalId,omitempty"`
	Principaltype interface{} `json:"PrincipalType,omitempty"`
}

// DescribePermissionSetRequest represents the DescribePermissionSetRequest schema from the OpenAPI specification
type DescribePermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// DeleteAccountAssignmentRequest represents the DeleteAccountAssignmentRequest schema from the OpenAPI specification
type DeleteAccountAssignmentRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
	Principalid interface{} `json:"PrincipalId"`
	Principaltype interface{} `json:"PrincipalType"`
	Targetid interface{} `json:"TargetId"`
	Targettype interface{} `json:"TargetType"`
}

// ListAccountAssignmentCreationStatusRequest represents the ListAccountAssignmentCreationStatusRequest schema from the OpenAPI specification
type ListAccountAssignmentCreationStatusRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filter interface{} `json:"Filter,omitempty"`
}

// CreateInstanceAccessControlAttributeConfigurationRequest represents the CreateInstanceAccessControlAttributeConfigurationRequest schema from the OpenAPI specification
type CreateInstanceAccessControlAttributeConfigurationRequest struct {
	Instanceaccesscontrolattributeconfiguration interface{} `json:"InstanceAccessControlAttributeConfiguration"`
	Instancearn interface{} `json:"InstanceArn"`
}

// DeleteAccountAssignmentResponse represents the DeleteAccountAssignmentResponse schema from the OpenAPI specification
type DeleteAccountAssignmentResponse struct {
	Accountassignmentdeletionstatus interface{} `json:"AccountAssignmentDeletionStatus,omitempty"`
}

// DeletePermissionSetRequest represents the DeletePermissionSetRequest schema from the OpenAPI specification
type DeletePermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// ListAccountsForProvisionedPermissionSetResponse represents the ListAccountsForProvisionedPermissionSetResponse schema from the OpenAPI specification
type ListAccountsForProvisionedPermissionSetResponse struct {
	Accountids interface{} `json:"AccountIds,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteInstanceAccessControlAttributeConfigurationRequest represents the DeleteInstanceAccessControlAttributeConfigurationRequest schema from the OpenAPI specification
type DeleteInstanceAccessControlAttributeConfigurationRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
}

// AttachManagedPolicyToPermissionSetResponse represents the AttachManagedPolicyToPermissionSetResponse schema from the OpenAPI specification
type AttachManagedPolicyToPermissionSetResponse struct {
}

// ListCustomerManagedPolicyReferencesInPermissionSetResponse represents the ListCustomerManagedPolicyReferencesInPermissionSetResponse schema from the OpenAPI specification
type ListCustomerManagedPolicyReferencesInPermissionSetResponse struct {
	Customermanagedpolicyreferences interface{} `json:"CustomerManagedPolicyReferences,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdatePermissionSetRequest represents the UpdatePermissionSetRequest schema from the OpenAPI specification
type UpdatePermissionSetRequest struct {
	Sessionduration interface{} `json:"SessionDuration,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
	Relaystate interface{} `json:"RelayState,omitempty"`
}

// ListPermissionSetsProvisionedToAccountResponse represents the ListPermissionSetsProvisionedToAccountResponse schema from the OpenAPI specification
type ListPermissionSetsProvisionedToAccountResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Permissionsets interface{} `json:"PermissionSets,omitempty"`
}

// ProvisionPermissionSetResponse represents the ProvisionPermissionSetResponse schema from the OpenAPI specification
type ProvisionPermissionSetResponse struct {
	Permissionsetprovisioningstatus interface{} `json:"PermissionSetProvisioningStatus,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// DetachCustomerManagedPolicyReferenceFromPermissionSetResponse represents the DetachCustomerManagedPolicyReferenceFromPermissionSetResponse schema from the OpenAPI specification
type DetachCustomerManagedPolicyReferenceFromPermissionSetResponse struct {
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// ListPermissionSetsProvisionedToAccountRequest represents the ListPermissionSetsProvisionedToAccountRequest schema from the OpenAPI specification
type ListPermissionSetsProvisionedToAccountRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Provisioningstatus interface{} `json:"ProvisioningStatus,omitempty"`
	Accountid interface{} `json:"AccountId"`
}

// CreateInstanceAccessControlAttributeConfigurationResponse represents the CreateInstanceAccessControlAttributeConfigurationResponse schema from the OpenAPI specification
type CreateInstanceAccessControlAttributeConfigurationResponse struct {
}

// GetPermissionsBoundaryForPermissionSetResponse represents the GetPermissionsBoundaryForPermissionSetResponse schema from the OpenAPI specification
type GetPermissionsBoundaryForPermissionSetResponse struct {
	Permissionsboundary interface{} `json:"PermissionsBoundary,omitempty"`
}

// UpdatePermissionSetResponse represents the UpdatePermissionSetResponse schema from the OpenAPI specification
type UpdatePermissionSetResponse struct {
}

// ListAccountAssignmentCreationStatusResponse represents the ListAccountAssignmentCreationStatusResponse schema from the OpenAPI specification
type ListAccountAssignmentCreationStatusResponse struct {
	Accountassignmentscreationstatus interface{} `json:"AccountAssignmentsCreationStatus,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateAccountAssignmentResponse represents the CreateAccountAssignmentResponse schema from the OpenAPI specification
type CreateAccountAssignmentResponse struct {
	Accountassignmentcreationstatus interface{} `json:"AccountAssignmentCreationStatus,omitempty"`
}

// GetInlinePolicyForPermissionSetResponse represents the GetInlinePolicyForPermissionSetResponse schema from the OpenAPI specification
type GetInlinePolicyForPermissionSetResponse struct {
	Inlinepolicy interface{} `json:"InlinePolicy,omitempty"`
}

// DeletePermissionSetResponse represents the DeletePermissionSetResponse schema from the OpenAPI specification
type DeletePermissionSetResponse struct {
}

// DescribeInstanceAccessControlAttributeConfigurationRequest represents the DescribeInstanceAccessControlAttributeConfigurationRequest schema from the OpenAPI specification
type DescribeInstanceAccessControlAttributeConfigurationRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
}

// PermissionSetProvisioningStatus represents the PermissionSetProvisioningStatus schema from the OpenAPI specification
type PermissionSetProvisioningStatus struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Permissionsetarn interface{} `json:"PermissionSetArn,omitempty"`
	Requestid interface{} `json:"RequestId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DescribeAccountAssignmentDeletionStatusResponse represents the DescribeAccountAssignmentDeletionStatusResponse schema from the OpenAPI specification
type DescribeAccountAssignmentDeletionStatusResponse struct {
	Accountassignmentdeletionstatus interface{} `json:"AccountAssignmentDeletionStatus,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tags interface{} `json:"Tags"`
	Instancearn interface{} `json:"InstanceArn"`
}

// ListPermissionSetsRequest represents the ListPermissionSetsRequest schema from the OpenAPI specification
type ListPermissionSetsRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AttachCustomerManagedPolicyReferenceToPermissionSetRequest represents the AttachCustomerManagedPolicyReferenceToPermissionSetRequest schema from the OpenAPI specification
type AttachCustomerManagedPolicyReferenceToPermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
	Customermanagedpolicyreference interface{} `json:"CustomerManagedPolicyReference"`
}

// PutInlinePolicyToPermissionSetRequest represents the PutInlinePolicyToPermissionSetRequest schema from the OpenAPI specification
type PutInlinePolicyToPermissionSetRequest struct {
	Inlinepolicy interface{} `json:"InlinePolicy"`
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// ProvisionPermissionSetRequest represents the ProvisionPermissionSetRequest schema from the OpenAPI specification
type ProvisionPermissionSetRequest struct {
	Permissionsetarn interface{} `json:"PermissionSetArn"`
	Targetid interface{} `json:"TargetId,omitempty"`
	Targettype interface{} `json:"TargetType"`
	Instancearn interface{} `json:"InstanceArn"`
}

// DescribePermissionSetProvisioningStatusRequest represents the DescribePermissionSetProvisioningStatusRequest schema from the OpenAPI specification
type DescribePermissionSetProvisioningStatusRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Provisionpermissionsetrequestid interface{} `json:"ProvisionPermissionSetRequestId"`
}

// DetachManagedPolicyFromPermissionSetResponse represents the DetachManagedPolicyFromPermissionSetResponse schema from the OpenAPI specification
type DetachManagedPolicyFromPermissionSetResponse struct {
}

// UpdateInstanceAccessControlAttributeConfigurationRequest represents the UpdateInstanceAccessControlAttributeConfigurationRequest schema from the OpenAPI specification
type UpdateInstanceAccessControlAttributeConfigurationRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Instanceaccesscontrolattributeconfiguration interface{} `json:"InstanceAccessControlAttributeConfiguration"`
}

// UpdateInstanceAccessControlAttributeConfigurationResponse represents the UpdateInstanceAccessControlAttributeConfigurationResponse schema from the OpenAPI specification
type UpdateInstanceAccessControlAttributeConfigurationResponse struct {
}

// ListPermissionSetProvisioningStatusResponse represents the ListPermissionSetProvisioningStatusResponse schema from the OpenAPI specification
type ListPermissionSetProvisioningStatusResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Permissionsetsprovisioningstatus interface{} `json:"PermissionSetsProvisioningStatus,omitempty"`
}

// PermissionSet represents the PermissionSet schema from the OpenAPI specification
type PermissionSet struct {
	Relaystate interface{} `json:"RelayState,omitempty"`
	Sessionduration interface{} `json:"SessionDuration,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Permissionsetarn interface{} `json:"PermissionSetArn,omitempty"`
}

// DescribePermissionSetProvisioningStatusResponse represents the DescribePermissionSetProvisioningStatusResponse schema from the OpenAPI specification
type DescribePermissionSetProvisioningStatusResponse struct {
	Permissionsetprovisioningstatus interface{} `json:"PermissionSetProvisioningStatus,omitempty"`
}

// ListPermissionSetProvisioningStatusRequest represents the ListPermissionSetProvisioningStatusRequest schema from the OpenAPI specification
type ListPermissionSetProvisioningStatusRequest struct {
	Filter interface{} `json:"Filter,omitempty"`
	Instancearn interface{} `json:"InstanceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// PermissionsBoundary represents the PermissionsBoundary schema from the OpenAPI specification
type PermissionsBoundary struct {
	Customermanagedpolicyreference interface{} `json:"CustomerManagedPolicyReference,omitempty"`
	Managedpolicyarn interface{} `json:"ManagedPolicyArn,omitempty"`
}

// AccessControlAttribute represents the AccessControlAttribute schema from the OpenAPI specification
type AccessControlAttribute struct {
	Value interface{} `json:"Value"`
	Key interface{} `json:"Key"`
}

// ListInstancesRequest represents the ListInstancesRequest schema from the OpenAPI specification
type ListInstancesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreatePermissionSetRequest represents the CreatePermissionSetRequest schema from the OpenAPI specification
type CreatePermissionSetRequest struct {
	Relaystate interface{} `json:"RelayState,omitempty"`
	Sessionduration interface{} `json:"SessionDuration,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Instancearn interface{} `json:"InstanceArn"`
	Name interface{} `json:"Name"`
}

// DescribeAccountAssignmentCreationStatusRequest represents the DescribeAccountAssignmentCreationStatusRequest schema from the OpenAPI specification
type DescribeAccountAssignmentCreationStatusRequest struct {
	Accountassignmentcreationrequestid interface{} `json:"AccountAssignmentCreationRequestId"`
	Instancearn interface{} `json:"InstanceArn"`
}

// OperationStatusFilter represents the OperationStatusFilter schema from the OpenAPI specification
type OperationStatusFilter struct {
	Status interface{} `json:"Status,omitempty"`
}

// ListAccountsForProvisionedPermissionSetRequest represents the ListAccountsForProvisionedPermissionSetRequest schema from the OpenAPI specification
type ListAccountsForProvisionedPermissionSetRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
	Provisioningstatus interface{} `json:"ProvisioningStatus,omitempty"`
	Instancearn interface{} `json:"InstanceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ListAccountAssignmentDeletionStatusRequest represents the ListAccountAssignmentDeletionStatusRequest schema from the OpenAPI specification
type ListAccountAssignmentDeletionStatusRequest struct {
	Filter interface{} `json:"Filter,omitempty"`
	Instancearn interface{} `json:"InstanceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribePermissionSetResponse represents the DescribePermissionSetResponse schema from the OpenAPI specification
type DescribePermissionSetResponse struct {
	Permissionset interface{} `json:"PermissionSet,omitempty"`
}

// DetachManagedPolicyFromPermissionSetRequest represents the DetachManagedPolicyFromPermissionSetRequest schema from the OpenAPI specification
type DetachManagedPolicyFromPermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Managedpolicyarn interface{} `json:"ManagedPolicyArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// ListAccountAssignmentDeletionStatusResponse represents the ListAccountAssignmentDeletionStatusResponse schema from the OpenAPI specification
type ListAccountAssignmentDeletionStatusResponse struct {
	Accountassignmentsdeletionstatus interface{} `json:"AccountAssignmentsDeletionStatus,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// PermissionSetProvisioningStatusMetadata represents the PermissionSetProvisioningStatusMetadata schema from the OpenAPI specification
type PermissionSetProvisioningStatusMetadata struct {
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Requestid interface{} `json:"RequestId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// PutInlinePolicyToPermissionSetResponse represents the PutInlinePolicyToPermissionSetResponse schema from the OpenAPI specification
type PutInlinePolicyToPermissionSetResponse struct {
}

// GetPermissionsBoundaryForPermissionSetRequest represents the GetPermissionsBoundaryForPermissionSetRequest schema from the OpenAPI specification
type GetPermissionsBoundaryForPermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// AttachManagedPolicyToPermissionSetRequest represents the AttachManagedPolicyToPermissionSetRequest schema from the OpenAPI specification
type AttachManagedPolicyToPermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Managedpolicyarn interface{} `json:"ManagedPolicyArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// InstanceMetadata represents the InstanceMetadata schema from the OpenAPI specification
type InstanceMetadata struct {
	Identitystoreid interface{} `json:"IdentityStoreId,omitempty"`
	Instancearn interface{} `json:"InstanceArn,omitempty"`
}

// CreateAccountAssignmentRequest represents the CreateAccountAssignmentRequest schema from the OpenAPI specification
type CreateAccountAssignmentRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
	Principalid interface{} `json:"PrincipalId"`
	Principaltype interface{} `json:"PrincipalType"`
	Targetid interface{} `json:"TargetId"`
	Targettype interface{} `json:"TargetType"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DetachCustomerManagedPolicyReferenceFromPermissionSetRequest represents the DetachCustomerManagedPolicyReferenceFromPermissionSetRequest schema from the OpenAPI specification
type DetachCustomerManagedPolicyReferenceFromPermissionSetRequest struct {
	Customermanagedpolicyreference interface{} `json:"CustomerManagedPolicyReference"`
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// DescribeAccountAssignmentDeletionStatusRequest represents the DescribeAccountAssignmentDeletionStatusRequest schema from the OpenAPI specification
type DescribeAccountAssignmentDeletionStatusRequest struct {
	Accountassignmentdeletionrequestid interface{} `json:"AccountAssignmentDeletionRequestId"`
	Instancearn interface{} `json:"InstanceArn"`
}

// CustomerManagedPolicyReference represents the CustomerManagedPolicyReference schema from the OpenAPI specification
type CustomerManagedPolicyReference struct {
	Name interface{} `json:"Name"`
	Path interface{} `json:"Path,omitempty"`
}

// DeletePermissionsBoundaryFromPermissionSetResponse represents the DeletePermissionsBoundaryFromPermissionSetResponse schema from the OpenAPI specification
type DeletePermissionsBoundaryFromPermissionSetResponse struct {
}

// DeleteInlinePolicyFromPermissionSetResponse represents the DeleteInlinePolicyFromPermissionSetResponse schema from the OpenAPI specification
type DeleteInlinePolicyFromPermissionSetResponse struct {
}

// DescribeInstanceAccessControlAttributeConfigurationResponse represents the DescribeInstanceAccessControlAttributeConfigurationResponse schema from the OpenAPI specification
type DescribeInstanceAccessControlAttributeConfigurationResponse struct {
	Instanceaccesscontrolattributeconfiguration interface{} `json:"InstanceAccessControlAttributeConfiguration,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusreason interface{} `json:"StatusReason,omitempty"`
}

// AccountAssignmentOperationStatus represents the AccountAssignmentOperationStatus schema from the OpenAPI specification
type AccountAssignmentOperationStatus struct {
	Requestid interface{} `json:"RequestId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Permissionsetarn interface{} `json:"PermissionSetArn,omitempty"`
	Principaltype interface{} `json:"PrincipalType,omitempty"`
	Targetid interface{} `json:"TargetId,omitempty"`
	Targettype interface{} `json:"TargetType,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Principalid interface{} `json:"PrincipalId,omitempty"`
}

// ListManagedPoliciesInPermissionSetRequest represents the ListManagedPoliciesInPermissionSetRequest schema from the OpenAPI specification
type ListManagedPoliciesInPermissionSetRequest struct {
	Permissionsetarn interface{} `json:"PermissionSetArn"`
	Instancearn interface{} `json:"InstanceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetInlinePolicyForPermissionSetRequest represents the GetInlinePolicyForPermissionSetRequest schema from the OpenAPI specification
type GetInlinePolicyForPermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// ListCustomerManagedPolicyReferencesInPermissionSetRequest represents the ListCustomerManagedPolicyReferencesInPermissionSetRequest schema from the OpenAPI specification
type ListCustomerManagedPolicyReferencesInPermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// DeletePermissionsBoundaryFromPermissionSetRequest represents the DeletePermissionsBoundaryFromPermissionSetRequest schema from the OpenAPI specification
type DeletePermissionsBoundaryFromPermissionSetRequest struct {
	Instancearn interface{} `json:"InstanceArn"`
	Permissionsetarn interface{} `json:"PermissionSetArn"`
}

// ListAccountAssignmentsResponse represents the ListAccountAssignmentsResponse schema from the OpenAPI specification
type ListAccountAssignmentsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Accountassignments interface{} `json:"AccountAssignments,omitempty"`
}

// PutPermissionsBoundaryToPermissionSetResponse represents the PutPermissionsBoundaryToPermissionSetResponse schema from the OpenAPI specification
type PutPermissionsBoundaryToPermissionSetResponse struct {
}

// AttachCustomerManagedPolicyReferenceToPermissionSetResponse represents the AttachCustomerManagedPolicyReferenceToPermissionSetResponse schema from the OpenAPI specification
type AttachCustomerManagedPolicyReferenceToPermissionSetResponse struct {
}

// ListManagedPoliciesInPermissionSetResponse represents the ListManagedPoliciesInPermissionSetResponse schema from the OpenAPI specification
type ListManagedPoliciesInPermissionSetResponse struct {
	Attachedmanagedpolicies interface{} `json:"AttachedManagedPolicies,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Value interface{} `json:"Value"`
	Key interface{} `json:"Key"`
}

// ListPermissionSetsResponse represents the ListPermissionSetsResponse schema from the OpenAPI specification
type ListPermissionSetsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Permissionsets interface{} `json:"PermissionSets,omitempty"`
}
