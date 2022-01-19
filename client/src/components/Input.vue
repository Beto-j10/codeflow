<script>
    export default {
        props: {
            modelValue: {
                type: [Number, String],
                default: 0,
            },
            // type: {
            //     type: String,
            //     default: 'number',
            // },
        },
        emits: ['update:modelValue', 'change'],
        methods: {
            checkValue(event) {
                const value = event.target.value;
                const type = event.target.type;
                if (type === 'number') {
                    if (value !== '') {
                        this.$emit('update:modelValue', value);
                        this.$emit('change', +value);
                    }
                }else {
                    this.$emit('update:modelValue', value);
                    this.$emit('change', value);
                }
            }
        }
    }

</script>

<template>
    <input
        :value="modelValue"
        @input="checkValue"
        type="number"
    >
</template>

<style scoped>
input {
    color: var(--text-color-secondary);
    border: 1px solid var(--border-color);
    border-radius: var(--dfNodeBorderRadius);
    border-radius: var(--border-radius-input);
    line-height: 2;
    padding: 0 10px;
    caret-color: var(--text-color-secondary);
    font-size: var(--font-size-small);
    min-width: 100px;
    width: 100%;
}

/* Chrome, Safari, Edge, Opera */
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

/* Firefox */
input[type="number"] {
    -moz-appearance: textfield;
}

input:focus {
    outline: none;
}
</style>