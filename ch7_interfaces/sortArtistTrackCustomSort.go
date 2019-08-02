package main

import (
	"fmt"
	"sort"
)
//Can't run like this: go run sortArtistTrackCustomSort.go
//we need to comment out main in another file and run: go run ch7_interfaces/*.go
//otherwise go will not compile other files in the same package
func main() {
	fmt.Println("---------before------------")
	printTracks(ArtistTracks)
	sort.Sort(customSort{
		tracks: ArtistTracks,
		less: func(x, y *Track) bool {
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
		},
	})
	fmt.Println("\n\n---------after------------")
	printTracks(ArtistTracks)

}

type customSort struct {
	tracks []*Track
	less   func(x, y *Track) bool
}

func (c customSort) Len() int {
	return len(c.tracks)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.tracks[i], c.tracks[j])
}

func (c customSort) Swap(i, j int) {
	c.tracks[i], c.tracks[j] = c.tracks[j], c.tracks[i]
}
