<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import store from '../../store';
import { registerStop } from '../../helpers/stopWatch';
import Node from '../layouts/Node.vue';
import moveTitle from '../../helpers/moveTitle';
import Input from '../Input.vue';
import { checkMounted, registerMounted } from '../../helpers/mountedNodes';
import { checkAllConnectedInputs } from '../../modules/checkConnections'

export default defineComponent({
    components: {
        Node,
        Input
    },
    setup() {
        const el = ref(null);
        const nodeId = ref(0);
        const num = ref(0);
        const varName = ref("")
        let df = null
        let nodeData = {};
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            if (!checkMounted(nodeId.value)) {
                store.initializeInputValues(nodeId.value, 0)
            }
            const sharedState = reactive(store.stateConnections[nodeId.value])
            varName.value = `Var${nodeId.value}`
            if (!checkMounted(nodeId.value)) {
                registerMounted(nodeId.value)

                const stop = watch(sharedState, () => {
                    console.log("sharedState Assign: ", sharedState)
                    if (sharedState[0].run) {
                        store.addVar(nodeId.value, varName.value, sharedState[1].input_1)
                        nodeData.data.num = sharedState[1].input_1
                    } else {
                        nodeData.data.num = 0
                        store.removeVar(nodeId.value)
                    }
                    df.updateNodeDataFromId(nodeId.value, nodeData.data);
                    num.value = nodeData.data.num;
                })

                registerStop(nodeId.value, stop)

            }

            nodeData = df.getNodeFromId(nodeId.value)
            nodeData.data.var = varName.value
            num.value = nodeData.data.num;
            moveTitle(nodeId.value)
            df.updateNodeDataFromId(nodeId.value, nodeData.data);
        });

        return {
            el,
            varName,
        }
    },
})
</script>
//TODO: change Component
<template>
    <div ref="el">
        <Node node-title="Assign">
            <Input v-model="varName" type="text" readonly />
        </Node>
    </div>
</template>

<style scoped>
.hide {
    display: none;
}
</style>