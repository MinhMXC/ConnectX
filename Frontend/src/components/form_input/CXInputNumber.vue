<script setup>
import InputNumber from "primevue/inputnumber";

const props = defineProps(['label', 'prefix', 'format'])
const model = defineModel()
</script>

<template>
  <div style="display: flex; align-items: center">
    <p style="margin-right: 10px; font-size: 18px; width: 25%">{{ props.label }}:</p>
    <InputNumber v-model="model" :placeholder="props.label" :prefix="props.prefix" :format="props.format"
                 input-class="p-inputtext-lg" style="width: 100%" />
  </div>
</template>

<style scoped>

</style>