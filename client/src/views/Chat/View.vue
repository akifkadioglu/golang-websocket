<template>
  <button
    @click="$router.push({ name: 'Home' })"
    class="flex space-x-3 items-center hover:underline"
  >
    <mdicon name="arrow-left" />
    <span>Back to Groups</span>
  </button>
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
          class="text-blue-700 border border-blue-700 hover:bg-blue-700 hover:text-white focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm p-2.5 text-center inline-flex items-center dark:border-blue-500 dark:text-blue-500 dark:hover:text-white dark:focus:ring-blue-800 dark:hover:bg-blue-500"
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
import { useRoute, useRouter } from "vue-router";
import { cryption } from "./../../functions";

/* data */
const name = ref("");
const text = ref("");
const ws = ref();
const messages = ref([]);
const route = useRoute();
const router = useRouter();

/* Mounted */
onMounted(() => {
  checkAuth();
  connectToSocket();
});

/* Methods */
async function checkAuth() {
  var token = localStorage.getItem("token");
  if (!token) {
    localStorage.clear();
    router.replace("/login");
    return;
  }

  var s = await fetch("http://https://socket-nwnt.onrender.com/check", {
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
      authorization: "Bearer " + token,
    },
  });
  if (s.status != 200) {
    localStorage.clear();
    router.replace("/login");
  }
  name.value = cryption.parseJwt(token).name;
}

function connectToSocket() {
  ws.value = new WebSocket("ws://https://socket-nwnt.onrender.com/socket/" + route.params.name);

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
    username: name.value,
    text: text.value.trim(),
  };
  var jsonMessage = JSON.stringify(message);
  ws.value.send(jsonMessage);
  text.value = "";
}
</script>
