The primary way this project gets used is as a library. The main entry point is
the exported global variable `Analyzer`. All the core logic goes in source files
in the root of the repository.

The tool can be used as a binary but it is not the primary use case. The binary
is in `cmd/intrange` and gets called via `go vet`.

Third-party libraries are avoided unless absolutely necessary. All testing is
done using only the standard libraryWith with only one exception for
`TestAnalyzer` which uses `github.com/gostaticanalysis/testutil` and
`golang.org/x/tools/go/analysis/analysistest`.

Performance is a primary concern. The tool is designed to be fast and efficient.
