package main

import (
	"context"
	"os"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-go/logger"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	apiKey, exist := os.LookupEnv("STEAM_API_KEY")
	if !exist || apiKey == "" {
		panic("STEAM_API_KEY is required")
	}
	config := &porter.GetPorterInformationResponse{
		BinarySummary: &librarian.PorterBinarySummary{
			SourceCodeAddress: "github.com/tuihub/librarian/pkg/tuihub-steam",
			BuildVersion:      version,
			BuildDate:         "",
			Name:              "tuihub-steam",
			Version:           version,
			Description:       "",
		},
		GlobalName: "github.com/tuihub/librarian/pkg/tuihub-steam",
		Region:     "",
		FeatureSummary: &librarian.FeatureSummary{
			AccountPlatforms: []*librarian.FeatureFlag{
				{
					Id: tuihub.WellKnownToString(
						librarian.WellKnownAccountPlatform_WELL_KNOWN_ACCOUNT_PLATFORM_STEAM,
					),
					Name:             "Steam",
					Description:      "",
					ConfigJsonSchema: "",
					RequireContext:   false,
				},
			},
			AppInfoSources: []*librarian.FeatureFlag{
				{
					Id: tuihub.WellKnownToString(
						librarian.WellKnownAppInfoSource_WELL_KNOWN_APP_INFO_SOURCE_STEAM,
					),
					Name:             "Steam",
					Description:      "",
					ConfigJsonSchema: "",
					RequireContext:   false,
				},
			},
		},
	}
	server, err := tuihub.NewPorter(
		context.Background(),
		config,
		internal.NewHandler(apiKey),
	)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if err = server.Run(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
