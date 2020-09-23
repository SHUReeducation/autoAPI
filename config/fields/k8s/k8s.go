package k8s

import "github.com/urfave/cli/v2"

type K8s struct {
	Host      *string `yaml:"host" json:"host" toml:"host"`
	Uri       *string `yaml:"uri" json:"uri" toml:"uri"`
	Namespace *string `yaml:"namespace" json:"namespace" toml:"namespace"`
}

func (k8s *K8s) MergeWith(other *K8s) {
	if other == nil {
		return
	}
	if k8s.Host == nil {
		k8s.Host = other.Host
	}
	if k8s.Uri == nil {
		k8s.Uri = other.Uri
	}
	if k8s.Namespace == nil {
		k8s.Namespace = other.Namespace
	}
}

func FromCommandLine(c *cli.Context) (*K8s, error) {
	if c.Bool("nok8s") {
		return nil, nil
	}
	var result K8s
	if host := c.String("host"); host != "" {
		result.Host = &host
	}
	return &result, nil
}
