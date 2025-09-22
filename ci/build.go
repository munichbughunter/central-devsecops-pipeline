package ci

import (
	"context"
	"fmt"

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

// Build and publish the default Python image in one step
func BuildDefaultPythonImageAndPublish(ctx context.Context, client *dagger.Client, imageTag string, githubUsername string, githubToken *dagger.Secret) (string, error) {
	image := client.Container().
		From("python:3.12-slim").
		WithExec([]string{"pip", "install", "--no-cache-dir", "bandit", "black", "flake8"}).
		WithLabel("maintainer", githubUsername).
		WithUser("appuser").
		WithWorkdir("/app")

	// Trivy Scan direct on the container before publishing		
	out, err := client.Trivy().ScanContainer(ctx, image, dagger.TrivyScanContainerOpts{
		Severity:      "HIGH,CRITICAL,MEDIUM",
		ExitCode:      1,
		Format:        "table",
		TrivyImageTag: "latest",
	})
	if err != nil {
		return "", err
	}
	
	fmt.Printf("trivy scan result:\n%s\n", out)
	
	// Reuse the generic publish function
	return publishImage(ctx, client, image, imageTag, githubUsername, githubToken, "default-python")
}

// Publish any container image to a registry
func publishImage(ctx context.Context, client *dagger.Client, image *dagger.Container, imageTag, githubUsername string, githubToken *dagger.Secret, repoName string) (string, error) {

	return image.
		WithRegistryAuth("ghcr.io", githubUsername, githubToken).
		Publish(ctx, fmt.Sprintf("ghcr.io/%s/%s:%s", githubUsername, repoName, imageTag))
}
