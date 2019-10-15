import sys
import image
import data_pb2
import pywt
import numpy

def apply_transform(data, mode):
    cA, (cH, cV, cD) = pywt.dwt2(data, mode)

    top = numpy.concatenate((cA, cH), 1)
    bottom = numpy.concatenate((cV, cD), 1)
 
    return numpy.concatenate((top, bottom), 0)

if __name__== "__main__":
    input_data = data_pb2.ProtoDWT()
    input_data.ParseFromString(sys.stdin.buffer.read())
    input_matrix = image.protobuf_to_matrix(input_data.data)
    output_matrix = apply_transform(input_matrix, input_data.mode)
    output_data = image.matrix_to_protobuf(output_matrix)

    sys.stdout.buffer.write(output_data.SerializeToString())
    sys.stdin.close()
    sys.stdout.close()
