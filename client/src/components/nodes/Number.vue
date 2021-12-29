<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick } from 'vue';
import NodeHeader from './NodeHeader.vue';
import store from '../../store';

export default defineComponent({
    components: {
        NodeHeader
    },
    data() {
        return {
            sharedState: store.state
        }
    },
    setup() {
        const el = ref(null);
        let df = null
        const nodeId = ref(0);
        const nodeData = ref({});
        const num = ref(0)
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const handleChange = (value) => {
            nodeData.value.data.num = value;
            df.updateNodeDataFromId(nodeId.value, nodeData.value.data);
            store.updateState()
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;
        });

        return {
            el,
            num,
            handleChange,
        }
    },
})
</script>

<template>
    <div ref="el">
        <NodeHeader title="Number" />
        <el-input-number v-model="num" :controls="false" @change="handleChange" df-num />
    </div>
</template>