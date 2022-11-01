### 构建runtime镜像
```
docker build -t xxx:yy .
```

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