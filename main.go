package linkchecker

import (
	"flag"
	"os"
	"strings"
)

var filePath string

func init() {
	flag.StringVar(&filePath, "file", "", "File to check")
	flag.Parse()
}

func main() {
	if filePath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

}

func getFilesLinks(filePath string) []string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	reader := strings.NewReader(string(content))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)
	}
}
