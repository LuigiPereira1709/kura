package executil

import (
	"context"
	"os/exec"
	"strings"
)

// Exec executes a command in the given context and returns the command, its stdout and stderr pipes, and any error encountered during execution.
func Exec(ctx context.Context, name string, args ...string) *exec.Cmd {
	return exec.CommandContext(ctx, name, args[0:]...)
}

func ExecShell(ctx context.Context, command string) *exec.Cmd {
	if strings.TrimSpace(command) == "" {
		return nil
	}

	return exec.CommandContext(ctx, "sh", "-c", command)
}
