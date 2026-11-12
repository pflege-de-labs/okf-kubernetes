---
type: concept
title: Feature Gates
description: This page contains an overview of the various feature gates an administrator can specify on different Kubernetes components.
resource: https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/
tags: [kubernetes, reference, command-line-tools-reference]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/command-line-tools-reference/feature-gates/index.md
    title: 'Kubernetes website source: reference/command-line-tools-reference/feature-gates/index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/
    title: Feature Gates | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/command-line-tools-reference/feature-gates/index.md
---

# Feature Gates

This page contains an overview of the various feature gates an administrator
can specify on different Kubernetes components.

See [feature stages](#feature-stages) for an explanation of the stages for a feature.

## Overview

Feature gates are a set of key=value pairs that describe Kubernetes features.
You can turn these features on or off using the `--feature-gates` command line flag
on each Kubernetes component.

## How to enable Feature Gates

To enable or disable a feature gate for a particular Kubernetes component, use the 
`--feature-gates` flag.

This flag accepts a comma-separated list of key=value pairs, where each key is a 
feature gate name and each value is either `true` (enable) or `false` (disable).

**Example usage:**

```shell
kube-apiserver --feature-gates=FeatureName1=true,FeatureName2=false
kubelet --feature-gates=GracefulNodeShutdown=true
```

Each Kubernetes component supports only the feature gates relevant to its functions.
Use `<component> -h` to list available feature gates for a specific component.

For detailed instructions on configuring feature gates in your cluster, see
[Configure Feature Gates](/tasks/administer-cluster/configure-feature-gates.md).

## Feature gates in Kubernetes v1.37 {#list-of-gates}

The following tables are a summary of the feature gates that you can set on
different Kubernetes components.

- The "Since" column contains the Kubernetes release when a feature is introduced
  or its release stage is changed.
- The "Until" column, if not empty, contains the last Kubernetes release in which
  you can still use a feature gate.
- If a feature is in the Alpha or Beta state, you can find the feature listed
  in the [Alpha/Beta feature gate table](#feature-gates-for-alpha-or-beta-features).
- If a feature is stable you can find all stages for that feature listed in the
  [Graduated/Deprecated feature gate table](#feature-gates-for-graduated-or-deprecated-features).
- The [Graduated/Deprecated feature gate table](#feature-gates-for-graduated-or-deprecated-features)
  also lists deprecated and withdrawn features.

**Note:**
For a reference to old feature gates that are removed, please refer to
[feature gates removed](/reference/command-line-tools-reference/feature-gates-removed.md).

### Feature gates for Alpha or Beta features

| Feature | Default | Stage | Since | Until |
|---|---|---|---|---|
| `APIResponseCompression` | `false` | alpha | 1.7 | 1.15 |
| `APIResponseCompression` | `true` | beta | 1.16 |  |
| `APIServerIdentity` | `false` | alpha | 1.20 | 1.25 |
| `APIServerIdentity` | `true` | beta | 1.26 |  |
| `APIServerWebhookAuthenticationToken` | `false` | alpha | 1.37 |  |
| `APIServingWithRoutine` | `false` | alpha | 1.30 |  |
| `AllowParsingUserUIDFromCertAuth` | `false` | alpha | 1.33 | 1.33 |
| `AllowParsingUserUIDFromCertAuth` | `true` | beta | 1.34 |  |
| `AllowUnsafeMalformedObjectDeletion` | `false` | alpha | 1.32 | 1.36 |
| `AllowUnsafeMalformedObjectDeletion` | `true` | beta | 1.37 |  |
| `AtomicFIFO` | `true` | beta | 1.36 |  |
| `AtomicWriteVolumeUserFields` | `false` | alpha | 1.37 |  |
| `AuthorizePodWebsocketUpgradeCreatePermission` | `true` | beta | 1.35 |  |
| `CBORServingAndStorage` | `false` | alpha | 1.32 |  |
| `CPUManagerPolicyAlphaOptions` | `false` | alpha | 1.23 |  |
| `CPUManagerPolicyBetaOptions` | `true` | beta | 1.23 |  |
| `CRDObservedGenerationTracking` | `false` | beta | 1.35 |  |
| `CRIListStreaming` | `false` | alpha | 1.36 |  |
| `CSIVolumeHealth` | `false` | alpha | 1.21 |  |
| `ClearingNominatedNodeNameAfterBinding` | `false` | alpha | 1.34 | 1.34 |
| `ClearingNominatedNodeNameAfterBinding` | `true` | beta | 1.35 |  |
| `CloudControllerManagerWatchBasedRoutesReconciliation` | `false` | alpha | 1.35 |  |
| `CloudControllerManagerWebhook` | `false` | alpha | 1.27 |  |
| `ComponentFlagz` | `false` | alpha | 1.32 | 1.35 |
| `ComponentFlagz` | `true` | beta | 1.36 |  |
| `ComponentStatusz` | `false` | alpha | 1.32 | 1.35 |
| `ComponentStatusz` | `true` | beta | 1.36 |  |
| `CompositePodGroup` | `false` | alpha | 1.37 |  |
| `ConcurrentWatchObjectDecode` | `false` | beta | 1.31 | 1.36 |
| `ConcurrentWatchObjectDecode` | `true` | beta | 1.37 |  |
| `ConstrainedImpersonation` | `false` | alpha | 1.35 | 1.35 |
| `ConstrainedImpersonation` | `true` | beta | 1.36 |  |
| `ContainerCheckpoint` | `false` | alpha | 1.25 | 1.29 |
| `ContainerCheckpoint` | `true` | beta | 1.30 |  |
| `ContainerRestartRules` | `false` | alpha | 1.34 | 1.34 |
| `ContainerRestartRules` | `true` | beta | 1.35 |  |
| `ContainerStopSignals` | `false` | alpha | 1.33 |  |
| `ContextualLogging` | `false` | alpha | 1.24 |  |
| `ContextualLogging` | `true` | beta | 1.30 |  |
| `ControllerManagerReleaseLeaderElectionLockOnExit` | `false` | alpha | 1.36 |  |
| `CoordinatedLeaderElection` | `false` | alpha | 1.31 | 1.32 |
| `CoordinatedLeaderElection` | `false` | beta | 1.33 |  |
| `CrossNamespaceVolumeDataSource` | `false` | alpha | 1.26 |  |
| `CustomCPUCFSQuotaPeriod` | `false` | alpha | 1.12 |  |
| `DRAConsumableCapacity` | `false` | alpha | 1.34 | 1.35 |
| `DRAConsumableCapacity` | `true` | beta | 1.36 |  |
| `DRADerivedAttributes` | `false` | alpha | 1.37 |  |
| `DRADeviceBindingConditions` | `false` | alpha | 1.34 | 1.35 |
| `DRADeviceBindingConditions` | `true` | beta | 1.36 |  |
| `DRADeviceCompatibilityGroups` | `false` | alpha | 1.37 |  |
| `DRAListTypeAttributes` | `false` | alpha | 1.36 |  |
| `DRANodeAllocatableResources` | `false` | alpha | 1.36 |  |
| `DRAOptionalNodeOperations` | `false` | alpha | 1.37 |  |
| `DRAPartitionableDevices` | `false` | alpha | 1.33 | 1.35 |
| `DRAPartitionableDevices` | `true` | beta | 1.36 |  |
| `DRAPartitionableDevicesType` | `false` | alpha | 1.37 |  |
| `DRAResourceClaimGranularStatusAuthorization` | `true` | beta | 1.36 |  |
| `DRAResourcePoolStatus` | `false` | alpha | 1.36 |  |
| `DRASchedulerFilterTimeout` | `false` | alpha | 1.34 |  |
| `DRAWorkloadResourceClaims` | `false` | alpha | 1.36 |  |
| `DRAWorkloadResourceClaims` | `false` | beta | 1.37 |  |
| `DeclarativeValidationBeta` | `true` | beta | 1.36 |  |
| `DefaultPodSysctls` | `false` | alpha | 1.37 |  |
| `DeploymentReplicaSetTerminatingReplicas` | `false` | alpha | 1.33 | 1.34 |
| `DeploymentReplicaSetTerminatingReplicas` | `true` | beta | 1.35 |  |
| `DetectCacheInconsistency` | `true` | beta | 1.34 |  |
| `DisableCPUQuotaWithExclusiveCPUs` | `true` | beta | 1.33 |  |
| `EmptyDirVolumeMode` | `false` | alpha | 1.37 |  |
| `EnvFiles` | `false` | alpha | 1.34 | 1.34 |
| `EnvFiles` | `true` | beta | 1.35 |  |
| `EtcdRangeStream` | `true` | beta | 1.37 |  |
| `EventedPLEG` | `false` | alpha | 1.26 |  |
| `ExcludeAdmissionWebhookVirtualResources` | `true` | beta | 1.37 |  |
| `ExtendWebSocketsToKubelet` | `true` | beta | 1.36 |  |
| `GRPCContainerProbeTLS` | `false` | alpha | 1.37 |  |
| `GenericWorkload` | `false` | alpha | 1.35 | 1.36 |
| `GenericWorkload` | `false` | beta | 1.37 |  |
| `GracefulNodeShutdown` | `false` | alpha | 1.20 | 1.20 |
| `GracefulNodeShutdown` | `true` | beta | 1.21 |  |
| `GracefulNodeShutdownBasedOnPodPriority` | `false` | alpha | 1.23 | 1.23 |
| `GracefulNodeShutdownBasedOnPodPriority` | `true` | beta | 1.24 |  |
| `H2CContainerProbe` | `false` | alpha | 1.37 |  |
| `HPAScaleToZero` | `false` | alpha | 1.16 | 1.36 |
| `HPAScaleToZero` | `true` | beta | 1.37 |  |
| `HugepageAwareEviction` | `true` | beta | 1.37 |  |
| `ImageVolumeWithDigest` | `false` | alpha | 1.35 |  |
| `InOrderInformers` | `true` | alpha | 1.33 | 1.33 |
| `InOrderInformers` | `true` | beta | 1.34 |  |
| `InPlacePodLevelResourcesVerticalScaling` | `false` | alpha | 1.35 | 1.35 |
| `InPlacePodLevelResourcesVerticalScaling` | `true` | beta | 1.36 |  |
| `InPlacePodVerticalScalingExclusiveCPUs` | `false` | alpha | 1.32 |  |
| `InPlacePodVerticalScalingExclusiveMemory` | `false` | alpha | 1.34 |  |
| `InPlacePodVerticalScalingMemoryBackedVolumes` | `false` | alpha | 1.37 |  |
| `InPlacePodVerticalScalingSchedulerPreemption` | `false` | alpha | 1.37 |  |
| `KubeProxyNFTablesLocalhostNodePorts` | `false` | alpha | 1.37 |  |
| `KubeletCrashLoopBackOffMax` | `false` | alpha | 1.32 | 1.34 |
| `KubeletCrashLoopBackOffMax` | `true` | beta | 1.35 |  |
| `KubeletEnsureSecretPulledImages` | `false` | alpha | 1.33 | 1.34 |
| `KubeletEnsureSecretPulledImages` | `true` | beta | 1.35 |  |
| `KubeletInUserNamespace` | `false` | alpha | 1.22 | 1.36 |
| `KubeletInUserNamespace` | `true` | beta | 1.37 |  |
| `KubeletSeparateDiskGC` | `false` | alpha | 1.29 | 1.30 |
| `KubeletSeparateDiskGC` | `true` | beta | 1.31 |  |
| `KubeletServiceAccountTokenForCredentialProviders` | `false` | alpha | 1.33 | 1.33 |
| `KubeletServiceAccountTokenForCredentialProviders` | `true` | beta | 1.34 |  |
| `ListFromCacheSnapshot` | `false` | alpha | 1.33 | 1.33 |
| `ListFromCacheSnapshot` | `true` | beta | 1.34 |  |
| `LocalStorageCapacityIsolationFSQuotaMonitoring` | `false` | alpha | 1.15 | 1.30 |
| `LocalStorageCapacityIsolationFSQuotaMonitoring` | `false` | beta | 1.31 |  |
| `LoggingAlphaOptions` | `false` | alpha | 1.24 |  |
| `LoggingBetaOptions` | `true` | beta | 1.24 |  |
| `ManifestBasedAdmissionControlConfig` | `false` | alpha | 1.36 | 1.36 |
| `ManifestBasedAdmissionControlConfig` | `true` | beta | 1.37 |  |
| `MatchLabelKeysInPodTopologySpread` | `false` | alpha | 1.25 | 1.26 |
| `MatchLabelKeysInPodTopologySpread` | `true` | beta | 1.27 |  |
| `MatchLabelKeysInPodTopologySpreadSelectorMerge` | `true` | beta | 1.34 |  |
| `MaxUnavailableStatefulSet` | `false` | alpha | 1.24 | 1.34 |
| `MaxUnavailableStatefulSet` | `true` | beta | 1.35.0 | 1.35.3 |
| `MaxUnavailableStatefulSet` | `false` | beta | 1.35.4 | 1.36 |
| `MaxUnavailableStatefulSet` | `true` | beta | 1.37 |  |
| `MemoryQoS` | `false` | alpha | 1.22 | 1.36 |
| `MemoryQoS` | `true` | beta | 1.37 |  |
| `MutablePVNodeAffinity` | `false` | alpha | 1.35 |  |
| `MutablePodResourcesForSuspendedJobs` | `false` | alpha | 1.35 | 1.35 |
| `MutablePodResourcesForSuspendedJobs` | `true` | beta | 1.36 |  |
| `MutableSchedulingDirectivesForSuspendedJobs` | `false` | alpha | 1.35 | 1.35 |
| `MutableSchedulingDirectivesForSuspendedJobs` | `true` | beta | 1.36 |  |
| `NativeHistograms` | `false` | alpha | 1.36 | 1.36 |
| `NativeHistograms` | `true` | beta | 1.37 |  |
| `NodeLifecycleConditions` | `false` | alpha | 1.37 |  |
| `NominatedNodeNameForExpectation` | `false` | alpha | 1.34 | 1.34 |
| `NominatedNodeNameForExpectation` | `true` | beta | 1.35 |  |
| `OpenAPIEnums` | `false` | alpha | 1.23 | 1.23 |
| `OpenAPIEnums` | `true` | beta | 1.24 |  |
| `OpportunisticBatching` | `true` | beta | 1.35 |  |
| `PersistentVolumeClaimUnusedSinceTime` | `false` | alpha | 1.36 | 1.36 |
| `PersistentVolumeClaimUnusedSinceTime` | `true` | beta | 1.37 |  |
| `PodAndContainerStatsFromCRI` | `false` | alpha | 1.23 | 1.36 |
| `PodAndContainerStatsFromCRI` | `false` | beta | 1.37 |  |
| `PodDeletionCost` | `false` | alpha | 1.21 | 1.21 |
| `PodDeletionCost` | `true` | beta | 1.22 |  |
| `PodGroupPreemptionPolicy` | `false` | alpha | 1.37 |  |
| `PodLevelResourceManagers` | `false` | alpha | 1.36 | 1.36 |
| `PodLevelResourceManagers` | `false` | beta | 1.37 |  |
| `PodLevelResources` | `false` | alpha | 1.32 | 1.33 |
| `PodLevelResources` | `true` | beta | 1.34 |  |
| `PodLogsQuerySplitStreams` | `false` | alpha | 1.32 |  |
| `PodTopologyLabelsAdmission` | `false` | alpha | 1.33 | 1.34 |
| `PodTopologyLabelsAdmission` | `true` | beta | 1.35 |  |
| `PodsAPI` | `false` | alpha | 1.36 | 1.36 |
| `PodsAPI` | `true` | beta | 1.37 |  |
| `PortForwardWebsockets` | `false` | alpha | 1.30 | 1.30 |
| `PortForwardWebsockets` | `true` | beta | 1.31 |  |
| `PreventStaticPodAPIReferences` | `true` | beta | 1.34 |  |
| `QOSReserved` | `false` | alpha | 1.11 |  |
| `ReduceDefaultCrashLoopBackOffDecay` | `false` | alpha | 1.33 |  |
| `ReloadKubeletServerCertificateFile` | `true` | beta | 1.31 |  |
| `RemoteRequestHeaderUID` | `false` | alpha | 1.32 |  |
| `ResourceHealthStatus` | `false` | alpha | 1.31 | 1.35 |
| `ResourceHealthStatus` | `true` | beta | 1.36 |  |
| `RestartAllContainersOnContainerExits` | `false` | alpha | 1.35 | 1.35 |
| `RestartAllContainersOnContainerExits` | `true` | beta | 1.36 |  |
| `RotateKubeletServerCertificate` | `false` | alpha | 1.7 | 1.11 |
| `RotateKubeletServerCertificate` | `true` | beta | 1.12 |  |
| `RuntimeClassInImageCriApi` | `false` | alpha | 1.29 |  |
| `SchedulerAsyncAPICalls` | `true` | beta | 1.34 |  |
| `SchedulerAsyncPreemption` | `false` | alpha | 1.32 | 1.32 |
| `SchedulerAsyncPreemption` | `true` | beta | 1.33 |  |
| `SchedulerPopFromBackoffQ` | `true` | beta | 1.33 |  |
| `ServiceAccountNodeAudienceRestriction` | `false` | beta | 1.32 | 1.32 |
| `ServiceAccountNodeAudienceRestriction` | `true` | beta | 1.33 |  |
| `ShardedListAndWatch` | `false` | alpha | 1.36 |  |
| `SizeBasedListCostEstimate` | `true` | beta | 1.34 |  |
| `StaleControllerConsistencyDaemonSet` | `true` | beta | 1.36 |  |
| `StaleControllerConsistencyJob` | `true` | beta | 1.36 |  |
| `StaleControllerConsistencyReplicaSet` | `true` | beta | 1.36 |  |
| `StaleControllerConsistencyStatefulSet` | `true` | beta | 1.36 |  |
| `StatefulSetRecreateStrategy` | `false` | alpha | 1.37 |  |
| `StorageCapacityScoring` | `false` | alpha | 1.33 | 1.36 |
| `StorageCapacityScoring` | `true` | beta | 1.37 |  |
| `StorageVersionAPI` | `false` | alpha | 1.20 |  |
| `StorageVersionHash` | `false` | alpha | 1.14 | 1.14 |
| `StorageVersionHash` | `true` | beta | 1.15 |  |
| `StorageVersionMigrator` | `false` | alpha | 1.30 | 1.34 |
| `StorageVersionMigrator` | `false` | beta | 1.35 |  |
| `StrictIPCIDRValidation` | `false` | alpha | 1.33 | 1.35 |
| `StrictIPCIDRValidation` | `true` | beta | 1.36 |  |
| `StructuredAuthenticationConfigurationEgressSelector` | `true` | beta | 1.34 |  |
| `StructuredAuthenticationConfigurationJWKSMetrics` | `true` | beta | 1.35 |  |
| `SupplementalGroupsPolicy` | `false` | alpha | 1.31 | 1.32 |
| `SupplementalGroupsPolicy` | `true` | beta | 1.33 |  |
| `SystemdWatchdog` | `true` | beta | 1.32 |  |
| `TaintTolerationComparisonOperators` | `false` | alpha | 1.35 |  |
| `TokenRequestServiceAccountUIDValidation` | `true` | beta | 1.34 |  |
| `TopologyAwareWorkloadScheduling` | `false` | alpha | 1.36 |  |
| `TopologyManagerPolicyAlphaOptions` | `false` | alpha | 1.26 |  |
| `TopologyManagerPolicyBetaOptions` | `false` | beta | 1.26 | 1.27 |
| `TopologyManagerPolicyBetaOptions` | `true` | beta | 1.28 |  |
| `TranslateStreamCloseWebsocketRequests` | `false` | alpha | 1.29 | 1.29 |
| `TranslateStreamCloseWebsocketRequests` | `true` | beta | 1.30 |  |
| `UnauthenticatedHTTP2DOSMitigation` | `false` | beta | 1.28 | 1.28 |
| `UnauthenticatedHTTP2DOSMitigation` | `true` | beta | 1.29 |  |
| `UnknownVersionInteroperabilityProxy` | `false` | alpha | 1.28 | 1.35 |
| `UnknownVersionInteroperabilityProxy` | `true` | beta | 1.36 |  |
| `UnlockWhileProcessingFIFO` | `true` | beta | 1.36 |  |
| `UserNamespacesHostNetworkSupport` | `false` | alpha | 1.35 |  |
| `VolumeBindMountOptions` | `false` | alpha | 1.37 |  |
| `VolumeLimitScaling` | `false` | alpha | 1.35 | 1.36 |
| `VolumeLimitScaling` | `true` | beta | 1.37 |  |
| `WatchList` | `false` | alpha | 1.27 | 1.31 |
| `WatchList` | `true` | beta | 1.32 | 1.32 |
| `WatchList` | `false` | beta | 1.33 | 1.33 |
| `WatchList` | `true` | beta | 1.34 |  |
| `WatchListClient` | `false` | beta | 1.30 | 1.34 |
| `WatchListClient` | `true` | beta | 1.35 |  |
| `WatchListCompression` | `true` | beta | 1.37 |  |
| `WindowsCPUAndMemoryAffinity` | `false` | alpha | 1.32 |  |
| `WindowsGracefulNodeShutdown` | `false` | alpha | 1.32 | 1.33 |
| `WindowsGracefulNodeShutdown` | `true` | beta | 1.34 |  |
| `WorkloadWithJob` | `false` | alpha | 1.36 |  |

### Feature gates for graduated or deprecated features

| Feature | Default | Stage | Since | Until |
|---|---|---|---|---|
| `APIServerTracing` | `false` | alpha | 1.22 | 1.26 |
| `APIServerTracing` | `true` | beta | 1.27 | 1.33 |
| `APIServerTracing` | `true` | stable | 1.34 |  |
| `AllowDNSOnlyNodeCSR` | `false` | deprecated | 1.31 |  |
| `AllowInsecureKubeletCertificateSigningRequests` | `false` | deprecated | 1.31 |  |
| `AllowServiceExternalIPs` | `true` | deprecated | 1.36 |  |
| `AnonymousAuthConfigurableEndpoints` | `false` | alpha | 1.31 | 1.31 |
| `AnonymousAuthConfigurableEndpoints` | `true` | beta | 1.32 | 1.33 |
| `AnonymousAuthConfigurableEndpoints` | `true` | stable | 1.34 |  |
| `AnyVolumeDataSource` | `false` | alpha | 1.18 | 1.23 |
| `AnyVolumeDataSource` | `true` | beta | 1.24 | 1.32 |
| `AnyVolumeDataSource` | `true` | stable | 1.33 |  |
| `AuthorizeNodeWithSelectors` | `false` | alpha | 1.31 | 1.31 |
| `AuthorizeNodeWithSelectors` | `true` | beta | 1.32 | 1.33 |
| `AuthorizeNodeWithSelectors` | `true` | stable | 1.34 |  |
| `AuthorizeWithSelectors` | `false` | alpha | 1.31 | 1.31 |
| `AuthorizeWithSelectors` | `true` | beta | 1.32 | 1.33 |
| `AuthorizeWithSelectors` | `true` | stable | 1.34 |  |
| `BtreeWatchCache` | `true` | beta | 1.32 | 1.32 |
| `BtreeWatchCache` | `true` | stable | 1.33 |  |
| `CPUManagerPolicyOptions` | `false` | alpha | 1.22 | 1.22 |
| `CPUManagerPolicyOptions` | `true` | beta | 1.23 | 1.32 |
| `CPUManagerPolicyOptions` | `true` | stable | 1.33 |  |
| `CRDValidationRatcheting` | `false` | alpha | 1.28 | 1.29 |
| `CRDValidationRatcheting` | `true` | beta | 1.30 | 1.32 |
| `CRDValidationRatcheting` | `true` | stable | 1.33 |  |
| `CSIServiceAccountTokenSecrets` | `true` | beta | 1.35 | 1.35 |
| `CSIServiceAccountTokenSecrets` | `true` | stable | 1.36 |  |
| `ChangeContainerStatusOnKubeletRestart` | `false` | deprecated | 1.35 |  |
| `ClusterTrustBundle` | `false` | alpha | 1.27 | 1.32 |
| `ClusterTrustBundle` | `false` | beta | 1.33 | 1.36 |
| `ClusterTrustBundle` | `true` | stable | 1.37 |  |
| `ClusterTrustBundleProjection` | `false` | alpha | 1.29 | 1.32 |
| `ClusterTrustBundleProjection` | `false` | beta | 1.33 | 1.36 |
| `ClusterTrustBundleProjection` | `true` | stable | 1.37 |  |
| `ConsistentListFromCache` | `false` | alpha | 1.28 | 1.30 |
| `ConsistentListFromCache` | `true` | beta | 1.31 | 1.33 |
| `ConsistentListFromCache` | `true` | stable | 1.34 |  |
| `CronJobsScheduledAnnotation` | `true` | beta | 1.28 | 1.31 |
| `CronJobsScheduledAnnotation` | `true` | stable | 1.32 |  |
| `CustomResourceFieldSelectors` | `false` | alpha | 1.30 | 1.30 |
| `CustomResourceFieldSelectors` | `true` | beta | 1.31 | 1.31 |
| `CustomResourceFieldSelectors` | `true` | stable | 1.32 |  |
| `DRAAdminAccess` | `false` | alpha | 1.32 | 1.33 |
| `DRAAdminAccess` | `true` | beta | 1.34 | 1.35 |
| `DRAAdminAccess` | `true` | stable | 1.36 |  |
| `DRADeviceTaintRules` | `false` | alpha | 1.35 | 1.35 |
| `DRADeviceTaintRules` | `false` | beta | 1.36 | 1.37 |
| `DRADeviceTaintRules` | `true` | stable | 1.37 |  |
| `DRADeviceTaints` | `false` | alpha | 1.33 | 1.35 |
| `DRADeviceTaints` | `true` | beta | 1.36 | 1.37 |
| `DRADeviceTaints` | `true` | stable | 1.37 |  |
| `DRAExtendedResource` | `false` | alpha | 1.34 | 1.35 |
| `DRAExtendedResource` | `true` | beta | 1.36 | 1.36 |
| `DRAExtendedResource` | `true` | stable | 1.37 |  |
| `DRAPrioritizedList` | `false` | alpha | 1.33 | 1.33 |
| `DRAPrioritizedList` | `true` | beta | 1.34 | 1.35 |
| `DRAPrioritizedList` | `true` | stable | 1.36 |  |
| `DRAResourceClaimDeviceStatus` | `false` | alpha | 1.32 | 1.32 |
| `DRAResourceClaimDeviceStatus` | `true` | beta | 1.33 | 1.36 |
| `DRAResourceClaimDeviceStatus` | `true` | stable | 1.37 |  |
| `DeclarativeValidation` | `true` | beta | 1.33 | 1.35 |
| `DeclarativeValidation` | `true` | stable | 1.36 |  |
| `DeclarativeValidationTakeover` | `false` | beta | 1.33 | 1.35 |
| `DeclarativeValidationTakeover` | `false` | deprecated | 1.36 | 1.36 |
| `DeclarativeValidationTakeover` | `false` | deprecated | 1.37 |  |
| `DisableAllocatorDualWrite` | `false` | alpha | 1.31 | 1.32 |
| `DisableAllocatorDualWrite` | `false` | beta | 1.33 | 1.33 |
| `DisableAllocatorDualWrite` | `true` | stable | 1.34 |  |
| `DisableNodeKubeProxyVersion` | `false` | alpha | 1.29 | 1.30 |
| `DisableNodeKubeProxyVersion` | `true` | beta | 1.31.0 | 1.31.0 |
| `DisableNodeKubeProxyVersion` | `false` | deprecated | 1.31.1 |  |
| `DisableNodeKubeProxyVersion` | `false` | deprecated | 1.32 | 1.32 |
| `DisableNodeKubeProxyVersion` | `true` | deprecated | 1.33 | 1.35 |
| `DisableNodeKubeProxyVersion` | `true` | deprecated | 1.36 |  |
| `DynamicResourceAllocation` | `false` | alpha | 1.30 | 1.31 |
| `DynamicResourceAllocation` | `false` | beta | 1.32 | 1.33 |
| `DynamicResourceAllocation` | `true` | stable | 1.34 | 1.34 |
| `DynamicResourceAllocation` | `true` | stable | 1.35 |  |
| `ElasticIndexedJob` | `true` | beta | 1.27 | 1.30 |
| `ElasticIndexedJob` | `true` | stable | 1.31 |  |
| `ExecProbeTimeout` | `true` | stable | 1.20 |  |
| `ExternalServiceAccountTokenSigner` | `false` | alpha | 1.32 | 1.33 |
| `ExternalServiceAccountTokenSigner` | `true` | beta | 1.34 | 1.35 |
| `ExternalServiceAccountTokenSigner` | `true` | stable | 1.36 |  |
| `GitRepoVolumeDriver` | `false` | deprecated | 1.33 |  |
| `HPAConfigurableTolerance` | `false` | alpha | 1.33 | 1.34 |
| `HPAConfigurableTolerance` | `false` | beta | 1.35 | 1.36 |
| `HPAConfigurableTolerance` | `true` | stable | 1.37 |  |
| `HostnameOverride` | `false` | alpha | 1.34 | 1.34 |
| `HostnameOverride` | `true` | beta | 1.35 | 1.36 |
| `HostnameOverride` | `true` | stable | 1.37 |  |
| `ImageMaximumGCAge` | `false` | alpha | 1.29 | 1.29 |
| `ImageMaximumGCAge` | `true` | beta | 1.30 | 1.34 |
| `ImageMaximumGCAge` | `true` | stable | 1.35 |  |
| `ImageVolume` | `false` | alpha | 1.31 | 1.32 |
| `ImageVolume` | `false` | beta | 1.33 | 1.34 |
| `ImageVolume` | `true` | beta | 1.35 | 1.35 |
| `ImageVolume` | `true` | stable | 1.36 |  |
| `InPlacePodVerticalScaling` | `false` | alpha | 1.27 | 1.32 |
| `InPlacePodVerticalScaling` | `true` | beta | 1.33 | 1.34 |
| `InPlacePodVerticalScaling` | `true` | stable | 1.35 |  |
| `InPlacePodVerticalScalingAllocatedStatus` | `false` | alpha | 1.32 | 1.32 |
| `InPlacePodVerticalScalingAllocatedStatus` | `false` | deprecated | 1.33 |  |
| `InformerResourceVersion` | `false` | alpha | 1.30 | 1.34 |
| `InformerResourceVersion` | `true` | stable | 1.35 |  |
| `JobBackoffLimitPerIndex` | `false` | alpha | 1.28 | 1.28 |
| `JobBackoffLimitPerIndex` | `true` | beta | 1.29 | 1.32 |
| `JobBackoffLimitPerIndex` | `true` | stable | 1.33 |  |
| `JobManagedBy` | `false` | alpha | 1.30 | 1.31 |
| `JobManagedBy` | `true` | beta | 1.32 | 1.34 |
| `JobManagedBy` | `true` | stable | 1.35 |  |
| `JobPodReplacementPolicy` | `false` | alpha | 1.28 | 1.28 |
| `JobPodReplacementPolicy` | `true` | beta | 1.29 | 1.33 |
| `JobPodReplacementPolicy` | `true` | stable | 1.34 |  |
| `JobSuccessPolicy` | `false` | alpha | 1.30 | 1.30 |
| `JobSuccessPolicy` | `true` | beta | 1.31 | 1.32 |
| `JobSuccessPolicy` | `true` | stable | 1.33 |  |
| `KMSv1` | `true` | deprecated | 1.28 | 1.28 |
| `KMSv1` | `false` | deprecated | 1.29 |  |
| `KubeProxyIPVS` | `true` | deprecated | 1.37 |  |
| `KubeletCgroupDriverFromCRI` | `false` | alpha | 1.28 | 1.30 |
| `KubeletCgroupDriverFromCRI` | `true` | beta | 1.31 |  |
| `KubeletCgroupDriverFromCRI` | `true` | stable | 1.34 |  |
| `KubeletFineGrainedAuthz` | `false` | alpha | 1.32 | 1.32 |
| `KubeletFineGrainedAuthz` | `true` | beta | 1.33 | 1.35 |
| `KubeletFineGrainedAuthz` | `true` | stable | 1.36 |  |
| `KubeletPSI` | `false` | alpha | 1.33 | 1.33 |
| `KubeletPSI` | `true` | beta | 1.34 | 1.35 |
| `KubeletPSI` | `true` | stable | 1.36 |  |
| `KubeletPodResourcesDynamicResources` | `false` | alpha | 1.27 | 1.33 |
| `KubeletPodResourcesDynamicResources` | `true` | beta | 1.34 | 1.35 |
| `KubeletPodResourcesDynamicResources` | `true` | stable | 1.36 |  |
| `KubeletPodResourcesGet` | `false` | alpha | 1.27 | 1.33 |
| `KubeletPodResourcesGet` | `true` | beta | 1.34 | 1.35 |
| `KubeletPodResourcesGet` | `true` | stable | 1.36 |  |
| `KubeletTracing` | `false` | alpha | 1.25 | 1.26 |
| `KubeletTracing` | `true` | beta | 1.27 | 1.33 |
| `KubeletTracing` | `true` | stable | 1.34 |  |
| `LogarithmicScaleDown` | `false` | alpha | 1.21 | 1.21 |
| `LogarithmicScaleDown` | `true` | beta | 1.22 | 1.30 |
| `LogarithmicScaleDown` | `true` | stable | 1.31 |  |
| `MatchLabelKeysInPodAffinity` | `false` | alpha | 1.29 | 1.30 |
| `MatchLabelKeysInPodAffinity` | `true` | beta | 1.31 | 1.32 |
| `MatchLabelKeysInPodAffinity` | `true` | stable | 1.33 |  |
| `MemoryManager` | `false` | alpha | 1.21 | 1.21 |
| `MemoryManager` | `true` | beta | 1.22 | 1.31 |
| `MemoryManager` | `true` | stable | 1.32 |  |
| `MultiCIDRServiceAllocator` | `false` | alpha | 1.27 | 1.30 |
| `MultiCIDRServiceAllocator` | `false` | beta | 1.31 | 1.32 |
| `MultiCIDRServiceAllocator` | `true` | stable | 1.33 |  |
| `MutableCSINodeAllocatableCount` | `false` | alpha | 1.33 | 1.33 |
| `MutableCSINodeAllocatableCount` | `false` | beta | 1.34 | 1.34 |
| `MutableCSINodeAllocatableCount` | `true` | beta | 1.35 | 1.35 |
| `MutableCSINodeAllocatableCount` | `true` | stable | 1.36 |  |
| `MutatingAdmissionPolicy` | `false` | alpha | 1.30 | 1.33 |
| `MutatingAdmissionPolicy` | `false` | beta | 1.34 | 1.35 |
| `MutatingAdmissionPolicy` | `true` | stable | 1.36 |  |
| `NFTablesProxyMode` | `false` | alpha | 1.29 | 1.30 |
| `NFTablesProxyMode` | `true` | beta | 1.31 | 1.32 |
| `NFTablesProxyMode` | `true` | stable | 1.33 |  |
| `NodeDeclaredFeatures` | `false` | alpha | 1.35 | 1.35 |
| `NodeDeclaredFeatures` | `true` | beta | 1.36 | 1.36 |
| `NodeDeclaredFeatures` | `true` | stable | 1.37 |  |
| `NodeInclusionPolicyInPodTopologySpread` | `false` | alpha | 1.25 | 1.25 |
| `NodeInclusionPolicyInPodTopologySpread` | `true` | beta | 1.26 | 1.32 |
| `NodeInclusionPolicyInPodTopologySpread` | `true` | stable | 1.33 |  |
| `NodeLogQuery` | `false` | alpha | 1.27 | 1.29 |
| `NodeLogQuery` | `false` | beta | 1.30 | 1.35 |
| `NodeLogQuery` | `true` | stable | 1.36 |  |
| `NodeSwap` | `false` | alpha | 1.22 | 1.27 |
| `NodeSwap` | `false` | beta | 1.28 | 1.29 |
| `NodeSwap` | `true` | beta | 1.30 | 1.33 |
| `NodeSwap` | `true` | stable | 1.34 |  |
| `OrderedNamespaceDeletion` | `false` | beta | 1.30 | 1.32 |
| `OrderedNamespaceDeletion` | `true` | beta | 1.33 | 1.33 |
| `OrderedNamespaceDeletion` | `true` | stable | 1.34 |  |
| `PodCertificateRequest` | `false` | alpha | 1.34 | 1.34 |
| `PodCertificateRequest` | `false` | beta | 1.35 | 1.36 |
| `PodCertificateRequest` | `true` | stable | 1.37 |  |
| `PodIndexLabel` | `true` | beta | 1.28 | 1.31 |
| `PodIndexLabel` | `true` | stable | 1.32 |  |
| `PodLifecycleSleepAction` | `false` | alpha | 1.29 | 1.29 |
| `PodLifecycleSleepAction` | `true` | beta | 1.30 | 1.33 |
| `PodLifecycleSleepAction` | `true` | stable | 1.34 |  |
| `PodLifecycleSleepActionAllowZero` | `false` | alpha | 1.32 | 1.32 |
| `PodLifecycleSleepActionAllowZero` | `true` | beta | 1.33 | 1.33 |
| `PodLifecycleSleepActionAllowZero` | `true` | stable | 1.34 |  |
| `PodObservedGenerationTracking` | `false` | alpha | 1.33 | 1.33 |
| `PodObservedGenerationTracking` | `true` | beta | 1.34 | 1.34 |
| `PodObservedGenerationTracking` | `true` | stable | 1.35 |  |
| `PodReadyToStartContainersCondition` | `false` | alpha | 1.28 | 1.28 |
| `PodReadyToStartContainersCondition` | `true` | beta | 1.29 | 1.36 |
| `PodReadyToStartContainersCondition` | `true` | stable | 1.37 |  |
| `PodSchedulingReadiness` | `false` | alpha | 1.26 | 1.26 |
| `PodSchedulingReadiness` | `true` | beta | 1.27 | 1.29 |
| `PodSchedulingReadiness` | `true` | stable | 1.30 |  |
| `PreferSameTrafficDistribution` | `false` | alpha | 1.33 | 1.33 |
| `PreferSameTrafficDistribution` | `true` | beta | 1.34 | 1.34 |
| `PreferSameTrafficDistribution` | `true` | stable | 1.35 |  |
| `ProcMountType` | `false` | alpha | 1.12 | 1.30 |
| `ProcMountType` | `false` | beta | 1.31 | 1.32 |
| `ProcMountType` | `true` | beta | 1.33 | 1.35 |
| `ProcMountType` | `true` | stable | 1.36 |  |
| `RecoverVolumeExpansionFailure` | `false` | alpha | 1.23 | 1.31 |
| `RecoverVolumeExpansionFailure` | `true` | beta | 1.32 | 1.33 |
| `RecoverVolumeExpansionFailure` | `true` | stable | 1.34 |  |
| `RecursiveReadOnlyMounts` | `false` | alpha | 1.30 | 1.30 |
| `RecursiveReadOnlyMounts` | `true` | beta | 1.31 | 1.32 |
| `RecursiveReadOnlyMounts` | `true` | stable | 1.33 |  |
| `RelaxedEnvironmentVariableValidation` | `false` | alpha | 1.30 | 1.31 |
| `RelaxedEnvironmentVariableValidation` | `true` | beta | 1.32 | 1.33 |
| `RelaxedEnvironmentVariableValidation` | `true` | stable | 1.34 |  |
| `RelaxedServiceNameValidation` | `false` | alpha | 1.34 | 1.35 |
| `RelaxedServiceNameValidation` | `true` | beta | 1.36 | 1.36 |
| `RelaxedServiceNameValidation` | `true` | stable | 1.37 |  |
| `ResilientWatchCacheInitialization` | `true` | beta | 1.31 | 1.33 |
| `ResilientWatchCacheInitialization` | `true` | stable | 1.34 |  |
| `RetryGenerateName` | `false` | alpha | 1.30 | 1.30 |
| `RetryGenerateName` | `true` | beta | 1.31 | 1.31 |
| `RetryGenerateName` | `true` | stable | 1.32 |  |
| `SELinuxChangePolicy` | `false` | alpha | 1.32 | 1.32 |
| `SELinuxChangePolicy` | `true` | beta | 1.33 | 1.35 |
| `SELinuxChangePolicy` | `true` | stable | 1.36 |  |
| `SELinuxMount` | `false` | alpha | 1.30 | 1.32 |
| `SELinuxMount` | `false` | beta | 1.33 | 1.36 |
| `SELinuxMount` | `true` | stable | 1.37 |  |
| `SELinuxMountReadWriteOncePod` | `false` | alpha | 1.25 | 1.26 |
| `SELinuxMountReadWriteOncePod` | `false` | beta | 1.27 | 1.27 |
| `SELinuxMountReadWriteOncePod` | `true` | beta | 1.28 | 1.35 |
| `SELinuxMountReadWriteOncePod` | `true` | stable | 1.36 |  |
| `SchedulerQueueingHints` | `true` | beta | 1.28 | 1.28 |
| `SchedulerQueueingHints` | `false` | beta | 1.29 | 1.31 |
| `SchedulerQueueingHints` | `true` | beta | 1.32 | 1.33 |
| `SchedulerQueueingHints` | `true` | stable | 1.34 |  |
| `SeparateCacheWatchRPC` | `true` | beta | 1.28 | 1.32 |
| `SeparateCacheWatchRPC` | `false` | deprecated | 1.33 |  |
| `SeparateTaintEvictionController` | `true` | beta | 1.29 | 1.33 |
| `SeparateTaintEvictionController` | `true` | stable | 1.34 |  |
| `ServiceAccountTokenJTI` | `false` | alpha | 1.29 | 1.29 |
| `ServiceAccountTokenJTI` | `true` | beta | 1.30 | 1.31 |
| `ServiceAccountTokenJTI` | `true` | stable | 1.32 |  |
| `ServiceAccountTokenNodeBinding` | `false` | alpha | 1.29 | 1.30 |
| `ServiceAccountTokenNodeBinding` | `true` | beta | 1.31 | 1.32 |
| `ServiceAccountTokenNodeBinding` | `true` | stable | 1.33 |  |
| `ServiceAccountTokenNodeBindingValidation` | `false` | alpha | 1.29 | 1.29 |
| `ServiceAccountTokenNodeBindingValidation` | `true` | beta | 1.30 | 1.31 |
| `ServiceAccountTokenNodeBindingValidation` | `true` | stable | 1.32 |  |
| `ServiceAccountTokenPodNodeInfo` | `false` | alpha | 1.29 | 1.29 |
| `ServiceAccountTokenPodNodeInfo` | `true` | beta | 1.30 | 1.31 |
| `ServiceAccountTokenPodNodeInfo` | `true` | stable | 1.32 |  |
| `ServiceTrafficDistribution` | `false` | alpha | 1.30 | 1.30 |
| `ServiceTrafficDistribution` | `true` | beta | 1.31 | 1.32 |
| `ServiceTrafficDistribution` | `true` | stable | 1.33 |  |
| `SidecarContainers` | `false` | alpha | 1.28 | 1.28 |
| `SidecarContainers` | `true` | beta | 1.29 | 1.32 |
| `SidecarContainers` | `true` | stable | 1.33 |  |
| `StatefulSetAutoDeletePVC` | `false` | alpha | 1.23 | 1.26 |
| `StatefulSetAutoDeletePVC` | `true` | beta | 1.27 | 1.31 |
| `StatefulSetAutoDeletePVC` | `true` | stable | 1.32 |  |
| `StatefulSetStartOrdinal` | `false` | alpha | 1.26 | 1.26 |
| `StatefulSetStartOrdinal` | `true` | beta | 1.27 | 1.30 |
| `StatefulSetStartOrdinal` | `true` | stable | 1.31 |  |
| `StorageNamespaceIndex` | `true` | beta | 1.30 | 1.32 |
| `StorageNamespaceIndex` | `true` | deprecated | 1.33 |  |
| `StreamingCollectionEncodingToJSON` | `true` | beta | 1.33 | 1.33 |
| `StreamingCollectionEncodingToJSON` | `true` | stable | 1.34 |  |
| `StreamingCollectionEncodingToProtobuf` | `true` | alpha | 1.33 | 1.33 |
| `StreamingCollectionEncodingToProtobuf` | `true` | stable | 1.34 |  |
| `StrictCostEnforcementForVAP` | `false` | beta | 1.30 | 1.31 |
| `StrictCostEnforcementForVAP` | `true` | stable | 1.32 |  |
| `StrictCostEnforcementForWebhooks` | `false` | beta | 1.31 | 1.31 |
| `StrictCostEnforcementForWebhooks` | `true` | stable | 1.32 |  |
| `StructuredAuthenticationConfiguration` | `false` | alpha | 1.29 | 1.29 |
| `StructuredAuthenticationConfiguration` | `true` | beta | 1.30 | 1.33 |
| `StructuredAuthenticationConfiguration` | `true` | stable | 1.34 |  |
| `StructuredAuthorizationConfiguration` | `false` | alpha | 1.29 | 1.29 |
| `StructuredAuthorizationConfiguration` | `true` | beta | 1.30 | 1.31 |
| `StructuredAuthorizationConfiguration` | `true` | stable | 1.32 |  |
| `TopologyManagerPolicyOptions` | `false` | alpha | 1.26 | 1.27 |
| `TopologyManagerPolicyOptions` | `true` | beta | 1.28 | 1.31 |
| `TopologyManagerPolicyOptions` | `true` | stable | 1.32 |  |
| `UserNamespacesSupport` | `false` | alpha | 1.28 | 1.29 |
| `UserNamespacesSupport` | `false` | beta | 1.30 | 1.32 |
| `UserNamespacesSupport` | `true` | beta | 1.33 | 1.35 |
| `UserNamespacesSupport` | `true` | stable | 1.36 |  |
| `VolumeAttributesClass` | `false` | alpha | 1.29 | 1.30 |
| `VolumeAttributesClass` | `false` | beta | 1.31 | 1.33 |
| `VolumeAttributesClass` | `true` | stable | 1.34 | 1.35 |
| `VolumeAttributesClass` | `true` | stable | 1.36 |  |
| `WatchCacheInitializationPostStartHook` | `false` | beta | 1.31 | 1.35 |
| `WatchCacheInitializationPostStartHook` | `true` | beta | 1.36 | 1.36 |
| `WatchCacheInitializationPostStartHook` | `true` | stable | 1.37 |  |
| `WatchFromStorageWithoutResourceVersion` | `false` | beta | 1.30 | 1.32 |
| `WatchFromStorageWithoutResourceVersion` | `false` | deprecated | 1.33 |  |
| `WinDSR` | `false` | alpha | 1.14 | 1.32 |
| `WinDSR` | `true` | beta | 1.33 | 1.33 |
| `WinDSR` | `true` | stable | 1.34 |  |
| `WinOverlay` | `false` | alpha | 1.14 | 1.19 |
| `WinOverlay` | `true` | beta | 1.20 | 1.33 |
| `WinOverlay` | `true` | stable | 1.34 |  |
| `WindowsHostNetwork` | `true` | alpha | 1.26 | 1.32 |
| `WindowsHostNetwork` | `false` | deprecated | 1.33 |  |

## Using a feature

### Feature stages

A feature can be in *Alpha*, *Beta* or *GA* stage.
An *Alpha* feature means:

* Disabled by default.
* Might be buggy. Enabling the feature may expose bugs.
* Support for feature may be dropped at any time without notice.
* The API may change in incompatible ways in a later software release without notice.
* Recommended for use only in short-lived testing clusters, due to increased
  risk of bugs and lack of long-term support.

A *Beta* feature means:

* Usually enabled by default. Beta API groups are
  [disabled by default](https://github.com/kubernetes/enhancements/tree/master/keps/sig-architecture/3136-beta-apis-off-by-default).
* The feature is well tested. Enabling the feature is considered safe.
* Support for the overall feature will not be dropped, though details may change.
* The schema and/or semantics of objects may change in incompatible ways in a
  subsequent beta or stable release. When this happens, we will provide instructions
  for migrating to the next version. This may require deleting, editing, and
  re-creating API objects. The editing process may require some thought.
  This may require downtime for applications that rely on the feature.
* Recommended for only non-business-critical uses because of potential for
  incompatible changes in subsequent releases. If you have multiple clusters
  that can be upgraded independently, you may be able to relax this restriction.

**Note:**
Please do try *Beta* features and give feedback on them!
After they exit beta, it may not be practical for us to make more changes.

A *General Availability* (GA) feature is also referred to as a *stable* feature. It means:

* The feature is always enabled; you cannot disable it.
* The corresponding feature gate is no longer needed.
* Stable versions of features will appear in released software for many subsequent versions.

## List of feature gates {#feature-gates}

Each feature gate is designed for enabling/disabling a specific feature.

* **`APIResponseCompression`**: Compress the API responses for `LIST` or `GET` requests.
* **`APIServerIdentity`**: Assign each API server an ID in a cluster, using a [Lease](/concepts/architecture/leases.md).
* **`APIServerTracing`**: Add support for distributed tracing in the API server. See [Traces for Kubernetes System Components](/concepts/cluster-administration/system-traces.md) for more details.
* **`APIServerWebhookAuthenticationToken`**: Enables the `kube-apiserver` to issue short-lived, scoped ServiceAccount tokens for authenticating to [admission webhooks](/reference/access-authn-authz/extensible-admission-controllers.md#authenticating-to-admission-webhooks). These tokens are bound to a specific ValidatingWebhookConfiguration or MutatingWebhookConfiguration and are scoped to particular API groups via attestation claims in the [TokenRequest](https://kubernetes.io/docs/reference/kubernetes-api/authentication-resources/token-request-v1/) API.
* **`APIServingWithRoutine`**: This feature gate enables an API server performance improvement: the API server can use separate goroutines (lightweight threads managed by the Go runtime) to serve [**watch**](/reference/using-api/api-concepts.md#efficient-detection-of-changes) requests.
* **`AllowDNSOnlyNodeCSR`**: Allow kubelet to request a certificate without any Node IP available, only with DNS names.
* **`AllowInsecureKubeletCertificateSigningRequests`**: Disable node admission validation of [CertificateSigningRequests](/reference/access-authn-authz/certificate-signing-requests.md#certificate-signing-requests) for kubelet signers. Unless you disable this feature gate, Kubernetes enforces that new kubelet certificates have a `commonName` matching `system:node:$nodeName`.
* **`AllowParsingUserUIDFromCertAuth`**: When this feature is enabled, the subject name attribute `1.3.6.1.4.1.57683.2` in an X.509 certificate will be parsed as the user UID during certificate authentication.
* **`AllowServiceExternalIPs`**: If set, kube-proxy will continue to respect the deprecated `externalIPs` field in Services.
* **`AllowUnsafeMalformedObjectDeletion`**: Enables the cluster operator to identify corrupt resource(s) using the **list** operation, and introduces an option `ignoreStoreReadErrorWithClusterBreakingPotential` that the operator can set to perform unsafe and force **delete** operation of such corrupt resource(s) using the Kubernetes API.
* **`AnonymousAuthConfigurableEndpoints`**: Enable [configurable endpoints for anonymous auth](/reference/access-authn-authz/authentication.md#anonymous-authenticator-configuration) for the API server.
* **`AnyVolumeDataSource`**: Enable use of any custom resource as the `DataSource` of a [PVC](/reference/glossary/persistent-volume-claim.md).
* **`AtomicFIFO`**: A client-go implementation of a FIFO queue that uses atomic operations to ensure events that come in batches, such as those from a ListAndWatch call, are processed in a single chunk. This is in contrast to the previous implementation which would process these events one by one, potentially causing the internal cache to become temporarily inconsistent with the API server. This feature gate can be toggled in the kube-controller-manager and any client-go based controller.
* **`AtomicWriteVolumeUserFields`**: This feature gate exists in the Kubernetes API server and kubelet. Used from the kube-apiserver, it allows users to set the `user` and `defaultUser` fields across `configMap`, `secret`, `downwardAPI` and `projected` volumes. In kubelet, if the `user` or `defaultUser` fields are specified for a volume, it sets the owner UID of the volume's data files during file creation.
* **`AuthorizeNodeWithSelectors`**: Make the [Node authorizer](/reference/access-authn-authz/node.md) use fine-grained selector authorization.
* **`AuthorizePodWebsocketUpgradeCreatePermission`**: When the `AuthorizePodWebsocketUpgradeCreatePermission` feature gate is `true`, clients must be authorized to **create** Pod subresources even when triggering their creation using a WebSocket. The connection upgrade request occurs for each of the following subresources: `pods/exec`, `pods/attach`, and `pods/portforward`. This feature gate fixes a security gap caused by the protocol transition: while SPDY requests utilize HTTP `POST` (naturally aligning with the **create** RBAC permission), the WebSocket protocol requires an HTTP `GET` request for the handshake. To correct this defect, a synthetic RBAC check is now applied to ensure WebSocket upgrades strictly enforce the **create** permission, matching the existing SPDY security model. You may want to disable this feature gate if you have existing clients or custom tooling that rely on the previous behavior—specifically, if they connect via WebSockets but do *not* currently hold the **create** RBAC permission.
* **`AuthorizeWithSelectors`**: Allows authorization to use field and label selectors. Enables `fieldSelector` and `labelSelector` fields in the [SubjectAccessReview API](/reference/kubernetes-api/definitions/subject-access-review-v1-authorization.md), passes field and label selector information to [authorization webhooks](/reference/access-authn-authz/webhook.md), enables `fieldSelector` and `labelSelector` functions in the [authorizer CEL library](https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#AuthzSelectors), and enables checking `fieldSelector` and `labelSelector` fields in [authorization webhook `matchConditions`](/reference/access-authn-authz/authorization.md#using-configuration-file-for-authorization).
* **`BtreeWatchCache`**: When enabled, the API server will replace the legacy HashMap-based _watch cache_ with a BTree-based implementation. This replacement may bring performance improvements.
* **`CBORServingAndStorage`**: Enables CBOR as a [supported encoding for requests and responses](/reference/using-api/api-concepts.md#cbor-encoding), and as the preferred storage encoding for custom resources.
* **`CPUManagerPolicyAlphaOptions`**: This allows fine-tuning of CPUManager policies, experimental, Alpha-quality options This feature gate guards *a group* of CPUManager options whose quality level is alpha. This feature gate will never graduate to beta or stable.
* **`CPUManagerPolicyBetaOptions`**: This allows fine-tuning of CPUManager policies, experimental, Beta-quality options This feature gate guards *a group* of CPUManager options whose quality level is beta. This feature gate will never graduate to stable.
* **`CPUManagerPolicyOptions`**: Allow fine-tuning of CPUManager policies.
* **`CRDObservedGenerationTracking`**: Allows for the observed generation to be tracked in CRD conditions. Setting to false will make it so CRD conditions will have the observed generation wiped.
* **`CRDValidationRatcheting`**: Enable updates to custom resources to contain violations of their OpenAPI schema if the offending portions of the resource update did not change. See [Validation Ratcheting](/tasks/extend-kubernetes/custom-resources/custom-resource-definitions.md#validation-ratcheting) for more details.
* **`CRIListStreaming`**: Enable streaming RPCs for CRI list operations (`ListContainers`, `ListPodSandbox`, `ListImages`). When enabled, the kubelet uses server-side streaming RPCs (e.g., `StreamContainers`, `StreamPodSandboxes`) that allow the container runtime to divide results across multiple response messages, bypassing the 16 MiB gRPC message size limit. This allows listing containers on nodes with thousands of containers without failures. If the container runtime does not support streaming RPCs, the kubelet falls back to unary RPCs.
* **`CSIServiceAccountTokenSecrets`**: Enables CSI drivers to opt-in for receiving service account tokens from kubelet through the dedicated secrets field in NodePublishVolumeRequest instead of the volume_context field.
* **`CSIVolumeHealth`**: Enable support for CSI volume health monitoring on node.
* **`ChangeContainerStatusOnKubeletRestart`**: Enable legacy writes to update container `ready` status after the kubelet detects a [restart](/concepts/workloads/pods/pod-lifecycle.md#kubelet-restarts). This feature gate was introduced to allow you revert the behavior to a previously used default. If you are satisfied with the default behavior, you do not need to enable this feature gate.
* **`ClearingNominatedNodeNameAfterBinding`**: Enable clearing `.status.nominatedNodeName` whenever Pods are bound to nodes.
* **`CloudControllerManagerWatchBasedRoutesReconciliation`**: Enables a watch-based route reconciliation mechanism (rather than reconciling at a fixed interval) within the cloud-controller-manager library.
* **`CloudControllerManagerWebhook`**: Enable webhooks in cloud controller manager.
* **`ClusterTrustBundle`**: This feature gate exists in the Kubernetes API server and the controller manager. Used from the kube-apiserver, it enables ClusterTrustBundle support. In the Kubernetes controller manager, it is used to control publishing of a ClusterTrustBundle for the `kubernetes.io/kube-apiserver-serving` signer.
* **`ClusterTrustBundleProjection`**: [`clusterTrustBundle` projected volume sources](/concepts/storage/projected-volumes.md#clustertrustbundle).
* **`ComponentFlagz`**: Enables the component's flagz endpoint. See [zpages](/reference/instrumentation/zpages.md) for more information.
* **`ComponentStatusz`**: Enables the component's statusz endpoint. See [zpages](/reference/instrumentation/zpages.md) for more information.
* **`CompositePodGroup`**: Enable hierarchical gang scheduling for [CompositePodGroups](/concepts/workloads/compositepodgroup-api.md) and PodGroups. Enabling the `CompositePodGroup` feature gate requires that the `GenericWorkload` and `TopologyAwareWorkloadScheduling` feature gates are enabled as well.
* **`ConcurrentWatchObjectDecode`**: Enable concurrent watch object decoding. This is to avoid starving the API server's watch cache when a conversion webhook is installed.
* **`ConsistentListFromCache`**: Enhance Kubernetes API server performance by serving consistent **list** requests directly from its watch cache, improving scalability and response times. To consistent list from cache Kubernetes requires a newer etcd version (v3.4.31+ or v3.5.13+), that includes fixes to watch progress request feature. If older etcd version is provided Kubernetes will automatically detect it and fallback to serving consistent reads from etcd. Progress notifications ensure watch cache is consistent with etcd while reducing the need for resource-intensive quorum reads from etcd. See the Kubernetes documentation on [Semantics for **get** and **list**](/reference/using-api/api-concepts.md#semantics-for-get-and-list) for more details.
* **`ConstrainedImpersonation`**: Enables impersonation that is constrained to specific requests instead of being all or nothing.
* **`ContainerCheckpoint`**: Enables the kubelet `checkpoint` API. See [Kubelet Checkpoint API](/reference/node/kubelet-checkpoint-api.md) for more details.
* **`ContainerRestartRules`**: Enables the ability to configure container-level restart policy and restart rules. See [Container Restart Policy and Rules](/concepts/workloads/pods/pod-lifecycle.md#container-restart-rules) for more details.
* **`ContainerStopSignals`**: Enables usage of the StopSignal lifecycle for containers for configuring custom stop signals using which the containers would be stopped.
* **`ContextualLogging`**: Enables extra details in log output of Kubernetes components that support contextual logging.
* **`ControllerManagerReleaseLeaderElectionLockOnExit`**: Enables the `kube-controller-manager` to actively release its leader election lock during leader transitions, rather than waiting for the lock's TTL to expire. This allows a new leader to be elected more quickly.
* **`CoordinatedLeaderElection`**: Enables the behaviors supporting the LeaseCandidate API, and also enables coordinated leader election for the Kubernetes control plane, deterministically.
* **`CronJobsScheduledAnnotation`**: Set the scheduled job time as an [annotation](/reference/glossary/annotation.md) on Jobs that were created on behalf of a CronJob.
* **`CrossNamespaceVolumeDataSource`**: Enable the usage of cross namespace volume data source to allow you to specify a source namespace in the `dataSourceRef` field of a PersistentVolumeClaim.
* **`CustomCPUCFSQuotaPeriod`**: Enable nodes to change `cpuCFSQuotaPeriod` in [kubelet config](/tasks/administer-cluster/kubelet-config-file.md).
* **`CustomResourceFieldSelectors`**: Enable `selectableFields` in the [CustomResourceDefinition](/reference/glossary/customresourcedefinition.md) API to allow filtering of custom resource **list**, **watch** and **deletecollection** requests.
* **`DRAAdminAccess`**: Enables support for requesting [admin access](/concepts/resource-management/dynamic-resource-allocation/dra-api.md#admin-access) in a ResourceClaim or a ResourceClaimTemplate. Admin access grants access to in-use devices and may enable additional permissions when making the device available in a container. Starting with Kubernetes v1.33, only users authorized to create ResourceClaim or ResourceClaimTemplate objects in namespaces labeled with `resource.kubernetes.io/admin-access: "true"` (case-sensitive) can use the `adminAccess` field. This ensures that non-admin users cannot misuse the feature. Starting with Kubernetes v1.34, this label has been updated to `resource.kubernetes.io/admin-access: "true"`.
* **`DRAConsumableCapacity`**: Enables device sharing across multiple ResourceClaims or requests. Additionally, if a device supports sharing, its resource (capacity) can be managed through a defined sharing policy.
* **`DRADerivedAttributes`**: Enables `derivedAttributes` in Dynamic Resource Allocation (DRA), letting ResourceClaim and ResourceClaimTemplate authors compute virtual device attributes with per-device CEL expressions, for use in `matchAttribute` and `distinctAttribute` constraints. For more information, see [Derived attributes](/concepts/resource-management/dynamic-resource-allocation/dra-api.md#derived-attributes) in the DRA API Objects documentation.
* **`DRADeviceBindingConditions`**: Enables support for DeviceBindingConditions in the DRA related fields. This allows for thorough device readiness checks and attachment processes before Bind phase.
* **`DRADeviceCompatibilityGroups`**: Enables support for [device compatibility groups](/concepts/resource-management/dynamic-resource-allocation/dra-features.md#device-compatibility-groups) in DRA. Drivers can declare opaque compatibility groups on each `consumesCounters` entry of a device in a ResourceSlice, and the scheduler only co-allocates devices drawing from the same counter set when their declared groups intersect. Requires `DRAPartitionableDevices` to be enabled.
* **`DRADeviceTaintRules`**: Enables support for [tainting devices through DeviceTaintRule objects](/concepts/resource-management/dynamic-resource-allocation/device-taints.md#device-taints-and-tolerations) when using dynamic resource allocation to manage devices. This feature gate has no effect unless you also enable the `DRADeviceTaints` feature gate.
* **`DRADeviceTaints`**: Enables support for [tainting devices and selectively tolerating those taints](/concepts/resource-management/dynamic-resource-allocation/device-taints.md#device-taints-and-tolerations) when using dynamic resource allocation to manage devices.
* **`DRAExtendedResource`**: Enables support for the [Extended Resource allocation by DRA](/concepts/configuration/manage-resources-containers.md#extended-resources-allocation-by-dra) feature. It makes it possible to specify an extended resource name in a DeviceClass.
* **`DRAListTypeAttributes`**: Enables list-type attribute fields (`bools`, `ints`, `strings`, `versions`) for devices in `ResourceSlice`, allowing a device to advertise multiple values for a single attribute. When enabled, `matchAttribute` uses set-intersection semantics (the sets of attribute values across all selected devices must have a non-empty intersection), and `distinctAttribute` uses pairwise-disjoint semantics (the sets must share no values). Scalar attributes remain backward-compatible, treated as singleton sets. Also adds the `includes()` helper function to CEL device selector expressions, which works on both scalar and list-type attributes. For more information, see [List type attributes](/concepts/resource-management/dynamic-resource-allocation/dra-api.md#list-type-attributes) in the Dynamic Resource Allocation documentation.
* **`DRANodeAllocatableResources`**: Enables the `kube-scheduler` to incorporate node allocatable resources (such as CPU, memory, and hugepages) managed by Dynamic Resource Allocation (DRA) into its standard node resource accounting. When enabled, DRA drivers can use the `nodeAllocatableResources` field on `ResourceSlice` devices to specify how their devices consume node allocatable resources. This field supports two different use cases: - `mapping`: For drivers that directly provide a native node resource (e.g., a CPU or Memory DRA driver). It supports scaling capacities or device counts. - `overhead`: For devices that require auxiliary node dependencies (e.g., an accelerator that consumes host memory). It supports per-pod or per-container costs. This allows the scheduler to combine these DRA allocations with standard Pod requests to prevent node over-subscription during Pod admission. It also exposes the `status.nodeAllocatableResourceClaimStatuses` field on the Pod API to track the resulting resource allocations. The `kubelet` consumes this to update Pod and container cgroup settings and adjust OOM scores. For more information, see [Node Allocatable Resources](/concepts/resource-management/dynamic-resource-allocation/how-dra-works.md#node-allocatable-resources) in the Dynamic Resource Allocation documentation.
* **`DRAOptionalNodeOperations`**: Enables support for optional node-local operations in Dynamic Resource Allocation (DRA). This allows drivers to declare that specific node operations (`NodePrepareResources` and/or `NodeUnprepareResources`) can be skipped for their devices, enabling the `kubelet` to bypass unnecessary gRPC calls. For more information, see [Optional node operations](/concepts/resource-management/dynamic-resource-allocation/dra-features.md#optional-node-operations) in the Dynamic Resource Allocation documentation.
* **`DRAPartitionableDevices`**: Enables support for requesting [Partitionable Devices](/concepts/resource-management/dynamic-resource-allocation/dra-features.md#partitionable-devices) for DRA. This lets drivers advertise multiple devices that maps to the same resources of a physical device.
* **`DRAPartitionableDevicesType`**: Enables the `PartitionTypeAttribute` field on ResourceSlices, which opts a [partitionable](/concepts/resource-management/dynamic-resource-allocation/dra-features.md#partitionable-devices) resource pool into the typed partition summary view of [ResourcePoolStatusRequest](/concepts/resource-management/dynamic-resource-allocation/dra-observability.md#resource-pool-partition-summary). The field names a device attribute (such as a MIG profile) whose value groups each partition type, so that a ResourcePoolStatusRequest can report how many devices of each partition type are still allocatable. This builds on the [`DRAPartitionableDevices`](/reference/command-line-tools-reference/feature-gates.md#DRAPartitionableDevices) and [`DRAResourcePoolStatus`](/reference/command-line-tools-reference/feature-gates.md#DRAResourcePoolStatus) feature gates, both of which must also be enabled.
* **`DRAPrioritizedList`**: Allows specifying a prioritized list of alternative devices that can be allocated to a request in a claim if the preferred alternative is not available.
* **`DRAResourceClaimDeviceStatus`**: Enables support the ResourceClaim.status.devices field and for setting this status from DRA drivers. It requires the `DynamicResourceAllocation` feature gate to be enabled.
* **`DRAResourceClaimGranularStatusAuthorization`**: Enables support for granular authorization of ResourceClaim status updates. This feature requires additional fine-grained access permissions when modifying specific fields within ResourceClaim status objects.
* **`DRAResourcePoolStatus`**: Enables the ResourcePoolStatusRequest API for querying the [availability of devices in DRA resource pools](/concepts/resource-management/dynamic-resource-allocation/dra-observability.md#resource-pool-status). When enabled, users can create ResourcePoolStatusRequest objects to get a point-in-time snapshot of device availability (total, allocated, available, and unavailable devices) for a specific driver and optionally a specific pool. A controller in kube-controller-manager processes these one-time requests and populates the status with pool information.
* **`DRASchedulerFilterTimeout`**: Enables aborting the per-node filter operation in the scheduler after a certain time (10 seconds by default, configurable in the DynamicResources scheduler plugin configuration).
* **`DRAWorkloadResourceClaims`**: Enables PodGroup resources from the [Workload API](/concepts/workloads/workload-api.md) to make requests for devices through [Dynamic Resource Allocation](/concepts/resource-management/dynamic-resource-allocation.md) that can be shared by their member Pods.
* **`DeclarativeValidation`**: Reports differences between declarative validation of in-tree Kubernetes APIs and the equivalent hand-written validation. When enabled, rules marked `+k8s:alpha` or `+k8s:beta` run alongside the hand-written validation, and the API server logs any discrepancy and counts it in the `declarative_validation_mismatch_total` metric. This gate controls only reporting, not which result the API server returns. Enforcement is: - No prefix: always enforced. - `+k8s:beta`: enforced when the [`DeclarativeValidationBeta` feature gate](/reference/command-line-tools-reference/feature-gates.md#DeclarativeValidationBeta) is enabled (the default). - `+k8s:alpha`: never enforced. This feature gate only operates on the `kube-apiserver` component.
* **`DeclarativeValidationBeta`**: The global safety switch for Beta validation rules (`+k8s:beta`). Disable it to revert those rules to shadow mode if regressions are found. In shadow mode, declarative validation still runs and mismatches are logged and metered, but they do not reject requests; hand-written validation stays authoritative. Enforcement for resources using `WithDeclarativeEnforcement()`: - No prefix: always enforced, regardless of this gate. - `+k8s:beta`: enforced when this gate is enabled (the default), otherwise shadowed. - `+k8s:alpha`: always shadowed. This gate takes effect independently of the [`DeclarativeValidation` feature gate](/reference/command-line-tools-reference/feature-gates.md#DeclarativeValidation), which controls only mismatch reporting.
* **`DeclarativeValidationTakeover`**: Deprecated: replaced by the [`DeclarativeValidationBeta` feature gate](/reference/command-line-tools-reference/feature-gates.md#DeclarativeValidationBeta). Before Kubernetes v1.36, this gate made declarative validation errors authoritative, replacing the hand-written errors for rules that had a declarative implementation. The API server no longer honors this gate. Use `DeclarativeValidationBeta` instead to control whether `+k8s:beta` rules are enforced. This feature gate only operates on the `kube-apiserver` component.
* **`DefaultPodSysctls`**: Enables the `defaultPodSysctls` field in [KubeletConfiguration](/reference/config-api/kubelet-config.v1beta1.md), allowing Node administrators to specify a default set of namespaced kernel parameters (sysctls) that the `kubelet` applies to all Pods on the Node. See [Setting Sysctls for All Pods](/tasks/administer-cluster/sysctl-cluster.md#setting-sysctls-for-all-pods) for more details.
* **`DeploymentReplicaSetTerminatingReplicas`**: Enables a new status field `.status.terminatingReplicas` in Deployments and ReplicaSets to allow tracking of terminating pods.
* **`DetectCacheInconsistency`**: Enable cache inconsistency detection in the API server.
* **`DisableAllocatorDualWrite`**: You can enable the `MultiCIDRServiceAllocator` feature gate. The API server supports migration from the old bitmap ClusterIP allocators to the new IPAddress allocators. The API server performs a dual-write on both allocators. This feature gate disables the dual write on the new Cluster IP allocators; you can enable this feature gate if you have completed the relevant stage of the migration.
* **`DisableCPUQuotaWithExclusiveCPUs`**: When the feature gate `DisableCPUQuotaWithExclusiveCPUs` is enabled (the default), then Kubernetes does **not** enforce CPU quota for Pods that use the [Guaranteed](/concepts/workloads/pods/pod-qos.md#guaranteed) [QoS class](/reference/glossary/qos-class.md). You can disable the `DisableCPUQuotaWithExclusiveCPUs` feature gate to restore the legacy behavior.
* **`DisableNodeKubeProxyVersion`**: Disable setting the `kubeProxyVersion` field of the Node.
* **`DynamicResourceAllocation`**: Enables support for resources with custom parameters and a lifecycle that is independent of a Pod. Allocation of resources is handled by the Kubernetes scheduler based on "structured parameters".
* **`ElasticIndexedJob`**: Enables Indexed Jobs to be scaled up or down by mutating both `spec.completions` and `spec.parallelism` together such that `spec.completions == spec.parallelism`. See docs on [elastic Indexed Jobs](/concepts/workloads/controllers/job.md#elastic-indexed-jobs) for more details.
* **`EmptyDirVolumeMode`**: Enables setting Unix permission bits on `emptyDir` volume directories using the `mode` field in `emptyDir` volume sources. When enabled, users can specify a value between `0000` and `01777` (octal) to control the directory permissions at creation time. If `mode` is not specified, the default `0777` behavior is preserved.
* **`EnvFiles`**: Support defining container's Environment Variable Values via File. See [Define Environment Variable Values Using An Init Container](/tasks/inject-data-application/define-environment-variable-via-file.md) for more details.
* **`EtcdRangeStream`**: Enables the kube-apiserver to use etcd's `RangeStream` RPC to stream large list responses from etcd, instead of fetching them in a single `Range` response. This reduces memory spikes in both etcd and the kube-apiserver when serving large `LIST` requests.
* **`EventedPLEG`**: Enable support for the kubelet to receive container life cycle events from the [container runtime](/reference/glossary/container-runtime.md) via an extension to [CRI](/reference/glossary/cri.md). (PLEG is an abbreviation for “Pod lifecycle event generator”). For this feature to be useful, you also need to enable support for container lifecycle events in each container runtime running in your cluster. If the container runtime does not announce support for container lifecycle events then the kubelet automatically switches to the legacy generic PLEG mechanism, even if you have this feature gate enabled.
* **`ExcludeAdmissionWebhookVirtualResources`**: Exclude non-persisted (virtual) authentication and authorization resources, such as `TokenReview` and `SubjectAccessReview`, from admission webhooks. This matches the set of resources that ValidatingAdmissionPolicy and MutatingAdmissionPolicy already exclude, and prevents a misbehaving webhook from blocking the cluster's own authentication and authorization requests. Disable this feature gate to restore the previous behavior of dispatching admission webhooks for these resources.
* **`ExecProbeTimeout`**: Ensure kubelet respects exec probe timeouts. This feature gate exists in case any of your existing workloads depend on a now-corrected fault where Kubernetes ignored exec probe timeouts. See [readiness probes](/tasks/configure-pod-container/configure-liveness-readiness-startup-probes.md#configure-probes).
* **`ExtendWebSocketsToKubelet`**: When ExtendWebSocketsToKubelet is enabled and a kubelet node advertises support, exec/attach/portforward streams are proxied directly to the kubelet rather than being translated or tunneled at the API server. Critically, the same stream translation and tunneling handlers used at the API server are now set up identically at the kubelet — the logic is simply moved closer to the container runtime. This feature depends on NodeDeclaredFeatures graduating to beta so that kubelet capability advertisement is reliable in production clusters.
* **`ExternalServiceAccountTokenSigner`**: Enable setting `--service-account-signing-endpoint` to make the kube-apiserver use [external signer](/reference/access-authn-authz/service-accounts-admin.md#external-serviceaccount-token-signing-and-key-management) for token signing and token verifying key management.
* **`GRPCContainerProbeTLS`**: Enables TLS support for gRPC container probes. When enabled, you can add the `mode` field to the `grpc` field in gRPC probes. Setting `mode: TLS` on a liveness, readiness, or startup probe causes the kubelet to connect over TLS (with `InsecureSkipVerify`). See [Configure Liveness, Readiness and Startup Probes](/tasks/configure-pod-container/configure-liveness-readiness-startup-probes.md#grpc-probe-tls).
* **`GenericWorkload`**: Enables support for the [Workload API](/concepts/workloads/workload-api.md) and [PodGroup API](/concepts/workloads/podgroup-api.md) to express scheduling requirements at the workload level. When enabled, Pods can reference a specific PodGroup to influence the way that they are scheduled. Starting in Kubernetes v1.37, this feature gate also encompasses [gang scheduling](/concepts/scheduling-eviction/gang-scheduling.md), and [workload-aware preemption](/concepts/scheduling-eviction/workload-aware-preemption.md).
* **`GitRepoVolumeDriver`**: This controls if the `gitRepo` volume plugin is supported or not. The `gitRepo` volume plugin is disabled by default starting v1.33 release. This provides a way for users to enable it.
* **`GracefulNodeShutdown`**: Enables support for graceful shutdown in kubelet. During a system shutdown, kubelet will attempt to detect the shutdown event and gracefully terminate pods running on the node. See [Graceful Node Shutdown](/concepts/cluster-administration/node-shutdown.md#graceful-node-shutdown) for more details.
* **`GracefulNodeShutdownBasedOnPodPriority`**: Enables the kubelet to check Pod priorities when shutting down a node gracefully.
* **`H2CContainerProbe`**: Enables HTTP/2 cleartext (h2c) support for HTTP container probes. When enabled, you can add the `protocol` field to the `httpGet` field in HTTP probes. Setting `protocol: HTTP2` on a liveness, readiness, or startup probe causes the `kubelet` to probe over h2c instead of HTTP/1.1. See [Configure Liveness, Readiness and Startup Probes](/tasks/configure-pod-container/configure-liveness-readiness-startup-probes.md#use-h2c-with-http-probes).
* **`HPAConfigurableTolerance`**: Enables setting a [tolerance threshold](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale#tolerance) for HorizontalPodAutoscaler metrics.
* **`HPAScaleToZero`**: Enables setting `minReplicas` to 0 for `HorizontalPodAutoscaler` resources when using custom or external metrics.
* **`HostnameOverride`**: Allows setting any FQDN as the pod's hostname.
* **`HugepageAwareEviction`**: Subtracts hugepage capacity from `memory.available` so the kubelet's eviction signal reflects actual regular-memory availability. Without this gate, hugepage reservations inflate `AvailableBytes`, delaying eviction and causing OOM kills on nodes with hugepages configured.
* **`ImageMaximumGCAge`**: Enables the kubelet configuration field `imageMaximumGCAge`, allowing an administrator to specify the age after which an image will be garbage collected.
* **`ImageVolume`**: Allow using the [`image`](/concepts/storage/volumes.md) volume source in a Pod. This volume source lets you mount a container image as a read-only volume.
* **`ImageVolumeWithDigest`**: For each [`image` volume](/concepts/storage/volumes.md#image) in a Pod, image digest as part of the pod's status.
* **`InOrderInformers`**: Force the informers to deliver watch stream events in order instead of out of order.
* **`InPlacePodLevelResourcesVerticalScaling`**: Enables the in-place vertical scaling of resources for a Pod (For example, changing a running Pod's pod-level CPU or memory requests/limits without needing to restart it). For details, see the documentation on [In-place Pod-level Resources Vertical Scaling](/tasks/configure-pod-container/resize-pod-resources.md).
* **`InPlacePodVerticalScaling`**: Enables in-place Pod vertical scaling.
* **`InPlacePodVerticalScalingAllocatedStatus`**: Enables the `allocatedResources` field in the container status. This feature requires the `InPlacePodVerticalScaling` gate be enabled as well.
* **`InPlacePodVerticalScalingExclusiveCPUs`**: Enable resource resizing for containers in Guaranteed pods with integer CPU requests. It applies only in nodes with `InPlacePodVerticalScaling` and `CPUManager` features enabled, and the CPUManager policy set to `static`.
* **`InPlacePodVerticalScalingExclusiveMemory`**: Allow resource resize for containers in Guaranteed Pods when the memory manager policy is set to `"Static"`. Applies only to nodes with `InPlacePodVerticalScaling` and memory manager features enabled.
* **`InPlacePodVerticalScalingMemoryBackedVolumes`**: Enables in-place vertical scaling for memory-backed `emptyDir` volume size limits.
* **`InPlacePodVerticalScalingSchedulerPreemption`**: Enables scheduler preemption of lower-priority pods to fulfill deferred in-place pod resize requests.
* **`InformerResourceVersion`**: Allow clients to use the `LastSyncResourceVersion()` call on informers, enabling them to perform actions based on the current resource version. When disabled, `LastSyncResourceVersion()` succeeds but returns an empty string. Used by kube-controller-manager for StorageVersionMigration.
* **`JobBackoffLimitPerIndex`**: Allows specifying the maximal number of pod retries per index in Indexed jobs.
* **`JobManagedBy`**: Allows to delegate reconciliation of a Job object to an external controller.
* **`JobPodReplacementPolicy`**: Allows you to specify pod replacement for terminating pods in a [Job](/concepts/workloads/controllers/job.md)
* **`JobSuccessPolicy`**: Allow users to specify when a Job can be declared as succeeded based on the set of succeeded pods.
* **`KMSv1`**: Enables KMS v1 API for encryption at rest. See [Using a KMS Provider for data encryption](/tasks/administer-cluster/kms-provider.md) for more details.
* **`KubeProxyIPVS`**: Enable support for the deprecated [`ipvs` proxy mode](/reference/networking/virtual-ips.md#proxy-mode-ipvs) in kube-proxy.
* **`KubeProxyNFTablesLocalhostNodePorts`**: Enables localhost NodePort Service proxying with the nftables mode of kube-proxy.
* **`KubeletCgroupDriverFromCRI`**: Enable detection of the kubelet cgroup driver configuration option from the [CRI](/reference/glossary/cri.md). This feature gate is now on for all clusters. However, it only works on nodes where there is a CRI container runtime that supports the `RuntimeConfig` CRI call. If the CRI supports this feature, the kubelet ignores the `cgroupDriver` configuration setting (or deprecated `--cgroup-driver` command line argument). If the container runtime doesn't support it, the kubelet falls back to using the driver configured using the `cgroupDriver` configuration setting. The kubelet will stop falling back to this configuration in Kubernetes 1.36. Thus, users must upgrade their CRI container runtime to a version that supports the `RuntimeConfig` CRI call by then. Admins can use the metric `kubelet_cri_losing_support` to see if there are any nodes in their cluster that will lose support in 1.36. The following CRI versions support this CRI call: * containerd: Support was added in v2.0.0 * CRI-O: Support was added in v1.28.0
* **`KubeletCrashLoopBackOffMax`**: Enables support for configurable per-node backoff maximums for restarting containers in the `CrashLoopBackOff` state. For more details, check the `crashLoopBackOff.maxContainerRestartPeriod` field in the [kubelet config file](/reference/config-api/kubelet-config.v1beta1.md).
* **`KubeletEnsureSecretPulledImages`**: Ensure that pods requesting an image are authorized to access the image with the provided credentials when the image is already present on the node. See [Ensure Image Pull Credential Verification](/concepts/containers/images.md#ensureimagepullcredentialverification).
* **`KubeletFineGrainedAuthz`**: Enable [fine-grained authorization](/reference/access-authn-authz/kubelet-authn-authz.md#fine-grained-authorization) for the kubelet's HTTP(s) API.
* **`KubeletInUserNamespace`**: Enables support for running kubelet in a [user namespace](/reference/glossary/userns.md). See [Running Kubernetes Node Components as a Non-root User](/tasks/administer-cluster/kubelet-in-userns.md).
* **`KubeletPSI`**: Enable kubelet to surface Pressure Stall Information (PSI) metrics in the Summary API and Prometheus metrics.
* **`KubeletPodResourcesDynamicResources`**: Extend the kubelet's [pod resources monitoring gRPC API](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md) endpoints List and Get to include resources allocated in ResourceClaims via [Dynamic Resource Allocation](/concepts/resource-management/dynamic-resource-allocation.md). Below is an example of GPU metrics to show how this API is consumed by [NVIDIA dcgm-exporter](https://github.com/NVIDIA/dcgm-exporter) to collect per pod GPU metrics allocated by [NVIDIA DRA driver](https://github.com/NVIDIA/k8s-dra-driver-gpu): ``` DCGM_FI_PROF_PCIE_RX_BYTES{gpu="0",UUID="GPU-a4f34abc-7715-3560-dcea-7238b9611a45",pci_bus_id="00000009:01:00.0",device="nvidia0",modelName="NVIDIA GH200 96GB HBM3",Hostname="sc-starwars-xxxx",container="ctr",dra_claim_name="single-gpu",dra_claim_namespace="gpu-test3",dra_device_name="gpu-0",dra_driver_name="gpu.nvidia.com",dra_pool_name="sc-starwars-xxxx",namespace="gpu-test3",pod="pod1"} 23792 DCGM_FI_PROF_PCIE_RX_BYTES{gpu="0",UUID="GPU-a4f34abc-7715-3560-dcea-7238b9611a45",pci_bus_id="00000009:01:00.0",device="nvidia0",modelName="NVIDIA GH200 96GB HBM3",Hostname="sc-starwars-xxxx",container="ctr",dra_claim_name="single-gpu",dra_claim_namespace="gpu-test3",dra_device_name="gpu-0",dra_driver_name="gpu.nvidia.com",dra_pool_name="sc-starwars-xxxx",namespace="gpu-test3",pod="pod2"} 23792 with Pod DRA info: container="ctr", dra_claim_name="single-gpu", dra_claim_namespace="gpu-test3", dra_device_name="gpu-0",dra_driver_name="gpu.nvidia.com", dra_pool_name="sc-starwars-xxxx", namespace="gpu-test3", pod="pod1" ```
* **`KubeletPodResourcesGet`**: Enable the `Get` gRPC endpoint on kubelet's for Pod resources. This API augments the [resource allocation reporting](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md#monitoring-device-plugin-resources).
* **`KubeletSeparateDiskGC`**: The split image filesystem feature enables kubelet to perform garbage collection of images (read-only layers) and/or containers (writeable layers) deployed on separate filesystems.
* **`KubeletServiceAccountTokenForCredentialProviders`**: Enable kubelet to send the service account token bound to the pod for which the image is being pulled to the credential provider plugin.
* **`KubeletTracing`**: Add support for distributed tracing in the kubelet. When enabled, kubelet CRI interface and authenticated http servers are instrumented to generate OpenTelemetry trace spans. See [Traces for Kubernetes System Components](/concepts/cluster-administration/system-traces.md) for more details.
* **`ListFromCacheSnapshot`**: Enables the API server to generate snapshots for the watch cache store and using them to serve LIST requests.
* **`LocalStorageCapacityIsolationFSQuotaMonitoring`**: When `LocalStorageCapacityIsolation` is enabled for [local ephemeral storage](/concepts/configuration/manage-resources-containers.md), the backing filesystem for [emptyDir volumes](/concepts/storage/volumes.md#emptydir) supports project quotas, and `UserNamespacesSupport` is enabled, project quotas are used to monitor `emptyDir` volume storage consumption rather than using filesystem walk, ensuring better performance and accuracy.
* **`LogarithmicScaleDown`**: Enable semi-random selection of pods to evict on controller scaledown based on logarithmic bucketing of pod timestamps.
* **`LoggingAlphaOptions`**: Allow fine-tuning of experimental, alpha-quality logging options.
* **`LoggingBetaOptions`**: Allow fine-tuning of experimental, beta-quality logging options.
* **`ManifestBasedAdmissionControlConfig`**: Enable loading admission webhooks and CEL-based admission policies from static manifest files on disk via the `staticManifestsDir` field in `AdmissionConfiguration`. These policies are active from API server startup, survive etcd unavailability, and can protect API-based admission resources from modification.
* **`MatchLabelKeysInPodAffinity`**: Enable the `matchLabelKeys` and `mismatchLabelKeys` fields for [pod (anti)affinity](/concepts/scheduling-eviction/assign-pod-node.md).
* **`MatchLabelKeysInPodTopologySpread`**: Enable the `matchLabelKeys` field for [Pod topology spread constraints](/concepts/scheduling-eviction/topology-spread-constraints.md).
* **`MatchLabelKeysInPodTopologySpreadSelectorMerge`**: Enable merging of selectors built from `matchLabelKeys` into `labelSelector` of [Pod topology spread constraints](/concepts/scheduling-eviction/topology-spread-constraints.md). This feature gate can be enabled when `matchLabelKeys` feature is enabled with the `MatchLabelKeysInPodTopologySpread` feature flag.
* **`MaxUnavailableStatefulSet`**: Enables setting the `maxUnavailable` field for the [rolling update strategy](/concepts/workloads/controllers/statefulset.md#rolling-updates) of a StatefulSet. The field specifies the maximum number of Pods that can be unavailable during the update.
* **`MemoryManager`**: Allows setting memory affinity for a container based on NUMA topology.
* **`MemoryQoS`**: Enable memory protection and usage throttling for Pods and containers using the cgroup v2 memory controller. When `memoryThrottlingFactor` is set, the `kubelet` sets `memory.high` for throttling on Burstable and BestEffort containers. The `kubelet` optionally sets `memory.min` and `memory.low` for tiered memory protection when `memoryReservationPolicy` is set to `TieredReservation`. Feature requires both - feature gate enablement and `kubelet` configuration setting.
* **`MultiCIDRServiceAllocator`**: Track IP address allocations for Service cluster IPs using IPAddress objects.
* **`MutableCSINodeAllocatableCount`**: Make the `.spec.drivers[*].allocatable.count` field of a CSINode mutable. Also, enable a CSIDriver field, `nodeAllocatableUpdatePeriodSeconds`. This allows periodic updates to a node's reported allocatable volume capacity, preventing stateful pods from becoming stuck due to outdated information that the kube-scheduler would otherwise rely upon.
* **`MutablePVNodeAffinity`**: Allow update to the `.spec.nodeAffinity` field of a PersistentVolume. See [Updates to node affinity](/concepts/storage/persistent-volumes.md#updates-to-node-affinity) for more details.
* **`MutablePodResourcesForSuspendedJobs`**: Enable the ability to patch pod templates for suspended Jobs, in order to change requests or limits for infrastructure resources.
* **`MutableSchedulingDirectivesForSuspendedJobs`**: Enable the ability to patch pod templates for suspended Jobs, in order to change the pod scheduling directives.
* **`MutatingAdmissionPolicy`**: Enable [MutatingAdmissionPolicy](/reference/access-authn-authz/mutating-admission-policy.md) support, which allows [CEL](/reference/using-api/cel.md) mutations to be applied during admission control. For Kubernetes v1.30 and v1.31, this feature gate existed but had no effect.
* **`NFTablesProxyMode`**: Allow running kube-proxy in [nftables mode](/reference/networking/virtual-ips.md#proxy-mode-nftables).
* **`NativeHistograms`**: Enables Kubernetes components to expose metrics in Prometheus Native Histogram format for improved efficiency and finer bucket resolution. See [Native Histograms](/reference/instrumentation/native-histograms.md) for more information.
* **`NodeDeclaredFeatures`**: Enables Nodes to report supported features via their `.status`. This enables the scheduler and admission controller to prevent operations on nodes lacking features required by the Pod. See [Node Declared Features](/concepts/scheduling-eviction/node-declared-features.md).
* **`NodeInclusionPolicyInPodTopologySpread`**: Enable using `nodeAffinityPolicy` and `nodeTaintsPolicy` in [Pod topology spread constraints](/concepts/scheduling-eviction/topology-spread-constraints.md) when calculating pod topology spread skew.
* **`NodeLifecycleConditions`**: Enables well-known Node conditions that report drain, maintenance, and [Graceful Node Shutdown](/concepts/cluster-administration/node-shutdown.md) lifecycle state. See [Node lifecycle conditions](/reference/node/node-lifecycle-conditions.md).
* **`NodeLogQuery`**: Enables querying logs of node services using the `/logs` endpoint.
* **`NodeSwap`**: Enable the kubelet to allocate swap memory for Kubernetes workloads on a node. Must be used with `KubeletConfiguration.failSwapOn` set to false. For more details, please see [swap memory](/concepts/architecture/nodes.md#swap-memory)
* **`NominatedNodeNameForExpectation`**: When enabled, kube-scheduler uses `.status.nominatedNodeName` to express where a Pod is going to be bound. The `.status.nominatedNodeName` field is set when kube-scheduler triggers preemption of pods, or anticipates that WaitOnPermit or PreBinding phase will take relatively long. Other components may read and use `.status.nominatedNodeName`, but should not set it. When disabled, kube-scheduler will only set `.status.nominatedNodeName` before triggering preemption.
* **`OpenAPIEnums`**: Enables populating "enum" fields of OpenAPI schemas in the spec returned from the API server.
* **`OpportunisticBatching`**: Enable reusing of scheduling results from the previous scheduling cycle for equivalent pods.
* **`OrderedNamespaceDeletion`**: While deleting namespace, the pods resources is going to be deleted before the rest of resources.
* **`PersistentVolumeClaimUnusedSinceTime`**: When enabled, the PVC protection controller adds an `Unused` condition to PersistentVolumeClaims that tracks whether the PVC is currently referenced by any non-terminal Pod. The condition's `lastTransitionTime` records when the PVC last transitioned between being in use and being unused.
* **`PodAndContainerStatsFromCRI`**: Configure the kubelet to gather container and pod stats from the CRI container runtime rather than gathering them from cAdvisor. As of 1.26, this also includes gathering metrics from CRI and emitting them over `/metrics/cadvisor` (rather than having cAdvisor emit them directly).
* **`PodCertificateRequest`**: Enable PodCertificateRequest objects and podCertificate projected volume sources.
* **`PodDeletionCost`**: Enable the [Pod Deletion Cost](/concepts/workloads/controllers/replicaset.md#pod-deletion-cost) feature which allows users to influence ReplicaSet downscaling order.
* **`PodGroupPreemptionPolicy`**: Enables the support for `PreemptionPolicy` field in [PodGroup API](/concepts/workloads/podgroup-api.md) and [Workload API](/concepts/workloads/workload-api.md). When enabled, if a PodGroup has `PreemptionPolicy: Never` it will not perform workload aware preemption.
* **`PodIndexLabel`**: Enables the Job controller and StatefulSet controller to add the pod index as a label when creating new pods. See [Job completion mode docs](/concepts/workloads/controllers/job.md#completion-mode) and [StatefulSet pod index label docs](/concepts/workloads/controllers/statefulset.md#pod-index-label) for more details.
* **`PodLevelResourceManagers`**: Enable _Pod-level resource managers_: the ability for the Topology, CPU, and Memory managers to use information from `.spec.resources` to perform NUMA alignment for an entire pod and manage resources flexibly for the containers within that pod.
* **`PodLevelResources`**: Enable _Pod level resources_: the ability to specify resource requests and limits at the Pod level, rather than only for specific containers.
* **`PodLifecycleSleepAction`**: Enables the `sleep` action in Container lifecycle hooks (`preStop` and `postStart`).
* **`PodLifecycleSleepActionAllowZero`**: Enables setting zero value for the `sleep` action in [container lifecycle hooks](/concepts/containers/container-lifecycle-hooks.md).
* **`PodLogsQuerySplitStreams`**: Enable fetching specific log streams (either stdout or stderr) from a container's log streams, using the Pod API.
* **`PodObservedGenerationTracking`**: Enables the kubelet to set `observedGeneration` in the Pod `.status`, and enables other components to set `observedGeneration` in pod conditions. This feature allows reflecting the `.metadata.generation` of the Pod at the time that the overall status, or some specific condition, was being recorded. Storing it helps avoid risks associated with _lost updates_.
* **`PodReadyToStartContainersCondition`**: Enable the kubelet to mark the [PodReadyToStartContainers](/concepts/workloads/pods/pod-lifecycle.md#pod-has-network) condition on pods. This feature gate was previously known as `PodHasNetworkCondition`, and the associated condition was named `PodHasNetwork`.
* **`PodSchedulingReadiness`**: Enable setting `schedulingGates` field to control a Pod's [scheduling readiness](/concepts/scheduling-eviction/pod-scheduling-readiness.md).
* **`PodTopologyLabelsAdmission`**: Enables the `PodTopologyLabels` admission plugin. See [Pod Topology Labels](/reference/access-authn-authz/admission-controllers.md#podtopologylabels) for details.
* **`PodsAPI`**: Enables the kubelet Pods API gRPC service. See [Kubelet Pods API](/reference/node/kubelet-pods-api.md) for more details.
* **`PortForwardWebsockets`**: Allow WebSocket streaming of the portforward sub-protocol (`port-forward`) from clients requesting version v2 (`v2.portforward.k8s.io`) of the sub-protocol.
* **`PreferSameTrafficDistribution`**: Allows usage of the values `PreferSameZone` and `PreferSameNode` in the Service [`trafficDistribution`](/reference/networking/virtual-ips.md#traffic-distribution) field.
* **`PreventStaticPodAPIReferences`**: Denies Pod admission if static Pods reference other API objects.
* **`ProcMountType`**: Enables control over the type proc mounts for containers by setting the `procMount` field of a Pod's `securityContext`.
* **`QOSReserved`**: Allows resource reservations at the QoS level preventing pods at lower QoS levels from bursting into resources requested at higher QoS levels (memory only for now).
* **`RecoverVolumeExpansionFailure`**: Enables users to edit their PVCs to smaller sizes so as they can recover from previously issued volume expansion failures. See [Recovering from Failure when Expanding Volumes](/concepts/storage/persistent-volumes.md#recovering-from-failure-when-expanding-volumes) for more details.
* **`RecursiveReadOnlyMounts`**: Enables support for recursive read-only mounts. For more details, see [read-only mounts](/concepts/storage/volumes.md#read-only-mounts).
* **`ReduceDefaultCrashLoopBackOffDecay`**: Enabled reduction of both the initial delay and the maximum delay accrued between container restarts for a node for containers in `CrashLoopBackOff` across the cluster to `1s` initial delay and `60s` maximum delay.
* **`RelaxedEnvironmentVariableValidation`**: Allow almost all printable ASCII characters in environment variables.
* **`RelaxedServiceNameValidation`**: Enables relaxed validation for Service object names, allowing the use of [RFC 1123 label names](/concepts/overview/working-with-objects/names.md#dns-label-names) instead of [RFC 1035 label names](/concepts/overview/working-with-objects/names.md#rfc-1035-label-names). This feature allows Service object names to start with a digit.
* **`ReloadKubeletServerCertificateFile`**: Enable the kubelet TLS server to update its certificate if the specified certificate file are changed. This feature is useful when specifying `tlsCertFile` and `tlsPrivateKeyFile` in kubelet configuration. The feature gate has no effect for other cases such as using TLS bootstrap.
* **`RemoteRequestHeaderUID`**: Enable the API server to accept UIDs (user IDs) via request header authentication. This will also make the `kube-apiserver`'s API aggregator add UIDs via standard headers when forwarding requests to the servers serving the aggregated API.
* **`ResilientWatchCacheInitialization`**: Enables resilient watchcache initialization to avoid controlplane overload.
* **`ResourceHealthStatus`**: Enable the `allocatedResourcesStatus` field within the `.status` for a Pod. The field reports additional details for each container in the Pod, with the health information for each device assigned to the Pod. Starting in v1.36 (beta), the health report includes an optional `message` field that provides additional human-readable context about the health status, such as error details or failure reasons. This feature applies to devices managed by both [Device Plugins](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md#device-plugin-and-unhealthy-devices) and [Dynamic Resource Allocation](/concepts/resource-management/dynamic-resource-allocation/dra-observability.md#device-health-monitoring). See [Device plugin and unhealthy devices](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md#device-plugin-and-unhealthy-devices) for more details.
* **`RestartAllContainersOnContainerExits`**: Enables the ability to specify `RestartAllContainers` as an action in container `restartPolicyRules`. When a container's exit matches a rule with this action, the entire Pod is terminated and restarted in-place. `RestartAllContainersOnContainerExits` depends on both the `ContainerRestartRules` and `NodeDeclaredFeatures` feature gates. If the dependent feature gates are not enabled, kubelet startup can fail. See [Restart All Containers](/concepts/workloads/pods/pod-lifecycle.md#restart-all-containers) for more details.
* **`RetryGenerateName`**: Enables retrying of object creation when the [API server](/reference/glossary/kube-apiserver.md) is expected to generate a [name](/concepts/overview/working-with-objects/names.md#names). When this feature is enabled, requests using `generateName` are retried automatically in case the control plane detects a name conflict with an existing object, up to a limit of 8 total attempts.
* **`RotateKubeletServerCertificate`**: Enable the rotation of the server TLS certificate on the kubelet. See [kubelet configuration](/reference/access-authn-authz/kubelet-tls-bootstrapping.md#kubelet-configuration) for more details.
* **`RuntimeClassInImageCriApi`**: Enables images to be pulled based on the [runtime class](/concepts/containers/runtime-class.md) of the pods that reference them.
* **`SELinuxChangePolicy`**: Enables `spec.securityContext.seLinuxChangePolicy` field. This field can be used to opt-out from applying the SELinux label to the pod volumes using mount options. This is required when a single volume that supports mounting with SELinux mount option is shared between Pods that have different SELinux labels, such as a privileged and unprivileged Pods. Enabling the `SELinuxChangePolicy` feature gate requires the feature gate `SELinuxMountReadWriteOncePod` to be enabled.
* **`SELinuxMount`**: Speeds up container startup by allowing kubelet to mount volumes for a Pod directly with the correct SELinux label instead of changing each file on the volumes recursively. It widens the performance improvements behind the `SELinuxMountReadWriteOncePod` feature gate by extending the implementation to all volumes.
* **`SELinuxMountReadWriteOncePod`**: Speeds up container startup by allowing kubelet to mount volumes for a Pod directly with the correct SELinux label instead of changing each file on the volumes recursively. The initial implementation focused on ReadWriteOncePod volumes.
* **`SchedulerAsyncAPICalls`**: Change the kube-scheduler to make the entire scheduling cycle free of blocking requests to the Kubernetes API server. Instead, interact with the Kubernetes API using asynchronous code.
* **`SchedulerAsyncPreemption`**: Enable running some expensive operations within the scheduler, associated with [preemption](/concepts/scheduling-eviction/pod-priority-preemption.md), asynchronously. Asynchronous processing of preemption improves overall Pod scheduling latency.
* **`SchedulerPopFromBackoffQ`**: Improves scheduling queue behavior by popping pods from the backoffQ when the activeQ is empty. This allows to process potentially schedulable pods ASAP, eliminating a penalty effect of the backoff queue.
* **`SchedulerQueueingHints`**: Enables scheduler [queueing hints](/concepts/scheduling-eviction/scheduling-framework.md#queueinghint), which benefits to reduce the useless requeuing. The scheduler retries scheduling pods if something changes in the cluster that could make the pod scheduled. Queueing hints are internal signals that allow the scheduler to filter the changes in the cluster that are relevant to the unscheduled pod, based on previous scheduling attempts.
* **`SeparateCacheWatchRPC`**: Allows the API server watch cache to create a watch on a dedicated RPC. This prevents watch cache from being starved by other watches.
* **`SeparateTaintEvictionController`**: Enables running the _taint based eviction_ controller, that performs [Taint-based Evictions](/concepts/scheduling-eviction/taint-and-toleration.md#taint-based-evictions), as a standalone controller (separate from the _node lifecycle_ controller).
* **`ServiceAccountNodeAudienceRestriction`**: This gate is used to restrict the audience for which the kubelet can request a service account token for.
* **`ServiceAccountTokenJTI`**: Controls whether JTIs (UUIDs) are embedded into generated service account tokens, and whether these JTIs are recorded into the Kubernetes audit log for future requests made by these tokens.
* **`ServiceAccountTokenNodeBinding`**: Controls whether the API server allows binding service account tokens to Node objects.
* **`ServiceAccountTokenNodeBindingValidation`**: Controls whether the apiserver will validate a Node reference in service account tokens.
* **`ServiceAccountTokenPodNodeInfo`**: Controls whether the apiserver embeds the node name and uid for the associated node when issuing service account tokens bound to Pod objects.
* **`ServiceTrafficDistribution`**: Allows usage of the optional `spec.trafficDistribution` field in Services. The field offers a way to express preferences for how traffic is distributed to Service endpoints.
* **`ShardedListAndWatch`**: Enable support for the `shardSelector` parameter on **list** and **watch** requests, allowing clients to receive a filtered subset of objects based on hash ranges of metadata fields (such as UID). See [Sharded list and watch](/reference/using-api/api-concepts.md#sharded-list-and-watch) for more details.
* **`SidecarContainers`**: Allow setting the `restartPolicy` of an init container to `Always` so that the container becomes a sidecar container (restartable init containers). See [Sidecar containers and restartPolicy](/concepts/workloads/pods/sidecar-containers.md) for more details.
* **`SizeBasedListCostEstimate`**: Enables APF to use size of objects for estimating request cost.
* **`StaleControllerConsistencyDaemonSet`**: Enables behavior within the DaemonSet controller to ensure that prior writes to the API server are observed before proceeding with additional reconciliation for the same DaemonSet. This is to prevent stale cache from causing incorrect or spurious updates to the DaemonSet.
* **`StaleControllerConsistencyJob`**: Enables behavior within the Job controller to ensure that prior writes to the API server are observed before proceeding with additional reconciliation for the same Job. This is to prevent stale cache from causing incorrect or spurious updates to the Job.
* **`StaleControllerConsistencyReplicaSet`**: Enables behavior within the ReplicaSet controller to ensure that prior writes to the API server are observed before proceeding with additional reconciliation for the same ReplicaSet. This is to prevent stale cache from causing incorrect or spurious updates to the ReplicaSet.
* **`StaleControllerConsistencyStatefulSet`**: Enables behavior within the StatefulSet controller to ensure that prior writes to the API server are observed before proceeding with additional reconciliation for the same StatefulSet. This is to prevent stale cache from causing incorrect or spurious updates to the StatefulSet.
* **`StatefulSetAutoDeletePVC`**: Allows the use of the optional `.spec.persistentVolumeClaimRetentionPolicy` field, providing control over the deletion of PVCs in a StatefulSet's lifecycle. See [PersistentVolumeClaim retention](/concepts/workloads/controllers/statefulset.md#persistentvolumeclaim-retention) for more details.
* **`StatefulSetRecreateStrategy`**: Enables the `Recreate` update strategy for StatefulSets, which deletes all of a StatefulSet's Pods before creating new Pods that reflect modifications made to the StatefulSet's `.spec.template`. See [Recreate](/concepts/workloads/controllers/statefulset.md#recreate) for details.
* **`StatefulSetStartOrdinal`**: Allow configuration of the start ordinal in a StatefulSet. See [Start ordinal](/concepts/workloads/controllers/statefulset.md#start-ordinal) for more details.
* **`StorageCapacityScoring`**: The feature gate `VolumeCapacityPriority` was used in v1.32 to support storage that are statically provisioned. Starting from v1.33, the new feature gate `StorageCapacityScoring` replaces the old `VolumeCapacityPriority` gate with added support to dynamically provisioned storage. When `StorageCapacityScoring` is enabled, the VolumeBinding plugin in the kube-scheduler is extended to score Nodes based on the storage capacity on each of them. This feature is applicable to CSI volumes that supported [Storage Capacity](/concepts/storage/storage-capacity.md), including local storage backed by a CSI driver.
* **`StorageNamespaceIndex`**: Enables a namespace indexer for namespace scoped resources in API server cache to accelerate list operations.
* **`StorageVersionAPI`**: Enable the [storage version API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#storageversion-v1alpha1-internal-apiserver-k8s-io).
* **`StorageVersionHash`**: Allow API servers to expose the storage version hash in the discovery.
* **`StorageVersionMigrator`**: Enables the migration of the [storage version](/concepts/overview/working-with-objects/storage-version.md) of a resource.
* **`StreamingCollectionEncodingToJSON`**: Allow the API server JSON encoder to encode collections item by item, instead of all at once.
* **`StreamingCollectionEncodingToProtobuf`**: Allow the API server Protobuf encoder to encode collections item by item, instead of all at once.
* **`StrictCostEnforcementForVAP`**: Apply strict CEL cost validation for ValidatingAdmissionPolicies.
* **`StrictCostEnforcementForWebhooks`**: Apply strict CEL cost validation for `matchConditions` within admission webhooks.
* **`StrictIPCIDRValidation`**: Use stricter validation for fields containing IP addresses and CIDR values. In particular, with this feature gate enabled, octets within IPv4 addresses are not allowed to have any leading `0`s, and IPv4-mapped IPv6 values (e.g. `::ffff:192.168.0.1`) are forbidden. These sorts of values can potentially cause security problems when different components interpret the same string as referring to different IP addresses (as in CVE-2021-29923). This tightening applies only to fields in build-in API kinds, and not to custom resource kinds, values in Kubernetes configuration files, or command-line arguments.
* **`StructuredAuthenticationConfiguration`**: Enable [structured authentication configuration](/reference/access-authn-authz/authentication.md#configuring-the-api-server) for the API server.
* **`StructuredAuthenticationConfigurationEgressSelector`**: Enables Egress Selector in Structured Authentication Configuration.
* **`StructuredAuthenticationConfigurationJWKSMetrics`**: Enables additional metrics for JSON Web Key Set (JWKS) operations in JWT authenticators configured via `--authentication-config`. When enabled, the API server records metrics about the last time JWKS was fetched and the hash value of the JWKS response. See the [metrics reference](/reference/instrumentation/metrics.md) for details.
* **`StructuredAuthorizationConfiguration`**: Enable structured authorization configuration, so that cluster administrators can specify more than one [authorization webhook](/reference/access-authn-authz/webhook.md) in the API server handler chain.
* **`SupplementalGroupsPolicy`**: Enables support for fine-grained SupplementalGroups control. For more details, see [Configure fine-grained SupplementalGroups control for a Pod](/tasks/configure-pod-container/security-context.md#supplementalgroupspolicy).
* **`SystemdWatchdog`**: Allow using systemd watchdog to monitor the health status of kubelet. See [Kubelet Systemd Watchdog](/reference/node/systemd-watchdog.md) for more details.
* **`TaintTolerationComparisonOperators`**: Enables numeric comparison operators (`Lt` and `Gt`) for [tolerations](/concepts/scheduling-eviction/taint-and-toleration.md).
* **`TokenRequestServiceAccountUIDValidation`**: This is used to ensure that the UID provided in the TokenRequest matches the UID of the ServiceAccount for which the token is being requested. It helps prevent misuse of the TokenRequest API by ensuring that tokens are only issued for the correct ServiceAccount.
* **`TopologyAwareWorkloadScheduling`**: Enable topology-aware scheduling for Workloads.
* **`TopologyManagerPolicyAlphaOptions`**: Allow fine-tuning of topology manager policies, experimental, Alpha-quality options. This feature gate guards *a group* of topology manager options whose quality level is alpha. This feature gate will never graduate to beta or stable.
* **`TopologyManagerPolicyBetaOptions`**: Allow fine-tuning of topology manager policies, experimental, Beta-quality options. This feature gate guards *a group* of topology manager options whose quality level is beta. This feature gate will never graduate to stable.
* **`TopologyManagerPolicyOptions`**: Enable [fine-tuning](/tasks/administer-cluster/topology-manager.md#topology-manager-policy-options) of topology manager policies.
* **`TranslateStreamCloseWebsocketRequests`**: Allow WebSocket streaming of the remote command sub-protocol (`exec`, `cp`, `attach`) from clients requesting version 5 (v5) of the sub-protocol.
* **`UnauthenticatedHTTP2DOSMitigation`**: Enables HTTP/2 Denial of Service (DoS) mitigations for unauthenticated clients. Kubernetes v1.28.0 through v1.28.2 do not include this feature gate.
* **`UnknownVersionInteroperabilityProxy`**: Proxy resource requests to the correct peer kube-apiserver when multiple kube-apiservers exist at varied versions. See [Mixed version proxy](/concepts/architecture/mixed-version-proxy.md) for more information.
* **`UnlockWhileProcessingFIFO`**: Enable use of a FIFO queue within client-go that unlocks while processing events. If not enabled, the queue instead holds the lock for the entire duration of processing events, which could lead to performance issues in high-throughput scenarios. This feature gate can be toggled in the kube-controller-manager and any client-go based controller. You can only enable this feature gate if the [AtomicFIFO](/reference/command-line-tools-reference/feature-gates.md#AtomicFIFO) feature gate is also enabled.
* **`UserNamespacesHostNetworkSupport`**: When enabled, pods are allowed to use both `hostNetwork` and [User Namespaces](/concepts/workloads/pods/user-namespaces.md) simultaneously.
* **`UserNamespacesSupport`**: Enable user namespace support for Pods.
* **`VolumeAttributesClass`**: Enable support for VolumeAttributesClasses. See [Volume Attributes Classes](/concepts/storage/volume-attributes-classes.md) for more information.
* **`VolumeBindMountOptions`**: Enables setting bind mount options (`noexec`, `nodev`, `nosuid`) per container volume mount using the `bindMountOptions` field in `volumeMounts`. When enabled, the kubelet passes these options to the container runtime, which applies them as Linux bind mount flags. The container runtime must support the `mount_options` field in the CRI `Mount` message. This field has no effect on Windows nodes.
* **`VolumeLimitScaling`**: Enables volume limit scaling for CSI drivers. This allows scheduler to co-ordinate better with cluster-autoscaler for storage limits. See [Storage Limits](/concepts/storage/storage-limits.md) for more information.
* **`WatchCacheInitializationPostStartHook`**: Enables post-start-hook for watchcache initialization to be part of readyz (with timeout).
* **`WatchFromStorageWithoutResourceVersion`**: Enables watches without `resourceVersion` to be served from storage.
* **`WatchList`**: Enable support for [streaming initial state of objects in watch requests](/reference/using-api/api-concepts.md#streaming-lists).
* **`WatchListClient`**: Allows an API client to request a stream of data rather than fetching a full list. This functionality is available in `client-go` and requires the [WatchList](/reference/command-line-tools-reference/feature-gates.md) feature to be enabled on the server. If the `WatchList` is not supported on the server, the client will seamlessly fall back to a standard list request.
* **`WatchListCompression`**: Enables compression for [_watch list_](/reference/using-api/api-concepts.md#streaming-lists) responses from the API server. This feature gate has no effect if `WatchList` is disabled.
* **`WinDSR`**: Allows kube-proxy to create DSR loadbalancers for Windows.
* **`WinOverlay`**: Allows kube-proxy to run in overlay mode for Windows.
* **`WindowsCPUAndMemoryAffinity`**: Add CPU and Memory Affinity support to Windows nodes with [CPUManager](/tasks/administer-cluster/cpu-management-policies.md#windows-support), [MemoryManager](/tasks/administer-cluster/memory-manager.md#windows-support) and topology manager.
* **`WindowsGracefulNodeShutdown`**: Enables support for windows node graceful shutdown in kubelet. During a system shutdown, kubelet will attempt to detect the shutdown event and gracefully terminate pods running on the node. See [Graceful Node Shutdown](/concepts/cluster-administration/node-shutdown.md#graceful-node-shutdown) for more details.
* **`WindowsHostNetwork`**: Enables support for joining Windows containers to a hosts' network namespace.
* **`WorkloadWithJob`**: Enables the Job controller to compile a Job's `.spec.scheduling` configuration into [Workload](/concepts/workloads/workload-api.md) and [PodGroup](/concepts/workloads/podgroup-api.md) objects before it creates any Pods. When `.spec.scheduling` is omitted, the Job defaults to the `Basic` scheduling policy. See [Integrate with Workload APIs](/concepts/workloads/controllers/job.md#integrate-with-workload-apis) for details.

## What's next

* The [deprecation policy](/reference/deprecation-policy.md) for Kubernetes explains
  the project's approach to removing features and components.
* Since Kubernetes 1.24, new beta APIs are not enabled by default.  When enabling a beta
  feature, you will also need to enable any associated API resources.
  For example, to enable a particular resource like
  `storage.k8s.io/v1beta1/csistoragecapacities`, set `--runtime-config=storage.k8s.io/v1beta1/csistoragecapacities`.
  See [API Versioning](/reference/using-api.md#api-versioning) for more details on the command line flags.
* See [Configure Feature Gates](/tasks/administer-cluster/configure-feature-gates.md)
  for step-by-step guidance on enabling feature gates.
