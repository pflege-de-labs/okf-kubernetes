---
type: section
title: Configuration
description: Configuration mechanisms within Kubernetes.
resource: https://kubernetes.io/docs/concepts/configuration/
tags: [kubernetes, concepts]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/configuration/_index.md
    title: 'Kubernetes website source: concepts/configuration/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/configuration/
    title: Configuration | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/configuration/_index.md
---

# Configuration

Kubernetes provides ways to _separate configuration from code_. This is a common
practice that aligns with DevOps as a practice, and with cloud native architecture.

If you have a [container image](/reference/glossary/image.md) that
allows separate configuration, you can deploy the **same** application code in
different contexts. For example, you can run tests against the built image, and
then run that exact same image in a production context. Doing that gives you
better confidence in your testing, compared with if you deployed the application
one way for tests and a different way in production.

To learn about configuring the `kubectl` command-line tool,
read [configure access to multiple clusters](/tasks/access-application-cluster/configure-access-multiple-clusters.md).

## Kubernetes configuration-related APIs {#api-kinds}

Kubernetes provides two main API kinds that you can use to store configuration,
ready for a Pod (or other component) to load and use:
[ConfigMap](/concepts/configuration/configmap.md), and
[Secret](/concepts/configuration/secret.md).

The special-use
[CertificateSigningRequest](/reference/access-authn-authz/certificate-signing-requests.md#certificate-signing-requests)
and
[ClusterTrustBundle](/reference/access-authn-authz/certificate-signing-requests.md#cluster-trust-bundles)
API kinds also hold configuration data that's relevant to some specific cases,
such as TLS configuration.

All of these configuration mechanisms, even for certificates, are designed so that configuration
can be reused and shared across multiple Pods. You typically don't have to write the configuration
out once per Pod, and you should design applications where possible so that they can work with
shared configuration.

To learn about using configuration in your applications, read
[inject data into applications](/tasks/inject-data-application.md).

### ConfigMaps

A ConfigMap is an API object used to store non-confidential data in key-value pairs.
[Pods](/reference/glossary/pod.md) can consume ConfigMaps as
environment variables, command-line arguments, or as configuration files in a
[volume](/reference/glossary/volume.md).

 

A ConfigMap allows you to decouple environment-specific configuration from your [container images](/reference/glossary/image.md), so that your applications are easily portable.

### Secrets

A Secret is an object that contains an amount of confidential data,
such as a password, a token, or a key.

Secrets are similar to [ConfigMaps](/reference/glossary/configmap.md)
but are specifically intended to hold confidential data. Read the page about [Secrets](/concepts/configuration/secret.md)
to learn about the benefits (and limitations) around information security and Secrets.

## Configuration via sidecar containers or init containers {#config-helper-overview}

You can also use an 
[init container](/reference/glossary/init-container.md) or a
[sidecar container](/reference/glossary/sidecar-container.md)
to provide configuration to your workload.

This is useful if you already have some means, outside of Kubernetes, to store
configuration or security keys or other information that your containers must
use, and that shouldn't be stored in container images.

### Sidecar container helper {#config-helper-sidecar}

A _sidecar_ configuration helper runs in the same Pod as the application container, so
each Pod has its own sidecar. A typical configuration sidecar fetches configuration data over the
Pod network, and then writes it into a [volume](/reference/glossary/volume.md)
that is mounted into both containers.

There are many variations on this basic pattern, that still have a helper
container writing configuration for the app container to consume.

### Init container helper {#config-helper-init-container}

Pods can have _init_ containers that start before the main application. Unlike sidecars,
init containers complete before the Pod's main (application) containers start up, so the
configuration only gets set once.

You have two main options for configuration using init containers:

#### Configuration via the filesystem {#init-container-shared-volume}

With this option, the init container fetches configuration and writes the configuration to
a file, or several files, typically to a Pod-local
[volume](/reference/glossary/volume.md).

#### Configuration via environment variables {#init-container-env-files}

**Feature state:** Kubernetes v1.35 [beta] (feature gate `EnvFiles`, enabled by default)

This is done using a local volume, but only the init container mounts that volume. The kubelet then reads from that volume before starting the app container.

Read [Define Environment Variable Values Using An Init Container](/tasks/inject-data-application/define-environment-variable-via-file.md)
to learn more.

## Cluster configuration

This section of the documentation is about configuration. However,
if you are doing platform-level or cluster-level configuration, read
the [cluster administration](/concepts/cluster-administration.md) pages
and look through the list of [cluster administration tasks](/tasks/administer-cluster.md).

## What's next

The following links are all relevant to the overall idea of _configuration_ and Kubernetes:

* [Configure Pods and Containers](/tasks/configure-pod-container.md)
* [Inject Data Into Applications](/tasks/inject-data-application.md)
* [Policies](/concepts/policy.md), another form of configuration
* [Management of Kubernetes Objects Using Configuration Files](/tasks/manage-kubernetes-objects/declarative-config.md),
  which is about Kubernetes [manifests](/reference/glossary/manifest.md)
  * [Kubernetes Object Management](/concepts/overview/working-with-objects/object-management.md) is also about object management / configuration
* [Configure Redis Using A ConfigMap](https://kubernetes.io/docs/tutorials/configuration/configure-redis-using-configmap/)
  (a tutorial)
* [Updating Configuration via a ConfigMap](https://kubernetes.io/docs/tutorials/configuration/updating-configuration-via-a-configmap/)
  (another tutorial)

and you can review the [list of configuration file formats](/reference/config-api.md)
in the reference section.

Within **this** section of the documentation, you can read about:
