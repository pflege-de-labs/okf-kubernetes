---
type: task
title: Manual Rotation of CA Certificates
description: This page shows how to manually rotate the certificate authority (CA) certificates.
resource: https://kubernetes.io/docs/tasks/tls/manual-rotation-of-ca-certificates/
tags: [kubernetes, tasks, tls]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/tls/manual-rotation-of-ca-certificates.md
    title: 'Kubernetes website source: tasks/tls/manual-rotation-of-ca-certificates.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/tls/manual-rotation-of-ca-certificates/
    title: Manual Rotation of CA Certificates | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/tls/manual-rotation-of-ca-certificates.md
---

# Manual Rotation of CA Certificates

This page shows how to manually rotate the certificate authority (CA) certificates.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

- For more information about authentication in Kubernetes, see
  [Authenticating](/reference/access-authn-authz/authentication.md).
- For more information about best practices for CA certificates, see
  [Single root CA](https://kubernetes.io/docs/setup/best-practices/certificates/#single-root-ca).

## Rotate the CA certificates manually

**Caution:**
Make sure to back up your certificate directory along with configuration files and any other necessary files.

This approach assumes operation of the Kubernetes control plane in an HA configuration with multiple API servers.
Graceful termination of the API server is also assumed, so clients can cleanly disconnect from one API server and
reconnect to another.

Configurations with a single API server will experience unavailability while the API server is being restarted.

1. Distribute the new CA certificates and private keys (for example: `ca.crt`, `ca.key`, `front-proxy-ca.crt`,
   and `front-proxy-ca.key`) to all your control plane nodes in the Kubernetes certificates directory.

1. Update the `--root-ca-file` flag for the [kube-controller-manager](/reference/glossary/kube-controller-manager.md) to include
   both old and new CA, then restart the kube-controller-manager.

   Any [ServiceAccount](/reference/glossary/service-account.md) created after this point will get
   Secrets that include both old and new CAs.

   **Note:**
   The files specified by the kube-controller-manager flags `--client-ca-file` and `--cluster-signing-cert-file`
   cannot be CA bundles. If these flags and `--root-ca-file` point to the same `ca.crt` file, which is now a
   bundle (includes both old and new CA), you will face an error. To workaround this problem, you can copy the new CA
   to a separate file and make the flags `--client-ca-file` and `--cluster-signing-cert-file` point to the copy.
   Once `ca.crt` is no longer a bundle, you can restore the problem flags to point to `ca.crt` and delete the copy.

   [Issue 1350](https://github.com/kubernetes/kubeadm/issues/1350) for kubeadm tracks an bug with the
   kube-controller-manager being unable to accept a CA bundle.
   

1. Wait for the controller manager to update `ca.crt` in the service account Secrets to include both old and new CA certificates.

    If any Pods are started before the new CA is used by API servers, the new Pods get this update and will trust both
    old and new CAs.

1. Restart all pods using in-cluster configurations (for example: kube-proxy, CoreDNS, etc) so they can use the
   updated certificate authority data from Secrets that link to ServiceAccounts.

   * Make sure CoreDNS, kube-proxy, and other Pods using in-cluster configurations are working as expected.

1. Append both old and new CA to the file against `--client-ca-file` and `--kubelet-certificate-authority`
   flag in the `kube-apiserver` configuration.

1. Append both old and new CA to the file against the `--client-ca-file` flag in the `kube-scheduler` configuration.

1. Update certificates for user accounts by replacing the content of `client-certificate-data` and `client-key-data`
   respectively.

   For information about creating certificates for individual user accounts, see
   [Configure certificates for user accounts](https://kubernetes.io/docs/setup/best-practices/certificates/#configure-certificates-for-user-accounts).

   Additionally, update the `certificate-authority-data` section in the kubeconfig files,
   respectively with Base64-encoded old and new certificate authority data

1. Update the `--root-ca-file` flag for the [Cloud Controller Manager](/reference/glossary/cloud-controller-manager.md) to include
   both old and new CA, then restart the cloud-controller-manager.

   **Note:**
   If your cluster does not have a cloud-controller-manager, you can skip this step.
   

1. Follow the steps below in a rolling fashion.

   1. Restart any other
      [aggregated API servers](/concepts/extend-kubernetes/api-extension/apiserver-aggregation.md) or
      webhook handlers to trust the new CA certificates.

   1. Restart the kubelet by updating the file against `clientCAFile` in the kubelet configuration and
      `certificate-authority-data` in `kubelet.conf` to use both the old and new CA on all nodes.

      If your kubelet is not using client certificate rotation, update `client-certificate-data` and
      `client-key-data` in `kubelet.conf` on all nodes, along with the kubelet client certificate file
      usually found in `/var/lib/kubelet/pki`.

   1. Restart API servers with the certificates (`apiserver.crt`, `apiserver-kubelet-client.crt`, and
      `front-proxy-client.crt`), signed by the new CA.
      You can use the existing private keys or new private keys.
      If you changed the private keys, then update these in the Kubernetes certificates directory as well.

      Since the Pods in your cluster trust both old and new CAs, there will be a momentary disconnection
      after which, pods' Kubernetes clients reconnect to the new API server.
      The new API server uses a certificate signed by the new CA.

      * Restart the [kube-scheduler](/reference/glossary/kube-scheduler.md) to use and
        trust the new CAs.
      * Make sure control plane components logs no TLS errors.

      **Note:**
      To generate certificates and private keys for your cluster using the `openssl` command line tool,
      see [Certificates (`openssl`)](/tasks/administer-cluster/certificates.md#openssl).
      You can also use [`cfssl`](/tasks/administer-cluster/certificates.md#cfssl).
      

    1. Annotate any DaemonSets and Deployments to trigger pod replacement in a safer rolling fashion.

      ```shell
      for namespace in $(kubectl get namespace -o jsonpath='{.items[*].metadata.name}'); do
          for name in $(kubectl get deployments -n $namespace -o jsonpath='{.items[*].metadata.name}'); do
              kubectl patch deployment -n ${namespace} ${name} -p '{"spec":{"template":{"metadata":{"annotations":{"ca-rotation": "1"}}}}}';
          done
          for name in $(kubectl get daemonset -n $namespace -o jsonpath='{.items[*].metadata.name}'); do
              kubectl patch daemonset -n ${namespace} ${name} -p '{"spec":{"template":{"metadata":{"annotations":{"ca-rotation": "1"}}}}}';
          done
      done
      ```

      **Note:**
      To limit the number of concurrent disruptions that your application experiences,
      see [configure pod disruption budget](/tasks/run-application/configure-pdb.md).
      

        Depending on how you use StatefulSets, you may also need to perform a similar rolling replacement.

1. If your cluster is using bootstrap tokens to join nodes, update the ConfigMap `cluster-info` in the `kube-public`
   namespace with new CA.

   ```shell
   base64_encoded_ca="$(base64 -w0 /etc/kubernetes/pki/ca.crt)"

   kubectl get cm/cluster-info --namespace kube-public -o yaml | \
       /bin/sed "s/\(certificate-authority-data:\).*/\1 ${base64_encoded_ca}/" | \
       kubectl apply -f -
   ```

1. Verify the cluster functionality.

    1. Check the logs from control plane components, along with the kubelet and the kube-proxy.
       Ensure those components are not reporting any TLS errors; see
       [looking at the logs](/tasks/debug/debug-cluster.md#looking-at-logs) for more details.

    1. Validate logs from any aggregated api servers and pods using in-cluster config.

1. Once the cluster functionality is successfully verified:

   1. Update all service account tokens to include new CA certificate only.

      * All pods using an in-cluster kubeconfig will eventually need to be restarted to pick up the new Secret,
        so that no Pods are relying on the old cluster CA.

   1. Restart the control plane components by removing the old CA from the kubeconfig files and the files against it
      `--client-ca-file`, `--root-ca-file` flags resp.

   1. On each node, restart the kubelet by removing the old CA from the file against the `clientCAFile` flag
      and from the kubelet kubeconfig file. You should carry this out as a rolling update.

      If your cluster lets you make this change, you can also roll it out by replacing nodes rather than
      reconfiguring them.
