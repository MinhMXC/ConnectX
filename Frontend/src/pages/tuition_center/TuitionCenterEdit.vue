<script setup>
import 'primeicons/primeicons.css';
import Back from "@/pages/Back.vue";
import { useRoute } from "vue-router";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
import useBackendGet from "@/composables/useBackendGet.js";
import TuitionCenterSetupForm from "@/components/form/tuition_center/TuitionCenterSetupForm.vue";

const route = useRoute();
const { data } = useBackendGet(`/tuition_center/${route.params.id}`);
</script>

<template>
  <FullScreenCenteredContainer>
    <div style="margin-bottom: 150px; min-width: 600px">
      <Back :navigate="() => $router.push(`/tuition_center/${route.params.id}`)" />
      <h1 id="tuition-center-setup-text">Editing Tuition Center Profile</h1>

      <TuitionCenterSetupForm :data="data" />
    </div>
  </FullScreenCenteredContainer>
</template>

<style scoped>
#tuition-center-setup-text {
  margin-top: 20px;
  margin-bottom: 40px;
}
</style>