<script>
import { defineComponent, ref, onMounted, getCurrentInstance, nextTick } from 'vue';
import store from '../../store';
import Node from '../layouts/Node.vue';
import moveTitle from '../../helpers/moveTitle';
import Input from '../Input.vue';

export default defineComponent({
    components: {
    Node,
    Input
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

            moveTitle(nodeId.value)
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
        <Node node-title="Number">
            <Input v-model.number="num" @change="handleChange"/>
        </Node>
    </div>
</template>

<style>
    #drawflow .NumericLiteral .output::before {
        content: "Number";
    }
</style>