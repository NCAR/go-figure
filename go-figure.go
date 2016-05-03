package gofigure

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sort"
)

/*gather returns a list files located in the passed paths.  Files are parsed in
ascii order, and in linear order from the passed paths. */
func gather(paths []string) (files sort.StringSlice) {
	files = sort.StringSlice{}
	for _, pa := range paths {
		pa = os.ExpandEnv(pa)
		lfiles := sort.StringSlice{}
		walk := func(p string, i os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !i.IsDir() {
				lfiles = append(lfiles, p)
			} else if pa != p {
				return filepath.SkipDir
			}
			return nil
		}
		filepath.Walk(pa, walk)
		for _, f := range lfiles {
			files = append(files, f)
		}
	}
	return
}

/*Parse populates the viper instance walking through the passed paths, loading any
files found in the directory.  It silently ignores any errors (bad files, etc)*/
func Parse(v *viper.Viper, paths []string) {
	for _, file := range gather(paths) {
		if f, e := os.Open(file); e == nil {
			v.MergeConfig(f)
			defer f.Close()
		}
	}
	return
}
