
export function transformToAST(data) {
    const preAST = {}
    const nonRootNodes = {}
    console.log("keys", Object.keys(data))
    Object.keys(data).forEach(key => {
        console.log("key", data[key])
        transformNode(data[key], preAST)
    })

    Object.keys(preAST).forEach(key => {
        console.log("key", preAST[key])
        bindingNodes(preAST[key], preAST, nonRootNodes)
    })

    Object.keys(nonRootNodes).forEach(key => {
        delete preAST[key]
    })

    let ot = preAST

    console.log("preAST#########22", JSON.stringify(ot, null, 2))
    console.log("nonRootNodes =====", JSON.stringify(nonRootNodes, null, 2))



}

function transformNode(node, preAST) {
    switch (node.class) {
        case "NumericLiteral":
            preAST[node.id] = {
                type: "NumericLiteral",
                id: node.id,
                value: node.data.num,
            }
            break;
        case "BinaryExpression":
            preAST[node.id] = {
                type: "BinaryExpression",
                id: node.id,
                operator: node.data.operator,
                left: node.inputs.input_1.connections.at(0).node,
                right: node.inputs.input_2.connections.at(0).node,
            }
            break;
        case "VariableDeclarator":
            preAST[node.id] = {
                type: "VariableDeclarator",
                id: node.id,
                init: node.inputs.input_1.connections.at(0).node,
            }
            break;
        case "ForStatement":
            preAST[node.id] = {
                type: "ForStatement",
                id: node.id,
                test: {
                    type: "BinaryExpression",
                    operator: node.data.operator,
                    left: node.inputs.input_1.connections.at(0).node,
                    right: node.inputs.input_2.connections.at(0).node,
                },
                update: {
                    type: "UpdateExpression",
                    operator: "++", //to implement
                    value: node.inputs.input_3.connections.at(0).node,
                },
                body: node.outputs.output_1.connections.at(0).node,
                // body: {
                //     type: "BlockStatement",
                //     body: []
                // }
            }
            break;

        default:
            break;
    }
}

function bindingNodes(node, preAST, nonRootNodes) {
    switch (node.type) {
        case "NumericLiteral":

            break;
        case "BinaryExpression":

            nonRootNodes[node.left] = ""
            nonRootNodes[node.right] = ""

            node.left = preAST[node.left]
            node.right = preAST[node.right]
            break;

        case "VariableDeclarator":

            nonRootNodes[node.init] = ""

            node.init = preAST[node.init]
            break;
        case "ForStatement":

            nonRootNodes[node.body] = ""

            node.body = preAST[node.body]
            break;

        default:
            break;
    }
}