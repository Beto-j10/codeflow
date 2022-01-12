<script>

import Drawflow from 'drawflow'
import 'drawflow/dist/drawflow.min.css'
import '../styles/drawflow.css'
import { onMounted, shallowRef, h, getCurrentInstance, render, readonly, ref, nextTick } from 'vue'
import Number from './nodes/Number.vue'
import NodeAdd from './nodes/Add.vue'
import NodeSub from './nodes/Subtraction.vue'
import NodeMult from './nodes/Multiplication.vue'
import NodeDiv from './nodes/Division.vue'
import NodeAssign from './nodes/Assign.vue'
import NodeFor from './nodes/For.vue'
import NodeBlock from './nodes/Block.vue'

import store from '../store'
import { checkConnections } from '../modules/constraints'
import { stopWatch } from '../helpers/stopWatch'
import { transformToAST } from '../modules/transformToAST'
import BaseLayout from './layouts/BaseLayout.vue'


export default {
    name: "drawflow",
    setup() {
        const listNodes = readonly([
            {
                name: "Number",
                color: "#49494970",
                item: "Num",
                input: 1,
                output: 1,
                data: {
                    num: 0,
                },
                class: "NumericLiteral",
            },
            {
                name: "Add",
                color: "#49433440",
                item: "Add",
                input: 2,
                output: 1,
                data: {
                    num: 0,
                    operator: "+",
                },
                class: "BinaryExpression",
            },
            {
                name: "Subtraction",
                color: "#49433440",
                item: "Sub",
                input: 2,
                output: 1,
                data: {
                    num: 0,
                    operator: "-",
                },
                class: "BinaryExpression",
            },
            {
                name: "Multiplication",
                color: "#49433440",
                item: "Mult",
                input: 2,
                output: 1,
                data: {
                    num: 0,
                    operator: "*",
                },
                class: "BinaryExpression",
            },
            {
                name: "Division",
                color: "#49433440",
                item: "Div",
                input: 2,
                output: 1,
                data: {
                    num: 0,
                    operator: "/",
                },
                class: "BinaryExpression",
            },
            {
                name: "Assign",
                color: "#49433440",
                item: "Assign",
                input: 1,
                output: 1,
                data: {
                    num: 0,
                },
                class: "VariableDeclarator",
            },
            {
                name: "For",
                color: "#49433440",
                item: "For",
                input: 3,
                output: 1,
                data: {
                    operator: "<",
                },
                class: "ForStatement",
            },
            {
                name: "Block",
                color: "#49433440",
                item: "Block",
                input: 0,
                output: 1,
                data: {},
                class: "BlockStatement",
            },
        ]);
        const editor = shallowRef({});
        const dialogVisible = ref(false);
        const dialogData = ref({});
        const Vue = { version: 3, h, render };
        const internalInstance = getCurrentInstance();
        internalInstance.appContext.app._context.config.globalProperties.$df = editor;
        function exportEditor() {
            dialogData.value = editor.value.export();
            transformToAST(dialogData.value.drawflow.Home.data);
            dialogVisible.value = true;
        }
        const drag = (ev) => {
            ev.dataTransfer.setData("node", ev.target.getAttribute("data-node"));
        };
        const drop = (ev) => {
            var data = ev.dataTransfer.getData("node");
            addNodeToDrawFlow(data, ev.clientX, ev.clientY);
        };
        const allowDrop = (ev) => {
        };
        function addNodeToDrawFlow(name, pos_x, pos_y) {
            pos_x = pos_x * (editor.value.precanvas.clientWidth / (editor.value.precanvas.clientWidth * editor.value.zoom)) - (editor.value.precanvas.getBoundingClientRect().x * (editor.value.precanvas.clientWidth / (editor.value.precanvas.clientWidth * editor.value.zoom)));
            pos_y = pos_y * (editor.value.precanvas.clientHeight / (editor.value.precanvas.clientHeight * editor.value.zoom)) - (editor.value.precanvas.getBoundingClientRect().y * (editor.value.precanvas.clientHeight / (editor.value.precanvas.clientHeight * editor.value.zoom)));
            const nodeSelected = listNodes.find(ele => ele.item == name);
            editor.value.addNode(name, nodeSelected.input, nodeSelected.output, pos_x, pos_y, nodeSelected.class, nodeSelected.data, name, "vue");
        }
        onMounted(() => {
            const id = document.getElementById("drawflow");
            editor.value = new Drawflow(id, Vue, internalInstance.appContext.app._context);
            editor.value.start();
            editor.value.registerNode("Num", Number, {}, {});
            editor.value.registerNode("Add", NodeAdd, {}, {});
            editor.value.registerNode("Sub", NodeSub, {}, {});
            editor.value.registerNode("Mult", NodeMult, {}, {});
            editor.value.registerNode("Div", NodeDiv, {}, {});
            editor.value.registerNode("Assign", NodeAssign, {}, {});
            editor.value.registerNode("For", NodeFor, {}, {});
            editor.value.registerNode("Block", NodeBlock, {}, {});
            editor.value.on("nodeRemoved", (id) => {
                stopWatch(id);
            });
            editor.value.on("connectionCreated", (ids) => {
                checkConnections(ids, editor.value);
                store.updateState();
            });
            editor.value.on("connectionRemoved", async () => {
                await nextTick();
                store.updateState();
            });
            // editor.value.import(
            //     {
            //         "drawflow": {
            //             "Home": {
            //                 "data": {
            //                     "1": {
            //                         "id": 1,
            //                         "name": "Node2",
            //                         "data": {
            //                             "script": "(req,res) => {\n console.log(req);\n}"
            //                         },
            //                         "class": "Node2",
            //                         "html": "Node2",
            //                         "typenode": "vue",
            //                         "inputs": {
            //                             "input_1": {
            //                                 "connections": [
            //         I have data that I want to write to a file, and open a file dialog for the user to choose where to save the file. It would be great if it worked in all browsers, but it has to work in Chrome. I want to do this all client-side.
            //                                         "node": "6",
            //                                         "input": "output_1"
            //                                     }
            //                                 ]
            //                             }
            //                         },
            //                         "outputs": {
            //                             "output_1": {
            //                                 "connections": [
            //                                 ]
            //                             },
            //                             "output_2": {
            //                                 "connections": [
            //                                 ]
            //                             }
            //                         },
            //                         "pos_x": 600,
            //                         "pos_y": 80
            //                     },
            //                     "6": {
            //                         "id": 6,
            //                         "name": "Node1",
            //                         "data": {
            //                             "url": "localhost/add",
            //                             "method": "post"
            //                         },
            //                         "class": "Node1",
            //                         "html": "Node1",
            //                         "typenode": "vue",
            //                         "inputs": {
            //                         },
            //                         "outputs": {
            //                             "output_2": {
            //                                 "connections": [
            //                                     {
            //                                         "node": "5",
            //                                         "output": "input_2"
            //                                     }
            //                                 ]
            //                             }
            //                         },
            //                         "pos_x": 137,
            //                         "pos_y": 89
            //                     }
            //                 }
            //             }
            //         }
            //     }
            // )
        });
        return {
            exportEditor,
            listNodes,
            drag,
            drop,
            allowDrop,
            dialogVisible,
            dialogData
        };
    },
    components: { BaseLayout }
}

