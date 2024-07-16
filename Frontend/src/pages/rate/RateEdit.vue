<script setup>

import router from "@/router.js";
import Back from "@/pages/Back.vue";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
import useBackendGet from "@/composables/useBackendGet.js";
import {useRoute} from "vue-router";
import RateForm from "@/components/form/RateForm.vue";

const route = useRoute();
const { data } = useBackendGet(`/rate/${route.params.id}`);
</script>

<template>
  <FullScreenCenteredContainer>
    <div style="margin-bottom: 300px; min-width: 900px;">
      <Back :navigate="() => router.back()" />
      <h1 id="editing-rate-text">Editing Rate</h1>

      <RateForm :data="data" />
    </div>
  </FullScreenCenteredContainer>
</template>

<style scoped>
#editing-rate-text {
  margin-top: 20px;
  margin-bottom: 50px;
}
</style>