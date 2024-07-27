<script setup>
import {computed, inject, ref, watch} from "vue";
import CXInputText from "@/components/form_input/CXInputText.vue";
import CXInputNumber from "@/components/form_input/CXInputNumber.vue";
import CXSelect from "@/components/form_input/CXSelect.vue";
import CXTextarea from "@/components/form_input/CXTextarea.vue";
import Button from "primevue/button";
import useBackendPost from "@/composables/useBackendPost.js";
import router from "@/router.js";
import FormStatusText from "@/components/FormStatusText.vue";
import useBackendPatch from "@/composables/useBackendPatch.js";

const props = defineProps(['data']);
const data = computed(() => props.data);
const name = ref("");
const age = ref();
const picture = ref("");
const gender = ref("Male");
const phone = ref();
const description = ref("");

watch(data, (newValue) => {
  name.value = newValue.name;
  age.value = newValue.age;
  picture.value = newValue.picture;
  gender.value = newValue.gender === null ? "Others" : newValue.gender ? "Female" : "Male";
  phone.value = newValue.phone;
  description.value = newValue.description;
});

const genderOptions = [ "Male", "Female", "Others" ];
const postObject = useBackendPost("/tutor/setup");
const patchObject = computed(() => useBackendPatch(`/tutor/${props.data?.id}`));

const post = postObject.post;
const patch = computed(() => patchObject.value.patch);

const status = computed(() => props.data ? patchObject.value.status.value : postObject.status.value);
const loading = computed(() => props.data ? patchObject.value.loading.value : postObject.loading.value);

const $cookies = inject("$cookies");

const submitOnClick = async () => {
  if (props.data) {
    await patch.value({
      name: name.value,
      age: age.value,
      picture: picture.value,
      gender: gender.value === "Others" ? null : gender.value === "Female",
      phone: String(phone.value),
      description: description.value
    });

    if (status.value === "Success") {
      await router.push(`/tutor/${props.data.id}`);
    }
  } else {
    await post({
      name: name.value,
      age: age.value,
      picture: picture.value,
      gender: gender.value === "Others" ? null : gender.value === "Female",
      phone: String(phone.value),
      description: description.value
    });

    if (status.value === "Success") {
      $cookies.set("user_type", "1");
      await router.push("/");
    }
  }
};

</script>

<template>
  <div class="vertical-form-container">
    <CXInputText label="Name" v-model="name" />
    <CXInputNumber label="Age" v-model="age" />
    <CXInputText label="Profile Pic" v-model="picture" />
    <CXSelect label="Gender" :options="genderOptions" v-model="gender" />
    <CXInputNumber label="Handphone" v-model="phone" prefix="+65 " />
    <CXTextarea label="Description" v-model="description" :rows="10" placeholder="Tell users more about yourself" />

    <FormStatusText :status="status" />

    <div style="display: flex; flex-direction: row-reverse">
      <Button id="submit-button" label="Submit" :loading="loading" @click="submitOnClick" />
    </div>
  </div>
</template>

<style scoped>

</style>