<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import moveTitle from '../../helpers/moveTitle';
import Node from '../layouts/Node.vue';
import Input from '../Input.vue';

export default defineComponent({
    components: {
        Node,
        Input
    },
    setup() {
        const el = ref(null);
        const nodeId = ref(0);
        const nodeName = ref('')
        let nodeData = {};
        let df = null
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData = df.getNodeFromId(nodeId.value)

            nodeName.value = nodeData.name;
            moveTitle(nodeId.value)

        });

        return {
            el,
            nodeName,
        }
    },
})
</script>

<template>
    <div ref="el">
        <Node :node-title="nodeName" isEmpty></Node>
    </div>
</template>

