### demo目标

通过demo展示如何利用k8s的schme进行对象版本转换，目标：

（1）深入理解scheme机制

（2）遇到项目组需要版本转换，可以直接封装使用

### Demo 结构

直接go run convertHpa.go/testconvert.go 即可。

 k8s.sh是为了下载k8s.io/kubernetes依赖包

```
.
├── convertHpa.go              // 展示如何将v1beta2版本hpa转换为v1版本
├── go.mod
├── go.sum
├── import_known_vesions.go
├── k8s.sh
└── testconvert.go            // 展示如何将v1beta1版本deploy转换为v1版本
```

这里有个疑问，为什么我 import_known_vesions.go 不起作用，scheme里面是空的，奇怪。。

后面我是在convertHpa.go里面自己install各个版本的。