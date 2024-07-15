<script setup>

import router from "@/router.js";
import Back from "@/pages/Back.vue";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
import QualificationForm from "@/components/form/QualificationForm.vue";
import useBackendGet from "@/composables/useBackendGet.js";
import {useRoute} from "vue-router";

const route = useRoute();
const { data } = useBackendGet(`/qualification/${route.params.id}`);
</script>

<template>
  <FullScreenCenteredContainer>
    <div style="margin-bottom: 300px; min-width: 900px;">
      <Back :navigate="() => router.back()" />
      <h1 id="editing-qualification-text">Editing Qualification</h1>

      <QualificationForm :data="data" />
    </div>
  </FullScreenCenteredContainer>
</template>

<style scoped>
#editing-qualification-text {
  margin-top: 20px;
  margin-bottom: 50px;
}
</style>