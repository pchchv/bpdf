package bpdf

import (
	"github.com/f-amaral/go-async/async"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/page"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/config"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/cache"
	"github.com/pchchv/bpdf/internal/providers/gofpdf"
)

type Bpdf struct {
	config   *entity.Config
	provider core.Provider
	cache    cache.Cache
	// Building
	cell          entity.Cell
	pages         []core.Page
	rows          []core.Row
	header        []core.Row
	footer        []core.Row
	headerHeight  float64
	footerHeight  float64
	currentHeight float64
	// Processing
	pool async.Processor[[]core.Page, []byte]
}

func (m *Bpdf) fillPageToAddNew() {
	var p core.Page
	space := m.cell.Height - m.currentHeight - m.footerHeight
	c := col.New(m.config.MaxGridSize)
	spaceRow := row.New(space)
	spaceRow.Add(c)
	m.rows = append(m.rows, spaceRow)
	m.rows = append(m.rows, m.footer...)
	if m.config.PageNumber != nil {
		p = page.New(*m.config.PageNumber)
	} else {
		p = page.New()
	}

	p.SetConfig(m.config)
	p.Add(m.rows...)
	m.pages = append(m.pages, p)
	m.rows = nil
	m.currentHeight = 0
}

func getProvider(cache cache.Cache, cfg *entity.Config) core.Provider {
	deps := gofpdf.NewBuilder().Build(cfg, cache)
	provider := gofpdf.New(deps)
	provider.SetMetadata(cfg.Metadata)
	provider.SetCompression(cfg.Compression)
	provider.SetProtection(cfg.Protection)
	return provider
}

func getConfig(configs ...*entity.Config) *entity.Config {
	if len(configs) > 0 {
		return configs[0]
	}

	return config.NewBuilder().Build()
}

func (m *Bpdf) processPage(pages []core.Page) ([]byte, error) {
	innerCtx := m.cell.Copy()
	innerProvider := getProvider(cache.NewMutexDecorator(cache.New()), m.config)
	for _, page := range pages {
		page.Render(innerProvider, innerCtx)
	}

	return innerProvider.GenerateBytes()
}
