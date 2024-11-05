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
	"github.com/pchchv/bpdf/node"
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

// GetCurrentConfig is responsible for returning the current settings from the file.
func (m *Bpdf) GetCurrentConfig() *entity.Config {
	return m.config
}

// GetStructure is responsible for return the component tree,
// this is useful on unit tests cases.
func (m *Bpdf) GetStructure() *node.Node[core.Structure] {
	m.fillPageToAddNew()
	node := node.New(core.Structure{
		Type:    "maroto",
		Details: m.config.ToMap(),
	})
	for _, p := range m.pages {
		inner := p.GetStructure()
		node.AddNext(inner)
	}
	return node
}

func (m *Bpdf) processPage(pages []core.Page) ([]byte, error) {
	innerCtx := m.cell.Copy()
	innerProvider := getProvider(cache.NewMutexDecorator(cache.New()), m.config)
	for _, page := range pages {
		page.Render(innerProvider, innerCtx)
	}

	return innerProvider.GenerateBytes()
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

func (m *Bpdf) addHeader() {
	for _, headerRow := range m.header {
		m.currentHeight += headerRow.GetHeight(m.provider, &m.cell)
		m.rows = append(m.rows, headerRow)
	}
}

func (m *Bpdf) addRow(r core.Row) {
	if len(r.GetColumns()) == 0 {
		r.Add(col.New())
	}

	maxHeight := m.cell.Height
	r.SetConfig(m.config)
	rowHeight := r.GetHeight(m.provider, &m.cell)
	sumHeight := rowHeight + m.currentHeight + m.footerHeight
	// Row smaller than the remain space on page
	if sumHeight < maxHeight {
		m.currentHeight += rowHeight
		m.rows = append(m.rows, r)
		return
	}

	// As row will extrapolate page, we will add empty space
	// on the page to force a new page
	m.fillPageToAddNew()
	m.addHeader()
	// AddRows row on the new page
	m.currentHeight += rowHeight
	m.rows = append(m.rows, r)
}

func (m *Bpdf) addRows(rows ...core.Row) {
	for _, row := range rows {
		m.addRow(row)
	}
}

func (m *Bpdf) setConfig() {
	for i, page := range m.pages {
		page.SetConfig(m.config)
		page.SetNumber(i+1, len(m.pages))
	}
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
