<template>
  <div
    v-show="$store.state.imgDialog.isopen"
    class="absolute inset-0 flex items-center justify-center bg-gray-700 bg-opacity-50"
  >
    <div class="p-2 mx-10 bg-white rounded-md shadow-xl">
      <div class="flex items-center space-x-5 justify-between">
        <h3 class="text-2xl">{{ $store.state.imgDialog.title }}</h3>
        <button @click="$store.state.imgDialog.isopen = false">
          <mdicon name="close" />
        </button>
      </div>
      <div class="mt-4">
        <img
          :src="$store.state.imgDialog.src"
          :alt="$store.state.imgDialog.title"
        />
      </div>
    </div>
  </div>
</template>
<script>
export default {};
</script>
<style>
img {
  width: 100%;
  max-width: 720px;
  aspect-ratio: 16 / 16;
}
</style>
