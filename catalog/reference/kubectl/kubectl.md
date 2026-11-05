---
type: tool-reference
title: kubectl
description: Command line tool for communicating with a Kubernetes cluster's control plane, using the Kubernetes API.
resource: https://kubernetes.io/docs/reference/kubectl/kubectl/
tags: [kubernetes, reference, kubectl]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubectl/kubectl.md
    title: 'Kubernetes website source: reference/kubectl/kubectl.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubectl/kubectl/
    title: kubectl | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubectl/kubectl.md
---

# kubectl

## Synopsis

Command line tool for communicating with a Kubernetes cluster's
[control plane](/reference/glossary/control-plane.md),
using the Kubernetes API.

Find more information in [Command line tool](/reference/kubectl.md) (`kubectl`).

```shell
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
<td colspan="2">--add-dir-header</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">If true, adds the file directory to the header of the log messages</td>
</tr>

<tr>
<td colspan="2">--alsologtostderr</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">log to standard error as well as files</td>
</tr>

<tr>
<td colspan="2">--as string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Username to impersonate for the operation</td>
</tr>

<tr>
<td colspan="2">--as-group stringArray</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Group to impersonate for the operation, this flag can be repeated to specify multiple groups.</td>
</tr>

<tr>
<td colspan="2">--azure-container-registry-config string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Path to the file containing Azure container registry configuration information.</td>
</tr>

<tr>
<td colspan="2">--cache-dir string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "$HOME/.kube/cache"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Default cache directory</td>
</tr>

<tr>
<td colspan="2">--certificate-authority string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Path to a cert file for the certificate authority</td>
</tr>

<tr>
<td colspan="2">--client-certificate string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Path to a client certificate file for TLS</td>
</tr>

<tr>
<td colspan="2">--client-key string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Path to a client key file for TLS</td>
</tr>

<tr>
<td colspan="2">--cloud-provider-gce-l7lb-src-cidrs cidrs&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 130.211.0.0/22,35.191.0.0/16</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">CIDRs opened in GCE firewall for L7 LB traffic proxy & health checks</td>
</tr>

<tr>
<td colspan="2">--cloud-provider-gce-lb-src-cidrs cidrs&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 130.211.0.0/22,209.85.152.0/22,209.85.204.0/22,35.191.0.0/16</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">CIDRs opened in GCE firewall for L4 LB traffic proxy & health checks</td>
</tr>

<tr>
<td colspan="2">--cluster string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">The name of the kubeconfig cluster to use</td>
</tr>

<tr>
<td colspan="2">--context string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">The name of the kubeconfig context to use</td>
</tr>

<tr>
<td colspan="2">--default-not-ready-toleration-seconds int&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 300</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Indicates the tolerationSeconds of the toleration for notReady:NoExecute that is added by default to every pod that does not already have such a toleration.</td>
</tr>

<tr>
<td colspan="2">--default-unreachable-toleration-seconds int&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 300</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Indicates the tolerationSeconds of the toleration for unreachable:NoExecute that is added by default to every pod that does not already have such a toleration.</td>
</tr>

<tr>
<td colspan="2">-h, --help</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">help for kubectl</td>
</tr>

<tr>
<td colspan="2">--insecure-skip-tls-verify</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure</td>
</tr>

<tr>
<td colspan="2">--kubeconfig string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Path to the kubeconfig file to use for CLI requests.</td>
</tr>

<tr>
<td colspan="2">--log-backtrace-at traceLocation&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: :0</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">when logging hits line file:N, emit a stack trace</td>
</tr>

<tr>
<td colspan="2">--log-dir string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">If non-empty, write log files in this directory</td>
</tr>

<tr>
<td colspan="2">--log-file string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">If non-empty, use this log file</td>
</tr>

<tr>
<td colspan="2">--log-file-max-size uint&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 1800</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Defines the maximum size a log file can grow to. Unit is megabytes. If the value is 0, the maximum file size is unlimited.</td>
</tr>

<tr>
<td colspan="2">--log-flush-frequency duration&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 5s</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Maximum number of seconds between log flushes</td>
</tr>

<tr>
<td colspan="2">--logtostderr&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: true</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">log to standard error instead of files</td>
</tr>

<tr>
<td colspan="2">--match-server-version</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Require server version to match client version</td>
</tr>

<tr>
<td colspan="2">-n, --namespace string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">If present, the namespace scope for this CLI request</td>
</tr>

<tr>
<td colspan="2">--one-output</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">If true, only write logs to their native severity level (vs also writing to each lower severity level)</td>
</tr>

<tr>
<td colspan="2">--password string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Password for basic authentication to the API server</td>
</tr>

<tr>
<td colspan="2">--profile string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "none"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Name of profile to capture. One of (none|cpu|heap|goroutine|threadcreate|block|mutex)</td>
</tr>

<tr>
<td colspan="2">--profile-output string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "profile.pprof"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Name of the file to write the profile to</td>
</tr>

<tr>
<td colspan="2">--request-timeout string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "0"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests.</td>
</tr>

<tr>
<td colspan="2">-s, --server string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">The address and port of the Kubernetes API server</td>
</tr>

<tr>
<td colspan="2">--skip-headers</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">If true, avoid header prefixes in the log messages</td>
</tr>

<tr>
<td colspan="2">--skip-log-headers</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">If true, avoid headers when opening log files</td>
</tr>

<tr>
<td colspan="2">--stderrthreshold severity&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 2</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">logs at or above this threshold go to stderr</td>
</tr>

<tr>
<td colspan="2">--tls-server-name string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used</td>
</tr>

<tr>
<td colspan="2">--token string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Bearer token for authentication to the API server</td>
</tr>

<tr>
<td colspan="2">--user string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">The name of the kubeconfig user to use</td>
</tr>

<tr>
<td colspan="2">--username string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Username for basic authentication to the API server</td>
</tr>

<tr>
<td colspan="2">-v, --v Level</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">number for the log level verbosity</td>
</tr>

<tr>
<td colspan="2">--version version[=true]</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Print version information and quit</td>
</tr>

<tr>
<td colspan="2">--vmodule moduleSpec</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">comma-separated list of pattern=N settings for file-filtered logging</td>
</tr>

<tr>
<td colspan="2">--warnings-as-errors</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Treat warnings received from the server as errors and exit with a non-zero exit code</td>
</tr>

</tbody>
</table>

## Environment variables

<table style="width: 100%; table-layout: fixed;">
<colgroup>
<col span="1" style="width: 10px;" />
<col span="1" />
</colgroup>
<tbody>

<tr>
<td colspan="2">KUBECONFIG</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Path to the kubectl configuration ("kubeconfig") file. Default: "$HOME/.kube/config"</td>
</tr>

<tr>
<td colspan="2">KUBECTL_EXPLAIN_OPENAPIV3</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">Toggles whether calls to `kubectl explain` use the new OpenAPIv3 data source available. OpenAPIV3 is enabled by default since Kubernetes 1.24.
</td>
</tr>

<tr>
<td colspan="2">KUBECTL_ENABLE_CMD_SHADOW</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">When set to true, external plugins can be used as subcommands for builtin commands if subcommand does not exist. In alpha stage, this feature can only be used for create command(e.g. kubectl create networkpolicy).
</td>
</tr>

<tr>
<td colspan="2">KUBECTL_PORT_FORWARD_WEBSOCKETS</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">When set to true, the kubectl port-forward command will attempt to stream using the websockets protocol. If the upgrade to websockets fails, the commands will fallback to use the current SPDY protocol.
</td>
</tr>

<tr>
<td colspan="2">KUBECTL_REMOTE_COMMAND_WEBSOCKETS</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">When set to true, the kubectl exec, cp, and attach commands will attempt to stream using the websockets protocol. If the upgrade to websockets fails, the commands will fallback to use the current SPDY protocol.
</td>
</tr>

<tr>
<td colspan="2">KUBECTL_KUBERC</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;">When set to true, kuberc file is taken into account to define user specific preferences.
</td>
</tr>

</tbody>
</table>

## See Also

* [kubectl annotate](/reference/kubectl/generated/kubectl_annotate.md) - Update the annotations on a resource
* [kubectl api-resources](/reference/kubectl/generated/kubectl_api-resources.md) - Print the supported API resources on the server
* [kubectl api-versions](/reference/kubectl/generated/kubectl_api-versions.md) - Print the supported API versions on the server,
  in the form of "group/version"
* [kubectl apply](/reference/kubectl/generated/kubectl_apply.md) - Apply a configuration to a resource by filename or stdin
* [kubectl attach](/reference/kubectl/generated/kubectl_attach.md) - Attach to a running container
* [kubectl auth](/reference/kubectl/generated/kubectl_auth.md) - Inspect authorization
* [kubectl autoscale](/reference/kubectl/generated/kubectl_autoscale.md) - Auto-scale a Deployment, ReplicaSet, or ReplicationController
* [kubectl certificate](/reference/kubectl/generated/kubectl_certificate.md) - Modify certificate resources.
* [kubectl cluster-info](/reference/kubectl/generated/kubectl_cluster-info.md) - Display cluster info
* [kubectl completion](/reference/kubectl/generated/kubectl_completion.md) - Output shell completion code for the specified shell (bash or zsh)
* [kubectl config](/reference/kubectl/generated/kubectl_config.md) - Modify kubeconfig files
* [kubectl cordon](/reference/kubectl/generated/kubectl_cordon.md) - Mark node as unschedulable
* [kubectl cp](/reference/kubectl/generated/kubectl_cp.md) - Copy files and directories to and from containers.
* [kubectl create](/reference/kubectl/generated/kubectl_create.md) - Create a resource from a file or from stdin.
* [kubectl debug](/reference/kubectl/generated/kubectl_debug.md) - Create debugging sessions for troubleshooting workloads and nodes
* [kubectl delete](/reference/kubectl/generated/kubectl_delete.md) - Delete resources by filenames,
  stdin, resources and names, or by resources and label selector
* [kubectl describe](/reference/kubectl/generated/kubectl_describe.md) - Show details of a specific resource or group of resources
* [kubectl diff](/reference/kubectl/generated/kubectl_diff.md) - Diff live version against would-be applied version
* [kubectl drain](/reference/kubectl/generated/kubectl_drain.md) - Drain node in preparation for maintenance
* [kubectl edit](/reference/kubectl/generated/kubectl_edit.md) - Edit a resource on the server
* [kubectl events](/reference/kubectl/generated/kubectl_events.md)  - List events
* [kubectl exec](/reference/kubectl/generated/kubectl_exec.md) - Execute a command in a container
* [kubectl explain](/reference/kubectl/generated/kubectl_explain.md) - Documentation of resources
* [kubectl expose](/reference/kubectl/generated/kubectl_expose.md) - Take a replication controller,
  service, deployment or pod and expose it as a new Kubernetes Service
* [kubectl get](/reference/kubectl/generated/kubectl_get.md) - Display one or many resources
* [kubectl kustomize](/reference/kubectl/generated/kubectl_kustomize.md) - Build a kustomization
  target from a directory or a remote url.
* [kubectl label](/reference/kubectl/generated/kubectl_label.md) - Update the labels on a resource
* [kubectl logs](/reference/kubectl/generated/kubectl_logs.md) - Print the logs for a container in a pod
* [kubectl options](/reference/kubectl/generated/kubectl_options.md) - Print the list of flags inherited by all commands
* [kubectl patch](/reference/kubectl/generated/kubectl_patch.md) - Update field(s) of a resource
* [kubectl plugin](/reference/kubectl/generated/kubectl_plugin.md) - Provides utilities for interacting with plugins.
* [kubectl port-forward](/reference/kubectl/generated/kubectl_port-forward.md) - Forward one or more local ports to a pod
* [kubectl proxy](/reference/kubectl/generated/kubectl_proxy.md) - Run a proxy to the Kubernetes API server
* [kubectl replace](/reference/kubectl/generated/kubectl_replace.md) - Replace a resource by filename or stdin
* [kubectl rollout](/reference/kubectl/generated/kubectl_rollout.md) - Manage the rollout of a resource
* [kubectl run](/reference/kubectl/generated/kubectl_run.md) - Run a particular image on the cluster
* [kubectl scale](/reference/kubectl/generated/kubectl_scale.md) - Set a new size for a Deployment, ReplicaSet or Replication Controller
* [kubectl set](/reference/kubectl/generated/kubectl_set.md) - Set specific features on objects
* [kubectl taint](/reference/kubectl/generated/kubectl_taint.md) - Update the taints on one or more nodes
* [kubectl top](/reference/kubectl/generated/kubectl_top.md) - Display Resource (CPU/Memory/Storage) usage.
* [kubectl uncordon](/reference/kubectl/generated/kubectl_uncordon.md) - Mark node as schedulable
* [kubectl version](/reference/kubectl/generated/kubectl_version.md) - Print the client and server version information
* [kubectl wait](/reference/kubectl/generated/kubectl_wait.md) - Experimental: Wait for a specific condition on one or many resources.
