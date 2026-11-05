---
type: concept
title: Feature Gates (removed)
description: This page contains list of feature gates that have been removed.
resource: https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates-removed/
tags: [kubernetes, reference, command-line-tools-reference]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/command-line-tools-reference/feature-gates-removed/index.md
    title: 'Kubernetes website source: reference/command-line-tools-reference/feature-gates-removed/index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates-removed/
    title: Feature Gates (removed) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/command-line-tools-reference/feature-gates-removed/index.md
---

# Feature Gates (removed)

This page contains list of feature gates that have been removed. The information on this page is for reference.
A removed feature gate is different from a GA'ed or deprecated one in that a removed one is
no longer recognized as a valid feature gate.
However, a GA'ed or a deprecated feature gate is still recognized by the corresponding Kubernetes
components although they are unable to cause any behavior differences in a cluster.

For feature gates that are still recognized by the Kubernetes components, please refer to
the [Alpha/Beta feature gate table](/reference/command-line-tools-reference/feature-gates.md#feature-gates-for-alpha-or-beta-features)
or the [Graduated/Deprecated feature gate table](/reference/command-line-tools-reference/feature-gates.md#feature-gates-for-graduated-or-deprecated-features)

### Feature gates that are removed

In the following table:

- The "From" column contains the Kubernetes release when a feature is introduced
  or its release stage is changed.
- The "To" column, if not empty, contains the last Kubernetes release in which
  you can still use a feature gate. If the feature stage is either "Deprecated"
  or "GA", the "To" column is the Kubernetes release when the feature is removed.

| Feature | Default | Stage | Since | Until |
|---|---|---|---|---|
| `APIListChunking` | `false` | alpha | 1.8 | 1.8 |
| `APIListChunking` | `true` | beta | 1.9 | 1.28 |
| `APIListChunking` | `true` | stable | 1.29 | 1.32 |
| `APIPriorityAndFairness` | `false` | alpha | 1.18 | 1.19 |
| `APIPriorityAndFairness` | `true` | beta | 1.20 | 1.28 |
| `APIPriorityAndFairness` | `true` | stable | 1.29 | 1.30 |
| `APISelfSubjectReview` | `false` | alpha | 1.26 | 1.26 |
| `APISelfSubjectReview` | `true` | beta | 1.27 | 1.27 |
| `APISelfSubjectReview` | `true` | stable | 1.28 | 1.29 |
| `Accelerators` | `false` | alpha | 1.6 | 1.10 |
| `Accelerators` | `<nil>` | deprecated | 1.11 | 1.11 |
| `AdmissionWebhookMatchConditions` | `false` | alpha | 1.27 | 1.27 |
| `AdmissionWebhookMatchConditions` | `true` | beta | 1.28 | 1.29 |
| `AdmissionWebhookMatchConditions` | `true` | stable | 1.30 | 1.32 |
| `AdvancedAuditing` | `false` | alpha | 1.7 | 1.7 |
| `AdvancedAuditing` | `true` | beta | 1.8 | 1.11 |
| `AdvancedAuditing` | `true` | stable | 1.12 | 1.27 |
| `AffinityInAnnotations` | `false` | alpha | 1.6 | 1.7 |
| `AffinityInAnnotations` | `<nil>` | deprecated | 1.8 | 1.8 |
| `AggregatedDiscoveryEndpoint` | `false` | alpha | 1.26 | 1.26 |
| `AggregatedDiscoveryEndpoint` | `true` | beta | 1.27 | 1.29 |
| `AggregatedDiscoveryEndpoint` | `true` | stable | 1.30 | 1.32 |
| `AllowExtTrafficLocalEndpoints` | `false` | beta | 1.4 | 1.6 |
| `AllowExtTrafficLocalEndpoints` | `true` | stable | 1.7 | 1.9 |
| `AllowInsecureBackendProxy` | `true` | beta | 1.17 | 1.20 |
| `AllowInsecureBackendProxy` | `true` | stable | 1.21 | 1.25 |
| `AllowServiceLBStatusOnNonLB` | `false` | deprecated | 1.29 | 1.34 |
| `AppArmor` | `true` | beta | 1.4 | 1.30 |
| `AppArmor` | `true` | stable | 1.31 | 1.32 |
| `AppArmorFields` | `true` | beta | 1.30 | 1.30 |
| `AppArmorFields` | `true` | stable | 1.31 | 1.32 |
| `AttachVolumeLimit` | `false` | alpha | 1.11 | 1.11 |
| `AttachVolumeLimit` | `true` | beta | 1.12 | 1.16 |
| `AttachVolumeLimit` | `true` | stable | 1.17 | 1.21 |
| `BalanceAttachedNodeVolumes` | `false` | alpha | 1.11 | 1.21 |
| `BalanceAttachedNodeVolumes` | `false` | deprecated | 1.22 | 1.22 |
| `BlockVolume` | `false` | alpha | 1.9 | 1.12 |
| `BlockVolume` | `true` | beta | 1.13 | 1.17 |
| `BlockVolume` | `true` | stable | 1.18 | 1.21 |
| `BoundServiceAccountTokenVolume` | `false` | alpha | 1.13 | 1.20 |
| `BoundServiceAccountTokenVolume` | `true` | beta | 1.21 | 1.21 |
| `BoundServiceAccountTokenVolume` | `true` | stable | 1.22 | 1.23 |
| `CPUManager` | `false` | alpha | 1.8 | 1.9 |
| `CPUManager` | `true` | beta | 1.10 | 1.25 |
| `CPUManager` | `true` | stable | 1.26 | 1.32 |
| `CRIContainerLogRotation` | `false` | alpha | 1.10 | 1.10 |
| `CRIContainerLogRotation` | `true` | beta | 1.11 | 1.20 |
| `CRIContainerLogRotation` | `true` | stable | 1.21 | 1.22 |
| `CSIBlockVolume` | `false` | alpha | 1.11 | 1.13 |
| `CSIBlockVolume` | `true` | beta | 1.14 | 1.17 |
| `CSIBlockVolume` | `true` | stable | 1.18 | 1.21 |
| `CSIDriverRegistry` | `false` | alpha | 1.12 | 1.13 |
| `CSIDriverRegistry` | `true` | beta | 1.14 | 1.17 |
| `CSIDriverRegistry` | `true` | stable | 1.18 | 1.21 |
| `CSIInlineVolume` | `false` | alpha | 1.15 | 1.15 |
| `CSIInlineVolume` | `true` | beta | 1.16 | 1.24 |
| `CSIInlineVolume` | `true` | stable | 1.25 | 1.26 |
| `CSIMigration` | `false` | alpha | 1.14 | 1.16 |
| `CSIMigration` | `true` | beta | 1.17 | 1.24 |
| `CSIMigration` | `true` | stable | 1.25 | 1.26 |
| `CSIMigrationAWS` | `false` | alpha | 1.14 | 1.16 |
| `CSIMigrationAWS` | `false` | beta | 1.17 | 1.22 |
| `CSIMigrationAWS` | `true` | beta | 1.23 | 1.24 |
| `CSIMigrationAWS` | `true` | stable | 1.25 | 1.26 |
| `CSIMigrationAWSComplete` | `false` | alpha | 1.17 | 1.20 |
| `CSIMigrationAWSComplete` | `<nil>` | deprecated | 1.21 | 1.21 |
| `CSIMigrationAzureDisk` | `false` | alpha | 1.15 | 1.18 |
| `CSIMigrationAzureDisk` | `false` | beta | 1.19 | 1.22 |
| `CSIMigrationAzureDisk` | `true` | beta | 1.23 | 1.23 |
| `CSIMigrationAzureDisk` | `true` | stable | 1.24 | 1.26 |
| `CSIMigrationAzureDiskComplete` | `false` | alpha | 1.17 | 1.20 |
| `CSIMigrationAzureDiskComplete` | `<nil>` | deprecated | 1.21 | 1.21 |
| `CSIMigrationAzureFile` | `false` | alpha | 1.15 | 1.20 |
| `CSIMigrationAzureFile` | `false` | beta | 1.21 | 1.23 |
| `CSIMigrationAzureFile` | `true` | beta | 1.24 | 1.25 |
| `CSIMigrationAzureFile` | `true` | stable | 1.26 | 1.29 |
| `CSIMigrationAzureFileComplete` | `false` | alpha | 1.17 | 1.20 |
| `CSIMigrationAzureFileComplete` | `<nil>` | deprecated | 1.21 | 1.21 |
| `CSIMigrationGCE` | `false` | alpha | 1.14 | 1.16 |
| `CSIMigrationGCE` | `false` | beta | 1.17 | 1.22 |
| `CSIMigrationGCE` | `true` | beta | 1.23 | 1.24 |
| `CSIMigrationGCE` | `true` | stable | 1.25 | 1.27 |
| `CSIMigrationGCEComplete` | `false` | alpha | 1.17 | 1.20 |
| `CSIMigrationGCEComplete` | `<nil>` | deprecated | 1.21 | 1.21 |
| `CSIMigrationOpenStack` | `false` | alpha | 1.14 | 1.17 |
| `CSIMigrationOpenStack` | `true` | beta | 1.18 | 1.23 |
| `CSIMigrationOpenStack` | `true` | stable | 1.24 | 1.25 |
| `CSIMigrationOpenStackComplete` | `false` | alpha | 1.17 | 1.20 |
| `CSIMigrationOpenStackComplete` | `<nil>` | deprecated | 1.21 | 1.21 |
| `CSIMigrationPortworx` | `false` | alpha | 1.23 | 1.24 |
| `CSIMigrationPortworx` | `false` | beta | 1.25 | 1.30 |
| `CSIMigrationPortworx` | `true` | beta | 1.31 | 1.32 |
| `CSIMigrationPortworx` | `true` | stable | 1.33 | 1.35 |
| `CSIMigrationRBD` | `false` | alpha | 1.23 | 1.27 |
| `CSIMigrationRBD` | `false` | deprecated | 1.28 | 1.30 |
| `CSIMigrationvSphere` | `false` | alpha | 1.18 | 1.18 |
| `CSIMigrationvSphere` | `false` | beta | 1.19 | 1.24 |
| `CSIMigrationvSphere` | `true` | beta | 1.25 | 1.25 |
| `CSIMigrationvSphere` | `true` | stable | 1.26 | 1.28 |
| `CSIMigrationvSphereComplete` | `false` | beta | 1.19 | 1.21 |
| `CSIMigrationvSphereComplete` | `<nil>` | deprecated | 1.22 | 1.22 |
| `CSINodeExpandSecret` | `false` | alpha | 1.25 | 1.26 |
| `CSINodeExpandSecret` | `true` | beta | 1.27 | 1.28 |
| `CSINodeExpandSecret` | `true` | stable | 1.29 | 1.30 |
| `CSINodeInfo` | `false` | alpha | 1.12 | 1.13 |
| `CSINodeInfo` | `true` | beta | 1.14 | 1.16 |
| `CSINodeInfo` | `true` | stable | 1.17 | 1.22 |
| `CSIPersistentVolume` | `false` | alpha | 1.9 | 1.9 |
| `CSIPersistentVolume` | `true` | beta | 1.10 | 1.12 |
| `CSIPersistentVolume` | `true` | stable | 1.13 | 1.16 |
| `CSIServiceAccountToken` | `false` | alpha | 1.20 | 1.20 |
| `CSIServiceAccountToken` | `true` | beta | 1.21 | 1.21 |
| `CSIServiceAccountToken` | `true` | stable | 1.22 | 1.24 |
| `CSIStorageCapacity` | `false` | alpha | 1.19 | 1.20 |
| `CSIStorageCapacity` | `true` | beta | 1.21 | 1.23 |
| `CSIStorageCapacity` | `true` | stable | 1.24 | 1.27 |
| `CSIVolumeFSGroupPolicy` | `false` | alpha | 1.19 | 1.19 |
| `CSIVolumeFSGroupPolicy` | `true` | beta | 1.20 | 1.22 |
| `CSIVolumeFSGroupPolicy` | `true` | stable | 1.23 | 1.25 |
| `CSRDuration` | `true` | beta | 1.22 | 1.23 |
| `CSRDuration` | `true` | stable | 1.24 | 1.25 |
| `CloudDualStackNodeIPs` | `false` | alpha | 1.27 | 1.28 |
| `CloudDualStackNodeIPs` | `true` | beta | 1.29 | 1.29 |
| `CloudDualStackNodeIPs` | `true` | stable | 1.30 | 1.31 |
| `ComponentSLIs` | `false` | alpha | 1.26 | 1.26 |
| `ComponentSLIs` | `true` | beta | 1.27 | 1.31 |
| `ComponentSLIs` | `true` | stable | 1.32 | 1.34 |
| `ConfigurableFSGroupPolicy` | `false` | alpha | 1.18 | 1.19 |
| `ConfigurableFSGroupPolicy` | `true` | beta | 1.20 | 1.22 |
| `ConfigurableFSGroupPolicy` | `true` | stable | 1.23 | 1.25 |
| `ConsistentHTTPGetHandlers` | `true` | stable | 1.25 | 1.30 |
| `ControllerManagerLeaderMigration` | `false` | alpha | 1.21 | 1.21 |
| `ControllerManagerLeaderMigration` | `true` | beta | 1.22 | 1.23 |
| `ControllerManagerLeaderMigration` | `true` | stable | 1.24 | 1.26 |
| `CronJobControllerV2` | `false` | alpha | 1.20 | 1.20 |
| `CronJobControllerV2` | `true` | beta | 1.21 | 1.21 |
| `CronJobControllerV2` | `true` | stable | 1.22 | 1.23 |
| `CronJobTimeZone` | `false` | alpha | 1.24 | 1.24 |
| `CronJobTimeZone` | `true` | beta | 1.25 | 1.26 |
| `CronJobTimeZone` | `true` | stable | 1.27 | 1.28 |
| `CustomPodDNS` | `false` | alpha | 1.9 | 1.9 |
| `CustomPodDNS` | `true` | beta | 1.10 | 1.13 |
| `CustomPodDNS` | `true` | stable | 1.14 | 1.16 |
| `CustomResourceDefaulting` | `false` | alpha | 1.15 | 1.15 |
| `CustomResourceDefaulting` | `true` | beta | 1.16 | 1.16 |
| `CustomResourceDefaulting` | `true` | stable | 1.17 | 1.18 |
| `CustomResourcePublishOpenAPI` | `false` | alpha | 1.14 | 1.14 |
| `CustomResourcePublishOpenAPI` | `true` | beta | 1.15 | 1.15 |
| `CustomResourcePublishOpenAPI` | `true` | stable | 1.16 | 1.18 |
| `CustomResourceSubresources` | `false` | alpha | 1.10 | 1.10 |
| `CustomResourceSubresources` | `true` | beta | 1.11 | 1.15 |
| `CustomResourceSubresources` | `true` | stable | 1.16 | 1.18 |
| `CustomResourceValidation` | `false` | alpha | 1.8 | 1.8 |
| `CustomResourceValidation` | `true` | beta | 1.9 | 1.15 |
| `CustomResourceValidation` | `true` | stable | 1.16 | 1.18 |
| `CustomResourceValidationExpressions` | `false` | alpha | 1.23 | 1.24 |
| `CustomResourceValidationExpressions` | `true` | beta | 1.25 | 1.28 |
| `CustomResourceValidationExpressions` | `true` | stable | 1.29 | 1.30 |
| `CustomResourceWebhookConversion` | `false` | alpha | 1.13 | 1.14 |
| `CustomResourceWebhookConversion` | `true` | beta | 1.15 | 1.15 |
| `CustomResourceWebhookConversion` | `true` | stable | 1.16 | 1.18 |
| `DRAControlPlaneController` | `false` | alpha | 1.26 | 1.31 |
| `DaemonSetUpdateSurge` | `false` | alpha | 1.21 | 1.21 |
| `DaemonSetUpdateSurge` | `true` | beta | 1.22 | 1.24 |
| `DaemonSetUpdateSurge` | `true` | stable | 1.25 | 1.26 |
| `DefaultHostNetworkHostPortsInPodTemplates` | `false` | deprecated | 1.28 | 1.30 |
| `DefaultPodTopologySpread` | `false` | alpha | 1.19 | 1.19 |
| `DefaultPodTopologySpread` | `true` | beta | 1.20 | 1.23 |
| `DefaultPodTopologySpread` | `true` | stable | 1.24 | 1.25 |
| `DelegateFSGroupToCSIDriver` | `false` | alpha | 1.22 | 1.22 |
| `DelegateFSGroupToCSIDriver` | `true` | beta | 1.23 | 1.25 |
| `DelegateFSGroupToCSIDriver` | `true` | stable | 1.26 | 1.27 |
| `DevicePluginCDIDevices` | `false` | alpha | 1.28 | 1.28 |
| `DevicePluginCDIDevices` | `true` | beta | 1.29 | 1.30 |
| `DevicePluginCDIDevices` | `true` | stable | 1.31 | 1.33 |
| `DevicePlugins` | `false` | alpha | 1.8 | 1.9 |
| `DevicePlugins` | `true` | beta | 1.10 | 1.25 |
| `DevicePlugins` | `true` | stable | 1.26 | 1.27 |
| `DisableAcceleratorUsageMetrics` | `false` | alpha | 1.19 | 1.19 |
| `DisableAcceleratorUsageMetrics` | `true` | beta | 1.20 | 1.24 |
| `DisableAcceleratorUsageMetrics` | `true` | stable | 1.25 | 1.27 |
| `DisableCloudProviders` | `false` | alpha | 1.22 | 1.28 |
| `DisableCloudProviders` | `true` | beta | 1.29 | 1.30 |
| `DisableCloudProviders` | `true` | stable | 1.31 | 1.32 |
| `DisableKubeletCloudCredentialProviders` | `false` | alpha | 1.23 | 1.28 |
| `DisableKubeletCloudCredentialProviders` | `true` | beta | 1.29 | 1.30 |
| `DisableKubeletCloudCredentialProviders` | `true` | stable | 1.31 | 1.32 |
| `DownwardAPIHugePages` | `false` | alpha | 1.20 | 1.20 |
| `DownwardAPIHugePages` | `false` | beta | 1.21 | 1.21 |
| `DownwardAPIHugePages` | `true` | beta | 1.22 | 1.26 |
| `DownwardAPIHugePages` | `true` | stable | 1.27 | 1.28 |
| `DryRun` | `false` | alpha | 1.12 | 1.12 |
| `DryRun` | `true` | beta | 1.13 | 1.18 |
| `DryRun` | `true` | stable | 1.19 | 1.27 |
| `DynamicAuditing` | `false` | alpha | 1.13 | 1.18 |
| `DynamicAuditing` | `<nil>` | deprecated | 1.19 | 1.19 |
| `DynamicKubeletConfig` | `false` | alpha | 1.4 | 1.10 |
| `DynamicKubeletConfig` | `true` | beta | 1.11 | 1.21 |
| `DynamicKubeletConfig` | `false` | deprecated | 1.22 | 1.25 |
| `DynamicProvisioningScheduling` | `false` | alpha | 1.11 | 1.11 |
| `DynamicProvisioningScheduling` | `<nil>` | deprecated | 1.12 |  |
| `DynamicVolumeProvisioning` | `true` | alpha | 1.3 | 1.7 |
| `DynamicVolumeProvisioning` | `true` | stable | 1.8 | 1.12 |
| `EfficientWatchResumption` | `false` | alpha | 1.20 | 1.20 |
| `EfficientWatchResumption` | `true` | beta | 1.21 | 1.23 |
| `EfficientWatchResumption` | `true` | stable | 1.24 | 1.32 |
| `EnableAggregatedDiscoveryTimeout` | `true` | deprecated | 1.16 | 1.17 |
| `EnableEquivalenceClassCache` | `false` | alpha | 1.8 | 1.12 |
| `EnableEquivalenceClassCache` | `<nil>` | deprecated | 1.13 | 1.23 |
| `EndpointSlice` | `false` | alpha | 1.16 | 1.16 |
| `EndpointSlice` | `false` | beta | 1.17 | 1.17 |
| `EndpointSlice` | `true` | beta | 1.18 | 1.20 |
| `EndpointSlice` | `true` | stable | 1.21 | 1.24 |
| `EndpointSliceNodeName` | `false` | alpha | 1.20 | 1.20 |
| `EndpointSliceNodeName` | `true` | stable | 1.21 | 1.24 |
| `EndpointSliceProxying` | `false` | alpha | 1.18 | 1.18 |
| `EndpointSliceProxying` | `true` | beta | 1.19 | 1.21 |
| `EndpointSliceProxying` | `true` | stable | 1.22 | 1.24 |
| `EndpointSliceTerminatingCondition` | `false` | alpha | 1.20 | 1.21 |
| `EndpointSliceTerminatingCondition` | `true` | beta | 1.22 | 1.25 |
| `EndpointSliceTerminatingCondition` | `true` | stable | 1.26 | 1.27 |
| `EphemeralContainers` | `false` | alpha | 1.16 | 1.22 |
| `EphemeralContainers` | `true` | beta | 1.23 | 1.24 |
| `EphemeralContainers` | `true` | stable | 1.25 | 1.26 |
| `EvenPodsSpread` | `false` | alpha | 1.16 | 1.17 |
| `EvenPodsSpread` | `true` | beta | 1.18 | 1.18 |
| `EvenPodsSpread` | `true` | stable | 1.19 | 1.21 |
| `ExpandCSIVolumes` | `false` | alpha | 1.14 | 1.15 |
| `ExpandCSIVolumes` | `true` | beta | 1.16 | 1.23 |
| `ExpandCSIVolumes` | `true` | stable | 1.24 | 1.26 |
| `ExpandInUsePersistentVolumes` | `false` | alpha | 1.11 | 1.14 |
| `ExpandInUsePersistentVolumes` | `true` | beta | 1.15 | 1.23 |
| `ExpandInUsePersistentVolumes` | `true` | stable | 1.24 | 1.26 |
| `ExpandPersistentVolumes` | `false` | alpha | 1.8 | 1.10 |
| `ExpandPersistentVolumes` | `true` | beta | 1.11 | 1.23 |
| `ExpandPersistentVolumes` | `true` | stable | 1.24 | 1.26 |
| `ExpandedDNSConfig` | `false` | alpha | 1.22 | 1.25 |
| `ExpandedDNSConfig` | `true` | beta | 1.26 | 1.27 |
| `ExpandedDNSConfig` | `true` | stable | 1.28 | 1.29 |
| `ExperimentalCriticalPodAnnotation` | `false` | alpha | 1.5 | 1.12 |
| `ExperimentalCriticalPodAnnotation` | `false` | deprecated | 1.13 | 1.16 |
| `ExperimentalHostUserNamespaceDefaulting` | `false` | beta | 1.5 | 1.27 |
| `ExperimentalHostUserNamespaceDefaulting` | `false` | deprecated | 1.28 | 1.29 |
| `ExternalPolicyForExternalIP` | `true` | stable | 1.18 | 1.22 |
| `GCERegionalPersistentDisk` | `true` | beta | 1.10 | 1.12 |
| `GCERegionalPersistentDisk` | `true` | stable | 1.13 | 1.16 |
| `GRPCContainerProbe` | `false` | alpha | 1.23 | 1.23 |
| `GRPCContainerProbe` | `true` | beta | 1.24 | 1.26 |
| `GRPCContainerProbe` | `true` | stable | 1.27 | 1.28 |
| `GangScheduling` | `false` | alpha | 1.35 | 1.36 |
| `GenericEphemeralVolume` | `false` | alpha | 1.19 | 1.20 |
| `GenericEphemeralVolume` | `true` | beta | 1.21 | 1.22 |
| `GenericEphemeralVolume` | `true` | stable | 1.23 | 1.24 |
| `HPAContainerMetrics` | `false` | alpha | 1.20 | 1.26 |
| `HPAContainerMetrics` | `true` | beta | 1.27 | 1.29 |
| `HPAContainerMetrics` | `true` | stable | 1.30 | 1.31 |
| `HonorPVReclaimPolicy` | `false` | alpha | 1.23 | 1.30 |
| `HonorPVReclaimPolicy` | `true` | beta | 1.31 | 1.32 |
| `HonorPVReclaimPolicy` | `true` | stable | 1.33 | 1.35 |
| `HugePageStorageMediumSize` | `false` | alpha | 1.18 | 1.18 |
| `HugePageStorageMediumSize` | `true` | beta | 1.19 | 1.21 |
| `HugePageStorageMediumSize` | `true` | stable | 1.22 | 1.24 |
| `HugePages` | `false` | alpha | 1.8 | 1.9 |
| `HugePages` | `true` | beta | 1.10 | 1.13 |
| `HugePages` | `true` | stable | 1.14 | 1.16 |
| `HyperVContainer` | `false` | alpha | 1.10 | 1.19 |
| `HyperVContainer` | `false` | deprecated | 1.20 | 1.20 |
| `IPTablesOwnershipCleanup` | `false` | alpha | 1.25 | 1.26 |
| `IPTablesOwnershipCleanup` | `true` | beta | 1.27 | 1.27 |
| `IPTablesOwnershipCleanup` | `true` | stable | 1.28 | 1.29 |
| `IPv6DualStack` | `false` | alpha | 1.15 | 1.20 |
| `IPv6DualStack` | `true` | beta | 1.21 | 1.22 |
| `IPv6DualStack` | `true` | stable | 1.23 | 1.24 |
| `IdentifyPodOS` | `false` | alpha | 1.23 | 1.23 |
| `IdentifyPodOS` | `true` | beta | 1.24 | 1.24 |
| `IdentifyPodOS` | `true` | stable | 1.25 | 1.26 |
| `ImmutableEphemeralVolumes` | `false` | alpha | 1.18 | 1.18 |
| `ImmutableEphemeralVolumes` | `true` | beta | 1.19 | 1.20 |
| `ImmutableEphemeralVolumes` | `true` | stable | 1.21 | 1.24 |
| `InTreePluginAWSUnregister` | `false` | alpha | 1.21 | 1.30 |
| `InTreePluginAzureDiskUnregister` | `false` | alpha | 1.21 | 1.30 |
| `InTreePluginAzureFileUnregister` | `false` | alpha | 1.21 | 1.30 |
| `InTreePluginGCEUnregister` | `false` | alpha | 1.21 | 1.30 |
| `InTreePluginOpenStackUnregister` | `false` | alpha | 1.21 | 1.30 |
| `InTreePluginPortworxUnregister` | `false` | alpha | 1.23 | 1.35 |
| `InTreePluginRBDUnregister` | `false` | alpha | 1.23 | 1.27 |
| `InTreePluginRBDUnregister` | `false` | deprecated | 1.28 | 1.30 |
| `InTreePluginvSphereUnregister` | `false` | alpha | 1.21 | 1.30 |
| `IndexedJob` | `false` | alpha | 1.21 | 1.21 |
| `IndexedJob` | `true` | beta | 1.22 | 1.23 |
| `IndexedJob` | `true` | stable | 1.24 | 1.25 |
| `IngressClassNamespacedParams` | `false` | alpha | 1.21 | 1.21 |
| `IngressClassNamespacedParams` | `true` | beta | 1.22 | 1.22 |
| `IngressClassNamespacedParams` | `true` | stable | 1.23 | 1.24 |
| `Initializers` | `false` | alpha | 1.7 | 1.13 |
| `Initializers` | `<nil>` | deprecated | 1.14 | 1.14 |
| `JobMutableNodeSchedulingDirectives` | `true` | beta | 1.23 | 1.26 |
| `JobMutableNodeSchedulingDirectives` | `true` | stable | 1.27 | 1.28 |
| `JobPodFailurePolicy` | `false` | alpha | 1.25 | 1.25 |
| `JobPodFailurePolicy` | `true` | beta | 1.26 | 1.30 |
| `JobPodFailurePolicy` | `true` | stable | 1.31 | 1.32 |
| `JobReadyPods` | `false` | alpha | 1.23 | 1.23 |
| `JobReadyPods` | `true` | beta | 1.24 | 1.28 |
| `JobReadyPods` | `true` | stable | 1.29 | 1.30 |
| `JobTrackingWithFinalizers` | `false` | alpha | 1.22 | 1.22 |
| `JobTrackingWithFinalizers` | `false` | beta | 1.23 | 1.24 |
| `JobTrackingWithFinalizers` | `true` | beta | 1.25 | 1.25 |
| `JobTrackingWithFinalizers` | `true` | stable | 1.26 | 1.28 |
| `KMSv2` | `false` | alpha | 1.25 | 1.26 |
| `KMSv2` | `true` | beta | 1.27 | 1.28 |
| `KMSv2` | `true` | stable | 1.29 | 1.31 |
| `KMSv2KDF` | `false` | beta | 1.28 | 1.28 |
| `KMSv2KDF` | `true` | stable | 1.29 | 1.31 |
| `KubeProxyDrainingTerminatingNodes` | `false` | alpha | 1.28 | 1.30 |
| `KubeProxyDrainingTerminatingNodes` | `true` | beta | 1.30 | 1.30 |
| `KubeProxyDrainingTerminatingNodes` | `true` | stable | 1.31 | 1.32 |
| `KubeletConfigFile` | `false` | alpha | 1.8 | 1.9 |
| `KubeletConfigFile` | `<nil>` | deprecated | 1.10 | 1.10 |
| `KubeletCredentialProviders` | `false` | alpha | 1.20 | 1.23 |
| `KubeletCredentialProviders` | `true` | beta | 1.24 | 1.25 |
| `KubeletCredentialProviders` | `true` | stable | 1.26 | 1.28 |
| `KubeletPluginsWatcher` | `false` | alpha | 1.11 | 1.11 |
| `KubeletPluginsWatcher` | `true` | beta | 1.12 | 1.12 |
| `KubeletPluginsWatcher` | `true` | stable | 1.13 | 1.16 |
| `KubeletPodResources` | `false` | alpha | 1.13 | 1.14 |
| `KubeletPodResources` | `true` | beta | 1.15 | 1.27 |
| `KubeletPodResources` | `true` | stable | 1.28 | 1.29 |
| `KubeletPodResourcesGetAllocatable` | `false` | alpha | 1.21 | 1.22 |
| `KubeletPodResourcesGetAllocatable` | `true` | beta | 1.23 | 1.27 |
| `KubeletPodResourcesGetAllocatable` | `true` | stable | 1.28 | 1.29 |
| `LegacyNodeRoleBehavior` | `false` | alpha | 1.16 | 1.18 |
| `LegacyNodeRoleBehavior` | `true` | beta | 1.19 | 1.20 |
| `LegacyNodeRoleBehavior` | `false` | stable | 1.21 | 1.22 |
| `LegacyServiceAccountTokenCleanUp` | `false` | alpha | 1.28 | 1.28 |
| `LegacyServiceAccountTokenCleanUp` | `true` | beta | 1.29 | 1.29 |
| `LegacyServiceAccountTokenCleanUp` | `true` | stable | 1.30 | 1.31 |
| `LegacyServiceAccountTokenNoAutoGeneration` | `true` | beta | 1.24 | 1.25 |
| `LegacyServiceAccountTokenNoAutoGeneration` | `true` | stable | 1.26 | 1.28 |
| `LegacyServiceAccountTokenTracking` | `false` | alpha | 1.26 | 1.26 |
| `LegacyServiceAccountTokenTracking` | `true` | beta | 1.27 | 1.27 |
| `LegacyServiceAccountTokenTracking` | `true` | stable | 1.28 | 1.29 |
| `LoadBalancerIPMode` | `false` | alpha | 1.29 | 1.30 |
| `LoadBalancerIPMode` | `true` | beta | 1.30 | 1.31 |
| `LoadBalancerIPMode` | `true` | stable | 1.32 | 1.34 |
| `LocalStorageCapacityIsolation` | `false` | alpha | 1.7 | 1.9 |
| `LocalStorageCapacityIsolation` | `true` | beta | 1.10 | 1.24 |
| `LocalStorageCapacityIsolation` | `true` | stable | 1.25 | 1.26 |
| `MinDomainsInPodTopologySpread` | `false` | alpha | 1.24 | 1.24 |
| `MinDomainsInPodTopologySpread` | `false` | beta | 1.25 | 1.26 |
| `MinDomainsInPodTopologySpread` | `true` | beta | 1.27 | 1.29 |
| `MinDomainsInPodTopologySpread` | `true` | stable | 1.30 | 1.31 |
| `MinimizeIPTablesRestore` | `false` | alpha | 1.26 | 1.26 |
| `MinimizeIPTablesRestore` | `true` | beta | 1.27 | 1.27 |
| `MinimizeIPTablesRestore` | `true` | stable | 1.28 | 1.29 |
| `MixedProtocolLBService` | `false` | alpha | 1.20 | 1.23 |
| `MixedProtocolLBService` | `true` | beta | 1.24 | 1.25 |
| `MixedProtocolLBService` | `true` | stable | 1.26 | 1.27 |
| `MountContainers` | `false` | alpha | 1.9 | 1.16 |
| `MountContainers` | `false` | deprecated | 1.17 | 1.17 |
| `MountPropagation` | `false` | alpha | 1.8 | 1.9 |
| `MountPropagation` | `true` | beta | 1.10 | 1.11 |
| `MountPropagation` | `true` | stable | 1.12 | 1.14 |
| `MultiCIDRRangeAllocator` | `false` | alpha | 1.25 | 1.28 |
| `NamespaceDefaultLabelName` | `true` | beta | 1.21 | 1.21 |
| `NamespaceDefaultLabelName` | `true` | stable | 1.22 | 1.23 |
| `NetworkPolicyEndPort` | `false` | alpha | 1.21 | 1.21 |
| `NetworkPolicyEndPort` | `true` | beta | 1.22 | 1.24 |
| `NetworkPolicyEndPort` | `true` | stable | 1.25 | 1.26 |
| `NetworkPolicyStatus` | `false` | alpha | 1.24 | 1.27 |
| `NewVolumeManagerReconstruction` | `false` | alpha | 1.25 | 1.26 |
| `NewVolumeManagerReconstruction` | `true` | beta | 1.27 | 1.29 |
| `NewVolumeManagerReconstruction` | `true` | stable | 1.30 | 1.31 |
| `NodeDisruptionExclusion` | `false` | alpha | 1.16 | 1.18 |
| `NodeDisruptionExclusion` | `true` | beta | 1.19 | 1.20 |
| `NodeDisruptionExclusion` | `true` | stable | 1.21 | 1.22 |
| `NodeLease` | `false` | alpha | 1.12 | 1.13 |
| `NodeLease` | `true` | beta | 1.14 | 1.16 |
| `NodeLease` | `true` | stable | 1.17 | 1.23 |
| `NodeOutOfServiceVolumeDetach` | `false` | alpha | 1.24 | 1.25 |
| `NodeOutOfServiceVolumeDetach` | `true` | beta | 1.26 | 1.27 |
| `NodeOutOfServiceVolumeDetach` | `true` | stable | 1.28 | 1.31 |
| `NonPreemptingPriority` | `false` | alpha | 1.15 | 1.18 |
| `NonPreemptingPriority` | `true` | beta | 1.19 | 1.23 |
| `NonPreemptingPriority` | `true` | stable | 1.24 | 1.25 |
| `OpenAPIV3` | `false` | alpha | 1.23 | 1.23 |
| `OpenAPIV3` | `true` | beta | 1.24 | 1.26 |
| `OpenAPIV3` | `true` | stable | 1.27 | 1.28 |
| `PDBUnhealthyPodEvictionPolicy` | `false` | alpha | 1.26 | 1.26 |
| `PDBUnhealthyPodEvictionPolicy` | `true` | beta | 1.27 | 1.30 |
| `PDBUnhealthyPodEvictionPolicy` | `true` | stable | 1.31 | 1.32 |
| `PVCProtection` | `false` | alpha | 1.9 | 1.9 |
| `PVCProtection` | `<nil>` | deprecated | 1.10 | 1.10 |
| `PersistentLocalVolumes` | `false` | alpha | 1.7 | 1.9 |
| `PersistentLocalVolumes` | `true` | beta | 1.10 | 1.13 |
| `PersistentLocalVolumes` | `true` | stable | 1.14 | 1.16 |
| `PersistentVolumeLastPhaseTransitionTime` | `false` | alpha | 1.28 | 1.28 |
| `PersistentVolumeLastPhaseTransitionTime` | `true` | beta | 1.29 | 1.30 |
| `PersistentVolumeLastPhaseTransitionTime` | `true` | stable | 1.31 | 1.32 |
| `PodAffinityNamespaceSelector` | `false` | alpha | 1.21 | 1.21 |
| `PodAffinityNamespaceSelector` | `true` | beta | 1.22 | 1.23 |
| `PodAffinityNamespaceSelector` | `true` | stable | 1.24 | 1.25 |
| `PodDisruptionBudget` | `false` | alpha | 1.3 | 1.4 |
| `PodDisruptionBudget` | `true` | beta | 1.5 | 1.20 |
| `PodDisruptionBudget` | `true` | stable | 1.21 | 1.25 |
| `PodDisruptionConditions` | `false` | alpha | 1.25 | 1.25 |
| `PodDisruptionConditions` | `true` | beta | 1.26 | 1.30 |
| `PodDisruptionConditions` | `true` | stable | 1.31 | 1.33 |
| `PodHasNetworkCondition` | `false` | alpha | 1.25 | 1.27 |
| `PodHostIPs` | `false` | alpha | 1.28 | 1.28 |
| `PodHostIPs` | `true` | beta | 1.29 | 1.30 |
| `PodHostIPs` | `true` | stable | 1.30 | 1.31 |
| `PodOverhead` | `false` | alpha | 1.16 | 1.17 |
| `PodOverhead` | `true` | beta | 1.18 | 1.23 |
| `PodOverhead` | `true` | stable | 1.24 | 1.25 |
| `PodPriority` | `false` | alpha | 1.8 | 1.10 |
| `PodPriority` | `true` | beta | 1.11 | 1.13 |
| `PodPriority` | `true` | stable | 1.14 | 1.18 |
| `PodReadinessGates` | `false` | alpha | 1.11 | 1.11 |
| `PodReadinessGates` | `true` | beta | 1.12 | 1.13 |
| `PodReadinessGates` | `true` | stable | 1.14 | 1.16 |
| `PodSecurity` | `false` | alpha | 1.22 | 1.22 |
| `PodSecurity` | `true` | beta | 1.23 | 1.24 |
| `PodSecurity` | `true` | stable | 1.25 | 1.27 |
| `PodShareProcessNamespace` | `false` | alpha | 1.10 | 1.11 |
| `PodShareProcessNamespace` | `true` | beta | 1.12 | 1.16 |
| `PodShareProcessNamespace` | `true` | stable | 1.17 | 1.19 |
| `PreferNominatedNode` | `false` | alpha | 1.21 | 1.21 |
| `PreferNominatedNode` | `true` | beta | 1.22 | 1.23 |
| `PreferNominatedNode` | `true` | stable | 1.24 | 1.25 |
| `ProbeTerminationGracePeriod` | `false` | alpha | 1.21 | 1.21 |
| `ProbeTerminationGracePeriod` | `false` | beta | 1.22 | 1.24 |
| `ProbeTerminationGracePeriod` | `true` | beta | 1.25 | 1.27 |
| `ProbeTerminationGracePeriod` | `true` | stable | 1.28 | 1.28 |
| `ProxyTerminatingEndpoints` | `false` | alpha | 1.22 | 1.25 |
| `ProxyTerminatingEndpoints` | `true` | beta | 1.26 | 1.27 |
| `ProxyTerminatingEndpoints` | `true` | stable | 1.28 | 1.29 |
| `ReadOnlyAPIDataVolumes` | `true` | beta | 1.8 | 1.9 |
| `ReadOnlyAPIDataVolumes` | `<nil>` | stable | 1.10 | 1.10 |
| `ReadWriteOncePod` | `false` | alpha | 1.22 | 1.26 |
| `ReadWriteOncePod` | `true` | beta | 1.27 | 1.28 |
| `ReadWriteOncePod` | `true` | stable | 1.29 | 1.30 |
| `RelaxedDNSSearchValidation` | `false` | alpha | 1.32 | 1.32 |
| `RelaxedDNSSearchValidation` | `true` | beta | 1.33 | 1.33 |
| `RelaxedDNSSearchValidation` | `true` | stable | 1.34 | 1.36 |
| `RemainingItemCount` | `false` | alpha | 1.15 | 1.15 |
| `RemainingItemCount` | `true` | beta | 1.16 | 1.28 |
| `RemainingItemCount` | `true` | stable | 1.29 | 1.32 |
| `RemoveSelfLink` | `false` | alpha | 1.16 | 1.19 |
| `RemoveSelfLink` | `true` | beta | 1.20 | 1.23 |
| `RemoveSelfLink` | `true` | stable | 1.24 | 1.29 |
| `RequestManagement` | `false` | alpha | 1.15 | 1.16 |
| `RequestManagement` | `<nil>` | deprecated | 1.17 | 1.17 |
| `ResourceLimitsPriorityFunction` | `false` | alpha | 1.9 | 1.18 |
| `ResourceLimitsPriorityFunction` | `<nil>` | deprecated | 1.19 | 1.19 |
| `ResourceQuotaScopeSelectors` | `false` | alpha | 1.11 | 1.11 |
| `ResourceQuotaScopeSelectors` | `true` | beta | 1.12 | 1.16 |
| `ResourceQuotaScopeSelectors` | `true` | stable | 1.17 | 1.18 |
| `RetroactiveDefaultStorageClass` | `false` | alpha | 1.25 | 1.25 |
| `RetroactiveDefaultStorageClass` | `true` | beta | 1.26 | 1.27 |
| `RetroactiveDefaultStorageClass` | `true` | stable | 1.28 | 1.28 |
| `RootCAConfigMap` | `false` | alpha | 1.13 | 1.19 |
| `RootCAConfigMap` | `true` | beta | 1.20 | 1.20 |
| `RootCAConfigMap` | `true` | stable | 1.21 | 1.22 |
| `RotateKubeletClientCertificate` | `true` | beta | 1.8 | 1.18 |
| `RotateKubeletClientCertificate` | `true` | stable | 1.19 | 1.21 |
| `RunAsGroup` | `true` | beta | 1.14 | 1.20 |
| `RunAsGroup` | `true` | stable | 1.21 | 1.22 |
| `RuntimeClass` | `false` | alpha | 1.12 | 1.13 |
| `RuntimeClass` | `true` | beta | 1.14 | 1.19 |
| `RuntimeClass` | `true` | stable | 1.20 | 1.24 |
| `SCTPSupport` | `false` | alpha | 1.12 | 1.18 |
| `SCTPSupport` | `true` | beta | 1.19 | 1.19 |
| `SCTPSupport` | `true` | stable | 1.20 | 1.22 |
| `ScheduleDaemonSetPods` | `false` | alpha | 1.11 | 1.11 |
| `ScheduleDaemonSetPods` | `true` | beta | 1.12 | 1.16 |
| `ScheduleDaemonSetPods` | `true` | stable | 1.17 | 1.18 |
| `SeccompDefault` | `false` | alpha | 1.22 | 1.24 |
| `SeccompDefault` | `true` | beta | 1.25 | 1.26 |
| `SeccompDefault` | `true` | stable | 1.27 | 1.28 |
| `SecurityContextDeny` | `false` | alpha | 1.27 | 1.29 |
| `SelectorIndex` | `false` | alpha | 1.18 | 1.18 |
| `SelectorIndex` | `true` | beta | 1.19 | 1.19 |
| `SelectorIndex` | `true` | stable | 1.20 | 1.25 |
| `ServerSideApply` | `false` | alpha | 1.14 | 1.15 |
| `ServerSideApply` | `true` | beta | 1.16 | 1.21 |
| `ServerSideApply` | `true` | stable | 1.22 | 1.31 |
| `ServerSideFieldValidation` | `false` | alpha | 1.23 | 1.24 |
| `ServerSideFieldValidation` | `true` | beta | 1.25 | 1.26 |
| `ServerSideFieldValidation` | `true` | stable | 1.27 | 1.31 |
| `ServiceAccountIssuerDiscovery` | `false` | alpha | 1.18 | 1.19 |
| `ServiceAccountIssuerDiscovery` | `true` | beta | 1.20 | 1.20 |
| `ServiceAccountIssuerDiscovery` | `true` | stable | 1.21 | 1.23 |
| `ServiceAppProtocol` | `false` | alpha | 1.18 | 1.18 |
| `ServiceAppProtocol` | `true` | beta | 1.19 | 1.19 |
| `ServiceAppProtocol` | `true` | stable | 1.20 | 1.22 |
| `ServiceIPStaticSubrange` | `false` | alpha | 1.24 | 1.24 |
| `ServiceIPStaticSubrange` | `true` | beta | 1.25 | 1.25 |
| `ServiceIPStaticSubrange` | `true` | stable | 1.26 | 1.27 |
| `ServiceInternalTrafficPolicy` | `false` | alpha | 1.21 | 1.21 |
| `ServiceInternalTrafficPolicy` | `true` | beta | 1.22 | 1.25 |
| `ServiceInternalTrafficPolicy` | `true` | stable | 1.26 | 1.27 |
| `ServiceLBNodePortControl` | `false` | alpha | 1.20 | 1.21 |
| `ServiceLBNodePortControl` | `true` | beta | 1.22 | 1.23 |
| `ServiceLBNodePortControl` | `true` | stable | 1.24 | 1.25 |
| `ServiceLoadBalancerClass` | `false` | alpha | 1.21 | 1.21 |
| `ServiceLoadBalancerClass` | `true` | beta | 1.22 | 1.23 |
| `ServiceLoadBalancerClass` | `true` | stable | 1.24 | 1.25 |
| `ServiceLoadBalancerFinalizer` | `false` | alpha | 1.15 | 1.15 |
| `ServiceLoadBalancerFinalizer` | `true` | beta | 1.16 | 1.16 |
| `ServiceLoadBalancerFinalizer` | `true` | stable | 1.17 | 1.20 |
| `ServiceNodeExclusion` | `false` | alpha | 1.8 | 1.18 |
| `ServiceNodeExclusion` | `true` | beta | 1.19 | 1.20 |
| `ServiceNodeExclusion` | `true` | stable | 1.21 | 1.22 |
| `ServiceNodePortStaticSubrange` | `false` | alpha | 1.27 | 1.27 |
| `ServiceNodePortStaticSubrange` | `true` | beta | 1.28 | 1.28 |
| `ServiceNodePortStaticSubrange` | `true` | stable | 1.29 | 1.30 |
| `ServiceTopology` | `false` | alpha | 1.17 | 1.19 |
| `ServiceTopology` | `false` | deprecated | 1.20 | 1.22 |
| `SetHostnameAsFQDN` | `false` | alpha | 1.19 | 1.19 |
| `SetHostnameAsFQDN` | `true` | beta | 1.20 | 1.21 |
| `SetHostnameAsFQDN` | `true` | stable | 1.22 | 1.24 |
| `SizeMemoryBackedVolumes` | `false` | alpha | 1.20 | 1.21 |
| `SizeMemoryBackedVolumes` | `true` | beta | 1.22 | 1.31 |
| `SizeMemoryBackedVolumes` | `true` | stable | 1.32 | 1.34 |
| `SkipReadOnlyValidationGCE` | `false` | alpha | 1.28 | 1.28 |
| `SkipReadOnlyValidationGCE` | `true` | deprecated | 1.29 | 1.30 |
| `StableLoadBalancerNodeSet` | `true` | beta | 1.27 | 1.29 |
| `StableLoadBalancerNodeSet` | `true` | stable | 1.30 | 1.31 |
| `StartupProbe` | `false` | alpha | 1.16 | 1.17 |
| `StartupProbe` | `true` | beta | 1.18 | 1.19 |
| `StartupProbe` | `true` | stable | 1.20 | 1.23 |
| `StatefulSetMinReadySeconds` | `false` | alpha | 1.22 | 1.22 |
| `StatefulSetMinReadySeconds` | `true` | beta | 1.23 | 1.24 |
| `StatefulSetMinReadySeconds` | `true` | stable | 1.25 | 1.26 |
| `StorageObjectInUseProtection` | `true` | beta | 1.10 | 1.10 |
| `StorageObjectInUseProtection` | `true` | stable | 1.11 | 1.24 |
| `StreamingProxyRedirects` | `false` | beta | 1.5 | 1.5 |
| `StreamingProxyRedirects` | `true` | beta | 1.6 | 1.17 |
| `StreamingProxyRedirects` | `true` | deprecated | 1.18 | 1.21 |
| `StreamingProxyRedirects` | `false` | deprecated | 1.22 | 1.24 |
| `SupportIPVSProxyMode` | `false` | alpha | 1.8 | 1.8 |
| `SupportIPVSProxyMode` | `false` | beta | 1.9 | 1.9 |
| `SupportIPVSProxyMode` | `true` | beta | 1.10 | 1.10 |
| `SupportIPVSProxyMode` | `true` | stable | 1.11 | 1.20 |
| `SupportNodePidsLimit` | `false` | alpha | 1.14 | 1.14 |
| `SupportNodePidsLimit` | `true` | beta | 1.15 | 1.19 |
| `SupportNodePidsLimit` | `true` | stable | 1.20 | 1.23 |
| `SupportPodPidsLimit` | `false` | alpha | 1.10 | 1.13 |
| `SupportPodPidsLimit` | `true` | beta | 1.14 | 1.19 |
| `SupportPodPidsLimit` | `true` | stable | 1.20 | 1.23 |
| `SuspendJob` | `false` | alpha | 1.21 | 1.21 |
| `SuspendJob` | `true` | beta | 1.22 | 1.23 |
| `SuspendJob` | `true` | stable | 1.24 | 1.25 |
| `Sysctls` | `true` | beta | 1.11 | 1.20 |
| `Sysctls` | `true` | stable | 1.21 | 1.22 |
| `TTLAfterFinished` | `false` | alpha | 1.12 | 1.20 |
| `TTLAfterFinished` | `true` | beta | 1.21 | 1.22 |
| `TTLAfterFinished` | `true` | stable | 1.23 | 1.24 |
| `TaintBasedEvictions` | `false` | alpha | 1.6 | 1.12 |
| `TaintBasedEvictions` | `true` | beta | 1.13 | 1.17 |
| `TaintBasedEvictions` | `true` | stable | 1.18 | 1.20 |
| `TaintNodesByCondition` | `false` | alpha | 1.8 | 1.11 |
| `TaintNodesByCondition` | `true` | beta | 1.12 | 1.16 |
| `TaintNodesByCondition` | `true` | stable | 1.17 | 1.18 |
| `TokenRequest` | `false` | alpha | 1.10 | 1.11 |
| `TokenRequest` | `true` | beta | 1.12 | 1.19 |
| `TokenRequest` | `true` | stable | 1.20 | 1.21 |
| `TokenRequestProjection` | `false` | alpha | 1.11 | 1.11 |
| `TokenRequestProjection` | `true` | beta | 1.12 | 1.19 |
| `TokenRequestProjection` | `true` | stable | 1.20 | 1.21 |
| `TopologyAwareHints` | `false` | alpha | 1.21 | 1.22 |
| `TopologyAwareHints` | `false` | beta | 1.23 | 1.23 |
| `TopologyAwareHints` | `true` | beta | 1.24 | 1.32 |
| `TopologyAwareHints` | `true` | stable | 1.33 | 1.35 |
| `TopologyManager` | `false` | alpha | 1.16 | 1.17 |
| `TopologyManager` | `true` | beta | 1.18 | 1.26 |
| `TopologyManager` | `true` | stable | 1.27 | 1.28 |
| `UserNamespacesPodSecurityStandards` | `false` | alpha | 1.29 | 1.34 |
| `UserNamespacesStatelessPodsSupport` | `false` | alpha | 1.25 | 1.27 |
| `ValidateProxyRedirects` | `false` | alpha | 1.12 | 1.13 |
| `ValidateProxyRedirects` | `true` | beta | 1.14 | 1.21 |
| `ValidateProxyRedirects` | `true` | deprecated | 1.22 | 1.24 |
| `ValidatingAdmissionPolicy` | `false` | alpha | 1.26 | 1.27 |
| `ValidatingAdmissionPolicy` | `false` | beta | 1.28 | 1.29 |
| `ValidatingAdmissionPolicy` | `true` | stable | 1.30 | 1.31 |
| `VolumeCapacityPriority` | `false` | alpha | 1.21 | 1.32 |
| `VolumePVCDataSource` | `false` | alpha | 1.15 | 1.15 |
| `VolumePVCDataSource` | `true` | beta | 1.16 | 1.17 |
| `VolumePVCDataSource` | `true` | stable | 1.18 | 1.21 |
| `VolumeScheduling` | `false` | alpha | 1.9 | 1.9 |
| `VolumeScheduling` | `true` | beta | 1.10 | 1.12 |
| `VolumeScheduling` | `true` | stable | 1.13 | 1.16 |
| `VolumeSnapshotDataSource` | `false` | alpha | 1.12 | 1.16 |
| `VolumeSnapshotDataSource` | `true` | beta | 1.17 | 1.19 |
| `VolumeSnapshotDataSource` | `true` | stable | 1.20 | 1.22 |
| `VolumeSubpath` | `true` | stable | 1.10 | 1.24 |
| `VolumeSubpathEnvExpansion` | `false` | alpha | 1.14 | 1.14 |
| `VolumeSubpathEnvExpansion` | `true` | beta | 1.15 | 1.16 |
| `VolumeSubpathEnvExpansion` | `true` | stable | 1.17 | 1.24 |
| `WarningHeaders` | `true` | beta | 1.19 | 1.21 |
| `WarningHeaders` | `true` | stable | 1.22 | 1.24 |
| `WatchBookmark` | `false` | alpha | 1.15 | 1.15 |
| `WatchBookmark` | `true` | beta | 1.16 | 1.16 |
| `WatchBookmark` | `true` | stable | 1.17 | 1.32 |
| `WindowsEndpointSliceProxying` | `false` | alpha | 1.19 | 1.20 |
| `WindowsEndpointSliceProxying` | `true` | beta | 1.21 | 1.21 |
| `WindowsEndpointSliceProxying` | `true` | stable | 1.22 | 1.24 |
| `WindowsGMSA` | `false` | alpha | 1.14 | 1.15 |
| `WindowsGMSA` | `true` | beta | 1.16 | 1.17 |
| `WindowsGMSA` | `true` | stable | 1.18 | 1.18 |
| `WindowsHostProcessContainers` | `false` | alpha | 1.22 | 1.22 |
| `WindowsHostProcessContainers` | `true` | beta | 1.23 | 1.25 |
| `WindowsHostProcessContainers` | `true` | stable | 1.26 | 1.27 |
| `WindowsRunAsUserName` | `false` | alpha | 1.16 | 1.16 |
| `WindowsRunAsUserName` | `true` | beta | 1.17 | 1.17 |
| `WindowsRunAsUserName` | `true` | stable | 1.18 | 1.20 |
| `WorkloadAwarePreemption` | `false` | alpha | 1.36 | 1.36 |
| `ZeroLimitedNominalConcurrencyShares` | `false` | beta | 1.29 | 1.29 |
| `ZeroLimitedNominalConcurrencyShares` | `true` | stable | 1.30 | 1.31 |

## Descriptions for removed feature gates

* **`APIListChunking`**: Enable the API clients to retrieve (`LIST` or `GET`) resources from API server in chunks.
* **`APIPriorityAndFairness`**: Enable managing request concurrency with prioritization and fairness at each server. (Renamed from `RequestManagement`)
* **`APISelfSubjectReview`**: Activate the `SelfSubjectReview` API which allows users to see the requesting subject's authentication information. See [API access to authentication information for a client](/reference/access-authn-authz/authentication.md#self-subject-review) for more details.
* **`Accelerators`**: Provided an early form of plugin to enable Nvidia GPU support when using Docker Engine; no longer available. See [Device Plugins](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md) for an alternative.
* **`AdmissionWebhookMatchConditions`**: Enable [match conditions](/reference/access-authn-authz/extensible-admission-controllers.md#matching-requests-matchconditions) on mutating & validating admission webhooks.
* **`AdvancedAuditing`**: Enable [advanced auditing](/tasks/debug/debug-cluster/audit.md#advanced-audit)
* **`AffinityInAnnotations`**: Enable setting [Pod affinity or anti-affinity](/concepts/scheduling-eviction/assign-pod-node.md#affinity-and-anti-affinity).
* **`AggregatedDiscoveryEndpoint`**: Enable a single HTTP endpoint `/discovery/<version>` which supports native HTTP caching with ETags containing all APIResources known to the API server.
* **`AllowExtTrafficLocalEndpoints`**: Enable a service to route external requests to node local endpoints.
* **`AllowInsecureBackendProxy`**: Enable the users to skip TLS verification of kubelets on Pod log requests.
* **`AllowServiceLBStatusOnNonLB`**: Enables `.status.ingress.loadBalancer` to be set on Services of types other than `LoadBalancer`.
* **`AppArmor`**: Enable use of AppArmor mandatory access control for Pods running on Linux nodes. See [AppArmor Tutorial](https://kubernetes.io/docs/tutorials/security/apparmor/) for more details.
* **`AppArmorFields`**: Enable AppArmor related security context settings. For more information about AppArmor and Kubernetes, read the [AppArmor](/concepts/security/linux-kernel-security-constraints.md#apparmor) section within [security features in the Linux kernel](/concepts/security/linux-kernel-security-constraints.md#linux-security-features).
* **`AttachVolumeLimit`**: Enable volume plugins to report limits on number of volumes that can be attached to a node. See [dynamic volume limits](/concepts/storage/storage-limits.md#dynamic-volume-limits) for more details.
* **`BalanceAttachedNodeVolumes`**: Include volume count on node to be considered for balanced resource allocation while scheduling. A node which has closer CPU, memory utilization, and volume count is favored by the scheduler while making decisions.
* **`BlockVolume`**: Enable the definition and consumption of raw block devices in Pods. See [Raw Block Volume Support](/concepts/storage/persistent-volumes.md#raw-block-volume-support) for more details.
* **`BoundServiceAccountTokenVolume`**: Migrate ServiceAccount volumes to use a projected volume consisting of a ServiceAccountTokenVolumeProjection. Cluster admins can use metric `serviceaccount_stale_tokens_total` to monitor workloads that are depending on the extended tokens. If there are no such workloads, turn off extended tokens by starting `kube-apiserver` with flag `--service-account-extend-token-expiration=false`. Check [Bound Service Account Tokens](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/1205-bound-service-account-tokens/README.md) for more details.
* **`CPUManager`**: Enable container level CPU affinity support, see [CPU Management Policies](/tasks/administer-cluster/cpu-management-policies.md).
* **`CRIContainerLogRotation`**: Enable container log rotation for CRI container runtime. The default max size of a log file is 10MB and the default max number of log files allowed for a container is 5. These values can be configured in the kubelet config. See [logging at node level](/concepts/cluster-administration/logging.md#logging-at-the-node-level) for more details.
* **`CSIBlockVolume`**: Enable external CSI volume drivers to support block storage. See [`csi` raw block volume support](/concepts/storage/volumes.md#csi-raw-block-volume-support) for more details.
* **`CSIDriverRegistry`**: Enable all logic related to the CSIDriver API object in `csi.storage.k8s.io`.
* **`CSIInlineVolume`**: Enable CSI Inline volumes support for pods.
* **`CSIMigration`**: Enables shims and translation logic to route volume operations from in-tree plugins to corresponding pre-installed CSI plugins
* **`CSIMigrationAWS`**: Enables shims and translation logic to route volume operations from the AWS-EBS in-tree plugin to EBS CSI plugin. Supports falling back to in-tree EBS plugin for mount operations to nodes that have the feature disabled or that do not have EBS CSI plugin installed and configured. Does not support falling back for provision operations, for those the CSI plugin must be installed and configured.
* **`CSIMigrationAWSComplete`**: Stops registering the EBS in-tree plugin in kubelet and volume controllers and enables shims and translation logic to route volume operations from the AWS-EBS in-tree plugin to EBS CSI plugin. Requires CSIMigration and CSIMigrationAWS feature flags enabled and EBS CSI plugin installed and configured on all nodes in the cluster. This flag has been deprecated in favor of the `InTreePluginAWSUnregister` feature flag which prevents the registration of in-tree EBS plugin.
* **`CSIMigrationAzureDisk`**: Enables shims and translation logic to route volume operations from the Azure-Disk in-tree plugin to AzureDisk CSI plugin. Supports falling back to in-tree AzureDisk plugin for mount operations to nodes that have the feature disabled or that do not have AzureDisk CSI plugin installed and configured. Does not support falling back for provision operations, for those the CSI plugin must be installed and configured. Requires CSIMigration feature flag enabled.
* **`CSIMigrationAzureDiskComplete`**: Stops registering the Azure-Disk in-tree plugin in kubelet and volume controllers and enables shims and translation logic to route volume operations from the Azure-Disk in-tree plugin to AzureDisk CSI plugin. Requires CSIMigration and CSIMigrationAzureDisk feature flags enabled and AzureDisk CSI plugin installed and configured on all nodes in the cluster. This flag has been deprecated in favor of the `InTreePluginAzureDiskUnregister` feature flag which prevents the registration of in-tree AzureDisk plugin.
* **`CSIMigrationAzureFile`**: Enables shims and translation logic to route volume operations from the Azure-File in-tree plugin to AzureFile CSI plugin. Supports falling back to in-tree AzureFile plugin for mount operations to nodes that have the feature disabled or that do not have AzureFile CSI plugin installed and configured. Does not support falling back for provision operations, for those the CSI plugin must be installed and configured. Requires CSIMigration feature flag enabled.
* **`CSIMigrationAzureFileComplete`**: Stops registering the Azure-File in-tree plugin in kubelet and volume controllers and enables shims and translation logic to route volume operations from the Azure-File in-tree plugin to AzureFile CSI plugin. Requires CSIMigration and CSIMigrationAzureFile feature flags enabled and AzureFile CSI plugin installed and configured on all nodes in the cluster. This flag has been deprecated in favor of the `InTreePluginAzureFileUnregister` feature flag which prevents the registration of in-tree AzureFile plugin.
* **`CSIMigrationGCE`**: Enables shims and translation logic to route volume operations from the GCE-PD in-tree plugin to PD CSI plugin. Supports falling back to in-tree GCE plugin for mount operations to nodes that have the feature disabled or that do not have PD CSI plugin installed and configured. Does not support falling back for provision operations, for those the CSI plugin must be installed and configured. Requires CSIMigration feature flag enabled.
* **`CSIMigrationGCEComplete`**: Stops registering the GCE-PD in-tree plugin in kubelet and volume controllers and enables shims and translation logic to route volume operations from the GCE-PD in-tree plugin to PD CSI plugin. Requires CSIMigration and CSIMigrationGCE feature flags enabled and PD CSI plugin installed and configured on all nodes in the cluster. This flag has been deprecated in favor of the `InTreePluginGCEUnregister` feature flag which prevents the registration of in-tree GCE PD plugin.
* **`CSIMigrationOpenStack`**: Enables shims and translation logic to route volume operations from the Cinder in-tree plugin to Cinder CSI plugin. Supports falling back to in-tree Cinder plugin for mount operations to nodes that have the feature disabled or that do not have Cinder CSI plugin installed and configured. Does not support falling back for provision operations, for those the CSI plugin must be installed and configured. Requires CSIMigration feature flag enabled.
* **`CSIMigrationOpenStackComplete`**: Stops registering the Cinder in-tree plugin in kubelet and volume controllers and enables shims and translation logic to route volume operations from the Cinder in-tree plugin to Cinder CSI plugin. Requires CSIMigration and CSIMigrationOpenStack feature flags enabled and Cinder CSI plugin installed and configured on all nodes in the cluster. This flag has been deprecated in favor of the `InTreePluginOpenStackUnregister` feature flag which prevents the registration of in-tree openstack cinder plugin.
* **`CSIMigrationPortworx`**: Enables shims and translation logic to route volume operations from the Portworx in-tree plugin to Portworx CSI plugin. Requires Portworx CSI driver to be installed and configured in the cluster.
* **`CSIMigrationRBD`**: Enables shims and translation logic to route volume operations from the RBD in-tree plugin to Ceph RBD CSI plugin. Requires CSIMigration and csiMigrationRBD feature flags enabled and Ceph CSI plugin installed and configured in the cluster. This feature gate was deprecated in favor of the `InTreePluginRBDUnregister` feature gate, which prevents the registration of in-tree RBD plugin.
* **`CSIMigrationvSphere`**: Enables shims and translation logic to route volume operations from the vSphere in-tree plugin to vSphere CSI plugin. Supports falling back to in-tree vSphere plugin for mount operations to nodes that have the feature disabled or that do not have vSphere CSI plugin installed and configured. Does not support falling back for provision operations, for those the CSI plugin must be installed and configured. Requires CSIMigration feature flag enabled.
* **`CSIMigrationvSphereComplete`**: Stops registering the vSphere in-tree plugin in kubelet and volume controllers and enables shims and translation logic to route volume operations from the vSphere in-tree plugin to vSphere CSI plugin. Requires CSIMigration and CSIMigrationvSphere feature flags enabled and vSphere CSI plugin installed and configured on all nodes in the cluster. This flag has been deprecated in favor of the `InTreePluginvSphereUnregister` feature flag which prevents the registration of in-tree vsphere plugin.
* **`CSINodeExpandSecret`**: Enable passing secret authentication data to a CSI driver for use during a `NodeExpandVolume` CSI operation.
* **`CSINodeInfo`**: Enable all logic related to the CSINodeInfo API object in `csi.storage.k8s.io`.
* **`CSIPersistentVolume`**: Enable discovering and mounting volumes provisioned through a [CSI (Container Storage Interface)](https://git.k8s.io/design-proposals-archive/storage/container-storage-interface.md) compatible volume plugin.
* **`CSIServiceAccountToken`**: Enable CSI drivers to receive the pods' service account token that they mount volumes for. See [Token Requests](https://kubernetes-csi.github.io/docs/token-requests.html).
* **`CSIStorageCapacity`**: Enables CSI drivers to publish storage capacity information and the Kubernetes scheduler to use that information when scheduling pods. See [Storage Capacity](/concepts/storage/storage-capacity.md). Check the [`csi` volume type](/concepts/storage/volumes.md#csi) documentation for more details.
* **`CSIVolumeFSGroupPolicy`**: Allows CSIDrivers to use the `fsGroupPolicy` field. This field controls whether volumes created by a CSIDriver support volume ownership and permission modifications when these volumes are mounted.
* **`CSRDuration`**: Allows clients to request a duration for certificates issued via the Kubernetes CSR API.
* **`CloudDualStackNodeIPs`**: Enables dual-stack `kubelet --node-ip` with external cloud providers. See [Configure IPv4/IPv6 dual-stack](/concepts/services-networking/dual-stack.md#configure-ipv4-ipv6-dual-stack) for more details.
* **`ComponentSLIs`**: Enable the `/metrics/slis` endpoint on Kubernetes components like kubelet, kube-scheduler, kube-proxy, kube-controller-manager, cloud-controller-manager allowing you to scrape health check metrics.
* **`ConfigurableFSGroupPolicy`**: Allows user to configure volume permission change policy for fsGroups when mounting a volume in a Pod. See [Configure volume permission and ownership change policy for Pods](/tasks/configure-pod-container/security-context.md#configure-volume-permission-and-ownership-change-policy-for-pods) for more details.
* **`ConsistentHTTPGetHandlers`**: Normalize HTTP get URL and Header passing for lifecycle handlers with probers.
* **`ControllerManagerLeaderMigration`**: Enables Leader Migration for [kube-controller-manager](/tasks/administer-cluster/controller-manager-leader-migration.md#initial-leader-migration-configuration) and [cloud-controller-manager](/tasks/administer-cluster/controller-manager-leader-migration.md#deploy-cloud-controller-manager) which allows a cluster operator to live migrate controllers from the kube-controller-manager into an external controller-manager (e.g. the cloud-controller-manager) in an HA cluster without downtime.
* **`CronJobControllerV2`**: Use an alternative implementation of the [CronJob](/reference/glossary/cronjob.md) controller. Otherwise, version 1 of the same controller is selected.
* **`CronJobTimeZone`**: Allow the use of the `timeZone` optional field in [CronJobs](/concepts/workloads/controllers/cron-jobs.md)
* **`CustomPodDNS`**: Enable customizing the DNS settings for a Pod using its `dnsConfig` property. Check [Pod's DNS Config](/concepts/services-networking/dns-pod-service.md#pod-dns-config) for more details.
* **`CustomResourceDefaulting`**: Enable CRD support for default values in OpenAPI v3 validation schemas.
* **`CustomResourcePublishOpenAPI`**: Enables publishing of CRD OpenAPI specs.
* **`CustomResourceSubresources`**: Enable `/status` and `/scale` subresources on resources created from [CustomResourceDefinition](/concepts/extend-kubernetes/api-extension/custom-resources.md).
* **`CustomResourceValidation`**: Enable schema based validation on resources created from [CustomResourceDefinition](/concepts/extend-kubernetes/api-extension/custom-resources.md).
* **`CustomResourceValidationExpressions`**: Enable expression language validation in CRD which will validate customer resource based on validation rules written in the `x-kubernetes-validations` extension.
* **`CustomResourceWebhookConversion`**: Enable webhook-based conversion on resources created from [CustomResourceDefinition](/concepts/extend-kubernetes/api-extension/custom-resources.md).
* **`DRAControlPlaneController`**: Enables support for resources with custom parameters and a lifecycle that is independent of a Pod. Allocation of resources is handled by a resource driver's control plane controller.
* **`DaemonSetUpdateSurge`**: Enables the DaemonSet workloads to maintain availability during update per node. See [Perform a Rolling Update on a DaemonSet](/tasks/manage-daemon/update-daemon-set.md).
* **`DefaultHostNetworkHostPortsInPodTemplates`**: This feature gate controls the point at which a default value for `.spec.containers[*].ports[*].hostPort` is assigned, for Pods using `hostNetwork: true`. The default since Kubernetes v1.28 is to only set a default value in Pods. Enabling this means a default will be assigned even to the `.spec` of an embedded [PodTemplate](/concepts/workloads/pods.md#pod-templates) (for example, in a Deployment), which is the way that older releases of Kubernetes worked. You should migrate your code so that it does not rely on the legacy behavior.
* **`DefaultPodTopologySpread`**: Enables the use of `PodTopologySpread` scheduling plugin to do [default spreading](/concepts/scheduling-eviction/topology-spread-constraints.md#internal-default-constraints).
* **`DelegateFSGroupToCSIDriver`**: If supported by the CSI driver, delegates the role of applying `fsGroup` from a Pod's `securityContext` to the driver by passing `fsGroup` through the NodeStageVolume and NodePublishVolume CSI calls.
* **`DevicePluginCDIDevices`**: Enable support to CDI device IDs in the [Device Plugin](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md) API.
* **`DevicePlugins`**: Enable the [device-plugins](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md) based resource provisioning on nodes.
* **`DisableAcceleratorUsageMetrics`**: [Disable accelerator metrics collected by the kubelet](/concepts/cluster-administration/system-metrics.md#disable-accelerator-metrics).
* **`DisableCloudProviders`**: Enabling this feature gate deactivated functionality in `kube-apiserver`, `kube-controller-manager` and `kubelet` that related to the `--cloud-provider` command line argument. In Kubernetes v1.31 and later, the only valid values for `--cloud-provider` are the empty string (no cloud provider integration), or "external" (integration via a separate cloud-controller-manager).
* **`DisableKubeletCloudCredentialProviders`**: Enabling the feature gate deactivated the legacy in-tree functionality within the kubelet, that allowed the kubelet to authenticate to a cloud provider container registry for container image pulls.
* **`DownwardAPIHugePages`**: Enables usage of hugepages in [downward API](/tasks/inject-data-application/downward-api-volume-expose-pod-information.md).
* **`DryRun`**: Enable server-side [dry run](/reference/using-api/api-concepts.md#dry-run) requests so that validation, merging, and mutation can be tested without committing.
* **`DynamicAuditing`**: Used to enable dynamic auditing before v1.19.
* **`DynamicKubeletConfig`**: Enable the dynamic configuration of kubelet. The feature is no longer supported outside of supported skew policy. The feature gate was removed from kubelet in 1.24.
* **`DynamicProvisioningScheduling`**: Extend the default scheduler to be aware of volume topology and handle PV provisioning. This feature was superseded by the `VolumeScheduling` feature in v1.12.
* **`DynamicVolumeProvisioning`**: Enable the [dynamic provisioning](/concepts/storage/dynamic-provisioning.md) of persistent volumes to Pods.
* **`EfficientWatchResumption`**: Allows for storage-originated bookmark (progress notify) events to be delivered to the users. This is only applied to watch operations.
* **`EnableAggregatedDiscoveryTimeout`**: Enable the five second timeout on aggregated discovery calls.
* **`EnableEquivalenceClassCache`**: Enable the scheduler to cache equivalence of nodes when scheduling Pods.
* **`EndpointSlice`**: Enables EndpointSlices for more scalable and extensible network endpoints. See [Enabling EndpointSlices](/concepts/services-networking/endpoint-slices.md).
* **`EndpointSliceNodeName`**: Enables EndpointSlice `nodeName` field.
* **`EndpointSliceProxying`**: When enabled, kube-proxy running on Linux will use EndpointSlices as the primary data source instead of Endpoints, enabling scalability and performance improvements. See [Enabling Endpoint Slices](/concepts/services-networking/endpoint-slices.md).
* **`EndpointSliceTerminatingCondition`**: Enables EndpointSlice `terminating` and `serving` condition fields.
* **`EphemeralContainers`**: Enable the ability to add [ephemeral containers](/reference/glossary/ephemeral-container.md) to running Pods.
* **`EvenPodsSpread`**: Enable pods to be scheduled evenly across topology domains. See [Pod Topology Spread Constraints](/concepts/scheduling-eviction/topology-spread-constraints.md).
* **`ExpandCSIVolumes`**: Enable the expanding of CSI volumes.
* **`ExpandInUsePersistentVolumes`**: Enable expanding in-use PVCs. See [Resizing an in-use PersistentVolumeClaim](/concepts/storage/persistent-volumes.md#resizing-an-in-use-persistentvolumeclaim).
* **`ExpandPersistentVolumes`**: Enable the expanding of persistent volumes. See [Expanding Persistent Volumes Claims](/concepts/storage/persistent-volumes.md#expanding-persistent-volumes-claims).
* **`ExpandedDNSConfig`**: Enable kubelet and kube-apiserver to allow more DNS search paths and longer list of DNS search paths. This feature requires container runtime support(Containerd: v1.5.6 or higher, CRI-O: v1.22 or higher). See [Expanded DNS Configuration](/concepts/services-networking/dns-pod-service.md#expanded-dns-configuration).
* **`ExperimentalCriticalPodAnnotation`**: Enable annotating specific pods as *critical* so that their [scheduling is guaranteed](/tasks/administer-cluster/guaranteed-scheduling-critical-addon-pods.md). This feature is deprecated by Pod Priority and Preemption as of v1.13.
* **`ExperimentalHostUserNamespaceDefaulting`**: Enabling the defaulting user namespace to host. This is for containers that are using other host namespaces, host mounts, or containers that are privileged or using specific non-namespaced capabilities (e.g. `MKNODE`, `SYS_MODULE` etc.). This should only be enabled if user namespace remapping is enabled in the Docker daemon.
* **`ExternalPolicyForExternalIP`**: Fix a bug where ExternalTrafficPolicy is not applied to Service ExternalIPs.
* **`GCERegionalPersistentDisk`**: Enable the regional PD feature on GCE.
* **`GRPCContainerProbe`**: Enables the gRPC probe method for {Liveness,Readiness,Startup}Probe. See [Configure Liveness, Readiness and Startup Probes](/tasks/configure-pod-container/configure-liveness-readiness-startup-probes.md#define-a-grpc-liveness-probe).
* **`GangScheduling`**: Enables the GangScheduling plugin in kube-scheduler, which implements "all-or-nothing" scheduling algorithm. The [Workload API](/concepts/workloads/workload-api.md) is used to express the requirements. This feature gate was removed in 1.37 and merged together with the `GenericWorkload` feature gate.
* **`GenericEphemeralVolume`**: Enables ephemeral, inline volumes that support all features of normal volumes (can be provided by third-party storage vendors, storage capacity tracking, restore from snapshot, etc.). See [Ephemeral Volumes](/concepts/storage/ephemeral-volumes.md).
* **`HPAContainerMetrics`**: Allow [HorizontalPodAutoscalers](/reference/glossary/horizontal-pod-autoscaler.md) to scale based on metrics from individual containers within target pods.
* **`HonorPVReclaimPolicy`**: Honor persistent volume reclaim policy when it is `Delete` irrespective of PV-PVC deletion ordering. For more details, check the [PersistentVolume deletion protection finalizer](/concepts/storage/persistent-volumes.md#persistentvolume-deletion-protection-finalizer) documentation.
* **`HugePageStorageMediumSize`**: Enable support for multiple sizes pre-allocated [huge pages](/tasks/manage-hugepages/scheduling-hugepages.md).
* **`HugePages`**: Enable the allocation and consumption of pre-allocated [huge pages](/tasks/manage-hugepages/scheduling-hugepages.md).
* **`HyperVContainer`**: Enable [Hyper-V isolation](https://docs.microsoft.com/en-us/virtualization/windowscontainers/manage-containers/hyperv-container) for Windows containers.
* **`IPTablesOwnershipCleanup`**: This causes kubelet to no longer create legacy iptables rules.
* **`IPv6DualStack`**: Enable [dual stack](/concepts/services-networking/dual-stack.md) support for IPv6.
* **`IdentifyPodOS`**: Allows the Pod OS field to be specified. This helps in identifying the OS of the pod authoritatively during the API server admission time.
* **`ImmutableEphemeralVolumes`**: Allows for marking individual Secrets and ConfigMaps as immutable for better safety and performance.
* **`InTreePluginAWSUnregister`**: Stops registering the aws-ebs in-tree plugin in kubelet and volume controllers.
* **`InTreePluginAzureDiskUnregister`**: Stops registering the azuredisk in-tree plugin in kubelet and volume controllers.
* **`InTreePluginAzureFileUnregister`**: Stops registering the azurefile in-tree plugin in kubelet and volume controllers.
* **`InTreePluginGCEUnregister`**: Stops registering the gce-pd in-tree plugin in kubelet and volume controllers.
* **`InTreePluginOpenStackUnregister`**: Stops registering the OpenStack cinder in-tree plugin in kubelet and volume controllers.
* **`InTreePluginPortworxUnregister`**: Stops registering the Portworx in-tree plugin in kubelet and volume controllers.
* **`InTreePluginRBDUnregister`**: Stops registering the RBD in-tree plugin within kubelet and volume controllers.
* **`InTreePluginvSphereUnregister`**: Stops registering the vSphere in-tree plugin in kubelet and volume controllers.
* **`IndexedJob`**: Allows the [Job](/concepts/workloads/controllers/job.md) controller to manage Pod completions per completion index.
* **`IngressClassNamespacedParams`**: Allow namespace-scoped parameters reference in `IngressClass` resource. This feature adds two fields - `Scope` and `Namespace` to `IngressClass.spec.parameters`.
* **`Initializers`**: Allow asynchronous coordination of object creation using the Initializers admission plugin.
* **`JobMutableNodeSchedulingDirectives`**: Allows updating node scheduling directives in the pod template of [Job](/concepts/workloads/controllers/job.md).
* **`JobPodFailurePolicy`**: Allow users to specify handling of pod failures based on container exit codes and pod conditions.
* **`JobReadyPods`**: Enables tracking the number of Pods that have a `Ready` [condition](/concepts/workloads/pods/pod-lifecycle.md#pod-conditions). The count of `Ready` pods is recorded in the [status](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/job-v1/#JobStatus) of a [Job](/concepts/workloads/controllers/job.md) status.
* **`JobTrackingWithFinalizers`**: Enables tracking [Job](/concepts/workloads/controllers/job.md) completions without relying on Pods remaining in the cluster indefinitely. The Job controller uses Pod finalizers and a field in the Job status to keep track of the finished Pods to count towards completion.
* **`KMSv2`**: Enables KMS v2 API for encryption at rest. See [Using a KMS Provider for data encryption](/tasks/administer-cluster/kms-provider.md) for more details.
* **`KMSv2KDF`**: Enables KMS v2 to generate single use data encryption keys. See [Using a KMS Provider for data encryption](/tasks/administer-cluster/kms-provider.md) for more details. If the `KMSv2` feature gate is not enabled in your cluster, the value of the `KMSv2KDF` feature gate has no effect.
* **`KubeProxyDrainingTerminatingNodes`**: Implement connection draining for terminating nodes for `externalTrafficPolicy: Cluster` services.
* **`KubeletConfigFile`**: Enable loading kubelet configuration from a file specified using a config file. See [setting kubelet parameters via a config file](/tasks/administer-cluster/kubelet-config-file.md) for more details.
* **`KubeletCredentialProviders`**: Enable kubelet exec credential providers for image pull credentials.
* **`KubeletPluginsWatcher`**: Enable probe-based plugin watcher utility to enable kubelet to discover plugins such as [CSI volume drivers](/concepts/storage/volumes.md#csi).
* **`KubeletPodResources`**: Enable the kubelet's pod resources gRPC endpoint. See [Support Device Monitoring](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/606-compute-device-assignment/README.md) for more details.
* **`KubeletPodResourcesGetAllocatable`**: Enable the kubelet's pod resources `GetAllocatableResources` functionality. This API augments the [resource allocation reporting](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md#monitoring-device-plugin-resources)
* **`LegacyNodeRoleBehavior`**: When disabled, legacy behavior in service load balancers and node disruption will ignore the `node-role.kubernetes.io/master` label in favor of the feature-specific labels provided by `NodeDisruptionExclusion` and `ServiceNodeExclusion`.
* **`LegacyServiceAccountTokenCleanUp`**: Enable cleaning up Secret-based [service account tokens](/concepts/security/service-accounts.md#get-a-token) when they are not used in a specified time (default to be one year).
* **`LegacyServiceAccountTokenNoAutoGeneration`**: Stop auto-generation of Secret-based [service account tokens](/concepts/security/service-accounts.md#get-a-token).
* **`LegacyServiceAccountTokenTracking`**: Track usage of Secret-based [service account tokens](/concepts/security/service-accounts.md#get-a-token).
* **`LoadBalancerIPMode`**: Allows setting `ipMode` for Services where `type` is set to `LoadBalancer`. See [Specifying IPMode of load balancer status](/concepts/services-networking/service.md#load-balancer-ip-mode) for more information.
* **`LocalStorageCapacityIsolation`**: Enable the consumption of [local ephemeral storage](/concepts/configuration/manage-resources-containers.md) and also the `sizeLimit` property of an [emptyDir volume](/concepts/storage/volumes.md#emptydir).
* **`MinDomainsInPodTopologySpread`**: Enable `minDomains` in [Pod topology spread constraints](/concepts/scheduling-eviction/topology-spread-constraints.md).
* **`MinimizeIPTablesRestore`**: Enables new performance improvement logics in the kube-proxy iptables mode.
* **`MixedProtocolLBService`**: Enable using different protocols in the same `LoadBalancer` type Service instance.
* **`MountContainers`**: Enable using utility containers on host as the volume mounter.
* **`MountPropagation`**: Enable sharing volume mounted by one container to other containers or pods. For more details, please see [mount propagation](/concepts/storage/volumes.md#mount-propagation).
* **`MultiCIDRRangeAllocator`**: Enables the MultiCIDR range allocator.
* **`NamespaceDefaultLabelName`**: Configure the API Server to set an immutable [label](/reference/glossary/label.md) `kubernetes.io/metadata.name` on all namespaces, containing the namespace name.
* **`NetworkPolicyEndPort`**: Allows you to define ports in a [NetworkPolicy](/concepts/services-networking/network-policies.md) rule as a range of port numbers.
* **`NetworkPolicyStatus`**: Enable the `status` subresource for NetworkPolicy objects.
* **`NewVolumeManagerReconstruction`**: Enables improved discovery of mounted volumes during kubelet startup. Since the associated code had been significantly refactored, Kubernetes versions 1.25 to 1.29 allowed you to opt-out in case the kubelet got stuck at the startup, or did not unmount volumes from terminated Pods. This refactoring was behind the `SELinuxMountReadWriteOncePod` feature gate in Kubernetes releases 1.25 and 1.26.
* **`NodeDisruptionExclusion`**: Enable use of the Node label `node.kubernetes.io/exclude-disruption` which prevents nodes from being evacuated during zone failures.
* **`NodeLease`**: Enable the new Lease API to report node heartbeats, which could be used as a node health signal.
* **`NodeOutOfServiceVolumeDetach`**: When a Node is marked out-of-service using the `node.kubernetes.io/out-of-service` taint, Pods on the node will be forcefully deleted if they can not tolerate this taint, and the volume detach operations for Pods terminating on the node will happen immediately. The deleted Pods can recover quickly on different nodes.
* **`NonPreemptingPriority`**: Enable `preemptionPolicy` field for PriorityClass and Pod.
* **`OpenAPIV3`**: Enables the API server to publish OpenAPI v3.
* **`PDBUnhealthyPodEvictionPolicy`**: Enables the `unhealthyPodEvictionPolicy` field of a `PodDisruptionBudget`. This specifies when unhealthy pods should be considered for eviction. Please see [Unhealthy Pod Eviction Policy](/tasks/run-application/configure-pdb.md#unhealthy-pod-eviction-policy) for more details.
* **`PVCProtection`**: Enable the prevention of a PersistentVolumeClaim (PVC) from being deleted when it is still used by any Pod.
* **`PersistentLocalVolumes`**: Enable the usage of `local` volume type in Pods. Pod affinity has to be specified if requesting a `local` volume.
* **`PersistentVolumeLastPhaseTransitionTime`**: Adds a new field to PersistentVolume which holds a timestamp of when the volume last transitioned its phase.
* **`PodAffinityNamespaceSelector`**: Enable the [Pod Affinity Namespace Selector](/concepts/scheduling-eviction/assign-pod-node.md#namespace-selector) and [CrossNamespacePodAffinity](/concepts/policy/resource-quotas.md#cross-namespace-pod-affinity-quota) quota scope features.
* **`PodDisruptionBudget`**: Enable the [PodDisruptionBudget](/tasks/run-application/configure-pdb.md) feature.
* **`PodDisruptionConditions`**: Enabled support for appending a dedicated pod condition indicating that the pod is being deleted due to a disruption.
* **`PodHasNetworkCondition`**: Enable the kubelet to mark the [PodHasNetwork](/concepts/workloads/pods/pod-lifecycle.md#pod-has-network) condition on pods. This was renamed to `PodReadyToStartContainersCondition` in 1.28.
* **`PodHostIPs`**: Enable the `status.hostIPs` field for pods and the [downward API](/reference/glossary/downward-api.md). The field lets you expose host IP addresses to workloads.
* **`PodOverhead`**: Enable the [PodOverhead](/concepts/scheduling-eviction/pod-overhead.md) feature to account for pod overheads.
* **`PodPriority`**: Enable the descheduling and preemption of Pods based on their [priorities](/concepts/scheduling-eviction/pod-priority-preemption.md).
* **`PodReadinessGates`**: Enable the setting of `PodReadinessGate` field for extending Pod readiness evaluation. See [Pod readiness gate](/concepts/workloads/pods/pod-lifecycle.md#pod-readiness-gate) for more details.
* **`PodSecurity`**: Enables the `PodSecurity` admission plugin.
* **`PodShareProcessNamespace`**: Enable the setting of `shareProcessNamespace` in a Pod for sharing a single process namespace between containers running in a pod. More details can be found in [Share Process Namespace between Containers in a Pod](/tasks/configure-pod-container/share-process-namespace.md).
* **`PreferNominatedNode`**: This flag tells the scheduler whether the nominated nodes will be checked first before looping through all the other nodes in the cluster.
* **`ProbeTerminationGracePeriod`**: Enable [setting probe-level `terminationGracePeriodSeconds`](/tasks/configure-pod-container/configure-liveness-readiness-startup-probes.md#probe-level-terminationgraceperiodseconds) on pods. See the [enhancement proposal](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2238-liveness-probe-grace-period) for more details.
* **`ProxyTerminatingEndpoints`**: Enable the kube-proxy to handle terminating endpoints when `ExternalTrafficPolicy=Local`.
* **`ReadOnlyAPIDataVolumes`**: Set [`configMap`](/concepts/storage/volumes.md#configmap), [`secret`](/concepts/storage/volumes.md#secret), [`downwardAPI`](/concepts/storage/volumes.md#downwardapi) and [`projected`](/concepts/storage/volumes.md#projected) [volumes](/reference/glossary/volume.md) to be mounted read-only. Since Kubernetes v1.10, these volume types are always read-only and you cannot opt out.
* **`ReadWriteOncePod`**: Enables the usage of `ReadWriteOncePod` PersistentVolume access mode.
* **`RelaxedDNSSearchValidation`**: Relax the server side validation for the DNS search string (`.spec.dnsConfig.searches`) for containers. For example, with this gate enabled, it is okay to include the `_` character in the DNS name search string.
* **`RemainingItemCount`**: Allow the API servers to show a count of remaining items in the response to a [chunking list request](/reference/using-api/api-concepts.md#retrieving-large-results-sets-in-chunks).
* **`RemoveSelfLink`**: Sets the `.metadata.selfLink` field to blank (empty string) for all objects and collections. This field has been deprecated since the Kubernetes v1.16 release. When this feature is enabled, the `.metadata.selfLink` field remains part of the Kubernetes API, but is always unset.
* **`RequestManagement`**: Enables managing request concurrency with prioritization and fairness at each API server. Deprecated by `APIPriorityAndFairness` since 1.17.
* **`ResourceLimitsPriorityFunction`**: Enable a scheduler priority function that assigns a lowest possible score of 1 to a node that satisfies at least one of the input Pod's cpu and memory limits. The intent is to break ties between nodes with same scores.
* **`ResourceQuotaScopeSelectors`**: Enable resource quota scope selectors.
* **`RetroactiveDefaultStorageClass`**: Allow assigning StorageClass to unbound PVCs retroactively.
* **`RootCAConfigMap`**: Configure the `kube-controller-manager` to publish a [ConfigMap](/reference/glossary/configmap.md) named `kube-root-ca.crt` to every namespace. This ConfigMap contains a CA bundle used for verifying connections to the kube-apiserver. See [Bound Service Account Tokens](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/1205-bound-service-account-tokens/README.md) for more details.
* **`RotateKubeletClientCertificate`**: Enable the rotation of the client TLS certificate on the kubelet. See [kubelet configuration](/reference/access-authn-authz/kubelet-tls-bootstrapping.md#kubelet-configuration) for more details.
* **`RunAsGroup`**: Enable control over the primary group ID set on the init processes of containers.
* **`RuntimeClass`**: Enable the [RuntimeClass](/concepts/containers/runtime-class.md) feature for selecting container runtime configurations.
* **`SCTPSupport`**: Enables the _SCTP_ `protocol` value in Pod, Service, Endpoints, EndpointSlice, and NetworkPolicy definitions.
* **`ScheduleDaemonSetPods`**: Enable DaemonSet Pods to be scheduled by the default scheduler instead of the DaemonSet controller.
* **`SeccompDefault`**: Enables the use of `RuntimeDefault` as the default seccomp profile for all workloads. The seccomp profile is specified in the `securityContext` of a Pod and/or a Container.
* **`SecurityContextDeny`**: This gate signals that the `SecurityContextDeny` admission controller is deprecated.
* **`SelectorIndex`**: Allows label and field based indexes in API server watch cache to accelerate list operations.
* **`ServerSideApply`**: Enables the [Sever Side Apply (SSA)](/reference/using-api/server-side-apply.md) feature on the API Server.
* **`ServerSideFieldValidation`**: Enables server-side field validation. This means the validation of resource schema is performed at the API server side rather than the client side (for example, the `kubectl create` or `kubectl apply` command line).
* **`ServiceAccountIssuerDiscovery`**: Enable OIDC discovery endpoints (issuer and JWKS URLs) for the service account issuer in the API server. See [Configure Service Accounts for Pods](/tasks/configure-pod-container/configure-service-account.md#service-account-issuer-discovery) for more details.
* **`ServiceAppProtocol`**: Enables the `appProtocol` field on Services and Endpoints.
* **`ServiceIPStaticSubrange`**: Enables a strategy for Services ClusterIP allocations, whereby the ClusterIP range is subdivided. Dynamic allocated ClusterIP addresses will be allocated preferently from the upper range allowing users to assign static ClusterIPs from the lower range with a low risk of collision. See [Avoiding collisions](/reference/networking/virtual-ips.md#avoiding-collisions) for more details.
* **`ServiceInternalTrafficPolicy`**: Enables the `internalTrafficPolicy` field on Services
* **`ServiceLBNodePortControl`**: Enables the `allocateLoadBalancerNodePorts` field on Services.
* **`ServiceLoadBalancerClass`**: Enables the `loadBalancerClass` field on Services. See [Specifying class of load balancer implementation](/concepts/services-networking/service.md#load-balancer-class) for more details.
* **`ServiceLoadBalancerFinalizer`**: Enable finalizer protection for Service load balancers.
* **`ServiceNodeExclusion`**: Enable the exclusion of nodes from load balancers created by a cloud provider. A node is eligible for exclusion if labelled with "`node.kubernetes.io/exclude-from-external-load-balancers`".
* **`ServiceNodePortStaticSubrange`**: Enables the use of different port allocation strategies for NodePort Services. For more details, see [reserve NodePort ranges to avoid collisions](/concepts/services-networking/service.md#avoid-nodeport-collisions).
* **`ServiceTopology`**: Enable service to route traffic based upon the Node topology of the cluster.
* **`SetHostnameAsFQDN`**: Enable the ability of setting Fully Qualified Domain Name(FQDN) as the hostname of a pod. See [Pod's `setHostnameAsFQDN` field](/concepts/services-networking/dns-pod-service.md#pod-sethostnameasfqdn-field).
* **`SizeMemoryBackedVolumes`**: Enable kubelets to determine the size limit for memory-backed volumes (mainly `emptyDir` volumes).
* **`SkipReadOnlyValidationGCE`**: Skip validation that GCE PersistentDisk volumes are in read-only mode.
* **`StableLoadBalancerNodeSet`**: Enables less load balancer re-configurations by the service controller (KCCM) as an effect of changing node state.
* **`StartupProbe`**: Enable the [startup](/concepts/workloads/pods/pod-lifecycle.md#when-should-you-use-a-startup-probe) probe in the kubelet.
* **`StatefulSetMinReadySeconds`**: Allows `minReadySeconds` to be respected by the StatefulSet controller.
* **`StorageObjectInUseProtection`**: Postpone the deletion of PersistentVolume or PersistentVolumeClaim objects if they are still being used.
* **`StreamingProxyRedirects`**: Instructs the API server to intercept (and follow) redirects from the backend (kubelet) for streaming requests. Examples of streaming requests include the `exec`, `attach` and `port-forward` requests.
* **`SupportIPVSProxyMode`**: Enable providing in-cluster service load balancing using IPVS. See [service proxies](/reference/networking/virtual-ips.md) for more details.
* **`SupportNodePidsLimit`**: Enable the support to limiting PIDs on the Node. The parameter `pid=<number>` in the `--system-reserved` and `--kube-reserved` options can be specified to ensure that the specified number of process IDs will be reserved for the system as a whole and for Kubernetes system daemons respectively.
* **`SupportPodPidsLimit`**: Enable the support to limiting PIDs in Pods.
* **`SuspendJob`**: Enable support to suspend and resume Jobs. For more details, see [the Jobs docs](/concepts/workloads/controllers/job.md).
* **`Sysctls`**: Enable support for namespaced kernel parameters (sysctls) that can be set for each pod. See [sysctls](/tasks/administer-cluster/sysctl-cluster.md) for more details.
* **`TTLAfterFinished`**: Allow a [TTL controller](/concepts/workloads/controllers/ttlafterfinished.md) to clean up resources after they finish execution.
* **`TaintBasedEvictions`**: Enable evicting pods from nodes based on taints on Nodes and tolerations on Pods. See [taints and tolerations](/concepts/scheduling-eviction/taint-and-toleration.md) for more details.
* **`TaintNodesByCondition`**: Enable automatic tainting nodes based on [node conditions](/concepts/architecture/nodes.md#condition).
* **`TokenRequest`**: Enable the `TokenRequest` endpoint on service account resources.
* **`TokenRequestProjection`**: Enable the injection of service account tokens into a Pod through a [`projected` volume](/concepts/storage/volumes.md#projected).
* **`TopologyAwareHints`**: Enables topology aware routing based on topology hints in EndpointSlices. See [Topology Aware Hints](/concepts/services-networking/topology-aware-routing.md) for more details.
* **`TopologyManager`**: Enable a mechanism to coordinate fine-grained hardware resource assignments for different components in Kubernetes. See [Control Topology Management Policies on a node](/tasks/administer-cluster/topology-manager.md).
* **`UserNamespacesPodSecurityStandards`**: Enable Pod Security Standards policies relaxation for pods that run with namespaces. You must set the value of this feature gate consistently across all nodes in your cluster, and you must also enable `UserNamespacesSupport` to use this feature. This feature gate is not part of Kubernetes v1.35 (and onwards) because, for those versions of Kubernetes, all supported kubelet versions will fail to create a pod if it requests a user namespace, but the node does not.
* **`UserNamespacesStatelessPodsSupport`**: Enable user namespace support for stateless Pods. This feature gate was superseded by the `UserNamespacesSupport` feature gate in the Kubernetes v1.28 release.
* **`ValidateProxyRedirects`**: This flag controls whether the API server should validate that redirects are only followed to the same host. Only used if the `StreamingProxyRedirects` flag is enabled.
* **`ValidatingAdmissionPolicy`**: Enable [ValidatingAdmissionPolicy](/reference/access-authn-authz/validating-admission-policy.md) support for CEL validations be used in Admission Control.
* **`VolumeCapacityPriority`**: Enable support for prioritizing nodes in different topologies based on available PV capacity. This feature is renamed to `StorageCapacityScoring` in v1.33.
* **`VolumePVCDataSource`**: Enable support for specifying an existing PVC as a DataSource.
* **`VolumeScheduling`**: Enable volume topology aware scheduling and make the PersistentVolumeClaim (PVC) binding aware of scheduling decisions. It also enables the usage of [`local`](/concepts/storage/volumes.md#local) volume type when used together with the `PersistentLocalVolumes` feature gate.
* **`VolumeSnapshotDataSource`**: Enable volume snapshot data source support.
* **`VolumeSubpath`**: Allow mounting a subpath of a volume in a container.
* **`VolumeSubpathEnvExpansion`**: Enable `subPathExpr` field for expanding environment variables into a `subPath`.
* **`WarningHeaders`**: Allow sending warning headers in API responses.
* **`WatchBookmark`**: Enable support for watch bookmark events.
* **`WindowsEndpointSliceProxying`**: When enabled, kube-proxy running on Windows will use EndpointSlices as the primary data source instead of Endpoints, enabling scalability and performance improvements. See [Enabling Endpoint Slices](/concepts/services-networking/endpoint-slices.md).
* **`WindowsGMSA`**: Enables passing of GMSA credential specs from pods to container runtimes.
* **`WindowsHostProcessContainers`**: Enables support for Windows HostProcess containers.
* **`WindowsRunAsUserName`**: Enable support for running applications in Windows containers with as a non-default user. See [Configuring RunAsUserName](/tasks/configure-pod-container/configure-runasusername.md) for more details.
* **`WorkloadAwarePreemption`**: Enables the support for [Workload-aware preemption](/concepts/scheduling-eviction/workload-aware-preemption.md). This feature gate was removed in 1.37 and merged together with the `GenericWorkload` feature gate. When enabled, if a PodGroup fails to schedule, the scheduler will use a workload-aware preemption algorithm to select victims to preempt instead of the default pod preemption algorithm.
* **`ZeroLimitedNominalConcurrencyShares`**: Allow [priority & fairness](/concepts/cluster-administration/flow-control.md) in the API server to use a zero value for the `nominalConcurrencyShares` field of the `limited` section of a priority level.
