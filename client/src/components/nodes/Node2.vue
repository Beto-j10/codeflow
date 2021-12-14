<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import NodeHeader from './NodeHeader.vue';
import store from '../store';
import { add } from '../operations';

export default defineComponent({
    components: {
        NodeHeader
    },
    setup() {
        const el = ref(null);
        let df = null
        const nodeId = ref(0);
        const dataNode = ref({});
        const num = ref(0)
        const sharedState = reactive(store.state)
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
            num.value = dataNode.value.data.num;
        });

        watch(sharedState, () => {
            add(df, nodeId.value)
            dataNode.value = df.getNodeFromId(nodeId.value)
            num.value = dataNode.value.data.num;
        })

        return {
            el,
            num,
        }
    },
})
</script>

<template>
    <div ref="el">
        <NodeHeader title="Add" />
        <el-input-number v-model="num" disabled :controls="false" df-num />
    </div>
</template>