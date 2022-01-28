package event

var tilstand string
var items = [] string{
	"kylling",
	"rev",
	"korn",
}




func Tilstander(tilstanden string)string{
	tilstand = tilstanden
	tilstand = "[kylling rev korn mennesker ---V \\____/___________/ Ø---]"
	return tilstanden
}

func FirstPut(item string)string  {
	tilstand = "["
	for i:=0; i < len(items); i++{
		if item != items[i]{
			tilstand += items[i] + " "
		}
	}

	if(item == items[1]){
		tilstand = ""+items[0]+ " spiser " +items[2]
	}
	if(item == items[2]){
		tilstand = ""+items[1]+ " spiser " +items[0]
	}
	if(item == items[0]) {
		items = append(items[:0])
		tilstand += " ---V \\_hs+"+item+"_/___________/ Ø---]"
	}else{
		tilstand = "Error!Skriv kylling!"
	}
	return Tilstander(tilstand)



}

//mennesker og elementer som skal i båten
//samme som i put funksjonen

func GetIn(item string)string  {
	tilstand = "[rev korn ---V \\_hs+"+item+"_/___________/ Ø---]"
	return Tilstander(tilstand)

}

func CrossRiver(item string)string  {
	tilstand = "[rev korn ---V \\___________\\_hs+"+item+"___/ Ø---]"
	return Tilstander(tilstand)

}

func TakeOut(item string)string  {
	tilstand = "[rev korn ---V \\___________\\_hs___/ Ø--- "+item+"]"
	return Tilstander(tilstand)

}

func GetOut(item string)string{
	tilstand = "[rev korn ---V \\___________\\____/ Ø--- hs "+item+"]"
	return Tilstander(tilstand)
}
