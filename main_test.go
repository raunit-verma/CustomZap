package main

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v4"
)

func BenchmarkCustomZap(t *testing.B) {
	a := GitRegistry{Id: 324, Name: "Raunit Verma", Url: "dfsdfksdkfjdkfj.com", Password: "password is this", SshPrivateKey: "sdfsdfkjsdkfjshdfasjfk;dfsdf", AccessToken: "asdfsdfjkfjksdfjkfjakfjkdkjflkdsjflkj", AuthMode: GitHostRequest{Id: 23234, Name: "sdfdfsdaffads", Active: true, WebhookUrl: "dffsdffsdfdf", WebhookSecret: "adfsfsdafdfsdaf", EventTypeHeader: "dsffsfsdaff", SecretHeader: "sfdfdsffsdfdsfdsfsdf", SecretValidator: "dfasdfdsaffdfdfsdfdf", UserId: 2343}, Active: true, UserId: 34323, GitHostId: 343}

	b := DockerArtifactStoreBean{Id: "adfdfdsfsdfdfds", PluginId: "dfdsafdsfsdfdfd", RegistryURL: "sdfdsfdsfdsfsdfsfdsfsdfds", RegistryType: a, IsOCICompliantRegistry: true, OCIRegistryConfig: map[string]string{"key": "value", "dasfdsfs": "sdfdsfdsf"}, IsPublic: false, RepositoryList: []string{"hello", "i am raunit"}, AWSAccessKeyId: "aws access key is this ", AWSSecretAccessKey: "no secret access key is found here", AWSRegion: "south-east-100", Username: "raunit", Password: "iraunit", IsDefault: true, Connection: "No connection found", Cert: "dsfdskfsdkfjsdklfjdskfjdskfjaskdfjsdkfjdkslf", Active: true, DisabledFields: a, User: 2332, DockerRegistryIpsConfig: a}

	c := GitHostRequest{Id: 234, Name: "raunit is here", Active: true, WebhookUrl: "sdfjsdklfsdklfsdjlfjsdklfjsd", WebhookSecret: "slkdfsdlkfjsdklfjsdkfjkldsfj", EventTypeHeader: "klsdfjdlskfjdksfjksdjf", SecretHeader: "kdfjklfjdskfjskfj", SecretValidator: "kdfjdlksfjdsklfjksld", UserId: 324324}

	d := Test{Username: "sdkfdjsklfdsjkfjsdkf", Password: "lksdfjlkdsafjdsklfjkldsfjklsfjkdslfj", FullName: "kldsfjlkdsfjkldsfjkdsjfklds", Class: 34, Email: "kldsdfjkldsfkldsfkld", Role: "lkdsfjkdfjksldjkls", CheckInt: 3432, Test2: &Test2{Check: 234, Numbers: []int{234, 324, 324, 32, 42, 3}, Username: "sdfjskfsdkfsd", Password: "sdklkfjdsklfdskf", FullName: "jdsflksdjkflsdjfkl", Class: 3432, Email: "kldsjfjkldjfkdsfjksd", Role: "lksdjfkllkfdsk", CheckInt: 32432, MyMap: &map[string]string{"kdsfjdsf": "fdsfsdafsdf", "dfdsaasdf": "dsfdfdsafsd"}, Test3: &Test3{Check: 234, Numbers: []int{234, 324, 324, 32, 42, 3}, Username: "sdfjskfsdkfsd", Password: "sdklkfjdsklfdskf", FullName: "jdsflksdjkflsdjfkl", Class: 3432, Email: "kldsjfjkldjfkdsfjksd", Role: "lksdjfkllkfdsk", CheckInt: 32432}}}

	e := ComprehensiveStruct{}
	err := faker.FakeData(&e)
	if err != nil {
		fmt.Println(err)
	}

	f := StructOne{}

	err = faker.FakeData(&f)
	if err != nil {
		fmt.Println(err)
	}
	// err = faker.FakeData(&c)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = faker.FakeData(&d)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	l, _ := NewSugardLogger("custom")
	// fmt.Println(getNil())
	// l.Infow("message ", "A is : ", Test2{MyMap: nil, Test3: getNil()})
	t.ResetTimer()

	CustomZap(l, a, &b, &c, &d, &e, f)
	for i := 0; i < t.N; i++ {
	}
	// fmt.Println(a, b, c, d)

}
