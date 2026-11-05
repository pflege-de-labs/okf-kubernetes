# Configuration APIs

* [Client Authentication (v1)](client-authentication.v1.md) - HTTP transports.</p>
* [Client Authentication (v1beta1)](client-authentication.v1beta1.md) - HTTP transports.</p>
* [Event Rate Limit Configuration (v1alpha1)](apiserver-eventratelimit.v1alpha1.md) - controller.</p>
* [Image Policy API (v1alpha1)](imagepolicy.v1alpha1.md) - More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata</p> Refer to the Kubernetes API documentation for the fields of the <code>metadata</code> field.</td> </tr> <tr><td><code>spec</code><br/> <a href="#imagepolicy-k8s-io-v1alpha1-ImageReviewSpec"...
* [Kubelet Configuration (v1)](kubelet-config.v1.md) - each exec credential provider.
* [Kubelet Configuration (v1alpha1)](kubelet-config.v1alpha1.md) - each exec credential provider.
* [Kubelet Configuration (v1beta1)](kubelet-config.v1beta1.md) - Only available when the LoggingAlphaOptions feature gate is enabled.</p> </td> </tr> <tr><td><code>json</code> <B>[Required]</B><br/> <a href="#JSONOptions"><code>JSONOptions</code></a> </td> <td> <p>[Alpha] JSON contains options for logging format &quot;json&quot;.
* [Kubelet CredentialProvider (v1)](kubelet-credentialprovider.v1.md) - Kubelet will pass this request object to the plugin via stdin.
* [WebhookAdmission Configuration (v1)](apiserver-webhookadmission.v1.md) - configurations to be loaded at startup.
* [kube-apiserver Admission (v1)](apiserver-admission.v1.md) - otherwise identical (parallel requests, requests when earlier requests did not modify etc) The UID is meant to track the round trip (request/response) between the KAS and the WebHook, not the user request.
* [kube-apiserver Audit Configuration (v1)](apiserver-audit.v1.md) - For non-resource requests, this is the lower-cased HTTP method.</p> </td> </tr> <tr><td><code>user</code> <B>[Required]</B><br/> <a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#userinfo-v1-authentication-k8s-io"><code>authentication/v1.UserInfo</code></a> </td> <td> ...
* [kube-apiserver Configuration (v1)](apiserver-config.v1.md) - The connection is insecure, and does not currently support TLS.
* [kube-apiserver Configuration (v1alpha1)](apiserver-config.v1alpha1.md) - The connection is insecure, and does not currently support TLS.
* [kube-apiserver Configuration (v1beta1)](apiserver-config.v1beta1.md) - The connection is insecure, and does not currently support TLS.
* [kube-controller-manager Configuration (v1alpha1)](kube-controller-manager-config.v1alpha1.md) - default value of 'application/json'.
* [kube-proxy Configuration (v1alpha1)](kube-proxy-config.v1alpha1.md) - Only available when the LoggingAlphaOptions feature gate is enabled.</p> </td> </tr> <tr><td><code>json</code> <B>[Required]</B><br/> <a href="#JSONOptions"><code>JSONOptions</code></a> </td> <td> <p>[Alpha] JSON contains options for logging format &quot;json&quot;.
* [kube-scheduler Configuration (v1)](kube-scheduler-config.v1.md) - default value of 'application/json'.
* [kubeadm Configuration (v1beta3)](kubeadm-config.v1beta3.md) - This version improves on the v1beta2 format by fixing some minor issues and adding a few new fields.</p> <p>A list of changes since v1beta2:</p> <ul> <li>The deprecated &quot;ClusterConfiguration.useHyperKubeImage&quot; field has been removed.
* [kubeadm Configuration (v1beta4)](kubeadm-config.v1beta4.md) - This version improves on the v1beta3 format by fixing some minor issues and adding a few new fields.</p> <p>A list of changes since v1beta3:</p> <p>v1.35:</p> <ul> <li>Add <code>httpEndpoints</code> field to <code>ClusterConfiguration.etcd.externalEtcd</code> that can be used to configure the HTT...
* [kubeconfig (v1)](kubeconfig.v1.md) - TODO(jlowdermilk): remove this after eliminating downstream dependencies.</p> </td> </tr> <tr><td><code>apiVersion</code><br/> <code>string</code> </td> <td> <p>Legacy field from pkg/api/types.go TypeMeta.
* [kuberc (v1alpha1)](kuberc.v1alpha1.md) - This is especially useful, when user doesn't want to explicitly set flags each time.</p> </td> </tr> <tr><td><code>aliases</code> <B>[Required]</B><br/> <a href="#kubectl-config-k8s-io-v1alpha1-AliasOverride"><code>[]AliasOverride</code></a> </td> <td> <p>aliases allow defining command aliases fo...
* [kuberc (v1beta1)](kuberc.v1beta1.md) - This is especially useful, when user doesn't want to explicitly set options each time.</p> </td> </tr> <tr><td><code>aliases</code> <B>[Required]</B><br/> <a href="#kubectl-config-k8s-io-v1beta1-AliasOverride"><code>[]AliasOverride</code></a> </td> <td> <p>aliases allow defining command aliases f...
