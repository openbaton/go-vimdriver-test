/*
 *  Copyright (c) 2017 Open Baton (http://openbaton.org)
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package main

import (
	"fmt"
	"time"

	"github.com/openbaton/go-openbaton/catalogue"
	"github.com/openbaton/go-openbaton/util"
	log "github.com/sirupsen/logrus"
)

type driver struct {
	*log.Logger
}

func (d driver) AddFlavour(vimInstance *catalogue.VIMInstance, deploymentFlavour *catalogue.DeploymentFlavour) (*catalogue.DeploymentFlavour, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return deploymentFlavour, nil
}

func (d driver) AddImage(vimInstance *catalogue.VIMInstance, image *catalogue.NFVImage, imageFile []byte) (*catalogue.NFVImage, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return image, nil
}

func (d driver) AddImageFromURL(vimInstance *catalogue.VIMInstance, image *catalogue.NFVImage, imageURL string) (*catalogue.NFVImage, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return image, nil
}

func (d driver) CopyImage(vimInstance *catalogue.VIMInstance, image *catalogue.NFVImage, imageFile []byte) (*catalogue.NFVImage, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return image, nil
}

func (d driver) CreateNetwork(vimInstance *catalogue.VIMInstance, network *catalogue.Network) (*catalogue.Network, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return network, nil
}

func (d driver) CreateSubnet(vimInstance *catalogue.VIMInstance, createdNetwork *catalogue.Network, subnet *catalogue.Subnet) (*catalogue.Subnet, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return subnet, nil
}

func (d driver) DeleteFlavour(vimInstance *catalogue.VIMInstance, extID string) (bool, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return true, nil
}

func (d driver) DeleteImage(vimInstance *catalogue.VIMInstance, image *catalogue.NFVImage) (bool, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return true, nil
}

func (d driver) DeleteNetwork(vimInstance *catalogue.VIMInstance, extID string) (bool, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return true, nil
}

func (d driver) DeleteServerByIDAndWait(vimInstance *catalogue.VIMInstance, id string) error {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	time.Sleep(3 * time.Second)
	return nil
}

func (d driver) DeleteSubnet(vimInstance *catalogue.VIMInstance, existingSubnetExtID string) (bool, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return true, nil
}

func (d driver) LaunchInstance(
	vimInstance *catalogue.VIMInstance,
	name, image, Flavour, keypair string,
	networks []*catalogue.VNFDConnectionPoint,
	secGroup []string,
	userData string) (*catalogue.Server, error) {

	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return newServer(), nil
}

func (d driver) LaunchInstanceAndWait(
	vimInstance *catalogue.VIMInstance,
	hostname, image, extID, keyPair string,
	networks []*catalogue.VNFDConnectionPoint,
	securityGroups []string,
	s string) (*catalogue.Server, error) {

	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return d.LaunchInstanceAndWaitWithIPs(vimInstance, hostname, image, extID, keyPair, networks, securityGroups, s, nil, nil)
}

func (d driver) LaunchInstanceAndWaitWithIPs(
	vimInstance *catalogue.VIMInstance,
	hostname, image, extID, keyPair string,
	network []*catalogue.VNFDConnectionPoint,
	securityGroups []string,
	s string,
	floatingIps map[string]string,
	keys []*catalogue.Key) (*catalogue.Server, error) {

	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	time.Sleep(3 * time.Second)

	return newServer(), nil
}

func (d driver) ListFlavours(vimInstance *catalogue.VIMInstance) ([]*catalogue.DeploymentFlavour, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return []*catalogue.DeploymentFlavour{
		{ExtID: "ext_id_1", FlavourKey: "flavour_name"},
		{ExtID: "ext_id_6", FlavourKey: "m1.nano"},
		{ExtID: "ext_id_2", FlavourKey: "m1.tiny"},
		{ExtID: "ext_id_3", FlavourKey: "m1.small"},
	}, nil
}

func (d driver) ListImages(vimInstance *catalogue.VIMInstance) ([]*catalogue.NFVImage, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	nfvImages := []*catalogue.NFVImage{
		{ExtID: "ext_id", Name: "ubuntu-14.04-server-cloudimg-amd64-disk1"},
	}

	for i := 1; i <= 20; i++ {
		nfvImages = append(nfvImages, &catalogue.NFVImage{
			ExtID: fmt.Sprintf("ext_id_%d", i),
			Name:  fmt.Sprintf("image_name_%d", i),
		})
	}

	return nfvImages, nil
}

func (d driver) ListNetworks(vimInstance *catalogue.VIMInstance) ([]*catalogue.Network, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return []*catalogue.Network{newNetwork("network-id")}, nil
}

func (d driver) ListServer(vimInstance *catalogue.VIMInstance) ([]*catalogue.Server, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return []*catalogue.Server{
		{
			Name:  "server_name",
			ExtID: "ext_id",
			Flavour: &catalogue.DeploymentFlavour{
				RAM:   10,
				VCPUs: 1,
			},
			IPs: make(map[string][]string),
		},
	}, nil
}

func (d driver) NetworkByID(vimInstance *catalogue.VIMInstance, id string) (*catalogue.Network, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return newNetwork(id), nil
}

func (d driver) Quota(vimInstance *catalogue.VIMInstance) (*catalogue.Quota, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return newQuota(), nil
}

func (d driver) SubnetsExtIDs(vimInstance *catalogue.VIMInstance, networkExtID string) ([]string, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return []string{}, nil
}

func (d driver) Type(vimInstance *catalogue.VIMInstance) (string, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return "test", nil
}

func (d driver) UpdateFlavour(vimInstance *catalogue.VIMInstance, deploymentFlavour *catalogue.DeploymentFlavour) (*catalogue.DeploymentFlavour, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return deploymentFlavour, nil
}

func (d driver) UpdateImage(vimInstance *catalogue.VIMInstance, image *catalogue.NFVImage) (*catalogue.NFVImage, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return image, nil
}

func (d driver) UpdateNetwork(vimInstance *catalogue.VIMInstance, network *catalogue.Network) (*catalogue.Network, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return network, nil
}

func (d driver) UpdateSubnet(vimInstance *catalogue.VIMInstance, createdNetwork *catalogue.Network, subnet *catalogue.Subnet) (*catalogue.Subnet, error) {
	tag := util.FuncName()

	d.WithFields(log.Fields{
		"tag":          tag,
		"vim-instance": vimInstance.ID,
	}).Debug("received request")

	return subnet, nil
}

func newServer() *catalogue.Server {
	return &catalogue.Server{
		Name:           "server_name",
		ExtID:          "ext_id",
		Created:        catalogue.NewDate(),
		FloatingIPs:    make(map[string]string),
		ExtendedStatus: "ACTIVE",
		Flavour: &catalogue.DeploymentFlavour{
			Disk:       100,
			ExtID:      "ext",
			FlavourKey: "m1.small",
			RAM:        2000,
			VCPUs:      4,
		},
		IPs: make(map[string][]string),
	}
}

func newNetwork(id string) *catalogue.Network {
	return &catalogue.Network{
		Name:  "network_name",
		ID:    id,
		ExtID: "ext_id",
	}
}

func newQuota() *catalogue.Quota {
	return &catalogue.Quota{
		Cores:       99999,
		FloatingIPs: 99999,
		Instances:   99999,
		KeyPairs:    99999,
		RAM:         99999,
		Tenant:      "test-tenant",
	}
}
