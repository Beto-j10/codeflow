<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import NodeHeader from './NodeHeader.vue';
import store from '../../store';
import { assign } from '../../modules/ops';
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

        //TODO: fix reactivity chain
        // check if the value of its input changed
        const stop = watch(sharedState, async () => {
            await nextTick()
            assign(df, nodeId.value)
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
//TODO: change Component
<template>
    <div ref="el" class="assign">
        <NodeHeader title="Assign" />
        <!-- <el-input-number v-model="num" :controls="false" df-num /> -->
        <div class="assign__container">
            <span class="assign__text">{{ num }}</span>
        </div>
    </div>
</template>

<style scoped>
.assign__container {
    /* background-color: #fff; */
    display: flex;
    align-items: center;
    justify-content: center;
    /* width: 100%;
    height: 100%; */
}
.assign__text {
    display: block;
    margin: 0 auto;
    color: #fff;
    text-align: center;
    font-size: 1.2rem;
}
</style>