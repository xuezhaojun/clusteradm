// Copyright Contributors to the Open Cluster Management project
package token

import (
	"k8s.io/cli-runtime/pkg/genericiooptions"
	genericclioptionsclusteradm "open-cluster-management.io/clusteradm/pkg/genericclioptions"
)

// Options is holding all the command-line options
type Options struct {
	//ClusteradmFlags: The generic options from the clusteradm cli-runtime.
	ClusteradmFlags *genericclioptionsclusteradm.ClusteradmFlags
	values          Values
	//The file to output the resources will be sent to the file.
	outputFile string
	//If true the bootstrap token will be used instead of the service account token
	useBootstrapToken bool
	//output format
	output string

	Streams genericiooptions.IOStreams
}

// Values: The values used in the template
type Values struct {
	//The values related to the hub
	Hub Hub `json:"hub"`
}

// Hub: The hub values for the template
type Hub struct {
	//TokenID: A token id allowing the cluster to connect back to the hub
	TokenID string `json:"tokenID"`
	//TokenSecret: A token secret allowing the cluster to connect back to the hub
	TokenSecret string `json:"tokenSecret"`
}

func newOptions(clusteradmFlags *genericclioptionsclusteradm.ClusteradmFlags, streams genericiooptions.IOStreams) *Options {
	return &Options{
		ClusteradmFlags: clusteradmFlags,
		Streams:         streams,
	}
}
