<script setup>
import Password from "primevue/password";

const props = defineProps(['label', 'feedback']);
const model = defineModel();
</script>

<template>
  <div style="display: flex; align-items: center">
    <p style="margin-right: 10px; font-size: 18px; width: 25%">{{ props.label }}:</p>
    <Password v-model="model" :placeholder="props.label" style="width: 100%"
              toggleMask input-class="p-inputtext-lg" :feedback="props.feedback" />
  </div>
</template>

<style scoped>
.p-password {
  flex-direction: column;
}
</style>