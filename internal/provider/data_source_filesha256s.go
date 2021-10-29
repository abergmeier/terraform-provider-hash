package provider

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	homedir "github.com/mitchellh/go-homedir"
)

func dataSourceFileSha256s() *schema.Resource {
	return &schema.Resource{
		Description: "Hashes all files with consistent ordering",

		ReadContext: dataSourceFileSha256sRead,

		Schema: map[string]*schema.Schema{
			"files": {
				Type:        schema.TypeSet,
				Description: "Set of file paths.",
				Elem:        schema.TypeString,
				Required:    true,
			},
			"sha256": {
				Type:        schema.TypeString,
				Description: "The combined digest of all files.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFileSha256sRead(ctx context.Context, r *schema.ResourceData, d interface{}) diag.Diagnostics {
	fi := r.Get("files")
	files := fi.([]string)

	sort.Strings(files)

	h := sha256.New()
	for _, path := range files {
		err := func() error {
			f, err := openFile(".", path)
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(h, f)
			return err
		}()
		if err != nil {
			return diag.FromErr(err)
		}
	}
	err := r.Set("sha256", hex.EncodeToString(h.Sum(nil)))
	return diag.FromErr(err)
}

func openFile(baseDir, path string) (*os.File, error) {
	path, err := homedir.Expand(path)
	if err != nil {
		return nil, fmt.Errorf("failed to expand ~: %s", err)
	}

	if !filepath.IsAbs(path) {
		path = filepath.Join(baseDir, path)
	}

	// Ensure that the path is canonical for the host OS
	path = filepath.Clean(path)

	return os.Open(path)
}
