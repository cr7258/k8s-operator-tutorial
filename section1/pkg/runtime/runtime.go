package runtime

import (
	"k8s-operator-tutorial/pkg/subscription"
)

func RunLoop(subscriptions []subscription.ISubscription) error {
	//var wg sync.WaitGroup
	for _, subscription := range subscriptions {
		//wg.Add(1)
		wiface, err := subscription.Subscribe()
		if err != nil {
			return err
		}
		go func() {
			for {
				select {
				case msg := <-wiface.ResultChan():
					subscription.Reconcile(msg.Object, msg.Type)
					//case isComplete := <-subscription.IsCompleted():
					//	if isComplete {
					//		wg.Done()
					//	}
				}
			}
		}()
	}
	//wg.Wait()
	return nil
}
