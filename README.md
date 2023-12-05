# goland

# 基本语法

# 设计模式

简单工厂模式
一般工厂模式
抽象工厂模式
装饰者模式
命令模式
单例模式
装饰器模式
观察者模式
代理模式
策略模式
模板模式
外观模式

# 推荐学习站点

# https://www.liwenzhou.com/posts/Go/golang-menu/
# https://github.com/coderit666/GoGuide
# https://gin-gonic.com/zh-cn/docs/examples/multipart-urlencoded-binding/
# https://www.cnblogs.com/liwenzhou/p/13629767.html
# https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/

# chatGPT-main.rar 人工智能源码
# https://cloud.sealos.io  云操作系统

# fanqiang

 1. https://7zy.com.cn/726.html 免费 
 2. https://wwwjs01.com/  收费 
 3. https://flm88.blog/s/sg0179 收费
 4. https://github.com/hwanz/SSR-V2ray-Trojan-vpn 免费
 5. https://github.com/yuchuanqicy/Over-The-Wall 免费
 6. https://github.com/jjqqkk/jjqqkk 免费+收费
 7. https://www.4spaces.org/1406.html 自建
 8. https://github.com/XX-net/XX-Net 收费
 9. https://iguge.xyz/ 收费

#python
1. Python 库让你相见恨晚 https://www.zhihu.com/question/24590883/answer/1220720307

# 大电商
https://www.youzan.com



```


3. template文件编写
```
language: net6evoc
fprocess: dotnet ./root.dll
welcome_message: |
```

4. 文件上传到template文件夹
![img_29.png](img_29.png)

5. 创建一个给予模板net6的自定义函数
```
faas-cli new net6test05 --lang netroot -p jxsdpengbin --gateway 172.16.12.146:31112
```
![img_30.png](img_30.png)

6. 构建
``` 
faas-cli build -f ./net6test05.yml
```
![img_31.png](img_31.png)

7. 推送镜像
```
docker push jxsdpengbin/net6test05:latest
```
![img_32.png](img_32.png)

8. deploy
```
faas-cli deploy -f ./net6test05.yml
```
![img_33.png](img_33.png)

9. 进入pod内部执行 https认证指令，
``` 
kubectl exec -it net6test05-59c44944db-z22dg -c net6test05 -n openfaas-fn bash
dotnet dev-certs https
``` 
![img_36.png](img_36.png)

10. 用nodeport做pod内的端口暴露到容器所在的ip和端口。例如：vim 111.yaml，注意空格对其

``` 
apiVersion: v1
kind: Service
metadata:
    annotations:
        prometheus.io.scrape: "false"
    name: net6test02
    namespace: openfaas-fn
spec:
    type: NodePort
    ports:
        - port: 5128
          name: dslgjd
          targetPort: 5128
          nodePort: 30007
        - port: 7128
          name: dlsjgl
          targetPort: 7128
          nodePort: 30017
``` 
应用转换指令
``` 
kubectl  apply -f 111.yaml
``` 

访问连接：
``` 
https://172.16.12.146:30017/Add?strNumbers=1,25
``` 
![img_35.png](img_35.png)

## 常用问题排查指令

1. 查询所有pod状态
``` 
kubectl get pods -n openfaas-fn
```
2. 查询pod启动详情
``` 
kubectl describe pod net6test02-58bc5494c8-gtr4n -n openfaas-fn
``` 
3. 查询pod启动日志
``` 
kubectl logs -f net6test05-59c44944db-z22dg -n openfaas-fn
``` 
4. 调用 指定ip端口
``` 
curl -v 172.20.3.162:8080
``` 

5. 进入pod内部执行 某些指令，例如:dotnet dev-certs https
``` 
kubectl exec -it net6test05-59c44944db-z22dg -c net6test05 -n openfaas-fn bash
``` 
6. 编辑某个容器的 deployment 文件。例如读写探针，端口等等
``` 
kubectl edit deployment net6test05 -n openfaas-fn
``` 
7. 查询svc暴露的端口
``` 
kubectl  get svc -n openfaas-fn
``` 

8. 查询某个容器的 id，名称，等信息
``` 
docker ps|grep net6test05
``` 

9. 根据容器的id 查询某个容器的State.Pid
``` 
docker inspect -f '{{.State.Pid}}'  d6010624a714
``` 

10. 根据容器的pid 进入pod 内容，查看pod内启动的程序，查询完成记得执行 exit退出pod
``` 
nsenter -t 11691 -n
netstat -tnlp
``` 
![img_34.png](img_34.png)



