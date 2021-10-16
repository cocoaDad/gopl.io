// 本程序相对于du3的区别在于，实现了为每一个指定的root目录计算和定期输出各自占用的总空间

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type item struct {
	size  int64
	count int
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan item)

	var n sync.WaitGroup
	var nFiles, nBytes int64
	var pFiles, pBytes []int64
	for count, root := range roots {
		n.Add(1)
		pFiles = append(pFiles, 0)
		pBytes = append(pBytes, 0)
		go func(count int, dir string) {
			WalkDir(count, dir, &n, fileSizes)
		}(count, root)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *vFlag {
		fmt.Println("vFlag on")
		tick = time.Tick(50 * time.Millisecond)
	}
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			// fmt.Println("<-fileSizes")
			if !ok {
				break loop
			}
			nFiles++
			nBytes += size.size
			pFiles[size.count]++
			pBytes[size.count] += size.size
		case <-tick:
			printDiskUsage("whole: ", nFiles, nBytes)
			for c := range roots {
				printDiskUsage(roots[c], pFiles[c], pBytes[c]) 
			}
		}
	}
	printDiskUsage("whole: ", nFiles, nBytes)
	for c := range roots {
printDiskUsage(roots[c], pFiles[c], pBytes[c])
	}
}

func WalkDir(count int, dir string, wg *sync.WaitGroup, fileSizes chan item) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go WalkDir(count, subDir, wg, fileSizes)
		} else {
			fileSizes <- item{entry.Size(), count}
		}
	}
}

func printDiskUsage(dir string, nfiles, nbytes int64) {
	fmt.Printf("%s %d files  %.1f GB\n", dir, nfiles, float64(nbytes)/1e9)
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
