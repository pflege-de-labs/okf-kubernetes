# Objects In Kubernetes

* [Kubernetes Object Management](object-management.md) - The `kubectl` command-line tool supports several different ways to create and manage Kubernetes objects.
* [Object Names and IDs](names.md) - Each object in your cluster has a _Name_ that is unique for that type of resource.
* [Labels and Selectors](labels.md) - _Labels_ are key/value pairs that are attached to objects such as Pods.
* [Namespaces](namespaces.md) - In Kubernetes, _namespaces_ provide a mechanism for isolating groups of resources within a single cluster.
* [Annotations](annotations.md) - You can use Kubernetes annotations to attach arbitrary non-identifying metadata to objects.
* [Field Selectors](field-selectors.md) - _Field selectors_ let you select Kubernetes objects based on the value of one or more resource fields.
* [Finalizers](finalizers.md) - Finalizers are namespaced keys that tell Kubernetes to wait until specific conditions are met before it fully deletes resources that are marked for deletion.
* [Owners and Dependents](owners-dependents.md) - In Kubernetes, some objects are *owners* of other objects.
* [Recommended Labels](common-labels.md) - You can visualize and manage Kubernetes objects with more tools than kubectl and the dashboard.
* [Storage Versions](storage-version.md) - The Kubernetes API server stores objects, relying on an etcd-compatible backing store (often, the backing storage is etcd itself).
