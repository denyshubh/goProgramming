package main

import "fmt"

type song struct {
	Title  string
	Artist string
	next   *song
}

type playlist struct {
	NowPlaying *song
	Head       *song
	Name       string
}

// Create a playlist
// Add song
// show all songs
// start playing
// Next/ Previous Song

func createPlaylist(name string) *playlist {
	return &playlist{
		Name: name,
	}
}

func (p *playlist) addSong(title, artist string) error {
	s := &song{
		Title:  title,
		Artist: artist,
	}
	if p.Head == nil {
		p.Head = s
	} else {
		currNode := p.Head
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = s
	}
	return nil
}

func (p *playlist) showAllSongs() error {
	currNode := p.Head
	if currNode == nil {
		fmt.Println("Playlist is empty")
		return nil
	}
	fmt.Printf("%+v\n", *currNode)
	for currNode.next != nil {
		currNode = currNode.next
		fmt.Printf("%+v\n", *currNode)
	}
	return nil
}

func (p *playlist) startPlaying() *song {
	p.NowPlaying = p.Head
	return p.NowPlaying
}

func (p *playlist) nextSong() *song {
	p.NowPlaying = p.NowPlaying.next
	return p.NowPlaying
}

func main() {
	playlistName := "myfavouritesong"
	playlist := createPlaylist(playlistName)
	fmt.Printf("A new playlist, %s, is created\n\n", playlistName)
	fmt.Println("Adding songs to the playlist ...")
	playlist.addSong("Ophelia", "The Lumineers")
	playlist.addSong("Shape of you", "Ed Sheeran")
	playlist.addSong("Stubborn Loave", "The Lumineers")
	playlist.addSong("Feels", "Calvin Harris")

	fmt.Println("Showing all songs in playlist...")
	playlist.showAllSongs()
	fmt.Println()

	nowPlaying := playlist.startPlaying()
	fmt.Printf("Now playing: %s by %s\n", nowPlaying.Title, nowPlaying.Artist)
	// fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Title, playlist.NowPlaying.Artist)

	playlist.nextSong()
	fmt.Println("Changing Next Song...")
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Title, playlist.NowPlaying.Artist)
}
