# Crisp

[![GitHub Release](https://img.shields.io/github/release/ahstn/crisp.svg?style=flat-square)](https://github.com/ahstn/crisp/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/ahstn/crisp?style=flat-square)](https://goreportcard.com/report/github.com/ahstn/crisp)
[![Software License](https://img.shields.io/github/license/ahstn/crisp.svg?style=flat-square)](LICENSE)

Crisp is a fast and minimal shell prompt written in [Go].

**WIP:** Will fill this out eventually

## Licence Complicance (GNU GPL 3.0)
This project was orginally forked from [Mímir] which has a [GNU GPLv3] license.
With this comes the need to state my changes which at the time of writing are as
follows:
* Split out "modules" (Git branch, Kube Context, etc) into their own packages.
* Altered the look of Kube Context.
* Adding Git dirty check.

[Go]: https://golang.org
[Mímir]: https://github.com/talal/mimir
[GNU GPLv3]: https://choosealicense.com/licenses/gpl-3.0/
