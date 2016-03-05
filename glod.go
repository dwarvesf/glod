package glod

type Glod interface {
	GetDirectLink(url string) ([]string, error)
}
