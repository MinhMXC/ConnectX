<script setup>
import 'primeicons/primeicons.css';
import Back from "@/pages/Back.vue";
import UserSetupForm from "@/components/form/user/UserSetupForm.vue";
import { useRoute } from "vue-router";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
import useBackendGet from "@/composables/useBackendGet.js";

const route = useRoute();
const { data } = useBackendGet(`/user/${route.params.id}`);
</script>

<template>
  <FullScreenCenteredContainer>
    <div style="margin-bottom: 150px; min-width: 600px">
      <Back :navigate="() => $router.push(`/user/${route.params.id}`)" />
      <h1 id="user-setup-text">Editing User Profile</h1>

      <UserSetupForm :data="data" />
    </div>
  </FullScreenCenteredContainer>
</template>

<style scoped>
#user-setup-text {
  margin-top: 20px;
  margin-bottom: 40px;
}
</style>