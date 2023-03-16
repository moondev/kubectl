# kubectl
go run kubectl

```
go run github.com/moondev/kubectl@latest
go run github.com/moondev/kubectl@release-1.26
go run github.com/moondev/kubectl@v1.26.2

go run -ldflags="-X k8s.io/client-go/pkg/version.gitVersion=v1.26.2 -X k8s.io/component-base/version.gitVersion=v1.26.2 -X k8s.io/component-base/version.gitMajor=1 -X k8s.io/component-base/version.gitMinor=26 " github.com/moondev/kubectl@v1.26.2 version



```


