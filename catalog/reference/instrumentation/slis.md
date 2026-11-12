---
type: reference
title: Kubernetes Component SLI Metrics
description: High-level indicators for measuring the reliability and performance of Kubernetes components.
resource: https://kubernetes.io/docs/reference/instrumentation/slis/
tags: [kubernetes, reference, instrumentation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/instrumentation/slis.md
    title: 'Kubernetes website source: reference/instrumentation/slis.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/instrumentation/slis/
    title: Kubernetes Component SLI Metrics | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/instrumentation/slis.md
---

# Kubernetes Component SLI Metrics

**Feature state:** Kubernetes v1.32 [stable] (feature gate `ComponentSLIs`, enabled by default)

By default, Kubernetes 1.37 publishes Service Level Indicator (SLI) metrics 
for each Kubernetes component binary. This metric endpoint is exposed on the serving 
HTTPS port of each component, at the path `/metrics/slis`. The 
`ComponentSLIs` [feature gate](/reference/command-line-tools-reference/feature-gates.md)
defaults to enabled for each Kubernetes component as of v1.27.

## SLI Metrics

With SLI metrics enabled, each Kubernetes component exposes two metrics,
labeled per healthcheck:

- a gauge (which represents the current state of the healthcheck)
- a counter (which records the cumulative counts observed for each healthcheck state)

You can use the metric information to calculate per-component availability statistics.
For example, the API server checks the health of etcd. You can work out and report how
available or unavailable etcd has been - as reported by its client, the API server.

The prometheus gauge data looks like this:

```
# HELP kubernetes_healthcheck [ALPHA] This metric records the result of a single healthcheck.
# TYPE kubernetes_healthcheck gauge
kubernetes_healthcheck{name="autoregister-completion",type="healthz"} 1
kubernetes_healthcheck{name="autoregister-completion",type="readyz"} 1
kubernetes_healthcheck{name="etcd",type="healthz"} 1
kubernetes_healthcheck{name="etcd",type="readyz"} 1
kubernetes_healthcheck{name="etcd-readiness",type="readyz"} 1
kubernetes_healthcheck{name="informer-sync",type="readyz"} 1
kubernetes_healthcheck{name="log",type="healthz"} 1
kubernetes_healthcheck{name="log",type="readyz"} 1
kubernetes_healthcheck{name="ping",type="healthz"} 1
kubernetes_healthcheck{name="ping",type="readyz"} 1
```

While the counter data looks like this:

```
# HELP kubernetes_healthchecks_total [ALPHA] This metric records the results of all healthcheck.
# TYPE kubernetes_healthchecks_total counter
kubernetes_healthchecks_total{name="autoregister-completion",status="error",type="readyz"} 1
kubernetes_healthchecks_total{name="autoregister-completion",status="success",type="healthz"} 15
kubernetes_healthchecks_total{name="autoregister-completion",status="success",type="readyz"} 14
kubernetes_healthchecks_total{name="etcd",status="success",type="healthz"} 15
kubernetes_healthchecks_total{name="etcd",status="success",type="readyz"} 15
kubernetes_healthchecks_total{name="etcd-readiness",status="success",type="readyz"} 15
kubernetes_healthchecks_total{name="informer-sync",status="error",type="readyz"} 1
kubernetes_healthchecks_total{name="informer-sync",status="success",type="readyz"} 14
kubernetes_healthchecks_total{name="log",status="success",type="healthz"} 15
kubernetes_healthchecks_total{name="log",status="success",type="readyz"} 15
kubernetes_healthchecks_total{name="ping",status="success",type="healthz"} 15
kubernetes_healthchecks_total{name="ping",status="success",type="readyz"} 15
```

## Using this data

The component SLIs metrics endpoint is intended to be scraped at a high frequency. Scraping
at a high frequency means that you end up with greater granularity of the gauge's signal, which
can be then used to calculate SLOs. The `/metrics/slis` endpoint provides the raw data necessary
to calculate an availability SLO for the respective Kubernetes component.
