package k8s

type RestartDeployment struct {
	Namespace      string `json:"namespace"  binding:"required"`
	DeploymentName string `json:"deployment_name" binding:"required"`
}

type ScaleDeployment struct {
	Namespace string `json:"namespace"`
	DeploymentName string `json:"deployment_name"`
	ScaleNumber *int32 `json:"scale_number"`
}