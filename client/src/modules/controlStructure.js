// export function forStructure(df, nodeId) {
//     let nodeData = {}
//     nodeData = df.getNodeFromId(nodeId)

//     const inputs = nodeData.inputs
//     const input1 = inputs.input_1.connections.at(0)
//     const input2 = inputs.input_2.connections.at(0)
//     const input3 = inputs.input_3.connections.at(0)
//     let inputValues = []
    
//     if (input1 && input2 && input3) {
//         for (const input in inputs) {
//             const inputId = inputs[input].connections[0].node
//             const inputData = df.getNodeFromId(inputId)
//             inputValues.push(inputData.data.num) //adds the value of the input to the array
//         }
//         transform("for", inputValues)
//     }
// }