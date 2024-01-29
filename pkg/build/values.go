// Package build is just a package of variables set at build time to store build information.
package build

var (
	//nolint: gochecknoglobals // There are set at build time, so they are global and cannot be constants unfortunately.
	Build = "snapshot"
	//nolint: gochecknoglobals // There are set at build time, so they are global and cannot be constants unfortunately.
	Commit = ""
	//nolint: gochecknoglobals // There are set at build time, so they are global and cannot be constants unfortunately.
	Date = ""
	//nolint: gochecknoglobals // There are set at build time, so they are global and cannot be constants unfortunately.
	Runtime = ""
	//nolint: gochecknoglobals // There are set at build time, so they are global and cannot be constants unfortunately.
	OS = ""
	//nolint: gochecknoglobals // There are set at build time, so they are global and cannot be constants unfortunately.
	Architecture = ""
	//nolint: gochecknoglobals // There are set at build time, so they are global and cannot be constants unfortunately.
	ARM = ""
	//nolint: gochecknoglobals // There are set at build time, so they are global and cannot be constants unfortunately.
	AMD64 = ""
)
