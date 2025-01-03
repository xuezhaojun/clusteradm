// Copyright Contributors to the Open Cluster Management project
package genericclioptions

import (
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"
	ocmfeature "open-cluster-management.io/api/feature"
	operatorv1 "open-cluster-management.io/api/operator/v1"
)

var HubMutableFeatureGate = featuregate.NewFeatureGate()
var SpokeMutableFeatureGate = featuregate.NewFeatureGate()

func init() {
	utilruntime.Must(HubMutableFeatureGate.Add(ocmfeature.DefaultHubWorkFeatureGates))
	utilruntime.Must(HubMutableFeatureGate.Add(ocmfeature.DefaultHubRegistrationFeatureGates))
	utilruntime.Must(HubMutableFeatureGate.Add(ocmfeature.DefaultHubAddonManagerFeatureGates))
	utilruntime.Must(SpokeMutableFeatureGate.Add(ocmfeature.DefaultSpokeRegistrationFeatureGates))
	utilruntime.Must(SpokeMutableFeatureGate.Add(ocmfeature.DefaultSpokeWorkFeatureGates))
}

func ConvertToFeatureGateAPI(featureGates featuregate.MutableFeatureGate, defaultFeatureGate map[featuregate.Feature]featuregate.FeatureSpec) []operatorv1.FeatureGate {
	var features []operatorv1.FeatureGate
	featureGatesMap := featureGates.GetAll()

	// enable user-specified feature gates
	for feature := range featureGatesMap {
		if _, ok := defaultFeatureGate[feature]; !ok {
			continue
		}
		if featureGates.Enabled(feature) {
			features = append(features, operatorv1.FeatureGate{Feature: string(feature), Mode: operatorv1.FeatureGateModeTypeEnable})
		} else if defaultFeatureGate[feature].Default {
			// Explicitly disable the feature gate that is enabled by default
			features = append(features, operatorv1.FeatureGate{Feature: string(feature), Mode: operatorv1.FeatureGateModeTypeDisable})
		}
	}

	// enable default feature gates
	for feature, spec := range defaultFeatureGate {
		if _, ok := featureGatesMap[feature]; !ok && spec.Default {
			features = append(features, operatorv1.FeatureGate{Feature: string(feature), Mode: operatorv1.FeatureGateModeTypeEnable})
		}
	}

	return features
}
