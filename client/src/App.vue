<template>
  <router-view v-slot="{ Component, route }">
    <component :is="$route.meta.layout">
      <div class="container mx-auto px-4 py-4" :key="route.name">
        <component :is="Component" />
      </div>
    </component>
  </router-view>
</template>
<script>
export default {};
</script>
