package main

import (
	"context"
	"log"

	v1 "github.com/mohuishou/blog-code/02-k8s-operator/operator-kubebuilder-clientset/api/job/v1"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	cfg, err := config.GetConfigWithContext("kind-kind")
	fatalf(err, "get config fail")

	scheme, err := v1.SchemeBuilder.Build()
	fatalf(err, "get scheme fail")

	c, err := client.New(cfg, client.Options{Scheme: scheme})
	fatalf(err, "new client fail")

	test := v1.Test{}
	err = c.Get(context.Background(), types.NamespacedName{
		Namespace: "default",
		Name:      "test-sample",
	}, &test)
	fatalf(err, "get resource fail")

	log.Printf("foo: %s", test.Spec.Foo)
}

func fatalf(err error, format string, args ...any) {
	if err == nil {
		return
	}
	err = errors.Wrapf(err, format, args...)
	log.Panicf("%+v", err)
}
