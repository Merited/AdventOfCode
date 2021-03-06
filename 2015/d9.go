package main

import (
	"strconv"
	"strings"
	"bufio"
	"os"
	"fmt"
)

type vertex struct {
	label string
	edgeContainer map[string]edge
}

type edge struct {
	start string
	end string
	cost int
}

type graph struct {
	vertexContainer map[string]vertex
}

func newGraph()(graph){
	g := graph{}
	g.vertexContainer = make(map[string]vertex)
	return g
}


func (g *graph) addEdge(begin string,finish string,weight int) {
	// both vertices must exist before we can even
	// make the edge
	e := edge{start: begin,end:finish,cost:weight}
	g.vertexContainer[begin].edgeContainer[finish] = e
	g.vertexContainer[finish].edgeContainer[begin] = e
}
func (g *graph) edgeList()([]edge){
	x := []edge{}

	for _,value := range g.vertexContainer{
		for i := range value.edgeContainer{
			x = append(x,value.edgeContainer[i])
		}
	}
	return x
}
func newVertex(t string)(vertex){
	b := vertex{}
	b.label = t
	b.edgeContainer = make(map[string]edge)
	return b
}
func (g *graph) addVertex(name string){
	p := newVertex(name)
	g.vertexContainer[name] = p
}
func d9a(){
	maxInt := 2147483647
	g:= newGraph()
	edgeList := []edge{}
	f,err := os.Open("d9.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	marked := map[string]int
	for i := range lines{
		if marked[n[0]] == 0 {
			marked[n[0]] = 2
		}
		n := strings.Split(lines[i]," ")
		number,_ := strconv.Atoi(n[4])
		if g.vertexContainer[n[0]].label == ""{
			v := newVertex(n[0])
			g.vertexContainer[n[0]] = v
		}
		if g.vertexContainer[n[2]].label == ""{
			v := newVertex(n[2])
			g.vertexContainer[n[2]] = v
		}
		h := edge{start:n[0],end:n[2],cost:number}
		g.vertexContainer[h.start].edgeContainer[h.end] = h
		g.vertexContainer[h.end].edgeContainer[h.start] = h
		edgeList = append(edgeList,h)
	}
	flag := true
	current := ""
	total := 0
	for flag {
		min := maxInt
		if marked[current] == 1 {
			return total
		}
		for k := range edgeList {
			if edgeList[k].cost < min && (marked[edgelist[k].start]==2 || marked[edgeList[k].end]==2) {
				min = edgeList[k].cost
			} 
		}
	}

	/*
	// graph done, now time for ffloyds
	// o could also be a 2D array
	o := make(map[string]map[string]int)
	for key,_ := range g.vertexContainer{
		o[key] = make(map[string]int)
		for key1,_ := range g.vertexContainer{
			o[key][key1] = maxInt
		}
		o[key][key] = 0
	}
	edgeList := g.edgeList()
	for e := range edgeList{
		x,y := edgeList[e].start,edgeList[e].end
		c := edgeList[e].cost
		fmt.Println(x,y,c)
		o[x][y] = c
		o[y][x] = c
	}
	*/
	/*
	for v,_ := range g.vertexContainer{
		for w,_ := range g.vertexContainer{
			for x,_ := range g.vertexContainer{
				if o[w][x] > (o[w][v] + o[v][x]){
					fmt.Println(strconv.Itoa(o[w][x]) + "Is longer than: " + strconv.Itoa(o[w][v] + o[v][x]))
					o[w][x] = (o[w][v] + o[v][x]) 
				}
			}
		}
	}
	*/
	fmt.Println(o)
	labelList := make(map[string]int)
	for key,_ := range g.vertexContainer{
		labelList[key] = 0
	}
	min := maxInt
	total := 0
	flag := true
	for flag {
		for key,value := range o{
			for key1,value1 := range value{
				
			}
		}
	}
	}
	/*
	min := maxInt
	for _,value := range o{
		minValue := 0
		for _,value2 := range value{
			minValue += value2
		}
		fmt.Println(minValue)
		if minValue < min{
			min = minValue
		}
	}
	fmt.Println(min)
	*/
//fmt.Println(edgeList)
}


func main(){
	d9a()
}