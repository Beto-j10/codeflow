export const data1 = {
    "drawflow": {
        "Home": {
            "data": {
                "1": {
                    "id": 1,
                    "name": "If",
                    "data": {
                        "operator": "<"
                    },
                    "class": "IfStatement",
                    "html": "If",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "2",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "2",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {},
                    "pos_x": 283,
                    "pos_y": 208
                },
                "2": {
                    "id": 2,
                    "name": "Num",
                    "data": {
                        "num": 0
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "1",
                                    "output": "input_1"
                                },
                                {
                                    "node": "1",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 80,
                    "pos_y": 253
                }
            }
        },
        "If5": {
            "data": {
                "3": {
                    "id": 3,
                    "name": "Num",
                    "data": {
                        "num": 13
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "4",
                                    "output": "input_1"
                                },
                                {
                                    "node": "4",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 33,
                    "pos_y": 113
                },
                "4": {
                    "id": 4,
                    "name": "Add",
                    "data": {
                        "num": 26,
                        "operator": "+"
                    },
                    "class": "BinaryExpression Addition ops",
                    "html": "Add",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "3",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "3",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {
                        "output_1": {
                            "connections": []
                        }
                    },
                    "pos_x": 231,
                    "pos_y": 142
                }
            }
        }
    }
}

export const data2 = {
    "drawflow": {
        "Home": {
            "data": {
            }
        },
        "If4": {
            "data": {
            }
        },
        "Else4": {
            "data": {
            }
        }
    }
}

export const data3 = {
    "drawflow": {
        "Home": {
            "data": {
                "3": {
                    "id": 3,
                    "name": "Num",
                    "data": {
                        "num": 1
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "7",
                                    "output": "input_1"
                                }
                            ]
                        }
                    },
                    "pos_x": 17,
                    "pos_y": 78
                },
                "4": {
                    "id": 4,
                    "name": "Num",
                    "data": {
                        "num": 2
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "7",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 18,
                    "pos_y": 203
                },
                "7": {
                    "id": 7,
                    "name": "If",
                    "data": {
                        "operator": "=="
                    },
                    "class": "IfStatement",
                    "html": "If",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "3",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "4",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {},
                    "pos_x": 224,
                    "pos_y": 113
                }
            }
        },
        "If7": {
            "data": {
                "8": {
                    "id": 8,
                    "name": "Num",
                    "data": {
                        "num": 9
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "9",
                                    "output": "input_1"
                                },
                                {
                                    "node": "9",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 9,
                    "pos_y": 177
                },
                "9": {
                    "id": 9,
                    "name": "Add",
                    "data": {
                        "num": 18,
                        "operator": "+"
                    },
                    "class": "BinaryExpression Addition ops",
                    "html": "Add",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "8",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "8",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {
                        "output_1": {
                            "connections": []
                        }
                    },
                    "pos_x": 207,
                    "pos_y": 150
                }
            }
        }
    }
}

export const data4 = {
    "drawflow": {
        "Home": {
            "data": {
                "3": {
                    "id": 3,
                    "name": "Num",
                    "data": {
                        "num": 1
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "7",
                                    "output": "input_1"
                                }
                            ]
                        }
                    },
                    "pos_x": 17,
                    "pos_y": 78
                },
                "4": {
                    "id": 4,
                    "name": "Num",
                    "data": {
                        "num": 2
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "7",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 18,
                    "pos_y": 203
                },
                "7": {
                    "id": 7,
                    "name": "If",
                    "data": {
                        "operator": "=="
                    },
                    "class": "IfStatement",
                    "html": "If",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "3",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "4",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {},
                    "pos_x": 224,
                    "pos_y": 113
                },
                "10": {
                    "id": 10,
                    "name": "If",
                    "data": {
                        "operator": "=="
                    },
                    "class": "IfStatement",
                    "html": "If",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "11",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "12",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {},
                    "pos_x": 233,
                    "pos_y": 331
                },
                "11": {
                    "id": 11,
                    "name": "Num",
                    "data": {
                        "num": 11
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "10",
                                    "output": "input_1"
                                }
                            ]
                        }
                    },
                    "pos_x": 20,
                    "pos_y": 324
                },
                "12": {
                    "id": 12,
                    "name": "Num",
                    "data": {
                        "num": 5
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "10",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 20,
                    "pos_y": 460
                }
            }
        },
        "If7": {
            "data": {
                "8": {
                    "id": 8,
                    "name": "Num",
                    "data": {
                        "num": 9
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "9",
                                    "output": "input_1"
                                },
                                {
                                    "node": "9",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 9,
                    "pos_y": 177
                },
                "9": {
                    "id": 9,
                    "name": "Add",
                    "data": {
                        "num": 18,
                        "operator": "+"
                    },
                    "class": "BinaryExpression Addition ops",
                    "html": "Add",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "8",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "8",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {
                        "output_1": {
                            "connections": []
                        }
                    },
                    "pos_x": 207,
                    "pos_y": 150
                }
            }
        },
        "If10": {
            "data": {
                "13": {
                    "id": 13,
                    "name": "Num",
                    "data": {
                        "num": 8
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "15",
                                    "output": "input_1"
                                }
                            ]
                        }
                    },
                    "pos_x": 29,
                    "pos_y": 79
                },
                "14": {
                    "id": 14,
                    "name": "Num",
                    "data": {
                        "num": 9
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "15",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 30,
                    "pos_y": 202
                },
                "15": {
                    "id": 15,
                    "name": "Add",
                    "data": {
                        "num": 17,
                        "operator": "+"
                    },
                    "class": "BinaryExpression Addition ops",
                    "html": "Add",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "13",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "14",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "17",
                                    "output": "input_1"
                                }
                            ]
                        }
                    },
                    "pos_x": 203,
                    "pos_y": 87
                },
                "16": {
                    "id": 16,
                    "name": "Mult",
                    "data": {
                        "num": 68,
                        "operator": "*"
                    },
                    "class": "BinaryExpression Multiplication ops",
                    "html": "Mult",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "18",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "19",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "20",
                                    "output": "input_1"
                                }
                            ]
                        }
                    },
                    "pos_x": 195,
                    "pos_y": 302
                },
                "17": {
                    "id": 17,
                    "name": "Assign",
                    "data": {
                        "num": 17,
                        "var": "Assign17"
                    },
                    "class": "VariableDeclarator Assign ops",
                    "html": "Assign",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "15",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {},
                    "pos_x": 401,
                    "pos_y": 97
                },
                "18": {
                    "id": 18,
                    "name": "Assign17",
                    "data": {
                        "num": 17,
                        "idParent": "17"
                    },
                    "class": "Identifier ops",
                    "html": "Assign17",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "16",
                                    "output": "input_1"
                                }
                            ]
                        }
                    },
                    "pos_x": 51,
                    "pos_y": 342
                },
                "19": {
                    "id": 19,
                    "name": "Num",
                    "data": {
                        "num": 4
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "16",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 25,
                    "pos_y": 421
                },
                "20": {
                    "id": 20,
                    "name": "Assign",
                    "data": {
                        "num": 68,
                        "var": "Assign20"
                    },
                    "class": "VariableDeclarator Assign ops",
                    "html": "Assign",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "16",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {},
                    "pos_x": 371,
                    "pos_y": 319
                },
                "21": {
                    "id": 21,
                    "name": "Assign20",
                    "data": {
                        "num": 68,
                        "idParent": "20"
                    },
                    "class": "Identifier ops",
                    "html": "Assign20",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "23",
                                    "output": "input_1"
                                }
                            ]
                        }
                    },
                    "pos_x": 43,
                    "pos_y": 548
                },
                "22": {
                    "id": 22,
                    "name": "Num",
                    "data": {
                        "num": 4
                    },
                    "class": "NumericLiteral",
                    "html": "Num",
                    "typenode": "vue",
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "23",
                                    "output": "input_2"
                                }
                            ]
                        }
                    },
                    "pos_x": 21,
                    "pos_y": 622
                },
                "23": {
                    "id": 23,
                    "name": "Div",
                    "data": {
                        "num": 17,
                        "operator": "/"
                    },
                    "class": "BinaryExpression Division ops",
                    "html": "Div",
                    "typenode": "vue",
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "node": "21",
                                    "input": "output_1"
                                }
                            ]
                        },
                        "input_2": {
                            "connections": [
                                {
                                    "node": "22",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {
                        "output_1": {
                            "connections": []
                        }
                    },
                    "pos_x": 234,
                    "pos_y": 496
                }
            }
        }
    }
}