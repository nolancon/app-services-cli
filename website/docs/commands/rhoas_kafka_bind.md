---
id: rhoas_kafka_bind
title: rhoas kafka bind
---
## rhoas kafka bind

Bind currently selected kafa cluster credentials to your cluster

### Synopsis

Bind command will use current kubernetes or openshift context (namespace/project you have selected)
		Bind command will retrieve credentials for your kafka and mount them as secret into your project.
		Additionally we going to provide extra metadata for the service binding operator:

		https://github.com/redhat-developer/service-binding-operator

		If your cluster has service binding operator installed, you would be able to bind your application with credentials directly from the console etc.

		To use command your cluster needs to have rhoas-operator installed:

		<Link to operator here>
		
		

```
rhoas kafka bind [flags]
```

### Options

```
      --dry-run oc apply -f kafka.yml   Provide yaml file containing changes without applying them to the cluster. Developers can use oc apply -f kafka.yml to apply it manually
  -h, --help                            help for bind
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Manage your Kafka clusters

###### Auto generated by spf13/cobra on 26-Nov-2020