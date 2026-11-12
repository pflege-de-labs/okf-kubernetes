---
type: tool-reference
title: kubectl
description: kubectl controls the Kubernetes cluster manager
resource: https://kubernetes.io/docs/reference/kubectl/generated/kubectl/
tags: [kubernetes, reference, kubectl]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubectl/generated/kubectl.md
    title: 'Kubernetes website source: reference/kubectl/generated/kubectl.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubectl/generated/kubectl/
    title: kubectl | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubectl/generated/kubectl.md
---

# kubectl

## Synopsis

kubectl controls the Kubernetes cluster manager.

 Find more information at: https://kubernetes.io/docs/reference/kubectl/

```
kubectl [flags]
```

## Options

   <table style="width: 100%; table-layout: fixed;">
<colgroup>
<col span="1" style="width: 10px;" />
<col span="1" />
</colgroup>
<tbody>

<tr>
<td colspan="2">--as string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Username to impersonate for the operation. User could be a regular user or a service account in a namespace.</p></td>
</tr>

<tr>
<td colspan="2">--as-group strings</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Group to impersonate for the operation, this flag can be repeated to specify multiple groups.</p></td>
</tr>

<tr>
<td colspan="2">--as-uid string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>UID to impersonate for the operation.</p></td>
</tr>

<tr>
<td colspan="2">--as-user-extra strings</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>User extras to impersonate for the operation, this flag can be repeated to specify multiple values for the same key.</p></td>
</tr>

<tr>
<td colspan="2">--cache-dir string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "$HOME/.kube/cache"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Default cache directory</p></td>
</tr>

<tr>
<td colspan="2">--certificate-authority string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to a cert file for the certificate authority</p></td>
</tr>

<tr>
<td colspan="2">--client-certificate string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to a client certificate file for TLS</p></td>
</tr>

<tr>
<td colspan="2">--client-key string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to a client key file for TLS</p></td>
</tr>

<tr>
<td colspan="2">--cluster string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The name of the kubeconfig cluster to use</p></td>
</tr>

<tr>
<td colspan="2">--context string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The name of the kubeconfig context to use</p></td>
</tr>

<tr>
<td colspan="2">--disable-compression</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>If true, opt-out of response compression for all requests to the server</p></td>
</tr>

<tr>
<td colspan="2">-h, --help</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for kubectl</p></td>
</tr>

<tr>
<td colspan="2">--insecure-skip-tls-verify</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure</p></td>
</tr>

<tr>
<td colspan="2">--kubeconfig string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to the kubeconfig file to use for CLI requests.</p></td>
</tr>

<tr>
<td colspan="2">--kuberc string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to the kuberc file to use for preferences. This can be disabled by exporting KUBECTL_KUBERC=false feature gate or turning off the feature KUBERC=off.</p></td>
</tr>

<tr>
<td colspan="2">--match-server-version</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Require server version to match client version</p></td>
</tr>

<tr>
<td colspan="2">-n, --namespace string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>If present, the namespace scope for this CLI request</p></td>
</tr>

<tr>
<td colspan="2">--password string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Password for basic authentication to the API server</p></td>
</tr>

<tr>
<td colspan="2">--profile string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "none"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Name of profile to capture. One of (none|cpu|heap|goroutine|threadcreate|block|mutex|trace)</p></td>
</tr>

<tr>
<td colspan="2">--profile-output string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "profile.pprof"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Name of the file to write the profile to</p></td>
</tr>

<tr>
<td colspan="2">--proxy-url string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Proxy URL to use for requests to the API server</p></td>
</tr>

<tr>
<td colspan="2">--request-timeout string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "0"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests.</p></td>
</tr>

<tr>
<td colspan="2">-s, --server string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The address and port of the Kubernetes API server</p></td>
</tr>

<tr>
<td colspan="2">--storage-driver-buffer-duration duration&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 1m0s</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Writes in the storage driver will be buffered for this duration, and committed to the non memory backends as a single transaction</p></td>
</tr>

<tr>
<td colspan="2">--storage-driver-db string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "cadvisor"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>database name</p></td>
</tr>

<tr>
<td colspan="2">--storage-driver-host string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "localhost:8086"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>database host:port</p></td>
</tr>

<tr>
<td colspan="2">--storage-driver-password string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "root"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>database password</p></td>
</tr>

<tr>
<td colspan="2">--storage-driver-secure</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>use secure connection with database</p></td>
</tr>

<tr>
<td colspan="2">--storage-driver-table string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "stats"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>table name</p></td>
</tr>

<tr>
<td colspan="2">--storage-driver-user string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "root"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>database username</p></td>
</tr>

<tr>
<td colspan="2">--tls-server-name string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used</p></td>
</tr>

<tr>
<td colspan="2">--token string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Bearer token for authentication to the API server</p></td>
</tr>

<tr>
<td colspan="2">--user string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The name of the kubeconfig user to use</p></td>
</tr>

<tr>
<td colspan="2">--username string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Username for basic authentication to the API server</p></td>
</tr>

<tr>
<td colspan="2">--version version[=true]</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>--version, --version=raw prints version information and quits; --version=vX.Y.Z... sets the reported version</p></td>
</tr>

<tr>
<td colspan="2">--warnings-as-errors</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Treat warnings received from the server as errors and exit with a non-zero exit code</p></td>
</tr>

</tbody>
</table>

## See Also

