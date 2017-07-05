package youtube

import (
	"git.dwarvesf.com/glod"
)

type Youtube struct {
}

// function that receive input is a link and output doesnt matter(but it override GetDirectLink of Glod interface)
func (youtube *Youtube) GetDirectLink(link string) ([]glod.Response, error) {
	var listSong []glod.Response
	return listSong, nil
}
