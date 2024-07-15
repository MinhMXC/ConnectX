<script setup>
import 'primeicons/primeicons.css';
import Back from "@/pages/Back.vue";
import { useRoute } from "vue-router";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
import useBackendGet from "@/composables/useBackendGet.js";
import TutorSetupForm from "@/components/form/tutor/TutorSetupForm.vue";

const route = useRoute();
const { data } = useBackendGet(`/tutor/${route.params.id}`);
</script>

<template>
  <FullScreenCenteredContainer>
    <div style="margin-bottom: 150px; min-width: 600px">
      <Back :navigate="() => $router.push(`/tutor/${route.params.id}`)" />
      <h1 id="tutor-setup-text">Editing Tutor Profile</h1>

      <TutorSetupForm :data="data" />
    </div>
  </FullScreenCenteredContainer>
</template>

<style scoped>
#tutor-setup-text {
  margin-top: 20px;
  margin-bottom: 40px;
}
</style>