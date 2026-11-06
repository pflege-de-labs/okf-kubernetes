# Apps

* [ControllerRevision](controller-revision-v1.md) - ControllerRevision implements an immutable snapshot of state data. Clients are responsible for serializing and deserializing the objects that contain their internal state. Once a ControllerRevision has been successfully created, it can not be updated. The API Server will fail validation of all requests that attempt to mutate the Data field. ControllerRevisions may, however, be deleted. Note that, due to its use by both the DaemonSet and StatefulSet controllers for update and rollback, this object is beta. However, it may be subject to name and representation changes in future releases, and clients should not depend on its stability. It is primarily for internal use by controllers.
* [DaemonSet](daemon-set-v1.md) - DaemonSet represents the configuration of a daemon set.
* [Deployment](deployment-v1.md) - Deployment enables declarative updates for Pods and ReplicaSets.
* [ReplicaSet](replica-set-v1.md) - ReplicaSet ensures that a specified number of pod replicas are running at any given time.
* [StatefulSet](stateful-set-v1.md) - StatefulSet represents a set of pods with consistent identities. Identities are defined as: - Network: A single stable DNS and hostname. - Storage: As many VolumeClaims as requested. The StatefulSet guarantees that a given network identity will always map to the same storage identity.
