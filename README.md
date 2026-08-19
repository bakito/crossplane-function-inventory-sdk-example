# crossplane-function-inventory-sdk-example

Example [Crossplane](https://www.crossplane.io/) project built with the
[crossplane-function-inventory-sdk](https://github.com/bakito/crossplane-function-inventory-sdk).
It packages composition functions (under `functions/`), their composite resource APIs (under
`apis/`), and pushes the resulting project package to `ghcr.io/bakito/crossplane-function-inventory-sdk-example`.

## Prerequisites

The only thing you need installed globally is [mise](https://mise.jdx.dev/) — it manages every
other tool this project needs, pinned to exact versions in [`mise.toml`](./mise.toml).

Install with:

```shell
# Trusts the mise.toml
mise trust .
# Installs all tools
mise install
# Setup mise env
eval $(mise activate --shims )
# Configures crossplane specific settings/tools
task setup
```

## Repository layout

- `functions/<name>/` — one Go module per composition function (Crossplane function-sdk-go).
- `apis/<name>/` — the `schema.yaml` (SimpleSchema source), generated `definition.yaml` (XRD),
  and `composition.yaml` for each composite resource.
- `schemas/` — Go types generated from the APIs (regenerated, not hand-edited).
- `examples/` — example composite resources (XRs) for local testing.
- `crossplane-project.yaml` — the Crossplane project manifest (dependencies, repository).
- `_output/` — build artifacts (`*.xpkg` package, `ko/*.tar` function image tarballs).

Everything under `functions/` and `apis/` is discovered dynamically by the `Taskfile.yml` — add
a new folder in either, and it's automatically picked up by `generate`, `build`, `test`, `lint`,
and `format`.

## Development workflow

All commands are run through `task` (see `task --list-all` for the full list with descriptions).

- **Changed an API schema or a function's tagged struct?** Run `task generate`
  to regenerate Go code (`go generate`), Go model types (`schemas/`), and XRD
  definitions (`apis/*/definition.yaml`).
- **Changed function code?** Run `task test` (go test), `task lint`
  (golangci-lint run), and `task format` (golangci-lint fmt) — each runs
  across every function.
- **Want to build?** Run `task build` to build every function image (via
  `ko`, not pushed) and package the project. Produces
  `_output/ko/<function-name>.tar` per function and
  `_output/crossplane-function-inventory-sdk-example.xpkg`.
- **Want to push?** Run `task publish VERSION=v0.1.0` to build and publish
  the package to the registry.

