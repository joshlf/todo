package impl

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleGenerate() {
	rand.Seed(28526) // Make the test deterministic
	fmt.Println(MakeTestTasksN(3, 2, 2))
	// Output:[{Id:8845887351047480521 Start:3662713084969698356 End:2245807622914294392 Completed:true Dependencies:[]}
	// {Id:2189650770990404615 Start:5025781435361181057 End:7492068751658652768 Completed:false Dependencies:[122872713384859104]}
	// {Id:4133444238857250179 Start:9048556065834103247 End:8632018550427270028 Completed:true Dependencies:[]}
	// {Id:5258043255585688963 Start:6560561761263074688 End:8147425073874716161 Completed:false Dependencies:[]}
	// {Id:591495124321297301 Start:4357109288288652689 End:3036062292035631935 Completed:false Dependencies:[]}
	// {Id:122872713384859104 Start:7882079866308123610 End:5381310628619860373 Completed:false Dependencies:[]}]

}

func TestGenerate(t *testing.T) {
	rand.Seed(21469) // Make the test deterministic
	for i := 1; i < 10; i++ {
		tasks := MakeTestTasksN(i, i, i)
		if !tasks.Acyclic() {
			t.Errorf("Generated cyclic graph: %v", tasks)
		}
	}
}
