package cicd

type CICD struct {
	GithubAction *bool `yaml:"GitHubAction" json:"GitHubAction"`
	K8s          *bool `yaml:"k8s" json:"k8s"`
}

func (cicd *CICD) Validate() error {
	if cicd.GithubAction == nil {
		t := true
		cicd.GithubAction = &t
	}
	if cicd.K8s == nil {
		t := true
		cicd.K8s = &t
	}
	return nil
}
