package exercises

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nottheeconomist/adventofcode_2015/puzzle"
)

var logger *log.Logger

// Node describes a gate in the circuit
type Node struct {
	// upstream and downstream are the connections coming into and going out of the Node.
	name       string
	upstream   []chan int16
	downstream []chan int16

	// f is the function to apply to the signal.
	f func(...int16) int16
}

// RegisterUpstream registers an upstream connection
func (n *Node) RegisterUpstream(ch chan int16) {
	n.upstream = append(n.upstream, ch)
	logger.Printf("Adding to upstream on %s\n", n.name)
}

// RegisterDownstream registers a downstream connection
func (n *Node) RegisterDownstream(ch chan int16) {
	n.downstream = append(n.downstream, ch)
}

// Start begins operations on the Node, effectively "Starting" the signal
func (n *Node) Start() {
	var result int16
	if len(n.upstream) == 0 {
		logger.Printf("%s has no upstream, sending nothing", n.name)
		logger.Println(*n)
		return
	} else if len(n.upstream) == 1 {
		a := <-n.upstream[0]
		result = n.f(a)
	} else if len(n.upstream) == 2 {
		a := <-n.upstream[0]
		b := <-n.upstream[1]
		result = n.f(a, b)
	} else {
		logger.Fatalf("Upstream is too long! Node named: %s has length %v\n",
			n.name, len(n.upstream))
	}
	for _, ch := range n.downstream {
		ch <- result
	}
}

// Binary functions
func pass(inputs ...int16) int16 {
	// pass automatically through
	if len(inputs) != 1 {
		panic("Input to pass is wrong")
	}
	return inputs[0]
}

func and(inputs ...int16) int16 {
	if len(inputs) != 2 {
		panic("Input to and is wrong")
	}
	return inputs[0] & inputs[1]
}

func or(inputs ...int16) int16 {
	if len(inputs) != 2 {
		panic("Input to or is wrong")
	}
	return inputs[0] | inputs[1]
}

func not(inputs ...int16) int16 {
	if len(inputs) != 1 {
		panic("Input to not is wrong")
	}
	return ^inputs[0]
}

func lshift(inputs ...int16) int16 {
	if len(inputs) != 2 {
		panic("Input to lshift is wrong")
	}
	a, n := inputs[0], uint(inputs[1])
	return a << n
}

func rshift(inputs ...int16) int16 {
	if len(inputs) != 2 {
		panic("Input to pass is wrong")
	}
	a, n := inputs[0], uint(inputs[1])
	return a >> n
}

// GetNode returns the *Node stored at name in nodeMap, if it exists. If it does
// not, then it creates a *Node, stores its reference in nodeMap, and returns it.
func GetNode(name string, nodesMap map[string]*Node) *Node {
	node, ok := nodesMap[name]
	if !ok {
		node = &Node{name: name,
			f: pass}
		nodesMap[name] = node
	}
	return node
}

