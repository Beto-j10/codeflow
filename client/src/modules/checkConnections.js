import store from "../store"

export function checkConnections(ids, editor) {
    // console.log("ids: ", ids)
    // console.log("Editor: ", editor)
    const nodeDataOutput = editor.getNodeFromId(ids.output_id)
    console.log("nodeDataOutput: ", nodeDataOutput.inputs)
    if (Object.keys(nodeDataOutput.inputs).length === 0) {
        console.log("nodeDataOutput.inputs is undefined")
    }
    const nodeDataInput = editor.getNodeFromId(ids.input_id)

    if (nodeDataOutput.class !== "ops") {
        return
    } else if (nodeDataInput.class !== "assign") {
        editor.removeSingleConnection(ids.output_id, ids.input_id, ids.output_class, ids.input_class)
        console.log("else", ids.output_class)
    }
}

export function checkConnected(nodeID, editor) {
    const isAllInputConnected = checkAllConnectedInputs(nodeID, editor)
    if (isAllInputConnected) {
        store.updateConnections(nodeID, isAllInputConnected)
    }
    console.log(store.stateConnections)
}

/**
 * check if all inputs are connected and set the run state of the node to true if they are
 * @param {number} nodeID 
 * @param {object} editor 
 */

export function checkAllConnectedInputs(nodeID, editor) {
    const nodeData = editor.getNodeFromId(nodeID)
    const inputs = nodeData.inputs
    for (const input of Object.keys(inputs)) {
        if (!inputs[input].connections.at(0)) {
            return false
        }
    }
    store.updateRun(nodeID, true)
    return true
}

export function checkAllConnectedOutputs(nodeID, editor) {
    const nodeData = editor.getNodeFromId(nodeID)
    const outputs = nodeData.outputs
    for (const output of Object.keys(outputs)) {
        if (!outputs[output].connections.at(0)) {
            return false
        }
    }
    return true
}

export function checkAssigment(df) {

}

export function checkNodeRemoved(id, editor) {
    const nodeData = editor.getNodeFromId(id)
    if (nodeData.class === "forStatement") {
        store.deleteVar(id)
    }

}

export function CheckClassOps(id, editor) {
    const nodeData = editor.getNodeFromId(id)
    const nodeClass = nodeData.class.split(" ").at(-1)
    if (nodeClass === "ops") {
        return true
    }
    return false
}