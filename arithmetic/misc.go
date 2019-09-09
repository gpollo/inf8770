package main

func bitsToByte(bits *[]bool) byte {
	count := len(*bits)
	if count > 8 {
		count = 8
	}

	var packed byte = 0
	for i := 0; i < count; i++ {
		if (*bits)[i] {
			packed += (1 << i)
		}
	}

	*bits = (*bits)[count:]

	return packed
}
