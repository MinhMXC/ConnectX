<script setup>
import {computed} from "vue";

  const props = defineProps(['status']);

  const style = computed(() => {
    return {
      fontWeight: 600,
      color: props.status === "Success" ? "green" : "red"
    };
  });
</script>

<template>
  <p :style="style">{{ props.status }}</p>
</template>

<style scoped>

</style>