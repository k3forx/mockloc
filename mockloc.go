package mockloc

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/packages"
)

const doc = "mockloc is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "mockloc",
	Doc:  doc,
	Run:  run,
}

func run(_ *analysis.Pass) (any, error) {
	cfg := packages.Config{
		Mode: packages.NeedName | packages.NeedFiles,
		Dir:  "./testdata/src/a",
	}

	pkgs, err := packages.Load(&cfg, "./...")
	if err != nil {
		return nil, err
	}

	var errs []error
	for _, pkg := range pkgs {
		pkgPath := pkg.PkgPath
		if strings.Contains(pkgPath, "mock") {
			continue
		}
		for _, goFile := range pkg.GoFiles {
			split := strings.Split(goFile, "/")
			f := split[len(split)-1]
			fmt.Printf("f: %+v\n", f)
			if strings.Contains(f, "mock") {
				errs = append(errs, fmt.Errorf("file: %s should not be directory of %s", f, pkgPath))
			}
		}
	}

	return nil, errors.Join(errs...)
}
