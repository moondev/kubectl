# kubectl
go run kubectl

```
go run github.com/moondev/kubectl@latest
go run github.com/moondev/kubectl@release-1.26
go run github.com/moondev/kubectl@v1.26.2



go run -v -ldflags="-X k8s.io/client-go/pkg/version.gitCommit=eb7bb41e -X k8s.io/client-go/pkg/version.gitTreeState=clean -X k8s.io/client-go/pkg/version.gitVersion=v1.26.2 -X k8s.io/client-go/pkg/version.gitMajor=1 -X k8s.io/client-go/pkg/version.gitMinor=26 -X k8s.io/client-go/pkg/version.buildDate=2022-07-12T15:56:10Z -X k8s.io/component-base/version.gitCommit=eb7bb41e -X k8s.io/component-base/version.gitTreeState=clean -X k8s.io/component-base/version.gitVersion=v1.26.2 -X k8s.io/component-base/version.gitMajor=1 -X k8s.io/component-base/version.gitMinor=26 -X k8s.io/component-base/version.buildDate=2022-07-12T15:56:10Z" github.com/moondev/kubectl@v1.26.2 version --client -o yaml
```