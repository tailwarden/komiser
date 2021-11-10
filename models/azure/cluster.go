package azure

type Cluster struct {
	Name              string `json:"name"`
	NodeResourceGroup string `json:"nodeResourceGroup"`
	KubernetesVersion string `json:"kubernetesVersion"`
}
