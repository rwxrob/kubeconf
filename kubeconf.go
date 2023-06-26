package kubeconf

import (
	"fmt"

	"github.com/rwxrob/run-go"
	cc "k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// CurrentContext calls kubectl config current-context. Often this is
// better retrieved as a part of the api.Config object returned by Load
// function.
func CurrentContext(a run.Args) (string, error) {
	cmd := []string{`kubectl`, `config`, `current-context`}
	cmd = append(cmd, a.List()...)
	return run.OutErr(cmd...)
}

// UseContext calls kubectl config use-context.
func UseContext(a run.Args) (string, error) {
	name, hasname := a[`name`]
	if !hasname {
		return "", fmt.Errorf(`name must not be empty`)
	}
	delete(a, `name`)
	cmd := []string{`kubectl`, `config`, `use-context`, name}
	cmd = append(cmd, a.List()...)
	return run.OutErr(cmd...)
}

// SetContext calls kubectl config set-context. Only the name argument
// must be non-zero length. All others will be ignored if empty. See
// kubectl config set-context --help for more.
func SetContext(a run.Args) (string, error) {
	name, hasname := a[`name`]
	if !hasname {
		return "", fmt.Errorf(`name must not be empty`)
	}
	delete(a, `name`)
	cmd := []string{`kubectl`, `config`, `set-context`, name}
	cmd = append(cmd, a.ListEq()...)
	return run.OutErr(cmd...)
}

// RenameContext calls kubectl config rename-context. Requires 'new' and
// 'old' run.Args fields to be set.
func RenameContext(a run.Args) (string, error) {
	name, hasnew := a[`new`]
	if !hasnew {
		return "", fmt.Errorf(`new name must not be empty`)
	}
	delete(a, `new`)
	old, hasold := a[`old`]
	if !hasold {
		return "", fmt.Errorf(`old name must not be empty`)
	}
	delete(a, `old`)
	cmd := []string{`kubectl`, `config`, `rename-context`, old, name}
	cmd = append(cmd, a.List()...)
	return run.OutErr(cmd...)
}

// DeleteContext calls kubectl config delete-context.
func DeleteContext(a run.Args) (string, error) {
	name, hasname := a[`name`]
	if !hasname {
		return "", fmt.Errorf(`name must not be empty`)
	}
	delete(a, `name`)
	cmd := []string{`kubectl`, `config`, `delete-context`, name}
	cmd = append(cmd, a.List()...)
	return run.OutErr(cmd...)
}

// SetCluster calls kubectl config set-cluster. Only the name argument
// must be non-zero length. All others will be ignored if empty. See
// kubectl config set-cluster --help for more.
func SetCluster(a run.Args) (string, error) {
	name, hasname := a[`name`]
	if !hasname {
		return "", fmt.Errorf(`name must not be empty`)
	}
	delete(a, `name`)
	cmd := []string{`kubectl`, `config`, `set-cluster`, name}
	cmd = append(cmd, a.ListEq()...)
	return run.OutErr(cmd...)
}

// DeleteCluster calls kubectl config delete-cluster.
func DeleteCluster(a run.Args) (string, error) {
	name, hasname := a[`name`]
	if !hasname {
		return "", fmt.Errorf(`name must not be empty`)
	}
	delete(a, `name`)
	cmd := []string{`kubectl`, `config`, `delete-cluster`, name}
	cmd = append(cmd, a.ListEq()...)
	return run.OutErr(cmd...)
}

// SetUser calls kubectl config set-credentials. Only the name argument
// must be non-zero length. All others will be ignored if empty. See
// kubectl config set-credentials --help for more. Note that changing
// the type of credential (from --token to --auth-provider) might leave
// residual fields in the configuration. In such cases, calling
// DeleteUser should be done first before creating a new one with
// SetUser. Note that the terms "user", "credentials" and "authinfo" are
// frequently used synonymously in Kubernetes documentation. Also note
// that unlike other applications Kubernetes uses id_access user.token
// instead of access_token.
func SetUser(a run.Args) (string, error) {
	name, hasname := a[`name`]
	if !hasname {
		return "", fmt.Errorf(`name must not be empty`)
	}
	delete(a, `name`)
	cmd := []string{`kubectl`, `config`, `set-credentials`, name}
	cmd = append(cmd, a.ListEq()...)
	return run.OutErr(cmd...)
}

// DeleteUser calls kubectl config delete-user.
func DeleteUser(a run.Args) (string, error) {
	name, hasname := a[`name`]
	if !hasname {
		return "", fmt.Errorf(`name must not be empty`)
	}
	delete(a, `name`)
	cmd := []string{`kubectl`, `config`, `delete-user`, name}
	cmd = append(cmd, a.ListEq()...)
	return run.OutErr(cmd...)
}

// Set calls kubectl config set. Only the name argument
// must be non-zero length. All others will be ignored if empty. See
// kubectl config set --help for more. Note that an empty value will
// still be set as empty.
func Set(a run.Args) (string, error) {
	name, hasname := a[`name`]
	if !hasname {
		return "", fmt.Errorf(`name must not be empty`)
	}
	delete(a, `name`)
	value, hasvalue := a[`value`]
	if !hasvalue {
		return "", fmt.Errorf(`value must not be empty`)
	}
	delete(a, `value`)
	cmd := []string{`kubectl`, `config`, `set`, name, value}
	cmd = append(cmd, a.ListEq()...)
	return run.OutErr(cmd...)
}

// Unset calls kubectl config unset. Only the name argument
// must be non-zero length. All others will be ignored if empty. See
// kubectl config unset --help for more. Note that an empty value will
// still be set as empty.
func Unset(a run.Args) (string, error) {
	name, hasname := a[`name`]
	if !hasname {
		return "", fmt.Errorf(`name must not be empty`)
	}
	delete(a, `name`)
	cmd := []string{`kubectl`, `config`, `unset`, name}
	cmd = append(cmd, a.ListEq()...)
	return run.OutErr(cmd...)
}

// Load returns an api.Config struct created using the default loading rules
// for the clientcmd/api (the same as kubectl itself). See kubectl
// config --help for more information. This is useful for viewing how
// kubectl actually resolved loading multiple configurations from
// different files.
func Load() (api.Config, error) {
	return cc.NewNonInteractiveDeferredLoadingClientConfig(
		cc.NewDefaultClientConfigLoadingRules(), nil).RawConfig()
}