// DaySeven runs the calculations for the Day Seven puzzle
func DaySeven(p *puzzle.Puzzle) {

	f, err := os.Create("log.log")
	if err != nil {
		panic("failed to open log")
	}
	logger = log.New(f, "", log.Lshortfile)

	nodesMap := make(map[string]*Node)
	for _, line := range strings.Split(dayseveninput, "\n") {
		fields := strings.Fields(line)

		logger.Printf("Building line: %s\n", line)
		target := GetNode(fields[len(fields)-1], nodesMap)
		chTarget := make(chan int16)
		target.RegisterUpstream(chTarget)

		switch len(fields) {
		case 3:
			// value -> target
			// or
			// from -> target
			if value, err := strconv.ParseInt(fields[0], 10, 16); err != nil {
				// fields[0] is a node name, not a value
				from := GetNode(fields[0], nodesMap)
				from.RegisterDownstream(chTarget)
				from.f = pass
			} else {
				go func() {
					value := int16(value)
					chTarget <- value
				}()
			}
		case 4:
			// NOT from -> target
			from := GetNode(fields[1], nodesMap)
			from.RegisterDownstream(chTarget)
			from.f = not
		case 5:
			// from1|val AND|OR from2|val -> target
			// or
			// from1|val LSHIFT|RSHIFT n -> target
			// TODO: Code currently fails because it assumes that
			// the first value will always be a node, but it doesn't have to be
			from1 := GetNode(fields[0], nodesMap)
			chFrom1 := make(chan int16)
			from1.RegisterDownstream(chFrom1)

			switch fields[1] {
			case "AND", "OR":
				from2 := GetNode(fields[2], nodesMap)
				chFrom2 := make(chan int16)
				from2.RegisterDownstream(chFrom2)

				// gates don't need to be recorded anywhere and can start immediately
				gate := &Node{upstream: []chan int16{chFrom1, chFrom2},
					downstream: []chan int16{chTarget}}
				if fields[1] == "AND" {
					gate.f = and
				} else {
					gate.f = or
				}
				go gate.Start()

			case "LSHIFT", "RSHIFT":
				chValue := make(chan int16)
				if n, err := strconv.ParseInt(fields[2], 10, 16); err != nil {
					panic("Bad input! " + line)
				} else {
					go func(chValue chan int16, n int16) { chValue <- n }(chValue, int16(n))
				}

				// gates don't need to be recorded anywhere and can start immediately
				gate := &Node{upstream: []chan int16{chFrom1, chValue},
					downstream: []chan int16{chTarget}}
				gate.RegisterDownstream(chTarget)
				if fields[1] == "LSHIFT" {
					gate.f = lshift
				} else {
					gate.f = rshift
				}
				go gate.Start()
			}
		}
	}
	logger.Println("All done building logic chain!")
	// register listener on node `a`
	listener := make(chan int16)
	if nodeA, ok := nodesMap["a"]; !ok {
		panic("No node named ``a``")
	} else {
		nodeA.RegisterDownstream(listener)
	}
	for _, node := range nodesMap {
		go node.Start()
	}
	result := <-listener
	p.AddSolution(strconv.Itoa(int(result)))
	logger.Printf("Logged node a's output: %v\n", result)
}

