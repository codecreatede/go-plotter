# go-plotter 

- annotation and genome annotation visualizer.
- takes a gff file and visualizes all the regions such as mRNA, cds, protein five prime UTR, and three prime UTR. 
- uses gonum for the visualization. 

```
git clone htts.github.com/go-plotter
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

```
-rw-r--r--. 1 gauavsablok gauavsablok 16751 Nov  1 19:27 barcdsMinus.png
-rw-r--r--. 1 gauavsablok gauavsablok 17014 Nov  1 19:27 barcdsPlus.png
-rw-r--r--. 1 gauavsablok gauavsablok 16751 Nov  1 19:27 barcds.png
-rw-r--r--. 1 gauavsablok gauavsablok 18691 Nov  1 19:27 barexonMinus.png
-rw-r--r--. 1 gauavsablok gauavsablok 18152 Nov  1 19:27 barexonPlus.png
-rw-r--r--. 1 gauavsablok gauavsablok 18691 Nov  1 19:27 barexon.png
-rw-r--r--. 1 gauavsablok gauavsablok 14515 Nov  1 19:27 barfiveMinus.png
-rw-r--r--. 1 gauavsablok gauavsablok 15755 Nov  1 19:27 barfivePlus.png
-rw-r--r--. 1 gauavsablok gauavsablok 14515 Nov  1 19:27 barfive.png
-rw-r--r--. 1 gauavsablok gauavsablok 14968 Nov  1 19:27 barmRNAMinus.png
-rw-r--r--. 1 gauavsablok gauavsablok 15935 Nov  1 19:27 barmRNAPlus.png
-rw-r--r--. 1 gauavsablok gauavsablok 14968 Nov  1 19:27 barmRNA.png
-rw-r--r--. 1 gauavsablok gauavsablok 16890 Nov  1 19:27 barproteinMinus.png
-rw-r--r--. 1 gauavsablok gauavsablok 17917 Nov  1 19:27 barproteinPlus.png
-rw-r--r--. 1 gauavsablok gauavsablok 16890 Nov  1 19:27 barprotein.png
-rw-r--r--. 1 gauavsablok gauavsablok 13854 Nov  1 19:27 barthreeMinus.png
-rw-r--r--. 1 gauavsablok gauavsablok 14599 Nov  1 19:27 barthreePlus.png
-rw-r--r--. 1 gauavsablok gauavsablok 13854 Nov  1 19:27 barthree.png
```

![](https://github.com/codecreatede/go-plotter/blob/main/barcds.png)

Gaurav Sablok
