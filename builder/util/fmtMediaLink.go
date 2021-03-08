package util

import (
	"io/ioutil"
	"log"
    "path/filepath"
	"os"
	"regexp"
    "net/url"
)

func fmtMediaLink(file string) {
    log.Println("fix medialink:"+file)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
    md := string(bytes)
    log.Println(file)

    reg := regexp.MustCompile(`\w{32}`)
    pageID := reg.FindString(file)

    reg = regexp.MustCompile(`\/(.*)\.md`)
    folderName := reg.FindStringSubmatch(file)[1]

    
	// regexp find path `!()[path]`
	reg = regexp.MustCompile(`\!\[.*\]\((.+)\)`)
    md = reg.ReplaceAllString(md,"![]($1)")
	for _, m := range reg.FindAllString(md, -1) {
        log.Println(m)
        removeReg := regexp.MustCompile(`\!\[.*\]\(.*\/([^\/]*)\)`)
        media := removeReg.FindStringSubmatch(m)[1]
        mediaName := pageID+"-" + media
        md = removeReg.ReplaceAllString(md,"![](../"+mediaName+")")

		extractMedia(media,mediaName, folderName)
	}
    reg = regexp.MustCompile(`\[(.+)\]\((.+)\)`)
    md = reg.ReplaceAllString(md,"![$1]($2)")

    log.Println(md)
	// fix media link to MD
	ioutil.WriteFile(file, []byte(md), os.ModePerm)
    extractMD(file)
}
func extractMedia(file string,newFile string,folder string) {
    file, _ = url.QueryUnescape(file)
	err := os.Rename(
		filepath.Join(contentPath, folder, file),
		filepath.Join(postPath, newFile),
	)
	if err != nil {
		panic(err)
	}
	log.Println(filepath.Join(postPath, newFile))
}

func extractMD(pathwithfile string) {
    _, file := filepath.Split(pathwithfile)
    err := os.Rename(
        filepath.Join(contentPath, file),
        filepath.Join(postPath, file),
    )
    if err != nil {
        panic(err)
    }
    log.Println(filepath.Join(postPath, file))
}
