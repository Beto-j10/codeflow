<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import NodeHeader from './NodeHeader.vue';
import store from '../../store';
import { multiplication } from '../../modules/ops';
import { registerStop } from '../../helpers/stopWatch';

export default defineComponent({
    components: {
        NodeHeader
    },
    setup() {
        const el = ref(null);
        let df = null
        const nodeId = ref(0);
        const nodeData = ref({});
        const num = ref(0)
        const sharedState = reactive(store.state)
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        // check if the value of one of its inputs changed
        const stop = watch(sharedState, () => {
            multiplication(df, nodeId.value)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;
        })

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;

            registerStop(nodeId.value, stop)
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
        <NodeHeader title="Multiplication" />
        <el-input-number v-model="num" :controls="false" df-num />
    </div>
</template>