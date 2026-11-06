# kubeadm_init

* [kubeadm_init_phase](kubeadm_init_phase.md) - Use this command to invoke single phase of the "init" workflow
* [kubeadm_init_phase_addon](kubeadm_init_phase_addon.md) - Install required addons for passing conformance tests
* [kubeadm_init_phase_addon_all](kubeadm_init_phase_addon_all.md) - Install all the addons
* [kubeadm_init_phase_addon_coredns](kubeadm_init_phase_addon_coredns.md) - Install the CoreDNS addon to a Kubernetes cluster
* [kubeadm_init_phase_addon_kube-proxy](kubeadm_init_phase_addon_kube-proxy.md) - Install the kube-proxy addon to a Kubernetes cluster
* [kubeadm_init_phase_bootstrap-token](kubeadm_init_phase_bootstrap-token.md) - Generates bootstrap tokens used to join a node to a cluster
* [kubeadm_init_phase_certs](kubeadm_init_phase_certs.md) - Certificate generation
* [kubeadm_init_phase_certs_all](kubeadm_init_phase_certs_all.md) - Generate all certificates
* [kubeadm_init_phase_certs_apiserver](kubeadm_init_phase_certs_apiserver.md) - Generate the certificate for serving the Kubernetes API, and save them into apiserver.crt and apiserver.key files.
* [kubeadm_init_phase_certs_apiserver-etcd-client](kubeadm_init_phase_certs_apiserver-etcd-client.md) - Generate the certificate the apiserver uses to access etcd, and save them into apiserver-etcd-client.crt and apiserver-etcd-client.key files.
* [kubeadm_init_phase_certs_apiserver-kubelet-client](kubeadm_init_phase_certs_apiserver-kubelet-client.md) - Generate the certificate for the API server to connect to kubelet, and save them into apiserver-kubelet-client.crt and apiserver-kubelet-client.key files.
* [kubeadm_init_phase_certs_ca](kubeadm_init_phase_certs_ca.md) - Generate the self-signed Kubernetes CA to provision identities for other Kubernetes components, and save them into ca.crt and ca.key files.
* [kubeadm_init_phase_certs_etcd-ca](kubeadm_init_phase_certs_etcd-ca.md) - Generate the self-signed CA to provision identities for etcd, and save them into etcd/ca.crt and etcd/ca.key files.
* [kubeadm_init_phase_certs_etcd-healthcheck-client](kubeadm_init_phase_certs_etcd-healthcheck-client.md) - Generate the certificate for liveness probes to healthcheck etcd, and save them into etcd/healthcheck-client.crt and etcd/healthcheck-client.key files.
* [kubeadm_init_phase_certs_etcd-peer](kubeadm_init_phase_certs_etcd-peer.md) - Generate the certificate for etcd nodes to communicate with each other, and save them into etcd/peer.crt and etcd/peer.key files.
* [kubeadm_init_phase_certs_etcd-server](kubeadm_init_phase_certs_etcd-server.md) - Generate the certificate for serving etcd, and save them into etcd/server.crt and etcd/server.key files.
* [kubeadm_init_phase_certs_front-proxy-ca](kubeadm_init_phase_certs_front-proxy-ca.md) - Generate the self-signed CA to provision identities for front proxy, and save them into front-proxy-ca.crt and front-proxy-ca.key files.
* [kubeadm_init_phase_certs_front-proxy-client](kubeadm_init_phase_certs_front-proxy-client.md) - Generate the certificate for the front proxy client, and save them into front-proxy-client.crt and front-proxy-client.key files.
* [kubeadm_init_phase_certs_sa](kubeadm_init_phase_certs_sa.md) - Generate a private key for signing service account tokens along with its public key
* [kubeadm_init_phase_control-plane](kubeadm_init_phase_control-plane.md) - Generate all static Pod manifest files necessary to establish the control plane
* [kubeadm_init_phase_control-plane_all](kubeadm_init_phase_control-plane_all.md) - Generate all static Pod manifest files
* [kubeadm_init_phase_control-plane_apiserver](kubeadm_init_phase_control-plane_apiserver.md) - Generates the kube-apiserver static Pod manifest
* [kubeadm_init_phase_control-plane_controller-manager](kubeadm_init_phase_control-plane_controller-manager.md) - Generates the kube-controller-manager static Pod manifest
* [kubeadm_init_phase_control-plane_scheduler](kubeadm_init_phase_control-plane_scheduler.md) - Generates the kube-scheduler static Pod manifest
* [kubeadm_init_phase_etcd](kubeadm_init_phase_etcd.md) - Generate static Pod manifest file for local etcd
* [kubeadm_init_phase_etcd_local](kubeadm_init_phase_etcd_local.md) - Generate the static Pod manifest file for a local, single-node local etcd instance
* [kubeadm_init_phase_kubeconfig](kubeadm_init_phase_kubeconfig.md) - Generate all kubeconfig files necessary to establish the control plane and the admin kubeconfig file
* [kubeadm_init_phase_kubeconfig_admin](kubeadm_init_phase_kubeconfig_admin.md) - Generate a kubeconfig file for the admin to use and for kubeadm itself
* [kubeadm_init_phase_kubeconfig_all](kubeadm_init_phase_kubeconfig_all.md) - Generate all kubeconfig files
* [kubeadm_init_phase_kubeconfig_controller-manager](kubeadm_init_phase_kubeconfig_controller-manager.md) - Generate a kubeconfig file for the controller manager to use
* [kubeadm_init_phase_kubeconfig_kubelet](kubeadm_init_phase_kubeconfig_kubelet.md) - Generate a kubeconfig file for the kubelet to use *only* for cluster bootstrapping purposes
* [kubeadm_init_phase_kubeconfig_scheduler](kubeadm_init_phase_kubeconfig_scheduler.md) - Generate a kubeconfig file for the scheduler to use
* [kubeadm_init_phase_kubeconfig_super-admin](kubeadm_init_phase_kubeconfig_super-admin.md) - Generate a kubeconfig file for the super-admin, and save it to super-admin.conf file.
* [kubeadm_init_phase_kubelet-finalize](kubeadm_init_phase_kubelet-finalize.md) - Updates settings relevant to the kubelet after TLS bootstrap
* [kubeadm_init_phase_kubelet-finalize_all](kubeadm_init_phase_kubelet-finalize_all.md) - Run all kubelet-finalize phases
* [kubeadm_init_phase_kubelet-finalize_enable-client-cert-rotation](kubeadm_init_phase_kubelet-finalize_enable-client-cert-rotation.md) - Enable kubelet client certificate rotation
* [kubeadm_init_phase_kubelet-start](kubeadm_init_phase_kubelet-start.md) - Write kubelet settings and (re)start the kubelet
* [kubeadm_init_phase_mark-control-plane](kubeadm_init_phase_mark-control-plane.md) - Mark a node as a control-plane
* [kubeadm_init_phase_preflight](kubeadm_init_phase_preflight.md) - Run pre-flight checks for kubeadm init.
* [kubeadm_init_phase_show-join-command](kubeadm_init_phase_show-join-command.md) - Show the join command for control-plane and worker node
* [kubeadm_init_phase_upload-certs](kubeadm_init_phase_upload-certs.md) - Upload certificates to kubeadm-certs
* [kubeadm_init_phase_upload-config](kubeadm_init_phase_upload-config.md) - Upload the kubeadm and kubelet configuration to a ConfigMap
* [kubeadm_init_phase_upload-config_all](kubeadm_init_phase_upload-config_all.md) - Upload all configuration to a config map
* [kubeadm_init_phase_upload-config_kubeadm](kubeadm_init_phase_upload-config_kubeadm.md) - Upload the kubeadm ClusterConfiguration to a ConfigMap called kubeadm-config in the kube-system namespace.
* [kubeadm_init_phase_upload-config_kubelet](kubeadm_init_phase_upload-config_kubelet.md) - Upload the kubelet component config to a ConfigMap
* [kubeadm_init_phase_wait-control-plane](kubeadm_init_phase_wait-control-plane.md) - Wait for the control plane to start
