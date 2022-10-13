import { checkAllConnectedOutputs } from '../modules/checkConnections'
import store from '../store'

const numberNodes = []

export function debugStartNumberNodes() {
    console.warn("Print startNumberNodes: ", numberNodes)
}

export function hasNumberNodes() {
    return numberNodes.length !== 0 ? true : false
}

export function registerNumberNode(nodeId) {
    numberNodes.push(nodeId)
}

function unregisterNumberNode() {
        numberNodes.splice(0);
}

export function startNumberNode(editor) {
    numberNodes.forEach(nodeId => {
    console.log("------------------------------------", nodeId)

        const isAllConnectedOutputs = checkAllConnectedOutputs(nodeId, editor)
        if (isAllConnectedOutputs) {
            store.updateConnections(nodeId, editor);
        }

    })
    unregisterNumberNode()
}