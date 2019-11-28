# 部署

本项目已经打包成[docker镜像](https://hub.docker.com/r/longfangsong/shuhelperConfigStorage/)。

## 支持服务
### postgresql数据库
migration文件位于本repo的 [migration](https://github.com/shuosc/shuhelperConfigStorage/tree/master/migration) 目录中。

建议使用 [golang-migrate](https://github.com/golang-migrate/migrate) 来进行 migrate。
```shell
migrate -source github://[你的Github用户名]:[你的Github Access Token]@shuosc/shuhelperConfigStorage/migration -database [你的postgrsql数据库url] up
```

## 服务本身
### 环境变量
- `PORT`: 服务端口
- `DB_ADDRESS`: 数据库url
- `JWT_SECRET`: jwt密钥

### k8s
在k8s下使用如下yaml，假设
- `JWT_SECRET`由k8s secret给出
- 数据库服务器在`shuhelper-config-storage-postgres-svc`
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: shuhelper-config-storage
spec:
  selector:
    matchLabels:
      run: shuhelper-config-storage
  replicas: 1
  template:
    metadata:
      labels:
        run: shuhelper-config-storage
    spec:
      containers:
      - name: shuhelper-config-storage
        image: shuosc/shuhelper-config-storage
        env:
        - name: PORT
          value: "8000"
        - name: DB_ADDRESS
          value: "postgres://shuosc@shuhelper-config-storage-postgres-svc:5432/shuhelper-config-storage?sslmode=disable"
        - name: JWT_SECRET
          valueFrom:
            secretKeyRef:
              name: shuosc-secret
              key: JWT_SECRET
        ports:
        - containerPort: 8000
---
apiVersion: v1
kind: Service
metadata:
  name: shuhelper-config-storage-svc
spec:
  selector:
     run: shuhelper-config-storage
  ports:
  - protocol: TCP
    port: 8000
    targetPort: 8000
```