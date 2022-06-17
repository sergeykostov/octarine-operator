package v1

import (
	coreV1 "k8s.io/api/core/v1"
)

type CBContainersRuntimeResolverSpec struct {
	// +kubebuilder:default:=<>
	Labels map[string]string `json:"labels,omitempty"`
	// +kubebuilder:default:=<>
	DeploymentAnnotations map[string]string `json:"deploymentAnnotations,omitempty"`
	// +kubebuilder:default:=<>
	PodTemplateAnnotations map[string]string `json:"podTemplateAnnotations,omitempty"`
	// +kubebuilder:default:=1
	ReplicasCount *int32 `json:"replicasCount,omitempty"`
	// +kubebuilder:default:=<>
	Env map[string]string `json:"env,omitempty"`
	// +kubebuilder:default:={repository:"cbartifactory/runtime-kubernetes-resolver"}
	Image CBContainersImageSpec `json:"image,omitempty"`
	// +kubebuilder:default:={requests: {memory: "64Mi", cpu: "200m"}, limits: {memory: "1024Mi", cpu: "900m"}}
	Resources coreV1.ResourceRequirements `json:"resources,omitempty"`
	// +kubebuilder:default:=<>
	Probes CBContainersHTTPProbesSpec `json:"probes,omitempty"`
	// +kubebuilder:default:={port: 7071}
	Prometheus CBContainersPrometheusSpec `json:"prometheus,omitempty"`
	// +kubebuilder:default:=<>
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// +kubebuilder:default:=<>
	Affinity *coreV1.Affinity `json:"affinity,omitempty"`
}

type CBContainersRuntimeSensorSpec struct {
	// +kubebuilder:default:=<>
	Labels map[string]string `json:"labels,omitempty"`
	// +kubebuilder:default:=<>
	DaemonSetAnnotations map[string]string `json:"daemonSetAnnotations,omitempty"`
	// +kubebuilder:default:=<>
	PodTemplateAnnotations map[string]string `json:"podTemplateAnnotations,omitempty"`
	// +kubebuilder:default:=<>
	Env map[string]string `json:"env,omitempty"`
	// +kubebuilder:default:={repository:"cbartifactory/runtime-kubernetes-sensor"}
	Image CBContainersImageSpec `json:"image,omitempty"`
	// +kubebuilder:default:={requests: {memory: "64Mi", cpu: "30m"}, limits: {memory: "1024Mi", cpu: "500m"}}
	Resources coreV1.ResourceRequirements `json:"resources,omitempty"`
	// +kubebuilder:default:=<>
	Probes CBContainersFileProbesSpec `json:"probes,omitempty"`
	// +kubebuilder:default:={port: 7071}
	Prometheus CBContainersPrometheusSpec `json:"prometheus,omitempty"`
}

// CBContainersRuntimeProtectionSpec defines the desired state of CBContainersRuntime
type CBContainersRuntimeProtectionSpec struct {
	// +kubebuilder:default:=true
	Enabled *bool `json:"enabled,omitempty"`
	// +kubebuilder:default:=<>
	Resolver CBContainersRuntimeResolverSpec `json:"resolver,omitempty"`
	// +kubebuilder:default:=<>
	Sensor CBContainersRuntimeSensorSpec `json:"sensor,omitempty"`
	// +kubebuilder:default:=8080
	InternalGrpcPort int32 `json:"internalGrpcPort,omitempty"`
	// +kubebuilder:default:="info"
	LogVerbosity string `json:"logVerbosity,omitempty"`
}
