<script setup>
import {ref} from "vue";
import CXInputText from "@/components/form_input/CXInputText.vue";
import CXInputNumber from "@/components/form_input/CXInputNumber.vue";
import CXSelect from "@/components/form_input/CXSelect.vue";
import CXTextarea from "@/components/form_input/CXTextarea.vue";
import Button from "primevue/button";
import useBackendPost from "@/composables/useBackendPost.js";
import router from "@/router.js";
import FormStatusText from "@/components/FormStatusText.vue";

const name = ref("");
const age = ref();
const picture = ref("");
const gender = ref("Male");
const handphone = ref();
const description = ref("");

const genderOptions = [ "Male", "Female", "Others" ];
const { post, status, loading } = useBackendPost("/tutor/setup");

const submitOnClick = async () => {
  await post({
    name: name.value,
    age: age.value,
    picture: picture.value,
    gender: gender.value === "Others" ? null : gender.value === "Female",
    phone: String(handphone.value),
    description: description.value
  })

  if (status === "Success") {
    await router.push("/temp");
  }
}

</script>

<template>
  <div class="vertical-form-container">
    <CXInputText label="Name" v-model="name" />
    <CXInputNumber label="Age" v-model="age" />
    <CXInputText label="Profile Pic" v-model="picture" />
    <CXSelect label="Gender" :options="genderOptions" v-model="gender" />
    <CXInputNumber label="Handphone" v-model="handphone" prefix="+65 " />
    <CXTextarea label="Description" v-model="description" :rows="10" placeholder="Tell users more about yourself" />

    <FormStatusText :status="status" />

    <div style="display: flex; flex-direction: row-reverse">
      <Button id="submit-button" label="Submit" :loading="loading" @click="submitOnClick" />
    </div>
  </div>
</template>

<style scoped>
.vertical-form-container {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>