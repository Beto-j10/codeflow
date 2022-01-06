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
    }else if (nodeDataInput.class !== "assign") {
            editor.removeSingleConnection(ids.output_id, ids.input_id, ids.output_class, ids.input_class)
            console.log("else", ids.output_class)
    }
}

export function checkAssigment(df) {

}