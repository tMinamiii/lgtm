package drawer

type Drawer interface {
	Draw(inputPath, outputPath string) error
}
