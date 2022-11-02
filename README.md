### 构建runtime镜像
```
docker build -t xxx:yy .
```

**将K8S的配置文件放入src/conf中，命名为k8sconfig**

### 使用容器运行代码
```
docker run -it -v ${pwd}/src:/app xxx:yy
```

***容器内部***
```
/bin/bash
cd /app
go get
make build_run
```