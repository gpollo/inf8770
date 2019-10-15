import data_pb2
import numpy

def protobuf_to_matrix(image_data):
    rows = image_data.rows

    size_x = len(rows[0].values)
    size_y = len(rows)

    matrix = []
    for j in range(size_y):
        row = []
        for i in range(size_x):
            row.append(rows[j].values[i])
        matrix.append(row)
    return numpy.matrix(matrix)

def matrix_to_protobuf(matrix):
    size_x = len(matrix[0])
    size_y = len(matrix)

    image_data = data_pb2.ImageData()
    for j in range(size_y):
        row = data_pb2.ImageRow()
        row.values.extend(matrix[j])
        image_data.rows.extend([row])
    return image_data
