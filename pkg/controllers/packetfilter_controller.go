package controllers

import (
	"fmt"
	edwinv1 "github.com/YRXING/edwin/pkg/apis/edwin/v1"
	"time"

	clientset "github.com/YRXING/edwin/pkg/client/clientset/versioned"
	pfscheme "github.com/YRXING/edwin/pkg/client/clientset/versioned/scheme"
	informer "github.com/YRXING/edwin/pkg/client/informers/externalversions/edwin/v1"
	v1 "github.com/YRXING/edwin/pkg/client/listers/edwin/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const controllerName = "packetfilter-controller"

const (
	SuccessSynced         = "Synced"
	MessageResourceSynced = "PacketFilter synced successfully"
)

type PacketFilterController struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclinetset kubernetes.Interface
	// pfclientset is a clientset for our own API group
	pfclientset clientset.Interface

	pflister v1.PacketFilterLister
	pfSynced cache.InformerSynced

	workqueue workqueue.RateLimitingInterface

	recorder record.EventRecorder
}

// NewPacketFilterController returns a new student controller
func NewPacketFilterController(clientset kubernetes.Interface,
	pfclientset clientset.Interface,
	pfinformer informer.PacketFilterInformer) *PacketFilterController {

	utilruntime.Must(pfscheme.AddToScheme(pfscheme.Scheme))
	klog.Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{
		Interface: clientset.CoreV1().Events(""),
	})

	recorder := eventBroadcaster.NewRecorder(pfscheme.Scheme, corev1.EventSource{
		Component: controllerName,
	})

	controller := &PacketFilterController{
		kubeclinetset: clientset,
		pfclientset:   pfclientset,
		pflister:      pfinformer.Lister(),
		pfSynced:      pfinformer.Informer().HasSynced,
		workqueue:     workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		recorder:      recorder,
	}
	klog.Info("Setting up event handlers")

	pfinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    controller.enqueuePacketFilter,
		UpdateFunc: controller.updatePacketFilter,
		DeleteFunc: controller.enqueuePacketFilterDelete,
	})

	return controller
}

func (c *PacketFilterController) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	klog.Info("starting controller, waiting for cache synced")
	if ok := cache.WaitForCacheSync(stopCh, c.pfSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}
	klog.Info("start workers")
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	klog.Infof("%s and all works stopped", controllerName)
	return nil
}

func (c *PacketFilterController) runWorker() {
	for c.processNextWorkItem() {

	}
}

func (c *PacketFilterController) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()
	if shutdown {
		return false
	}

	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool

		// get changed resource's key from queue
		if key, ok = obj.(string); !ok {
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}

		// process business
		if err := c.syncHandler(key); err != nil {
			return fmt.Errorf("error syncing %s: %s", key, err.Error())
		}

		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced %s", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}
	return true
}

// process business
func (c *PacketFilterController) syncHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// get Object from cache
	packetfilter, err := c.pflister.PacketFilters(namespace).Get(name)
	if err != nil {
		// if object was deleted
		if errors.IsNotFound(err) {
			klog.Infof("Object was deleted")
			// TODO:delete logic....
			return nil
		}
		utilruntime.HandleError(fmt.Errorf("failed to list packetfilter by: %s %s", namespace, name))
		return err
	}

	klog.Infof("this is the expect status of this Object: %#v", packetfilter)
	//TODO: The actual status is obtained from the business level.
	//TODO: The actual status should be compared with the expected status,
	//TODO: and the response should be made according to the difference

	c.recorder.Event(packetfilter, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)

	return nil
}

// add event
// data should be put in cache first and the put in queue
func (c *PacketFilterController) enqueuePacketFilter(obj interface{}) {
	var (
		key string
		err error
	)

	// make keys for API Object in the format <namespace>/<name>
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}

	// put the key in queue
	c.workqueue.AddRateLimited(key)
}

func (c *PacketFilterController) updatePacketFilter(oldObj, newObj interface{}) {
	oldPf := oldObj.(*edwinv1.PacketFilter)
	newPf := newObj.(*edwinv1.PacketFilter)

	if oldPf.ResourceVersion == newPf.ResourceVersion {
		return
	}

	c.enqueuePacketFilter(newObj)
}

// delete event
func (c *PacketFilterController) enqueuePacketFilterDelete(obj interface{}) {
	var (
		key string
		err error
	)

	key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}

	c.workqueue.AddRateLimited(key)
}
