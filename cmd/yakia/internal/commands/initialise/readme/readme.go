package readme

type Readme struct {
	Name        string
	ReadmeItems []ReadmeItem
}

type ReadmeItem struct {
	Name        string
	Description string
	Usage       string
}

func New() *Readme {

	return &Readme{
		Name: "项目名称",
	}

}
