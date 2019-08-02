package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)
func main() {
	fmt.Println("---------before------------")
	printTracks(ArtistTracks)
	sort.Sort(byArtist(ArtistTracks))
	fmt.Println("\n\n\n---------after------------")
	printTracks(ArtistTracks)
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var ArtistTracks = []*Track{
	{"Go", "Delilah", "From the roots up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m387s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) (d time.Duration) {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return
}

type byArtist []*Track

//implementation of sort.Interface
func (b byArtist) Len() int {
	return len(b)
}

func (b byArtist) Less(i, j int) bool {
	return b[i].Artist < b[j].Artist
}

func (b byArtist) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//*[]Track is different with []*Track
// a pointer to an Track array vs a array of pointer to Track
// we can't range over a pointer to an Track array
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	var tw = new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, v := range tracks {
		fmt.Fprintf(tw, format, v.Title, v.Artist, v.Album, v.Year, v.Length)
	}
	tw.Flush()
}
