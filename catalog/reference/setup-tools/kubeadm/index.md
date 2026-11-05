# Kubeadm

* [Kubeadm Generated](generated.md) - See the contents of this section. ([contents](generated/index.md))
* [kubeadm init](kubeadm-init.md) - This command initializes a Kubernetes control plane node.
* [kubeadm join](kubeadm-join.md) - This command initializes a new Kubernetes node and joins it to the existing cluster.
* [kubeadm upgrade](kubeadm-upgrade.md) - behind one command, with support for both planning an upgrade and actually performing it.
* [kubeadm upgrade phases](kubeadm-upgrade-phase.md) - Using the phases of `kubeadm upgrade apply`, you can choose to execute the separate steps of the initial upgrade of a control plane node.
* [kubeadm config](kubeadm-config.md) - During `kubeadm init`, kubeadm uploads the `ClusterConfiguration` object to your cluster in a ConfigMap called `kubeadm-config` in the `kube-system` namespace.
* [kubeadm reset](kubeadm-reset.md) - Performs a best effort revert of changes made by `kubeadm init` or `kubeadm join`.
* [kubeadm token](kubeadm-token.md) - Bootstrap tokens are used for establishing bidirectional trust between a node joining the cluster and a control-plane node, as described in authenticating with bootstrap tokens.
* [kubeadm version](kubeadm-version.md) - This command prints the version of kubeadm.
* [kubeadm alpha](kubeadm-alpha.md) - from the community.
* [kubeadm certs](kubeadm-certs.md) - For more details on how these commands can be used, see Certificate Management with kubeadm.
* [kubeadm init phase](kubeadm-init-phase.md) - Hence, you can let kubeadm do some of the work and you can fill in the gaps if you wish to apply customization.
* [kubeadm join phase](kubeadm-join-phase.md) - Hence, you can let kubeadm do some of the work and you can fill in the gaps if you wish to apply customization.
* [kubeadm kubeconfig](kubeadm-kubeconfig.md) - For examples on how to use `kubeadm kubeconfig user` see Generating kubeconfig files for additional users.
* [kubeadm reset phase](kubeadm-reset-phase.md) - Hence, you can let kubeadm do some of the work and you can fill in the gaps if you wish to apply customization.
* [Implementation details](implementation-details.md) - bare Kubernetes cluster from scratch, that aligns with the best-practices.
