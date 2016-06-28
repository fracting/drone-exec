package script

import (
	"bytes"
	"fmt"

	"github.com/drone/drone-exec/parser"
	"github.com/drone/drone-plugin-go/plugin"
	"github.com/samalba/dockerclient"
)

// Encode encodes the build script as a command in the
// provided Container config. For linux, the build script
// is embedded as the container entrypoint command, base64
// encoded as a one-line script.
func Encode(w *plugin.Workspace, c *dockerclient.ContainerConfig, n *parser.DockerNode) {
	var buf bytes.Buffer
	buf.WriteString(setupScript)
	buf.WriteString(forceYesScript)

	if w != nil && w.Keys != nil && w.Netrc != nil {
		buf.WriteString(writeKey(
			w.Keys.Private,
		))
		buf.WriteString(writeNetrc(
			w.Netrc.Machine,
			w.Netrc.Login,
			w.Netrc.Password,
		))
	}

	if len(n.Shell) > 0 {
		buf.WriteString(fmt.Sprintf("%s << \"EOF\"", n.Shell))
	    buf.WriteString(setupScript)
	}

	buf.WriteString(writeCmds(n.Commands))

	if len(n.Shell) > 0 {
		buf.WriteString("EOF")
	}

	buf.WriteString(teardownScript)

	c.Entrypoint = entrypoint
	c.Cmd = []string{encode(buf.Bytes())}
}
