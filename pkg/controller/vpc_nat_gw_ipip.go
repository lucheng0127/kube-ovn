package controller

import (
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

func (c *Controller) enqueueAddVpcNatGwIpip(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	klog.V(3).Infof("enqueue add vpc-nat-gw-ipip %s", key)
	c.addOrUpdateVpcNatGatewayIpipQueue.Add(key)
}

func (c *Controller) enqueueUpdateVpcNatGwIpip(_, newObj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(newObj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	klog.V(3).Infof("enqueue update vpc-nat-gw-ipip %s", key)
	c.addOrUpdateVpcNatGatewayIpipQueue.Add(key)
}

func (c *Controller) enqueueDeleteVpcNatGwIpip(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	klog.V(3).Infof("enqueue del vpc-nat-gw-ipip %s", key)
	c.delVpcNatGatewayIpipQueue.Add(key)
}

func (c *Controller) runAddOrUpdateVpcNatGwIpipWorker() {
	for c.processNextWorkItem("addOrUpdateVpcNatGatewayIpip", c.addOrUpdateVpcNatGatewayIpipQueue, c.handleAddOrUpdateVpcNatGwIpip) {
	}
}

func (c *Controller) runDelVpcNatGwIpipWorker() {
	for c.processNextWorkItem("delVpcNatGatewayIpip", c.delVpcNatGatewayIpipQueue, c.handleDelVpcNatGwIpip) {
	}
}

func (c *Controller) handleAddOrUpdateVpcNatGwIpip(key string) error {
	// TODO(shawnlu): Implement it
	klog.Infof("add or update vpc nat gw ipip %s", key)
	return nil
}

func (c *Controller) handleDelVpcNatGwIpip(key string) error {
	// TODO(shawnlu): Implement it
	klog.Infof("delete vpc nat gw ipip %s", key)
	return nil
}
