package main

import (
    "os"
    "github.com/nsf/termbox-go"
    "errors"
)

type Coord struct {
    x, y int
}

type NextCoord func(current Coord) (next Coord, err error)

func nextFromDimensions(dimensions Coord) NextCoord {
    return func(current Coord) (next Coord, err error) {
        var nullCoord Coord
        next = current
        next.x++
        if next.x >= dimensions.x {
            next.y++
            next.x = nullCoord.x
        }
        if next.y >= dimensions.y {
            next = nullCoord
            err = errors.New("logic: coordinate exceeds screen boundary")
        }
        return
    }
}

type Keypair struct {
    ascii rune
    meta termbox.Key
}

func eventsToChannels(keychan chan Keypair, dimchan chan Coord) {
    for {
        ev := termbox.PollEvent()
        if ev.Err != nil {
            panic(ev.Err)
        }
        if ev.Type == termbox.EventKey {
            pair := Keypair{ev.Ch, ev.Key}
            keychan <- pair
        }
        if ev.Type == termbox.EventResize {
            dimensions := Coord{ ev.Width, ev.Height }
            dimchan <- dimensions
        }
    }
}

func getFirstDimensions() Coord {
    ev := termbox.PollEvent()
    dimensions := Coord{ ev.Width, ev.Height }
    return dimensions
}

func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()
    termbox.SetInputMode(termbox.InputEsc)
    termbox.Flush()
    dimensions := getFirstDimensions()
    getNext := nextFromDimensions(dimensions)

    var keychan chan Keypair
    var dimchan chan Coord
    go eventsToChannels(keychan, dimchan)

    var thisCoord Coord

    for {
        select {
            case keys := <- keychan:
                if keys.meta == termbox.KeyCtrlF {
                    os.Exit(1)
                }
                termbox.SetCell(
                    thisCoord.x,
                    thisCoord.y,
                    keys.ascii,
                    termbox.ColorWhite,
                    termbox.ColorBlack,
                )
                termbox.Flush()
                thisCoord, err = getNext(thisCoord)
                if err != nil {
                    panic(err)
                }
            case dimensions := <- dimchan:
                getNext = nextFromDimensions(dimensions)
        }
    }
}

