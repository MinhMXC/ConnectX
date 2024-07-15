<script setup>
import DatePicker from "primevue/datepicker";

const props = defineProps(['label']);
const model = defineModel();
</script>

<template>
  <div style="display: flex; align-items: center">
    <p style="margin-right: 10px; font-size: 18px; width: 25%">{{ props.label }}:</p>
    <DatePicker v-model="model" :placeholder="props.label" dateFormat="dd/mm/yy"
                input-class="p-inputtext-lg" style="width: 100%" />
  </div>
</template>

<style scoped>

</style>