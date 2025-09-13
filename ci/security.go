package ci

import (
	"context"

	"github.com/munichbughunter/central-devsecops-pipeline/internal/dagger"
)

func RunSAST(ctx context.Context, dag *dagger.Client, src *dagger.Directory) (string, error) {
	container := dag.Container().
		// Go 1.23 ist die neueste verfügbare stabile Version
		From("golang:1.23").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		// Git und curl installieren
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "git", "curl"}).
		// GOTOOLCHAIN auf auto setzen, damit Go 1.25 automatisch heruntergeladen wird
		WithEnvVariable("GOTOOLCHAIN", "auto").
		// Projekt-Dependencies auflösen
		WithExec([]string{"go", "mod", "download"}).
		// golangci-lint installieren (neueste Version)
		WithExec([]string{"bash", "-c", "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin"})

	// Nur spezifische Verzeichnisse scannen (Paket-basiert)
	return container.WithExec([]string{"/usr/local/bin/golangci-lint", "run", "./ci/...", "."}).Stdout(ctx)
}
