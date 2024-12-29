package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

const gopher = `               ..-xZY""""""""""Y0+..
          ..Z"!                ..  ."4,...JJ..
  .Y"77"HY^ .,77?7!..      ./^    .7.  Ta    .W.
 K   ..J^ .7        .i    ,'         1.  WMN.  b
.]  (M#   \....       t  .!.NN,       t   WF  .F
 W,  d'  .(MM#H       (  .,MM@d       c    N..Y
  ?5j%    t?WB^      .\   1.?^       .'    ,]
    d      1.       .^.gNaJ(+.     .?'      W
   .F        ?7<(?7!.?MMMM@:.  ??!          -|
   .]              J         1              .]
   ,\              .=</7J?7,(^               F
   -}                .' (  r                 @
   ,]                .i."i.^                 @
   .]                                        @
    N                                        @
    d.                                       @
    ({                                       @
 ..?4)                                       #_?!,
r.  -}                                       H. <.
(..?d~                                       d 77!
    d                                        d
    d                                        J~
    W                                        ({
    k                                        ,}
    H                                        -}
    H                                        J'
    W                                        K
    W                                       .F
    d.                                      d
    .h                                     .^
     .N.                                  .3
       Tx                               .V'
       .?^?=.                        .,!  1.
      r.!  ..SJ..               ...v"!7, .,(.
      1\ .?     _7""TTTVVTY""""^'       1.,?`

func main() {
	fd := int(os.Stdout.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(gopher, "\n")
	gopherWidth := 0
	for _, line := range lines {
		if len(line) > gopherWidth {
			gopherWidth = len(line)
		}
	}
	os.Stdout.Write([]byte("\x1b[2J"))
	for i := 1; i <= width+gopherWidth; i++ {
		for index, line := range lines {
			if index >= height {
				break
			}
			var builder strings.Builder
			builder.WriteString(fmt.Sprintf("\x1b[%d;0H\x1b[K", index+1))
			if i <= len(line) {
				builder.WriteString(strings.Repeat(" ", width-i))
				builder.WriteString(line[:i])
			} else if i <= width {
				builder.WriteString(strings.Repeat(" ", width-i))
				builder.WriteString(line)
			} else if i-width < len(line) {
				builder.WriteString(line[i-width:])

			}
			os.Stdout.Write([]byte(builder.String()))
		}
		time.Sleep(40 * time.Millisecond)
	}
	os.Stdout.Write([]byte("\n"))
}
