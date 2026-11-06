---
type: task
title: Managing Secrets using Kustomize
description: Creating Secret objects using kustomization.yaml file.
resource: https://kubernetes.io/docs/tasks/configmap-secret/managing-secret-using-kustomize/
tags: [kubernetes, tasks, configmap-secret]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/configmap-secret/managing-secret-using-kustomize.md
    title: 'Kubernetes website source: tasks/configmap-secret/managing-secret-using-kustomize.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configmap-secret/managing-secret-using-kustomize/
    title: Managing Secrets using Kustomize | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/configmap-secret/managing-secret-using-kustomize.md
---

# Managing Secrets using Kustomize

`kubectl` supports using the [Kustomize object management tool](/tasks/manage-kubernetes-objects/kustomization.md) to manage Secrets
and ConfigMaps. You create a *resource generator* using Kustomize, which
generates a Secret that you can apply to the API server using `kubectl`.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

## Create a Secret

You can generate a Secret by defining a `secretGenerator` in a
`kustomization.yaml` file that references other existing files, `.env` files, or
literal values. For example, the following instructions create a kustomization
file for the username `admin` and the password `1f2d1e2e67df`.

**Note:**
The `stringData` field for a Secret does not work well with server-side apply.

### Create the kustomization file

**Literals:**

```yaml
secretGenerator:
- name: database-creds
  literals:
  - username=admin
  - password=1f2d1e2e67df
```
**Files:**

1.  Store the credentials in files. The filenames are the keys of the secret:

    ```shell
    echo -n 'admin' > ./username.txt
    echo -n '1f2d1e2e67df' > ./password.txt
    ```
    The `-n` flag ensures that there's no newline character at the end of your
    files.

1.  Create the `kustomization.yaml` file:

    ```yaml
    secretGenerator:
    - name: database-creds
      files:
      - username.txt
      - password.txt
    ```

**.env files:**

You can also define the secretGenerator in the `kustomization.yaml` file by
providing `.env` files. For example, the following `kustomization.yaml` file
pulls in data from an `.env.secret` file:

```yaml
secretGenerator:
- name: db-user-pass
  envs:
  - .env.secret
```

In all cases, you don't need to encode the values in base64. The name of the YAML
file **must** be `kustomization.yaml` or `kustomization.yml`.

### Apply the kustomization file

To create the Secret, apply the directory that contains the kustomization file:

```shell
kubectl apply -k <directory-path>
```

The output is similar to:

```
secret/database-creds-5hdh7hhgfk created
```

When a Secret is generated, the Secret name is created by hashing
the Secret data and appending the hash value to the name. This ensures that
a new Secret is generated each time the data is modified.

To verify that the Secret was created and to decode the Secret data,

```shell
kubectl get -k <directory-path> -o jsonpath='{.data}' 
```

The output is similar to:

```
{ "password": "MWYyZDFlMmU2N2Rm", "username": "YWRtaW4=" }
```

```
echo 'MWYyZDFlMmU2N2Rm' | base64 --decode
```

The output is similar to:

```
1f2d1e2e67df
```

For more information, refer to
[Managing Secrets using kubectl](/tasks/configmap-secret/managing-secret-using-kubectl.md#verify-the-secret) and
[Declarative Management of Kubernetes Objects Using Kustomize](/tasks/manage-kubernetes-objects/kustomization.md).

## Edit a Secret {#edit-secret}

1.  In your `kustomization.yaml` file, modify the data, such as the `password`.
1.  Apply the directory that contains the kustomization file:

    ```shell
    kubectl apply -k <directory-path>
    ```

    The output is similar to:

    ```
    secret/db-user-pass-6f24b56cc8 created
    ```

The edited Secret is created as a new `Secret` object, instead of updating the
existing `Secret` object. You might need to update references to the Secret in
your Pods.

## Clean up

To delete a Secret, use `kubectl`:

```shell
kubectl delete secret db-user-pass
```

## What's next

- Read more about the [Secret concept](/concepts/configuration/secret.md)
- Learn how to [manage Secrets using kubectl](/tasks/configmap-secret/managing-secret-using-kubectl.md)
- Learn how to [manage Secrets using config file](/tasks/configmap-secret/managing-secret-using-config-file.md)
