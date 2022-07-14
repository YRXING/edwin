package main

import (
	"flag"
	clientset "github.com/YRXING/edwin/pkg/client/clientset/versioned"
	edwinInformers "github.com/YRXING/edwin/pkg/client/informers/externalversions"
	"github.com/YRXING/edwin/pkg/constants"
	"github.com/YRXING/edwin/pkg/controllers"
	"github.com/YRXING/edwin/pkg/signals"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"time"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig file")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API Server."+
		"Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
	flag.Parse()

	// process signals
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Infof("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	edwinClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building CRD clientset: %s", err.Error())
	}

	edwinInformerFactory := edwinInformers.NewSharedInformerFactory(edwinClient, time.Second*30)

	pfController := controllers.NewPacketFilterController(kubeClient, edwinClient, edwinInformerFactory.Edwin().V1().PacketFilters())

	go edwinInformerFactory.Start(stopCh)

	if err = pfController.Run(constants.THREADINESS, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}
