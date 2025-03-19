package controller

import (
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

func (c *Controller) enqueueAddVpcBmsConnection(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	klog.V(3).Infof("enqueue add vpc-bms-connection %s", key)
	c.addOrUpdateVpcBmsConnectionQueue.Add(key)
}

func (c *Controller) enqueueUpdateVpcBmsConnection(_, newObj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(newObj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	klog.V(3).Infof("enqueue update vpc-bms-connection %s", key)
	c.addOrUpdateVpcBmsConnectionQueue.Add(key)
}

func (c *Controller) enqueueDeleteVpcBmsConnection(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	klog.V(3).Infof("enqueue del vpc-bms-connection %s", key)
	c.delVpcBmsConnectionQueue.Add(key)
}

func (c *Controller) runAddOrUpdateVpcBmsConnectionWorker() {
	for c.processNextWorkItem("addOrUpdateVpcBmsConnection", c.addOrUpdateVpcNatGatewayIpipQueue, c.handleAddOrUpdateVpcBmsConnection) {
	}
}

func (c *Controller) runDelVpcBmsConnectionWorker() {
	for c.processNextWorkItem("delVpcBmsConnection", c.delVpcNatGatewayIpipQueue, c.handleDelVpcBmsConnection) {
	}
}

func (c *Controller) handleAddOrUpdateVpcBmsConnection(key string) error {
	// TODO(shawnlu): Implement it
	klog.Infof("add or update vpc bms connection %s", key)
	return nil
}

func (c *Controller) handleDelVpcBmsConnection(key string) error {
	// TODO(shawnlu): Implement it
	klog.Infof("delete vpc bms connection %s", key)
	return nil
}
