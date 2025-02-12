<script>
  import { inject } from "vue";
  export default {
    props: ["title"],
    setup() {
      const selectedTab = inject("selectedTab");

      return {
        selectedTab,
      };
    },
  };
</script>

<template>
  <div class="tab-content" v-show="selectedTab === title">
    <slot />
  </div>
</template>