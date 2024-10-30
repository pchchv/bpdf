
func TestNewBorderColorStyler(t *testing.T) {
	sut := cellwriter.NewBorderColorStyler(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderColorStyler", fmt.Sprintf("%T", sut))
}
