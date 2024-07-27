<script setup>
import Checkbox from "primevue/checkbox";
import {computed, inject, ref, watch} from "vue";
import Button from "primevue/button";
import CXSelect from "@/components/form_input/CXSelect.vue";
import useBackendPost from "@/composables/useBackendPost.js";
import FormStatusText from "@/components/FormStatusText.vue";
import router from "@/router.js";
import CXInputText from "@/components/form_input/CXInputText.vue";
import useBackendPatch from "@/composables/useBackendPatch.js";

const props = defineProps(['data']);
const data = computed(() => props.data);
const username = ref("");
const picture = ref("");
const gender = ref("Male");
const isParent = ref(false);
const genderOptions = [ "Male", "Female", "Others" ];

watch(data, (newValue) => {
  username.value = newValue.username;
  picture.value = newValue.picture;
  gender.value = newValue.gender === null ? "Others" : newValue.gender ? "Female" : "Male";
  isParent.value = newValue.is_parent;
});

const postObject = useBackendPost("/user/setup");
const patchObject = computed(() => useBackendPatch(`/user/${props.data?.id}`));

const post = postObject.post;
const patch = computed(() => patchObject.value.patch);

const status = computed(() => props.data ? patchObject.value.status.value : postObject.status.value);
const loading = computed(() => props.data ? patchObject.value.loading.value : postObject.loading.value);

const $cookies = inject("$cookies");

const submitOnClick = async () => {
  if (props.data) {
    await patch.value({
      username: username.value,
      picture: picture.value,
      gender: gender.value === "Others" ? null : gender.value === "Female",
      is_parent: isParent.value
    });

    if (status.value === "Success") {
      await router.push(`/user/${props.data?.id}`);
    }
  } else {
    await post({
      username: username.value,
      picture: picture.value,
      gender: gender.value === "Others" ? null : gender.value === "Female",
      is_parent: isParent.value
    });

    if (status.value === "Success") {
      $cookies.set("user_type", "0");
      await router.push("/");
    }
  }
};
</script>

<template>
  <div class="vertical-form-container">
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