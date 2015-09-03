package glod

type Glod interface {
	GetDirectLink(link string) ([]string, error)
}
