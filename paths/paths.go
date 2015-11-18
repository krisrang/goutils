package paths

import (
	"github.com/kardianos/osext"
)

// SelfPath returns the location of currently running executable.
func SelfPath() string {
	filename, _ := osext.Executable()

	return filename
}
