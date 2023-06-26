/*
Package kubeconf provides safe read and write operations involving the
complicated (and often buggy) official Kubernetes configuration
(KUBECONFIG) as used by kubectl.

The Load function provides read-only access to configuration data
combined appropriately from several possible "origins" per KUBECONFIG
into a single api.Config struct (the same used by kubectl itself). This
is the *only* appropriate way of handling many possible combinations,
some, such as contexts, can reference multiple files in a single entry.
Note that there is a significant difference between the output of
kubectl config view and the official api.Config object used internal my
kubectl itself ("named lists" instead of maps, for example). Install the
kubectl-kubeconf plugin to view this data as JSON directly.

Writing and updating configuration values is handled in the only safe
way possible: by using the kubectl command itself.  Unfortunately, not
all of the actions supported by kubectl config behave in a consistent
manner. Some make modifications and deletions to other fields as
side-effects, while others happily accept empty values and overwrite
good values with blanks. This risk is reduced by using one of the
following encapsulating functions which all take a run.Args map which is
expanded into long-form (getopt-style) arguments with equal signs
allowing any option to kubectl to be passed through. To pass an empty
value, explicitly set it to an empty string in run.Args. Otherwise, omit
any unwanted argument keys.

    current-context CurrentContext Displays the current-context.
    use-context     UseContext     Sets the current-context in a kubeconfig file.
    set-context     SetContext     Sets a context entry in kubeconfig.
    rename-context  RenameContext  Renames a context from the kubeconfig file.
    delete-context  DeleteContext  Delete the specified context from the kubeconfig.

    set-cluster     SetCluster     Sets a cluster entry in kubeconfig.
    delete-cluster  DeleteCluster  Delete the specified cluster from the kubeconfig.

    set-credentials SetUser        Sets a user entry in kubeconfig.
    delete-user     DeleteUser     Delete the specified user from the kubeconfig.

    set             Set            Sets an individual value in a kubeconfig file.
    unset           Unset          Unsets an individual value in a kubeconfig file.

Note that the get-* and view actions have been omitted in favor of using
Load function.

WARNING

* Never assume $HOME/.kube/config is the only configuration file.

* Never assume $KUBECONFIG only contains a single path.

* Never change any value by accessing that file directly unless
  api.Config.*.LocationOrOrigin is observed! It is usually better to just call
  the equivalent using kubectl config set or similar.

* Removing a cluster or user does *not* remove these named references
	from contexts pointing to them. This is expected behavior for kubectl
	but may cause confusion with login plugins and such when the target of
	that named reference is found to not exist.

*/
package kubeconf
