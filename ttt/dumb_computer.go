package ttt

import "math/rand"

type DumbComputer struct {}

func ( c *DumbComputer ) Move( b Board ) ( move int ) {
  for move = rand.Intn( 9 ); b.Spaces()[move] != b.Blank(); {
    move = rand.Intn( 9 )
  }
  return
}
