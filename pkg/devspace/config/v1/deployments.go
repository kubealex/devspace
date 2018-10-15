package v1

// DeploymentConfig defines the configuration how the devspace should be deployed
type DeploymentConfig struct {
	Name      *string        `yaml:"name"`
	Namespace *string        `yaml:"namespace"`
	Helm      *HelmConfig    `yaml:"helm,omitempty"`
	Kubectl   *KubectlConfig `yaml:"kubectl,omitempty"`
}

// HelmConfig defines the specific helm options used during deployment
type HelmConfig struct {
	ReleaseName *string                      `yaml:"releaseName,omitempty"`
	ChartPath   *string                      `yaml:"chartPath,omitempty"`
	Values      *map[interface{}]interface{} `yaml:"values,omitempty"`
}

// KubectlConfig defines the specific kubectl options used during deployment
type KubectlConfig struct {
	CmdPath   *string    `yaml:"cmdPath,omitempty"`
	Manifests *[]*string `yaml:"manifests,omitempty"`
}
