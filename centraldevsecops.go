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

// func (m *CentralDevSecOpsPipeline) Build(ctx context.Context, src *dagger.Directory) (string, error) {
// 	return ci.RunBuild(ctx, src)
// }

// func (m *CentralDevSecOpsPipeline) Test(ctx context.Context, src *dagger.Directory) (string, error) {
// 	return pipeline.RunTests(ctx, src)
// }

func (m *CentralDevsecopsPipeline) RunSAST(ctx context.Context, src *dagger.Directory) (string, error) {
	// Dagger Client erzeugen
	client := dagger.Connect()

	return ci.RunSAST(ctx, client, src)
}
