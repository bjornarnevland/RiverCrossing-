package state


var tilstand string

func ViewState()string{
	return Tilstander()
}

func Tilstander()string{
	tilstand = "[kylling rev korn menneske ---V \\____/___________/ Ø---]"
	return tilstand
}
