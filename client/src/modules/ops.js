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

    }
    nodeData.data.num = addNum;
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

    }
    nodeData.data.num = inputValues[0] - inputValues[1]; //subtracts the value of input 2 from input 1
    df.updateNodeDataFromId(nodeId, nodeData.data);
}

/**
 * the value of input 1 subtracts the value of input 2
 * and adds the result to the value of the node
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

    }
    nodeData.data.num = multipliedNums;
    df.updateNodeDataFromId(nodeId, nodeData.data);
}