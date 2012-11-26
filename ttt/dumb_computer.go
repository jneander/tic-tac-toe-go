package ttt

import "math/rand"

type DumbComputer struct {
  mark string
}

func NewDumbComputer() *DumbComputer {
  return new( DumbComputer )
}

func ( c *DumbComputer ) Move( b Board ) ( move int ) {
  for move = rand.Intn( 9 ); b.Spaces()[move] != b.Blank(); {
    move = rand.Intn( 9 )
  }
  return
}

func ( c *DumbComputer ) SetMark( mark string ) {
  c.mark = mark
}

func ( c *DumbComputer ) GetMark() string {
  return c.mark
}
