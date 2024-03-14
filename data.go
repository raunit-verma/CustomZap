package main

const LOG = "log"

type ComprehensiveStruct struct {
	IntValue        int     `log:"hide"` // Basic integer
	IntPointerValue *int    `log:"hide"` // Pointer to an integer
	StringValue     string  `log:"hide"`
	BoolValue       bool    // Boolean (true or false)
	Float32Value    float32 // Single-precision floating-point number
	Float64Value    float64 // Double-precision floating-point number
	ByteValue       byte    // Single byte (unsigned integer)
	RuneValue       rune    // Unicode character
	ArrayValue      [5]int  // Fixed-size array of integers
	//SliceValue      []string       // Dynamically sized slice of strings
	//MapValue      map[string]int // Key-value map (string keys, integer values)
	StructPointer *NestedStruct // Pointer to a nested struct
}

// NestedStruct demonstrates all data types within a nested struct
type NestedStruct struct {
	NestedIntValue    int            `log:"hide"`
	NestedIntPointer  *int           `log:"hide"`
	NestedStringValue string         `log:"hide"`
	BoolValue         bool           `log:"hide"` // Boolean (true or false)
	Float32Value      float32        `log:"hide"` // Single-precision floating-point number
	Float64Value      float64        `log:"hide"` // Double-precision floating-point number
	ByteValue         byte           `log:"hide"` // Single byte (unsigned integer)
	RuneValue         rune           `log:"hide"` // Unicode character
	ArrayValue        [5]int         `log:"hide"` // Fixed-size array of integers
	SliceValue        []string       `log:"hide"` // Dynamically sized slice of strings
	MapValue          map[string]int `log:"hide"` // Key-value map (string keys, integer values)
}

type StructOne struct {
	Integer      int
	Float        float64
	Bool         bool
	String       string
	IntegerPtr   *int
	FloatPtr     *float64
	BoolPtr      *bool
	StringPtr    *string
	NestedStruct StructTwo
	ArrayOfInt   [3]int
}

type StructTwo struct {
	AnotherInt   int
	AnotherFloat float64
	AnotherBool  bool
}
type GitRegistry struct {
	Id            int            `json:"id,omitempty" validate:"number"`
	Name          string         `json:"name,omitempty" validate:"required"`
	Url           string         `json:"url,omitempty"`
	UserName      string         `json:"userName,omitempty"`
	Password      string         `json:"password,omitempty" log:"hide"`
	SshPrivateKey string         `json:"sshPrivateKey,omitempty" log:"hide"`
	AccessToken   string         `json:"accessToken,omitempty" log:"hide"`
	AuthMode      GitHostRequest `json:"authMode,omitempty" validate:"required"`
	Active        bool           `json:"active"`
	UserId        int32          `json:"-"`
	GitHostId     int            `json:"gitHostId"`
}

type DockerArtifactStoreBean struct {
	Id                      string            `json:"id" validate:"required"`
	PluginId                string            `json:"pluginId,omitempty" validate:"required"`
	RegistryURL             string            `json:"registryUrl" validate:"required"`
	RegistryType            GitRegistry       `json:"registryType" validate:"required"`
	IsOCICompliantRegistry  bool              `json:"isOCICompliantRegistry"`
	OCIRegistryConfig       map[string]string `json:"ociRegistryConfig,omitempty"`
	IsPublic                bool              `json:"isPublic"`
	RepositoryList          []string          `json:"repositoryList,omitempty"`
	AWSAccessKeyId          string            `json:"awsAccessKeyId,omitempty"`
	AWSSecretAccessKey      string            `json:"awsSecretAccessKey,omitempty" log:"hide"`
	AWSRegion               string            `json:"awsRegion,omitempty"`
	Username                string            `json:"username,omitempty"`
	Password                string            `json:"password,omitempty" log:"hide"`
	IsDefault               bool              `json:"isDefault"`
	Connection              string            `json:"connection"`
	Cert                    string            `json:"cert"`
	Active                  bool              `json:"active"`
	DisabledFields          GitRegistry       `json:"disabledFields"`
	User                    int32             `json:"-"`
	DockerRegistryIpsConfig GitRegistry       `json:"ipsConfig,omitempty"`
}

type GitHostRequest struct {
	Id              int    `json:"id,omitempty" validate:"number"`
	Name            string `json:"name,omitempty" validate:"required"`
	Active          bool   `json:"active"`
	WebhookUrl      string `json:"webhookUrl"`
	WebhookSecret   string `json:"webhookSecret" log:"hide"`
	EventTypeHeader string `json:"eventTypeHeader"`
	SecretHeader    string `json:"secretHeader"`
	SecretValidator string `json:"secretValidator"`
	UserId          int32  `json:"-"`
}
type Test2 struct {
	Password string `log:"hide"`
	Check    int    `log:"true"`
	Numbers  []int  `log:"false"`
	Test3    *Test3
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Class    int    `json:"class"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role"`
	CheckInt int    `log:"hide"`
	MyMap    *map[string]string
}

type Test3 struct {
	Check    int    `log:"true"`
	Numbers  []int  `log:"false"`
	Username string `json:"df" log:"hide"`
	Password string `json:"password,omitempty" log:"hide"`
	FullName string `json:"fullname"`
	Class    int    `json:"class" log:"hide"`
	Email    string `json:"email,omitempty" log:"hide"`
	Role     string `json:"role"`
	CheckInt int    `log:"hide"`
}

type Test struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty" log:"hide"`
	FullName string `json:"fullname"`
	Class    int    `json:"class"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role"`
	CheckInt int    `log:"hide"`
	Test2    *Test2 `log:"true"`
}
