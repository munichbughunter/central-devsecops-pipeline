// A generated module for CentralDevsecopsPipeline functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"

	"github.com/munichbughunter/central-devsecops-pipeline/internal/dagger"

	"github.com/munichbughunter/central-devsecops-pipeline/ci"
)

type CentralDevsecopsPipeline struct{}

func (m *CentralDevsecopsPipeline) BuildPython(
	ctx context.Context,
	src *dagger.Directory,
	pythonVersion string,
) (string, error) {
	client := dagger.Connect()
	if pythonVersion == "" {
		pythonVersion = "3.11"
	}
	return ci.RunPythonBuild(ctx, client, src, pythonVersion)
}

// Build and publish the default Python image via the central pipeline
func (m *CentralDevsecopsPipeline) BuildDefaultPythonImageAndPublish(
	ctx context.Context,
	imageTag string,
	githubUsername string,
	githubToken *dagger.Secret,
) (string, error) {
	// Verbindung zu Dagger aufbauen
	client := dagger.Connect()

	// Aufruf der CI-Funktion
	return ci.BuildDefaultPythonImageAndPublish(ctx, client, imageTag, githubUsername, githubToken)
}

// ScanContainerImage scannt ein Container Image auf Vulnerabilities
func (m *CentralDevsecopsPipeline) ScanImage(
	ctx context.Context,
	imageRef string,
) (string, error) {
	client := dagger.Connect()
	return ci.RunImageScan(ctx, client, imageRef)
}
