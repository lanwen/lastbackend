//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package storage

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/common/types"
	"golang.org/x/oauth2"
)

type IUtil interface {
	Key(ctx context.Context, pattern ...string) string
}

type IStorage interface {
	Activity() IActivity
	Build() IBuild
	Hook() IHook
	Image() IImage
	Namespace() INamespace
	Service() IService
	Pod() IPod
	Node() INode
	System() ISystem
	Vendor() IVendor
	Volume() IVolume
	Endpoint() IEndpoint
}

type IActivity interface {
	Insert(ctx context.Context, activity *types.Activity) error
	ListProjectActivity(ctx context.Context, id string) ([]*types.Activity, error)
	ListServiceActivity(ctx context.Context, id string) ([]*types.Activity, error)
	RemoveByProject(ctx context.Context, id string) error
	RemoveByService(ctx context.Context, id string) error
}

type IBuild interface {
	GetByID(ctx context.Context, imageID, id string) (*types.Build, error)
	ListByImage(ctx context.Context, id string) ([]*types.Build, error)
	Insert(ctx context.Context, image string, build *types.Build) error
}

type IHook interface {
	Get(ctx context.Context, id string) (*types.Hook, error)
	Insert(ctx context.Context, hook *types.Hook) error
	Remove(ctx context.Context, id string) error
}

type INamespace interface {
	GetByName(ctx context.Context, name string) (*types.Namespace, error)
	List(ctx context.Context) ([]*types.Namespace, error)
	Insert(ctx context.Context, namespace *types.Namespace) error
	Update(ctx context.Context, project *types.Namespace) error
	Remove(ctx context.Context, id string) error
}

type IService interface {
	CountByNamespace(ctx context.Context, namespace string) (int, error)
	GetByName(ctx context.Context, namespace, name string) (*types.Service, error)
	GetByPodName(ctx context.Context, name string) (*types.Service, error)
	ListByNamespace(ctx context.Context, namespace string) ([]*types.Service, error)
	Insert(ctx context.Context, service *types.Service) error
	Update(ctx context.Context, service *types.Service) error
	UpdateSpec(ctx context.Context, service *types.Service) error
	Remove(ctx context.Context, service *types.Service) error
	RemoveByNamespace(ctx context.Context, namespace string) error

	Watch(ctx context.Context, service chan *types.Service) error
	SpecWatch(ctx context.Context, service chan *types.Service) error
	PodsWatch(ctx context.Context, service chan *types.Service) error
	BuildsWatch(ctx context.Context, service chan *types.Service) error
}

type IPod interface {
	GetByName(ctx context.Context, namespace, name string) (*types.Pod, error)
	ListByNamespace(ctx context.Context, namespace string) (map[string]*types.Pod, error)
	ListByService(ctx context.Context, namespace, service string) ([]*types.Pod, error)
	Upsert(ctx context.Context, namespace string, pod *types.Pod) error
	Update(ctx context.Context, namespace string, pod *types.Pod) error
	Remove(ctx context.Context, namespace string, pod *types.Pod) error
	Watch(ctx context.Context, pod chan *types.Pod) error
}

type IImage interface {
	Get(ctx context.Context, name string) (*types.Image, error)
	Insert(ctx context.Context, source *types.Image) error
	Update(ctx context.Context, image *types.Image) error
}

type IVendor interface {
	Insert(ctx context.Context, owner, name, host, serviceID string, token *oauth2.Token) error
	Get(ctx context.Context, name string) (*types.Vendor, error)
	List(ctx context.Context) (map[string]*types.Vendor, error)
	Update(ctx context.Context, vendor *types.Vendor) error
	Remove(ctx context.Context, vendorName string) error
}

type IVolume interface {
	GetByToken(ctx context.Context, token string) (*types.Volume, error)
	ListByNamespace(ctx context.Context, namespace string) ([]*types.Volume, error)
	Insert(ctx context.Context, volume *types.Volume) error
	Remove(ctx context.Context, id string) error
}

type INode interface {
	List(ctx context.Context) ([]*types.Node, error)

	Get(ctx context.Context, hostname string) (*types.Node, error)
	Insert(ctx context.Context, node *types.Node) error

	Update(ctx context.Context, node *types.Node) error

	InsertPod(ctx context.Context, meta *types.NodeMeta, pod *types.PodNodeSpec) error
	UpdatePod(ctx context.Context, meta *types.NodeMeta, pod *types.PodNodeSpec) error
	RemovePod(ctx context.Context, meta *types.NodeMeta, pod *types.PodNodeSpec) error

	Remove(ctx context.Context, meta *types.Node) error
	Watch(ctx context.Context, node chan *types.Node) error
}

type ISystem interface {
	ProcessSet(ctx context.Context, process *types.Process) error

	Elect(ctx context.Context, process *types.Process) (bool, error)
	ElectUpdate(ctx context.Context, process *types.Process) error
	ElectWait(ctx context.Context, process *types.Process, lead chan bool) error
}

type IEndpoint interface {
	Get(ctx context.Context, name string) ([]string, error)
	Upsert(ctx context.Context, name string, ips []string) error
	Remove(ctx context.Context, name string) error
	Watch(ctx context.Context, endpoint chan string) error
}
