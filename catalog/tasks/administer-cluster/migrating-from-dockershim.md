---
type: task
title: Migrating from dockershim
description: This section presents information you need to know when migrating from dockershim to other container runtimes.
resource: https://kubernetes.io/docs/tasks/administer-cluster/migrating-from-dockershim/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/administer-cluster/migrating-from-dockershim/_index.md
    title: 'Kubernetes website source: tasks/administer-cluster/migrating-from-dockershim/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/migrating-from-dockershim/
    title: Migrating from dockershim | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/administer-cluster/migrating-from-dockershim/_index.md
---

# Migrating from dockershim

This section presents information you need to know when migrating from
dockershim to other container runtimes.

Since the announcement of [dockershim deprecation](https://kubernetes.io/blog/2020/12/08/kubernetes-1-20-release-announcement/#dockershim-deprecation)
in Kubernetes 1.20, there were questions on how this will affect various workloads and Kubernetes
installations. Our [Dockershim Removal FAQ](https://kubernetes.io/blog/2022/02/17/dockershim-faq/) is there to help you
to understand the problem better.

Dockershim was removed from Kubernetes with the release of v1.24.
If you use Docker Engine via dockershim as your container runtime and wish to upgrade to v1.24,
it is recommended that you either migrate to another runtime or find an alternative means to obtain Docker Engine support.
Check out the [container runtimes](https://kubernetes.io/docs/setup/production-environment/container-runtimes/)
section to know your options.

The version of Kubernetes with dockershim (1.23) is out of support and the v1.24
will run out of support [soon](/releases.md#release-v1-24). Make sure to
[report issues](https://github.com/kubernetes/kubernetes/issues) you encountered
with the migration so the issues can be fixed in a timely manner and your cluster would be
ready for dockershim removal. After v1.24 running out of support, you will need
to contact your Kubernetes provider for support or upgrade multiple versions at a time
if there are critical issues affecting your cluster.

Your cluster might have more than one kind of node, although this is not a common
configuration.

These tasks will help you to migrate:

* [Check whether Dockershim removal affects you](/tasks/administer-cluster/migrating-from-dockershim/check-if-dockershim-removal-affects-you.md)
* [Migrating telemetry and security agents from dockershim](/tasks/administer-cluster/migrating-from-dockershim/migrating-telemetry-and-security-agents.md)

## What's next

* Check out [container runtimes](https://kubernetes.io/docs/setup/production-environment/container-runtimes/)
  to understand your options for an alternative.
* If you find a defect or other technical concern relating to migrating away from dockershim,
  you can [report an issue](https://github.com/kubernetes/kubernetes/issues/new/choose)
  to the Kubernetes project.
