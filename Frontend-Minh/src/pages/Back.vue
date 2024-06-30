<script setup>
const props = defineProps(['navigate'])
</script>

<template>
  <div id="back-container" @click="props.navigate">
    <i id="back-icon" class="pi pi-arrow-left"></i>
    <p id="back-text">Back</p>
  </div>
</template>

<style scoped>
#back-container {
  display: flex;
  align-items: center;
}

#back-icon {
  font-size: 30px;
}

#back-text {
  margin-left: 10px;
  font-size: 35px;
  margin-bottom: 4px;
}
</style>