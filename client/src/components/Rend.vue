<script>
    import store from "../store";
    export default {
        data() {
            return {
                sharedState: store
            }
        },
        methods: {
            handleClick(){
                store.addVar(6)
            }
        }
    }
</script>

//TODO: change key
<template>
    <li v-for="varData of sharedState.stateVars.vars" :key="varData.id">
        {{varData}}
    </li>
    <li v-for="varData of sharedState.stateModules.modules" :key="varData">
        {{varData}}
    </li>
    <button @click="handleClick"></button>
</template>