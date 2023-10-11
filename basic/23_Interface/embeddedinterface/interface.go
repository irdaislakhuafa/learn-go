package embeddedinterface

type count1 interface {
	Large() float64
	Around() float64
}

type count2 interface {
	Volume() float64
}

type Count interface {
	count1
	count2
}
