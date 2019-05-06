package kubernetes

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/talal/go-bits/color"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Contexts []struct {
		Context struct {
			Cluster   string `yaml:"cluster"`
			Namespace string `yaml:"namespace"`
			User      string `yaml:"user"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
}

// Context returns the current Kubernetes context and namespace.
// $KUBECONFIG will be used if it's set, otherwise default to ~/.kube/config
func Context() string {
	if os.Getenv("CRISP_KUBE") == "0" {
		return ""
	}
	path := os.Getenv("HOME") + "/.kube/config"
	if env, ok := os.LookupEnv("KUBECONFIG"); ok {
		path = env
	}

	context, namespace := clusterInfo(path)
	return fmt.Sprintf(" ☸️ %v:%v", color.Sprintf(color.Yellow, context),
		color.Sprintf(color.White, namespace))
}

func clusterInfo(configPath string) (string, string) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", ""
	}

	cfg := Config{}
	if err = yaml.Unmarshal(buf, &cfg); err != nil {
		fmt.Println(err)
	}

	if cfg.CurrentContext == "" {
		return "", ""
	}

	for _, c := range cfg.Contexts {
		if c.Name == cfg.CurrentContext {
			if c.Context.Namespace == "" {
				c.Context.Namespace = "default"
			}
			return strings.TrimSpace(cfg.CurrentContext),
				strings.TrimSpace(c.Context.Namespace)
		}
	}

	return cfg.CurrentContext, ""
}
