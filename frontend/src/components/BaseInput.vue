<template>
  <div class="input-wrapper">
    <label>{{ label }}</label>
    <input :value="modelValue" @input="$emit('update:modelValue', $event.target.value)" v-bind="$attrs"/>
  </div>
</template>

<script setup>
  defineProps({
    label: {
      type: String,
      default: "",
    },
    modelValue: {
      type: String,
      default: "",
    },
  });
</script>

<style scoped>
  input::-webkit-outer-spin-button,
  input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
  }

  .input-wrapper {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 1rem;

    input {
      border: none;
      border-radius: 0.2rem;
      outline: none;
      height: 2rem;
      padding: 0 0.5rem;
      background-color: rgba(240, 240, 240, 1);
      -webkit-font-smoothing: antialiased;

      .input:hover {
        border: none;
        background-color: rgba(255, 255, 255, 1);
      }

      .input:focus {
        border: none;
        background-color: rgba(255, 255, 255, 1);
      }
    }
  }

</style>