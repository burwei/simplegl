package simplegl

func NewSimpleCube() *[]float32{
	return &[]float32{
		//  X, Y, Z, U, V
		// Bottom (-Y)
		-10.0, -10.0, -10.0, 0.0, 0.0,
		10.0, -10.0, -10.0, 1.0, 0.0,
		-10.0, -10.0, 10.0, 0.0, 1.0,
		10.0, -10.0, -10.0, 1.0, 0.0,
		10.0, -10.0, 10.0, 1.0, 1.0,
		-10.0, -10.0, 10.0, 0.0, 1.0,
	
		// Top (+Y)
		-10.0, 10.0, -10.0, 0.0, 0.0,
		-10.0, 10.0, 10.0, 0.0, 1.0,
		10.0, 10.0, -10.0, 1.0, 0.0,
		10.0, 10.0, -10.0, 1.0, 0.0,
		-10.0, 10.0, 10.0, 0.0, 1.0,
		10.0, 10.0, 10.0, 1.0, 1.0,
	
		// Front (+Z)
		-10.0, -10.0, 10.0, 1.0, 0.0,
		10.0, -10.0, 10.0, 0.0, 0.0,
		-10.0, 10.0, 10.0, 1.0, 1.0,
		10.0, -10.0, 10.0, 0.0, 0.0,
		10.0, 10.0, 10.0, 0.0, 1.0,
		-10.0, 10.0, 10.0, 1.0, 1.0,
	
		// Back (-Z)
		-10.0, -10.0, -10.0, 0.0, 0.0,
		-10.0, 10.0, -10.0, 0.0, 1.0,
		10.0, -10.0, -10.0, 1.0, 0.0,
		10.0, -10.0, -10.0, 1.0, 0.0,
		-10.0, 10.0, -10.0, 0.0, 1.0,
		10.0, 10.0, -10.0, 1.0, 1.0,
	
		// Left (-X)
		-10.0, -10.0, 10.0, 0.0, 1.0,
		-10.0, 10.0, -10.0, 1.0, 0.0,
		-10.0, -10.0, -10.0, 0.0, 0.0,
		-10.0, -10.0, 10.0, 0.0, 1.0,
		-10.0, 10.0, 10.0, 1.0, 1.0,
		-10.0, 10.0, -10.0, 1.0, 0.0,
	
		// Right (+X)
		10.0, -10.0, 10.0, 1.0, 1.0,
		10.0, -10.0, -10.0, 1.0, 0.0,
		10.0, 10.0, -10.0, 0.0, 0.0,
		10.0, -10.0, 10.0, 1.0, 1.0,
		10.0, 10.0, -10.0, 0.0, 0.0,
		10.0, 10.0, 10.0, 0.0, 1.0,
	}
}