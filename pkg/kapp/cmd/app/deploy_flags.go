// Copyright 2024 The Carvel Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"fmt"

	ctlapp "carvel.dev/kapp/pkg/kapp/app"
	"github.com/spf13/cobra"
)

type DeployFlags struct {
	ctlapp.PrepareResourcesOpts
	Patch      bool
	AllowEmpty bool

	ExistingNonLabeledResourcesCheck            bool
	ExistingNonLabeledResourcesCheckConcurrency int
	OverrideOwnershipOfExistingResources        bool
	OwnershipOverrideAllowedApps                []string

	AppChangesMaxToKeep int

	DefaultLabelScopingRules bool

	Logs            bool
	LogsAll         bool
	AppMetadataFile string

	DisableGKScoping bool
}

func (s *DeployFlags) Set(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&s.AllowCheck, "allow-check", false, "Enable client-side allowing")
	cmd.Flags().StringSliceVar(&s.AllowedNamespaces, "allow-ns", nil, "Set allowed namespace for resources (does not apply to the app itself)")
	cmd.Flags().BoolVar(&s.AllowAllNamespaces, "allow-all-ns", false, "Set to allow all namespaces for resources (does not apply to the app itself)")
	cmd.Flags().BoolVar(&s.AllowCluster, "allow-cluster", false, "Set to allow cluster level for resources (does not apply to the app itself)")

	cmd.Flags().StringVar(&s.IntoNamespace, "into-ns", "", "Place resources into namespace")
	cmd.Flags().StringSliceVar(&s.MapNamespaces, "map-ns", nil, "Map resources from one namespace into another (could be specified multiple times)")

	cmd.Flags().BoolVarP(&s.Patch, "patch", "p", false, "Add or update existing resources only, never delete any")
	cmd.Flags().BoolVar(&s.AllowEmpty, "dangerous-allow-empty-list-of-resources", false, "Allow to apply empty set of resources (same as running kapp delete)")

	cmd.Flags().BoolVar(&s.ExistingNonLabeledResourcesCheck, "existing-non-labeled-resources-check",
		true, "Find and consider existing non-labeled resources in diff")
	cmd.Flags().IntVar(&s.ExistingNonLabeledResourcesCheckConcurrency, "existing-non-labeled-resources-check-concurrency",
		100, "Concurrency to check for existing non-labeled resources")
	cmd.Flags().BoolVar(&s.OverrideOwnershipOfExistingResources, "dangerous-override-ownership-of-existing-resources",
		false, "Steal existing resources from another app")
	cmd.Flags().StringSliceVar(&s.OwnershipOverrideAllowedApps, "dangerous-ownership-override-allowed-apps", nil,
		"Specify existing apps in the same namespace that existing resources can be stolen from")

	cmd.Flags().BoolVar(&s.DefaultLabelScopingRules, "default-label-scoping-rules",
		true, "Use default label scoping rules")

	cmd.Flags().IntVar(&s.AppChangesMaxToKeep, "app-changes-max-to-keep", ctlapp.AppChangesMaxToKeepDefault, "Maximum number of app changes to keep")

	cmd.Flags().BoolVar(&s.Logs, "logs", true, fmt.Sprintf("Show logs from Pods annotated as '%s'", deployLogsAnnKey))
	cmd.Flags().BoolVar(&s.LogsAll, "logs-all", false, "Show logs from all Pods")
	cmd.Flags().StringVar(&s.AppMetadataFile, "app-metadata-file-output", "", "Set filename to write app metadata")

	cmd.Flags().BoolVar(&s.DisableGKScoping, "dangerous-disable-gk-scoping",
		false, "Disable scoping of resource searching to used GroupKinds")
}
