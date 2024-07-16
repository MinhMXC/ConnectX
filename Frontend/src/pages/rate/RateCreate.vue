<script setup>

import router from "@/router.js";
import Back from "@/pages/Back.vue";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
import RateForm from "@/components/form/RateForm.vue";
</script>

<template>
  <FullScreenCenteredContainer>
    <div style="margin-bottom: 300px; min-width: 900px;">
      <Back :navigate="() => router.back()" />
      <h1 id="creating-rate-text">Creating Rate</h1>

      <RateForm />
    </div>
  </FullScreenCenteredContainer>
</template>

<style scoped>
#creating-rate-text {
  margin-top: 20px;
  margin-bottom: 50px;
}
</style>