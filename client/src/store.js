import { reactive } from 'vue';
import { checkAllConnectedInputs } from './modules/checkConnections';

const store = {

    stateConnections: reactive({

    }),

    stateVars: reactive({
        vars: [],
    }),

    stateModules: reactive({
        modules: [],
    }),

    initializeInputValues(nodeID, ...values) {
        this.stateConnections[nodeID] = [{run:false},{}]
        values.forEach((value, index) => {
            const input = `input_${index+1}`
            this.stateConnections[nodeID][1][input] = value
        })
        console.log("initialize", this.stateConnections)
    },

    deleteNodeConnections(nodeID) {
        delete this.stateConnections[nodeID];
    },

    updateConnections(nodeID, df) {
        const nodeData = df.getNodeFromId(nodeID)
        const outputs = nodeData.outputs
        Object.keys(outputs).forEach(output => {
            outputs[output].connections.forEach(connection => {
                const input = connection.output
                checkAllConnectedInputs(connection.node, df)
                this.stateConnections[connection.node][1][input] = nodeData.data.num
            })
        })
    },

    updateRun(nodeID, state) {
        this.stateConnections[nodeID][0].run = state
    },

    updateConnectionsForConnectionRemoved(nodeID, input) {
        this.updateRun(nodeID, false)
        this.stateConnections[nodeID][1][input] = 0
    },

    updateConnectionsForConnectionCreated(ids, editor){
        const nodeDataOutput = editor.getNodeFromId(ids.output_id)
        this.stateConnections[ids.input_id][1][ids.input_class] = nodeDataOutput.data.num
        checkAllConnectedInputs(ids.input_id, editor)
    },

    addVar(id, name) {
        this.stateVars.vars.push({
            name: name,
            color: "#49433440",
            item: name,
            input: 0,
            output: 1,
            data: {
                num: 0,
                idParent: id,
            },
            class: "Identifier ops",
        })
    },

    deleteVar(id, editor) {
        const varIndex = this.stateVars.vars.findIndex(({ data }) => data.idParent === id)
        if (varIndex > -1) {
            const varName = this.stateVars.vars[varIndex].name
            if (varName.startsWith("F")) {
                this.stateVars.vars.splice(varIndex, 1)
                this.removeModule(`For${id}`, editor)
            } else if (varName.startsWith("A")) {
                this.stateVars.vars.splice(varIndex, 1)
            }
        }
    },

    addModule(name, df) {
        df.addModule(name)
        this.stateModules.modules.push(name)
    },

    removeModule(name, editor) {
        this.stateModules.modules.splice(this.stateModules.modules.indexOf(name), 1)
        editor.removeModule(name)
    }

}

export default store