<script setup>
import Select from "primevue/select";

const props = defineProps(['label', 'options'])
const model = defineModel()
</script>

<template>
  <div style="display: flex; align-items: center">
    <p style="margin-right: 10px; font-size: 18px; width: 20%">{{ props.label }}:</p>
    <Select v-model="model" :options="props.options" :placeholder="props.label"
            input-class="p-inputtext-lg" style="width: 100%" @input="event => { props.model.value =  event.target.value }" />
  </div>
</template>

<style scoped>

</style>