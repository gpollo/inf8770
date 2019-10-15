import sys
import image
import data_pb2
import pywt
import numpy

def apply_transform(data, mode):
    splitted = numpy.split(data, 2, 0)
    top = splitted[0]
    bottom = splitted[1]

    splitted = numpy.split(top, 2, 1)
    cA = splitted[0]
    cH = splitted[1]

    splitted = numpy.split(bottom, 2, 1)
    cV = splitted[0]
    cD = splitted[1]
    
    return pywt.idwt2((cA, (cH, cV, cD)), mode)

if __name__== "__main__":
    input_data = data_pb2.PythonDWT()
    input_data.ParseFromString(sys.stdin.buffer.read())
    input_matrix = image.protobuf_to_matrix(input_data.data)
    output_matrix = apply_transform(input_matrix, input_data.mode)
    output_data = image.matrix_to_protobuf(output_matrix)

    sys.stdout.buffer.write(output_data.SerializeToString())
    sys.stdin.close()
    sys.stdout.close()
