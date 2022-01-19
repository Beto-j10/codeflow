<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick, watch, reactive } from 'vue';
import store from '../../store';
import { add } from '../../modules/ops';
import { registerStop } from '../../helpers/stopWatch';
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
        let df = null
        const nodeId = ref(0);
        const nodeData = ref({});
        const num = ref(0)
        const sharedState = reactive(store.state)
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        // check if the value of one of its inputs changed
        const stop = watch(sharedState, () => {
            add(df, nodeId.value)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;
        })

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData.value = df.getNodeFromId(nodeId.value)
            num.value = nodeData.value.data.num;

            registerStop(nodeId.value, stop)
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