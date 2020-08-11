package cicd

type CICD struct {
	GithubAction *bool `yaml:"GitHubAction" json:"GitHubAction"`
	K8s          *bool `yaml:"k8s" json:"k8s"`
}

func (cicd *CICD) MergeWithDefault() error {
	t := true
	if cicd.GithubAction == nil {
		cicd.GithubAction = &t
	}
	if cicd.K8s == nil {
		cicd.K8s = &t
	}
	return nil
}

func (cicd *CICD) MergeWith(other *CICD) {
	if other == nil {
		return
	}
	if cicd.GithubAction == nil {
		cicd.GithubAction = other.GithubAction
	}
	if cicd.K8s == nil {
		cicd.K8s = other.K8s
	}
}
