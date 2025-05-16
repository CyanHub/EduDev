package casbin

import (
	"testing"
)

func TestCasbinWithFile(t *testing.T) {
	CasbinWithFile("./model.conf", "./policy.csv")
}

func TestCasbinDynamic(t *testing.T) {
	CasbinDynamic("./model.conf", "./policy.csv")
}

func TestCasbinSavePolicy(t *testing.T) {
	CasbinSavePolicy("./model.conf", "./policy.csv")
}

func TestCasbinWithAdapter(t *testing.T) {
	CasbinWithAdapter("./model.conf", "./policy.csv")
}
