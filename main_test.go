package main

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v4"
)

func BenchmarkCustomZap(t *testing.B) {
	a := GitRegistry{}
	b := DockerArtifactStoreBean{}
	c := GitHostRequest{}
	d := Test{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	err = faker.FakeData(&b)
	if err != nil {
		fmt.Println(err)
	}
	err = faker.FakeData(&c)
	if err != nil {
		fmt.Println(err)
	}
	err = faker.FakeData(&d)

	if err != nil {
		fmt.Println(err)
	}

	l, _ := NewSugardLogger("custom")

	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		CustomZap(l, &a, &b, &c, &d)
	}

}
