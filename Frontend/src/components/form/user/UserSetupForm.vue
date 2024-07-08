<script setup>
import Checkbox from "primevue/checkbox";
import {ref} from "vue";
import Button from "primevue/button";
import CXSelect from "@/components/form_input/CXSelect.vue";
import useBackendPost from "@/composables/useBackendPost.js";
import FormStatusText from "@/components/FormStatusText.vue";
import router from "@/router.js";
import CXInputText from "@/components/form_input/CXInputText.vue";

const username = ref("");
const picture = ref("");
const gender = ref("Male");
const isParent = ref(false);
const genderOptions = [ "Male", "Female", "Others" ]

const { post, status, loading } = useBackendPost("/user/setup");

const submitOnClick = async () => {
  await post({
    username: username.value,
    picture: picture.value,
    gender: gender.value === "Others" ? null : gender.value === "Female",
    is_parent: isParent.value
  });

  if (status === "Success") {
    await router.push("/temp");
  }
}
</script>

<template>
  <div class="ctn-ctr-col-no-align" style="gap: 10px">
    <CXInputText label="Username" v-model="username" />
    <CXInputText label="Profile Pic" v-model="picture" />
    <CXSelect label="Gender" :options="genderOptions" v-model="gender" />

    <div>
      <Checkbox id="is-parent-checkbox" v-model="isParent" binary />
      <label for="is-parent-checkbox">Are you a parent?</label>
    </div>

    <FormStatusText :status="status" />

    <div style="display: flex; flex-direction: row-reverse">
      <Button id="submit-button" label="Submit" :loading="loading" @click="submitOnClick" />
    </div>
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