package normalizer

import "gopkg.in/bblfsh/sdk.v2/driver"

var Transforms = driver.Transforms{
	Namespace:   "java",
	Preprocess:  Preprocess,
	Normalize:   Normalize,
	Annotations: Native,
	Code:        Code,
}
