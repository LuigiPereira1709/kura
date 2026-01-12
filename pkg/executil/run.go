package executil

import (
	"context"
	"os/exec"
)

// ExecCommand executes a command in the given context and returns the command, its stdout and stderr pipes, and any error encountered during execution.
func ExecCommand(ctx context.Context, command ...string) *exec.Cmd {
	if len(command) == 0 {
		return nil
	}
	cmd := exec.CommandContext(ctx, command[0], command[1:]...)

	return cmd
}
