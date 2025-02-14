/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AsyncOperationName string

const (
	AsyncOperationName_CreateMultiRegionAccessPoint    AsyncOperationName = "CreateMultiRegionAccessPoint"
	AsyncOperationName_DeleteMultiRegionAccessPoint    AsyncOperationName = "DeleteMultiRegionAccessPoint"
	AsyncOperationName_PutMultiRegionAccessPointPolicy AsyncOperationName = "PutMultiRegionAccessPointPolicy"
)

type BucketCannedACL string

const (
	BucketCannedACL_private            BucketCannedACL = "private"
	BucketCannedACL_public_read        BucketCannedACL = "public-read"
	BucketCannedACL_public_read_write  BucketCannedACL = "public-read-write"
	BucketCannedACL_authenticated_read BucketCannedACL = "authenticated-read"
)

type BucketLocationConstraint string

const (
	BucketLocationConstraint_EU             BucketLocationConstraint = "EU"
	BucketLocationConstraint_eu_west_1      BucketLocationConstraint = "eu-west-1"
	BucketLocationConstraint_us_west_1      BucketLocationConstraint = "us-west-1"
	BucketLocationConstraint_us_west_2      BucketLocationConstraint = "us-west-2"
	BucketLocationConstraint_ap_south_1     BucketLocationConstraint = "ap-south-1"
	BucketLocationConstraint_ap_southeast_1 BucketLocationConstraint = "ap-southeast-1"
	BucketLocationConstraint_ap_southeast_2 BucketLocationConstraint = "ap-southeast-2"
	BucketLocationConstraint_ap_northeast_1 BucketLocationConstraint = "ap-northeast-1"
	BucketLocationConstraint_sa_east_1      BucketLocationConstraint = "sa-east-1"
	BucketLocationConstraint_cn_north_1     BucketLocationConstraint = "cn-north-1"
	BucketLocationConstraint_eu_central_1   BucketLocationConstraint = "eu-central-1"
)

type BucketVersioningStatus string

const (
	BucketVersioningStatus_Enabled   BucketVersioningStatus = "Enabled"
	BucketVersioningStatus_Suspended BucketVersioningStatus = "Suspended"
)

type DeleteMarkerReplicationStatus string

const (
	DeleteMarkerReplicationStatus_Enabled  DeleteMarkerReplicationStatus = "Enabled"
	DeleteMarkerReplicationStatus_Disabled DeleteMarkerReplicationStatus = "Disabled"
)

type ExistingObjectReplicationStatus string

const (
	ExistingObjectReplicationStatus_Enabled  ExistingObjectReplicationStatus = "Enabled"
	ExistingObjectReplicationStatus_Disabled ExistingObjectReplicationStatus = "Disabled"
)

type ExpirationStatus string

const (
	ExpirationStatus_Enabled  ExpirationStatus = "Enabled"
	ExpirationStatus_Disabled ExpirationStatus = "Disabled"
)

type Format string

const (
	Format_CSV     Format = "CSV"
	Format_Parquet Format = "Parquet"
)

type GeneratedManifestFormat string

const (
	GeneratedManifestFormat_S3InventoryReport_CSV_20211130 GeneratedManifestFormat = "S3InventoryReport_CSV_20211130"
)

type JobManifestFieldName string

const (
	JobManifestFieldName_Ignore    JobManifestFieldName = "Ignore"
	JobManifestFieldName_Bucket    JobManifestFieldName = "Bucket"
	JobManifestFieldName_Key       JobManifestFieldName = "Key"
	JobManifestFieldName_VersionId JobManifestFieldName = "VersionId"
)

type JobManifestFormat string

const (
	JobManifestFormat_S3BatchOperations_CSV_20180820 JobManifestFormat = "S3BatchOperations_CSV_20180820"
	JobManifestFormat_S3InventoryReport_CSV_20161130 JobManifestFormat = "S3InventoryReport_CSV_20161130"
)

type JobReportFormat string

const (
	JobReportFormat_Report_CSV_20180820 JobReportFormat = "Report_CSV_20180820"
)

type JobReportScope string

const (
	JobReportScope_AllTasks        JobReportScope = "AllTasks"
	JobReportScope_FailedTasksOnly JobReportScope = "FailedTasksOnly"
)

