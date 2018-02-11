package GamePlus

type IGamePlus interface {
	Init()
	Run()
	End()
	About() string
}
