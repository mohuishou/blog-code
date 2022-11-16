package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	jobv1 "github.com/mohuishou/blog-code/02-k8s-operator/operator-kubebuilder-clientset/api/job/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	cfg, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	fatalf(err, "get kube config fail")

	// 获取 client
	gvr := schema.GroupVersionResource{
		Group:    jobv1.GroupVersion.Group,
		Version:  jobv1.GroupVersion.Version,
		Resource: "tests",
	}
	client := dynamic.NewForConfigOrDie(cfg).Resource(gvr)

	ctx := context.Background()
	res, err := client.Namespace("default").Get(ctx, "test-sample", v1.GetOptions{})
	fatalf(err, "get resource fail")

	b, err := res.MarshalJSON()
	fatalf(err, "get json byte fail")

	test := jobv1.Test{}
	err = json.Unmarshal(b, &test)
	fatalf(err, "get json byte fail")

	log.Printf("foo: %s", test.Spec.Foo)
}

func fatalf(err error, format string, args ...any) {
	if err == nil {
		return
	}
	panic(fmt.Sprintf(format, args...))
}
