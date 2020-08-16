package monster

import "testing"

func TestStore(t *testing.T) {
	monster := &Monster{
		Name: "albert",
		Age: 21,
		Skill: "king",
	}

	res := monster.Store("../data/monster.json")
	if !res {
		t.Fatalf("Store(\"../data/monster.json\") 出错, 期望值=%v, 实际值=%v\n", true, res)
	} else {
		t.Logf("Store(\"../data/monster.json\") 执行成功")
	}
}

func TestReStore(t *testing.T) {
	monster := &Monster{}
	
	res := monster.ReStore("../data/monster.json")
	if !res {
		t.Fatalf("ReStore(\"../data/monster.json\") 出错, 期望值=%v, 实际值=%v\n", true, res)
	} else {
		t.Logf("ReStore(\"../data/monster.json\") 执行成功")
	}
}