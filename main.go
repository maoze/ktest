package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var cfg *rest.Config
	var err error

	cfg, err = clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		panic(err)
	}
	c, err := clientset.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	watcher, err := c.CoreV1().Pods("").Watch(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for {
		select {
		case event := <-watcher.ResultChan():
			fmt.Println(event)
		}
	}
}
