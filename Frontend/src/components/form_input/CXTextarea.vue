<script setup>
import Textarea from "primevue/textarea";

const props = defineProps(['label', 'rows', 'placeholder']);
const model = defineModel();
</script>

<template>
  <div style="display: flex; align-items: center">
    <p style="margin-right: 10px; font-size: 18px; width: 20%">{{ props.label }}:</p>
    <Textarea v-model="model" :placeholder="props.placeholder" :rows="props.rows" style="width: 100%" autoResize />
  </div>
</template>

<style scoped>

</style>