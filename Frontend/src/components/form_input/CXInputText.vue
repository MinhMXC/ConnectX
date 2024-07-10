<script setup>
import InputText from "primevue/inputtext";

const props = defineProps(['label']);
const model = defineModel();
</script>

<template>
  <div style="display: flex; align-items: center">
    <p style="margin-right: 10px; font-size: 18px; width: 25%">{{ props.label }}:</p>
    <InputText v-model="model" :placeholder="props.label" size="large" style="width: 100%" />
  </div>
</template>

<style scoped>

</style>