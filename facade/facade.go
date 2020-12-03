package facade

import "fmt"

type API interface {
	Test() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type apiImpl struct {
	a AModuleApi
	b BModuleApi
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type AModuleApi interface {
	TestA() string
}

type aModuleImpl struct{}

func NewAModuleAPI() AModuleApi {
	return &aModuleImpl{}
}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

type BModuleApi interface {
	TestB() string
}

type bModuleImpl struct{}

func NewBModuleAPI() BModuleApi {
	return &bModuleImpl{}
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
