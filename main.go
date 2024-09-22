package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

// Embed the image and audio assets.
//
//go:embed assets/img/paper.png
var paperPng []byte

//go:embed assets/audio/Stamp_old_3_16b_.wav
var stampWav []byte

type Game struct {
	mousePressedLastFrame bool
	img                   *ebiten.Image
	audioContext          *audio.Context
	player                *audio.Player
}

// NewGame initializes the game by loading assets.
func NewGame() (*Game, error) {
	game := &Game{}

	// Load the embedded image
	imgData, _, err := image.Decode(bytes.NewReader(paperPng))
	if err != nil {
		return nil, err
	}
	game.img = ebiten.NewImageFromImage(imgData)
	log.Println("Image loaded successfully")

	return game, nil
}

// initializeAudio sets up the audio context and player.
func (g *Game) initializeAudio() error {
	if g.audioContext != nil && g.player != nil {
		return nil // Already initialized
	}

	// Initialize audio context
	g.audioContext = audio.NewContext(44100)
	log.Println("Audio context initialized")

	// Load the embedded audio
	audioBuffer := bytes.NewReader(stampWav)
	wavStream, err := wav.DecodeWithSampleRate(44100, audioBuffer)
	if err != nil {
		return err
	}

	g.player, err = g.audioContext.NewPlayer(wavStream)
	if err != nil {
		return err
	}
	log.Println("Audio player created successfully")

	return nil
}

// Update handles the game logic.
func (g *Game) Update() error {
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mousePressed && !g.mousePressedLastFrame {
		x, y := ebiten.CursorPosition()
		imgBounds := g.img.Bounds()
		if x >= 0 && x < imgBounds.Dx() && y >= 0 && y < imgBounds.Dy() {
			// Initialize audio on first click
			if err := g.initializeAudio(); err != nil {
				log.Printf("Audio initialization failed: %v", err)
				return nil // Continue running even if audio fails
			}

			// Play the audio
			if g.player != nil {
				g.player.Rewind()
				g.player.Play()
				log.Println("Audio played")
			}
		}
	}
	g.mousePressedLastFrame = mousePressed
	return nil
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.img, op)
}

// Layout defines the screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}

func main() {
	ebiten.SetWindowSize(640, 360)
	ebiten.SetWindowTitle("Alchemist of the Shadow Bureau")

	game, err := NewGame()
	if err != nil {
		log.Fatalf("Failed to initialize game: %v", err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
