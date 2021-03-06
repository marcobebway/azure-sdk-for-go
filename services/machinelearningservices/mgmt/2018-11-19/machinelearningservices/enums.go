package machinelearningservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AllocationState enumerates the values for allocation state.
type AllocationState string

const (
	// Resizing ...
	Resizing AllocationState = "Resizing"
	// Steady ...
	Steady AllocationState = "Steady"
)

// PossibleAllocationStateValues returns an array of possible values for the AllocationState const type.
func PossibleAllocationStateValues() []AllocationState {
	return []AllocationState{Resizing, Steady}
}

// ComputeType enumerates the values for compute type.
type ComputeType string

const (
	// ComputeTypeAKS ...
	ComputeTypeAKS ComputeType = "AKS"
	// ComputeTypeAmlCompute ...
	ComputeTypeAmlCompute ComputeType = "AmlCompute"
	// ComputeTypeDatabricks ...
	ComputeTypeDatabricks ComputeType = "Databricks"
	// ComputeTypeDataFactory ...
	ComputeTypeDataFactory ComputeType = "DataFactory"
	// ComputeTypeDataLakeAnalytics ...
	ComputeTypeDataLakeAnalytics ComputeType = "DataLakeAnalytics"
	// ComputeTypeHDInsight ...
	ComputeTypeHDInsight ComputeType = "HDInsight"
	// ComputeTypeVirtualMachine ...
	ComputeTypeVirtualMachine ComputeType = "VirtualMachine"
)

// PossibleComputeTypeValues returns an array of possible values for the ComputeType const type.
func PossibleComputeTypeValues() []ComputeType {
	return []ComputeType{ComputeTypeAKS, ComputeTypeAmlCompute, ComputeTypeDatabricks, ComputeTypeDataFactory, ComputeTypeDataLakeAnalytics, ComputeTypeHDInsight, ComputeTypeVirtualMachine}
}

// ComputeTypeBasicCompute enumerates the values for compute type basic compute.
type ComputeTypeBasicCompute string

const (
	// ComputeTypeAKS1 ...
	ComputeTypeAKS1 ComputeTypeBasicCompute = "AKS"
	// ComputeTypeAmlCompute1 ...
	ComputeTypeAmlCompute1 ComputeTypeBasicCompute = "AmlCompute"
	// ComputeTypeCompute ...
	ComputeTypeCompute ComputeTypeBasicCompute = "Compute"
	// ComputeTypeDatabricks1 ...
	ComputeTypeDatabricks1 ComputeTypeBasicCompute = "Databricks"
	// ComputeTypeDataFactory1 ...
	ComputeTypeDataFactory1 ComputeTypeBasicCompute = "DataFactory"
	// ComputeTypeDataLakeAnalytics1 ...
	ComputeTypeDataLakeAnalytics1 ComputeTypeBasicCompute = "DataLakeAnalytics"
	// ComputeTypeHDInsight1 ...
	ComputeTypeHDInsight1 ComputeTypeBasicCompute = "HDInsight"
	// ComputeTypeVirtualMachine1 ...
	ComputeTypeVirtualMachine1 ComputeTypeBasicCompute = "VirtualMachine"
)

// PossibleComputeTypeBasicComputeValues returns an array of possible values for the ComputeTypeBasicCompute const type.
func PossibleComputeTypeBasicComputeValues() []ComputeTypeBasicCompute {
	return []ComputeTypeBasicCompute{ComputeTypeAKS1, ComputeTypeAmlCompute1, ComputeTypeCompute, ComputeTypeDatabricks1, ComputeTypeDataFactory1, ComputeTypeDataLakeAnalytics1, ComputeTypeHDInsight1, ComputeTypeVirtualMachine1}
}

// ComputeTypeBasicComputeNodesInformation enumerates the values for compute type basic compute nodes
// information.
type ComputeTypeBasicComputeNodesInformation string

const (
	// ComputeTypeBasicComputeNodesInformationComputeTypeAmlCompute ...
	ComputeTypeBasicComputeNodesInformationComputeTypeAmlCompute ComputeTypeBasicComputeNodesInformation = "AmlCompute"
	// ComputeTypeBasicComputeNodesInformationComputeTypeComputeNodesInformation ...
	ComputeTypeBasicComputeNodesInformationComputeTypeComputeNodesInformation ComputeTypeBasicComputeNodesInformation = "ComputeNodesInformation"
)

// PossibleComputeTypeBasicComputeNodesInformationValues returns an array of possible values for the ComputeTypeBasicComputeNodesInformation const type.
func PossibleComputeTypeBasicComputeNodesInformationValues() []ComputeTypeBasicComputeNodesInformation {
	return []ComputeTypeBasicComputeNodesInformation{ComputeTypeBasicComputeNodesInformationComputeTypeAmlCompute, ComputeTypeBasicComputeNodesInformationComputeTypeComputeNodesInformation}
}

// ComputeTypeBasicComputeSecrets enumerates the values for compute type basic compute secrets.
type ComputeTypeBasicComputeSecrets string

const (
	// ComputeTypeBasicComputeSecretsComputeTypeAKS ...
	ComputeTypeBasicComputeSecretsComputeTypeAKS ComputeTypeBasicComputeSecrets = "AKS"
	// ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets ...
	ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets ComputeTypeBasicComputeSecrets = "ComputeSecrets"
	// ComputeTypeBasicComputeSecretsComputeTypeDatabricks ...
	ComputeTypeBasicComputeSecretsComputeTypeDatabricks ComputeTypeBasicComputeSecrets = "Databricks"
	// ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine ...
	ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine ComputeTypeBasicComputeSecrets = "VirtualMachine"
)

// PossibleComputeTypeBasicComputeSecretsValues returns an array of possible values for the ComputeTypeBasicComputeSecrets const type.
func PossibleComputeTypeBasicComputeSecretsValues() []ComputeTypeBasicComputeSecrets {
	return []ComputeTypeBasicComputeSecrets{ComputeTypeBasicComputeSecretsComputeTypeAKS, ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets, ComputeTypeBasicComputeSecretsComputeTypeDatabricks, ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Unknown ...
	Unknown ProvisioningState = "Unknown"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Creating, Deleting, Failed, Succeeded, Unknown, Updating}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{SystemAssigned}
}

// Status enumerates the values for status.
type Status string

const (
	// Disabled ...
	Disabled Status = "Disabled"
	// Enabled ...
	Enabled Status = "Enabled"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Disabled, Enabled}
}

// UnderlyingResourceAction enumerates the values for underlying resource action.
type UnderlyingResourceAction string

const (
	// Delete ...
	Delete UnderlyingResourceAction = "Delete"
	// Detach ...
	Detach UnderlyingResourceAction = "Detach"
)

// PossibleUnderlyingResourceActionValues returns an array of possible values for the UnderlyingResourceAction const type.
func PossibleUnderlyingResourceActionValues() []UnderlyingResourceAction {
	return []UnderlyingResourceAction{Delete, Detach}
}

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// Count ...
	Count UsageUnit = "Count"
)

// PossibleUsageUnitValues returns an array of possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{Count}
}

// VMPriority enumerates the values for vm priority.
type VMPriority string

const (
	// Dedicated ...
	Dedicated VMPriority = "Dedicated"
	// LowPriority ...
	LowPriority VMPriority = "LowPriority"
)

// PossibleVMPriorityValues returns an array of possible values for the VMPriority const type.
func PossibleVMPriorityValues() []VMPriority {
	return []VMPriority{Dedicated, LowPriority}
}
