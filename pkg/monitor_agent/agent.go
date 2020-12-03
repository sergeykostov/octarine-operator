package monitor_agent

import (
	"encoding/json"
	"fmt"
	"github.com/octarinesec/octarine-operator/pkg/monitor_agent/grpc_client"
	"github.com/octarinesec/octarine-operator/pkg/monitor_agent/health_checker"
	pb "github.com/octarinesec/octarine-operator/pkg/monitor_agent/protobuf"
	"github.com/octarinesec/octarine-operator/pkg/octarine_api"
	"github.com/octarinesec/octarine-operator/pkg/types"
	admissions "k8s.io/api/admissionregistration/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"time"
)

var (
	logger = logf.Log.WithName("monitor_agent")
)

// Monitoring agent - monitors the installed Octarine components and sends health reports to Octarine backend
type MonitorAgent struct {
	// The OctarineSpec of the monitored Octarine deployment
	OctarineSpec *types.OctarineSpec

	// The interval for sending health reports to the backend
	Interval time.Duration

	// Health checker used for retrieving the health status from the cluster
	healthChecker *health_checker.HealthChecker

	// GRPC client for sending health reports to the backend
	grpcClient *grpc_client.GRPCClient

	// Channel for stopping the agent
	stopChan chan struct{}
}

func NewAgent(namespace string, octarineSpec *types.OctarineSpec, k8sClientset *kubernetes.Clientset) (*MonitorAgent, error) {
	healthChecker := health_checker.NewHealthChecker(logger, namespace, k8sClientset)
	apiClient := octarine_api.NewOctarineApiClient(octarineSpec.Global.Octarine.Account,
		octarineSpec.Global.Octarine.AccessToken, octarineSpec.Global.Octarine.Api)

	grpcClient, err := grpc_client.NewGRPCClient(apiClient, octarineSpec)
	if err != nil {
		logger.Error(err, "error creating grpc connection")
		return nil, err
	}

	return &MonitorAgent{
		OctarineSpec:  octarineSpec,
		Interval:      20 * time.Second,
		healthChecker: healthChecker,
		grpcClient:    grpcClient,
		stopChan:      make(chan struct{}),
	}, nil
}

func (agent *MonitorAgent) Start() {
	go agent.run()
}

func (agent *MonitorAgent) Stop() {
	close(agent.stopChan)
}

func (agent *MonitorAgent) buildValidatingWebhookMessage(webhook admissions.ValidatingWebhookConfiguration) *pb.WebhookHealthReport {
	return &pb.WebhookHealthReport{
		WebhookType: pb.WebhookHealthReport_VALIDATING,
		Uid:         string(webhook.UID),
	}

}

func (agent *MonitorAgent) buildServiceMessage(kind pb.ServiceHealthReport_Kind, name string, replicas *int32,
	containers []corev1.Container, statusObj interface{}, lables map[string]string) (*pb.ServiceHealthReport, error) {

	specContainers := make(map[string]*pb.ContainerSpec)
	for _, container := range containers {
		specContainers[container.Name] = &pb.ContainerSpec{
			Image: container.Image,
		}
	}

	status, err := json.Marshal(statusObj)
	if err != nil {
		return nil, fmt.Errorf("error marshaling status for %v: %v", kind, name)
	}

	spec := &pb.ServiceSpec{
		Containers: specContainers,
	}
	if replicas != nil {
		spec.Replicas = *replicas
	}

	return &pb.ServiceHealthReport{
		Kind:     kind,
		Spec:     spec,
		Status:   status,
		Replicas: make(map[string]*pb.ReplicaHealth),
		Labels:   lables,
	}, nil
}

func (agent *MonitorAgent) buildDeploymentMessage(dep appsv1.Deployment) (*pb.ServiceHealthReport, error) {
	return agent.buildServiceMessage(pb.ServiceHealthReport_DEPLOYMENT, dep.Name, dep.Spec.Replicas, dep.Spec.Template.Spec.Containers, dep.Status, dep.Labels)
}

func (agent *MonitorAgent) buildDaemonSetMessage(daemon appsv1.DaemonSet) (*pb.ServiceHealthReport, error) {
	return agent.buildServiceMessage(pb.ServiceHealthReport_DAEMONSET, daemon.Name, nil, daemon.Spec.Template.Spec.Containers, daemon.Status, daemon.Labels)
}

func (agent *MonitorAgent) buildReplicaMessage(pod corev1.Pod) (*pb.ReplicaHealth, error) {
	specContainers := make(map[string]*pb.ContainerSpec)
	for _, container := range pod.Spec.Containers {
		specContainers[container.Name] = &pb.ContainerSpec{
			Image: container.Image,
		}
	}

	status, err := json.Marshal(pod.Status)
	if err != nil {
		return nil, fmt.Errorf("error marshaling status for pod: %v", pod.Name)
	}

	spec := &pb.ReplicaSpec{
		Containers: specContainers,
	}

	return &pb.ReplicaHealth{
		Node:   pod.Spec.NodeName,
		Spec:   spec,
		Status: status,
	}, nil
}

func (agent *MonitorAgent) buildHealthMessage() (*pb.HealthReport, error) {
	services, err := agent.createHealthReportServices()
	if err != nil {
		return nil, err
	}
	webhooks, err := agent.createHealthReportWebhooks(err)
	if err != nil {
		return nil, err
	}
	enabledComponents := map[string]bool{"nodeguard": bool(agent.OctarineSpec.Nodeguard.Enabled),
		"guardrails": bool(agent.OctarineSpec.Guardrails.Enabled)}
	return &pb.HealthReport{
		Account:           agent.OctarineSpec.Global.Octarine.Account,
		Domain:            agent.OctarineSpec.Global.Octarine.Domain,
		Services:          services,
		Webhooks:          webhooks,
		EnabledComponents: enabledComponents,
		Version:           fmt.Sprintf("%v", agent.OctarineSpec.Global.Octarine.Version),
	}, nil
}