* [kubectl annotate](/reference/kubectl/generated/kubectl_annotate.md)	 - Update the annotations on a resource
* [kubectl api-resources](/reference/kubectl/generated/kubectl_api-resources.md)	 - Print the supported API resources on the server
* [kubectl api-versions](/reference/kubectl/generated/kubectl_api-versions.md)	 - Print the supported API versions on the server, in the form of "group/version"
* [kubectl apply](/reference/kubectl/generated/kubectl_apply.md)	 - Apply a configuration to a resource by file name or stdin
* [kubectl attach](/reference/kubectl/generated/kubectl_attach.md)	 - Attach to a running container
* [kubectl auth](/reference/kubectl/generated/kubectl_auth.md)	 - Inspect authorization
* [kubectl autoscale](/reference/kubectl/generated/kubectl_autoscale.md)	 - Auto-scale a deployment, replica set, stateful set, or replication controller
* [kubectl certificate](/reference/kubectl/generated/kubectl_certificate.md)	 - Modify certificate resources
* [kubectl cluster-info](/reference/kubectl/generated/kubectl_cluster-info.md)	 - Display cluster information
* [kubectl completion](/reference/kubectl/generated/kubectl_completion.md)	 - Output shell completion code for the specified shell (bash, zsh, fish, or powershell)
* [kubectl config](/reference/kubectl/generated/kubectl_config.md)	 - Modify kubeconfig files
* [kubectl cordon](/reference/kubectl/generated/kubectl_cordon.md)	 - Mark node as unschedulable
* [kubectl cp](/reference/kubectl/generated/kubectl_cp.md)	 - Copy files and directories to and from containers
* [kubectl create](/reference/kubectl/generated/kubectl_create.md)	 - Create a resource from a file or from stdin
* [kubectl debug](/reference/kubectl/generated/kubectl_debug.md)	 - Create debugging sessions for troubleshooting workloads and nodes
* [kubectl delete](/reference/kubectl/generated/kubectl_delete.md)	 - Delete resources by file names, stdin, resources and names, or by resources and label selector
* [kubectl describe](/reference/kubectl/generated/kubectl_describe.md)	 - Show details of a specific resource or group of resources
* [kubectl diff](/reference/kubectl/generated/kubectl_diff.md)	 - Diff the live version against a would-be applied version
* [kubectl drain](/reference/kubectl/generated/kubectl_drain.md)	 - Drain node in preparation for maintenance
* [kubectl edit](/reference/kubectl/generated/kubectl_edit.md)	 - Edit a resource on the server
* [kubectl events](/reference/kubectl/generated/kubectl_events.md)	 - List events
* [kubectl exec](/reference/kubectl/generated/kubectl_exec.md)	 - Execute a command in a container
* [kubectl explain](/reference/kubectl/generated/kubectl_explain.md)	 - Get documentation for a resource
* [kubectl expose](/reference/kubectl/generated/kubectl_expose.md)	 - Take a replication controller, service, deployment or pod and expose it as a new Kubernetes service
* [kubectl get](/reference/kubectl/generated/kubectl_get.md)	 - Display one or many resources
* [kubectl kuberc](/reference/kubectl/generated/kubectl_kuberc.md)	 - Manage kuberc configuration files
* [kubectl kustomize](/reference/kubectl/generated/kubectl_kustomize.md)	 - Build a kustomization target from a directory or URL
* [kubectl label](/reference/kubectl/generated/kubectl_label.md)	 - Update the labels on a resource
* [kubectl logs](/reference/kubectl/generated/kubectl_logs.md)	 - Print the logs for a container in a pod
* [kubectl options](/reference/kubectl/generated/kubectl_options.md)	 - Print the list of flags inherited by all commands
* [kubectl patch](/reference/kubectl/generated/kubectl_patch.md)	 - Update fields of a resource
* [kubectl plugin](/reference/kubectl/generated/kubectl_plugin.md)	 - Provides utilities for interacting with plugins
* [kubectl port-forward](/reference/kubectl/generated/kubectl_port-forward.md)	 - Forward one or more local ports to a pod
* [kubectl proxy](/reference/kubectl/generated/kubectl_proxy.md)	 - Run a proxy to the Kubernetes API server
* [kubectl replace](/reference/kubectl/generated/kubectl_replace.md)	 - Replace a resource by file name or stdin
* [kubectl rollout](/reference/kubectl/generated/kubectl_rollout.md)	 - Manage the rollout of a resource
* [kubectl run](/reference/kubectl/generated/kubectl_run.md)	 - Run a particular image on the cluster
* [kubectl scale](/reference/kubectl/generated/kubectl_scale.md)	 - Set a new size for a deployment, replica set, or replication controller
* [kubectl set](/reference/kubectl/generated/kubectl_set.md)	 - Set specific features on objects
* [kubectl taint](/reference/kubectl/generated/kubectl_taint.md)	 - Update the taints on one or more nodes
* [kubectl top](/reference/kubectl/generated/kubectl_top.md)	 - Display resource (CPU/memory) usage
* [kubectl uncordon](/reference/kubectl/generated/kubectl_uncordon.md)	 - Mark node as schedulable
* [kubectl version](/reference/kubectl/generated/kubectl_version.md)	 - Print the client and server version information
* [kubectl wait](/reference/kubectl/generated/kubectl_wait.md)	 - Wait for a specific condition on one or many resources
