import { reactive } from 'vue';

const store = {

    state: reactive({
        stateToggle: false
    }),

    updateState() {
        this.state.stateToggle = !this.state.stateToggle
    },
}

export default store