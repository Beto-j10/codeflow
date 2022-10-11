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
        const selected = ref('')
        const sharedState = reactive(store.stateVars)
        let nodeData = {};
        let df = null
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        function updateValueConnections() {
            const isAllConnectedOutputs = checkAllConnectedOutputs(nodeId.value, df)
            if (isAllConnectedOutputs) {
                store.updateConnections(nodeId.value, df);
            }
        }

        function updateSelectedValue() {
            nodeData.data.num = store.stateVars[selected.value].num
            df.updateNodeDataFromId(nodeId.value, nodeData.data);
            updateValueConnections()
        }

        function handleChange(e) {
            if (selected.value !== "") {
                updateSelectedValue()
            }
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            nodeData = df.getNodeFromId(nodeId.value)

            // const sharedState = reactive(store.stateConnections[nodeData.data.idParent])
            // sharedState = reactive(store.stateVars)
            nodeName.value = nodeData.name;
            // nodeData.data.num = sharedState[1].input_1
            num.value = nodeData.data.num;
            df.updateNodeDataFromId(nodeId.value, nodeData.data);
            moveTitle(nodeId.value)

            if (!checkMounted(nodeId.value)) {
                registerMounted(nodeId.value)

                const stop = watch(sharedState, () => {
                    // if (sharedState[0].run) {
                    //     nodeData.data.num = sharedState[1].input_1
                    // } else {
                    //     nodeData.data.num = 0
                    // }
                    // df.updateNodeDataFromId(nodeId.value, nodeData.data);
                    // num.value = nodeData.data.num;

                    // check if node connection or node was removed from its parent, if yes then blank selector
                    if (!Object.hasOwn(store.stateVars, selected.value)) {
                        selected.value = ""
                        nodeData.data.num = 0
                        df.updateNodeDataFromId(nodeId.value, nodeData.data);
                        updateValueConnections()
                    } else {
                        updateSelectedValue()
                    }
                })

                registerStop(nodeId.value, stop)

            }

        });

        return {
            el,
            num,
            selected,
            sharedState,
            handleChange
        }
    },
})
</script>

<template>
    <div ref="el">
        <!-- <Node node-title="Var" isEmpty> -->
        <Node node-title="Var">
            <el-select v-model="selected" class="m-2" placeholder="Select" size="small" @change="handleChange">
                <el-option v-for="item in sharedState" :key="item.name" :label="item.name" :value="item.idParent" />
            </el-select>
            <!-- <span>{{sharedState[selected].num}}</span> -->
            <!-- <Input v-model.number="value" readonly /> -->
        </Node>
    </div>
</template>

<style>
#drawflow .Identifier .output::before {
    content: "Number";
}
</style>