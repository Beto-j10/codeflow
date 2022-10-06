/**
 * Adds to the value of the node the sum of the values of its inputs
 * @param {object} df
 * @param {string} nodeId
 */
export function add(df, nodeId) {

    let nodeData = {}
    nodeData = df.getNodeFromId(nodeId)

    const inputs = nodeData.inputs
    let addNum = 0
    if (inputs.input_1.connections.at(0) && inputs.input_2.connections.at(0)) {
        for (const input in inputs) {
            const inputId = inputs[input].connections[0].node
            const inputData = df.getNodeFromId(inputId)
            addNum += inputData.data.num
        }
        nodeData.data.num = Math.round(addNum * 100) / 100
    }
    df.updateNodeDataFromId(nodeId, nodeData.data);
}

/**
 * the value of input 1 subtracts the value of input 2
 * and adds the result to the value of the node
 * @param {object} df 
 * @param {string} nodeId 
 */

export function subtraction(df, nodeId) {

    let nodeData = {}
    nodeData = df.getNodeFromId(nodeId)

    const inputs = nodeData.inputs
    let inputValues = []

    if (inputs.input_1.connections.at(0) && inputs.input_2.connections.at(0)) {
        for (const input in inputs) {
            const inputId = inputs[input].connections[0].node
            const inputData = df.getNodeFromId(inputId)
            inputValues.push(inputData.data.num) //adds the value of the input to the array
        }

        let result = inputValues[0] - inputValues[1] //subtracts the value of input 2 from input 1
        nodeData.data.num = Math.round(result * 100) / 100

    } else {
        nodeData.data.num = 0
    }
    df.updateNodeDataFromId(nodeId, nodeData.data);
}

/**
 * Add to the value of the node the multiplication of the values of its inputs
 * @param {object} df 
 * @param {string} nodeId 
 */

export function multiplication(df, nodeId) {

    let nodeData = {}
    nodeData = df.getNodeFromId(nodeId)

    const inputs = nodeData.inputs
    let multipliedNums = 1

    if (inputs.input_1.connections.at(0) && inputs.input_2.connections.at(0)) {
        for (const input in inputs) {
            const inputId = inputs[input].connections[0].node
            const inputData = df.getNodeFromId(inputId)
            multipliedNums *= inputData.data.num
        }

        nodeData.data.num = Math.round(multipliedNums * 100) / 100

    } else {
        nodeData.data.num = 0
    }
    df.updateNodeDataFromId(nodeId, nodeData.data);
}


/**
 * the value of input 1 is divided into the value of input 2
 * and adds the result to the value of the node
 * @param {object} df 
 * @param {string} nodeId 
 */

export function division(df, nodeId) {

    let nodeData = {}
    nodeData = df.getNodeFromId(nodeId)

    const inputs = nodeData.inputs
    let inputValues = []
    if (inputs.input_1.connections.at(0) && inputs.input_2.connections.at(0)) {
        for (const input in inputs) {
            const inputId = inputs[input].connections[0].node
            const inputData = df.getNodeFromId(inputId)
            inputValues.push(inputData.data.num) //adds the value of the input to the array
        }

        // Check that the divisor is not 0
        if (inputValues[1] === 0) {
            nodeData.data.num = 0
            df.updateNodeDataFromId(nodeId, nodeData.data);
            return
        }

        let result = inputValues[0] / inputValues[1]; //divides the value of input 1 by input 2
        nodeData.data.num = Math.round(result * 100) / 100

    } else {
        nodeData.data.num = 0
    }
    df.updateNodeDataFromId(nodeId, nodeData.data);
}

/**
 * assigns the value of its input to the node value
 * @param {object} df 
 * @param {string} nodeId 
 * @param {string} varName
 */

/**
 * gets the value of the node and assigns it to the variable
 * @param {object} df 
 * @param {string} nodeId 
 */

export function assignValue(df, nodeId) {

    const nodeData = df.getNodeFromId(nodeId)
    const parentNodeData = df.getNodeFromId(nodeData.data.idParent)

    parentNodeData.data.idChild = nodeId
    nodeData.data.num = parentNodeData.data.num
    df.updateNodeDataFromId(nodeId, nodeData.data);
    df.updateNodeDataFromId(nodeData.data.idParent, parentNodeData.data);
}