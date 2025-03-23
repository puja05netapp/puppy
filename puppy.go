package puppy

import (
	"github.com/puja05netapp/new_pub"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func Bigbark() string {
	return new_pub.WhenGroupUp(Bark())
}

func Bigbarks() string {
	return new_pub.WhenGroupUp(Barks())
}
