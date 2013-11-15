# Lua for go-plugins

A runtime for [go-plugins](https://github.com/progrium/go-plugins) powered by [golua](https://github.com/aarzilli/golua) and [luar](https://github.com/stevedonovan/luar).

## Usage

Just register the runtime:

	plugins.RegisterRuntime(luaplugins.GetRuntime())


## License

BSD