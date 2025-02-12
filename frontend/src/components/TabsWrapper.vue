<script>
  import { ref, provide } from "vue";
  export default {
    setup(props, { slots }) {
      const tabTitles = ref(slots.default().map((tab) => tab.props.title));
      const selectedTab = ref(tabTitles.value[0]);

      provide("selectedTab", selectedTab);

      return {
        selectedTab,
        tabTitles,
      };
    }
  };
</script>

<template>
  <div class="tabs-wrapper">
    <ul class="tabs-header">
      <li
        v-for="title in tabTitles"
        :key="title"
        :class="{ active: selectedTab === title }"
        @click="selectedTab = title"
      >
        {{ title }}
      </li>
    </ul>
    <slot />
  </div>
</template>

<style scoped>
.tabs {
  display: flex;
  width: 100%;
  margin: 0 auto;
}

.tabs-header {
  display: flex;
  list-style: none;
  padding: 0;
  margin-bottom: 0.2rem;
}

.tabs-header li {
  padding: 0.5rem 1rem;
  background-color: #f0f0f0;
  color: black;
  border-bottom: none;
  border-radius: 0.2rem 0.2rem 0 0;
  cursor: pointer;
  transition: 0.4s all ease-out;
}

.tabs-header li.active {
  background-color: #fff;
  color: #333;
}
</style>