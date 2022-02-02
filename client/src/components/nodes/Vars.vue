<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import store from '../../store';
import { registerStop } from '../../helpers/stopWatch';
import moveTitle from '../../helpers/moveTitle';
import Node from '../layouts/Node.vue';
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
        const num = ref(0)
        const sharedState = reactive(store.state)
        const nodeName = ref('')
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        // check if the value of one of its inputs changed
        const stop = watch(sharedState, () => {
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;
        })

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData.value = df.getNodeFromId(nodeId.value)
            nodeName.value = nodeData.value.name;
            num.value = nodeData.value.data.num;

            moveTitle(nodeId.value)
            if (!checkMounted(nodeId.value)) {
                registerMounted(nodeId.value)
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