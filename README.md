# go-plotter 

- annotation and genome annotation visualizer.
- takes a gff file and visualizes all the regions such as mRNA, cds, protein five prime UTR, and three prime UTR. 
- uses gonum for the visualization. 

```
git clone htts.github.com/go-goannotate
go run main.go

```
- detail usage 

```
╰─$ go run main.go -h
annotate and visualize your genome

Usage:
  golanannotate [flags]

Flags:
  -A, --annotationfile string   genome annotation (default "path to the annotation file")
  -h, --help                    help for golanannotate
exit status 1
```
- it will produce all the genome annotation figures 

![](https://github.com/codecreatede/go-plotter/blob/main/barcds.png)

Gaurav Sablok
