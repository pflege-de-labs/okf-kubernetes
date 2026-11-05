# API Access Control

* [Manifest-Based Admission Control](manifest-admission-control.md) - This page provides an overview of manifest-based admission control configuration.
* [Mutating Admission Policy](mutating-admission-policy.md) - This page provides an overview of _MutatingAdmissionPolicies_.
* [Validating Admission Policy](validating-admission-policy.md) - This page provides an overview of Validating Admission Policy.
* [Authenticating](authentication.md) - This page provides an overview of authentication in Kubernetes, with a focus on authentication to the Kubernetes API.
* [Authenticating with Bootstrap Tokens](bootstrap-tokens.md) - Bootstrap tokens are a simple bearer token that is meant to be used when creating new clusters or joining new nodes to an existing cluster.
* [Authorization](authorization.md) - Details of Kubernetes authorization mechanisms and supported authorization modes.
* [Using RBAC Authorization](rbac.md) - Role-based access control (RBAC) is a method of regulating access to computer or network resources based on the roles of individual users within your organization.
* [Using Node Authorization](node.md) - Node authorization is a special-purpose authorization mode that specifically authorizes API requests made by kubelets.
* [Webhook Mode](webhook.md) - A WebHook is an HTTP callback: an HTTP POST that occurs when something happens; a simple event-notification via HTTP POST.
* [Using ABAC Authorization](abac.md) - Attribute-based access control (ABAC) defines an access control paradigm whereby access rights are granted to users through the use of policies which combine attributes together.
* [Admission Control in Kubernetes](admission-controllers.md) - This page provides an overview of _admission controllers_.
* [Dynamic Admission Control](extensible-admission-controllers.md) - In addition to compiled-in admission plugins, admission plugins can be developed as extensions and run as webhooks configured at runtime.
* [Managing Service Accounts](service-accounts-admin.md) - A _ServiceAccount_ provides an identity for processes that run in a Pod.
* [User Impersonation](user-impersonation.md) - User _impersonation_ is a method of allowing authenticated users to act as another user, group, or service account through HTTP headers.
* [Certificates and Certificate Signing Requests](certificate-signing-requests.md) - Kubernetes certificate and trust bundle APIs enable automation of X.509 credential provisioning by providing a programmatic interface for clients of the Kubernetes API to request and obtain X.509 certificates from a Certificate Authority (CA).
* [Mapping PodSecurityPolicies to Pod Security Standards](psp-to-pod-security-standards.md) - The tables below enumerate the configuration parameters on `PodSecurityPolicy` objects, whether the field mutates and/or validates pods, and how the configuration values map to the Pod Security Standards.
* [Kubelet authentication/authorization](kubelet-authn-authz.md) - A kubelet's HTTPS endpoint exposes APIs which give access to data of varying sensitivity, and allow you to perform operations with varying levels of power on the node and within containers.
* [TLS bootstrapping](kubelet-tls-bootstrapping.md) - In a Kubernetes cluster, the components on the worker nodes - kubelet and kube-proxy - need to communicate with Kubernetes control plane components, specifically kube-apiserver.
