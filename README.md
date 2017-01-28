[![Coverage Status](https://coveralls.io/repos/github/dwarvesf/glod/badge.svg?branch=develop)](https://coveralls.io/github/dwarvesf/glod?branch=develop)

[![GoDoc](https://godoc.org/github.com/dwarvesf/glod?status.svg)](https://godoc.org/github.com/dwarvesf/glod)
[![Build Status](https://travis-ci.org/dwarvesf/glod.svg?branch=master)](https://travis-ci.org/dwarvesf/glod)
[![Code Climate](https://codeclimate.com/github/dwarvesf/glod/badges/gpa.svg)](https://codeclimate.com/github/dwarvesf/glod)
[![Test Coverage](https://codeclimate.com/github/dwarvesf/glod/badges/coverage.svg)](https://codeclimate.com/github/dwarvesf/glod/cov000erage)

# Introduction:

Being inspired by [youtube-dl](https://github.com/rg3/youtube-dl), **Glod** is a small library to help retrieve direct URL from multiple media sources that written in Go.

Glod which is abbreviated of Gloddson a.k.a Glod Gloddson, one of those unconventional Dwarfs, who with Giamo Casanunda, Hwel the playwright, Mad, Sharn and Cheery Littlebottom, express rebellion against the	limited life of a conventional Dwarf. Glod is a musician, a horn player, who comes to Ankh-Morpork apparently after some years' experience elsewhere, since he is first encountered applying for membership in the Musicians' Guild. In the office he meets Imp y Celyn and Lias Bluestone, two more newcomers. They join up to try to raise the Guild membership fees and form The Band With Rocks In.

If you are building your music streaming website, the command line tool to download media files or even a crawler ... you do not have to build it from scratch, glod is here to help.

# Installation:

Assume that you had Go installed
  
```
$ go get -u github.com/dwarvesf/glod
```

# Usage:

``` go
const (
	initNhacCuatui string = "nhaccuatui"
)

if strings.Contains(url, initNhacCuatui) {
	glod = &nct.NhacCuaTui{}
}

// url is the link inputed, listStream is list of URLs that permanently downloadable link
// In case url is song's link, listStream contains one item
listStream, error := glod.GetDirectLink(url)
```

## Current status and TODO

* [x] Mp3 Zing
* [x] Nhaccuatui
* [x] Soundcloud
* [x] Youtube
* [x] Chiasenhac
* [x] Facebook
* [x] Vimeo
* [ ] Lynda
* [ ] Udemy
* [ ] Flickr
* [ ] Slideshare
* [ ] Dropbox
* [ ] PornHub :sunglasses:

# Implementation

We have created an command-line tool that implements glod. Check it out [glod-cli](https://github.com/dwarvesf/glod-cli)!

# Contributing

* Fork it!
* Create your feature branch (for example soundcloud):

```
$ git checkout -b feature/soundcloud
```

* Write your function download, remember to override **GetDirectLink()** function
* Commit your changes:

```
$ git commit -am "Add function download for soundcloud"
```

* Push to the branch:

```
$ git push origin feature/soundcloud
```

* Submit your pull request

# License

Copyright 2016 Dwarves Foundation

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
