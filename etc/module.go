package etc


type Configuration struct {
	Name string
	Version string

	// 监听GPIO
	Listen listen

	Control control
}

type listen struct {
	Pin []int
	File string
	Interval int
}

type control struct {
	Pin []int
	File string
	Interval int
}
