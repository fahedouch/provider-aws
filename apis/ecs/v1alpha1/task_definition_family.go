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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TaskDefinitionFamilyParameters defines the desired state of TaskDefinitionFamily
type TaskDefinitionFamilyParameters struct {
	// Region is which region the TaskDefinitionFamily will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// A list of container definitions in JSON format that describe the different
	// containers that make up your task.
	// +kubebuilder:validation:Required
	ContainerDefinitions []*ContainerDefinition `json:"containerDefinitions"`
	// The number of CPU units used by the task. It can be expressed as an integer
	// using CPU units (for example, 1024) or as a string using vCPUs (for example,
	// 1 vCPU or 1 vcpu) in a task definition. String values are converted to an
	// integer indicating the CPU units when the task definition is registered.
	//
	// Task-level CPU and memory parameters are ignored for Windows containers.
	// We recommend specifying container-level resources for Windows containers.
	//
	// If you're using the EC2 launch type, this field is optional. Supported values
	// are between 128 CPU units (0.125 vCPUs) and 10240 CPU units (10 vCPUs). If
	// you do not specify a value, the parameter is ignored.
	//
	// If you're using the Fargate launch type, this field is required and you must
	// use one of the following values, which determines your range of supported
	// values for the memory parameter:
	//
	// The CPU units cannot be less than 1 vCPU when you use Windows containers
	// on Fargate.
	//
	//    * 256 (.25 vCPU) - Available memory values: 512 (0.5 GB), 1024 (1 GB),
	//    2048 (2 GB)
	//
	//    * 512 (.5 vCPU) - Available memory values: 1024 (1 GB), 2048 (2 GB), 3072
	//    (3 GB), 4096 (4 GB)
	//
	//    * 1024 (1 vCPU) - Available memory values: 2048 (2 GB), 3072 (3 GB), 4096
	//    (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)
	//
	//    * 2048 (2 vCPU) - Available memory values: 4096 (4 GB) and 16384 (16 GB)
	//    in increments of 1024 (1 GB)
	//
	//    * 4096 (4 vCPU) - Available memory values: 8192 (8 GB) and 30720 (30 GB)
	//    in increments of 1024 (1 GB)
	//
	//    * 8192 (8 vCPU) - Available memory values: 16 GB and 60 GB in 4 GB increments
	//    This option requires Linux platform 1.4.0 or later.
	//
	//    * 16384 (16vCPU) - Available memory values: 32GB and 120 GB in 8 GB increments
	//    This option requires Linux platform 1.4.0 or later.
	CPU *string `json:"cpu,omitempty"`
	// The amount of ephemeral storage to allocate for the task. This parameter
	// is used to expand the total amount of ephemeral storage available, beyond
	// the default amount, for tasks hosted on Fargate. For more information, see
	// Fargate task storage (https://docs.aws.amazon.com/AmazonECS/latest/userguide/using_data_volumes.html)
	// in the Amazon ECS User Guide for Fargate.
	//
	// For tasks using the Fargate launch type, the task requires the following
	// platforms:
	//
	//    * Linux platform version 1.4.0 or later.
	//
	//    * Windows platform version 1.0.0 or later.
	EphemeralStorage *EphemeralStorage `json:"ephemeralStorage,omitempty"`
	// You must specify a family for a task definition. You can use it track multiple
	// versions of the same task definition. The family is used as a name for your
	// task definition. Up to 255 letters (uppercase and lowercase), numbers, underscores,
	// and hyphens are allowed.
	// +kubebuilder:validation:Required
	Family *string `json:"family"`
	// The Elastic Inference accelerators to use for the containers in the task.
	InferenceAccelerators []*InferenceAccelerator `json:"inferenceAccelerators,omitempty"`
	// The IPC resource namespace to use for the containers in the task. The valid
	// values are host, task, or none. If host is specified, then all containers
	// within the tasks that specified the host IPC mode on the same container instance
	// share the same IPC resources with the host Amazon EC2 instance. If task is
	// specified, all containers within the specified task share the same IPC resources.
	// If none is specified, then IPC resources within the containers of a task
	// are private and not shared with other containers in a task or on the container
	// instance. If no value is specified, then the IPC resource namespace sharing
	// depends on the Docker daemon setting on the container instance. For more
	// information, see IPC settings (https://docs.docker.com/engine/reference/run/#ipc-settings---ipc)
	// in the Docker run reference.
	//
	// If the host IPC mode is used, be aware that there is a heightened risk of
	// undesired IPC namespace expose. For more information, see Docker security
	// (https://docs.docker.com/engine/security/security/).
	//
	// If you are setting namespaced kernel parameters using systemControls for
	// the containers in the task, the following will apply to your IPC resource
	// namespace. For more information, see System Controls (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html)
	// in the Amazon Elastic Container Service Developer Guide.
	//
	//    * For tasks that use the host IPC mode, IPC namespace related systemControls
	//    are not supported.
	//
	//    * For tasks that use the task IPC mode, IPC namespace related systemControls
	//    will apply to all containers within a task.
	//
	// This parameter is not supported for Windows containers or tasks run on Fargate.
	IPCMode *string `json:"ipcMode,omitempty"`
	// The amount of memory (in MiB) used by the task. It can be expressed as an
	// integer using MiB (for example ,1024) or as a string using GB (for example,
	// 1GB or 1 GB) in a task definition. String values are converted to an integer
	// indicating the MiB when the task definition is registered.
	//
	// Task-level CPU and memory parameters are ignored for Windows containers.
	// We recommend specifying container-level resources for Windows containers.
	//
	// If using the EC2 launch type, this field is optional.
	//
	// If using the Fargate launch type, this field is required and you must use
	// one of the following values. This determines your range of supported values
	// for the cpu parameter.
	//
	// The CPU units cannot be less than 1 vCPU when you use Windows containers
	// on Fargate.
	//
	//    * 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25
	//    vCPU)
	//
	//    * 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values:
	//    512 (.5 vCPU)
	//
	//    * 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168
	//    (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	//    * Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) -
	//    Available cpu values: 2048 (2 vCPU)
	//
	//    * Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) -
	//    Available cpu values: 4096 (4 vCPU)
	//
	//    * Between 16 GB and 60 GB in 4 GB increments - Available cpu values: 8192
	//    (8 vCPU) This option requires Linux platform 1.4.0 or later.
	//
	//    * Between 32GB and 120 GB in 8 GB increments - Available cpu values: 16384
	//    (16 vCPU) This option requires Linux platform 1.4.0 or later.
	Memory *string `json:"memory,omitempty"`
	// The Docker networking mode to use for the containers in the task. The valid
	// values are none, bridge, awsvpc, and host. If no network mode is specified,
	// the default is bridge.
	//
	// For Amazon ECS tasks on Fargate, the awsvpc network mode is required. For
	// Amazon ECS tasks on Amazon EC2 Linux instances, any network mode can be used.
	// For Amazon ECS tasks on Amazon EC2 Windows instances, <default> or awsvpc
	// can be used. If the network mode is set to none, you cannot specify port
	// mappings in your container definitions, and the tasks containers do not have
	// external connectivity. The host and awsvpc network modes offer the highest
	// networking performance for containers because they use the EC2 network stack
	// instead of the virtualized network stack provided by the bridge mode.
	//
	// With the host and awsvpc network modes, exposed container ports are mapped
	// directly to the corresponding host port (for the host network mode) or the
	// attached elastic network interface port (for the awsvpc network mode), so
	// you cannot take advantage of dynamic host port mappings.
	//
	// When using the host network mode, you should not run containers using the
	// root user (UID 0). It is considered best practice to use a non-root user.
	//
	// If the network mode is awsvpc, the task is allocated an elastic network interface,
	// and you must specify a NetworkConfiguration value when you create a service
	// or run a task with the task definition. For more information, see Task Networking
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html)
	// in the Amazon Elastic Container Service Developer Guide.
	//
	// If the network mode is host, you cannot run multiple instantiations of the
	// same task on a single container instance when port mappings are used.
	//
	// For more information, see Network settings (https://docs.docker.com/engine/reference/run/#network-settings)
	// in the Docker run reference.
	NetworkMode *string `json:"networkMode,omitempty"`
	// The process namespace to use for the containers in the task. The valid values
	// are host or task. If host is specified, then all containers within the tasks
	// that specified the host PID mode on the same container instance share the
	// same process namespace with the host Amazon EC2 instance. If task is specified,
	// all containers within the specified task share the same process namespace.
	// If no value is specified, the default is a private namespace. For more information,
	// see PID settings (https://docs.docker.com/engine/reference/run/#pid-settings---pid)
	// in the Docker run reference.
	//
	// If the host PID mode is used, be aware that there is a heightened risk of
	// undesired process namespace expose. For more information, see Docker security
	// (https://docs.docker.com/engine/security/security/).
	//
	// This parameter is not supported for Windows containers or tasks run on Fargate.
	PIDMode *string `json:"pidMode,omitempty"`
	// An array of placement constraint objects to use for the task. You can specify
	// a maximum of 10 constraints for each task. This limit includes constraints
	// in the task definition and those specified at runtime.
	PlacementConstraints []*TaskDefinitionPlacementConstraint `json:"placementConstraints,omitempty"`
	// The configuration details for the App Mesh proxy.
	//
	// For tasks hosted on Amazon EC2 instances, the container instances require
	// at least version 1.26.0 of the container agent and at least version 1.26.0-1
	// of the ecs-init package to use a proxy configuration. If your container instances
	// are launched from the Amazon ECS-optimized AMI version 20190301 or later,
	// then they contain the required versions of the container agent and ecs-init.
	// For more information, see Amazon ECS-optimized AMI versions (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-ami-versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	ProxyConfiguration *ProxyConfiguration `json:"proxyConfiguration,omitempty"`
	// The task launch type that Amazon ECS validates the task definition against.
	// A client exception is returned if the task definition doesn't validate against
	// the compatibilities specified. If no value is specified, the parameter is
	// omitted from the response.
	RequiresCompatibilities []*string `json:"requiresCompatibilities,omitempty"`
	// The operating system that your tasks definitions run on. A platform family
	// is specified only for tasks using the Fargate launch type.
	//
	// When you specify a task definition in a service, this value must match the
	// runtimePlatform value of the service.
	RuntimePlatform *RuntimePlatform `json:"runtimePlatform,omitempty"`
	// The metadata that you apply to the task definition to help you categorize
	// and organize them. Each tag consists of a key and an optional value. You
	// define both of them.
	//
	// The following basic restrictions apply to tags:
	//
	//    * Maximum number of tags per resource - 50
	//
	//    * For each resource, each tag key must be unique, and each tag key can
	//    have only one value.
	//
	//    * Maximum key length - 128 Unicode characters in UTF-8
	//
	//    * Maximum value length - 256 Unicode characters in UTF-8
	//
	//    * If your tagging schema is used across multiple services and resources,
	//    remember that other services may have restrictions on allowed characters.
	//    Generally allowed characters are: letters, numbers, and spaces representable
	//    in UTF-8, and the following characters: + - = . _ : / @.
	//
	//    * Tag keys and values are case-sensitive.
	//
	//    * Do not use aws:, AWS:, or any upper or lowercase combination of such
	//    as a prefix for either keys or values as it is reserved for Amazon Web
	//    Services use. You cannot edit or delete tag keys or values with this prefix.
	//    Tags with this prefix do not count against your tags per resource limit.
	Tags                           []*Tag `json:"tags,omitempty"`
	CustomTaskDefinitionParameters `json:",inline"`
}

