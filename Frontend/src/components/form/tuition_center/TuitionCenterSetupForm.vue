<script setup>
import {computed, ref, watch} from "vue";
import CXInputText from "@/components/form_input/CXInputText.vue";
import CXInputNumber from "@/components/form_input/CXInputNumber.vue";
import CXTextarea from "@/components/form_input/CXTextarea.vue";
import Button from "primevue/button";
import useBackendPost from "@/composables/useBackendPost.js";
import router from "@/router.js";
import FormStatusText from "@/components/FormStatusText.vue";
import useBackendPatch from "@/composables/useBackendPatch.js";

const props = defineProps(['data']);
const data = computed(() => props.data);
const name = ref("");
const picture = ref("");
const phone = ref();
const address = ref("");
const addressLink = ref("");
const description = ref("");
const website = ref("");

watch(data, (newValue) => {
  name.value = newValue.name;
  picture.value = newValue.picture;
  phone.value = newValue.phone;
  address.value = newValue.address;
  addressLink.value = newValue.address_link;
  description.value = newValue.description;
  website.value = newValue.website;
});

const postObject = useBackendPost("/tuition_center/setup");
const patchObject = computed(() => useBackendPatch(`/tuition_center/${props.data?.id}`));

const post = postObject.post;
const patch = computed(() => patchObject.value.patch);

const status = computed(() => props.data ? patchObject.value.status.value : postObject.status.value);
const loading = computed(() => props.data ? patchObject.value.loading.value : postObject.loading.value);

const submitOnclick = async () => {
  if (props.data) {
    await patch.value({
      name: name.value,
      picture: picture.value,
      phone: String(phone.value),
      address: address.value,
      address_link: addressLink.value,
      description: description.value,
      website: website.value
    });

    if (status.value === "Success") {
      await router.push(`/tuition_center/${props.data?.id}`);
    }
  } else {
    await post({
      name: name.value,
      picture: picture.value,
      phone: String(phone.value),
      address: address.value,
      address_link: addressLink.value,
      description: description.value,
      website: website.value
    });

    if (status.value === "Success") {
      await router.push("/");
    }
  }
};

</script>

<template>
  <div class="vertical-form-container">
    <CXInputText label="Name" v-model="name" />
    <CXInputText label="Profile Pic" v-model="picture" />
    <CXInputNumber label="Handphone" v-model="phone" prefix="+65 " />
    <CXInputText label="Address" v-model="address" />
    <CXInputText label="Address Link" v-model="addressLink" />
    <CXTextarea label="Website" v-model="website" :rows="1" placeholder="So users can find out more about you" />
    <CXTextarea label="Description" v-model="description" :rows="10" placeholder="Tell users more about yourself" />

    <FormStatusText :status="status" />

    <div style="display: flex; flex-direction: row-reverse">
      <Button id="submit-button" label="Submit" :loading="loading" @click="submitOnclick" />
    </div>
  </div>
</template>

<style scoped>

</style>