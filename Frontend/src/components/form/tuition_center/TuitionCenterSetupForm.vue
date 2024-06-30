<script setup>
import {ref} from "vue";
import CXInputText from "@/components/form_input/CXInputText.vue";
import CXInputNumber from "@/components/form_input/CXInputNumber.vue";
import CXTextarea from "@/components/form_input/CXTextarea.vue";
import Button from "primevue/button";
import useBackendPost from "@/composables/useBackendPost.js";
import router from "@/router.js";
import FormStatusText from "@/components/FormStatusText.vue";

const name = ref("");
const handphone = ref();
const address = ref("");
const addressLink = ref("");
const description = ref("");
const website = ref("");

const { post, status, loading } = useBackendPost("/tuition_center/setup");

const submitOnclick = async () => {
  await post({
    name: name.value,
    handphone: handphone.value,
    address: address.value,
    address_link: addressLink.value,
    description: description.value,
    website: website.value
  })

  if (status === "Success") {
    await router.push("/temp");
  }
}

</script>

<template>
  <div class="vertical-form-container">
    <CXInputText label="Name" v-model="name" />
    <CXInputNumber label="Handphone" v-model="handphone" prefix="+65 " />
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
.vertical-form-container {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>