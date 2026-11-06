---
type: concept
title: Automatic Cleanup for Finished Jobs
description: A time-to-live mechanism to clean up old Jobs that have finished execution.
resource: https://kubernetes.io/docs/concepts/workloads/controllers/ttlafterfinished/
tags: [kubernetes, concepts, workloads]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/controllers/ttlafterfinished.md
    title: 'Kubernetes website source: concepts/workloads/controllers/ttlafterfinished.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/workloads/controllers/ttlafterfinished/
    title: Automatic Cleanup for Finished Jobs | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/workloads/controllers/ttlafterfinished.md
---

# Automatic Cleanup for Finished Jobs

**Feature state:** Kubernetes v1.23 [stable]

When your Job has finished, it's useful to keep that Job in the API (and not immediately delete the Job)
so that you can tell whether the Job succeeded or failed.

Kubernetes' TTL-after-finished [controller](/reference/glossary/controller.md) provides a
TTL (time to live) mechanism to limit the lifetime of Job objects that
have finished execution.

## Cleanup for finished Jobs

The TTL-after-finished controller is only supported for Jobs. You can use this mechanism to clean
up finished Jobs (either `Complete` or `Failed`) automatically by specifying the
`.spec.ttlSecondsAfterFinished` field of a Job, as in this
[example](/concepts/workloads/controllers/job.md#clean-up-finished-jobs-automatically).

The TTL-after-finished controller assumes that a Job is eligible to be cleaned up
TTL seconds after the Job has finished. The timer starts once the
status condition of the Job changes to show that the Job is either `Complete` or `Failed`; once the TTL has
expired, that Job becomes eligible for
[cascading](/concepts/architecture/garbage-collection.md#cascading-deletion) removal. When the
TTL-after-finished controller cleans up a job, it will delete it cascadingly, that is to say it will delete
its dependent objects together with it.

Kubernetes honors object lifecycle guarantees on the Job, such as waiting for
[finalizers](/concepts/overview/working-with-objects/finalizers.md).

You can set the TTL seconds at any time. Here are some examples for setting the
`.spec.ttlSecondsAfterFinished` field of a Job:

* Specify this field in the Job manifest, so that a Job can be cleaned up
  automatically some time after it finishes.
* Manually set this field of existing, already finished Jobs, so that they become eligible
  for cleanup.
* Use a
  [mutating admission webhook](/reference/access-authn-authz/admission-controllers.md#mutatingadmissionwebhook)
  to set this field dynamically at Job creation time. Cluster administrators can
  use this to enforce a TTL policy for finished jobs.
* Use a
  [mutating admission webhook](/reference/access-authn-authz/admission-controllers.md#mutatingadmissionwebhook)
  to set this field dynamically after the Job has finished, and choose
  different TTL values based on job status, labels. For this case, the webhook needs
  to detect changes to the `.status` of the Job and only set a TTL when the Job
  is being marked as completed.
* Write your own controller to manage the cleanup TTL for Jobs that match a particular
  [selector](/reference/glossary/selector.md).

## Caveats

### Updating TTL for finished Jobs

You can modify the TTL period, e.g. `.spec.ttlSecondsAfterFinished` field of Jobs,
after the job is created or has finished. If you extend the TTL period after the
existing `ttlSecondsAfterFinished` period has expired, Kubernetes doesn't guarantee
to retain that Job, even if an update to extend the TTL returns a successful API
response.

### Time skew

Because the TTL-after-finished controller uses timestamps stored in the Kubernetes jobs to
determine whether the TTL has expired or not, this feature is sensitive to time
skew in your cluster, which may cause the control plane to clean up Job objects
at the wrong time.

Clocks aren't always correct, but the difference should be
very small. Please be aware of this risk when setting a non-zero TTL.

## What's next

* Read [Clean up Jobs automatically](/concepts/workloads/controllers/job.md#clean-up-finished-jobs-automatically)

* Refer to the [Kubernetes Enhancement Proposal](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/592-ttl-after-finish/README.md)
  (KEP) for adding this mechanism.
