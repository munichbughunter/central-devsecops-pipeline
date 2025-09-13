package ci

import (
	"context"

	"github.com/munichbughunter/central-devsecops-pipeline/internal/dagger"
)

func RunPythonBuild(ctx context.Context, client *dagger.Client, src *dagger.Directory, pythonVersion string) (string, error) {
	return client.Container().
		From("python:"+pythonVersion).
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"python", "app.py"}).
		Stdout(ctx)
}
