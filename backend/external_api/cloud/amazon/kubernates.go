package amazon

type IKubernates interface {
	ListCluster()
	CreateCluster()
	DeleteCluster()

	ListNodeGroup()
	CreateNodeGroup()
	DeleteNodeGroup()
	ScaleNodeGroup()
}
