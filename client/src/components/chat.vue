<template>
  <div class="container m-5 mx-auto">
    <div class="input-group mb-5">
      <input v-model="username" class="form-control" type="text" />
    </div>

    <div class="input-group">
      <input
        v-model="text"
        @keyup.enter="sendText"
        class="form-control"
        type="text"
      />
      <button class="btn btn-outline-secondary" @click="sendText">send</button>
    </div>
      <Messages :messages="messages" />
  </div>
</template>

<script setup>
/* imports */
import Messages from "./messages.vue";
import { ref, onMounted } from "vue";

/* data */
const username = ref("");
const text = ref("");
const ws = ref();
const messages = ref([]);

/* Mounted */
onMounted(() => {
  console.log("Starting connection to WebSocket Server");
  ws.value = new WebSocket("ws://172.168.10.243:3000/socket");
  ws.value.onmessage = function (event) {
    messages.value.unshift(JSON.parse(event.data));
  };

  ws.value.onopen = function (event) {
    console.log(event);
    console.log("Successfully connected to the golang websocket server...");
  };
});

/* Methods */
function sendText() {
  if (text.value.trim() == "") {
    alert("text can not be empty");
    return;
  }
  var message = {
    username: username.value,
    text: text.value.trim(),
  };
  var jsonMessage = JSON.stringify(message);
  ws.value.send(jsonMessage);
  text.value = "";
}
</script>
