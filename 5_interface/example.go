package main

// external library can't modifie the source code
type external struct {
}

func (a *external) EX1() {

}

func (a *external) EX2() {

}

// internal library we can change source code
type internal struct{}

func (i *internal) IN1() {

}

func (i *internal) IN2() {

}

// add method like external
func (i *internal) EX1() {

}

func (i *internal) EX2() {

}

/* so if we want to accept external and internal what we gonna do*/
func use(e *external) {

}

type myinterface interface {
	EX1()
	EX2()
}

/* accept interface */
func use1(m myinterface) {
	m.EX1()
}

func exUse() {

	in := &internal{}
	use1(in)

	ex := &external{}
	use1(ex)

	// NOTE: in oop world we can use adapter pattern. but need new layer to abstract
}
