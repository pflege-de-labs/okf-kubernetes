# Workload Management

* [Deployments](deployment.md) - A Deployment manages a set of Pods to run an application workload, usually one that doesn't maintain state.
* [ReplicaSet](replicaset.md) - A ReplicaSet's purpose is to maintain a stable set of replica Pods running at any given time. Usually, you define a Deployment and let that Deployment manage ReplicaSets automatically.
* [StatefulSets](statefulset.md) - A StatefulSet runs a group of Pods, and maintains a sticky identity for each of those Pods. This is useful for managing applications that need persistent storage or a stable, unique network identity.
* [DaemonSet](daemonset.md) - A DaemonSet defines Pods that provide node-local facilities. These might be fundamental to the operation of your cluster, such as a networking helper tool, or be part of an add-on.
* [Jobs](job.md) - Jobs represent one-off tasks that run to completion and then stop.
* [Automatic Cleanup for Finished Jobs](ttlafterfinished.md) - A time-to-live mechanism to clean up old Jobs that have finished execution.
* [CronJob](cron-jobs.md) - A CronJob starts one-time Jobs on a repeating schedule.
* [ReplicationController](replicationcontroller.md) - Legacy API for managing workloads that can scale horizontally. Superseded by the Deployment and ReplicaSet APIs.
