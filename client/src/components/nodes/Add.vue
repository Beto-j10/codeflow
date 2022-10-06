<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import store from '../../store';
// import { add } from '../../modules/ops';
import { registerStop } from '../../helpers/stopWatch';
import moveTitle from '../../helpers/moveTitle';
import Node from '../layouts/Node.vue';
import Input from '../Input.vue';
import { checkMounted, registerMounted } from '../../helpers/mountedNodes';
import { checkAllConnectedOutputs } from '../../modules/checkConnections';

export default defineComponent({
    components: {
        Node,
        Input
    },
    setup() {
        const el = ref(null);
        const nodeId = ref(0);
        const num = ref(0)
        let df = null
        let nodeData = {};
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;


        // check if the value of one of its inputs changed

        function add(values) {
            console.log("Values op Add: ", values)
            const result = values.reduce((acc, val) => acc + val, 0);
            nodeData.data.num = Math.round(result * 100) / 100
            df.updateNodeDataFromId(nodeId.value, nodeData.data);
            num.value = nodeData.data.num;
            console.log("Result opp Add: ", num.value)
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            if (!checkMounted(nodeId.value)) {
                store.initializeInputValues(nodeId.value, 0, 0)
            }
            const sharedState = reactive(store.stateConnections[nodeId.value])
            if (!checkMounted(nodeId.value)) {
                registerMounted(nodeId.value)

                const stop = watch(sharedState, () => {
                    console.log("sharedState Add: ", sharedState)
                    if (sharedState[0].run ) {
                        add(Object.values(sharedState[1]))
                    }else{
                        nodeData.data.num = 0
                        df.updateNodeDataFromId(nodeId.value, nodeData.data);
                        num.value = nodeData.data.num;
                    }
                    const isAllConnectedOutputs = checkAllConnectedOutputs(nodeId.value, df)
                    if (isAllConnectedOutputs) {
                        store.updateConnections(nodeId.value, df);
                    }
                })

                registerStop(nodeId.value, stop)

            }

            nodeData = df.getNodeFromId(nodeId.value)
            num.value = nodeData.data.num;
            moveTitle(nodeId.value)
        });
        return {
            el,
            num,
        }
    },
})
</script>

<template>
    <div ref="el">
        <Node node-title="Addition" width="150px">
            <Input v-model.number="num" readonly />
        </Node>
    </div>
</template>