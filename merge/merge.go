// Package merge implements PDF merge.
package merge

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func mergePdfs(readers []io.ReadSeeker, writer io.Writer, dividerPage bool) error {
	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	return api.MergeRaw(readers, writer, dividerPage, conf)
}
