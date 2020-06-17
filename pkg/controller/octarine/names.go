package octarine

import (
	"fmt"
	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
)

func registrySecretName() (types.NamespacedName, error) {
	ns, err := k8sutil.GetWatchNamespace()
	if err != nil {
		return types.NamespacedName{}, err
	}

	return types.NamespacedName{
		Name:      "octarine-registry-secret",
		Namespace: ns,
	}, nil
}

func guardrailsSecretName() (types.NamespacedName, error) {
	ns, err := k8sutil.GetWatchNamespace()
	if err != nil {
		return types.NamespacedName{}, err
	}

	return types.NamespacedName{
		Name:      "octarine-guardrails-tls",
		Namespace: ns,
	}, nil
}

func guardrailsServiceName() (types.NamespacedName, error) {
	ns, err := k8sutil.GetWatchNamespace()
	if err != nil {
		return types.NamespacedName{}, err
	}

	return types.NamespacedName{
		Name:      "octarine-guardrails",
		Namespace: ns,
	}, nil
}

func guardrailsWebhookName(octarine *unstructured.Unstructured) string {
	return fmt.Sprintf("%s-octarine-guardrails", octarine.GetName())
}
