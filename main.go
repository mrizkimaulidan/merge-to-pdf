package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
)

func init() {
	// bcdca97974e853709f5b4967f864dce23091a3f54f27e17262f997512f812d42
	err := license.SetMeteredKey("bcdca97974e853709f5b4967f864dce23091a3f54f27e17262f997512f812d42")

	if err != nil {
		panic(err)
	}
}

func MergeToPdf(folderPath *string, imagePaths *[]string, imageExtension *string, outputPath *string) error {
	c := creator.New()

	for _, imgPath := range *imagePaths {
		fullPath := *folderPath + imgPath + "." + *imageExtension

		img, err := c.NewImageFromFile(fullPath)

		if err != nil {
			panic(err)
		}

		img.ScaleToWidth(612.0)

		// Use page width of 612 points, and calculate the height proportionally based on the image.
		// Standard PPI is 72 points per inch, thus a width of 8.5"
		height := 612.0 * img.Height() / img.Width()
		c.SetPageSize(creator.PageSize{612, height})
		c.NewPage()
		img.SetPos(0, 0)
		_ = c.Draw(img)
	}

	err := c.WriteToFile(*outputPath)

	return err
}

func FileName(files *[]fs.FileInfo) ([]string, string) {
	var sliceOfFileNameOnlyName []string
	var imageExtension string

	for _, file := range *files {
		fileName := file.Name()
		splittedOfFileName := strings.Split(fileName, ".")

		sliceOfFileNameOnlyName = append(sliceOfFileNameOnlyName, splittedOfFileName[0])
		imageExtension = splittedOfFileName[1]

		sort.Slice(sliceOfFileNameOnlyName, func(i, j int) bool {
			numA, _ := strconv.Atoi(sliceOfFileNameOnlyName[i])
			numB, _ := strconv.Atoi(sliceOfFileNameOnlyName[j])
			return numA < numB
		})
	}

	return sliceOfFileNameOnlyName, imageExtension
}

func main() {
	folderPath := flag.String("folder-path", "null", "Input your folder path here")
	outputPath := flag.String("output-path", "null", "Input your output path here")
	flag.Parse()

	files, _ := ioutil.ReadDir(*folderPath)

	filePath, imageExtension := FileName(&files)

	err := MergeToPdf(folderPath, &filePath, &imageExtension, outputPath)

	if err != nil {
		panic(err)
	}

	fmt.Println("Output at :", *outputPath)
}
