---
type: reference
title: Articles on dockershim Removal and on Using CRI-compatible Runtimes
description: This is a list of articles and other pages that are either about the Kubernetes' deprecation and removal of _dockershim_, or about using CRI-compatible container runtimes, in connection with that removal.
resource: https://kubernetes.io/docs/reference/node/topics-on-dockershim-and-cri-compatible-runtimes/
tags: [kubernetes, reference, node]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/node/topics-on-dockershim-and-cri-compatible-runtimes.md
    title: 'Kubernetes website source: reference/node/topics-on-dockershim-and-cri-compatible-runtimes.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/node/topics-on-dockershim-and-cri-compatible-runtimes/
    title: Articles on dockershim Removal and on Using CRI-compatible Runtimes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/node/topics-on-dockershim-and-cri-compatible-runtimes.md
---

# Articles on dockershim Removal and on Using CRI-compatible Runtimes

This is a list of articles and other pages that are either
about the Kubernetes' deprecation and removal of _dockershim_,
or about using CRI-compatible container runtimes,
in connection with that removal.

## Kubernetes project

* Kubernetes blog: [Dockershim Removal FAQ](https://kubernetes.io/blog/2020/12/02/dockershim-faq/) (originally published 2020/12/02)

* Kubernetes blog: [Updated: Dockershim Removal FAQ](https://kubernetes.io/blog/2022/02/17/dockershim-faq/) (updated published 2022/02/17)

* Kubernetes blog: [Kubernetes is Moving on From Dockershim: Commitments and Next Steps](https://kubernetes.io/blog/2022/01/07/kubernetes-is-moving-on-from-dockershim/) (published 2022/01/07)

* Kubernetes blog: [Dockershim removal is coming. Are you ready?](https://kubernetes.io/blog/2021/11/12/are-you-ready-for-dockershim-removal/) (published 2021/11/12)

* Kubernetes documentation: [Migrating from dockershim](/tasks/administer-cluster/migrating-from-dockershim.md)

* Kubernetes documentation: [Container Runtimes](https://kubernetes.io/docs/setup/production-environment/container-runtimes/)

* Kubernetes enhancement proposal: [KEP-2221: Removing dockershim from kubelet](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/2221-remove-dockershim/README.md)

* Kubernetes enhancement proposal issue: [Removing dockershim from kubelet](https://github.com/kubernetes/enhancements/issues/2221) (_k/enhancements#2221_)

You can provide feedback via the GitHub issue [**Dockershim removal feedback & issues**](https://github.com/kubernetes/kubernetes/issues/106917). (_k/kubernetes/#106917_)

## External sources {#third-party}

* Amazon Web Services EKS documentation: [Amazon EKS is ending support for Dockershim](https://docs.aws.amazon.com/eks/latest/userguide/dockershim-deprecation.html)

* CNCF conference video: [Lessons Learned Migrating Kubernetes from Docker to containerd Runtime](https://www.youtube.com/watch?v=uDOu6rK4yOk) (Ana Caylin, at KubeCon Europe 2019)

* Docker.com blog: [What developers need to know about Docker, Docker Engine, and Kubernetes v1.20](https://www.docker.com/blog/what-developers-need-to-know-about-docker-docker-engine-and-kubernetes-v1-20/) (published 2020/12/04)

* "_Google Open Source_" channel on YouTube: [Learn Kubernetes with Google - Migrating from Dockershim to Containerd](https://youtu.be/fl7_4hjT52g)

* Microsoft Apps on Azure blog: [Dockershim deprecation and AKS](https://techcommunity.microsoft.com/t5/apps-on-azure-blog/dockershim-deprecation-and-aks/ba-p/3055902) (published 2022/01/21)

* Mirantis blog: [The Future of Dockershim is cri-dockerd](https://www.mirantis.com/blog/the-future-of-dockershim-is-cri-dockerd/) (published 2021/04/21)

* Mirantis: [Mirantis/cri-dockerd](https://mirantis.github.io/cri-dockerd/) Official Documentation

* Tripwire: [How Dockershim’s Forthcoming Deprecation Affects Your Kubernetes](https://www.tripwire.com/state-of-security/security-data-protection/cloud/how-dockershim-forthcoming-deprecation-affects-your-kubernetes/) (published 2021/07/01)
