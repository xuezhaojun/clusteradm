// Copyright Contributors to the Open Cluster Management project
package addon

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"open-cluster-management.io/clusteradm/pkg/cmd/addon/create"
	"open-cluster-management.io/clusteradm/pkg/cmd/addon/disable"
	"open-cluster-management.io/clusteradm/pkg/cmd/addon/enable"
	genericclioptionsclusteradm "open-cluster-management.io/clusteradm/pkg/genericclioptions"
)

// NewCmd provides a cobra command wrapping NewCmdImportCluster
func NewCmd(clusteradmFlags *genericclioptionsclusteradm.ClusteradmFlags, streams genericiooptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addon",
		Short: "addon options",
		Long:  "there are 3 addon options: create, enable and disable",
	}

	cmd.AddCommand(enable.NewCmd(clusteradmFlags, streams))
	cmd.AddCommand(disable.NewCmd(clusteradmFlags, streams))
	cmd.AddCommand(create.NewCmd(clusteradmFlags, streams))

	return cmd
}
