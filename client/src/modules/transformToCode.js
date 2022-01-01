function transform(type, values) {
    switch (type) {
        case "for":
            const code = `for i in range(${values[0]}, ${values[1]},${values[2]}):print (i)`
            break;
        default:
            break;
    }
}