//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersAgent) DeepCopyInto(out *CBContainersAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersAgent.
func (in *CBContainersAgent) DeepCopy() *CBContainersAgent {
	if in == nil {
		return nil
	}
	out := new(CBContainersAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CBContainersAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersAgentList) DeepCopyInto(out *CBContainersAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CBContainersAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersAgentList.
func (in *CBContainersAgentList) DeepCopy() *CBContainersAgentList {
	if in == nil {
		return nil
	}
	out := new(CBContainersAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CBContainersAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersAgentSpec) DeepCopyInto(out *CBContainersAgentSpec) {
	*out = *in
	out.ApiGatewaySpec = in.ApiGatewaySpec
	in.CoreSpec.DeepCopyInto(&out.CoreSpec)
	in.HardeningSpec.DeepCopyInto(&out.HardeningSpec)
	in.RuntimeSpec.DeepCopyInto(&out.RuntimeSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersAgentSpec.
func (in *CBContainersAgentSpec) DeepCopy() *CBContainersAgentSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersAgentStatus) DeepCopyInto(out *CBContainersAgentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersAgentStatus.
func (in *CBContainersAgentStatus) DeepCopy() *CBContainersAgentStatus {
	if in == nil {
		return nil
	}
	out := new(CBContainersAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersApiGatewaySpec) DeepCopyInto(out *CBContainersApiGatewaySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersApiGatewaySpec.
func (in *CBContainersApiGatewaySpec) DeepCopy() *CBContainersApiGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersApiGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersClusterMonitorSpec) DeepCopyInto(out *CBContainersClusterMonitorSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodTemplateAnnotations != nil {
		in, out := &in.PodTemplateAnnotations, &out.PodTemplateAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Image = in.Image
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	out.Probes = in.Probes
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersClusterMonitorSpec.
func (in *CBContainersClusterMonitorSpec) DeepCopy() *CBContainersClusterMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersClusterMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersCommonProbesSpec) DeepCopyInto(out *CBContainersCommonProbesSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersCommonProbesSpec.
func (in *CBContainersCommonProbesSpec) DeepCopy() *CBContainersCommonProbesSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersCommonProbesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersCoreSpec) DeepCopyInto(out *CBContainersCoreSpec) {
	*out = *in
	out.EventsGatewaySpec = in.EventsGatewaySpec
	in.MonitorSpec.DeepCopyInto(&out.MonitorSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersCoreSpec.
func (in *CBContainersCoreSpec) DeepCopy() *CBContainersCoreSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersCoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersEventsGatewaySpec) DeepCopyInto(out *CBContainersEventsGatewaySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersEventsGatewaySpec.
func (in *CBContainersEventsGatewaySpec) DeepCopy() *CBContainersEventsGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersEventsGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersFileProbesSpec) DeepCopyInto(out *CBContainersFileProbesSpec) {
	*out = *in
	out.CBContainersCommonProbesSpec = in.CBContainersCommonProbesSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersFileProbesSpec.
func (in *CBContainersFileProbesSpec) DeepCopy() *CBContainersFileProbesSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersFileProbesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersHTTPProbesSpec) DeepCopyInto(out *CBContainersHTTPProbesSpec) {
	*out = *in
	out.CBContainersCommonProbesSpec = in.CBContainersCommonProbesSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersHTTPProbesSpec.
func (in *CBContainersHTTPProbesSpec) DeepCopy() *CBContainersHTTPProbesSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersHTTPProbesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersHardeningEnforcerSpec) DeepCopyInto(out *CBContainersHardeningEnforcerSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodTemplateAnnotations != nil {
		in, out := &in.PodTemplateAnnotations, &out.PodTemplateAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReplicasCount != nil {
		in, out := &in.ReplicasCount, &out.ReplicasCount
		*out = new(int32)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Prometheus.DeepCopyInto(&out.Prometheus)
	out.Image = in.Image
	in.Resources.DeepCopyInto(&out.Resources)
	out.Probes = in.Probes
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersHardeningEnforcerSpec.
func (in *CBContainersHardeningEnforcerSpec) DeepCopy() *CBContainersHardeningEnforcerSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersHardeningEnforcerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersHardeningSpec) DeepCopyInto(out *CBContainersHardeningSpec) {
	*out = *in
	out.EventsGatewaySpec = in.EventsGatewaySpec
	in.EnforcerSpec.DeepCopyInto(&out.EnforcerSpec)
	in.StateReporterSpec.DeepCopyInto(&out.StateReporterSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersHardeningSpec.
func (in *CBContainersHardeningSpec) DeepCopy() *CBContainersHardeningSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersHardeningSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersHardeningStateReporterSpec) DeepCopyInto(out *CBContainersHardeningStateReporterSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodTemplateAnnotations != nil {
		in, out := &in.PodTemplateAnnotations, &out.PodTemplateAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Image = in.Image
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	out.Probes = in.Probes
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersHardeningStateReporterSpec.
func (in *CBContainersHardeningStateReporterSpec) DeepCopy() *CBContainersHardeningStateReporterSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersHardeningStateReporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersImageSpec) DeepCopyInto(out *CBContainersImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersImageSpec.
func (in *CBContainersImageSpec) DeepCopy() *CBContainersImageSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersPrometheusSpec) DeepCopyInto(out *CBContainersPrometheusSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersPrometheusSpec.
func (in *CBContainersPrometheusSpec) DeepCopy() *CBContainersPrometheusSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersPrometheusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersRuntimeResolverSpec) DeepCopyInto(out *CBContainersRuntimeResolverSpec) {
	*out = *in
	out.EventsGatewaySpec = in.EventsGatewaySpec
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodTemplateAnnotations != nil {
		in, out := &in.PodTemplateAnnotations, &out.PodTemplateAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReplicasCount != nil {
		in, out := &in.ReplicasCount, &out.ReplicasCount
		*out = new(int32)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Image = in.Image
	in.Resources.DeepCopyInto(&out.Resources)
	out.Probes = in.Probes
	in.Prometheus.DeepCopyInto(&out.Prometheus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersRuntimeResolverSpec.
func (in *CBContainersRuntimeResolverSpec) DeepCopy() *CBContainersRuntimeResolverSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersRuntimeResolverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersRuntimeSensorSpec) DeepCopyInto(out *CBContainersRuntimeSensorSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DaemonSetAnnotations != nil {
		in, out := &in.DaemonSetAnnotations, &out.DaemonSetAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodTemplateAnnotations != nil {
		in, out := &in.PodTemplateAnnotations, &out.PodTemplateAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Image = in.Image
	in.Resources.DeepCopyInto(&out.Resources)
	out.Probes = in.Probes
	in.Prometheus.DeepCopyInto(&out.Prometheus)
	if in.VerbosityLevel != nil {
		in, out := &in.VerbosityLevel, &out.VerbosityLevel
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersRuntimeSensorSpec.
func (in *CBContainersRuntimeSensorSpec) DeepCopy() *CBContainersRuntimeSensorSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersRuntimeSensorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CBContainersRuntimeSpec) DeepCopyInto(out *CBContainersRuntimeSpec) {
	*out = *in
	in.ResolverSpec.DeepCopyInto(&out.ResolverSpec)
	in.SensorSpec.DeepCopyInto(&out.SensorSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CBContainersRuntimeSpec.
func (in *CBContainersRuntimeSpec) DeepCopy() *CBContainersRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(CBContainersRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}