type JobStatus string

const (
	JobStatus_Active     JobStatus = "Active"
	JobStatus_Cancelled  JobStatus = "Cancelled"
	JobStatus_Cancelling JobStatus = "Cancelling"
	JobStatus_Complete   JobStatus = "Complete"
	JobStatus_Completing JobStatus = "Completing"
	JobStatus_Failed     JobStatus = "Failed"
	JobStatus_Failing    JobStatus = "Failing"
	JobStatus_New        JobStatus = "New"
	JobStatus_Paused     JobStatus = "Paused"
	JobStatus_Pausing    JobStatus = "Pausing"
	JobStatus_Preparing  JobStatus = "Preparing"
	JobStatus_Ready      JobStatus = "Ready"
	JobStatus_Suspended  JobStatus = "Suspended"
)

type MFADelete string

const (
	MFADelete_Enabled  MFADelete = "Enabled"
	MFADelete_Disabled MFADelete = "Disabled"
)

type MFADeleteStatus string

const (
	MFADeleteStatus_Enabled  MFADeleteStatus = "Enabled"
	MFADeleteStatus_Disabled MFADeleteStatus = "Disabled"
)

type MetricsStatus string

const (
	MetricsStatus_Enabled  MetricsStatus = "Enabled"
	MetricsStatus_Disabled MetricsStatus = "Disabled"
)

type MultiRegionAccessPointStatus string

const (
	MultiRegionAccessPointStatus_READY                       MultiRegionAccessPointStatus = "READY"
	MultiRegionAccessPointStatus_INCONSISTENT_ACROSS_REGIONS MultiRegionAccessPointStatus = "INCONSISTENT_ACROSS_REGIONS"
	MultiRegionAccessPointStatus_CREATING                    MultiRegionAccessPointStatus = "CREATING"
	MultiRegionAccessPointStatus_PARTIALLY_CREATED           MultiRegionAccessPointStatus = "PARTIALLY_CREATED"
	MultiRegionAccessPointStatus_PARTIALLY_DELETED           MultiRegionAccessPointStatus = "PARTIALLY_DELETED"
	MultiRegionAccessPointStatus_DELETING                    MultiRegionAccessPointStatus = "DELETING"
)

type NetworkOrigin string

const (
	NetworkOrigin_Internet NetworkOrigin = "Internet"
	NetworkOrigin_VPC      NetworkOrigin = "VPC"
)

type ObjectLambdaAccessPointAliasStatus string

const (
	ObjectLambdaAccessPointAliasStatus_PROVISIONING ObjectLambdaAccessPointAliasStatus = "PROVISIONING"
	ObjectLambdaAccessPointAliasStatus_READY        ObjectLambdaAccessPointAliasStatus = "READY"
)

type ObjectLambdaAllowedFeature string

const (
	ObjectLambdaAllowedFeature_GetObject_Range       ObjectLambdaAllowedFeature = "GetObject-Range"
	ObjectLambdaAllowedFeature_GetObject_PartNumber  ObjectLambdaAllowedFeature = "GetObject-PartNumber"
	ObjectLambdaAllowedFeature_HeadObject_Range      ObjectLambdaAllowedFeature = "HeadObject-Range"
	ObjectLambdaAllowedFeature_HeadObject_PartNumber ObjectLambdaAllowedFeature = "HeadObject-PartNumber"
)

type ObjectLambdaTransformationConfigurationAction string

const (
	ObjectLambdaTransformationConfigurationAction_GetObject     ObjectLambdaTransformationConfigurationAction = "GetObject"
	ObjectLambdaTransformationConfigurationAction_HeadObject    ObjectLambdaTransformationConfigurationAction = "HeadObject"
	ObjectLambdaTransformationConfigurationAction_ListObjects   ObjectLambdaTransformationConfigurationAction = "ListObjects"
	ObjectLambdaTransformationConfigurationAction_ListObjectsV2 ObjectLambdaTransformationConfigurationAction = "ListObjectsV2"
)

type OperationName string

