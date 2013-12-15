package main

import (
    "fmt"
    "github.com/nsf/termbox-go"
)

type key struct {
	x int
	y int
	ch rune
}

var K_1 = []key{{4,4,'1'}}
var K_BANG = []key{{4,4,'!'}}
var K_2 = []key{{7,4,'2'}}
var K_AT = []key{{7,4,'@'}}
var K_3 = []key{{10,4,'3'}}
var K_HASH = []key{{10,4,'#'}}
var K_4 = []key{{13,4,'4'}}
var K_DOLLAR = []key{{13,4,'$'}}
var K_5 = []key{{16,4,'5'}}
var K_PERCENT = []key{{16,4,'%'}}
var K_6 = []key{{19,4,'6'}}
var K_CARET = []key{{19,4,'^'}}
var K_7 = []key{{22,4,'7'}}
var K_AMPERSAND = []key{{22,4,'&'}}
var K_8 = []key{{25,4,'8'}}
var K_ASTERISK = []key{{25,4,'*'}}
var K_9 = []key{{28,4,'9'}}
var K_0 = []key{{31,4,'0'}}
var K_MINUS = []key{{34,4,'-'}}
var K_MINUS_SHIFT = []key{{34,4,'_'}}
var K_EQUALS = []key{{37,4,'='}}
var K_EQUALS_SHIFT = []key{{37,4,'+'}}
var K_BACKSLASH = []key{{40,4,'\\'}}
var K_BACKSLASH_SHIFT = []key{{40,4,'|'}}
var K_BACKSPACE = []key{{44,4,0x2190},{45,4,0x2500},{46,4,0x2500}}
var K_K_SLASH = []key{{68,4,'/'}}
var K_K_STAR = []key{{71,4,'*'}}
var K_K_MINUS = []key{{74,4,'-'}}
var K_q = []key{{6,6,'q'}}
var K_Q = []key{{6,6,'Q'}}
var K_w = []key{{9,6,'w'}}
var K_W = []key{{9,6,'W'}}
var K_e = []key{{12,6,'e'}}
var K_E = []key{{12,6,'E'}}
var K_r = []key{{15,6,'r'}}
var K_R = []key{{15,6,'R'}}
var K_t = []key{{18,6,'t'}}
var K_T = []key{{18,6,'T'}}
var K_y = []key{{21,6,'y'}}
var K_Y = []key{{21,6,'Y'}}
var K_u = []key{{24,6,'u'}}
var K_U = []key{{24,6,'U'}}
var K_i = []key{{27,6,'i'}}
var K_I = []key{{27,6,'I'}}
var K_o = []key{{30,6,'o'}}
var K_O = []key{{30,6,'O'}}
var K_p = []key{{33,6,'p'}}
var K_P = []key{{33,6,'P'}}
var K_ENTER = []key{
    {43,6,0x2591},{44,6,0x2591},{45,6,0x2591},{46,6,0x2591},
    {43,7,0x2591},{44,7,0x2591},{45,7,0x21B5},{46,7,0x2591},
    {41,8,0x2591},{42,8,0x2591},{43,8,0x2591},{44,8,0x2591},
    {45,8,0x2591},{46,8,0x2591},}
var K_K_7 = []key{{65,6,'7'}}
var K_K_8 = []key{{68,6,'8'}}
var K_K_9 = []key{{71,6,'9'}}
var K_K_PLUS = []key{{74,6,' '},{74,7,'+'},{74,8,' '}}
var K_a = []key{{7,8,'a'}}
var K_A = []key{{7,8,'A'}}
var K_s = []key{{10,8,'s'}}
var K_S = []key{{10,8,'S'}}
var K_d = []key{{13,8,'d'}}
var K_D = []key{{13,8,'D'}}
var K_f = []key{{16,8,'f'}}
var K_F = []key{{16,8,'F'}}
var K_g = []key{{19,8,'g'}}
var K_G = []key{{19,8,'G'}}
var K_h = []key{{22,8,'h'}}
var K_H = []key{{22,8,'H'}}
var K_j = []key{{25,8,'j'}}
var K_J = []key{{25,8,'J'}}
var K_k = []key{{28,8,'k'}}
var K_K = []key{{28,8,'K'}}
var K_l = []key{{31,8,'l'}}
var K_L = []key{{31,8,'L'}}
var K_SEMICOLON = []key{{34,8,';'}}
var K_PARENTHESIS = []key{{34,8,':'}}
var K_QUOTE = []key{{37,8,'\''}}
var K_DOUBLEQUOTE = []key{{37,8,'"'}}
var K_K_4 = []key{{65,8,'4'}}
var K_K_5 = []key{{68,8,'5'}}
var K_K_6 = []key{{71,8,'6'}}
var K_z = []key{{9,10,'z'}}
var K_Z = []key{{9,10,'Z'}}
var K_x = []key{{12,10,'x'}}
var K_X = []key{{12,10,'X'}}
var K_c = []key{{15,10,'c'}}
var K_C = []key{{15,10,'C'}}
var K_v = []key{{18,10,'v'}}
var K_V = []key{{18,10,'V'}}
var K_b = []key{{21,10,'b'}}
var K_B = []key{{21,10,'B'}}
var K_n = []key{{24,10,'n'}}
var K_N = []key{{24,10,'N'}}
var K_m = []key{{27,10,'m'}}
var K_M = []key{{27,10,'M'}}
var K_COMMA = []key{{30,10,','}}
var K_LANB = []key{{30,10,'<'}}
var K_PERIOD = []key{{33,10,'.'}}
var K_RANB = []key{{33,10,'>'}}
var K_SLASH = []key{{36,10,'/'}}
var K_QUESTION = []key{{36,10,'?'}}
var K_K_1 = []key{{65,10,'1'}}
var K_K_2 = []key{{68,10,'2'}}
var K_K_3 = []key{{71,10,'3'}}
var K_K_ENTER = []key{{74,10,0x2591},{74,11,0x2591},{74,12,0x2591}}
var K_K_0 = []key{{65,12,' '},{66,12,'0'},{67,12,' '},{68,12,' '}}
var K_K_PERIOD = []key{{71,12,'.'}}

func keyChannel(c *chan<-) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	// termbox.Flush()

func main() {
    var c chan
    go keyChannel(&c)
    for {
        ev := termbox.PollEvent()
        if ev.Type == termbox.EventKey {
            if ev.Key == termbox.KeyCtrlC {
                break
            }
            if ev.Type == termbox.EventKey {
                    c <- ev.Ch
            }
        }
    }
}
