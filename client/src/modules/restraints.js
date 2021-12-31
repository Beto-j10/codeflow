export function checkConnections(ids, editor) {
    // console.log("ids: ", ids)
    // console.log("Editor: ", editor)
    const nodeData = editor.getNodeFromId(ids.output_id)
    console.log("nodeData: ", nodeData.name)
    const nodeData2 = editor.getNodeFromId(ids.input_id)
    console.log("nodeData2: ", nodeData2.name)
}

export function checkAssigment(df) {

}