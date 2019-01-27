package config

import (
	"fmt"
	"testing"
	"testing/quick"

	"github.com/nvkv/halp/pkg/testhelpers/v1"
)

func TestParser(t *testing.T) {
	checkParser := func(i int) bool {
		sheetid := testhelpers.RandomASCIIString()
		credspath := testhelpers.RandomASCIIString()
		tokenfile := testhelpers.RandomASCIIString()
		rangeStr := testhelpers.RandomASCIIString()
		cfg := Config{
			Datasource: DatasourceConfig{
				GoogleSheets: GSheetsConfig{
					SheetID:             sheetid,
					CredentialsFilePath: credspath,
					TokenFilePath:       tokenfile,
					Range:               rangeStr,
				},
			},
		}

		configStr := fmt.Sprintf(`
datasource {
  google_sheets {
    sheet_id         = "%s"
    credentials_file = "%s"
    token_file       = "%s"
    range            = "%s"
  }
}
`,
			sheetid,
			credspath,
			tokenfile,
			rangeStr,
		)

		parsedCfg, err := ParseConfig(configStr)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if *parsedCfg != cfg {
			fmt.Printf("Configs mismatch:\n %#v \n\n %#v\n", parsedCfg, cfg)
			return false
		}
		return true
	}
	testingConfig := &quick.Config{MaxCount: 10000}
	if err := quick.Check(checkParser, testingConfig); err != nil {
		t.Error(err)
	}
}
