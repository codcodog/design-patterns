package chain

func ExampleChain() {
	case1 := &Case1{}
	case2 := &Case2{}
	case3 := &Case3{}

	case1.setNextCase(case2)
	case2.setNextCase(case3)

	case1.handleRequest(89)
	case1.handleRequest(899)
	case1.handleRequest(8999)
	// Output:
	// case1 handle 89 RMB
	// case2 handle 899 RMB
	// case3 handle 8999 RMB
}
