package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

/*
Author Gaurav Sablok
Universitat Potsdam
Date 2024-11-1

golang plotter for the whole genome annotations. Given a gff file, plots the strand
specific annotations for mRNA, cds, exons and others.

*/

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}

var annotationfile string

var rootCmd = &cobra.Command{
	Use:  "golanannotate",
	Long: "annotate and visualize your genome",
	Run:  annotateFunc,
}

func init() {
	rootCmd.Flags().
		StringVarP(&annotationfile, "annotationfile", "A", "path to the annotation file", "genome annotation")
}

func annotateFunc(cmd *cobra.Command, args []string) {
	type mRNADetails struct {
		mRNAParent string
		mRNAStrand string
		mRNAStart  int
		mRNAEnd    int
	}

	type exonDetails struct {
		exonParent string
		exonStrand string
		exonStart  int
		exonEnd    int
	}

	type proteinDetails struct {
		proteinParent string
		proteinStrand string
		proteinStart  int
		proteinEnd    int
	}

	type cdsDetails struct {
		cdsParent string
		cdsStrand string
		cdsStart  int
		cdsEnd    int
	}

	type fiveDetails struct {
		fiveParent string
		fiveStrand string
		fiveStart  int
		fiveEnd    int
	}

	type threeDetails struct {
		threeParent string
		threeStrand string
		threeStart  int
		threeEnd    int
	}

	mRNADet := []mRNADetails{}
	exonDet := []exonDetails{}
	cdsDet := []cdsDetails{}
	proteinDet := []proteinDetails{}
	threeDet := []threeDetails{}
	fiveDet := []fiveDetails{}

	annotateOpen, err := os.Open(annotationfile)
	if err != nil {
		log.Fatal(err)
	}
	annotateRead := bufio.NewScanner(annotateOpen)

	for annotateRead.Scan() {
		line := annotateRead.Text()
		if strings.Split(line, "\t")[2] == "mRNA" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			mRNADet = append(mRNADet, mRNADetails{
				mRNAParent: strings.Split(string(line), "\t")[2],
				mRNAStrand: strings.Split(string(line), "\t")[6],
				mRNAStart:  start,
				mRNAEnd:    end,
			})
		}
		if strings.Split(line, "\t")[2] == "exon" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			exonDet = append(exonDet, exonDetails{
				exonParent: strings.Split(string(line), "\t")[2],
				exonStrand: strings.Split(string(line), "\t")[6],
				exonStart:  start,
				exonEnd:    end,
			})
		}
		if strings.Split(line, "\t")[2] == "CDS" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			cdsDet = append(cdsDet, cdsDetails{
				cdsParent: strings.Split(string(line), "\t")[2],
				cdsStrand: strings.Split(string(line), "\t")[6],
				cdsStart:  start,
				cdsEnd:    end,
			})
		}
		if strings.Split(line, "\t")[2] == "protein" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			proteinDet = append(proteinDet, proteinDetails{
				proteinParent: strings.Split(string(line), "\t")[2],
				proteinStrand: strings.Split(string(line), "\t")[6],
				proteinStart:  start,
				proteinEnd:    end,
			})
		}
		if strings.Split(line, "\t")[2] == "five_prime_UTR" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			fiveDet = append(fiveDet, fiveDetails{
				fiveParent: strings.Split(string(line), "\t")[2],
				fiveStrand: strings.Split(string(line), "\t")[6],
				fiveStart:  start,
				fiveEnd:    end,
			})
		}

		if strings.Split(line, "\t")[2] == "three_prime_UTR" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			threeDet = append(threeDet, threeDetails{
				threeParent: strings.Split(string(line), "\t")[2],
				threeStrand: strings.Split(string(line), "\t")[6],
				threeStart:  start,
				threeEnd:    end,
			})
		}
	}

	exonLengthPlot := []int{}
	mRNALengthPlot := []int{}
	cdsLengthPlot := []int{}
	proteinLengthPlot := []int{}
	threeLengthPlot := []int{}
	fiveLengthPlot := []int{}

	for i := range exonDet {
		exonLengthPlot = append(exonLengthPlot, exonDet[i].exonEnd-exonDet[i].exonStart)
	}

	for i := range mRNADet {
		mRNALengthPlot = append(mRNALengthPlot, mRNADet[i].mRNAEnd-mRNADet[i].mRNAStart)
	}

	for i := range cdsDet {
		cdsLengthPlot = append(cdsLengthPlot, cdsDet[i].cdsEnd-cdsDet[i].cdsStart)
	}

	for i := range proteinDet {
		proteinLengthPlot = append(
			proteinLengthPlot,
			proteinDet[i].proteinEnd-proteinDet[i].proteinStart,
		)
	}

	for i := range threeDet {
		threeLengthPlot = append(threeLengthPlot, threeDet[i].threeEnd-threeDet[i].threeStart)
	}

	for i := range fiveDet {
		fiveLengthPlot = append(fiveLengthPlot, fiveDet[i].fiveEnd-fiveDet[i].fiveStart)
	}

	exonPlusLengthPlot := []int{}
	mRNAPlusLengthPlot := []int{}
	cdsPlusLengthPlot := []int{}
	proteinPlusLengthPlot := []int{}
	threePlusLengthPlot := []int{}
	fivePlusLengthPlot := []int{}

	exonMinusLengthPlot := []int{}
	mRNAMinusLengthPlot := []int{}
	cdsMinusLengthPlot := []int{}
	proteinMinusLengthPlot := []int{}
	threeMinusLengthPlot := []int{}
	fiveMinusLengthPlot := []int{}

	for i := range exonDet {
		if exonDet[i].exonStrand == "+" {
			exonPlusLengthPlot = append(exonPlusLengthPlot, exonDet[i].exonEnd-exonDet[i].exonStart)
		}
		if exonDet[i].exonStrand == "-" {
			exonMinusLengthPlot = append(
				exonMinusLengthPlot,
				exonDet[i].exonEnd-exonDet[i].exonStart,
			)
		}
	}

	for i := range mRNADet {
		if mRNADet[i].mRNAStrand == "+" {
			mRNAPlusLengthPlot = append(mRNAPlusLengthPlot, mRNADet[i].mRNAEnd-mRNADet[i].mRNAStart)
		}
		if mRNADet[i].mRNAStrand == "-" {
			mRNAMinusLengthPlot = append(
				mRNAMinusLengthPlot,
				mRNADet[i].mRNAEnd-mRNADet[i].mRNAStart,
			)
		}
	}

	for i := range cdsDet {
		if cdsDet[i].cdsStrand == "+" {
			cdsPlusLengthPlot = append(cdsPlusLengthPlot, cdsDet[i].cdsEnd-cdsDet[i].cdsStart)
		}
		if cdsDet[i].cdsStrand == "-" {
			cdsMinusLengthPlot = append(
				cdsMinusLengthPlot,
				cdsDet[i].cdsEnd-cdsDet[i].cdsStart,
			)
		}
	}

	for i := range fiveDet {
		if fiveDet[i].fiveStrand == "+" {
			fivePlusLengthPlot = append(fivePlusLengthPlot, fiveDet[i].fiveEnd-fiveDet[i].fiveStart)
		}
		if fiveDet[i].fiveStrand == "-" {
			fiveMinusLengthPlot = append(
				fiveMinusLengthPlot,
				fiveDet[i].fiveEnd-fiveDet[i].fiveStart,
			)
		}
	}

	for i := range threeDet {
		if threeDet[i].threeStrand == "+" {
			threePlusLengthPlot = append(
				threePlusLengthPlot,
				threeDet[i].threeEnd-threeDet[i].threeStart,
			)
		}
		if threeDet[i].threeStrand == "-" {
			threeMinusLengthPlot = append(
				threeMinusLengthPlot,
				threeDet[i].threeEnd-threeDet[i].threeStart,
			)
		}
	}

	for i := range proteinDet {
		if proteinDet[i].proteinStrand == "+" {
			proteinPlusLengthPlot = append(
				proteinPlusLengthPlot,
				proteinDet[i].proteinEnd-proteinDet[i].proteinStart,
			)
		}
		if proteinDet[i].proteinStrand == "-" {
			proteinMinusLengthPlot = append(
				proteinMinusLengthPlot,
				proteinDet[i].proteinEnd-proteinDet[i].proteinStart,
			)
		}
	}

	var mRNAPlotter plotter.Values
	var mRNAPlusPlotter plotter.Values
	var mRNAMinusPlotter plotter.Values

	for i := range mRNALengthPlot {
		mRNAPlotter = append(mRNAPlotter, float64(mRNALengthPlot[i]))
	}
	mRNA := plot.New()
	mRNA.Title.Text = "Bar chart"
	mRNA.Y.Label.Text = "mRNA"
	w := vg.Points(14)

	barsmRNA, err := plotter.NewBarChart(mRNAPlotter, w)
	if err != nil {
		panic(err)
	}

	barsmRNA.LineStyle.Width = vg.Length(0)
	barsmRNA.Color = plotutil.Color(0)
	barsmRNA.Offset = -w
	mRNA.Add(barsmRNA)
	mRNA.Legend.Add("mRNA", barsmRNA)
	if err := mRNA.Save(20*vg.Inch, 10*vg.Inch, "barmRNA.png"); err != nil {
		panic(err)
	}

	for i := range mRNAPlusLengthPlot {
		mRNAPlusPlotter = append(mRNAPlusPlotter, float64(mRNAPlusLengthPlot[i]))
	}
	mRNAPlus := plot.New()
	mRNAPlus.Title.Text = "Bar chart"
	mRNAPlus.Y.Label.Text = "mRNAPlus"

	barsmRNAPlus, err := plotter.NewBarChart(mRNAPlusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsmRNAPlus.LineStyle.Width = vg.Length(0)
	barsmRNAPlus.Color = plotutil.Color(0)
	barsmRNAPlus.Offset = -w
	mRNAPlus.Add(barsmRNAPlus)
	mRNAPlus.Legend.Add("mRNAPlus", barsmRNAPlus)
	if err := mRNAPlus.Save(20*vg.Inch, 10*vg.Inch, "barmRNAPlus.png"); err != nil {
		panic(err)
	}

	for i := range mRNAMinusLengthPlot {
		mRNAMinusPlotter = append(mRNAPlotter, float64(mRNAMinusLengthPlot[i]))
	}
	mRNAMinus := plot.New()
	mRNAMinus.Title.Text = "Bar chart"
	mRNAMinus.Y.Label.Text = "mRNA"

	barsmRNAMinus, err := plotter.NewBarChart(mRNAMinusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsmRNAMinus.LineStyle.Width = vg.Length(0)
	barsmRNAMinus.Color = plotutil.Color(0)
	barsmRNAMinus.Offset = -w
	mRNAMinus.Add(barsmRNA)
	mRNAMinus.Legend.Add("mRNAMinus", barsmRNAMinus)
	if err := mRNA.Save(20*vg.Inch, 10*vg.Inch, "barmRNAMinus.png"); err != nil {
		panic(err)
	}

	var cdsPlotter plotter.Values
	var cdsPlusPlotter plotter.Values
	var cdsMinusPlotter plotter.Values

	for i := range cdsLengthPlot {
		cdsPlotter = append(cdsPlotter, float64(cdsLengthPlot[i]))
	}
	cds := plot.New()
	cds.Title.Text = "Bar chart"
	cds.Y.Label.Text = "cds"

	barscds, err := plotter.NewBarChart(cdsPlotter, w)
	if err != nil {
		panic(err)
	}

	barscds.LineStyle.Width = vg.Length(0)
	barscds.Color = plotutil.Color(0)
	barscds.Offset = -w
	cds.Add(barscds)
	cds.Legend.Add("cds", barscds)
	if err := cds.Save(20*vg.Inch, 10*vg.Inch, "barcds.png"); err != nil {
		panic(err)
	}

	for i := range cdsPlusLengthPlot {
		cdsPlusPlotter = append(cdsPlusPlotter, float64(cdsPlusLengthPlot[i]))
	}
	cdsPlus := plot.New()
	cdsPlus.Title.Text = "Bar chart"
	cdsPlus.Y.Label.Text = "cdsPlus"

	barscdsPlus, err := plotter.NewBarChart(cdsPlusPlotter, w)
	if err != nil {
		panic(err)
	}

	barscdsPlus.LineStyle.Width = vg.Length(0)
	barscdsPlus.Color = plotutil.Color(0)
	barscdsPlus.Offset = -w
	cdsPlus.Add(barscdsPlus)
	cdsPlus.Legend.Add("cdsPlus", barscdsPlus)
	if err := cdsPlus.Save(20*vg.Inch, 10*vg.Inch, "barcdsPlus.png"); err != nil {
		panic(err)
	}

	for i := range cdsMinusLengthPlot {
		cdsMinusPlotter = append(cdsPlotter, float64(cdsMinusLengthPlot[i]))
	}
	cdsMinus := plot.New()
	cdsMinus.Title.Text = "Bar chart"
	cdsMinus.Y.Label.Text = "cds"

	barscdsMinus, err := plotter.NewBarChart(cdsMinusPlotter, w)
	if err != nil {
		panic(err)
	}

	barscdsMinus.LineStyle.Width = vg.Length(0)
	barscdsMinus.Color = plotutil.Color(0)
	barscdsMinus.Offset = -w
	cdsMinus.Add(barscds)
	cdsMinus.Legend.Add("cdsMinus", barscdsMinus)
	if err := cds.Save(20*vg.Inch, 10*vg.Inch, "barcdsMinus.png"); err != nil {
		panic(err)
	}

	var exonPlotter plotter.Values
	var exonPlusPlotter plotter.Values
	var exonMinusPlotter plotter.Values

	for i := range exonLengthPlot {
		exonPlotter = append(exonPlotter, float64(exonLengthPlot[i]))
	}
	exon := plot.New()
	exon.Title.Text = "Bar chart"
	exon.Y.Label.Text = "exon"

	barsexon, err := plotter.NewBarChart(exonPlotter, w)
	if err != nil {
		panic(err)
	}

	barsexon.LineStyle.Width = vg.Length(0)
	barsexon.Color = plotutil.Color(0)
	barsexon.Offset = -w
	exon.Add(barsexon)
	exon.Legend.Add("exon", barsexon)
	if err := exon.Save(20*vg.Inch, 10*vg.Inch, "barexon.png"); err != nil {
		panic(err)
	}

	for i := range exonPlusLengthPlot {
		exonPlusPlotter = append(exonPlusPlotter, float64(exonPlusLengthPlot[i]))
	}
	exonPlus := plot.New()
	exonPlus.Title.Text = "Bar chart"
	exonPlus.Y.Label.Text = "exonPlus"

	barsexonPlus, err := plotter.NewBarChart(exonPlusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsexonPlus.LineStyle.Width = vg.Length(0)
	barsexonPlus.Color = plotutil.Color(0)
	barsexonPlus.Offset = -w
	exonPlus.Add(barsexonPlus)
	exonPlus.Legend.Add("exonPlus", barsexonPlus)
	if err := exonPlus.Save(20*vg.Inch, 10*vg.Inch, "barexonPlus.png"); err != nil {
		panic(err)
	}

	for i := range exonMinusLengthPlot {
		exonMinusPlotter = append(exonPlotter, float64(exonMinusLengthPlot[i]))
	}
	exonMinus := plot.New()
	exonMinus.Title.Text = "Bar chart"
	exonMinus.Y.Label.Text = "exon"

	barsexonMinus, err := plotter.NewBarChart(exonMinusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsexonMinus.LineStyle.Width = vg.Length(0)
	barsexonMinus.Color = plotutil.Color(0)
	barsexonMinus.Offset = -w
	exonMinus.Add(barsexon)
	exonMinus.Legend.Add("exonMinus", barsexonMinus)
	if err := exon.Save(20*vg.Inch, 10*vg.Inch, "barexonMinus.png"); err != nil {
		panic(err)
	}

	var proteinPlotter plotter.Values
	var proteinPlusPlotter plotter.Values
	var proteinMinusPlotter plotter.Values

	for i := range proteinLengthPlot {
		proteinPlotter = append(proteinPlotter, float64(proteinLengthPlot[i]))
	}
	protein := plot.New()
	protein.Title.Text = "Bar chart"
	protein.Y.Label.Text = "protein"

	barsprotein, err := plotter.NewBarChart(proteinPlotter, w)
	if err != nil {
		panic(err)
	}

	barsprotein.LineStyle.Width = vg.Length(0)
	barsprotein.Color = plotutil.Color(0)
	barsprotein.Offset = -w
	protein.Add(barsprotein)
	protein.Legend.Add("protein", barsprotein)
	if err := protein.Save(20*vg.Inch, 10*vg.Inch, "barprotein.png"); err != nil {
		panic(err)
	}

	for i := range proteinPlusLengthPlot {
		proteinPlusPlotter = append(proteinPlusPlotter, float64(proteinPlusLengthPlot[i]))
	}
	proteinPlus := plot.New()
	proteinPlus.Title.Text = "Bar chart"
	proteinPlus.Y.Label.Text = "proteinPlus"

	barsproteinPlus, err := plotter.NewBarChart(proteinPlusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsproteinPlus.LineStyle.Width = vg.Length(0)
	barsproteinPlus.Color = plotutil.Color(0)
	barsproteinPlus.Offset = -w
	proteinPlus.Add(barsproteinPlus)
	proteinPlus.Legend.Add("proteinPlus", barsproteinPlus)
	if err := proteinPlus.Save(20*vg.Inch, 10*vg.Inch, "barproteinPlus.png"); err != nil {
		panic(err)
	}

	for i := range proteinMinusLengthPlot {
		proteinMinusPlotter = append(proteinPlotter, float64(proteinMinusLengthPlot[i]))
	}
	proteinMinus := plot.New()
	proteinMinus.Title.Text = "Bar chart"
	proteinMinus.Y.Label.Text = "protein"

	barsproteinMinus, err := plotter.NewBarChart(proteinMinusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsproteinMinus.LineStyle.Width = vg.Length(0)
	barsproteinMinus.Color = plotutil.Color(0)
	barsproteinMinus.Offset = -w
	proteinMinus.Add(barsprotein)
	proteinMinus.Legend.Add("proteinMinus", barsproteinMinus)
	if err := protein.Save(20*vg.Inch, 10*vg.Inch, "barproteinMinus.png"); err != nil {
		panic(err)
	}

	var threePlotter plotter.Values
	var threePlusPlotter plotter.Values
	var threeMinusPlotter plotter.Values

	for i := range threeLengthPlot {
		threePlotter = append(threePlotter, float64(threeLengthPlot[i]))
	}
	three := plot.New()
	three.Title.Text = "Bar chart"
	three.Y.Label.Text = "three"

	barsthree, err := plotter.NewBarChart(threePlotter, w)
	if err != nil {
		panic(err)
	}

	barsthree.LineStyle.Width = vg.Length(0)
	barsthree.Color = plotutil.Color(0)
	barsthree.Offset = -w
	three.Add(barsthree)
	three.Legend.Add("three", barsthree)
	if err := three.Save(20*vg.Inch, 10*vg.Inch, "barthree.png"); err != nil {
		panic(err)
	}

	for i := range threePlusLengthPlot {
		threePlusPlotter = append(threePlusPlotter, float64(threePlusLengthPlot[i]))
	}
	threePlus := plot.New()
	threePlus.Title.Text = "Bar chart"
	threePlus.Y.Label.Text = "threePlus"

	barsthreePlus, err := plotter.NewBarChart(threePlusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsthreePlus.LineStyle.Width = vg.Length(0)
	barsthreePlus.Color = plotutil.Color(0)
	barsthreePlus.Offset = -w
	threePlus.Add(barsthreePlus)
	threePlus.Legend.Add("threePlus", barsthreePlus)
	if err := threePlus.Save(20*vg.Inch, 10*vg.Inch, "barthreePlus.png"); err != nil {
		panic(err)
	}

	for i := range threeMinusLengthPlot {
		threeMinusPlotter = append(threePlotter, float64(threeMinusLengthPlot[i]))
	}
	threeMinus := plot.New()
	threeMinus.Title.Text = "Bar chart"
	threeMinus.Y.Label.Text = "three"

	barsthreeMinus, err := plotter.NewBarChart(threeMinusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsthreeMinus.LineStyle.Width = vg.Length(0)
	barsthreeMinus.Color = plotutil.Color(0)
	barsthreeMinus.Offset = -w
	threeMinus.Add(barsthree)
	threeMinus.Legend.Add("threeMinus", barsthreeMinus)
	if err := three.Save(20*vg.Inch, 10*vg.Inch, "barthreeMinus.png"); err != nil {
		panic(err)
	}

	var fivePlotter plotter.Values
	var fivePlusPlotter plotter.Values
	var fiveMinusPlotter plotter.Values

	for i := range fiveLengthPlot {
		fivePlotter = append(fivePlotter, float64(fiveLengthPlot[i]))
	}
	five := plot.New()
	five.Title.Text = "Bar chart"
	five.Y.Label.Text = "five"

	barsfive, err := plotter.NewBarChart(fivePlotter, w)
	if err != nil {
		panic(err)
	}

	barsfive.LineStyle.Width = vg.Length(0)
	barsfive.Color = plotutil.Color(0)
	barsfive.Offset = -w
	five.Add(barsfive)
	five.Legend.Add("five", barsfive)
	if err := five.Save(20*vg.Inch, 10*vg.Inch, "barfive.png"); err != nil {
		panic(err)
	}

	for i := range fivePlusLengthPlot {
		fivePlusPlotter = append(fivePlusPlotter, float64(fivePlusLengthPlot[i]))
	}
	fivePlus := plot.New()
	fivePlus.Title.Text = "Bar chart"
	fivePlus.Y.Label.Text = "fivePlus"

	barsfivePlus, err := plotter.NewBarChart(fivePlusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsfivePlus.LineStyle.Width = vg.Length(0)
	barsfivePlus.Color = plotutil.Color(0)
	barsfivePlus.Offset = -w
	fivePlus.Add(barsfivePlus)
	fivePlus.Legend.Add("fivePlus", barsfivePlus)
	if err := fivePlus.Save(20*vg.Inch, 10*vg.Inch, "barfivePlus.png"); err != nil {
		panic(err)
	}

	for i := range fiveMinusLengthPlot {
		fiveMinusPlotter = append(fivePlotter, float64(fiveMinusLengthPlot[i]))
	}
	fiveMinus := plot.New()
	fiveMinus.Title.Text = "Bar chart"
	fiveMinus.Y.Label.Text = "five"

	barsfiveMinus, err := plotter.NewBarChart(fiveMinusPlotter, w)
	if err != nil {
		panic(err)
	}

	barsfiveMinus.LineStyle.Width = vg.Length(0)
	barsfiveMinus.Color = plotutil.Color(0)
	barsfiveMinus.Offset = -w
	fiveMinus.Add(barsfive)
	fiveMinus.Legend.Add("fiveMinus", barsfiveMinus)
	if err := five.Save(20*vg.Inch, 10*vg.Inch, "barfiveMinus.png"); err != nil {
		panic(err)
	}

}
