<template>
    <el-container>
        <el-header class="header">
            <h3>CodeFlow</h3>
            <el-button type="primary" @click="exportEditor">Export</el-button>
        </el-header>
        <el-container class="container">
            <el-aside width="150px" class="column">
                <ul>
                    <li
                        v-for="n in listNodes"
                        :key="n"
                        draggable="true"
                        :data-node="n.item"
                        @dragstart="drag($event)"
                    >
                        <div class="node" :style="`background: ${n.color}`">{{ n.name }}</div>
                    </li>
                </ul>
            </el-aside>
            <el-main>
                <div
                    id="drawflow"
                    @drop.prevent="drop($event)"
                    @dragover.prevent="allowDrop($event)"
                ></div>
            </el-main>
        </el-container>
    </el-container>
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
<script>

import Drawflow from 'drawflow'
import styleDrawflow from 'drawflow/dist/drawflow.min.css'
import style from '../assets/style.css'
import { onMounted, shallowRef, h, getCurrentInstance, render, readonly, ref } from 'vue'
import Number from './nodes/Number.vue'
import NodeAdd from './nodes/Add.vue'
import NodeSub from './nodes/Subtraction.vue'


export default {
    name: 'drawflow',
    setup() {
        const listNodes = readonly([
            {
                name: 'Number',
                color: '#49494970',
                item: 'Num',
                input: 0,
                output: 1,
                data: {
                    num: 0,
                }
            },
            {
                name: 'Add',
                color: '#49433440',
                item: 'Addx',
                input: 2,
                output: 1,
                data: {
                    num: 0,
                }
            },
            {
                name: 'Subtraction',
                color: '#49433440',
                item: 'Sub',
                input: 2,
                output: 1,
                data: {
                    num: 0,
                }
            },
        ])

        const editor = shallowRef({})
        const dialogVisible = ref(false)
        const dialogData = ref({})
        const Vue = { version: 3, h, render };
        const internalInstance = getCurrentInstance()
        internalInstance.appContext.app._context.config.globalProperties.$df = editor;

        function exportEditor() {
            dialogData.value = editor.value.export();
            dialogVisible.value = true;
        }

        const drag = (ev) => {
            ev.dataTransfer.setData("node", ev.target.getAttribute('data-node'));
        }
        const drop = (ev) => {
            var data = ev.dataTransfer.getData("node");
            addNodeToDrawFlow(data, ev.clientX, ev.clientY);

        }
        const allowDrop = (ev) => {
        }

        function addNodeToDrawFlow(name, pos_x, pos_y) {
            pos_x = pos_x * (editor.value.precanvas.clientWidth / (editor.value.precanvas.clientWidth * editor.value.zoom)) - (editor.value.precanvas.getBoundingClientRect().x * (editor.value.precanvas.clientWidth / (editor.value.precanvas.clientWidth * editor.value.zoom)));
            pos_y = pos_y * (editor.value.precanvas.clientHeight / (editor.value.precanvas.clientHeight * editor.value.zoom)) - (editor.value.precanvas.getBoundingClientRect().y * (editor.value.precanvas.clientHeight / (editor.value.precanvas.clientHeight * editor.value.zoom)));

            const nodeSelected = listNodes.find(ele => ele.item == name);
            editor.value.addNode(name, nodeSelected.input, nodeSelected.output, pos_x, pos_y, name, nodeSelected.data, name, 'vue');

        }


        onMounted(() => {

            const id = document.getElementById("drawflow");
            editor.value = new Drawflow(id, Vue, internalInstance.appContext.app._context);
            editor.value.start();

            editor.value.registerNode('Num', Number, {}, {});
            editor.value.registerNode('Addx', NodeAdd, {}, {});
            editor.value.registerNode('Sub', NodeSub, {}, {});

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
        })

        return {
            exportEditor, listNodes, drag, drop, allowDrop, dialogVisible, dialogData
        }

    }
}

</script>
<style scoped>
.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #494949;
}
.container {
    min-height: calc(100vh - 100px);
}
.column {
    border-right: 1px solid #494949;
}
.column ul {
    padding-inline-start: 0px;
    padding: 10px 10px;
}
.column li {
    background: transparent;
}

.node {
    border-radius: 8px;
    border: 2px solid #494949;
    display: block;
    height: 60px;
    line-height: 40px;
    padding: 10px;
    margin: 10px 0px;
    cursor: move;
}
#drawflow {
    width: 100%;
    height: 100%;
    text-align: initial;
    background: #2b2c30;
    background-size: 20px 20px;
    background-image: radial-gradient(#494949 1px, transparent 1px);
}
</style>