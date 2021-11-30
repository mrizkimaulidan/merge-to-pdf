# merge-to-pdf

Mass merge image to pdf file.

Clone :

```bash
$ git clone https://github.com/mrizkimaulidan/merge-to-pdf.git
```

Download required depedencies :

```bash
$ go mod download
```

Build :

```bash
$ go build main.go
```

Usage :

```bash
$ ./main.exe --folder-path=/path/to/images/folder --output-path=/path/output.pdf
```
**Important** : Make sure the file name of the images are orderedly using numbers. For example :  1.jpg, 2.jpg, 3.jpg, etc..

Because if you are not rename the images file name orderedly by numbers, there will be an error in the output pdf file and the image of the file will be arranged randomly.

Example :

```bash
$ ./main.exe --folder-path=D:/images/ --output-path=D:/images/output.pdf
```

This tool was made from the beautiful [Unidoc](https://unidoc.io/) library. Thanks and love to **Unidoc**.

Unidoc Repository : [https://github.com/unidoc/unipdf](https://github.com/unidoc/unipdf)

Only tested on Windows, I don't know if it can run smoothly on Linux or MacOS. Just open an issue or pull requests if something went wrong on Linux or MacOS.