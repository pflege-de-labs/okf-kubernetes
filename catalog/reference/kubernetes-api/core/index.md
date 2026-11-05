# Core

* [ComponentStatus](component-status-v1.md) - ComponentStatus (and ComponentStatusList) holds the cluster validation info. Deprecated: This API is deprecated in v1.19+
* [ConfigMap](config-map-v1.md) - ConfigMap holds configuration data for pods to consume.
* [Endpoints](endpoints-v1.md) - Endpoints is a collection of endpoints that implement the actual service. Example: Name: &#34;mysvc&#34;, Subsets: [ { Addresses: [{&#34;ip&#34;: &#34;10.10.1.1&#34;}, {&#34;ip&#34;: &#34;10.10.2.2&#34;}], Ports: [{&#34;name&#34;: &#34;a&#34;, &#34;port&#34;: 8675}, {&#34;name&#34;: &#34;b&#34;, &#34;port&#34;: 309}] }, { Addresses: [{&#34;ip&#34;: &#34;10.10.3.3&#34;}], Ports: [{&#34;name&#34;: &#34;a&#34;, &#34;port&#34;: 93}, {&#34;name&#34;: &#34;b&#34;, &#34;port&#34;: 76}] }, ] Endpoints is a legacy API and does not contain information about all Service features. Use discoveryv1.EndpointSlice for complete information about Service endpoints. Deprecated: This API is deprecated in v1.33+. Use discoveryv1.EndpointSlice.
* [Event](event-v1.md) - Event is a report of an event somewhere in the cluster. Events have a limited retention time and triggers and messages may evolve with time. Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason. Events should be treated as informative, best-effort, supplemental data.
* [LimitRange](limit-range-v1.md) - LimitRange sets resource usage limits for each kind of resource in a Namespace.
* [Namespace](namespace-v1.md) - Namespace provides a scope for Names. Use of multiple namespaces is optional.
* [Node](node-v1.md) - Node is a worker node in Kubernetes. Each node will have a unique identifier in the cache (i.e. in etcd).
* [PersistentVolume](persistent-volume-v1.md) - PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
* [PersistentVolumeClaim](persistent-volume-claim-v1.md) - PersistentVolumeClaim is a user&#39;s request for and claim to a persistent volume
* [Pod](pod-v1.md) - Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.
* [PodTemplate](pod-template-v1.md) - PodTemplate describes a template for creating copies of a predefined pod.
* [ReplicationController](replication-controller-v1.md) - ReplicationController represents the configuration of a replication controller.
* [ResourceQuota](resource-quota-v1.md) - ResourceQuota sets aggregate quota restrictions enforced per namespace
* [Secret](secret-v1.md) - Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than MaxSecretSize bytes.
* [Service](service-v1.md) - Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.
* [ServiceAccount](service-account-v1.md) - ServiceAccount binds together: * a name, understood by users, and perhaps by peripheral systems, for an identity * a principal that can be authenticated and authorized * a set of secrets
