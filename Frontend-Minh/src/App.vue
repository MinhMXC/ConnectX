<script setup>
</script>

<template>
  <RouterView :key="$route.fullPath" />
</template>

<style>
#app {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
