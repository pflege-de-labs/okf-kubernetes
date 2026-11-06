# Storage

* [Volumes](volumes.md) - Kubernetes _volumes_ provide a way for containers in a Pod to access and share data via the filesystem.
* [Persistent Volumes](persistent-volumes.md) - This document describes _persistent volumes_ in Kubernetes.
* [Projected Volumes](projected-volumes.md) - This document describes _projected volumes_ in Kubernetes.
* [Ephemeral Volumes](ephemeral-volumes.md) - This document describes _ephemeral volumes_ in Kubernetes.
* [Storage Classes](storage-classes.md) - This document describes the concept of a StorageClass in Kubernetes.
* [Volume Attributes Classes](volume-attributes-classes.md) - This page assumes that you are familiar with StorageClasses, volumes and PersistentVolumes in Kubernetes.
* [Dynamic Volume Provisioning](dynamic-provisioning.md) - Dynamic volume provisioning allows storage volumes to be created on-demand.
* [Volume Snapshots](volume-snapshots.md) - In Kubernetes, a _VolumeSnapshot_ represents a snapshot of a volume on a storage system.
* [Volume Snapshot Classes](volume-snapshot-classes.md) - This document describes the concept of VolumeSnapshotClass in Kubernetes.
* [CSI Volume Cloning](volume-pvc-datasource.md) - This document describes the concept of cloning existing CSI Volumes in Kubernetes.
* [Volume Populators and Data Sources](volume-populators-and-data-sources.md) - This document describes _volume populators_ and _data sources_ in Kubernetes.
* [Storage Capacity](storage-capacity.md) - Storage capacity is limited and may vary depending on the node on which a pod runs: network-attached storage might not be accessible by all nodes, or storage is local to a node to begin with.
* [Node-specific Volume Limits](storage-limits.md) - This page describes the maximum number of volumes that can be attached to a Node for various cloud providers.
* [Local ephemeral storage](ephemeral-storage.md) - Nodes have local ephemeral storage, backed by locally-attached writeable devices or, sometimes, by RAM.
* [Volume Health Monitoring](volume-health-monitoring.md) - CSI volume health monitoring lets a CSI driver report health problems with a volume or with its storage backend directly to Kubernetes.
* [Windows Storage](windows-storage.md) - This page provides an storage overview specific to the Windows operating system.
