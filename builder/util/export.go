package util

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
    "os"
    "regexp"

	"github.com/kjk/notionapi"
)

func NotionExport(client *notionapi.Client, pageID string) {
	createContentDir()
	zipByte, err := client.ExportPages(
		pageID,
		notionapi.ExportTypeMarkdown,
		false,
	)
	if err != nil {
		panic(err)
	}

	zipReader, err := zip.NewReader(
		bytes.NewReader(zipByte),
		int64(len(zipByte)),
	)
	if err != nil {
		panic(err)
	}

	for _, zipFile := range zipReader.File {
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			log.Println(err)
			continue
		}
        filePath := filepath.Join(contentPath, zipFile.Name)
        r := regexp.MustCompile(`\w{32}`)
        filePath = r.ReplaceAllString(filePath,pageID)
        os.MkdirAll(filepath.Dir(filePath),os.ModePerm)
		err = ioutil.WriteFile(
			filePath,
			unzippedFileBytes,
			0644,
		)
		if err != nil {
			log.Println(err)
			continue
		}
	}

}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
