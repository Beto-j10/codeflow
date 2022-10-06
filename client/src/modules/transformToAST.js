import { transformToCode } from "./transformToCode"

export function transformToAST(data) {
    const preAST = {}
    const nonRootNodes = {}

    Object.keys(data).forEach(moduleName => {
        const module = data[moduleName].data

        preAST[moduleName] = {}
        nonRootNodes[moduleName] = {}

        Object.keys(module).forEach(node => {
            transformNode(module[node], preAST[moduleName])
        })

        Object.keys(preAST[moduleName]).forEach(node => {
            bindingNodes(preAST[moduleName][node], preAST[moduleName], nonRootNodes[moduleName])
        })

        Object.keys(nonRootNodes[moduleName]).forEach(key => {
            delete preAST[moduleName][key]
        })
        if (/(if|else|for)/i.test(moduleName)) {

            const splitModuleName = moduleName.split(/(if|else|for)/i)
            const moduleType = splitModuleName[1].toLowerCase()
            const moduleId = splitModuleName[2]

            switch (moduleType) {
                case "if":
                    Object.keys(preAST[moduleName]).forEach(node => {
                        preAST.Home[moduleId].consequent.body.push(preAST[moduleName][node])
                    })
                    delete preAST[moduleName]
                    break;
                case "else":
                    preAST.Home[moduleId].alternate = { type: "BlockStatement", body: [] }
                    Object.keys(preAST[moduleName]).forEach(node => {
                        preAST.Home[moduleId].alternate.body.push(preAST[moduleName][node])
                    })
                    delete preAST[moduleName]
                    break;
                case "for":
                    Object.keys(preAST[moduleName]).forEach(node => {
                        preAST.Home[moduleId].body.body.push(preAST[moduleName][node])
                    })
                    delete preAST[moduleName]
                    break;
                default:
                    break;
            }
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

    const toCode = transformToCode(AST.program)
    console.log(toCode)

}

function transformNode(node, module) {
    const nodeClass = node.class.split(' ')[0]

    switch (nodeClass) {
        case "NumericLiteral":
            module[node.id] = {
                type: "NumericLiteral",
                id: node.id,
                value: node.data.num,
            }
            break;
        case "BinaryExpression":
            module[node.id] = {
                type: "BinaryExpression",
                id: node.id,
                operator: node.data.operator,
                left: node.inputs.input_1.connections.at(0).node,
                right: node.inputs.input_2.connections.at(0).node,
            }
            break;
        case "VariableDeclarator":
            module[node.id] = {
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
            module[node.id] = {
                type: "Identifier",
                id: node.id,
                name: node.name,
            }
            break;
        case "ForStatement":
            module[node.id] = {
                type: "ForStatement",
                id: node.id,
                init: {
                    type: "Identifier",
                    name: node.data.var,
                },
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
        case "IfStatement":
            module[node.id] = {
                type: "IfStatement",
                id: node.id,
                test: {
                    type: "BinaryExpression",
                    operator: node.data.operator,
                    left: node.inputs.input_1.connections.at(0).node,
                    right: node.inputs.input_2.connections.at(0).node,
                },
                consequent: {
                    type: "BlockStatement",
                    body: []
                },
                alternate: null
            }

        default:
            break;
    }

    const moduleJSON = JSON.stringify(module, null, 2)
    // console.log("moduleJSON", moduleJSON)
}

function bindingNodes(node, module, nonRootNodes) {
    switch (node.type) {
        case "NumericLiteral":

            break;
        case "BinaryExpression":

            nonRootNodes[node.left] = ""
            nonRootNodes[node.right] = ""

            node.left = module[node.left]
            node.right = module[node.right]
            break;

        case "VariableDeclarator":

            nonRootNodes[node.init] = ""

            node.init = module[node.init]
            break;
        case "ForStatement":

            nonRootNodes[node.test.left] = ""
            nonRootNodes[node.test.right] = ""
            nonRootNodes[node.update.value] = ""

            node.test.left = module[node.test.left]
            node.test.right = module[node.test.right]
            node.update.value = module[node.update.value]
            break;
        case "IfStatement":

            nonRootNodes[node.test.left] = ""
            nonRootNodes[node.test.right] = ""

            node.test.left = module[node.test.left]
            node.test.right = module[node.test.right]
            break;


        default:
            break;
    }
}

function bindingBody() {

}