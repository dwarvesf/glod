glod - download songs, videos(single or playlist) from sources: zing, nhaccuatui, youtube

## Introduction: glod which is abbreviated of Gloddson a.k.a Glod Gloddson, one of those unconventional Dwarfs, who with Giamo Casanunda, Hwel the playwright, Mad, Sharn and Cheery Littlebottom, express rebellion against the limited life of a conventional Dwarf. Glod is a musician, a horn player, who comes to Ankh-Morpork apparently after some years' experience elsewhere, since he is first encountered applying for membership in the Musicians' Guild. In the office he meets Imp y Celyn and Lias Bluestone, two more newcomers. They join up to try to raise the Guild membership fees and form The Band With Rocks In.

## Installation: go get github.com/dwarvesf/glod

## Usage:
 ```go 

	const (
		initNhacCuatui string = "nhaccuatui"
	)
	if strings.Contains(url, initNhacCuatui) {
		glod = &nct.NhacCuaTui{}
	} 
	listStream, error := glod.GetDirectLink(url) //url is the link inputed, listStream is list of url that permanently downloadable link
	//in case url is song's link, listStream contains one item
 ```
## Contributing
  * Fork it!*
  * Create your feature branch(for example soundcloud): git checkout -b feature/soundcloud*
  * Write your function download, remember to override GetDirectLink function*
  * Commit your changes: git commit -am "Add function download for soundcloud"*
  * Push to the branch: git push origin feature/soundcloud*
  * Submit a pull request*
