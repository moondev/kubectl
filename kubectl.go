package main
//go:generate go run -ldflags="'-X k8s.io/client-go/pkg/version.gitVersion=v1.26.2' '-X k8s.io/component-base/version.gitVersion=v1.26.2' '-X k8s.io/component-base/version.gitMajor=1' '-X k8s.io/component-base/version.gitMinor=23'" github.com/moondev/kubectl@v1.26.2

	"k8s.io/component-base/cli"
	"k8s.io/kubectl/pkg/cmd"
	"k8s.io/kubectl/pkg/cmd/util"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	command := cmd.NewDefaultKubectlCommand()
	if err := cli.RunNoErrOutput(command); err != nil {
		// Pretty-print the error and exit with an error.
		util.CheckErr(err)
	}
}
