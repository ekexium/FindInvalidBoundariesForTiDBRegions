A tool to find regions with invalid boundaries. It's used to help recover from a bad situation caused by https://github.com/tikv/tikv/issues/10542.

Usage:
1. `pd-ctl region > region.json`
2. put json under the project directory
3. `go run main.go`

How to recover:
Merge the regions with invalid boundaries with normal ones.