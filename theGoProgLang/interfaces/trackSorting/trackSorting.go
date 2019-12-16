package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (a byArtist) Len() int           { return len(a) }
func (a byArtist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byArtist) Less(i, j int) bool { return a[i].Artist < a[j].Artist }

type sortTrack struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (a sortTrack) Len() int           { return len(a.t) }
func (a sortTrack) Swap(i, j int)      { a.t[i], a.t[j] = a.t[j], a.t[i] }
func (a sortTrack) Less(i, j int) bool { return a.less(a.t[i], a.t[j]) }

func main() {
	printTracks(tracks)

	fmt.Println()
	sort.Sort(byArtist(tracks))

	printTracks(tracks)

	titleSort := func(x, y *Track) bool {
		return x.Title < y.Title
	}

	orderTrack := func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}

	sort.Sort(sortTrack{tracks, titleSort})

	fmt.Println()
	printTracks(tracks)
	fmt.Println(sort.IsSorted(sortTrack{tracks, orderTrack}))
	sort.Sort(sortTrack{tracks, orderTrack})

	fmt.Println()
	printTracks(tracks)
	fmt.Println(sort.IsSorted(sortTrack{tracks, orderTrack}))

}
