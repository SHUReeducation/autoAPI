package cicd

type CICD struct {
	GithubAction *bool `yaml:"GitHubAction"`
	K8s          *bool `yaml:"k8s"`
}