const dayseveninput = `af AND ah -> ai
NOT lk -> ll
hz RSHIFT 1 -> is
NOT go -> gp
du OR dt -> dv
x RSHIFT 5 -> aa
at OR az -> ba
eo LSHIFT 15 -> es
ci OR ct -> cu
b RSHIFT 5 -> f
fm OR fn -> fo
NOT ag -> ah
v OR w -> x
g AND i -> j
an LSHIFT 15 -> ar
1 AND cx -> cy
jq AND jw -> jy
iu RSHIFT 5 -> ix
gl AND gm -> go
NOT bw -> bx
jp RSHIFT 3 -> jr
hg AND hh -> hj
bv AND bx -> by
er OR es -> et
kl OR kr -> ks
et RSHIFT 1 -> fm
e AND f -> h
u LSHIFT 1 -> ao
he RSHIFT 1 -> hx
eg AND ei -> ej
bo AND bu -> bw
dz OR ef -> eg
dy RSHIFT 3 -> ea
gl OR gm -> gn
da LSHIFT 1 -> du
au OR av -> aw
gj OR gu -> gv
eu OR fa -> fb
lg OR lm -> ln
e OR f -> g
NOT dm -> dn
NOT l -> m
aq OR ar -> as
gj RSHIFT 5 -> gm
hm AND ho -> hp
ge LSHIFT 15 -> gi
jp RSHIFT 1 -> ki
hg OR hh -> hi
lc LSHIFT 1 -> lw
km OR kn -> ko
eq LSHIFT 1 -> fk
1 AND am -> an
gj RSHIFT 1 -> hc
aj AND al -> am
gj AND gu -> gw
ko AND kq -> kr
ha OR gz -> hb
bn OR by -> bz
iv OR jb -> jc
NOT ac -> ad
bo OR bu -> bv
d AND j -> l
bk LSHIFT 1 -> ce
de OR dk -> dl
dd RSHIFT 1 -> dw
hz AND ik -> im
NOT jd -> je
fo RSHIFT 2 -> fp
hb LSHIFT 1 -> hv
lf RSHIFT 2 -> lg
gj RSHIFT 3 -> gl
ki OR kj -> kk
NOT ak -> al
ld OR le -> lf
ci RSHIFT 3 -> ck
1 AND cc -> cd
NOT kx -> ky
fp OR fv -> fw
ev AND ew -> ey
dt LSHIFT 15 -> dx
NOT ax -> ay
bp AND bq -> bs
NOT ii -> ij
ci AND ct -> cv
iq OR ip -> ir
x RSHIFT 2 -> y
fq OR fr -> fs
bn RSHIFT 5 -> bq
0 -> c
14146 -> b
d OR j -> k
z OR aa -> ab
gf OR ge -> gg
df OR dg -> dh
NOT hj -> hk
NOT di -> dj
fj LSHIFT 15 -> fn
lf RSHIFT 1 -> ly
b AND n -> p
jq OR jw -> jx
gn AND gp -> gq
x RSHIFT 1 -> aq
ex AND ez -> fa
NOT fc -> fd
bj OR bi -> bk
as RSHIFT 5 -> av
hu LSHIFT 15 -> hy
NOT gs -> gt
fs AND fu -> fv
dh AND dj -> dk
bz AND cb -> cc
dy RSHIFT 1 -> er
hc OR hd -> he
fo OR fz -> ga
t OR s -> u
b RSHIFT 2 -> d
NOT jy -> jz
hz RSHIFT 2 -> ia
kk AND kv -> kx
ga AND gc -> gd
fl LSHIFT 1 -> gf
bn AND by -> ca
NOT hr -> hs
NOT bs -> bt
lf RSHIFT 3 -> lh
au AND av -> ax
1 AND gd -> ge
jr OR js -> jt
fw AND fy -> fz
NOT iz -> ja
c LSHIFT 1 -> t
dy RSHIFT 5 -> eb
bp OR bq -> br
NOT h -> i
1 AND ds -> dt
ab AND ad -> ae
ap LSHIFT 1 -> bj
br AND bt -> bu
NOT ca -> cb
NOT el -> em
s LSHIFT 15 -> w
gk OR gq -> gr
ff AND fh -> fi
kf LSHIFT 15 -> kj
fp AND fv -> fx
lh OR li -> lj
bn RSHIFT 3 -> bp
jp OR ka -> kb
lw OR lv -> lx
iy AND ja -> jb
dy OR ej -> ek
1 AND bh -> bi
NOT kt -> ku
ao OR an -> ap
ia AND ig -> ii
NOT ey -> ez
bn RSHIFT 1 -> cg
fk OR fj -> fl
ce OR cd -> cf
eu AND fa -> fc
kg OR kf -> kh
jr AND js -> ju
iu RSHIFT 3 -> iw
df AND dg -> di
dl AND dn -> do
la LSHIFT 15 -> le
fo RSHIFT 1 -> gh
NOT gw -> gx
NOT gb -> gc
ir LSHIFT 1 -> jl
x AND ai -> ak
he RSHIFT 5 -> hh
1 AND lu -> lv
NOT ft -> fu
gh OR gi -> gj
lf RSHIFT 5 -> li
x RSHIFT 3 -> z
b RSHIFT 3 -> e
he RSHIFT 2 -> hf
NOT fx -> fy
jt AND jv -> jw
hx OR hy -> hz
jp AND ka -> kc
fb AND fd -> fe
hz OR ik -> il
ci RSHIFT 1 -> db
fo AND fz -> gb
fq AND fr -> ft
gj RSHIFT 2 -> gk
cg OR ch -> ci
cd LSHIFT 15 -> ch
jm LSHIFT 1 -> kg
ih AND ij -> ik
fo RSHIFT 3 -> fq
fo RSHIFT 5 -> fr
1 AND fi -> fj
1 AND kz -> la
iu AND jf -> jh
cq AND cs -> ct
dv LSHIFT 1 -> ep
hf OR hl -> hm
km AND kn -> kp
de AND dk -> dm
dd RSHIFT 5 -> dg
NOT lo -> lp
NOT ju -> jv
NOT fg -> fh
cm AND co -> cp
ea AND eb -> ed
dd RSHIFT 3 -> df
gr AND gt -> gu
ep OR eo -> eq
cj AND cp -> cr
lf OR lq -> lr
gg LSHIFT 1 -> ha
et RSHIFT 2 -> eu
NOT jh -> ji
ek AND em -> en
jk LSHIFT 15 -> jo
ia OR ig -> ih
gv AND gx -> gy
et AND fe -> fg
lh AND li -> lk
1 AND io -> ip
kb AND kd -> ke
kk RSHIFT 5 -> kn
id AND if -> ig
NOT ls -> lt
dw OR dx -> dy
dd AND do -> dq
lf AND lq -> ls
NOT kc -> kd
dy AND ej -> el
1 AND ke -> kf
et OR fe -> ff
hz RSHIFT 5 -> ic
dd OR do -> dp
cj OR cp -> cq
NOT dq -> dr
kk RSHIFT 1 -> ld
jg AND ji -> jj
he OR hp -> hq
hi AND hk -> hl
dp AND dr -> ds
dz AND ef -> eh
hz RSHIFT 3 -> ib
db OR dc -> dd
hw LSHIFT 1 -> iq
he AND hp -> hr
NOT cr -> cs
lg AND lm -> lo
hv OR hu -> hw
il AND in -> io
NOT eh -> ei
gz LSHIFT 15 -> hd
gk AND gq -> gs
1 AND en -> eo
NOT kp -> kq
et RSHIFT 5 -> ew
lj AND ll -> lm
he RSHIFT 3 -> hg
et RSHIFT 3 -> ev
as AND bd -> bf
cu AND cw -> cx
jx AND jz -> ka
b OR n -> o
be AND bg -> bh
1 AND ht -> hu
1 AND gy -> gz
NOT hn -> ho
ck OR cl -> cm
ec AND ee -> ef
lv LSHIFT 15 -> lz
ks AND ku -> kv
NOT ie -> if
hf AND hl -> hn
1 AND r -> s
ib AND ic -> ie
hq AND hs -> ht
y AND ae -> ag
NOT ed -> ee
bi LSHIFT 15 -> bm
dy RSHIFT 2 -> dz
ci RSHIFT 2 -> cj
NOT bf -> bg
NOT im -> in
ev OR ew -> ex
ib OR ic -> id
bn RSHIFT 2 -> bo
dd RSHIFT 2 -> de
bl OR bm -> bn
as RSHIFT 1 -> bl
ea OR eb -> ec
ln AND lp -> lq
kk RSHIFT 3 -> km
is OR it -> iu
iu RSHIFT 2 -> iv
as OR bd -> be
ip LSHIFT 15 -> it
iw OR ix -> iy
kk RSHIFT 2 -> kl
NOT bb -> bc
ci RSHIFT 5 -> cl
ly OR lz -> ma
z AND aa -> ac
iu RSHIFT 1 -> jn
cy LSHIFT 15 -> dc
cf LSHIFT 1 -> cz
as RSHIFT 3 -> au
cz OR cy -> da
kw AND ky -> kz
lx -> a
iw AND ix -> iz
lr AND lt -> lu
jp RSHIFT 5 -> js
aw AND ay -> az
jc AND je -> jf
lb OR la -> lc
NOT cn -> co
kh LSHIFT 1 -> lb
1 AND jj -> jk
y OR ae -> af
ck AND cl -> cn
kk OR kv -> kw
NOT cv -> cw
kl AND kr -> kt
iu OR jf -> jg
at AND az -> bb
jp RSHIFT 2 -> jq
iv AND jb -> jd
jn OR jo -> jp
x OR ai -> aj
ba AND bc -> bd
jl OR jk -> jm
b RSHIFT 1 -> v
o AND q -> r
NOT p -> q
k AND m -> n
as RSHIFT 2 -> at`
