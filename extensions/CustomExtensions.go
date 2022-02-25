package extensions

type CustomExtensions string

const (
	TimeOut CustomExtensions = "x-timeout"
)

func (c CustomExtensions) String() string {
	return string(c)
}
