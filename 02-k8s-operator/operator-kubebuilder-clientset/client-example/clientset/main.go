package main

import (
	"context"
	"fmt"
	"log"

	clientv1 "github.com/mohuishou/blog-code/02-k8s-operator/operator-kubebuilder-clientset/pkg/clientset/v1/typed/job/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	cfg, err := config.GetConfigWithContext("kind-kind")
	fatalf(err, "get config fail")

	client := clientv1.NewForConfigOrDie(cfg)

	test, err := client.Tests("default").Get(context.Background(), "test-sample", v1.GetOptions{})
	fatalf(err, "new client fail")

	log.Printf("foo: %s", test.Spec.Foo)
}

func fatalf(err error, format string, args ...any) {
	if err == nil {
		return
	}
	panic(fmt.Sprintf(format, args...))
}
