package zip

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

func EmptyZip() *ZipFile {
	return &ZipFile{}
}

// Steps to Zip a File
// 1. Create the .zip file you want to write to
// 2. Create the writer to the .zip file
// 3. In the zip writer, make a virtual file
// 4. In the virtual file, add the contents to the related physical file
// 4b. Close the virtual file writer
// 5. Repeat steps 3 and 4 for each file you want to zip
// 6. Close

// ~~~~~~~~~~~~~~~~~~~~ ZipFile
type ZipFile struct {
	writer zip.Writer
	name   string
	file   *os.File
}

func NewZipFile(name string) *ZipFile {
	name = constructZipName(name)

	file, err := os.Create(name)
	handle(err)

	return &ZipFile{writer: *zip.NewWriter(file), name: name, file: file}
	// return &ZipFile{}
}

func constructZipName(name string) string {
	if search("zip", strings.Split(name, ".")) > -1 {
		return name
	}
	return format("%s.zip", name)
}

func (z *ZipFile) Name() string { return z.name }

func (z *ZipFile) AddZipFile(filename string, contents io.Reader) {
	fileWriter, err := z.writer.Create(filename)
	handle(err)

	_, err = io.Copy(fileWriter, contents)
	handle(err)
}

func (z *ZipFile) WriteAndClose() {
	z.writer.Close()
	z.file.Close()
}

func ZipFolderN(folderPath, outputName string) {
	z := NewZipFile(outputName)

	// Crawl through folder
	crawler := ZipCrawler{ZipFile: *z, path: folderPath}
	crawler.Crawl(folderPath)

	z.WriteAndClose()
}

type ZipCrawler struct {
	ZipFile
	path string
	root string
}

// HandleFile writes the file contents to the zip file in the location specified by the path.
func (c *ZipCrawler) HandleFile(path string) {
	file, err := files.OpenFile(path)

	helpers.Handle(err)

	c.AddZipFile(path, file)
}

// HandleFolder recursively call Crawl on the folder being zipped.
func (c *ZipCrawler) HandleFolder(path string) {
	originalPath := c.path

	dirContents, err := os.ReadDir(path)
	handle(err)

	for _, dirEntry := range dirContents {
		// for each dirEntry, handle if if it is a folder or file
		entryPath := path + dirEntry.Name()
		print(entryPath)

		// c.Crawl(entryPath)
	}

	c.path = originalPath
}

func (c *ZipCrawler) Crawl(path string) {
	// NOTE: I am trying to rewrite the ZipFolder function, and that uses the ZipCrawler
	// In order to do so, I need to work on this however this is in a temporary state haitus
	c.path = path

	// 1 check if is folder or directory
	isDir := true

	if isDir {
		c.HandleFolder(path)
	} else {
		c.HandleFile(path)
	}
}

func ZipFolder(path, output string) {
	// here we create the zip zipFile
	zipFile, err := os.Create(format("%s.zip", output))
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	// helpers.Print(file.Name())

	// here we create the zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// action performed at each file
	walker := func(path string, info os.FileInfo, err error) error {
		printf("Crawling: %v", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// here we open the fileToZip that we want to zip
		fileToZip, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		if path[0] == "/"[0] { // TODO not really proper, im just following the advice above and should do the job
			err = errors.New("not allowed to have an absolute path")
			return err
		}

		// HERE is the actual file processing, above is error checking

		// here, in the zip Writer, we create the virtual file fileBeingZipped
		fileBeingZipped, err := zipWriter.Create(path)
		if err != nil {
			return err
		}

		// here we copy the contents in the physical file to the virtual file being zipped
		_, err = io.Copy(fileBeingZipped, fileToZip)
		if err != nil {
			return err
		}

		return nil
	}

	// performs function `walker` on each file within path, recursively
	err = filepath.Walk(path, walker)
	if err != nil {
		panic(err)
	}
}