const (
	OperationName_LambdaInvoke            OperationName = "LambdaInvoke"
	OperationName_S3PutObjectCopy         OperationName = "S3PutObjectCopy"
	OperationName_S3PutObjectAcl          OperationName = "S3PutObjectAcl"
	OperationName_S3PutObjectTagging      OperationName = "S3PutObjectTagging"
	OperationName_S3DeleteObjectTagging   OperationName = "S3DeleteObjectTagging"
	OperationName_S3InitiateRestoreObject OperationName = "S3InitiateRestoreObject"
	OperationName_S3PutObjectLegalHold    OperationName = "S3PutObjectLegalHold"
	OperationName_S3PutObjectRetention    OperationName = "S3PutObjectRetention"
	OperationName_S3ReplicateObject       OperationName = "S3ReplicateObject"
)

type OutputSchemaVersion string

const (
	OutputSchemaVersion_V_1 OutputSchemaVersion = "V_1"
)

type OwnerOverride string

const (
	OwnerOverride_Destination OwnerOverride = "Destination"
)

type ReplicaModificationsStatus string

const (
	ReplicaModificationsStatus_Enabled  ReplicaModificationsStatus = "Enabled"
	ReplicaModificationsStatus_Disabled ReplicaModificationsStatus = "Disabled"
)

type ReplicationRuleStatus string

const (
	ReplicationRuleStatus_Enabled  ReplicationRuleStatus = "Enabled"
	ReplicationRuleStatus_Disabled ReplicationRuleStatus = "Disabled"
)

type ReplicationStatus string

const (
	ReplicationStatus_COMPLETED ReplicationStatus = "COMPLETED"
	ReplicationStatus_FAILED    ReplicationStatus = "FAILED"
	ReplicationStatus_REPLICA   ReplicationStatus = "REPLICA"
	ReplicationStatus_NONE      ReplicationStatus = "NONE"
)

type ReplicationStorageClass string

const (
	ReplicationStorageClass_STANDARD            ReplicationStorageClass = "STANDARD"
	ReplicationStorageClass_REDUCED_REDUNDANCY  ReplicationStorageClass = "REDUCED_REDUNDANCY"
	ReplicationStorageClass_STANDARD_IA         ReplicationStorageClass = "STANDARD_IA"
	ReplicationStorageClass_ONEZONE_IA          ReplicationStorageClass = "ONEZONE_IA"
	ReplicationStorageClass_INTELLIGENT_TIERING ReplicationStorageClass = "INTELLIGENT_TIERING"
	ReplicationStorageClass_GLACIER             ReplicationStorageClass = "GLACIER"
	ReplicationStorageClass_DEEP_ARCHIVE        ReplicationStorageClass = "DEEP_ARCHIVE"
	ReplicationStorageClass_OUTPOSTS            ReplicationStorageClass = "OUTPOSTS"
	ReplicationStorageClass_GLACIER_IR          ReplicationStorageClass = "GLACIER_IR"
)

type ReplicationTimeStatus string

const (
	ReplicationTimeStatus_Enabled  ReplicationTimeStatus = "Enabled"
	ReplicationTimeStatus_Disabled ReplicationTimeStatus = "Disabled"
)

type RequestedJobStatus string

const (
	RequestedJobStatus_Cancelled RequestedJobStatus = "Cancelled"
	RequestedJobStatus_Ready     RequestedJobStatus = "Ready"
)

type S3CannedAccessControlList string

const (
	S3CannedAccessControlList_private                   S3CannedAccessControlList = "private"
	S3CannedAccessControlList_public_read               S3CannedAccessControlList = "public-read"
	S3CannedAccessControlList_public_read_write         S3CannedAccessControlList = "public-read-write"
	S3CannedAccessControlList_aws_exec_read             S3CannedAccessControlList = "aws-exec-read"
	S3CannedAccessControlList_authenticated_read        S3CannedAccessControlList = "authenticated-read"
	S3CannedAccessControlList_bucket_owner_read         S3CannedAccessControlList = "bucket-owner-read"
	S3CannedAccessControlList_bucket_owner_full_control S3CannedAccessControlList = "bucket-owner-full-control"
)

type S3ChecksumAlgorithm string

