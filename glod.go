/*
Being inspired by youtube-dl, Glod is a small library to help retrieve direct URL from multiple media sources that written in Go.

The package is named after Gloddson a.k.a Glod Gloddson, one of those unconventional Dwarfs, who with Giamo Casanunda, Hwel the playwright, Mad, Sharn and Cheery Littlebottom, express rebellion against the limited life of a conventional Dwarf. Glod is a musician, a horn player, who comes to Ankh-Morpork apparently after some years' experience elsewhere, since he is first encountered applying for membership in the Musicians' Guild. In the office he meets Imp y Celyn and Lias Bluestone, two more newcomers. They join up to try to raise the Guild membership fees and form The Band With Rocks In.

If you are building your music streaming website, the command line tool to download media files or even a media crawler, etc., you do not need to build it from scratch, glod is here to help. Package glod provides client library for other developer that wants to develop media related apps.

Glod includes several packages which are the source of media website, such as YouTube, SoundCloud, NhacCuaTui, Zing Mp3 ... Check out the sub-packages for more.
*/
package glod

// Source is the main interface and it defines the main methods of the package
type Source interface {
	GetDirectLink(url string) ([]Response, error)
}

type Response struct {
	Artist    string
	StreamURL string
	Title     string
}
