<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import store from '../../store';
import { assign } from '../../modules/ops';
import { registerStop } from '../../helpers/stopWatch';
import Node from '../layouts/Node.vue';
import moveTitle from '../../helpers/moveTitle';
import Input from '../Input.vue';
import { checkMounted, registerMounted } from '../../helpers/mountedNodes';

export default defineComponent({
    components: {
    Node,
    Input
},
    setup() {
        const el = ref(null);
        let df = null
        const nodeId = ref(0);
        const nodeData = ref({});
        const num = ref(0);
        const varName = ref("")
        const sharedState = reactive(store.state)
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        //TODO: fix reactivity chain
        // check if the value of its input changed
        const stop = watch(sharedState, async () => {
            await nextTick()
            assign(df, nodeId.value, varName.value)
            // nodeData.value = df.getNodeFromId(nodeId.value)
            // num.value = nodeData.value.data.num;
        })

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;
            varName.value = `var${nodeId.value}`

            nodeData.value.data.var = varName.value;
            df.updateNodeDataFromId(nodeId.value, nodeData.value.data);

            moveTitle(nodeId.value)

            if (!checkMounted(nodeId.value)) {
                registerMounted(nodeId.value)
                registerStop(nodeId.value, stop)
                store.addVar(nodeId.value ,varName.value)
                
            }
        });

        return {
            el,
            // num,
            varName,
        }
    },
})
</script>
//TODO: change Component
<template>
    <div ref="el">
        <Node node-title="Assign">
            <Input v-model="varName" type="text" readonly/>
        </Node>
    </div>
</template>

<style scoped>
    .hide {
        display: none;
    }
</style>