// TaskDefinitionFamilySpec defines the desired state of TaskDefinitionFamily
type TaskDefinitionFamilySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TaskDefinitionFamilyParameters `json:"forProvider"`
}

// TaskDefinitionFamilyObservation defines the observed state of TaskDefinitionFamily
type TaskDefinitionFamilyObservation struct {
	// The full description of the registered task definition.
	TaskDefinition *TaskDefinition_SDK `json:"taskDefinition,omitempty"`
}

// TaskDefinitionFamilyStatus defines the observed state of TaskDefinitionFamily.
type TaskDefinitionFamilyApiStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TaskDefinitionFamilyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TaskDefinitionFamily is the Schema for the TaskDefinitionFamilies API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TaskDefinitionFamily struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TaskDefinitionFamilySpec      `json:"spec"`
	Status            TaskDefinitionFamilyApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TaskDefinitionFamilyList contains a list of TaskDefinitionFamilies
type TaskDefinitionFamilyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TaskDefinitionFamily `json:"items"`
}

// Repository type metadata.
var (
	TaskDefinitionFamilyKind             = "TaskDefinitionFamily"
	TaskDefinitionFamilyGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TaskDefinitionFamilyKind}.String()
	TaskDefinitionFamilyKindAPIVersion   = TaskDefinitionFamilyKind + "." + GroupVersion.String()
	TaskDefinitionFamilyGroupVersionKind = GroupVersion.WithKind(TaskDefinitionFamilyKind)
)

func init() {
	SchemeBuilder.Register(&TaskDefinitionFamily{}, &TaskDefinitionFamilyList{})
}
