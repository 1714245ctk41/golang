package main

type iShirt interface{
	getLogo() string
	setLogo(logo string)
	getSize() int
	setSize(size int)
}

type shirt struct{
	logo string
	size int
}

func (s *shirt) getLogo() string{
	return s.logo
}
func (s *shirt) setLogo(logo string){
	s.logo = logo
}

func (s *shirt) getSize() int{
	return s.size
}
func (s *shirt) setSize(size int){
	s.size = size
}