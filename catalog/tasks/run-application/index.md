# Run Applications

* [Run a Stateless Application Using a Deployment](run-stateless-application-deployment.md) - This page shows how to run an application using a Kubernetes Deployment object.
* [Horizontal Manual Scaling for a Deployment](scale-deployment.md) - This page shows how to manually scale a Deployment horizontally, by changing its replica count.
* [Update a Deployment Without Downtime](update-deployment-rolling.md) - This page shows how to update a running Deployment to a new version using a rolling update.
* [Run a Single-Instance Stateful Application](run-single-instance-stateful-application.md) - This page shows you how to run a single-instance stateful application in Kubernetes using a PersistentVolume and a Deployment.
* [Run a Replicated Stateful Application](run-replicated-stateful-application.md) - This page shows how to run a replicated stateful application using a StatefulSet.
* [Scale a StatefulSet](scale-stateful-set.md) - This task shows how to scale a StatefulSet.
* [Delete a StatefulSet](delete-stateful-set.md) - This task shows you how to delete a StatefulSet.
* [Force Delete StatefulSet Pods](force-delete-stateful-set-pod.md) - This page shows how to delete Pods which are part of a stateful set, and explains the considerations to keep in mind when doing so.
* [HorizontalPodAutoscaler Walkthrough](horizontal-pod-autoscale-walkthrough.md) - A HorizontalPodAutoscaler (HPA for short) automatically updates a workload resource (such as a Deployment or StatefulSet), with the aim of automatically scaling the workload to match demand.
* [Specifying a Disruption Budget for your Application](configure-pdb.md) - This page shows how to limit the number of concurrent disruptions that your application experiences, allowing for higher availability while permitting the cluster administrator to manage the clusters nodes.
* [Accessing the Kubernetes API from a Pod](access-api-from-pod.md) - This guide demonstrates how to access the Kubernetes API from within a pod.
