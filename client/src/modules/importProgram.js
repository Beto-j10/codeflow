import { data1 } from "../components/data";
import { data2 } from "../components/data";
import { data3 } from "../components/data";
import { data4 } from "../components/data";
import { splitModuleName } from "../helpers/splitModuleName";
import store from "../store";

export function importProgram(editor) {
    // Object.keys(data1.drawflow).forEach(key => {
    //     if (key !== 'Home') {
    //         const { moduleType, moduleID } = splitModuleName(key);
    //         if (moduleType !== 'else') {
    //             // store.addModule(moduleType, moduleID, editor);
    //         } else {
    //             store.addModuleElse(moduleType, moduleID, editor);
    //         }
    //     }
    // });
    //TODO: delete setIsImport. This doesn't allow to create new modules
    store.setIsImport(true)
    editor.import(data4)

    // editor.changeModule('If7')
}