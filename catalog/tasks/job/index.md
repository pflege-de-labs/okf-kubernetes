# Run Jobs

* [Running Automated Tasks with a CronJob](automated-tasks-with-cron-jobs.md) - This page shows how to run automated tasks using Kubernetes CronJob object.
* [Coarse Parallel Processing Using a Work Queue](coarse-parallel-processing-work-queue.md) - In this example, you will run a Kubernetes Job with multiple parallel worker processes.
* [Fine Parallel Processing Using a Work Queue](fine-parallel-processing-work-queue.md) - In this example, you will run a Kubernetes Job that runs multiple parallel tasks as worker processes, each running as a separate Pod.
* [Indexed Job for Parallel Processing with Static Work Assignment](indexed-parallel-processing-static.md) - In this example, you will run a Kubernetes Job that uses multiple parallel worker processes.
* [Job with Pod-to-Pod Communication](job-with-pod-to-pod-communication.md) - In this example, you will run a Job in Indexed completion mode configured such that the pods created by the Job can communicate with each other using pod hostnames rather than pod IP addresses.
* [Parallel Processing using Expansions](parallel-processing-expansion.md) - This task demonstrates running multiple Jobs based on a common template.
* [Handling retriable and non-retriable pod failures with Pod failure policy](pod-failure-policy.md) - This document shows you how to use the Pod failure policy, in combination with the default Pod backoff failure policy, to improve the control over the handling of container- or Pod-level failure within a Job.
