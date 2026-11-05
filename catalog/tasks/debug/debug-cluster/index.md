# Troubleshooting Clusters

* [Auditing](audit.md) - Kubernetes _auditing_ provides a security-relevant, chronological set of records documenting the sequence of actions in a cluster.
* [Debugging Kubernetes Nodes With Kubectl](kubectl-node-debug.md) - This page shows how to debug a node running on the Kubernetes cluster using `kubectl debug` command.
* [Developing and debugging services locally using telepresence](local-debugging.md) - Kubernetes applications usually consist of multiple, separate services, each running in its own container.
* [Windows debugging tips](windows.md) - Ensure that your pause image is compatible with your Windows OS version.
* [Troubleshooting kubectl](troubleshoot-kubectl.md) - This documentation is about investigating and diagnosing kubectl related issues.
* [Resource metrics pipeline](resource-metrics-pipeline.md) - For Kubernetes, the _Metrics API_ offers a basic set of metrics to support automatic scaling and similar use cases.
* [Tools for Monitoring Resources](resource-usage-monitoring.md) - To scale an application and provide a reliable service, you need to understand how the application behaves when it is deployed.
* [Monitor Node Health](monitor-node-health.md) - You can run Node Problem Detector as a `DaemonSet` or as a standalone daemon.
* [Debugging Kubernetes nodes with crictl](crictl.md) - You can use it to inspect and debug container runtimes and applications on a Kubernetes node.
* [Troubleshooting Topology Management](topology.md) - Kubernetes keeps many aspects of how pods execute on nodes abstracted from the user.
