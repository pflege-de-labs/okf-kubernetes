# Rbac

* [ClusterRole](cluster-role-v1.md) - ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.
* [ClusterRoleBinding](cluster-role-binding-v1.md) - ClusterRoleBinding references a ClusterRole, but not contain it. It can reference a ClusterRole in the global namespace, and adds who information via Subject.
* [Role](role-v1.md) - Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.
* [RoleBinding](role-binding-v1.md) - RoleBinding references a role, but does not contain it. It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in. RoleBindings in a given namespace only have effect in that namespace.
