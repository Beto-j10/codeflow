import { reactive } from 'vue';
import { checkAllConnectedInputs, CheckClassOps } from './modules/checkConnections';

const store = {

    stateConnections: reactive({

    }),

    stateVars: reactive({
        vars: [],
    }),

    stateModules: reactive({
        modules: [],
    }),

    printStates(){
        console.log("stateConnections", this.stateConnections)
        console.log("stateVars", this.stateVars.vars)
        console.log("stateModules", this.stateModules.modules)
    },

    initializeInputValues(nodeID, ...values) {
        this.stateConnections[nodeID] = [{ run: false }, {}]
        values.forEach((value, index) => {
            const input = `input_${index + 1}`
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

                // check if the connection is connected to a node that is an ops node
                if (CheckClassOps(connection.node, df)) {                
                    const input = connection.output
                    checkAllConnectedInputs(connection.node, df)
                    this.stateConnections[connection.node][1][input] = nodeData.data.num
                }
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

    updateConnectionsForConnectionCreated(ids, editor) {
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

    removeVar(varIndex) {
        this.stateVars.vars.splice(varIndex, 1)
    },

    addModule(name, id, df) {
        df.addModule(name)
        this.stateModules.modules.push({ name, id })
    },

    addModuleElse(name, id, df) {
        df.addModule(name)
        const moduleIndex = this.stateModules.modules.findIndex((module) => module.id === id)
        this.stateModules.modules[moduleIndex].else = name
        this.stateModules.modules.push({ name, id: `${id}.1` })
    },

    checkModuleElse(name){
        const moduleIndex = this.stateModules.modules.findIndex((module) => module.name === name)
        if (moduleIndex !== -1) {
            return true
        }
        return false
    },

    removeModule(moduleIndex, editor) {
        editor.removeModule(this.stateModules.modules[moduleIndex].name)
        this.stateModules.modules.splice(moduleIndex, 1)
    },

    removeModuleELse(id, editor) {
        const moduleIndex = this.stateModules.modules.findIndex((module) => module.id === id)
        editor.removeModule(this.stateModules.modules[moduleIndex].else)
        const moduleElseIndex = this.stateModules.modules.findIndex((module) => module.id === `${id}.1`)
        this.stateModules.modules.splice(moduleElseIndex, 1)
        delete this.stateModules.modules[moduleIndex].else
    },

    deleteState(id, editor) {
        const varIndex = this.stateVars.vars.findIndex(({ data }) => data.idParent === id)
        const moduleIndex = this.stateModules.modules.findIndex((module) => module.id === id)

        // not a var and not a module
        if (varIndex === -1 && moduleIndex === -1) {
            return
        }

        // is a var
        if (varIndex !== -1 && moduleIndex === -1) {
            this.removeVar(varIndex)
            return
        }

        // is module For
        if (varIndex !== -1 && moduleIndex !== -1) {
            this.removeVar(varIndex)
            this.removeModule(moduleIndex, editor)
            return

        }

        // is module If
        // check if module has else
        if (this.stateModules.modules[moduleIndex]?.else) {
            this.removeModuleELse(id, editor)
        }
        this.removeModule(moduleIndex, editor)
    },

}

export default store