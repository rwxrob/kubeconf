package kubeconf_test

/*
func ExampleUseContext() {
	out, err := kubeconf.UseContext(run.Args{`name`: `foo`})
	fmt.Print(out, err)
	// Output:
	// Switched to context "foo".
	// <nil>
}

func ExampleRenameContext() {

	out, err := kubeconf.RenameContext(run.Args{
		`old`: `foo`,
		`new`: `bar`,
	})
	fmt.Print(out, err)

	// Output:
	// bar
}

func ExampleSetContext() {
	out, err := kubeconf.SetContext(run.Args{
		`name`:      `foo`,
		`cluster`:   `fcluster`,
		`user`:      `some`,
		`namespace`: `fnamespace`,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	// Output:
	// Context "foo" modified.
}

func ExampleDeleteCluster() {
	out, err := kubeconf.DeleteCluster(run.Args{`name`: `fcluster`})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
	// Output:
	// nothing
}

func ExampleSetCluster() {
	out, err := kubeconf.SetCluster(run.Args{
		`name`:   `foo`,
		`server`: `http://localhost:8443`,
		`user`:   `cuser`,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	// Output:
	// Cluster "foo" set.
}

func ExampleSetUser() {
	out, err := kubeconf.SetUser(run.Args{
		`name`:  `herman`,
		`token`: `mybearertoken`,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	// Output:
	// User "herman" set.
}

func ExampleCurrentContext() {
	out, err := kubeconf.CurrentContext(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	// Output:
	// foo
}

func ExampleDeleteUser() {
	out, err := kubeconf.DeleteUser(run.Args{`name`: `blah`})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
	// Output:
	// nothing
}

func ExampleSet() {
	out, err := kubeconf.Set(run.Args{
		`name`:  `users.bill.token`,
		`value`: `mybearertoken`,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	// Output:
	// Property "users.bill.token" set.
}

func ExampleUnset() {
	out, err := kubeconf.Unset(run.Args{
		`name`: `users.bill.token`,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	// Output:
	// Property "users.bill.token" unset.

}
*/
