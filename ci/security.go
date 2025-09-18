package ci

import (
	"context"

	"github.com/munichbughunter/central-devsecops-pipeline/internal/dagger"
)

// RunImageScan führt einen Trivy Vulnerability Scan auf einem Container Image durch
func RunImageScan(ctx context.Context, dag *dagger.Client, imageRef string) (string, error) {
	return dag.
		Trivy().
		ScanImage(ctx, imageRef)

	// // Trivy Module aus dem Daggerverse laden
	// trivy := dag.Module("github.com/jpadams/daggerverse/trivy@44c178290a412483e785a436aabc46c707842a62")

	// // Container Image scannen
	// result := trivy.Call("ScanImage", dagger.ModuleCallOpts{
	// 	Args: []dagger.CallArgument{
	// 		{Name: "imageRef", Value: imageRef},
	// 		{Name: "format", Value: "table"}, // Optionen: table, json, sarif
	// 		{Name: "severity", Value: "HIGH,CRITICAL"}, // Nur High und Critical Vulnerabilities
	// 	},
	// })

	// return result.Stdout(ctx)
}
