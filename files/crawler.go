package fileIO

import "os"

type Crawler interface {
	HandleFile(filename os.DirEntry)
	HandleFolder(path os.DirEntry)
	Crawl(path string)
}
