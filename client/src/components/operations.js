export function add(df, nodeId) {

    let dataNode = {}
    dataNode = df.getNodeFromId(nodeId)

    const inputs = dataNode.inputs
    let addNum = 0

    if (inputs.input_1.connections.at(0) && inputs.input_2.connections.at(0)) {
        for (const input in inputs) {
            const idInput = inputs[input].connections[0].node
            const dataInput = df.getNodeFromId(idInput)
            addNum += dataInput.data.num
        }

    }
    dataNode.data.num = addNum;
    df.updateNodeDataFromId(nodeId, dataNode.data);
}