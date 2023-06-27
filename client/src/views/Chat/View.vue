<template>
  <div class="flex justify-between">
    <button
      @click="$router.replace({ name: 'Home' })"
      class="flex space-x-3 items-center hover:underline hover:bg-gray-100 px-3 py-1 rounded-lg"
    >
      <mdicon name="arrow-left" />
      <span>Back to Home</span>
    </button>
    <button
      @click="copyURL"
      class="flex space-x-3 items-center hover:bg-gray-100 px-3 py-1 rounded-lg"
    >
      <mdicon name="share-variant-outline" />
    </button>
  </div>

  <div class="container m-5 mx-auto">
    <div class="md:flex md:items-center mb-6">
      <div class="flex w-full space-x-3">
        <input
          v-model="text"
          @keyup.enter="sendText"
          class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500"
          id="inline-full-name"
          type="text"
          placeholder="Text here"
        />
        <button
          type="button"
          class="text-blue-700 border border-blue-700 hover:bg-blue-700 hover:text-white focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm p-2.5 text-center inline-flex items-center"
          @click="sendText"
        >
          <mdicon name="arrow-right" />
        </button>
      </div>
    </div>
  </div>
  <Messages :messages="messages" />
</template>

<script setup>
/* imports */
import Messages from "./Components/Messages.vue";
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { cryption, auth } from "./../../functions";
import { useSnackbar } from "vue3-snackbar";

/* data */
const name = ref("");
const text = ref("");
const picture = ref("");
const ws = ref();
const messages = ref([]);
const route = useRoute();
const snackbar = useSnackbar();

/* Mounted */
onMounted(() => {
  checkAuth();
  connectToSocket();
});

/* Methods */
async function checkAuth() {
  var token = localStorage.getItem("token");
  await auth.check(token);
  name.value = cryption.parseJwt(token).name;
  picture.value = cryption.parseJwt(token).picture;
}

function connectToSocket() {
  ws.value = new WebSocket(
    "wss://socket-nwnt.onrender.com/socket/" + route.params.name
  );

  ws.value.onmessage = function (event) {
    messages.value.unshift(JSON.parse(event.data));
  };
}

function sendText() {
  if (text.value.trim() == "") {
    alert("text can not be empty");
    return;
  }
  var message = {
    picture: picture.value,
    name: name.value,
    text: text.value.trim(),
  };
  var jsonMessage = JSON.stringify(message);
  ws.value.send(jsonMessage);
  text.value = "";
}
async function copyURL() {
  await navigator.clipboard.writeText(window.location.href);
  snackbar.add({
    type: "success",
    text: "Link copied to clipboard!",
  });
}
</script>
