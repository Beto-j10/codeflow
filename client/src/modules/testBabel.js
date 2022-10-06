import * as babel from '@babel/core';
import traverse from "@babel/traverse";
import * as t from "@babel/types";
import { parse } from "@babel/parser";
import generate from '@babel/generator';
import template from '@babel/template';

import compiler from '../service/compiler';

export default function fun() {
    const code = `
    
    let varrr = 69;
    for (varrr; varrr < 12; varrr++) {
        const max = 6
        let element = 0
        element += varrr;
    }
    for (let i = 4; i < 12; i++) {
        const max = 6
        let element = 0
        element += i;
    }
    const op2 = 111 + 222 + 444;
    const num = 888 + 369;
    let op3 = num + 987;
    const flag = true
    if (flag === 5) {
        const op = 88 * 44 + 22 + 11 * 66 + 77; 
    }
    `;

    const astP = template.ast(code)
    console.log("####", JSON.stringify(astP, null, 2), "#####")


    const ast = parse(code)
    // console.log("####", JSON.stringify(ast, null, 2), "#####")


    // traverse(ast, {
    //     enter(path) {
    //         switch (path.node.type) {
    //             case 'VariableDeclarator':
    //                 console.log("I'm a: ", path.node.id.name)
    //                 break;
    //             case 'ForStatement':
    //                 const forMin = path.node.init.declarations[0].init.value
    //                 const forMax = path.node.test.right.value
    //                 console.log("I'm a: ", forMin)
    //                 console.log("I'm a: ", forMax)
    //                 transformToCode('for', forMin, forMax)
    //                 path.skip()
    //                 break;
    //             case 'IfStatement':
    //                 console.log("I'm a: ", path.node.type)
    //                 break;
    //         }
    //     }
    // });

    const output = generate(ast, code);
    // console.log(output.code); // 'const x = 1;'

}

function transformToCode(type, forMin, forMax) {
    const code = `for i in range(${forMin}, ${forMax}):print (i)`
    console.log("##Pyhton: ", code)
    const dataCompiler = {
        //TODO: add environment vars
        script:         code,
        language:       'python3',
        versionIndex:   '3'
    }
    // compiler('http://localhost:5000/v1/api/compiler',dataCompiler)
        // .then(response => console.log(response))
}