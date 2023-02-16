package k8s
// import (
// 	"k8s.io/client-go/tools/clientcmd"
//   "k8s.io/client-go/kubernetes"
// )
// 	// 使用本地 ~/.kube/config 创建配置
// 	kubeConfigPath := os.ExpandEnv("$HOME/.kube/config")
// 	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
// 	if err != nil {    
// 		log.Fatal(err)
// 	}// 使用上面的配置获取连接
// 	c, err := kubernetes.NewForConfig(config)
// 	if err != nil {    
// 		log.Fatal(err)
// 	}

// import (    
// 	"k8s.io/client-go/restmapper"    
// 	"k8s.io/client-go/dynamic")// 获取支持的资源类型列表
// 	resources, err := restmapper.GetAPIGroupResources(c.Discovery())
// 	if err != nil {    
// 		log.Fatal(err)
// 	}// 创建 'Discovery REST Mapper'，获取查询的资源的类型
// 	mapper:= restmapper.NewDiscoveryRESTMapper(resourcesAvailable)
// 	// 获取 'Dynamic REST Interface'，获取一个指定资源类型的 REST 接口
// 	dynamicREST, err := dynamic.NewForConfig(config)
// 	if err != nil {    
// 		log.Fatal(err)
// 	}
// 	finalYAML, err := templates.Read(myFilePath)
// 	if err != nil {log.Fatal(err)}

// 	objectsInYAML := bytes.Split(yamlBytes, []byte("---"))
// 	if len(objectsInYAML) == 0 {    
// 		return nil, nil
// 	}

// 	import(    "k8s.io/apimachinery/pkg/runtime/serializer/yaml")...
// 	for _, objectInYAML := range objectsInYAML {    
// 		runtimeObject, groupVersionAndKind, err :=     yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).        
// 		Decode(objectInYAML.Raw, nil, nil)    
// 		if err != nil {       log.Fatal(err)    }...