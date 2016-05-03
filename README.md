#go-figure
*go-figure* is a multi-file, multi-configuration package configuration  to handle old-school configuration techniques such searching for as multiple files in multiple folders (~/something.d, /etc/something.d, ...).

It does expand standard shell parameters in the path list via os.ExpandEnv().

#Usage

Generally the following should be done:

```go
    ...
    v := viper.New()
    v.SetConfigType("yml|json|yaml")
    gofigure.Parse(v, []string{"${HOME}/.app.d", "${CFGD}", "/etc/app.d"})
    ...
    //do something with v

```
