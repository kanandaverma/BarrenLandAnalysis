package main

import (
	"fmt"
	"strings"
	"strconv"
	"container/list"
	"sort"
	"bufio"
	"os"
)

func main() {
	var b Barren
	b.ReadInputSTDIN()
	b.EditInput(b.Input)
	b.RestartMatrix_MarkRectangles()
	b.Fertile()
	if b.Output=="240000"{
		fmt.Println("The whole field is barren or the coordinates do not form a specified rectangle")
	}else{
		fmt.Println("Output",b.Output)
	}

}

const X=400
const Y=600

type Barren struct{
	Input string
	AllRect list.List
	Queue list.List
	MapArea map[int]int
	ColorMatrix [X][Y]int
	Output string
}

func (Barren *Barren) ReadInputSTDIN(){
	var n int
	fmt.Println("enter number of fertile rectangle")
	fmt.Scan(&n)
	Barren.Set(n)
}

func (Barren *Barren) Set(n int){
	a := make([]string, n)
	fmt.Println("enter coords of fertile rectangles")
	reader := bufio.NewReader(os.Stdin)
	for i:=0;i<n;i++{
		text, _ := reader.ReadString('\n')
		a[i]=text
	}
	Barren.Input=strings.Join(a, ",")
	Barren.Input=strings.Replace(Barren.Input,"\n", "", -1)
}

func (Barren *Barren) EditInput(Input string){
	var parts []string = strings.Split(Input, ",")
	for i:=0;i<len(parts);i++{
		parts[i]=strings.Replace(parts[i], "\"", "", -1)
		parts[i]=strings.Replace(parts[i], "\\{|\\}", "", -1)
		parts[i]=strings.Replace(parts[i], "^ " , "", -1)
		parts[i]=strings.Replace(parts[i], "{", "", -1)
		parts[i]=strings.Replace(parts[i], "}", "", -1)
		if parts[i]!=""{
			var coord []string=strings.Split(parts[i], " ")
			if len(coord)<4{
				fmt.Println("Please enter coordinates of rectangles properly")
				os.Exit(1)
			}
			zero,_:=strconv.Atoi(coord[0])
			one,_:=strconv.Atoi(coord[1])
			two,_:=strconv.Atoi(coord[2])
			three,_:=strconv.Atoi(coord[3])
			if zero>=400 || two>=400{
				fmt.Println("Coordinates on X axis go outside the field")
				os.Exit(1)
			}else if one>=600 || three>=600{
				fmt.Println("Coordinates on Y axis go outside the field")
				os.Exit(1)
			} else if zero==0 && one==0 && two==399 && three==499{
				fmt.Println("The whole field is fertile")
				os.Exit(1)
			}
			Barren.AllRect.PushBack([]int{zero,one,two,three})
		}
	}

}

func (Barren *Barren) RestartMatrix_MarkRectangles(){
	for i:=0;i<X;i++{
		for j:=0;j<Y;j++{
			Barren.ColorMatrix[i][j]=0
		}
	}

	for e := Barren.AllRect.Front(); e !=nil; e=e.Next() {
		b:=e.Value.([]int)
		for i:=b[0];i<=b[2];i++{
			for j:=b[1];j<=b[3];j++{
				Barren.ColorMatrix[i][j]=1
			}
		}
	}

}

func (Barren *Barren) AddToQueue(i int, j int){
	if Barren.ColorMatrix[i][j]==0 {
		temp:=[]int{i,j}
		Barren.Queue.PushBack(temp)
	}
}

func (Barren *Barren) Fertile(){
	Barren.MapArea = make(map[int]int)
	land:=1
	i:=0
	j:=0
	for ;i<X && j<Y;{
		if Barren.Queue.Front()==nil{
			node:=[]int{i,j}
			if Barren.ColorMatrix[i][j]==0{
				land++
				Barren.MapArea[land]=0
				Barren.Queue.PushBack(node)
			}
			if i==X-1{
				i=0
				j++
			}else{
				i++
			}
		}
		if Barren.Queue.Front()!=nil{
			n:=Barren.Queue.Front()
			node:=n.Value.([]int)
			Barren.Queue.Remove(Barren.Queue.Front())
			x:=node[0]
			y:=node[1]

			if Barren.ColorMatrix[x][y]==0{
				if x > 0 {
					Barren.AddToQueue(x-1, y)
				}
				if x < (X - 1) {
					Barren.AddToQueue(x+1, y)
				}
				if y > 0{
					Barren.AddToQueue(x, y-1)
				}
				if y < (Y - 1){
					Barren.AddToQueue(x, y+1)
				}
				Barren.ColorMatrix[x][y]=land
				Barren.MapArea[land]=Barren.MapArea[land] + 1
			}
		}
	}

	result:= make([]int, len(Barren.MapArea))
	k:=0
	for _,v:=range Barren.MapArea{
		result[k]=v
		k++
	}
	sort.Ints(result)
	s:=strings.Trim(strings.Join(strings.Split(fmt.Sprint(result), " "), " "), "[]")
	Barren.Output=strings.Replace(s, "\\[|\\]|,", "", -1)
}




