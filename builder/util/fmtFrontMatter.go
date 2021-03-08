package util

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

    "github.com/gosimple/slug"
	"github.com/kjk/notionapi"
)

const descLen = 120

type frontMatter struct {
	PageID      string
	Title       string
	Time        int64
}

func (fm *frontMatter) Check(block *notionapi.Block) {
	switch block.Type {
	case notionapi.BlockPage:
		// frontMatter.Date
		fm.Time = block.CreatedTime
		// frontMatter.Title
		fm.Title = block.Title
	}
}

func (fm frontMatter) Print() {
	log.Println(fm.PageID)
	log.Println(fm.Title)
	log.Println(fm.Date())
}

func (fm frontMatter) Date() string {
	return time.Unix(fm.Time/1000, 0).Format("2006-01-02")
}

func (fm frontMatter) GetFile() string {
	var ret string
    log.Println(fm.PageID)
	err := filepath.Walk(contentPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, fm.PageID) {
			ret = path
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return ret
}

func (fm frontMatter) AddFrontMatter(in string) string {
	md := in

    // remove first H1
    reg := regexp.MustCompile(`^# (.+)\n\n`)

    log.Println(md)
    //get the post title
    orginTitle := reg.FindStringSubmatch(md)[0]
    titleName := reg.FindStringSubmatch(md)[1]

    //remove first H1
    md = reg.ReplaceAllString(md,"")

    //remove tag,make tag to ["a","b"]
    reg = regexp.MustCompile(`tag: (.+,|.+)\n`)
    beforeTags := reg.FindStringSubmatch(md)[1]
    replacer := strings.NewReplacer(" ","")
    afterTags := replacer.Replace(beforeTags)
    slice := strings.Split(afterTags,",")
    temp := []string{}
    for _, value := range slice{
        temp = append(temp,"\"" + value + "\"")
    }
    tags := strings.Join(temp,",")
    tags = "[" + tags + "]"

    //post information date title slug tags
    str := "+++\n" + "date = \"" + time.Now().Format("2006-01-02") + "\"\n" + "title = \"" + titleName +"\"\n" + "slug = \"" +slug.Make(titleName) + "\"\n" + "tags = " + tags + "\n+++\n" + orginTitle
    md = reg.ReplaceAllString(md,str)

    reg = regexp.MustCompile(`Author:.+\n`)
	md = reg.ReplaceAllString(md, "")

    reg = regexp.MustCompile(`Finish: .+\n`)
    md = reg.ReplaceAllString(md,"")

    reg = regexp.MustCompile(`發布時間:.+`)
    md = reg.ReplaceAllString(md,"")

    reg = regexp.MustCompile(``)
	return md
}

func fmtFrontMatter(file string, fm frontMatter) {
    log.Println(file)
	mdBytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
    tempMD := bytes.NewBuffer(mdBytes)
	// add frontmatter
	doneMD := fm.AddFrontMatter(tempMD.String())
    log.Println("fmtFrontMatter:"+file)
	ioutil.WriteFile(file, []byte(doneMD), os.ModePerm)

	// mv post to folder
}
