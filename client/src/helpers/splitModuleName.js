/**
 * returns the module type and the module ID
 * @param {string} Module
 * @returns {{moduleType: string, moduleID: string}}
 */
export function splitModuleName(moduleName) {
    const splitModuleName = moduleName.split(/(if|else|for)/i)
    const moduleType = splitModuleName[1].toLowerCase()
    const moduleID = splitModuleName[2]
    return { moduleType, moduleID }
}