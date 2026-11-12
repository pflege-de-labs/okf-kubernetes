---
type: tool-reference
title: WebhookAdmission Configuration (v1)
description: configurations to be loaded at startup.
resource: https://kubernetes.io/docs/reference/config-api/apiserver-webhookadmission.v1/
tags: [kubernetes, reference, config-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/config-api/apiserver-webhookadmission.v1.md
    title: 'Kubernetes website source: reference/config-api/apiserver-webhookadmission.v1.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/config-api/apiserver-webhookadmission.v1/
    title: WebhookAdmission Configuration (v1) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/config-api/apiserver-webhookadmission.v1.md
---

# WebhookAdmission Configuration (v1)

<p>Package v1 is the v1 version of the API.</p>

## Resource Types 

- [WebhookAdmission](#apiserver-config-k8s-io-v1-WebhookAdmission)
  

## `WebhookAdmission`     {#apiserver-config-k8s-io-v1-WebhookAdmission}
    

<p>WebhookAdmission provides configuration for the webhook admission controller.</p>

<table class="table">
<thead><tr><th width="30%">Field</th><th>Description</th></tr></thead>
<tbody>
    
<tr><td><code>apiVersion</code><br/>string</td><td><code>apiserver.config.k8s.io/v1</code></td></tr>
<tr><td><code>kind</code><br/>string</td><td><code>WebhookAdmission</code></td></tr>
    
  
<tr><td><code>kubeConfigFile</code> <B>[Required]</B><br/>
<code>string</code>
</td>
<td>
   <p>KubeConfigFile is the path to the kubeconfig file.</p>
</td>
</tr>
<tr><td><code>staticManifestsDir</code><br/>
<code>string</code>
</td>
<td>
   <p>StaticManifestsDir is the path to a directory containing static webhook
configurations to be loaded at startup. Files with extensions .yaml,
.yml, and .json are read. Only admissionregistration.k8s.io/v1
ValidatingWebhookConfiguration and MutatingWebhookConfiguration
resources are supported.
Using this field requires the ManifestBasedAdmissionControlConfig
feature gate to be enabled.</p>
</td>
</tr>
</tbody>
</table>
