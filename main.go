package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"log"
	"strconv"
	"path/filepath"
	"strings"
)

const (
	Separator     = os.PathSeparator
	strFullPipe   = "├───"
	strLastPipe   = "└───"
	strSimplePipe = "│"
)

type FileCust struct {
	basePath string
	isFile   bool
	intSize  int
}

type FileMap map[string]FileCust

var buildSubPipes []bool
var addPipe = true

func sortMapByKey(filesMap FileMap) []string {
	keys := make([]string, 0, len(filesMap))
	for k := range filesMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	var filesMap = make(map[string]FileCust)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if filepath.Base(path) != "." {
			if printFiles {
				filesMap[path] = FileCust{
					isFile:   !info.IsDir(),
					basePath: filepath.Base(path),
					intSize:  int(info.Size()),
				}
			} else if info.IsDir() {
				filesMap[path] = FileCust{
					isFile:   false,
					basePath: filepath.Base(path),
					intSize:  int(info.Size()),
				}
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	//SORTING Map by Key
	keys := sortMapByKey(filesMap)

	checkedRoutes := make(map[string]bool)

	//Add range of root directories
	for _, k := range keys {
		fl := strings.Split(k, string(Separator))
		if !filesMap[k].isFile {
			if _, exist := checkedRoutes[fl[0]]; !exist {
				checkedRoutes[fl[0]] = true
			}
		} else if filesMap[k].isFile {
			if _, exist := checkedRoutes[fl[0]]; !exist {
				checkedRoutes[fl[0]] = false
			}
		}
	}

	//Print resulted sorted map
	countFolders := 0
	filesInTheRoot := int(len(checkedRoutes))
	for _, k := range keys {

		f := strings.Split(k, string(Separator))
		if !filesMap[k].isFile {
			if val, exist := checkedRoutes[f[0]]; exist {
				countFolders++

				if val {
					delete(checkedRoutes, f[0])
					strPipe := generatePipeType(int(countFolders), filesInTheRoot)
					fmt.Println(strPipe + f[0])

					isPrintRootPipe := true
					if countFolders >= filesInTheRoot {
						isPrintRootPipe = false
					}
					printTreeRecursive(f[0], printFiles, true, 0, isPrintRootPipe, false)
				}

			}
		} else {

			if len(f) == 1 {
				countFolders++
				strPipe := generatePipeType(int(countFolders), filesInTheRoot)
				size := printSize(int(filesMap[k].intSize))
				fmt.Println(strPipe + f[0] + size)
			}
		}
	}
	return nil
}

func generatePipeType(countFolders, filesInTheRoot int) string {
	strPipe := strFullPipe

	if int(countFolders) == filesInTheRoot {
		strPipe = strLastPipe
	}

	return strPipe
}

func printSize(fileSize int) string {
	size := " (" + strconv.Itoa(fileSize) + "b)"
	if fileSize == 0 {
		size = " (empty)"
	}

	return size
}

func printTreeRecursive(path string, printFiles, isRoot bool, intCount int, isPrintRootPipe, lastFile bool) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// Add additional tab and pipe repeating
	intCount += 1

	if intCount-1 == 1 {
		buildSubPipes = make([]bool, 0)
	}

	intfilterFiles:=0
	for _, f := range files {
		if f.IsDir() {
			intfilterFiles++
		}
	}
	//Set filter counter in current folder
	countFile := 0

	for _, f := range files {


		addPipe = false
		intNumberOfFilesInDir:=int(len(files))
		if printFiles {
			countFile++
			//WHEN PRINT FILES AND DIRECTORIES
			if countFile == len(files) {
				addPipe = false
			} else if countFile < len(files) {
				addPipe = true
			}
			if len(files) == 1 && len(buildSubPipes) <= 1 {
				addPipe = true
			} else if len(files) > 1 {
				addPipe = false
			}

			if intCount >= 2 && !lastFile {
				addPipe = true
			}

			if int(countFile) == int(len(files)) {
				lastFile = true
			}
		}else{
			if f.IsDir(){
				countFile++
			}
			//WHEN PRINT ONLY DIRECTORIES
			if intCount >= 2 && !lastFile {
				addPipe = true
			}
			if int(countFile) == int(intfilterFiles) {
				lastFile = true
			}
			intNumberOfFilesInDir=int(intfilterFiles)
		}


		if !isRoot {

			if len(buildSubPipes) < intCount-1 {
				buildSubPipes = append(buildSubPipes, addPipe)
			} else if len(buildSubPipes) > intCount-1 {
				buildSubPipes = make([]bool, 0)
				buildSubPipes = append(buildSubPipes, addPipe)
			}

		} else {
			buildSubPipes = make([]bool, 0)
		}
		strPipeRes := buildSubPipesForInnerFiles(isPrintRootPipe, int(countFile), int(intNumberOfFilesInDir), buildSubPipes)
		if f.IsDir() {

			pathFull := string(path + string(Separator) + f.Name())
			if !isRoot {
				fmt.Println(strPipeRes + f.Name())
			} else {
				fmt.Println(strPipeRes + f.Name())
			}

			// Print recursively folder content
			printTreeRecursive(pathFull, printFiles, false, intCount, isPrintRootPipe, lastFile)
		} else {
			if printFiles {
				//Get file size in string format
				size := printSize(int(f.Size()))

				if isRoot && !f.IsDir() {
					fmt.Println(strPipeRes + f.Name() + size)
				} else {
					fmt.Println(strPipeRes + f.Name() + size)
				}
			}
		}

	}
}

func buildSubPipesForInnerFiles(isPrintRoot bool, countFile, intFiles int, buildSubPipes []bool) string {
	strPipe := generatePipeType(countFile, intFiles)
	var strPipeResult string

	if isPrintRoot {
		strPipeResult += strSimplePipe + "\t"
	} else {
		strPipeResult += "\t"
	}


		for _, val := range buildSubPipes {
			if val {
				strPipeResult += strSimplePipe + "\t"
			} else {
				strPipeResult += "\t"
			}
		}

	strPipeResult += strPipe

	return strPipeResult
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}

}
