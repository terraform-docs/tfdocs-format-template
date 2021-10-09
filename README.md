# tfdocs-format-template

`tfdocs-format-template` is a minimal [terraform-docs] plugin that is built with
[plugin SDK] and  is meant to be used as a template for implementing new plugin.
Most of the time you only need to modify and reorder the go template defined in
`sections.tmpl`.

## Requirements

- [Go] 1.16+
- [terraform-docs] v0.16+

Optionally for releasing:

- [gox]

## Installation

In order to install a plugin the following steps are needed:

- download the plugin and place it in `~/.tfdocs.d/plugins` (or `./.tfdocs.d/plugins`)
- make sure the plugin file name is `tfdocs-format-<NAME>` (e.g `<NAME>` should be `template`)
- modify `formatter` of `.terraform-docs.yml` file to be `<NAME>`

**Important notes:**

- if the plugin file name is different than the example above, `terraform-docs` won't
be able to to pick it up nor register it properly.
- you can only use plugin thorough `.terraform-docs.yml` file and it cannot be used
with CLI arguments

## Building

Create a repository from this by clicking `Use this template` and then clone the
newly created repository locally and run the following command:

```bash
make
```

You can then install the plugin to `~/.tfdocs.d/plugins` by running the following:

```bash
make install
```

Additionally you can override the destination of installation:

```bash
PLUGIN_FOLDER=/path/to/plugin/folder make install
```

Note that the plugin has to be built for target OS and architecture (`make build`
and `make install` do that,) but if you want to redistribute the plugin for other
people to use you have to cross-compile it (for example you can use [gox].)

## Docker

If you don't want cross-compile your plugin, you can pack it to the docker image. 
So the only dependency needed the user/or a server machine is `docker`.

### Building and publishing docker image

```bash
make docker
```

Before you will build your docker image you have to replace very few things in the `Makefile` and the `Dockerfile`.

Things to replace:

`Makefile`:
- DOCKER_REGISTRY (could be any docker registry)
- PROJECT_OWNER (how the the place file in your registry. In our company it's something like `devops/tools`)
- BUILD_NAME (your project name should follow the patter `tfdocs-format-<uniqueName>` )

`Dockerfile`:
- all **PROJECT_NAME** entries should be replaced with your project name( and it should follow the pattern `tfdocs-format-<uniqueName>`)

After all changes and `make docker` successful execution, you can push your new perfect image to the registry by:
```bash
make push
```

### Using docker image with your custom plugin

I'm a little lazy, and I made a simple sh script for easy usage of docker image with terraform-docs and the custom plugin:

```shell
#!/bin/bash

img=<link to the docker image with terraform-docs and custom plugin>

dir=$1
out_file=$2

function printMarkdownTo() {
  pushd $1
    (docker run --rm -v `pwd`:`pwd` -w `pwd` $img .) > $2
  popd
}

printMarkdownTo ${dir} ${out_file}

```

## Links

[terraform-docs]: https://github.com/terraform-docs/terraform-docs
[plugin SDK]: https://github.com/terraform-docs/terraform-docs/tree/master/plugin
[Go]: https://golang.org/
[gox]: https://github.com/mitchellh/gox
