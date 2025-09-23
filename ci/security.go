package ci

import (
	"context"

	"github.com/munichbughunter/central-devsecops-pipeline/internal/dagger"
)

// RunImageScan erzeugt eine SARIF-Datei für ein bestehendes Image (mit GHCR-Auth)
func RunImageScan(ctx context.Context, dag *dagger.Client, imageRef string, githubUsername string, githubToken *dagger.Secret) (*dagger.File, error) {
	ctr := dag.
		Container().
		WithRegistryAuth("ghcr.io", githubUsername, githubToken).
		From(imageRef)

	out, err := dag.Trivy().ScanContainer(ctx, ctr, dagger.TrivyScanContainerOpts{
		Severity:      "HIGH,CRITICAL,MEDIUM",
		ExitCode: 0,      // nicht failen; wir wollen SARIF erzeugen
		Format:   "sarif",
		TrivyImageTag: "latest",
	})
	if err != nil {
		return nil, err
	}
	return dag.File("trivy-report.sarif", out), nil
}

// RunSemgrepScan erzeugt eine SARIF-Datei für ein bestehendes Image (mit GHCR-Auth)
func RunSemgrepScan(ctx context.Context, dag *dagger.Client, src *dagger.Directory) (*dagger.File) {
	return dag.
		Container().
		From("returntocorp/semgrep:latest").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{
			"semgrep",
			"--config=auto",
			"--sarif",
			"--output=semgrep-report.sarif",
			".",
		}).
		File("semgrep-report.sarif")
}
