# Manage Kubernetes Objects

* [Declarative Management of Kubernetes Objects Using Configuration Files](declarative-config.md) - Kubernetes objects can be created, updated, and deleted by storing multiple object configuration files in a directory and using `kubectl apply` to recursively create and update those objects as needed.
* [Declarative Management of Kubernetes Objects Using Kustomize](kustomization.md) - to customize Kubernetes objects through a kustomization file.
* [Managing Kubernetes Objects Using Imperative Commands](imperative-command.md) - Kubernetes objects can quickly be created, updated, and deleted directly using imperative commands built into the `kubectl` command-line tool.
* [Imperative Management of Kubernetes Objects Using Configuration Files](imperative-config.md) - Kubernetes objects can be created, updated, and deleted by using the `kubectl` command-line tool along with an object configuration file written in YAML or JSON.
* [Update API Objects in Place Using kubectl patch](update-api-object-kubectl-patch.md) - Use kubectl patch to update Kubernetes API objects in place. Do a strategic merge patch or a JSON merge patch.
* [Migrate Kubernetes Objects Using Storage Version Migration](storage-version-migration.md) - Kubernetes relies on API data being actively re-written, to support some maintenance activities related to at rest storage.