const (
	S3ChecksumAlgorithm_CRC32  S3ChecksumAlgorithm = "CRC32"
	S3ChecksumAlgorithm_CRC32C S3ChecksumAlgorithm = "CRC32C"
	S3ChecksumAlgorithm_SHA1   S3ChecksumAlgorithm = "SHA1"
	S3ChecksumAlgorithm_SHA256 S3ChecksumAlgorithm = "SHA256"
)

type S3GlacierJobTier string

const (
	S3GlacierJobTier_BULK     S3GlacierJobTier = "BULK"
	S3GlacierJobTier_STANDARD S3GlacierJobTier = "STANDARD"
)

type S3GranteeTypeIdentifier string

const (
	S3GranteeTypeIdentifier_id           S3GranteeTypeIdentifier = "id"
	S3GranteeTypeIdentifier_emailAddress S3GranteeTypeIdentifier = "emailAddress"
	S3GranteeTypeIdentifier_uri          S3GranteeTypeIdentifier = "uri"
)

type S3MetadataDirective string

const (
	S3MetadataDirective_COPY    S3MetadataDirective = "COPY"
	S3MetadataDirective_REPLACE S3MetadataDirective = "REPLACE"
)

type S3ObjectLockLegalHoldStatus string

const (
	S3ObjectLockLegalHoldStatus_OFF S3ObjectLockLegalHoldStatus = "OFF"
	S3ObjectLockLegalHoldStatus_ON  S3ObjectLockLegalHoldStatus = "ON"
)

type S3ObjectLockMode string

const (
	S3ObjectLockMode_COMPLIANCE S3ObjectLockMode = "COMPLIANCE"
	S3ObjectLockMode_GOVERNANCE S3ObjectLockMode = "GOVERNANCE"
)

type S3ObjectLockRetentionMode string

const (
	S3ObjectLockRetentionMode_COMPLIANCE S3ObjectLockRetentionMode = "COMPLIANCE"
	S3ObjectLockRetentionMode_GOVERNANCE S3ObjectLockRetentionMode = "GOVERNANCE"
)

type S3Permission string

const (
	S3Permission_FULL_CONTROL S3Permission = "FULL_CONTROL"
	S3Permission_READ         S3Permission = "READ"
	S3Permission_WRITE        S3Permission = "WRITE"
	S3Permission_READ_ACP     S3Permission = "READ_ACP"
	S3Permission_WRITE_ACP    S3Permission = "WRITE_ACP"
)

type S3SSEAlgorithm string

const (
	S3SSEAlgorithm_AES256 S3SSEAlgorithm = "AES256"
	S3SSEAlgorithm_KMS    S3SSEAlgorithm = "KMS"
)

type S3StorageClass string

const (
	S3StorageClass_STANDARD            S3StorageClass = "STANDARD"
	S3StorageClass_STANDARD_IA         S3StorageClass = "STANDARD_IA"
	S3StorageClass_ONEZONE_IA          S3StorageClass = "ONEZONE_IA"
	S3StorageClass_GLACIER             S3StorageClass = "GLACIER"
	S3StorageClass_INTELLIGENT_TIERING S3StorageClass = "INTELLIGENT_TIERING"
	S3StorageClass_DEEP_ARCHIVE        S3StorageClass = "DEEP_ARCHIVE"
	S3StorageClass_GLACIER_IR          S3StorageClass = "GLACIER_IR"
)

type SSEKMSEncryptedObjectsStatus string

const (
	SSEKMSEncryptedObjectsStatus_Enabled  SSEKMSEncryptedObjectsStatus = "Enabled"
	SSEKMSEncryptedObjectsStatus_Disabled SSEKMSEncryptedObjectsStatus = "Disabled"
)

type TransitionStorageClass string

const (
	TransitionStorageClass_GLACIER             TransitionStorageClass = "GLACIER"
	TransitionStorageClass_STANDARD_IA         TransitionStorageClass = "STANDARD_IA"
	TransitionStorageClass_ONEZONE_IA          TransitionStorageClass = "ONEZONE_IA"
	TransitionStorageClass_INTELLIGENT_TIERING TransitionStorageClass = "INTELLIGENT_TIERING"
	TransitionStorageClass_DEEP_ARCHIVE        TransitionStorageClass = "DEEP_ARCHIVE"
)
