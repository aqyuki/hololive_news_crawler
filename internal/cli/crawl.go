package cli

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aqyuki/hololive_news_crawler/internal/crawler"
	"github.com/aqyuki/hololive_news_crawler/internal/extractor/hololive"
	cfs "github.com/aqyuki/hololive_news_crawler/internal/fs"
	"github.com/aqyuki/hololive_news_crawler/pkg/logging"
	"github.com/spf13/cobra"
)

const (
	targetURL = "https://hololive.hololivepro.com/news"
)

func NewCrawlCommand(ctx context.Context) *cobra.Command {
	// outputPath is a flag to specify the output path
	var outputPath string

	// format is a flag to specify the output format
	var format string

	cmd := &cobra.Command{
		Use:   "crawl",
		Short: "Crawl the web site and save the data",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logging.FromContext(ctx)

			extractor := hololive.HololiveNewsExtractor{}
			urlExtractor := hololive.HololiveNewsURLExtractor{}

			c := crawler.NewCrawler(&extractor)
			c.SetURLExtractor(&urlExtractor)

			sites, err := c.Crawl(ctx, targetURL)
			if err != nil {
				logger.Error("failed to crawl", slog.Any("error", err))
				return err
			}

			if outputPath != "" {
				var encode cfs.Encode

				switch format {
				case "json":
					encode = cfs.EncodeJSON
				case "yaml":
					encode = cfs.EncodeYAML
				default:
					encode = cfs.EncodeJSON
				}

				if err := cfs.EncodeSave(outputPath, sites, encode); err != nil {
					logger.Error("failed to save", slog.Any("error", err))
					return err
				}
			} else {
				for _, site := range sites {
					fmt.Printf("Title : %s\t URL : %s\t Description : %s\t\n", site.Title, site.URL, site.Description)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&outputPath, "output", "o", "", "output path")
	cmd.Flags().StringVarP(&format, "format", "f", "json", "output format")

	return cmd
}
