package subscription

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

type PodSubscription struct {
	watcherInterface watch.Interface
	ClientSet        kubernetes.Interface
	Ctx              context.Context
	Completion       chan bool
}

func (p *PodSubscription) Reconcile(object runtime.Object, event watch.EventType) {
	pod := object.(*v1.Pod)
	klog.Infof("PodSubscription event type %s for %s", event, pod.Name)

	switch event {
	case watch.Added:
		klog.Infof("Pod %s added", pod.Name)
	case watch.Modified:
		klog.Infof("Pod %s modified", pod.Name)
	case watch.Deleted:
		klog.Infof("Pod %s deleted", pod.Name)
	}
}

func (p *PodSubscription) Subscribe() (watch.Interface, error) {
	var err error
	p.watcherInterface, err = p.ClientSet.CoreV1().Pods("").Watch(p.Ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	klog.Info("Started watch stream for PodSubscription")
	return p.watcherInterface, nil
}

func (p *PodSubscription) IsCompleted() <-chan bool {
	return p.Completion
}