</script>

<template>
    <BaseLayout>
        <template #header>
            <h3 class="title">CodeFlow</h3>
            <el-button type="primary" @click="exportEditor">Export</el-button>
        </template>

        <template #sidebar>
            <ul class="list">
                <li
                    class="list__item"
                    v-for="n in listNodes"
                    :key="n"
                    draggable="true"
                    :data-node="n.item"
                    @dragstart="drag($event)"
                >
                    <!-- <div class="list__node" :style="`background: ${n.color}`">{{ n.name }}</div> -->
                    <div class="list__node">{{ n.name }}</div>
                </li>
            </ul>
        </template>

        <template #main>
            <div id="drawflow" @drop.prevent="drop($event)" @dragover.prevent="allowDrop($event)"></div>
        </template>
    </BaseLayout>

    <el-dialog v-model="dialogVisible" title="Export" width="50%">
        <span>Data:</span>
        <pre><code>{{ dialogData }}</code></pre>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogVisible = false">Cancel</el-button>
                <el-button type="primary" @click="dialogVisible = false">Confirm</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<style scoped>
.title {
    font-size: 2.4rem;
}

.list {
    list-style: none;
    padding: 0;
    margin: 0;
    border-right: 1px solid var(--border-color-dark);
}

.list__item {
    border-bottom: 1px solid var(--border-color-dark);
    cursor: move;
    padding: 12px 12px 12px var(--container-padding);
}

#drawflow {
    width: 100%;
    height: 100%;
    text-align: initial;
}
</style>