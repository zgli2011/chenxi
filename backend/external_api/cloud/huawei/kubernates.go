package huawei

type IKubernates interface {
	ListCluster()
	CreateCluster()
	DeleteCluster()

	ListNodeGroup()
	CreateNodeGroup()
	DeleteNodeGroup()
	ScaleNodeGroup()
}
