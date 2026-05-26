package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strings"
)

var outputDir = flag.String("output-dir", "manifest_staging/charts/gatekeeper", "The root directory in which to write the Helm chart")

var kindRegex = regexp.MustCompile(`(?m)^kind:[\s]+([\S]+)[\s]*$`)

// use exactly two spaces to be sure we are capturing metadata.name.
var nameRegex = regexp.MustCompile(`(?m)^  name:[\s]+([\S]+)[\s]*$`)

const (
	DeploymentKind     = "Deployment"
	ServiceAccountKind = "ServiceAccount"
	end                = "{{- end }}"
)

func isRbacKind(str string) bool { _ = "STUB: not implemented"; return false }

func extractKind(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func extractName(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func extractCRDKind(obj string) (string, error) { _ = "STUB: not implemented"; return "", nil }

type kindSet struct {
	byKind map[string][]string
}

func (ks *kindSet) Add(obj string) error { _ = "STUB: not implemented"; return nil }

func (ks *kindSet) Write() error { _ = "STUB: not implemented"; return nil }

// Inject extra volume mounts at the start of the volumeMounts section for stability

// Inject extra volumes at the start of the volumes section for stability

// Inject export-related volume mount and possible export sidecar

// Inject extra mounts/volumes at the headers for stability

func doReplacements(obj string) string { _ = "STUB: not implemented"; return "" }

func copyStaticFiles(root string, subdirs ...string) error { _ = "STUB: not implemented"; return nil }

// #nosec G304

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	kinds := kindSet{byKind: make(map[string][]string)}
	b := strings.Builder{}
	notate := func() {
		obj := doReplacements(b.String())
		b.Reset()
		if err := kinds.Add(obj); err != nil {
			log.Fatalf("Error adding object: %s, %s", err, b.String())
		}
	}

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "---") {
			if b.Len() > 0 {
				notate()
			}
		} else {
			b.WriteString(scanner.Text())
			b.WriteString("\n")
		}
	}
	if b.Len() > 0 {
		notate()
	}
	if err := copyStaticFiles("cmd/build/helmify/static"); err != nil {
		log.Fatal(err)
	}
	if err := kinds.Write(); err != nil {
		log.Fatal(err)
	}
}
