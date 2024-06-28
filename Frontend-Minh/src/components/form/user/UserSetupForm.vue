<script setup>
import Checkbox from "primevue/checkbox";
import {ref} from "vue";
import Button from "primevue/button";
import CXSelect from "@/components/form_input/CXSelect.vue";
import useBackendPost from "@/composables/useBackendPost.js";
import FormStatusText from "@/components/FormStatusText.vue";

const gender = ref("");
const isParent = ref(false);
const genderOptions = [ "Male", "Female", "Others" ]

const { post, status, loading } = useBackendPost("/user/setup");

const submitOnClick = async () => {
  await post({
    gender: gender.value === "Others" ? null : gender.value === "Female",
    is_parent: isParent.value
  });
}
</script>

<template>
  <CXSelect label="Gender" :options="genderOptions" v-model="gender" />

  <Checkbox id="is-parent-checkbox" v-model="isParent" binary />
  <label for="is-parent-checkbox">Are you a parent?</label>

  <br />
  <br />

  <FormStatusText :status="status" />

  <div style="display: flex; flex-direction: row-reverse">
    <Button id="submit-button" label="Submit" :loading="loading" @click="submitOnClick" />
  </div>
</template>

<style scoped>
#is-parent-checkbox {
  margin-right: 10px;
  margin-top: 15px;
}

#submit-button {
  margin-top: 5px;
}
</style>