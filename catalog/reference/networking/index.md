# Networking Reference

* [Protocols for Services](service-protocols.md) - If you configure a Service, you can select from any network protocol that Kubernetes supports.
* [Ports and Protocols](ports-and-protocols.md) - When running Kubernetes in an environment with strict network boundaries, such as on-premises datacenter with physical network firewalls or Virtual Networks in Public Cloud, it is useful to be aware of the ports and protocols used by Kubernetes components.
* [Virtual IPs and Service Proxies](virtual-ips.md) - Every node in a Kubernetes cluster runs a kube-proxy (unless you have deployed your own alternative component in place of `kube-proxy`).
