<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import store from '../../store';
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
        const nodeName = ref('')
        let nodeData = {};
        let df = null
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData = df.getNodeFromId(nodeId.value)

            const sharedState = reactive(store.stateConnections[nodeData.data.idParent])
            nodeName.value = nodeData.name;
            nodeData.data.num = sharedState[1].input_1
            num.value = nodeData.data.num;
            df.updateNodeDataFromId(nodeId.value, nodeData.data);
            moveTitle(nodeId.value)

            if (!checkMounted(nodeId.value)) {
                registerMounted(nodeId.value)

                const stop = watch(sharedState, () => {
                    if (sharedState[0].run) {
                        nodeData.data.num = sharedState[1].input_1
                    } else {
                        nodeData.data.num = 0
                    }
                    df.updateNodeDataFromId(nodeId.value, nodeData.data);
                    num.value = nodeData.data.num;

                    const isAllConnectedOutputs = checkAllConnectedOutputs(nodeId.value, df)
                    if (isAllConnectedOutputs) {
                        store.updateConnections(nodeId.value, df);
                    }

                })

                registerStop(nodeId.value, stop)

            }

        });

        return {
            el,
            num,
            nodeName,
        }
    },
})
</script>

<template>
    <div ref="el">
        <Node :node-title="nodeName" isEmpty>
            <Input v-model.number="num" readonly />
        </Node>
    </div>
</template>

<style>
#drawflow .Identifier .output::before {
    content: "Number";
}
</style>