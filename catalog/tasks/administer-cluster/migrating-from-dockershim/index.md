# Migrating from dockershim

* [Changing the Container Runtime on a Node from Docker Engine to containerd](change-runtime-containerd.md) - This task outlines the steps needed to update your container runtime to containerd from Docker.
* [Find Out What Container Runtime is Used on a Node](find-out-runtime-you-use.md) - This page outlines steps to find out what container runtime the nodes in your cluster use.
* [Troubleshooting CNI plugin-related errors](troubleshooting-cni-plugin-related-errors.md) - To avoid CNI plugin-related errors, verify that you are using or upgrading to a container runtime that has been tested to work correctly with your version of Kubernetes.
* [Check whether dockershim removal affects you](check-if-dockershim-removal-affects-you.md) - The `dockershim` component of Kubernetes allows the use of Docker as a Kubernetes's container runtime.
* [Migrating telemetry and security agents from dockershim](migrating-telemetry-and-security-agents.md) - Kubernetes' support for direct integration with Docker Engine is deprecated and has been removed.
