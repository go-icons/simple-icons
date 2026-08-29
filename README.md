# simple-icons

[![ci](https://github.com/go-icons/simple-icons/actions/workflows/ci.yml/badge.svg)](https://github.com/go-icons/simple-icons/actions/workflows/ci.yml)
![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-icons/simple-icons.svg)](https://pkg.go.dev/github.com/go-icons/simple-icons)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-blue.svg)](LICENSE)

The [Simple Icons](https://simpleicons.org) brand-logo set (CC0-1.0), embedded as
SVG and keyed by the icon's own name — for pure-Go UIs that render their own icons.

````go
import simpleicons "github.com/go-icons/simple-icons"

svg := simpleicons.Icon("reddit") // the Reddit logo, as an SVG string
````

`Icon(name)` returns the brand glyph by name (`reddit`, `instagram`, `x`,
`mastodon`, `bluesky`, `rss`, …), or `""` when unknown; `Has` and `Names`
enumerate the full set.

It is a **data package**: it returns SVG strings and draws nothing. Rendering is
a separate concern — a renderer such as
[go-widgets/toolkit](https://github.com/go-widgets/toolkit)'s `SVGIcon`, over the
[go-gfx](https://github.com/go-gfx/gfx) SVG rasteriser, turns the SVG into a
glyph. Simple Icons are single-path monochrome fills, so the renderer's ink
recolours them.

## Licence

The Go code is BSD-3-Clause (`LICENSE`). The embedded Simple Icons artwork is
CC0-1.0 (`SIMPLE-ICONS-LICENSE`) — redistributed unmodified. Brand names and
logos are trademarks of their respective owners.
