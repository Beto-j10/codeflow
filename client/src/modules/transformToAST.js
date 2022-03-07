
export function transformToAST(data) {
    const preAST = {}
    const nonRootNodes = {}

    Object.keys(data).forEach(key => {
        const module = data[key].data
        const moduleName = key

        preAST[moduleName] = {}
        nonRootNodes[moduleName] = {}

        console.log("keys", Object.keys(module))
        Object.keys(module).forEach(key => {
            console.log("key", module[key])
            transformNode(module[key], preAST[moduleName])
        })

        Object.keys(preAST[moduleName]).forEach(key => {
            console.log("key", preAST[moduleName][key])
            bindingNodes(preAST[moduleName][key], preAST[moduleName], nonRootNodes[moduleName])
        })

        Object.keys(nonRootNodes[moduleName]).forEach(key => {
            delete preAST[moduleName][key]
        })

        if (moduleName !== "Home") {
            console.log("IFModulePreAST", preAST)
            const nodeId = moduleName.at(-1)
            Object.keys(preAST[moduleName]).forEach(key => {
                preAST.Home[nodeId].body.body.push(preAST[moduleName][key])
            })
            delete preAST[moduleName]
        }
    })
    const AST = {
        "program": {
            "type": "Program",
            "sourceType": "script",
            "body": []
        }
    }
    Object.keys(preAST["Home"]).forEach(key => {
        AST.program.body.push(preAST["Home"][key])
    })

    let ot = preAST

    console.log("preAST#########22", JSON.stringify(ot, null, 2))
    console.log("AST77", JSON.stringify(AST, null, 2))
    console.log("nonRootNodes =====", JSON.stringify(nonRootNodes, null, 2))



}

function transformNode(node, preAST) {
    const nodeClass = node.class.split(' ')[0]

    switch (nodeClass) {
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
                idType: {
                    type: "Identifier",
                    name: node.data.var,
                },
                init: node.inputs.input_1.connections.at(0).node,
            }
            break;
        case "Identifier":
            preAST[node.id] = {
                type: "identifier",
                id: node.id,
                name: node.name,
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
                body: {
                    type: "BlockStatement",
                    body: []
                }
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

            nonRootNodes[node.test.left] = ""
            nonRootNodes[node.test.right] = ""
            nonRootNodes[node.update.value] = ""

            node.test.left = preAST[node.test.left]
            node.test.right = preAST[node.test.right]
            node.update.value = preAST[node.update.value]
            break;

        default:
            break;
    }
}

function bindingBody() {

}