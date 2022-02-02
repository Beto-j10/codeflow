import { reactive } from 'vue';

const store = {

    state: reactive({
        stateToggle: false,
    }),

    stateVars: reactive({
        vars: [],
    }),

    stateModules: reactive({
        modules: [],
    }),


    updateState() {
        this.state.stateToggle = !this.state.stateToggle
    },

    addVar(id, name) {
        this.stateVars.vars.push({
            name: name,
            color: "#49433440",
            item: name,
            input: 0,
            output: 1,
            data: {
                idParent: id,
            },
            class: "Identifier",
        })
    },

    deleteVar(id, editor) {
        const varIndex = this.stateVars.vars.findIndex(({ data }) => data.idParent === id)
        if (varIndex > -1) {
            const varName = this.stateVars.vars[varIndex].name
            if (varName.startsWith("F")) {
                this.stateVars.vars.splice(varIndex, 1)
                this.removeModule(`for_${id}`, editor)
            }else if (varName.startsWith("A")) {
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