package main

type Component interface {
	search(name string)
}

func main() {

	// aggregation relationship between folder and file
	// folder is the aggretate

	file1 := &File{name: "file 1"}
	file2 := &File{name: "file 2"}
	file3 := &File{name: "file 3"}

	folder1 := &Folder{name: "folder 1"}
	folder1.add(file1)

	folder2 := &Folder{name: "folder 2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("skyes")
}
