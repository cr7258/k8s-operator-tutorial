package main

import (
	"context"
	"flag"
	"k8s-operator-tutorial/pkg/runtime"
	"k8s-operator-tutorial/pkg/subscription"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"math/rand"
	"time"
)

var (
	minWatchTimeout = 5 * time.Minute
	timeoutSeconds  = int64(minWatchTimeout.Seconds() * (rand.Float64() + 1.0))
	masterURL       string
	kubeconfig      string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	klog.Info("Got watcher client...")

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	klog.Info("Built config from flags...")
	defaultKubernetesClientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	context := context.TODO()

	podSubscription := &subscription.PodSubscription{
		ClientSet:  defaultKubernetesClientSet,
		Ctx:        context,
		Completion: make(chan bool),
	}

	if err := runtime.RunLoop([]subscription.ISubscription{
		podSubscription,
	}); err != nil {
		klog.Fatalf("Error running loop: %s", err.Error())
	}
	select {}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")

}
