## Common

	"/common/postping", handler.CommonPostTest)

​	





## Admin

```
/admin/postping

```

#### Pod

**/admin/getPod**

```
POST {"namespace"[string], "pod"[string]}
```

**/admin/listPods**

```
POST {}
```

**/admin/podsMetrics**

```
POST {}
```

**/admin/podMemory**

```
POST {"namespace"[string], "pod"[string]}
```

**/admin/podCpu**

```
POST {"namespace"[string], "pod"[string]}
```







## Root
#### Namespace

**/root/getNamespace**

```
POST {"namespace"[string]}
```

**/root/listNamespaces**

```
POST {}
```

**/root/createNamespace**

```
POST {"namespace"[string]}
```

**/root/deleteNamespace**

```
POST {"namespace"[string]}
```

#### Node

**/root/getNode**

```
POST {"node"[string]}
```

**/root/listNodes**

```
POST {}
```

**/root/nodesMetrics**

```
POST {}
```

