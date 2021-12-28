export function add(df, nodeId) {

    let nodeData = {}
    nodeData = df.getNodeFromId(nodeId)

    const inputs = nodeData.inputs
    let addNum = 0
    // console.log("#####", inputs)
    if (inputs.input_1.connections.at(0) && inputs.input_2.connections.at(0)) {
        for (const input in inputs) {
            console.log("#####", input)
            const inputId = inputs[input].connections[0].node
            const inputData = df.getNodeFromId(inputId)
            addNum += inputData.data.num
        }

    }
    nodeData.data.num = addNum;
    df.updateNodeDataFromId(nodeId, nodeData.data);
}

export function sub(df, nodeId) {

    let nodeData = {}
    nodeData = df.getNodeFromId(nodeId)

    const inputs = nodeData.inputs
    let subNum = []

    if (inputs.input_1.connections.at(0) && inputs.input_2.connections.at(0)) {
        for (const input in inputs) {
            const inputId = inputs[input].connections[0].node
            const inputData = df.getNodeFromId(inputId)
            subNum.push(inputData.data.num)
        }

    }
    nodeData.data.num = subNum[0] - subNum[1];
    df.updateNodeDataFromId(nodeId, nodeData.data);
}