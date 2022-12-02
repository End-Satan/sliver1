//go:build server

package assets

import "embed"

var (
	//go:embed fs/sliver.asc fs/*.txt fs/*.zip fs/windows/amd64/* fs/lib/*.a
	assetsFs embed.FS
)
