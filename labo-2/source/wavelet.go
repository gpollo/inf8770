package main

// TODO: what to do if dimensions not multiple of 2?

func (d ImageData) GetXLowPassFilter() ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX/2, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX/2; i++ {
			data[j][i] = (d[j][2*i+1] + d[j][2*i+1]) / 2
		}
	}

	return data
}

func (d ImageData) GetYLowPassFilter() ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY/2)

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX; i++ {
			data[j][i] = (d[2*j+0][i] + d[2*j+1][i]) / 2
		}
	}

	return data
}

func (d ImageData) GetXHighPassFilter() ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX/2, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX/2; i++ {
			data[j][i] = (d[j][2*i+1] - d[j][2*i+1]) / 2
		}
	}

	return data
}

func (d ImageData) GetYHighPassFilter() ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY/2)

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX; i++ {
			data[j][i] = (d[2*j+0][i] - d[2*j+1][i]) / 2
		}
	}

	return data
}
