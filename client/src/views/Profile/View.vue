<template>
  <div class="text-center font-bold text-2xl my-5">
    {{ user.email }}
  </div>
  <div class="flex space-x-5 items-center">
    <div>
      <img class="rounded" :src="user.picture" alt="Default avatar" />
    </div>
    <div class="flex w-full space-x-3 items-center">
      <input
        v-model="user.name"
        @keyup.enter="changeName"
        class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500"
        id="inline-full-name"
        type="text"
        placeholder="Name"
      />
      <button
        v-if="!isLoading"
        type="button"
        class="text-blue-700 border border-blue-700 hover:bg-blue-700 hover:text-white focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm p-2.5 text-center inline-flex items-center"
        @click="changeName"
      >
        <mdicon name="check" />
      </button>
      <div v-else role="status">
        <svg
          aria-hidden="true"
          class="inline w-9 h-9 mr-2 text-gray-200 animate-spin fill-gray-600"
          viewBox="0 0 100 101"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
            fill="currentColor"
          />
          <path
            d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
            fill="currentFill"
          />
        </svg>
      </div>
    </div>
  </div>
  <div class="flex justify-center">
    <button @click="logout" class="flex">
      <mdicon name="logout" />
      <span>Logout</span>
    </button>
  </div>
</template>

<script setup>
/* imports */
import { useRouter } from "vue-router";
import { cryption, auth } from "./../../functions";
import { ref, onMounted } from "vue";
import { useSnackbar } from "vue3-snackbar";

/* data */
const router = useRouter();
const user = ref({ picture: "", email: "", name: "" });
const token = ref("");
const isLoading = ref(false);
const snackbar = useSnackbar();

onMounted(() => {
  token.value = localStorage.getItem("token");
  confirmationsFromToken();
});

/* methods */
async function confirmationsFromToken() {
  await auth.check(token.value);
  user.value.name = cryption.parseJwt(token.value).name;
  user.value.picture = cryption.parseJwt(token.value).picture;
  user.value.email = cryption.parseJwt(token.value).email;
}

function logout() {
  localStorage.clear();
  router.replace("/login");
}

async function changeName() {
  isLoading.value = true;
  const requestOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      authorization: "Bearer " + token.value,
    },
    body: JSON.stringify({ name: user.value.name }),
  };
  await fetch("https://socket-nwnt.onrender.com/change-name", requestOptions)
    .then(async (result) => {
      const data = await result.json();
      snackbar.add({
        type: "success",
        text: data.message,
      });
      localStorage.setItem("token", data.token);
    })
    .catch(async (err) => {
      console.log(err);
    });

  isLoading.value = false;
}
</script>
