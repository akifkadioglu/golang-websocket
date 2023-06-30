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
  <MessageInput :isLoading="isLoading" @sendText="sendText" />
  <Messages :messages="messages" />
</template>

<script setup>
/* imports */
import Messages from "./Components/Messages.vue";
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { cryption, auth } from "./../../functions";
import { useSnackbar } from "vue3-snackbar";
import { Status } from "./status";
import "vue3-emoji-picker/css";
import MessageInput from "./Components/MessageInput.vue";

/* data */
const name = ref("");
const picture = ref("");
const ws = ref(WebSocket);
const messages = ref([]);
const route = useRoute();
const snackbar = useSnackbar();
const isLoading = ref(true);

/* Mounted */
onMounted(async () => {
  await checkAuth();
  connectToSocket();
});

/* Methods */
async function checkAuth() {
  isLoading.value = true;
  var token = localStorage.getItem("token");
  await auth.check(token);
  name.value = cryption.parseJwt(token).name;
  picture.value = cryption.parseJwt(token).picture;
  isLoading.value = false;
}

async function connectToSocket() {
  ws.value = new WebSocket(
    "wss://socket-nwnt.onrender.com/socket/" + route.params.name
  );
  ws.value.onmessage = function (event) {
    var data = JSON.parse(event.data);
    switch (data.status) {
      case Status.MESSAGE:
        messages.value.unshift(data);
        break;

      case Status.CONNECT:
        snackbar.add({
          type: "warning",
          text: data.name + " is connected to chat",
          picture: data.picture,
        });
        break;
    }
  };
  var connect = {
    status: Status.CONNECT,
    picture: picture.value,
    name: name.value,
  };
  var jsonConnect = JSON.stringify(connect);
  ws.value.onopen = function (e) {
    ws.value.send(jsonConnect);
  };
}

function sendText(text) {
  if (text.trim() == "") {
    snackbar.add({
      type: "error",
      text: "text can not be empty!",
    });
    return;
  }
  var message = {
    status: Status.MESSAGE,
    picture: picture.value,
    name: name.value,
    text: text.trim(),
  };
  var jsonMessage = JSON.stringify(message);
  ws.value.send(jsonMessage);
}
async function copyURL() {
  await navigator.clipboard.writeText(window.location.href);
  snackbar.add({
    type: "success",
    text: "Link copied to clipboard!",
  });
}
</script>
