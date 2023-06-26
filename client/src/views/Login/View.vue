<template>
  <div class="rounded-xl text-center space-y-2 py-5 px-2">
    <Divider text="Login" />
    <div class="space-y-5">
      <div v-for="(item, index) in Logins" :key="index" class="px-6 sm:px-0">
        <button
          @click="googleSignIn"
          type="button"
          class="text-white w-36 space-x-3 focus:outline-none font-medium rounded-lg text-sm px-5 py-2.5 inline-flex items-center justify-center"
          :style="'background-color:' + item.color"
        >
          <mdicon :name="item.logo" />
        </button>
      </div>
    </div>
  </div>
</template>
<script>
import Divider from "../../components/Divider.vue";

export default {
  components: {
    Divider,
  },
  data() {
    return {
      Logins: [
        {
          color: "#4285F4",
          logo: "google",
        },
      ],
    };
  },
  methods: {
    googleSignIn() {
      window.open("http://127.0.0.1:3000/google/login","_self");
    },
  },
};
</script>
