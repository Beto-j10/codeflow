export function transformToCode(node) {
    const language = 'python'
    switch (language) {
        case "javascript":
            return transformToCodeJavascript(node);
        case "python":
            return transformToCodePython(node);
        default:
            return transformToCodeJavascript(node);
    }
}

function transformToCodePython(node) {
    switch (node.type) {
        case "Program":
        case "BlockStatement":
            let body = ''
            for (const child of node.body) {
                body += `${transformToCode(child)}\n`
            }
            return body;

        case "NumericLiteral":
            return node.value;
        case "BinaryExpression":
            return `${transformToCode(node.left)} ${node.operator} ${transformToCode(node.right)}`;
        case "VariableDeclarator":
            return `${node.idType.name} = ${transformToCode(node.init)}`;
        case "Identifier":
            return node.name;
        case "ForStatement":
            return `for ${transformToCode(node.init.name)} in range (${transformToCode(node.test.left)}, ${transformToCode(node.update.right)}, ${node.update.value.value}):\n${transformToCode(node.body)}`;
        case "IfStatement":
            return `if ${transformToCode(node.test)}:\n ${transformToCode(node.consequent)} ${node.alternate ? `else:\n\t${transformToCode(node.alternate)}` : ''}`;
        default:
            return "";

    }
}

function transformToCodeJavascript(node) {
    switch (node.type) {
        case "Program":
            for (const child of node.body) {
                transformToCode(child)
            }
            break;
        case "NumericLiteral":
            return node.value;
        case "BinaryExpression":
            return `(${transformToCode(node.left)} ${node.operator} ${transformToCode(node.right)})`;
        case "VariableDeclarator":
            return `${node.idType.name} = ${transformToCode(node.init)}`;
        case "Identifier":
            return node.name;
        case "ForStatement":
            return `for (${transformToCode(node.test.left)} ${node.test.operator} ${transformToCode(node.test.right)}; ${transformToCode(node.update.left)} ${node.update.operator} ${transformToCode(node.update.right)}; ${transformToCode(node.body)})`;
        case "IfStatement":
            return `if (${transformToCode(node.test)}) {${transformToCode(node.consequent)}} else {${transformToCode(node.alternate)}}`;
        case "BlockStatement":
            return `{${transformToCode(node.body)}}`;
        default:
            return "";

    }
}