// Create the webhooks map for the HealthReport.
// If could not create a part of the health report, returns the error.
func (agent *MonitorAgent) createHealthReportWebhooks(err error) (map[string]*pb.WebhookHealthReport, error) {
	webhooks := make(map[string]*pb.WebhookHealthReport)
	validatingWebhooks, err := agent.healthChecker.GetValidatingWebhookConfigurations()
	if err != nil {
		return nil, err
	}
	agent.addValidatingWebhooks(validatingWebhooks, webhooks)
	return webhooks, nil
}

// Create the services map for the HealthReport.
// If could not create a part of the health report, returns the error.
func (agent *MonitorAgent) createHealthReportServices() (map[string]*pb.ServiceHealthReport, error) {
	services := make(map[string]*pb.ServiceHealthReport)

	pods, err := agent.healthChecker.GetPods()
	if err != nil {
		return nil, err
	}
	replicasSets, err := agent.healthChecker.GetReplicaSets()
	if err != nil {
		return nil, err
	}
	deployments, err := agent.healthChecker.GetDeployments()
	if err != nil {
		return nil, err
	}
	daemonSets, err := agent.healthChecker.GetDaemonSets()
	if err != nil {
		return nil, err
	}
	agent.addDeploymentsServices(deployments, services)
	agent.addDaemonSetsServices(daemonSets, services)
	agent.updateServicesReplicasByPodsAndReplicaSets(pods, replicasSets, services)
	return services, nil
}

// Update the services replicas attribute by the pods and replica sets data.
func (agent *MonitorAgent) updateServicesReplicasByPodsAndReplicaSets(pods map[string]corev1.Pod, repSets map[string]appsv1.ReplicaSet, services map[string]*pb.ServiceHealthReport) {
	for podName, pod := range pods {
		if len(pod.OwnerReferences) < 1 {
			logger.Info("found pod with no parent", "pod", podName)
			continue
		}
		owner := pod.OwnerReferences[0]
		ownerName := owner.Name
		if owner.Kind == "ReplicaSet" {
			if rs, ok := repSets[owner.Name]; ok && len(rs.OwnerReferences) > 0 {
				ownerName = rs.OwnerReferences[0].Name
			} else {
				logger.Info("couldn't determine pod parent", "pod", podName)
				continue
			}
		}

		if serviceMsg, ok := services[ownerName]; ok {
			replicaMsg, err := agent.buildReplicaMessage(pod)
			if err != nil {
				logger.Error(err, "error getting pod data", "pod", podName)
				continue
			}

			serviceMsg.Replicas[podName] = replicaMsg
		}
	}
}

// Update the services map with the found daemon sets.
// If the daemon set message could not have been created successfully, logs an error and skip this daemon
func (agent *MonitorAgent) addDaemonSetsServices(daemons map[string]appsv1.DaemonSet, services map[string]*pb.ServiceHealthReport) {
	for daemonName, daemon := range daemons {
		serviceMsg, err := agent.buildDaemonSetMessage(daemon)
		if err != nil {
			logger.Error(err, "error building DaemonSet message")
			continue
		}

		if _, ok := services[daemonName]; ok {
			logger.Info("duplicate service name", "service", daemonName)
		}

		services[daemonName] = serviceMsg
	}
}

// Update the services map with the found deployments.
// If the deployment message could not have been created successfully, logs an error and skip this deployment.
func (agent *MonitorAgent) addDeploymentsServices(deps map[string]appsv1.Deployment, services map[string]*pb.ServiceHealthReport) {
	for depName, dep := range deps {
		serviceMsg, err := agent.buildDeploymentMessage(dep)
		if err != nil {
			logger.Error(err, "error building Deployment message")
			continue
		}
		if _, ok := services[depName]; ok {
			logger.Info("duplicate service name", "service", depName)
		}
		services[depName] = serviceMsg
	}
}

func (agent *MonitorAgent) getOctarineValidatingWebhooks() []string {
	guardrailsValidatingWebhook := fmt.Sprintf("%s-guardrails", agent.OctarineSpec.Metadata.Name)
	return []string{guardrailsValidatingWebhook}

}

// Update the webhooks map with the octarine validating webhooks.
// If an octarine validating webhook was not found, logs an error and skip this validating webhook.
func (agent *MonitorAgent) addValidatingWebhooks(foundWebhooks map[string]admissions.ValidatingWebhookConfiguration, webhooks map[string]*pb.WebhookHealthReport) {
	for _, webhookName := range agent.getOctarineValidatingWebhooks() {
		if webhook, ok := foundWebhooks[webhookName]; ok {
			webhookMsg := agent.buildValidatingWebhookMessage(webhook)
			if _, ok := webhooks[webhookName]; ok {
				logger.Info("duplicate webhook name", "webhook", webhookName)
			}
			webhooks[webhookName] = webhookMsg
		} else {
			logger.Info("octarine validating webhook not found.", "webhook", webhookName)
		}
	}
}

func (agent *MonitorAgent) run() {
	for {
		select {
		case <-time.After(agent.Interval):
			message, err := agent.buildHealthMessage()
			if err != nil {
				logger.Error(err, "error building health message")
			}
			err = agent.grpcClient.SendMonitorMessage(message)
			if err != nil {
				logger.Error(err, "error reporting message to backend")
			}
		case <-agent.stopChan:
			return
		}
	}
}
