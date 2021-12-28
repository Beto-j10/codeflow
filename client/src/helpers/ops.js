export function add(df, nodeId) {

    let nodeData = {}
    nodeData = df.getNodeFromId(nodeId)

    const inputs = nodeData.inputs
    let addNum = 0

    if (inputs.input_1.connections.at(0) && inputs.input_2.connections.at(0)) {
        for (const input in inputs) {
            const idInput = inputs[input].connections[0].node
            const dataInput = df.getNodeFromId(idInput)
            addNum += dataInput.data.num
        }

    }
    nodeData.data.num = addNum;
    df.updateNodeDataFromId(nodeId, nodeData.data);
}