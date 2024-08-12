package backingservices

import (
	"testing"

	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
)

func TestConstellationMessagingDeploymentWithAffinity(t *testing.T) {

	var affintiyBasePath = "constellation-messaging.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0]."

	helmChartParser := NewHelmConfigParser(
		NewHelmTestFromTemplate(t, helmChartRelativePath, map[string]string{
			"constellation-messaging.enabled": "true",
			"constellation-messaging.name":    "constellation-messaging",
			affintiyBasePath + "key":          "kubernetes.io/os",
			affintiyBasePath + "operator":     "In",
			affintiyBasePath + "values[0]":    "linux",
		},
			[]string{"charts/constellation-messaging/templates/messaging-deployment.yaml"}),
	)

	var cllnMessagingDeploymentObj appsv1.Deployment
	helmChartParser.getResourceYAML(SearchResourceOption{
		Name: "constellation-messaging",
		Kind: "Deployment",
	}, &cllnMessagingDeploymentObj)

	deploymentAffinity := cllnMessagingDeploymentObj.Spec.Template.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution
	require.Equal(t, "kubernetes.io/os", deploymentAffinity.NodeSelectorTerms[0].MatchExpressions[0].Key)
	require.Equal(t, "In", string(deploymentAffinity.NodeSelectorTerms[0].MatchExpressions[0].Operator))
	require.Equal(t, "linux", deploymentAffinity.NodeSelectorTerms[0].MatchExpressions[0].Values[0])
}
