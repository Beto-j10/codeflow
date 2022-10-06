<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import store from '../../store';
// import { multiplication } from '../../modules/ops';
import { registerStop } from '../../helpers/stopWatch';
import Node from '../layouts/Node.vue';
import Input from '../Input.vue';
import moveTitle from '../../helpers/moveTitle';
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

        function multiplication(values) {
            const result = values.reduce((acc, val) => acc * val, 1);
            nodeData.data.num = Math.round(result * 100) / 100
            df.updateNodeDataFromId(nodeId.value, nodeData.data);
            num.value = nodeData.data.num;
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
                    console.log("sharedState Multiplication: ", sharedState)
                    if (sharedState[0].run ) {
                        multiplication(Object.values(sharedState[1]))
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
        <Node node-title="Multiplication" width="150px">
            <Input v-model.number="num" readonly />
        </Node>
    </div>
</template>