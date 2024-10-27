package designfilesystem

import "strings"

type FileSystem struct {
	validPaths map[string]int
}

type Node struct {
	path  string
	value int
}

func Constructor() FileSystem {
	fileSystem := FileSystem{
		validPaths: make(map[string]int),
	}
	fileSystem.validPaths["/"] = -1

	return fileSystem
}

func (this *FileSystem) CreatePath(path string, value int) bool {

	if path == "" || path == "/" {
		return false
	}
	if _, ok := this.validPaths[path]; ok {
		return false
	}
	splitPath := strings.Split(path, "/")

	pathOfParent := strings.Join(splitPath[:len(splitPath)-1], "/")

	if !strings.HasPrefix(pathOfParent, "/") {
		pathOfParent = "/" + pathOfParent
	}

	//parentNode, ok := this.validPaths[pathOfParent]
	_, ok := this.validPaths[pathOfParent]
	if !ok {
		return false
	}

	this.validPaths[path] = value

	return true
}

func (this *FileSystem) Get(path string) int {
	if path == "" || path == "/" {
		return -1
	}

	if currentPath, ok := this.validPaths[path]; ok {
		return currentPath
	}

	return -1
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */
