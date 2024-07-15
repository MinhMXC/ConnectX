<script setup>

import router from "@/router.js";
import Back from "@/pages/Back.vue";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
import QualificationForm from "@/components/form/QualificationForm.vue";
</script>

<template>
  <FullScreenCenteredContainer>
    <div style="margin-bottom: 300px; min-width: 900px;">
      <Back :navigate="() => router.back()" />
      <h1 id="creating-qualification-text">Creating Qualification</h1>

      <QualificationForm />
    </div>
  </FullScreenCenteredContainer>
</template>

<style scoped>
#creating-qualification-text {
  margin-top: 20px;
  margin-bottom: 50px;
}
</style>